package golang

import (
	"fmt"
	"log"
	"net"
	"os"
	"testing"

	"github.com/prasek/go-grpc-info/grpcinfo"
	"github.com/prasek/protoer/proto"
	protoer "github.com/prasek/protoer/proto/golang"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const port = 6000

func TestMain(m *testing.M) {
	proto.SetProtoer(protoer.NewProtoer(nil))

	go startGrpc()

	code := m.Run()
	os.Exit(code)
}

func TestOptions(t *testing.T) {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure())
	require.Nil(t, err)
	client := NewTestServiceClient(conn)
	resp, err := client.Simple(context.Background(), &TestRequest{Foo: "abc", Bar: "123"})
	require.Nil(t, err)
	require.Equal(t, "abc", resp.Foo)
	require.Equal(t, "123", resp.Bar)
	require.Equal(t, true, resp.Sopt1)
	require.Equal(t, true, resp.Mopt1)
	require.Equal(t, CustomOption{Name: "test123", Value: 55}, *resp.Mopt2)
}

type TestService struct{}

func (s *TestService) Simple(ctx context.Context, req *TestRequest) (*TestResponse, error) {
	out := &TestResponse{Foo: req.Foo, Bar: req.Bar}

	mi := ToMethodInfo(ctx)
	if mi == nil {
		return out, nil
	}

	out.Sopt1 = mi.Service().GetBoolExtension(E_Sopt1, false)
	out.Mopt1 = mi.Method().GetBoolExtension(E_Mopt1, false)
	mopt2, err := mi.Method().GetExtension(E_Mopt2)
	out.Mopt2, _ = mopt2.(*CustomOption)

	if err != nil {
		return nil, err
	}
	sopt, err := mi.Service().GetExtension(E_Sopt1)
	if err != nil {
		return nil, err
	}
	if out.Sopt1 != *(sopt.(*bool)) {
		return nil, fmt.Errorf("genextension mismatch")
	}
	return out, nil
}

func UnaryInterceptor(reg grpcinfo.Registry) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		mi := reg.GetMethodInfo(info.FullMethod)
		ctx = WithMethodInfo(ctx, mi)
		return handler(ctx, req)
	}
}

type ctxKey struct{}

func WithMethodInfo(ctx context.Context, mi grpcinfo.MethodInfo) context.Context {
	return context.WithValue(ctx, ctxKey{}, mi)
}

func ToMethodInfo(ctx context.Context) grpcinfo.MethodInfo {
	if mi, ok := ctx.Value(ctxKey{}).(grpcinfo.MethodInfo); ok {
		return mi
	}
	return nil
}

func startGrpc() {
	reg := grpcinfo.NewRegistry()
	server := grpc.NewServer(grpc.UnaryInterceptor(UnaryInterceptor(reg)))
	RegisterTestServiceServer(server, &TestService{})

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Printf("%v", err)
		os.Exit(1)
	}

	//after services are registered
	err = reg.Load(server)
	if err != nil {
		log.Printf("%v", err)
		os.Exit(1)
	}

	server.Serve(l)
}

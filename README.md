# service and method info for grpc-go interceptors
[![Build Status](https://travis-ci.org/prasek/go-grpc-info.svg?branch=master)](https://travis-ci.org/prasek/go-grpc-info/branches)

## Getting started

```go
import (
	"github.com/prasek/protoer/proto"
	"github.com/prasek/protoer/proto/gogo"
	"github.com/prasek/go-grpc-info/grpcinfo"
)

func main() {
	proto.SetProtoer(gogo.NewProtoer(nil))

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

type TestService struct{}

func (s *TestService) Simple(ctx context.Context, req *TestRequest) (*TestResponse, error) {
	out := &TestResponse{Foo: req.Foo, Bar: req.Bar}

	mi := ToMethodInfo(ctx)
	if mi == nil {
		return out, nil
	}

	out.Sopt1 = mi.Service().GetBoolExtension(E_Sopt1, false)
	out.Mopt1 = mi.Method().GetBoolExtension(E_Mopt1, false)
	mopt2, _ := mi.Method().GetExtension(E_Mopt2)
	out.Mopt2, _ = mopt2.(*CustomOption)
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
```

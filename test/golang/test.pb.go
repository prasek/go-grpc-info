// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

/*
Package golang is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	TestRequest
	TestResponse
	CustomOption
*/
package golang

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TestRequest struct {
	Foo string `protobuf:"bytes,1,opt,name=foo" json:"foo,omitempty"`
	Bar string `protobuf:"bytes,2,opt,name=bar" json:"bar,omitempty"`
}

func (m *TestRequest) Reset()                    { *m = TestRequest{} }
func (m *TestRequest) String() string            { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()               {}
func (*TestRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TestRequest) GetFoo() string {
	if m != nil {
		return m.Foo
	}
	return ""
}

func (m *TestRequest) GetBar() string {
	if m != nil {
		return m.Bar
	}
	return ""
}

type TestResponse struct {
	Foo   string        `protobuf:"bytes,1,opt,name=foo" json:"foo,omitempty"`
	Bar   string        `protobuf:"bytes,2,opt,name=bar" json:"bar,omitempty"`
	Sopt1 bool          `protobuf:"varint,3,opt,name=sopt1" json:"sopt1,omitempty"`
	Mopt1 bool          `protobuf:"varint,4,opt,name=mopt1" json:"mopt1,omitempty"`
	Mopt2 *CustomOption `protobuf:"bytes,5,opt,name=mopt2" json:"mopt2,omitempty"`
}

func (m *TestResponse) Reset()                    { *m = TestResponse{} }
func (m *TestResponse) String() string            { return proto.CompactTextString(m) }
func (*TestResponse) ProtoMessage()               {}
func (*TestResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TestResponse) GetFoo() string {
	if m != nil {
		return m.Foo
	}
	return ""
}

func (m *TestResponse) GetBar() string {
	if m != nil {
		return m.Bar
	}
	return ""
}

func (m *TestResponse) GetSopt1() bool {
	if m != nil {
		return m.Sopt1
	}
	return false
}

func (m *TestResponse) GetMopt1() bool {
	if m != nil {
		return m.Mopt1
	}
	return false
}

func (m *TestResponse) GetMopt2() *CustomOption {
	if m != nil {
		return m.Mopt2
	}
	return nil
}

type CustomOption struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value int32  `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
}

func (m *CustomOption) Reset()                    { *m = CustomOption{} }
func (m *CustomOption) String() string            { return proto.CompactTextString(m) }
func (*CustomOption) ProtoMessage()               {}
func (*CustomOption) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CustomOption) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CustomOption) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

var E_Mopt1 = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50059,
	Name:          "test.mopt1",
	Tag:           "varint,50059,opt,name=mopt1",
	Filename:      "test.proto",
}

var E_Mopt2 = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*CustomOption)(nil),
	Field:         50060,
	Name:          "test.mopt2",
	Tag:           "bytes,50060,opt,name=mopt2",
	Filename:      "test.proto",
}

var E_Sopt1 = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.ServiceOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50061,
	Name:          "test.sopt1",
	Tag:           "varint,50061,opt,name=sopt1",
	Filename:      "test.proto",
}

func init() {
	proto.RegisterType((*TestRequest)(nil), "test.TestRequest")
	proto.RegisterType((*TestResponse)(nil), "test.TestResponse")
	proto.RegisterType((*CustomOption)(nil), "test.CustomOption")
	proto.RegisterExtension(E_Mopt1)
	proto.RegisterExtension(E_Mopt2)
	proto.RegisterExtension(E_Sopt1)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TestService service

type TestServiceClient interface {
	Simple(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) Simple(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error) {
	out := new(TestResponse)
	err := grpc.Invoke(ctx, "/test.TestService/Simple", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TestService service

type TestServiceServer interface {
	Simple(context.Context, *TestRequest) (*TestResponse, error)
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_Simple_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Simple(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.TestService/Simple",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Simple(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Simple",
			Handler:    _TestService_Simple_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x89, 0x4d, 0xab, 0x6e, 0x7b, 0xa8, 0xab, 0x87, 0xa5, 0x07, 0x0d, 0x39, 0x05, 0xb1,
	0x09, 0x4d, 0xc1, 0x4a, 0x8f, 0xea, 0x49, 0x10, 0x21, 0xf5, 0xa2, 0xb7, 0x24, 0xdd, 0xa6, 0xc1,
	0x24, 0xb3, 0xee, 0x6e, 0xfa, 0x0a, 0x15, 0xd4, 0xf7, 0xe9, 0xa3, 0x78, 0xf6, 0xe4, 0x63, 0x48,
	0x76, 0x53, 0x08, 0x08, 0xea, 0x6d, 0xe6, 0xcb, 0xfe, 0x33, 0xf3, 0xff, 0x41, 0x48, 0x52, 0x21,
	0x5d, 0xc6, 0x41, 0x02, 0x36, 0xab, 0x7a, 0x60, 0x25, 0x00, 0x49, 0x46, 0x3d, 0xc5, 0xa2, 0x72,
	0xe1, 0xcd, 0xa9, 0x88, 0x79, 0xca, 0x24, 0x70, 0xfd, 0xce, 0x1e, 0xa1, 0xee, 0x3d, 0x15, 0x32,
	0xa0, 0xcf, 0x25, 0x15, 0x12, 0xf7, 0x51, 0x6b, 0x01, 0x40, 0x0c, 0xcb, 0x70, 0xf6, 0x83, 0xaa,
	0xac, 0x48, 0x14, 0x72, 0xb2, 0xa3, 0x49, 0x14, 0x72, 0xfb, 0xc5, 0x40, 0x3d, 0xad, 0x11, 0x0c,
	0x0a, 0x41, 0xff, 0x23, 0xc2, 0x47, 0xa8, 0x2d, 0x80, 0xc9, 0x11, 0x69, 0x59, 0x86, 0xb3, 0x17,
	0xe8, 0xa6, 0xa2, 0xb9, 0xa2, 0xa6, 0xa6, 0xaa, 0xc1, 0x8e, 0xa6, 0x3e, 0x69, 0x5b, 0x86, 0xd3,
	0xf5, 0xb1, 0xab, 0x7c, 0x5d, 0x95, 0x42, 0x42, 0x7e, 0xc7, 0x64, 0x0a, 0x85, 0x7e, 0xe9, 0xdb,
	0x17, 0xa8, 0xd7, 0xc4, 0x18, 0x23, 0xb3, 0x08, 0x73, 0x5a, 0x9f, 0xa2, 0xea, 0x6a, 0xc7, 0x2a,
	0xcc, 0x4a, 0xaa, 0xae, 0x69, 0x07, 0xba, 0xf1, 0x1f, 0xb4, 0xef, 0x19, 0xe5, 0xab, 0x34, 0xa6,
	0xf8, 0x1a, 0x75, 0x66, 0x69, 0xce, 0x32, 0x8a, 0x0f, 0xf4, 0xb6, 0x46, 0x28, 0x03, 0xdc, 0x44,
	0xda, 0xb3, 0x7d, 0xf8, 0xb1, 0x21, 0xc6, 0xe7, 0x86, 0x74, 0xd1, 0x6e, 0xf5, 0x71, 0xe4, 0x8f,
	0xfb, 0x93, 0x81, 0xf9, 0xb5, 0x21, 0xc6, 0xf4, 0xbc, 0x36, 0x85, 0x8f, 0x5d, 0x1d, 0xbf, 0xbb,
	0x8d, 0xdf, 0xbd, 0xa5, 0x72, 0x09, 0x73, 0x7d, 0xac, 0x20, 0xaf, 0xeb, 0x56, 0xc3, 0xf6, 0xf4,
	0xa6, 0xb6, 0xfd, 0xa7, 0xee, 0x4d, 0xe9, 0x7e, 0x0b, 0x66, 0x3a, 0xa9, 0xe3, 0xc6, 0x27, 0x3f,
	0x66, 0xd5, 0x96, 0xb7, 0xc3, 0xde, 0xd7, 0xcd, 0x3f, 0x72, 0x79, 0xf6, 0x78, 0x9a, 0xa4, 0x72,
	0x59, 0x46, 0x6e, 0x0c, 0xb9, 0xc7, 0x78, 0x28, 0xe8, 0x93, 0x97, 0xc0, 0x30, 0xe1, 0x2c, 0x1e,
	0xa6, 0xc5, 0x02, 0xbc, 0x6a, 0xa7, 0x97, 0x40, 0x16, 0x16, 0x49, 0xd4, 0x51, 0x53, 0xc7, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x53, 0x03, 0x1c, 0x36, 0x7a, 0x02, 0x00, 0x00,
}
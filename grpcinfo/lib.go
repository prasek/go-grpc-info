package grpcinfo

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"

	dpb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/prasek/protoer/proto"
	"google.golang.org/grpc"
)

type Registry interface {
	Load(srv *grpc.Server) error
	GetMethodInfo(fqn string) MethodInfo
}

type MethodInfo interface {
	Service() Info
	Method() Info
}

type Info interface {
	GetExtension(ext interface{}) (interface{}, error)
	GetBoolExtension(ext interface{}, ifnotset bool) bool
}

func NewRegistry() Registry {
	return &registry{
		method: make(map[string]*methodInfo),
	}
}

type registry struct {
	method map[string]*methodInfo
}

func (r *registry) Load(svr *grpc.Server) error {

	for name, info := range svr.GetServiceInfo() {
		file, ok := info.Metadata.(string)
		if !ok {
			return fmt.Errorf("Service %q has unexpected metadata. Expecting a string, got %v", name, info.Metadata)
		}
		fd, err := loadFileDescriptorProto(file)
		if err != nil {
			return err
		}

		pkg := fd.GetPackage()

		merge := func(a, b string) string {
			if a == "" {
				return b
			} else {
				return a + "." + b
			}
		}

		for i := range fd.Service {
			svc := fd.Service[i]
			fqn := merge(pkg, svc.GetName())
			for j := range svc.Method {
				m := svc.Method[j]
				fqnMethod := fmt.Sprintf("/%s/%s", fqn, m.GetName())
				r.method[fqnMethod] = &methodInfo{service: &service{proto: svc}, method: &method{proto: m}}
			}
		}
	}
	return nil
}

func (r *registry) GetMethodInfo(fqn string) MethodInfo {
	return r.method[fqn]
}

type methodInfo struct {
	service *service
	method  *method
}

func (mi *methodInfo) Service() Info {
	return mi.service
}
func (mi *methodInfo) Method() Info {
	return mi.method
}

type service struct {
	proto *dpb.ServiceDescriptorProto
}

func (s *service) GetExtension(ext interface{}) (interface{}, error) {
	return proto.GetExtension(s.proto.GetOptions(), ext)
}

func (s *service) GetBoolExtension(ext interface{}, ifnotset bool) bool {
	return proto.GetBoolExtension(s.proto.GetOptions(), ext, ifnotset)
}

type method struct {
	proto *dpb.MethodDescriptorProto
}

func (m *method) GetExtension(ext interface{}) (interface{}, error) {
	return proto.GetExtension(m.proto.GetOptions(), ext)
}

func (m *method) GetBoolExtension(ext interface{}, ifnotset bool) bool {
	return proto.GetBoolExtension(m.proto.GetOptions(), ext, ifnotset)
}

// loadFileDescriptor loads a registered descriptor and decodes it. If the given
// name cannot be loaded but is a known standard name, an alias will be tried by the proto,
// so the standard files can be loaded even if linked against older "known bad"
// versions of packages.
func loadFileDescriptorProto(file string) (*dpb.FileDescriptorProto, error) {
	fdb := proto.FileDescriptor(file)
	if fdb == nil {
		return nil, fmt.Errorf("Missing file descriptor %s.", file)
	}

	fd, err := decodeFileDescriptorProto(file, fdb)
	if err != nil {
		return nil, err
	}

	// the file descriptor may have been laoded with an alias,
	// so we ensure the specified name to ensure it can be linked.
	fd.Name = proto.String(file)

	return fd, nil
}

// decodeFileDescriptorProto decodes the bytes of a registered file descriptor.
// Registered file descriptors are first "proto encoded" (e.g. binary format
// for the descriptor protos) and then gzipped. So this function gunzips and
// then unmarshals into a descriptor proto.
func decodeFileDescriptorProto(element string, fdb []byte) (*dpb.FileDescriptorProto, error) {
	raw, err := decompress(fdb)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress %q descriptor: %v", element, err)
	}
	fd := dpb.FileDescriptorProto{}
	if err := proto.Unmarshal(raw, &fd); err != nil {
		return nil, fmt.Errorf("bad descriptor for %q: %v", element, err)
	}
	return &fd, nil
}

func decompress(b []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("bad gzipped descriptor: %v", err)
	}
	out, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("bad gzipped descriptor: %v", err)
	}
	return out, nil
}

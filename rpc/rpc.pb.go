// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
	AddNetworkRequest
	AddNetworkReply
	DelNetworkRequest
	DelNetworkReply
*/
package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type AddNetworkRequest struct {
	K8S_POD_NAME      string `protobuf:"bytes,1,opt,name=K8S_POD_NAME,json=K8SPODNAME" json:"K8S_POD_NAME,omitempty"`
	K8S_POD_NAMESPACE string `protobuf:"bytes,2,opt,name=K8S_POD_NAMESPACE,json=K8SPODNAMESPACE" json:"K8S_POD_NAMESPACE,omitempty"`
	Netns             string `protobuf:"bytes,3,opt,name=Netns" json:"Netns,omitempty"`
	IfName            string `protobuf:"bytes,4,opt,name=IfName" json:"IfName,omitempty"`
}

func (m *AddNetworkRequest) Reset()                    { *m = AddNetworkRequest{} }
func (m *AddNetworkRequest) String() string            { return proto.CompactTextString(m) }
func (*AddNetworkRequest) ProtoMessage()               {}
func (*AddNetworkRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddNetworkRequest) GetK8S_POD_NAME() string {
	if m != nil {
		return m.K8S_POD_NAME
	}
	return ""
}

func (m *AddNetworkRequest) GetK8S_POD_NAMESPACE() string {
	if m != nil {
		return m.K8S_POD_NAMESPACE
	}
	return ""
}

func (m *AddNetworkRequest) GetNetns() string {
	if m != nil {
		return m.Netns
	}
	return ""
}

func (m *AddNetworkRequest) GetIfName() string {
	if m != nil {
		return m.IfName
	}
	return ""
}

type AddNetworkReply struct {
	Success      bool   `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
	IPv4Addr     string `protobuf:"bytes,2,opt,name=IPv4Addr" json:"IPv4Addr,omitempty"`
	IPv4Subnet   string `protobuf:"bytes,3,opt,name=IPv4Subnet" json:"IPv4Subnet,omitempty"`
	DeviceNumber int32  `protobuf:"varint,4,opt,name=DeviceNumber" json:"DeviceNumber,omitempty"`
}

func (m *AddNetworkReply) Reset()                    { *m = AddNetworkReply{} }
func (m *AddNetworkReply) String() string            { return proto.CompactTextString(m) }
func (*AddNetworkReply) ProtoMessage()               {}
func (*AddNetworkReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddNetworkReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *AddNetworkReply) GetIPv4Addr() string {
	if m != nil {
		return m.IPv4Addr
	}
	return ""
}

func (m *AddNetworkReply) GetIPv4Subnet() string {
	if m != nil {
		return m.IPv4Subnet
	}
	return ""
}

func (m *AddNetworkReply) GetDeviceNumber() int32 {
	if m != nil {
		return m.DeviceNumber
	}
	return 0
}

type DelNetworkRequest struct {
	K8S_POD_NAME      string `protobuf:"bytes,1,opt,name=K8S_POD_NAME,json=K8SPODNAME" json:"K8S_POD_NAME,omitempty"`
	K8S_POD_NAMESPACE string `protobuf:"bytes,2,opt,name=K8S_POD_NAMESPACE,json=K8SPODNAMESPACE" json:"K8S_POD_NAMESPACE,omitempty"`
	IPv4Addr          string `protobuf:"bytes,3,opt,name=IPv4Addr" json:"IPv4Addr,omitempty"`
}

func (m *DelNetworkRequest) Reset()                    { *m = DelNetworkRequest{} }
func (m *DelNetworkRequest) String() string            { return proto.CompactTextString(m) }
func (*DelNetworkRequest) ProtoMessage()               {}
func (*DelNetworkRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DelNetworkRequest) GetK8S_POD_NAME() string {
	if m != nil {
		return m.K8S_POD_NAME
	}
	return ""
}

func (m *DelNetworkRequest) GetK8S_POD_NAMESPACE() string {
	if m != nil {
		return m.K8S_POD_NAMESPACE
	}
	return ""
}

func (m *DelNetworkRequest) GetIPv4Addr() string {
	if m != nil {
		return m.IPv4Addr
	}
	return ""
}

type DelNetworkReply struct {
	Success      bool   `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
	IPv4Addr     string `protobuf:"bytes,2,opt,name=IPv4Addr" json:"IPv4Addr,omitempty"`
	DeviceNumber int32  `protobuf:"varint,3,opt,name=DeviceNumber" json:"DeviceNumber,omitempty"`
}

func (m *DelNetworkReply) Reset()                    { *m = DelNetworkReply{} }
func (m *DelNetworkReply) String() string            { return proto.CompactTextString(m) }
func (*DelNetworkReply) ProtoMessage()               {}
func (*DelNetworkReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DelNetworkReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *DelNetworkReply) GetIPv4Addr() string {
	if m != nil {
		return m.IPv4Addr
	}
	return ""
}

func (m *DelNetworkReply) GetDeviceNumber() int32 {
	if m != nil {
		return m.DeviceNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*AddNetworkRequest)(nil), "rpc.AddNetworkRequest")
	proto.RegisterType((*AddNetworkReply)(nil), "rpc.AddNetworkReply")
	proto.RegisterType((*DelNetworkRequest)(nil), "rpc.DelNetworkRequest")
	proto.RegisterType((*DelNetworkReply)(nil), "rpc.DelNetworkReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CNIBackend service

type CNIBackendClient interface {
	AddNetwork(ctx context.Context, in *AddNetworkRequest, opts ...grpc.CallOption) (*AddNetworkReply, error)
	DelNetwork(ctx context.Context, in *DelNetworkRequest, opts ...grpc.CallOption) (*DelNetworkReply, error)
}

type cNIBackendClient struct {
	cc *grpc.ClientConn
}

func NewCNIBackendClient(cc *grpc.ClientConn) CNIBackendClient {
	return &cNIBackendClient{cc}
}

func (c *cNIBackendClient) AddNetwork(ctx context.Context, in *AddNetworkRequest, opts ...grpc.CallOption) (*AddNetworkReply, error) {
	out := new(AddNetworkReply)
	err := grpc.Invoke(ctx, "/rpc.CNIBackend/AddNetwork", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cNIBackendClient) DelNetwork(ctx context.Context, in *DelNetworkRequest, opts ...grpc.CallOption) (*DelNetworkReply, error) {
	out := new(DelNetworkReply)
	err := grpc.Invoke(ctx, "/rpc.CNIBackend/DelNetwork", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CNIBackend service

type CNIBackendServer interface {
	AddNetwork(context.Context, *AddNetworkRequest) (*AddNetworkReply, error)
	DelNetwork(context.Context, *DelNetworkRequest) (*DelNetworkReply, error)
}

func RegisterCNIBackendServer(s *grpc.Server, srv CNIBackendServer) {
	s.RegisterService(&_CNIBackend_serviceDesc, srv)
}

func _CNIBackend_AddNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CNIBackendServer).AddNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.CNIBackend/AddNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CNIBackendServer).AddNetwork(ctx, req.(*AddNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CNIBackend_DelNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CNIBackendServer).DelNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.CNIBackend/DelNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CNIBackendServer).DelNetwork(ctx, req.(*DelNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CNIBackend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.CNIBackend",
	HandlerType: (*CNIBackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNetwork",
			Handler:    _CNIBackend_AddNetwork_Handler,
		},
		{
			MethodName: "DelNetwork",
			Handler:    _CNIBackend_DelNetwork_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0xa5, 0x56, 0x10, 0x26, 0x24, 0x0d, 0x1b, 0x42, 0x1a, 0x4c, 0x0c, 0xd9, 0x93, 0xf1, 0xd0,
	0x18, 0xf5, 0xc0, 0xc1, 0x4b, 0xa1, 0x24, 0x12, 0x62, 0x69, 0xda, 0x83, 0x47, 0x02, 0xed, 0xa8,
	0x84, 0x85, 0xd6, 0x6d, 0x0b, 0x72, 0xf0, 0x6e, 0xe2, 0xc1, 0x5f, 0x36, 0xbb, 0x56, 0x29, 0xf4,
	0x66, 0xe2, 0x6d, 0xde, 0xdb, 0x99, 0xcc, 0x7b, 0xb3, 0x0f, 0x6a, 0x3c, 0xf2, 0x8d, 0x88, 0x87,
	0x49, 0x48, 0x54, 0x1e, 0xf9, 0xf4, 0x53, 0x81, 0x86, 0x19, 0x04, 0x36, 0x26, 0x9b, 0x90, 0x2f,
	0x5c, 0x7c, 0x49, 0x31, 0x4e, 0x48, 0x07, 0xea, 0xa3, 0xae, 0x37, 0x71, 0xc6, 0xd6, 0xc4, 0x36,
	0xef, 0x07, 0xba, 0xd2, 0x51, 0xce, 0x6b, 0x2e, 0x8c, 0xba, 0x9e, 0x33, 0xb6, 0x04, 0x43, 0x2e,
	0xa0, 0x91, 0xef, 0xf0, 0x1c, 0xb3, 0x3f, 0xd0, 0x8f, 0x64, 0x9b, 0xb6, 0x6b, 0x93, 0x34, 0x69,
	0x42, 0xd9, 0xc6, 0x64, 0x15, 0xeb, 0xaa, 0x7c, 0xff, 0x06, 0xa4, 0x05, 0x95, 0xe1, 0xa3, 0x3d,
	0x5d, 0xa2, 0x7e, 0x2c, 0xe9, 0x0c, 0xd1, 0x0f, 0x05, 0xb4, 0xbc, 0xa2, 0x88, 0x6d, 0x89, 0x0e,
	0x27, 0x5e, 0xea, 0xfb, 0x18, 0xc7, 0x52, 0x4a, 0xd5, 0xfd, 0x81, 0xa4, 0x0d, 0xd5, 0xa1, 0xb3,
	0xbe, 0x31, 0x83, 0x80, 0x67, 0xeb, 0x7f, 0x31, 0x39, 0x03, 0x10, 0xb5, 0x97, 0xce, 0x56, 0x98,
	0x64, 0xcb, 0x73, 0x0c, 0xa1, 0x50, 0xb7, 0x70, 0x3d, 0xf7, 0xd1, 0x4e, 0x97, 0x33, 0xe4, 0x52,
	0x47, 0xd9, 0xdd, 0xe3, 0xe8, 0x1b, 0x34, 0x2c, 0x64, 0xff, 0x7a, 0x9e, 0xbc, 0x05, 0x75, 0xdf,
	0x02, 0x5d, 0x80, 0x96, 0x5f, 0xff, 0xf7, 0x5b, 0x1c, 0x7a, 0x55, 0x8b, 0x5e, 0xaf, 0xde, 0x15,
	0x80, 0xbe, 0x3d, 0xec, 0x4d, 0xfd, 0x05, 0xae, 0x02, 0x72, 0x0b, 0xb0, 0xfb, 0x07, 0xd2, 0x32,
	0x44, 0x72, 0x0a, 0x51, 0x69, 0x37, 0x0b, 0x7c, 0xc4, 0xb6, 0xb4, 0x24, 0xa6, 0x77, 0xca, 0xb3,
	0xe9, 0xc2, 0x25, 0xb3, 0xe9, 0x03, 0x8b, 0xb4, 0xd4, 0xbb, 0x84, 0xd3, 0x79, 0x68, 0x3c, 0x89,
	0x47, 0x7c, 0x9d, 0x2e, 0x23, 0x86, 0xb1, 0xf1, 0x8c, 0x8c, 0x85, 0x9b, 0x90, 0xb3, 0xa0, 0xa7,
	0xdd, 0x89, 0xfa, 0x41, 0xd4, 0x8e, 0xc8, 0xb2, 0xa3, 0xcc, 0x2a, 0x32, 0xd4, 0xd7, 0x5f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x0b, 0x98, 0xa4, 0x09, 0xe1, 0x02, 0x00, 0x00,
}

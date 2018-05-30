// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	PodInformation
	DeviceInformation
	VfInformation
*/
package api

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

type PodInformation struct {
	PodName      string `protobuf:"bytes,1,opt,name=podName" json:"podName,omitempty"`
	PodNamespace string `protobuf:"bytes,2,opt,name=podNamespace" json:"podNamespace,omitempty"`
}

func (m *PodInformation) Reset()                    { *m = PodInformation{} }
func (m *PodInformation) String() string            { return proto.CompactTextString(m) }
func (*PodInformation) ProtoMessage()               {}
func (*PodInformation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PodInformation) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

func (m *PodInformation) GetPodNamespace() string {
	if m != nil {
		return m.PodNamespace
	}
	return ""
}

type DeviceInformation struct {
	VFs []*VfInformation `protobuf:"bytes,1,rep,name=VFs" json:"VFs,omitempty"`
}

func (m *DeviceInformation) Reset()                    { *m = DeviceInformation{} }
func (m *DeviceInformation) String() string            { return proto.CompactTextString(m) }
func (*DeviceInformation) ProtoMessage()               {}
func (*DeviceInformation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DeviceInformation) GetVFs() []*VfInformation {
	if m != nil {
		return m.VFs
	}
	return nil
}

type VfInformation struct {
	Pciaddr string `protobuf:"bytes,3,opt,name=pciaddr" json:"pciaddr,omitempty"`
	Pfname  string `protobuf:"bytes,1,opt,name=pfname" json:"pfname,omitempty"`
	Vfid    int32  `protobuf:"varint,2,opt,name=vfid" json:"vfid,omitempty"`
}

func (m *VfInformation) Reset()                    { *m = VfInformation{} }
func (m *VfInformation) String() string            { return proto.CompactTextString(m) }
func (*VfInformation) ProtoMessage()               {}
func (*VfInformation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *VfInformation) GetPciaddr() string {
	if m != nil {
		return m.Pciaddr
	}
	return ""
}

func (m *VfInformation) GetPfname() string {
	if m != nil {
		return m.Pfname
	}
	return ""
}

func (m *VfInformation) GetVfid() int32 {
	if m != nil {
		return m.Vfid
	}
	return 0
}

func init() {
	proto.RegisterType((*PodInformation)(nil), "api.PodInformation")
	proto.RegisterType((*DeviceInformation)(nil), "api.DeviceInformation")
	proto.RegisterType((*VfInformation)(nil), "api.VfInformation")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SendPodInformation service

type SendPodInformationClient interface {
	SendPodInformation(ctx context.Context, in *PodInformation, opts ...grpc.CallOption) (*DeviceInformation, error)
}

type sendPodInformationClient struct {
	cc *grpc.ClientConn
}

func NewSendPodInformationClient(cc *grpc.ClientConn) SendPodInformationClient {
	return &sendPodInformationClient{cc}
}

func (c *sendPodInformationClient) SendPodInformation(ctx context.Context, in *PodInformation, opts ...grpc.CallOption) (*DeviceInformation, error) {
	out := new(DeviceInformation)
	err := grpc.Invoke(ctx, "/api.SendPodInformation/SendPodInformation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SendPodInformation service

type SendPodInformationServer interface {
	SendPodInformation(context.Context, *PodInformation) (*DeviceInformation, error)
}

func RegisterSendPodInformationServer(s *grpc.Server, srv SendPodInformationServer) {
	s.RegisterService(&_SendPodInformation_serviceDesc, srv)
}

func _SendPodInformation_SendPodInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PodInformation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendPodInformationServer).SendPodInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SendPodInformation/SendPodInformation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendPodInformationServer).SendPodInformation(ctx, req.(*PodInformation))
	}
	return interceptor(ctx, in, info, handler)
}

var _SendPodInformation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.SendPodInformation",
	HandlerType: (*SendPodInformationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendPodInformation",
			Handler:    _SendPodInformation_SendPodInformation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0xf2, 0xe3, 0xe2, 0x0b, 0xc8,
	0x4f, 0xf1, 0xcc, 0x4b, 0xcb, 0x2f, 0xca, 0x4d, 0x2c, 0xc9, 0xcc, 0xcf, 0x13, 0x92, 0xe0, 0x62,
	0x2f, 0xc8, 0x4f, 0xf1, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71,
	0x85, 0x94, 0xb8, 0x78, 0xa0, 0xcc, 0xe2, 0x82, 0xc4, 0xe4, 0x54, 0x09, 0x26, 0xb0, 0x34, 0x8a,
	0x98, 0x92, 0x25, 0x97, 0xa0, 0x4b, 0x6a, 0x59, 0x66, 0x72, 0x2a, 0xb2, 0x91, 0x2a, 0x5c, 0xcc,
	0x61, 0x6e, 0xc5, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x42, 0x7a, 0x20, 0x27, 0x84, 0xa5,
	0x21, 0x29, 0x08, 0x02, 0x49, 0x2b, 0x85, 0x72, 0xf1, 0xa2, 0x88, 0x82, 0x5d, 0x92, 0x9c, 0x99,
	0x98, 0x92, 0x52, 0x24, 0xc1, 0x0c, 0x75, 0x09, 0x84, 0x2b, 0x24, 0xc6, 0xc5, 0x56, 0x90, 0x96,
	0x87, 0x70, 0x22, 0x94, 0x27, 0x24, 0xc4, 0xc5, 0x52, 0x96, 0x96, 0x99, 0x02, 0x76, 0x19, 0x6b,
	0x10, 0x98, 0x6d, 0x14, 0xc9, 0x25, 0x14, 0x9c, 0x9a, 0x97, 0x82, 0xe6, 0x4b, 0x67, 0xac, 0xa2,
	0xc2, 0x60, 0xb7, 0xa1, 0x0a, 0x4a, 0x89, 0x81, 0x05, 0x31, 0x7c, 0xa5, 0xc4, 0x90, 0xc4, 0x06,
	0x0e, 0x48, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x8c, 0xbd, 0x80, 0x55, 0x01, 0x00,
	0x00,
}
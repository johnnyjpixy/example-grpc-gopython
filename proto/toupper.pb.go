// Code generated by protoc-gen-go. DO NOT EDIT.
// source: toupper.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The request message
type UpperRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpperRequest) Reset()         { *m = UpperRequest{} }
func (m *UpperRequest) String() string { return proto.CompactTextString(m) }
func (*UpperRequest) ProtoMessage()    {}
func (*UpperRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9263bae945d71a4, []int{0}
}

func (m *UpperRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpperRequest.Unmarshal(m, b)
}
func (m *UpperRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpperRequest.Marshal(b, m, deterministic)
}
func (m *UpperRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpperRequest.Merge(m, src)
}
func (m *UpperRequest) XXX_Size() int {
	return xxx_messageInfo_UpperRequest.Size(m)
}
func (m *UpperRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpperRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpperRequest proto.InternalMessageInfo

func (m *UpperRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message
type UpperReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpperReply) Reset()         { *m = UpperReply{} }
func (m *UpperReply) String() string { return proto.CompactTextString(m) }
func (*UpperReply) ProtoMessage()    {}
func (*UpperReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9263bae945d71a4, []int{1}
}

func (m *UpperReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpperReply.Unmarshal(m, b)
}
func (m *UpperReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpperReply.Marshal(b, m, deterministic)
}
func (m *UpperReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpperReply.Merge(m, src)
}
func (m *UpperReply) XXX_Size() int {
	return xxx_messageInfo_UpperReply.Size(m)
}
func (m *UpperReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UpperReply.DiscardUnknown(m)
}

var xxx_messageInfo_UpperReply proto.InternalMessageInfo

func (m *UpperReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*UpperRequest)(nil), "proto.UpperRequest")
	proto.RegisterType((*UpperReply)(nil), "proto.UpperReply")
}

func init() {
	proto.RegisterFile("toupper.proto", fileDescriptor_b9263bae945d71a4)
}

var fileDescriptor_b9263bae945d71a4 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0xc9, 0x2f, 0x2d,
	0x28, 0x48, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x4a, 0x5c,
	0x3c, 0xa1, 0x20, 0xd1, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xbc,
	0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49, 0x8d, 0x8b, 0x0b,
	0xaa, 0xa6, 0x20, 0xa7, 0x52, 0x48, 0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x1d, 0xa6,
	0x08, 0xc6, 0x35, 0xb2, 0xe1, 0x62, 0x0f, 0xc9, 0x07, 0xab, 0x14, 0x32, 0xe4, 0x62, 0x85, 0x30,
	0x84, 0x21, 0xd6, 0xe9, 0x21, 0x5b, 0x22, 0x25, 0x88, 0x2a, 0x58, 0x90, 0x53, 0xa9, 0xc4, 0x90,
	0xc4, 0x06, 0x16, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xd6, 0xe6, 0x20, 0xa8, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ToUpperClient is the client API for ToUpper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ToUpperClient interface {
	// Sends a greeting
	Upper(ctx context.Context, in *UpperRequest, opts ...grpc.CallOption) (*UpperReply, error)
}

type toUpperClient struct {
	cc grpc.ClientConnInterface
}

func NewToUpperClient(cc grpc.ClientConnInterface) ToUpperClient {
	return &toUpperClient{cc}
}

func (c *toUpperClient) Upper(ctx context.Context, in *UpperRequest, opts ...grpc.CallOption) (*UpperReply, error) {
	out := new(UpperReply)
	err := c.cc.Invoke(ctx, "/proto.ToUpper/Upper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToUpperServer is the server API for ToUpper service.
type ToUpperServer interface {
	// Sends a greeting
	Upper(context.Context, *UpperRequest) (*UpperReply, error)
}

// UnimplementedToUpperServer can be embedded to have forward compatible implementations.
type UnimplementedToUpperServer struct {
}

func (*UnimplementedToUpperServer) Upper(ctx context.Context, req *UpperRequest) (*UpperReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upper not implemented")
}

func RegisterToUpperServer(s *grpc.Server, srv ToUpperServer) {
	s.RegisterService(&_ToUpper_serviceDesc, srv)
}

func _ToUpper_Upper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpperRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToUpperServer).Upper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ToUpper/Upper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToUpperServer).Upper(ctx, req.(*UpperRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ToUpper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ToUpper",
	HandlerType: (*ToUpperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upper",
			Handler:    _ToUpper_Upper_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "toupper.proto",
}

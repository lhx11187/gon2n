// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/gon2n.proto

package gon2n

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

type SupernodeManagerCreateArgs struct {
	ListenPort           int64    `protobuf:"varint,1,opt,name=ListenPort,proto3" json:"ListenPort,omitempty"`
	ManagementPort       int64    `protobuf:"varint,2,opt,name=ManagementPort,proto3" json:"ManagementPort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SupernodeManagerCreateArgs) Reset()         { *m = SupernodeManagerCreateArgs{} }
func (m *SupernodeManagerCreateArgs) String() string { return proto.CompactTextString(m) }
func (*SupernodeManagerCreateArgs) ProtoMessage()    {}
func (*SupernodeManagerCreateArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2514c45d8f9ec34, []int{0}
}

func (m *SupernodeManagerCreateArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SupernodeManagerCreateArgs.Unmarshal(m, b)
}
func (m *SupernodeManagerCreateArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SupernodeManagerCreateArgs.Marshal(b, m, deterministic)
}
func (m *SupernodeManagerCreateArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupernodeManagerCreateArgs.Merge(m, src)
}
func (m *SupernodeManagerCreateArgs) XXX_Size() int {
	return xxx_messageInfo_SupernodeManagerCreateArgs.Size(m)
}
func (m *SupernodeManagerCreateArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_SupernodeManagerCreateArgs.DiscardUnknown(m)
}

var xxx_messageInfo_SupernodeManagerCreateArgs proto.InternalMessageInfo

func (m *SupernodeManagerCreateArgs) GetListenPort() int64 {
	if m != nil {
		return m.ListenPort
	}
	return 0
}

func (m *SupernodeManagerCreateArgs) GetManagementPort() int64 {
	if m != nil {
		return m.ManagementPort
	}
	return 0
}

type SupernodeManagerCreateReply struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SupernodeManagerCreateReply) Reset()         { *m = SupernodeManagerCreateReply{} }
func (m *SupernodeManagerCreateReply) String() string { return proto.CompactTextString(m) }
func (*SupernodeManagerCreateReply) ProtoMessage()    {}
func (*SupernodeManagerCreateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2514c45d8f9ec34, []int{1}
}

func (m *SupernodeManagerCreateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SupernodeManagerCreateReply.Unmarshal(m, b)
}
func (m *SupernodeManagerCreateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SupernodeManagerCreateReply.Marshal(b, m, deterministic)
}
func (m *SupernodeManagerCreateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupernodeManagerCreateReply.Merge(m, src)
}
func (m *SupernodeManagerCreateReply) XXX_Size() int {
	return xxx_messageInfo_SupernodeManagerCreateReply.Size(m)
}
func (m *SupernodeManagerCreateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SupernodeManagerCreateReply.DiscardUnknown(m)
}

var xxx_messageInfo_SupernodeManagerCreateReply proto.InternalMessageInfo

func (m *SupernodeManagerCreateReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*SupernodeManagerCreateArgs)(nil), "gon2n.SupernodeManagerCreateArgs")
	proto.RegisterType((*SupernodeManagerCreateReply)(nil), "gon2n.SupernodeManagerCreateReply")
}

func init() { proto.RegisterFile("proto/gon2n.proto", fileDescriptor_e2514c45d8f9ec34) }

var fileDescriptor_e2514c45d8f9ec34 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xcf, 0xcf, 0x33, 0xca, 0xd3, 0x03, 0xb3, 0x85, 0x58, 0xc1, 0x1c, 0xa5, 0x14,
	0x2e, 0xa9, 0xe0, 0xd2, 0x82, 0xd4, 0xa2, 0xbc, 0xfc, 0x94, 0x54, 0xdf, 0xc4, 0xbc, 0xc4, 0xf4,
	0xd4, 0x22, 0xe7, 0xa2, 0xd4, 0xc4, 0x92, 0x54, 0xc7, 0xa2, 0xf4, 0x62, 0x21, 0x39, 0x2e, 0x2e,
	0x9f, 0xcc, 0xe2, 0x92, 0xd4, 0xbc, 0x80, 0xfc, 0xa2, 0x12, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xe6,
	0x20, 0x24, 0x11, 0x21, 0x35, 0x2e, 0x3e, 0x88, 0xa6, 0xdc, 0xd4, 0xbc, 0x12, 0xb0, 0x1a, 0x26,
	0xb0, 0x1a, 0x34, 0x51, 0x25, 0x5d, 0x2e, 0x69, 0xec, 0xb6, 0x04, 0xa5, 0x16, 0xe4, 0x54, 0x0a,
	0xf1, 0x71, 0x31, 0x79, 0xa6, 0x80, 0x8d, 0xe7, 0x0c, 0x62, 0xf2, 0x4c, 0x31, 0x4a, 0xe5, 0x12,
	0x40, 0x57, 0x2e, 0x14, 0xc8, 0xc5, 0x06, 0xd1, 0x22, 0xa4, 0xa8, 0x07, 0xf1, 0x07, 0x6e, 0x77,
	0x4b, 0x29, 0xe1, 0x55, 0x02, 0xb6, 0x54, 0x89, 0x21, 0x89, 0x0d, 0x1c, 0x12, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x3c, 0x88, 0x69, 0xd6, 0x1e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SupernodeManagerClient is the client API for SupernodeManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SupernodeManagerClient interface {
	Create(ctx context.Context, in *SupernodeManagerCreateArgs, opts ...grpc.CallOption) (*SupernodeManagerCreateReply, error)
}

type supernodeManagerClient struct {
	cc *grpc.ClientConn
}

func NewSupernodeManagerClient(cc *grpc.ClientConn) SupernodeManagerClient {
	return &supernodeManagerClient{cc}
}

func (c *supernodeManagerClient) Create(ctx context.Context, in *SupernodeManagerCreateArgs, opts ...grpc.CallOption) (*SupernodeManagerCreateReply, error) {
	out := new(SupernodeManagerCreateReply)
	err := c.cc.Invoke(ctx, "/gon2n.SupernodeManager/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SupernodeManagerServer is the server API for SupernodeManager service.
type SupernodeManagerServer interface {
	Create(context.Context, *SupernodeManagerCreateArgs) (*SupernodeManagerCreateReply, error)
}

// UnimplementedSupernodeManagerServer can be embedded to have forward compatible implementations.
type UnimplementedSupernodeManagerServer struct {
}

func (*UnimplementedSupernodeManagerServer) Create(ctx context.Context, req *SupernodeManagerCreateArgs) (*SupernodeManagerCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterSupernodeManagerServer(s *grpc.Server, srv SupernodeManagerServer) {
	s.RegisterService(&_SupernodeManager_serviceDesc, srv)
}

func _SupernodeManager_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupernodeManagerCreateArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupernodeManagerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gon2n.SupernodeManager/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupernodeManagerServer).Create(ctx, req.(*SupernodeManagerCreateArgs))
	}
	return interceptor(ctx, in, info, handler)
}

var _SupernodeManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gon2n.SupernodeManager",
	HandlerType: (*SupernodeManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SupernodeManager_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/gon2n.proto",
}

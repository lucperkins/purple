// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kv.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Location struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_2216fe83c9c12408, []int{0}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type Value struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_2216fe83c9c12408, []int{1}
}

func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

func (m *Value) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type GetResponse struct {
	Value                *Value   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2216fe83c9c12408, []int{2}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetValue() *Value {
	if m != nil {
		return m.Value
	}
	return nil
}

type PutRequest struct {
	Location             *Location `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Value                *Value    `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PutRequest) Reset()         { *m = PutRequest{} }
func (m *PutRequest) String() string { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()    {}
func (*PutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2216fe83c9c12408, []int{3}
}

func (m *PutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRequest.Unmarshal(m, b)
}
func (m *PutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRequest.Marshal(b, m, deterministic)
}
func (m *PutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRequest.Merge(m, src)
}
func (m *PutRequest) XXX_Size() int {
	return xxx_messageInfo_PutRequest.Size(m)
}
func (m *PutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutRequest proto.InternalMessageInfo

func (m *PutRequest) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *PutRequest) GetValue() *Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Location)(nil), "proto.Location")
	proto.RegisterType((*Value)(nil), "proto.Value")
	proto.RegisterType((*GetResponse)(nil), "proto.GetResponse")
	proto.RegisterType((*PutRequest)(nil), "proto.PutRequest")
}

func init() { proto.RegisterFile("kv.proto", fileDescriptor_2216fe83c9c12408) }

var fileDescriptor_2216fe83c9c12408 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8e, 0x3f, 0x4f, 0x84, 0x40,
	0x10, 0xc5, 0xc3, 0x11, 0x4e, 0x9c, 0x23, 0x51, 0xa7, 0xba, 0x10, 0x0b, 0xdd, 0x42, 0x8d, 0x26,
	0x97, 0x78, 0x7e, 0x05, 0xcd, 0x15, 0x5a, 0x5c, 0xb6, 0xa0, 0xb3, 0x40, 0x32, 0x85, 0x01, 0x76,
	0x50, 0x66, 0x49, 0x28, 0xfd, 0xe6, 0x86, 0x65, 0x11, 0xff, 0xc4, 0x6a, 0x77, 0xe6, 0xbd, 0xf7,
	0x9b, 0x07, 0x71, 0xd9, 0x6d, 0x9a, 0x77, 0x16, 0xc6, 0xc8, 0x3d, 0x69, 0x52, 0x70, 0x5d, 0xb3,
	0x19, 0x97, 0xea, 0x14, 0xe2, 0x27, 0x2e, 0x72, 0x79, 0x65, 0x83, 0xc7, 0x10, 0x96, 0xd4, 0xaf,
	0x83, 0xb3, 0xe0, 0xea, 0x50, 0x0f, 0x5f, 0x75, 0x0e, 0x51, 0x96, 0x57, 0x96, 0x70, 0x0d, 0x07,
	0x05, 0x1b, 0x21, 0x23, 0x4e, 0x4e, 0xf4, 0x34, 0xaa, 0x5b, 0x58, 0xed, 0x48, 0x34, 0xb5, 0x0d,
	0x9b, 0x96, 0x50, 0x41, 0xd4, 0x0d, 0x09, 0x67, 0x5b, 0x6d, 0x93, 0xf1, 0xcc, 0xc6, 0x51, 0xf4,
	0x28, 0xa9, 0x67, 0x80, 0xbd, 0x15, 0x4d, 0x6f, 0x96, 0x5a, 0xc1, 0x1b, 0x88, 0x2b, 0xdf, 0xc0,
	0x87, 0x8e, 0x7c, 0x68, 0x2a, 0xa6, 0xbf, 0x0c, 0x33, 0x7e, 0xf1, 0x2f, 0x7e, 0xfb, 0x11, 0xc0,
	0xe2, 0x31, 0xc3, 0x6b, 0x08, 0x77, 0x24, 0xf8, 0x1b, 0x96, 0xa2, 0x5f, 0x7c, 0x6f, 0x7d, 0x01,
	0xe1, 0xde, 0x0a, 0x9e, 0x78, 0x69, 0x6e, 0x97, 0x4e, 0x17, 0x1e, 0xea, 0x46, 0x7a, 0xbc, 0x84,
	0xe5, 0x3d, 0x55, 0x24, 0xf4, 0x17, 0xfb, 0xc3, 0xf8, 0xb2, 0x74, 0xc3, 0xdd, 0x67, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x64, 0xf6, 0xf7, 0x6c, 0x7e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KVClient is the client API for KV service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KVClient interface {
	Get(ctx context.Context, in *Location, opts ...grpc.CallOption) (*GetResponse, error)
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*Empty, error)
	Delete(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Empty, error)
}

type kVClient struct {
	cc *grpc.ClientConn
}

func NewKVClient(cc *grpc.ClientConn) KVClient {
	return &kVClient{cc}
}

func (c *kVClient) Get(ctx context.Context, in *Location, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/proto.KV/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.KV/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) Delete(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.KV/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KVServer is the server API for KV service.
type KVServer interface {
	Get(context.Context, *Location) (*GetResponse, error)
	Put(context.Context, *PutRequest) (*Empty, error)
	Delete(context.Context, *Location) (*Empty, error)
}

func RegisterKVServer(s *grpc.Server, srv KVServer) {
	s.RegisterService(&_KV_serviceDesc, srv)
}

func _KV_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Location)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KV/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Get(ctx, req.(*Location))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KV/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Location)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KV/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Delete(ctx, req.(*Location))
	}
	return interceptor(ctx, in, info, handler)
}

var _KV_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.KV",
	HandlerType: (*KVServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _KV_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _KV_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _KV_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kv.proto",
}

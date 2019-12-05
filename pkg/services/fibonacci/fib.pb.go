// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fib.proto

package fibonacci

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

type FibNthRequest struct {
	N                    int32    `protobuf:"varint,1,opt,name=N,proto3" json:"N,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibNthRequest) Reset()         { *m = FibNthRequest{} }
func (m *FibNthRequest) String() string { return proto.CompactTextString(m) }
func (*FibNthRequest) ProtoMessage()    {}
func (*FibNthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2db00d511e02f9dc, []int{0}
}

func (m *FibNthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibNthRequest.Unmarshal(m, b)
}
func (m *FibNthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibNthRequest.Marshal(b, m, deterministic)
}
func (m *FibNthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibNthRequest.Merge(m, src)
}
func (m *FibNthRequest) XXX_Size() int {
	return xxx_messageInfo_FibNthRequest.Size(m)
}
func (m *FibNthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FibNthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FibNthRequest proto.InternalMessageInfo

func (m *FibNthRequest) GetN() int32 {
	if m != nil {
		return m.N
	}
	return 0
}

type FibNthResponse struct {
	Value                int32    `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibNthResponse) Reset()         { *m = FibNthResponse{} }
func (m *FibNthResponse) String() string { return proto.CompactTextString(m) }
func (*FibNthResponse) ProtoMessage()    {}
func (*FibNthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2db00d511e02f9dc, []int{1}
}

func (m *FibNthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibNthResponse.Unmarshal(m, b)
}
func (m *FibNthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibNthResponse.Marshal(b, m, deterministic)
}
func (m *FibNthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibNthResponse.Merge(m, src)
}
func (m *FibNthResponse) XXX_Size() int {
	return xxx_messageInfo_FibNthResponse.Size(m)
}
func (m *FibNthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FibNthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FibNthResponse proto.InternalMessageInfo

func (m *FibNthResponse) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type FibSeqUpperLimit struct {
	UpperBound           int32    `protobuf:"varint,1,opt,name=UpperBound,proto3" json:"UpperBound,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibSeqUpperLimit) Reset()         { *m = FibSeqUpperLimit{} }
func (m *FibSeqUpperLimit) String() string { return proto.CompactTextString(m) }
func (*FibSeqUpperLimit) ProtoMessage()    {}
func (*FibSeqUpperLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_2db00d511e02f9dc, []int{2}
}

func (m *FibSeqUpperLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibSeqUpperLimit.Unmarshal(m, b)
}
func (m *FibSeqUpperLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibSeqUpperLimit.Marshal(b, m, deterministic)
}
func (m *FibSeqUpperLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibSeqUpperLimit.Merge(m, src)
}
func (m *FibSeqUpperLimit) XXX_Size() int {
	return xxx_messageInfo_FibSeqUpperLimit.Size(m)
}
func (m *FibSeqUpperLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_FibSeqUpperLimit.DiscardUnknown(m)
}

var xxx_messageInfo_FibSeqUpperLimit proto.InternalMessageInfo

func (m *FibSeqUpperLimit) GetUpperBound() int32 {
	if m != nil {
		return m.UpperBound
	}
	return 0
}

type FibSeqResponse struct {
	Seq                  []int32  `protobuf:"varint,1,rep,packed,name=Seq,proto3" json:"Seq,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibSeqResponse) Reset()         { *m = FibSeqResponse{} }
func (m *FibSeqResponse) String() string { return proto.CompactTextString(m) }
func (*FibSeqResponse) ProtoMessage()    {}
func (*FibSeqResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2db00d511e02f9dc, []int{3}
}

func (m *FibSeqResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibSeqResponse.Unmarshal(m, b)
}
func (m *FibSeqResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibSeqResponse.Marshal(b, m, deterministic)
}
func (m *FibSeqResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibSeqResponse.Merge(m, src)
}
func (m *FibSeqResponse) XXX_Size() int {
	return xxx_messageInfo_FibSeqResponse.Size(m)
}
func (m *FibSeqResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FibSeqResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FibSeqResponse proto.InternalMessageInfo

func (m *FibSeqResponse) GetSeq() []int32 {
	if m != nil {
		return m.Seq
	}
	return nil
}

func init() {
	proto.RegisterType((*FibNthRequest)(nil), "fibonacci.FibNthRequest")
	proto.RegisterType((*FibNthResponse)(nil), "fibonacci.FibNthResponse")
	proto.RegisterType((*FibSeqUpperLimit)(nil), "fibonacci.FibSeqUpperLimit")
	proto.RegisterType((*FibSeqResponse)(nil), "fibonacci.FibSeqResponse")
}

func init() { proto.RegisterFile("fib.proto", fileDescriptor_2db00d511e02f9dc) }

var fileDescriptor_2db00d511e02f9dc = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0xcb, 0x4c, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x02, 0x31, 0xf3, 0xf3, 0x12, 0x93, 0x93, 0x33, 0x95, 0x64,
	0xb9, 0x78, 0xdd, 0x32, 0x93, 0xfc, 0x4a, 0x32, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84,
	0x78, 0xb8, 0x18, 0xfd, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x18, 0xfd, 0x94, 0xd4, 0xb8,
	0xf8, 0x60, 0xd2, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x22, 0x5c, 0xac, 0x61, 0x89, 0x39,
	0xa5, 0xa9, 0x50, 0x35, 0x10, 0x8e, 0x92, 0x11, 0x97, 0x80, 0x5b, 0x66, 0x52, 0x70, 0x6a, 0x61,
	0x68, 0x41, 0x41, 0x6a, 0x91, 0x4f, 0x66, 0x6e, 0x66, 0x89, 0x90, 0x1c, 0x17, 0x17, 0x98, 0xe7,
	0x94, 0x5f, 0x9a, 0x97, 0x02, 0x55, 0x8e, 0x24, 0xa2, 0xa4, 0x04, 0x36, 0x3b, 0x38, 0xb5, 0x10,
	0x6e, 0xb6, 0x00, 0x17, 0x73, 0x70, 0x6a, 0xa1, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0x6b, 0x10, 0x88,
	0x69, 0x34, 0x8d, 0x91, 0x8b, 0xd3, 0x0d, 0xe6, 0x58, 0x21, 0x27, 0x2e, 0x4e, 0xf7, 0xd4, 0x12,
	0xbf, 0x92, 0x0c, 0xb7, 0xcc, 0x24, 0x21, 0x09, 0x3d, 0xb8, 0x2f, 0xf4, 0x50, 0xbc, 0x20, 0x25,
	0x89, 0x45, 0x06, 0x62, 0x83, 0x12, 0x83, 0x90, 0x2b, 0xd8, 0x0c, 0x88, 0xc5, 0x42, 0xd2, 0xa8,
	0x2a, 0x51, 0xdc, 0x8f, 0x6e, 0x0c, 0x92, 0x43, 0x95, 0x18, 0x92, 0xd8, 0xc0, 0x21, 0x69, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x05, 0x77, 0x49, 0x8f, 0x56, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FibonacciClient is the client API for Fibonacci service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FibonacciClient interface {
	GetNthFib(ctx context.Context, in *FibNthRequest, opts ...grpc.CallOption) (*FibNthResponse, error)
	GetFibSeq(ctx context.Context, in *FibSeqUpperLimit, opts ...grpc.CallOption) (*FibSeqResponse, error)
}

type fibonacciClient struct {
	cc *grpc.ClientConn
}

func NewFibonacciClient(cc *grpc.ClientConn) FibonacciClient {
	return &fibonacciClient{cc}
}

func (c *fibonacciClient) GetNthFib(ctx context.Context, in *FibNthRequest, opts ...grpc.CallOption) (*FibNthResponse, error) {
	out := new(FibNthResponse)
	err := c.cc.Invoke(ctx, "/fibonacci.Fibonacci/GetNthFib", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fibonacciClient) GetFibSeq(ctx context.Context, in *FibSeqUpperLimit, opts ...grpc.CallOption) (*FibSeqResponse, error) {
	out := new(FibSeqResponse)
	err := c.cc.Invoke(ctx, "/fibonacci.Fibonacci/GetFibSeq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FibonacciServer is the server API for Fibonacci service.
type FibonacciServer interface {
	GetNthFib(context.Context, *FibNthRequest) (*FibNthResponse, error)
	GetFibSeq(context.Context, *FibSeqUpperLimit) (*FibSeqResponse, error)
}

// UnimplementedFibonacciServer can be embedded to have forward compatible implementations.
type UnimplementedFibonacciServer struct {
}

func (*UnimplementedFibonacciServer) GetNthFib(ctx context.Context, req *FibNthRequest) (*FibNthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNthFib not implemented")
}
func (*UnimplementedFibonacciServer) GetFibSeq(ctx context.Context, req *FibSeqUpperLimit) (*FibSeqResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFibSeq not implemented")
}

func RegisterFibonacciServer(s *grpc.Server, srv FibonacciServer) {
	s.RegisterService(&_Fibonacci_serviceDesc, srv)
}

func _Fibonacci_GetNthFib_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FibNthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FibonacciServer).GetNthFib(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fibonacci.Fibonacci/GetNthFib",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FibonacciServer).GetNthFib(ctx, req.(*FibNthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fibonacci_GetFibSeq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FibSeqUpperLimit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FibonacciServer).GetFibSeq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fibonacci.Fibonacci/GetFibSeq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FibonacciServer).GetFibSeq(ctx, req.(*FibSeqUpperLimit))
	}
	return interceptor(ctx, in, info, handler)
}

var _Fibonacci_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fibonacci.Fibonacci",
	HandlerType: (*FibonacciServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNthFib",
			Handler:    _Fibonacci_GetNthFib_Handler,
		},
		{
			MethodName: "GetFibSeq",
			Handler:    _Fibonacci_GetFibSeq_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fib.proto",
}
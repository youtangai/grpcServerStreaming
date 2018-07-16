// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloService.proto

package helloService

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

type HelloRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_helloService_f816705c59bbc8ff, []int{0}
}
func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (dst *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(dst, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type HelloResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_helloService_f816705c59bbc8ff, []int{1}
}
func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (dst *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(dst, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "HelloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServerStreamingServiceClient is the client API for HelloServerStreamingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServerStreamingServiceClient interface {
	Greet(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (HelloServerStreamingService_GreetClient, error)
}

type helloServerStreamingServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServerStreamingServiceClient(cc *grpc.ClientConn) HelloServerStreamingServiceClient {
	return &helloServerStreamingServiceClient{cc}
}

func (c *helloServerStreamingServiceClient) Greet(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (HelloServerStreamingService_GreetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HelloServerStreamingService_serviceDesc.Streams[0], "/HelloServerStreamingService/Greet", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServerStreamingServiceGreetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HelloServerStreamingService_GreetClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type helloServerStreamingServiceGreetClient struct {
	grpc.ClientStream
}

func (x *helloServerStreamingServiceGreetClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServerStreamingServiceServer is the server API for HelloServerStreamingService service.
type HelloServerStreamingServiceServer interface {
	Greet(*HelloRequest, HelloServerStreamingService_GreetServer) error
}

func RegisterHelloServerStreamingServiceServer(s *grpc.Server, srv HelloServerStreamingServiceServer) {
	s.RegisterService(&_HelloServerStreamingService_serviceDesc, srv)
}

func _HelloServerStreamingService_Greet_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServerStreamingServiceServer).Greet(m, &helloServerStreamingServiceGreetServer{stream})
}

type HelloServerStreamingService_GreetServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type helloServerStreamingServiceGreetServer struct {
	grpc.ServerStream
}

func (x *helloServerStreamingServiceGreetServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _HelloServerStreamingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "HelloServerStreamingService",
	HandlerType: (*HelloServerStreamingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Greet",
			Handler:       _HelloServerStreamingService_Greet_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "helloService.proto",
}

func init() { proto.RegisterFile("helloService.proto", fileDescriptor_helloService_f816705c59bbc8ff) }

var fileDescriptor_helloService_f816705c59bbc8ff = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0x48, 0xcd, 0xc9,
	0xc9, 0x0f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2,
	0xe0, 0xe2, 0xf1, 0x00, 0x89, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49, 0x70, 0xb1,
	0xe7, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8,
	0x4a, 0x9a, 0x5c, 0xbc, 0x50, 0x95, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0xb8, 0x95, 0x1a, 0x79,
	0x72, 0x49, 0x7b, 0xc0, 0xac, 0x4a, 0x2d, 0x0a, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xcd, 0xcc, 0x4b,
	0x87, 0xda, 0x2c, 0xa4, 0xc5, 0xc5, 0xea, 0x5e, 0x94, 0x9a, 0x5a, 0x22, 0xc4, 0xab, 0x87, 0x6c,
	0xb7, 0x14, 0x9f, 0x1e, 0x8a, 0x05, 0x4a, 0x0c, 0x06, 0x8c, 0x49, 0x6c, 0x60, 0x67, 0x1a, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xa4, 0xa6, 0x47, 0xbc, 0x00, 0x00, 0x00,
}

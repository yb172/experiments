// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wordgen/internets.proto

/*
Package wordgen is a generated protocol buffer package.

It is generated from these files:
	wordgen/internets.proto
	wordgen/number.proto
	wordgen/word.proto

It has these top-level messages:
	GetWordReq
	GetWordResp
	GenerateNumberReq
	GenerateNumberResp
	GenerateWordReq
	GenerateWordResp
*/
package wordgen

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

// GetWordReq contains data reuqired or helpful for word generation
type GetWordReq struct {
}

func (m *GetWordReq) Reset()                    { *m = GetWordReq{} }
func (m *GetWordReq) String() string            { return proto.CompactTextString(m) }
func (*GetWordReq) ProtoMessage()               {}
func (*GetWordReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// GetWordResp contains reults of word generation
type GetWordResp struct {
	Word string `protobuf:"bytes,1,opt,name=word" json:"word,omitempty"`
}

func (m *GetWordResp) Reset()                    { *m = GetWordResp{} }
func (m *GetWordResp) String() string            { return proto.CompactTextString(m) }
func (*GetWordResp) ProtoMessage()               {}
func (*GetWordResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetWordResp) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

func init() {
	proto.RegisterType((*GetWordReq)(nil), "wordgen.GetWordReq")
	proto.RegisterType((*GetWordResp)(nil), "wordgen.GetWordResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for InternetsWordGenerator service

type InternetsWordGeneratorClient interface {
	// GetWord searches the internet and returns random word
	GetWord(ctx context.Context, in *GetWordReq, opts ...grpc.CallOption) (*GetWordResp, error)
}

type internetsWordGeneratorClient struct {
	cc *grpc.ClientConn
}

func NewInternetsWordGeneratorClient(cc *grpc.ClientConn) InternetsWordGeneratorClient {
	return &internetsWordGeneratorClient{cc}
}

func (c *internetsWordGeneratorClient) GetWord(ctx context.Context, in *GetWordReq, opts ...grpc.CallOption) (*GetWordResp, error) {
	out := new(GetWordResp)
	err := grpc.Invoke(ctx, "/wordgen.InternetsWordGenerator/GetWord", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for InternetsWordGenerator service

type InternetsWordGeneratorServer interface {
	// GetWord searches the internet and returns random word
	GetWord(context.Context, *GetWordReq) (*GetWordResp, error)
}

func RegisterInternetsWordGeneratorServer(s *grpc.Server, srv InternetsWordGeneratorServer) {
	s.RegisterService(&_InternetsWordGenerator_serviceDesc, srv)
}

func _InternetsWordGenerator_GetWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternetsWordGeneratorServer).GetWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wordgen.InternetsWordGenerator/GetWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternetsWordGeneratorServer).GetWord(ctx, req.(*GetWordReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _InternetsWordGenerator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wordgen.InternetsWordGenerator",
	HandlerType: (*InternetsWordGeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWord",
			Handler:    _InternetsWordGenerator_GetWord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wordgen/internets.proto",
}

func init() { proto.RegisterFile("wordgen/internets.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0xcf, 0x2f, 0x4a,
	0x49, 0x4f, 0xcd, 0xd3, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0x2d, 0x29, 0xd6, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x4a, 0x28, 0xf1, 0x70, 0x71, 0xb9, 0xa7, 0x96, 0x84, 0xe7,
	0x17, 0xa5, 0x04, 0xa5, 0x16, 0x2a, 0x29, 0x72, 0x71, 0xc3, 0x79, 0xc5, 0x05, 0x42, 0x42, 0x5c,
	0x2c, 0x20, 0x75, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x51, 0x00, 0x97, 0x98,
	0x27, 0xcc, 0x30, 0x90, 0x42, 0xf7, 0xd4, 0xbc, 0xd4, 0xa2, 0xc4, 0x92, 0xfc, 0x22, 0x21, 0x33,
	0x2e, 0x76, 0xa8, 0x66, 0x21, 0x61, 0x3d, 0xa8, 0xf9, 0x7a, 0x08, 0xc3, 0xa5, 0x44, 0x30, 0x05,
	0x8b, 0x0b, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0x4e, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x10,
	0xb1, 0x5f, 0x47, 0xad, 0x00, 0x00, 0x00,
}

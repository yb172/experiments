// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wordgen/word.proto

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

// GenerateWordReq contains data reuqired or helpful for word generation
type GenerateWordReq struct {
}

func (m *GenerateWordReq) Reset()                    { *m = GenerateWordReq{} }
func (m *GenerateWordReq) String() string            { return proto.CompactTextString(m) }
func (*GenerateWordReq) ProtoMessage()               {}
func (*GenerateWordReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// GenerateWordResp contains reults of word generation
type GenerateWordResp struct {
	Word string `protobuf:"bytes,1,opt,name=word" json:"word,omitempty"`
}

func (m *GenerateWordResp) Reset()                    { *m = GenerateWordResp{} }
func (m *GenerateWordResp) String() string            { return proto.CompactTextString(m) }
func (*GenerateWordResp) ProtoMessage()               {}
func (*GenerateWordResp) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *GenerateWordResp) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

func init() {
	proto.RegisterType((*GenerateWordReq)(nil), "wordgen.GenerateWordReq")
	proto.RegisterType((*GenerateWordResp)(nil), "wordgen.GenerateWordResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for WordGenerator service

type WordGeneratorClient interface {
	// GenerateWord searches the internet and returns random word
	GenerateWord(ctx context.Context, in *GenerateWordReq, opts ...grpc.CallOption) (*GenerateWordResp, error)
}

type wordGeneratorClient struct {
	cc *grpc.ClientConn
}

func NewWordGeneratorClient(cc *grpc.ClientConn) WordGeneratorClient {
	return &wordGeneratorClient{cc}
}

func (c *wordGeneratorClient) GenerateWord(ctx context.Context, in *GenerateWordReq, opts ...grpc.CallOption) (*GenerateWordResp, error) {
	out := new(GenerateWordResp)
	err := grpc.Invoke(ctx, "/wordgen.WordGenerator/GenerateWord", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WordGenerator service

type WordGeneratorServer interface {
	// GenerateWord searches the internet and returns random word
	GenerateWord(context.Context, *GenerateWordReq) (*GenerateWordResp, error)
}

func RegisterWordGeneratorServer(s *grpc.Server, srv WordGeneratorServer) {
	s.RegisterService(&_WordGenerator_serviceDesc, srv)
}

func _WordGenerator_GenerateWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateWordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordGeneratorServer).GenerateWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wordgen.WordGenerator/GenerateWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordGeneratorServer).GenerateWord(ctx, req.(*GenerateWordReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _WordGenerator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wordgen.WordGenerator",
	HandlerType: (*WordGeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateWord",
			Handler:    _WordGenerator_GenerateWord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wordgen/word.proto",
}

func init() { proto.RegisterFile("wordgen/word.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0xcf, 0x2f, 0x4a,
	0x49, 0x4f, 0xcd, 0xd3, 0x07, 0xd1, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0x31,
	0x25, 0x41, 0x2e, 0x7e, 0xf7, 0xd4, 0xbc, 0xd4, 0xa2, 0xc4, 0x92, 0xd4, 0xf0, 0xfc, 0xa2, 0x94,
	0xa0, 0xd4, 0x42, 0x25, 0x35, 0x2e, 0x01, 0x54, 0xa1, 0xe2, 0x02, 0x21, 0x21, 0x2e, 0x16, 0x90,
	0x0e, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0xdb, 0x28, 0x8c, 0x8b, 0x17, 0x24, 0x0f,
	0x55, 0x9b, 0x5f, 0x24, 0xe4, 0xca, 0xc5, 0x83, 0xac, 0x51, 0x48, 0x42, 0x0f, 0x6a, 0x8b, 0x1e,
	0x9a, 0x15, 0x52, 0x92, 0x38, 0x64, 0x8a, 0x0b, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0x4e, 0x34, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x70, 0x79, 0xe4, 0xf1, 0xb8, 0x00, 0x00, 0x00,
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/dspa5.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/dspa5.proto

It has these top-level messages:
	Announcement
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Announcement_Level int32

const (
	// no chime (other systems may not display text)
	Announcement_NOTSET Announcement_Level = 0
	// no chime
	Announcement_DEBUG Announcement_Level = 10
	// attention chime
	Announcement_INFO Announcement_Level = 20
	// warning chime
	Announcement_WARNING Announcement_Level = 30
	// alarm chime either end
	Announcement_ERROR Announcement_Level = 40
	// alarm chime either end, message repeated twice
	Announcement_CRITICAL Announcement_Level = 50
)

var Announcement_Level_name = map[int32]string{
	0:  "NOTSET",
	10: "DEBUG",
	20: "INFO",
	30: "WARNING",
	40: "ERROR",
	50: "CRITICAL",
}
var Announcement_Level_value = map[string]int32{
	"NOTSET":   0,
	"DEBUG":    10,
	"INFO":     20,
	"WARNING":  30,
	"ERROR":    40,
	"CRITICAL": 50,
}

func (x Announcement_Level) String() string {
	return proto1.EnumName(Announcement_Level_name, int32(x))
}
func (Announcement_Level) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// similar to a python log handler
type Announcement struct {
	Message string             `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Level   Announcement_Level `protobuf:"varint,2,opt,name=level,enum=proto.Announcement_Level" json:"level,omitempty"`
}

func (m *Announcement) Reset()                    { *m = Announcement{} }
func (m *Announcement) String() string            { return proto1.CompactTextString(m) }
func (*Announcement) ProtoMessage()               {}
func (*Announcement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Announcement) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Announcement) GetLevel() Announcement_Level {
	if m != nil {
		return m.Level
	}
	return Announcement_NOTSET
}

func init() {
	proto1.RegisterType((*Announcement)(nil), "proto.Announcement")
	proto1.RegisterEnum("proto.Announcement_Level", Announcement_Level_name, Announcement_Level_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Dspa06 service

type Dspa06Client interface {
	// system will transform message into fragments and stream them back as they
	// are announced.
	Speak(ctx context.Context, in *Announcement, opts ...grpc.CallOption) (Dspa06_SpeakClient, error)
}

type dspa06Client struct {
	cc *grpc.ClientConn
}

func NewDspa06Client(cc *grpc.ClientConn) Dspa06Client {
	return &dspa06Client{cc}
}

func (c *dspa06Client) Speak(ctx context.Context, in *Announcement, opts ...grpc.CallOption) (Dspa06_SpeakClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Dspa06_serviceDesc.Streams[0], c.cc, "/proto.Dspa06/Speak", opts...)
	if err != nil {
		return nil, err
	}
	x := &dspa06SpeakClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Dspa06_SpeakClient interface {
	Recv() (*Announcement, error)
	grpc.ClientStream
}

type dspa06SpeakClient struct {
	grpc.ClientStream
}

func (x *dspa06SpeakClient) Recv() (*Announcement, error) {
	m := new(Announcement)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Dspa06 service

type Dspa06Server interface {
	// system will transform message into fragments and stream them back as they
	// are announced.
	Speak(*Announcement, Dspa06_SpeakServer) error
}

func RegisterDspa06Server(s *grpc.Server, srv Dspa06Server) {
	s.RegisterService(&_Dspa06_serviceDesc, srv)
}

func _Dspa06_Speak_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Announcement)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(Dspa06Server).Speak(m, &dspa06SpeakServer{stream})
}

type Dspa06_SpeakServer interface {
	Send(*Announcement) error
	grpc.ServerStream
}

type dspa06SpeakServer struct {
	grpc.ServerStream
}

func (x *dspa06SpeakServer) Send(m *Announcement) error {
	return x.ServerStream.SendMsg(m)
}

var _Dspa06_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Dspa06",
	HandlerType: (*Dspa06Server)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Speak",
			Handler:       _Dspa06_Speak_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/dspa5.proto",
}

func init() { proto1.RegisterFile("proto/dspa5.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x29, 0x2e, 0x48, 0x34, 0xd5, 0x03, 0xb3, 0x85, 0x58, 0xc1, 0x94, 0xd2, 0x4a,
	0x46, 0x2e, 0x1e, 0xc7, 0xbc, 0xbc, 0xfc, 0xd2, 0xbc, 0xe4, 0xd4, 0xdc, 0xd4, 0xbc, 0x12, 0x21,
	0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x18, 0x57, 0x48, 0x9f, 0x8b, 0x35, 0x27, 0xb5, 0x2c, 0x35, 0x47, 0x82, 0x49, 0x81, 0x51,
	0x83, 0xcf, 0x48, 0x12, 0x62, 0x90, 0x1e, 0xb2, 0x6e, 0x3d, 0x1f, 0x90, 0x82, 0x20, 0x88, 0x3a,
	0x25, 0x3f, 0x2e, 0x56, 0x30, 0x5f, 0x88, 0x8b, 0x8b, 0xcd, 0xcf, 0x3f, 0x24, 0xd8, 0x35, 0x44,
	0x80, 0x41, 0x88, 0x93, 0x8b, 0xd5, 0xc5, 0xd5, 0x29, 0xd4, 0x5d, 0x80, 0x4b, 0x88, 0x83, 0x8b,
	0xc5, 0xd3, 0xcf, 0xcd, 0x5f, 0x40, 0x44, 0x88, 0x9b, 0x8b, 0x3d, 0xdc, 0x31, 0xc8, 0xcf, 0xd3,
	0xcf, 0x5d, 0x40, 0x0e, 0xa4, 0xc2, 0x35, 0x28, 0xc8, 0x3f, 0x48, 0x40, 0x43, 0x88, 0x87, 0x8b,
	0xc3, 0x39, 0xc8, 0x33, 0xc4, 0xd3, 0xd9, 0xd1, 0x47, 0xc0, 0xc8, 0xc8, 0x96, 0x8b, 0xcd, 0xa5,
	0xb8, 0x20, 0xd1, 0xc0, 0x4c, 0xc8, 0x98, 0x8b, 0x35, 0xb8, 0x20, 0x35, 0x31, 0x5b, 0x48, 0x18,
	0x8b, 0x23, 0xa4, 0xb0, 0x09, 0x1a, 0x30, 0x26, 0xb1, 0x81, 0x45, 0x8d, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xe7, 0x69, 0x8f, 0xb3, 0x0d, 0x01, 0x00, 0x00,
}

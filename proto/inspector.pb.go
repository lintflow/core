// Code generated by protoc-gen-go.
// source: proto/inspector.proto
// DO NOT EDIT!

/*
Package lintflow_core is a generated protocol buffer package.

It is generated from these files:
	proto/inspector.proto
	proto/lookupd.proto
	proto/reporter.proto
	proto/resourcer.proto
	proto/validator.proto

It has these top-level messages:
	Task
	Progress
	Service
	ListRequest
	ListResponse
	RegisterRequest
	RegisterResponse
	Problem
	ConfigRequest
	Resource
	ValidationTask
	ValidateProgress
*/
package lintflow_core

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
const _ = proto.ProtoPackageIsVersion1

type Task struct {
	Validators *Task_Args `protobuf:"bytes,1,opt,name=validators" json:"validators,omitempty"`
	Reporters  *Task_Args `protobuf:"bytes,2,opt,name=reporters" json:"reporters,omitempty"`
	Resourcers *Task_Args `protobuf:"bytes,3,opt,name=resourcers" json:"resourcers,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Task) GetValidators() *Task_Args {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *Task) GetReporters() *Task_Args {
	if m != nil {
		return m.Reporters
	}
	return nil
}

func (m *Task) GetResourcers() *Task_Args {
	if m != nil {
		return m.Resourcers
	}
	return nil
}

type Task_Args struct {
	Service *Service `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	Config  []byte   `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (m *Task_Args) Reset()                    { *m = Task_Args{} }
func (m *Task_Args) String() string            { return proto.CompactTextString(m) }
func (*Task_Args) ProtoMessage()               {}
func (*Task_Args) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Task_Args) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type Progress struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Total   int64  `protobuf:"varint,2,opt,name=total" json:"total,omitempty"`
	Current int64  `protobuf:"varint,3,opt,name=current" json:"current,omitempty"`
}

func (m *Progress) Reset()                    { *m = Progress{} }
func (m *Progress) String() string            { return proto.CompactTextString(m) }
func (*Progress) ProtoMessage()               {}
func (*Progress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Task)(nil), "lintflow.core.Task")
	proto.RegisterType((*Task_Args)(nil), "lintflow.core.Task.Args")
	proto.RegisterType((*Progress)(nil), "lintflow.core.Progress")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for InspectorService service

type InspectorServiceClient interface {
	// run task to inspectation
	Inspect(ctx context.Context, in *Task, opts ...grpc.CallOption) (InspectorService_InspectClient, error)
	// return avialability services for usage
	Services(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type inspectorServiceClient struct {
	cc *grpc.ClientConn
}

func NewInspectorServiceClient(cc *grpc.ClientConn) InspectorServiceClient {
	return &inspectorServiceClient{cc}
}

func (c *inspectorServiceClient) Inspect(ctx context.Context, in *Task, opts ...grpc.CallOption) (InspectorService_InspectClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_InspectorService_serviceDesc.Streams[0], c.cc, "/lintflow.core.InspectorService/Inspect", opts...)
	if err != nil {
		return nil, err
	}
	x := &inspectorServiceInspectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InspectorService_InspectClient interface {
	Recv() (*Progress, error)
	grpc.ClientStream
}

type inspectorServiceInspectClient struct {
	grpc.ClientStream
}

func (x *inspectorServiceInspectClient) Recv() (*Progress, error) {
	m := new(Progress)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *inspectorServiceClient) Services(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/lintflow.core.InspectorService/Services", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for InspectorService service

type InspectorServiceServer interface {
	// run task to inspectation
	Inspect(*Task, InspectorService_InspectServer) error
	// return avialability services for usage
	Services(context.Context, *ListRequest) (*ListResponse, error)
}

func RegisterInspectorServiceServer(s *grpc.Server, srv InspectorServiceServer) {
	s.RegisterService(&_InspectorService_serviceDesc, srv)
}

func _InspectorService_Inspect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Task)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InspectorServiceServer).Inspect(m, &inspectorServiceInspectServer{stream})
}

type InspectorService_InspectServer interface {
	Send(*Progress) error
	grpc.ServerStream
}

type inspectorServiceInspectServer struct {
	grpc.ServerStream
}

func (x *inspectorServiceInspectServer) Send(m *Progress) error {
	return x.ServerStream.SendMsg(m)
}

func _InspectorService_Services_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(InspectorServiceServer).Services(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _InspectorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lintflow.core.InspectorService",
	HandlerType: (*InspectorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Services",
			Handler:    _InspectorService_Services_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Inspect",
			Handler:       _InspectorService_Inspect_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x91, 0xcb, 0x4a, 0xc3, 0x50,
	0x10, 0x86, 0x9b, 0xb6, 0xf6, 0x32, 0x5e, 0x90, 0xa9, 0xd6, 0x12, 0x37, 0x92, 0x95, 0xab, 0x58,
	0x2a, 0x88, 0xe0, 0xca, 0x85, 0x0b, 0xc5, 0x45, 0x39, 0xfa, 0x02, 0x31, 0x3d, 0x2d, 0x87, 0x86,
	0x4c, 0x9c, 0x39, 0xa9, 0x2f, 0xe0, 0x33, 0xf8, 0xbc, 0xe6, 0x4a, 0x35, 0x14, 0xba, 0x3b, 0xff,
	0xcc, 0xf7, 0xcf, 0xfc, 0x87, 0x81, 0xf3, 0x84, 0xc9, 0xd2, 0x8d, 0x89, 0x25, 0xd1, 0xa1, 0x25,
	0xf6, 0x0b, 0x8d, 0xc7, 0x91, 0x89, 0xed, 0x32, 0xa2, 0x2f, 0x3f, 0x24, 0xd6, 0xee, 0xa8, 0xa4,
	0x22, 0xa2, 0x75, 0x9a, 0x2c, 0x4a, 0xc6, 0xfb, 0x6e, 0x43, 0xf7, 0x3d, 0x90, 0x35, 0xde, 0x03,
	0x6c, 0x82, 0xc8, 0x2c, 0x82, 0xcc, 0x2f, 0x13, 0xe7, 0xca, 0xb9, 0x3e, 0x9c, 0x4d, 0xfc, 0x7f,
	0x13, 0xfc, 0x1c, 0xf4, 0x1f, 0x79, 0x25, 0xea, 0x0f, 0x8b, 0x77, 0x30, 0x64, 0x9d, 0x10, 0x5b,
	0x9d, 0x19, 0xdb, 0x7b, 0x8c, 0x5b, 0x34, 0xdf, 0xc8, 0x5a, 0x28, 0xe5, 0x30, 0x37, 0x76, 0xf6,
	0x6d, 0xdc, 0xb2, 0xee, 0x1c, 0xba, 0x79, 0x0d, 0xa7, 0xd0, 0x17, 0xcd, 0x1b, 0x13, 0xea, 0x2a,
	0xf0, 0xb8, 0x61, 0x7f, 0x2b, 0xbb, 0xaa, 0xc6, 0x70, 0x0c, 0xbd, 0x90, 0xe2, 0xa5, 0x59, 0x15,
	0x41, 0x8f, 0x54, 0xa5, 0xbc, 0x17, 0x18, 0xcc, 0x99, 0x56, 0xd9, 0x0e, 0xc1, 0x13, 0x68, 0x9b,
	0x45, 0x31, 0x70, 0xa8, 0xb2, 0x17, 0x9e, 0xc1, 0x81, 0x25, 0x1b, 0x44, 0x85, 0xa5, 0xa3, 0x4a,
	0x81, 0x13, 0xe8, 0x87, 0x29, 0xb3, 0x8e, 0x6d, 0x11, 0xbd, 0xa3, 0x6a, 0x39, 0xfb, 0x71, 0xe0,
	0xf4, 0xb9, 0x3e, 0x45, 0x95, 0x00, 0x1f, 0xa0, 0x5f, 0xd5, 0x70, 0xb4, 0xe3, 0x8f, 0xee, 0x45,
	0xa3, 0x58, 0xa7, 0xf1, 0x5a, 0x53, 0x07, 0x9f, 0x60, 0x50, 0xcd, 0x11, 0x74, 0x1b, 0xe0, 0xab,
	0x11, 0xab, 0xf4, 0x67, 0xaa, 0xc5, 0xba, 0x97, 0x3b, 0x7b, 0x92, 0x50, 0x2c, 0xda, 0x6b, 0x7d,
	0xf4, 0x8a, 0x93, 0xdf, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x96, 0x00, 0xde, 0x2f, 0x02,
	0x00, 0x00,
}

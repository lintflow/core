// Code generated by protoc-gen-go.
// source: proto/validator.proto
// DO NOT EDIT!

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

type ValidationTask struct {
	Config    []byte               `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	Reporter  *ValidationTask_Args `protobuf:"bytes,2,opt,name=reporter" json:"reporter,omitempty"`
	Resourcer *ValidationTask_Args `protobuf:"bytes,3,opt,name=resourcer" json:"resourcer,omitempty"`
}

func (m *ValidationTask) Reset()                    { *m = ValidationTask{} }
func (m *ValidationTask) String() string            { return proto.CompactTextString(m) }
func (*ValidationTask) ProtoMessage()               {}
func (*ValidationTask) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *ValidationTask) GetReporter() *ValidationTask_Args {
	if m != nil {
		return m.Reporter
	}
	return nil
}

func (m *ValidationTask) GetResourcer() *ValidationTask_Args {
	if m != nil {
		return m.Resourcer
	}
	return nil
}

type ValidationTask_Args struct {
	Service *Service `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	Config  []byte   `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (m *ValidationTask_Args) Reset()                    { *m = ValidationTask_Args{} }
func (m *ValidationTask_Args) String() string            { return proto.CompactTextString(m) }
func (*ValidationTask_Args) ProtoMessage()               {}
func (*ValidationTask_Args) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 0} }

func (m *ValidationTask_Args) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type ValidateProgress struct {
	Reporter     *ValidateProgress_Progress `protobuf:"bytes,1,opt,name=reporter" json:"reporter,omitempty"`
	Resourser    *ValidateProgress_Progress `protobuf:"bytes,2,opt,name=resourser" json:"resourser,omitempty"`
	LinkToReport string                     `protobuf:"bytes,3,opt,name=link_to_report,json=linkToReport" json:"link_to_report,omitempty"`
}

func (m *ValidateProgress) Reset()                    { *m = ValidateProgress{} }
func (m *ValidateProgress) String() string            { return proto.CompactTextString(m) }
func (*ValidateProgress) ProtoMessage()               {}
func (*ValidateProgress) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *ValidateProgress) GetReporter() *ValidateProgress_Progress {
	if m != nil {
		return m.Reporter
	}
	return nil
}

func (m *ValidateProgress) GetResourser() *ValidateProgress_Progress {
	if m != nil {
		return m.Resourser
	}
	return nil
}

type ValidateProgress_Progress struct {
	//  ID of service
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// all data
	Total int64 `protobuf:"varint,2,opt,name=total" json:"total,omitempty"`
	// already processed
	Current int64 `protobuf:"varint,3,opt,name=current" json:"current,omitempty"`
}

func (m *ValidateProgress_Progress) Reset()                    { *m = ValidateProgress_Progress{} }
func (m *ValidateProgress_Progress) String() string            { return proto.CompactTextString(m) }
func (*ValidateProgress_Progress) ProtoMessage()               {}
func (*ValidateProgress_Progress) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1, 0} }

func init() {
	proto.RegisterType((*ValidationTask)(nil), "lintflow.core.ValidationTask")
	proto.RegisterType((*ValidationTask_Args)(nil), "lintflow.core.ValidationTask.Args")
	proto.RegisterType((*ValidateProgress)(nil), "lintflow.core.ValidateProgress")
	proto.RegisterType((*ValidateProgress_Progress)(nil), "lintflow.core.ValidateProgress.Progress")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for ValidatorService service

type ValidatorServiceClient interface {
	Validate(ctx context.Context, in *ValidationTask, opts ...grpc.CallOption) (ValidatorService_ValidateClient, error)
}

type validatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewValidatorServiceClient(cc *grpc.ClientConn) ValidatorServiceClient {
	return &validatorServiceClient{cc}
}

func (c *validatorServiceClient) Validate(ctx context.Context, in *ValidationTask, opts ...grpc.CallOption) (ValidatorService_ValidateClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ValidatorService_serviceDesc.Streams[0], c.cc, "/lintflow.core.ValidatorService/Validate", opts...)
	if err != nil {
		return nil, err
	}
	x := &validatorServiceValidateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ValidatorService_ValidateClient interface {
	Recv() (*ValidateProgress, error)
	grpc.ClientStream
}

type validatorServiceValidateClient struct {
	grpc.ClientStream
}

func (x *validatorServiceValidateClient) Recv() (*ValidateProgress, error) {
	m := new(ValidateProgress)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ValidatorService service

type ValidatorServiceServer interface {
	Validate(*ValidationTask, ValidatorService_ValidateServer) error
}

func RegisterValidatorServiceServer(s *grpc.Server, srv ValidatorServiceServer) {
	s.RegisterService(&_ValidatorService_serviceDesc, srv)
}

func _ValidatorService_Validate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ValidationTask)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ValidatorServiceServer).Validate(m, &validatorServiceValidateServer{stream})
}

type ValidatorService_ValidateServer interface {
	Send(*ValidateProgress) error
	grpc.ServerStream
}

type validatorServiceValidateServer struct {
	grpc.ServerStream
}

func (x *validatorServiceValidateServer) Send(m *ValidateProgress) error {
	return x.ServerStream.SendMsg(m)
}

var _ValidatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lintflow.core.ValidatorService",
	HandlerType: (*ValidatorServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Validate",
			Handler:       _ValidatorService_Validate_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptor4 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x4e, 0xf2, 0x40,
	0x14, 0xfd, 0x5a, 0x3e, 0xf9, 0xb9, 0x20, 0x31, 0xa3, 0x12, 0xd2, 0xc4, 0x68, 0x1a, 0x17, 0xac,
	0x46, 0x82, 0x7b, 0xa3, 0x89, 0x71, 0xe1, 0xc2, 0x90, 0x91, 0xb8, 0x25, 0xa5, 0x0c, 0x64, 0x42,
	0xd3, 0x4b, 0xee, 0x0c, 0xf8, 0x1e, 0x3e, 0xaf, 0x0b, 0xcb, 0x0c, 0x2d, 0xc5, 0x18, 0x65, 0x37,
	0xf7, 0xe7, 0x9c, 0x7b, 0xce, 0x69, 0xe1, 0x7c, 0x49, 0x68, 0xf0, 0x66, 0x1d, 0x25, 0x6a, 0x1a,
	0x19, 0x24, 0x6e, 0x6b, 0x76, 0x9c, 0xa8, 0xd4, 0xcc, 0x12, 0x7c, 0xe7, 0x31, 0x92, 0x0c, 0x4e,
	0xdd, 0x56, 0x82, 0xb8, 0x58, 0x2d, 0xa7, 0x6e, 0x27, 0xfc, 0xf4, 0xa0, 0xfd, 0xe6, 0x70, 0x0a,
	0xd3, 0x51, 0xa4, 0x17, 0xac, 0x03, 0xd5, 0x18, 0xd3, 0x99, 0x9a, 0x77, 0xbd, 0x2b, 0xaf, 0xd7,
	0x12, 0xdb, 0x8a, 0xdd, 0x41, 0x9d, 0xe4, 0x12, 0xc9, 0x48, 0xea, 0xfa, 0xd9, 0xa4, 0x39, 0x08,
	0xf9, 0xde, 0x05, 0xbe, 0x4f, 0xc4, 0x1f, 0x68, 0xae, 0x45, 0x81, 0x61, 0xf7, 0xd0, 0x20, 0xa9,
	0x71, 0x45, 0x71, 0x46, 0x50, 0x39, 0x98, 0x60, 0x07, 0x0a, 0x86, 0xf0, 0x7f, 0xd3, 0x62, 0x7d,
	0xa8, 0x69, 0x49, 0x6b, 0x15, 0x4b, 0x2b, 0xb1, 0x39, 0xe8, 0x7c, 0xe3, 0x79, 0x75, 0x53, 0x91,
	0xaf, 0x95, 0x3c, 0xf9, 0x65, 0x4f, 0xe1, 0x87, 0x0f, 0x27, 0xdb, 0xa3, 0x72, 0x48, 0x38, 0xcf,
	0x6e, 0x69, 0xf6, 0x58, 0x32, 0xea, 0xf8, 0x7b, 0x3f, 0xeb, 0x2c, 0x20, 0x3c, 0x7f, 0x94, 0xec,
	0x3e, 0xe5, 0x76, 0x75, 0x91, 0xd7, 0xe1, 0x34, 0x3b, 0x28, 0xbb, 0x86, 0x76, 0x86, 0x5a, 0x8c,
	0x0d, 0x8e, 0x1d, 0xb7, 0xcd, 0xae, 0x21, 0x5a, 0x9b, 0xee, 0x08, 0x85, 0xed, 0x05, 0xcf, 0x50,
	0x2f, 0xf4, 0xb7, 0xc1, 0x57, 0x53, 0xab, 0xbc, 0x21, 0xb2, 0x17, 0x3b, 0x83, 0x23, 0x83, 0x26,
	0x4a, 0xac, 0x8a, 0x8a, 0x70, 0x05, 0xeb, 0x42, 0x2d, 0x5e, 0x11, 0xc9, 0xd4, 0x11, 0x56, 0x44,
	0x5e, 0x0e, 0x26, 0x45, 0x26, 0x48, 0xdb, 0x24, 0xd9, 0x0b, 0xd4, 0x73, 0xb5, 0xec, 0xe2, 0xd7,
	0xaf, 0x16, 0x5c, 0xfe, 0xe1, 0x32, 0xfc, 0xd7, 0xf7, 0x26, 0x55, 0xfb, 0xfb, 0xdd, 0x7e, 0x05,
	0x00, 0x00, 0xff, 0xff, 0xf3, 0x58, 0xfc, 0xf9, 0xbb, 0x02, 0x00, 0x00,
}

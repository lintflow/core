// Code generated by protoc-gen-go.
// source: proto/lookupd.proto
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

type Service_Type int32

const (
	Service_RESOURCER Service_Type = 0
	Service_LINTER    Service_Type = 1
	Service_REPORTER  Service_Type = 2
)

var Service_Type_name = map[int32]string{
	0: "RESOURCER",
	1: "LINTER",
	2: "REPORTER",
}
var Service_Type_value = map[string]int32{
	"RESOURCER": 0,
	"LINTER":    1,
	"REPORTER":  2,
}

func (x Service_Type) String() string {
	return proto.EnumName(Service_Type_name, int32(x))
}
func (Service_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

type ListRequest_SearchType int32

const (
	ListRequest_ANY       ListRequest_SearchType = 0
	ListRequest_RESOURCER ListRequest_SearchType = 1
	ListRequest_LINTER    ListRequest_SearchType = 2
	ListRequest_REPORTER  ListRequest_SearchType = 3
)

var ListRequest_SearchType_name = map[int32]string{
	0: "ANY",
	1: "RESOURCER",
	2: "LINTER",
	3: "REPORTER",
}
var ListRequest_SearchType_value = map[string]int32{
	"ANY":       0,
	"RESOURCER": 1,
	"LINTER":    2,
	"REPORTER":  3,
}

func (x ListRequest_SearchType) String() string {
	return proto.EnumName(ListRequest_SearchType_name, int32(x))
}
func (ListRequest_SearchType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 0} }

// тип сервис описывает объект сервиса
type Service struct {
	// юникальный идентификатор сервиса
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// адрес сервиса в сети
	//  как пример localhost:45679
	Address     string   `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Name        string   `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	Tags        []string `protobuf:"bytes,5,rep,name=tags" json:"tags,omitempty"`
	// декларирует тип сервиса
	Type Service_Type `protobuf:"varint,6,opt,name=type,enum=lintflow.core.Service_Type" json:"type,omitempty"`
	// конфигруатор сервиса
	// там описана специфика севиса для выполенкния задач
	TaskConfig []byte `protobuf:"bytes,7,opt,name=task_config,json=taskConfig,proto3" json:"task_config,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type ListRequest struct {
	Type ListRequest_SearchType `protobuf:"varint,1,opt,name=type,enum=lintflow.core.ListRequest_SearchType" json:"type,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type ListResponse struct {
	Services []*Service `protobuf:"bytes,1,rep,name=services" json:"services,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ListResponse) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

// запрос на регистрацию сервиса
type RegisterRequest struct {
	Service *Service `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *RegisterRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

// ответ на регистрацию
type RegisterResponse struct {
	Ok bool `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
}

func (m *RegisterResponse) Reset()                    { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()               {}
func (*RegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func init() {
	proto.RegisterType((*Service)(nil), "lintflow.core.Service")
	proto.RegisterType((*ListRequest)(nil), "lintflow.core.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "lintflow.core.ListResponse")
	proto.RegisterType((*RegisterRequest)(nil), "lintflow.core.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "lintflow.core.RegisterResponse")
	proto.RegisterEnum("lintflow.core.Service_Type", Service_Type_name, Service_Type_value)
	proto.RegisterEnum("lintflow.core.ListRequest_SearchType", ListRequest_SearchType_name, ListRequest_SearchType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for LookupdService service

type LookupdServiceClient interface {
	// метод позволяет сервисам реистриваться сервисам
	// для того чтобы потом их можно было использовать инспектору
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type lookupdServiceClient struct {
	cc *grpc.ClientConn
}

func NewLookupdServiceClient(cc *grpc.ClientConn) LookupdServiceClient {
	return &lookupdServiceClient{cc}
}

func (c *lookupdServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := grpc.Invoke(ctx, "/lintflow.core.LookupdService/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lookupdServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/lintflow.core.LookupdService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LookupdService service

type LookupdServiceServer interface {
	// метод позволяет сервисам реистриваться сервисам
	// для того чтобы потом их можно было использовать инспектору
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
}

func RegisterLookupdServiceServer(s *grpc.Server, srv LookupdServiceServer) {
	s.RegisterService(&_LookupdService_serviceDesc, srv)
}

func _LookupdService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(LookupdServiceServer).Register(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _LookupdService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(LookupdServiceServer).List(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _LookupdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lintflow.core.LookupdService",
	HandlerType: (*LookupdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _LookupdService_Register_Handler,
		},
		{
			MethodName: "List",
			Handler:    _LookupdService_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor1 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x92, 0x4f, 0xcf, 0xd2, 0x40,
	0x10, 0xc6, 0x69, 0x8b, 0x14, 0xa6, 0x80, 0x64, 0x4c, 0xcc, 0x06, 0x12, 0x21, 0x9b, 0x98, 0x70,
	0x2a, 0xa6, 0x9e, 0xbc, 0x98, 0x20, 0xe1, 0x60, 0x82, 0x60, 0x16, 0x3c, 0x78, 0x32, 0xb5, 0x5d,
	0xb0, 0x01, 0xbb, 0xb5, 0x5b, 0x34, 0xde, 0x3c, 0xfa, 0x39, 0xde, 0x4f, 0xfa, 0x6e, 0xb7, 0xed,
	0xcb, 0x9f, 0xc0, 0x6d, 0xfa, 0xcc, 0x33, 0xcf, 0xfe, 0x76, 0x3b, 0xf0, 0x22, 0x49, 0x45, 0x26,
	0x26, 0x07, 0x21, 0xf6, 0xc7, 0x24, 0x74, 0xf5, 0x17, 0x76, 0x0e, 0x51, 0x9c, 0x6d, 0x0f, 0xe2,
	0x8f, 0x1b, 0x88, 0x94, 0xd3, 0x7f, 0x26, 0xd8, 0x6b, 0x9e, 0xfe, 0x8e, 0x02, 0x8e, 0x5d, 0x30,
	0xa3, 0x90, 0x18, 0x23, 0x63, 0xdc, 0x62, 0xaa, 0x42, 0x02, 0xb6, 0x1f, 0x86, 0x29, 0x97, 0x92,
	0x98, 0x5a, 0xac, 0x3e, 0x11, 0xa1, 0x1e, 0xfb, 0x3f, 0x39, 0xb1, 0xb4, 0xac, 0x6b, 0x1c, 0x81,
	0x13, 0x72, 0x19, 0xa4, 0x51, 0x92, 0x45, 0x22, 0x26, 0x75, 0xdd, 0x3a, 0x97, 0xf2, 0xa9, 0xcc,
	0xdf, 0x49, 0xf2, 0x6c, 0x64, 0xe5, 0x53, 0x79, 0x8d, 0x13, 0xa5, 0xfd, 0x4d, 0x38, 0x69, 0x28,
	0x7b, 0xd7, 0x1b, 0xb8, 0x17, 0x74, 0x6e, 0x49, 0xe6, 0x6e, 0x94, 0x85, 0x69, 0x23, 0x0e, 0xc1,
	0xc9, 0x7c, 0xb9, 0xff, 0x16, 0x88, 0x78, 0x1b, 0xed, 0x88, 0xad, 0xe6, 0xda, 0x0c, 0x72, 0x69,
	0xa6, 0x15, 0xaa, 0x12, 0x73, 0x3b, 0x76, 0xa0, 0xc5, 0xe6, 0xeb, 0xd5, 0x17, 0x36, 0x9b, 0xb3,
	0x5e, 0x0d, 0x01, 0x1a, 0x8b, 0x8f, 0xcb, 0x8d, 0xaa, 0x0d, 0x6c, 0x43, 0x93, 0xcd, 0x3f, 0xaf,
	0x58, 0xfe, 0x65, 0xd2, 0xff, 0x06, 0x38, 0x8b, 0x48, 0x66, 0x8c, 0xff, 0x3a, 0x72, 0x99, 0xe1,
	0xbb, 0x12, 0xc9, 0xd0, 0x48, 0xaf, 0xaf, 0x90, 0xce, 0x9c, 0x0a, 0xcf, 0x4f, 0x83, 0x1f, 0x27,
	0x38, 0xfa, 0x1e, 0xe0, 0xa4, 0xa1, 0x0d, 0xd6, 0x74, 0xf9, 0x55, 0x9d, 0x7d, 0x81, 0x62, 0x9c,
	0xa1, 0x98, 0x17, 0x28, 0x16, 0xfd, 0x00, 0xed, 0x22, 0x5f, 0x26, 0x22, 0x96, 0x1c, 0x3d, 0x68,
	0xca, 0xe2, 0x09, 0xa4, 0xc2, 0xb1, 0xc6, 0x8e, 0xf7, 0xf2, 0xf6, 0x0b, 0xb1, 0x27, 0x1f, 0x9d,
	0xc1, 0x73, 0xc6, 0x77, 0x2a, 0x85, 0xa7, 0xd5, 0x8d, 0xde, 0x80, 0x5d, 0xb6, 0xf5, 0xa5, 0xee,
	0xa7, 0x54, 0x36, 0x4a, 0xa1, 0x77, 0x0a, 0x29, 0x61, 0xd4, 0x7a, 0x88, 0xbd, 0x0e, 0x68, 0x32,
	0x55, 0x79, 0x0f, 0x06, 0x74, 0x17, 0xc5, 0x6e, 0x55, 0x1b, 0xf4, 0x49, 0xdd, 0xa6, 0x1c, 0xc3,
	0x57, 0x57, 0x67, 0x5c, 0x41, 0xf5, 0x87, 0x77, 0xfb, 0xc5, 0x79, 0xb4, 0x86, 0x53, 0xa8, 0xe7,
	0xcf, 0x81, 0xfd, 0xfb, 0xff, 0xa0, 0x3f, 0xb8, 0xd9, 0xab, 0x22, 0xbe, 0x37, 0xf4, 0xd6, 0xbf,
	0x7d, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x88, 0x53, 0x4e, 0x0c, 0x03, 0x00, 0x00,
}

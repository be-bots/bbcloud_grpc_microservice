// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: system.proto

package model

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SystemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SystemRequest) Reset() {
	*x = SystemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemRequest) ProtoMessage() {}

func (x *SystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemRequest.ProtoReflect.Descriptor instead.
func (*SystemRequest) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{0}
}

func (x *SystemRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type System struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DsId string `protobuf:"bytes,2,opt,name=ds_id,json=dsId,proto3" json:"ds_id,omitempty"`
	UsId string `protobuf:"bytes,3,opt,name=us_id,json=usId,proto3" json:"us_id,omitempty"`
}

func (x *System) Reset() {
	*x = System{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *System) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*System) ProtoMessage() {}

func (x *System) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use System.ProtoReflect.Descriptor instead.
func (*System) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{1}
}

func (x *System) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *System) GetDsId() string {
	if x != nil {
		return x.DsId
	}
	return ""
}

func (x *System) GetUsId() string {
	if x != nil {
		return x.UsId
	}
	return ""
}

type Systems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Systems []*System `protobuf:"bytes,1,rep,name=systems,proto3" json:"systems,omitempty"`
}

func (x *Systems) Reset() {
	*x = Systems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Systems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Systems) ProtoMessage() {}

func (x *Systems) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Systems.ProtoReflect.Descriptor instead.
func (*Systems) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{2}
}

func (x *Systems) GetSystems() []*System {
	if x != nil {
		return x.Systems
	}
	return nil
}

var File_system_proto protoreflect.FileDescriptor

var file_system_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x0d, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x06, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x13, 0x0a,
	0x05, 0x64, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x73,
	0x49, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x73, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x07, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x52, 0x07, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x7a, 0x0a, 0x0d, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x37,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_system_proto_rawDescOnce sync.Once
	file_system_proto_rawDescData = file_system_proto_rawDesc
)

func file_system_proto_rawDescGZIP() []byte {
	file_system_proto_rawDescOnce.Do(func() {
		file_system_proto_rawDescData = protoimpl.X.CompressGZIP(file_system_proto_rawDescData)
	})
	return file_system_proto_rawDescData
}

var file_system_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_system_proto_goTypes = []interface{}{
	(*SystemRequest)(nil), // 0: model.SystemRequest
	(*System)(nil),        // 1: model.System
	(*Systems)(nil),       // 2: model.Systems
	(*empty.Empty)(nil),   // 3: google.protobuf.Empty
}
var file_system_proto_depIdxs = []int32{
	1, // 0: model.Systems.systems:type_name -> model.System
	0, // 1: model.SystemService.GetSystem:input_type -> model.SystemRequest
	3, // 2: model.SystemService.GetAllSystems:input_type -> google.protobuf.Empty
	1, // 3: model.SystemService.GetSystem:output_type -> model.System
	2, // 4: model.SystemService.GetAllSystems:output_type -> model.Systems
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_system_proto_init() }
func file_system_proto_init() {
	if File_system_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_system_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_system_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*System); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_system_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Systems); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_system_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_system_proto_goTypes,
		DependencyIndexes: file_system_proto_depIdxs,
		MessageInfos:      file_system_proto_msgTypes,
	}.Build()
	File_system_proto = out.File
	file_system_proto_rawDesc = nil
	file_system_proto_goTypes = nil
	file_system_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SystemServiceClient is the client API for SystemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SystemServiceClient interface {
	GetSystem(ctx context.Context, in *SystemRequest, opts ...grpc.CallOption) (*System, error)
	GetAllSystems(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Systems, error)
}

type systemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemServiceClient(cc grpc.ClientConnInterface) SystemServiceClient {
	return &systemServiceClient{cc}
}

func (c *systemServiceClient) GetSystem(ctx context.Context, in *SystemRequest, opts ...grpc.CallOption) (*System, error) {
	out := new(System)
	err := c.cc.Invoke(ctx, "/model.SystemService/GetSystem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) GetAllSystems(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Systems, error) {
	out := new(Systems)
	err := c.cc.Invoke(ctx, "/model.SystemService/GetAllSystems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemServiceServer is the server API for SystemService service.
type SystemServiceServer interface {
	GetSystem(context.Context, *SystemRequest) (*System, error)
	GetAllSystems(context.Context, *empty.Empty) (*Systems, error)
}

// UnimplementedSystemServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSystemServiceServer struct {
}

func (*UnimplementedSystemServiceServer) GetSystem(context.Context, *SystemRequest) (*System, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystem not implemented")
}
func (*UnimplementedSystemServiceServer) GetAllSystems(context.Context, *empty.Empty) (*Systems, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllSystems not implemented")
}

func RegisterSystemServiceServer(s *grpc.Server, srv SystemServiceServer) {
	s.RegisterService(&_SystemService_serviceDesc, srv)
}

func _SystemService_GetSystem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).GetSystem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.SystemService/GetSystem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).GetSystem(ctx, req.(*SystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_GetAllSystems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).GetAllSystems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.SystemService/GetAllSystems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).GetAllSystems(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _SystemService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.SystemService",
	HandlerType: (*SystemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSystem",
			Handler:    _SystemService_GetSystem_Handler,
		},
		{
			MethodName: "GetAllSystems",
			Handler:    _SystemService_GetAllSystems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system.proto",
}

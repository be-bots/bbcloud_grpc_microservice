// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: gutter.proto

package model

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type GutterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rfid string `protobuf:"bytes,1,opt,name=rfid,proto3" json:"rfid,omitempty"`
}

func (x *GutterRequest) Reset() {
	*x = GutterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gutter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GutterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GutterRequest) ProtoMessage() {}

func (x *GutterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gutter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GutterRequest.ProtoReflect.Descriptor instead.
func (*GutterRequest) Descriptor() ([]byte, []int) {
	return file_gutter_proto_rawDescGZIP(), []int{0}
}

func (x *GutterRequest) GetRfid() string {
	if x != nil {
		return x.Rfid
	}
	return ""
}

type Gutter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rfid     string               `protobuf:"bytes,1,opt,name=rfid,proto3" json:"rfid,omitempty"`
	ProdDate *timestamp.Timestamp `protobuf:"bytes,2,opt,name=prod_date,json=prodDate,proto3" json:"prod_date,omitempty"`
	Status   string               `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Gutter) Reset() {
	*x = Gutter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gutter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gutter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gutter) ProtoMessage() {}

func (x *Gutter) ProtoReflect() protoreflect.Message {
	mi := &file_gutter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gutter.ProtoReflect.Descriptor instead.
func (*Gutter) Descriptor() ([]byte, []int) {
	return file_gutter_proto_rawDescGZIP(), []int{1}
}

func (x *Gutter) GetRfid() string {
	if x != nil {
		return x.Rfid
	}
	return ""
}

func (x *Gutter) GetProdDate() *timestamp.Timestamp {
	if x != nil {
		return x.ProdDate
	}
	return nil
}

func (x *Gutter) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Gutters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gutters []*Gutter `protobuf:"bytes,1,rep,name=gutters,proto3" json:"gutters,omitempty"`
}

func (x *Gutters) Reset() {
	*x = Gutters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gutter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gutters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gutters) ProtoMessage() {}

func (x *Gutters) ProtoReflect() protoreflect.Message {
	mi := &file_gutter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gutters.ProtoReflect.Descriptor instead.
func (*Gutters) Descriptor() ([]byte, []int) {
	return file_gutter_proto_rawDescGZIP(), []int{2}
}

func (x *Gutters) GetGutters() []*Gutter {
	if x != nil {
		return x.Gutters
	}
	return nil
}

var File_gutter_proto protoreflect.FileDescriptor

var file_gutter_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0d, 0x47, 0x75, 0x74, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x66, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x72, 0x66, 0x69, 0x64, 0x22, 0x6d, 0x0a, 0x06, 0x47, 0x75, 0x74, 0x74,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x66, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x72, 0x66, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x32, 0x0a, 0x07, 0x47, 0x75, 0x74, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x67, 0x75, 0x74, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x75, 0x74, 0x74,
	0x65, 0x72, 0x52, 0x07, 0x67, 0x75, 0x74, 0x74, 0x65, 0x72, 0x73, 0x32, 0x7a, 0x0a, 0x0d, 0x47,
	0x75, 0x74, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x47, 0x75, 0x74, 0x74, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x47, 0x75, 0x74, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x75, 0x74, 0x74, 0x65, 0x72, 0x12, 0x37,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x75, 0x74, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x47, 0x75, 0x74, 0x74, 0x65, 0x72, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gutter_proto_rawDescOnce sync.Once
	file_gutter_proto_rawDescData = file_gutter_proto_rawDesc
)

func file_gutter_proto_rawDescGZIP() []byte {
	file_gutter_proto_rawDescOnce.Do(func() {
		file_gutter_proto_rawDescData = protoimpl.X.CompressGZIP(file_gutter_proto_rawDescData)
	})
	return file_gutter_proto_rawDescData
}

var file_gutter_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_gutter_proto_goTypes = []interface{}{
	(*GutterRequest)(nil),       // 0: model.GutterRequest
	(*Gutter)(nil),              // 1: model.Gutter
	(*Gutters)(nil),             // 2: model.Gutters
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*empty.Empty)(nil),         // 4: google.protobuf.Empty
}
var file_gutter_proto_depIdxs = []int32{
	3, // 0: model.Gutter.prod_date:type_name -> google.protobuf.Timestamp
	1, // 1: model.Gutters.gutters:type_name -> model.Gutter
	0, // 2: model.GutterService.GetGutter:input_type -> model.GutterRequest
	4, // 3: model.GutterService.GetAllGutters:input_type -> google.protobuf.Empty
	1, // 4: model.GutterService.GetGutter:output_type -> model.Gutter
	2, // 5: model.GutterService.GetAllGutters:output_type -> model.Gutters
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_gutter_proto_init() }
func file_gutter_proto_init() {
	if File_gutter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gutter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GutterRequest); i {
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
		file_gutter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gutter); i {
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
		file_gutter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gutters); i {
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
			RawDescriptor: file_gutter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gutter_proto_goTypes,
		DependencyIndexes: file_gutter_proto_depIdxs,
		MessageInfos:      file_gutter_proto_msgTypes,
	}.Build()
	File_gutter_proto = out.File
	file_gutter_proto_rawDesc = nil
	file_gutter_proto_goTypes = nil
	file_gutter_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GutterServiceClient is the client API for GutterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GutterServiceClient interface {
	GetGutter(ctx context.Context, in *GutterRequest, opts ...grpc.CallOption) (*Gutter, error)
	GetAllGutters(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Gutters, error)
}

type gutterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGutterServiceClient(cc grpc.ClientConnInterface) GutterServiceClient {
	return &gutterServiceClient{cc}
}

func (c *gutterServiceClient) GetGutter(ctx context.Context, in *GutterRequest, opts ...grpc.CallOption) (*Gutter, error) {
	out := new(Gutter)
	err := c.cc.Invoke(ctx, "/model.GutterService/GetGutter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gutterServiceClient) GetAllGutters(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Gutters, error) {
	out := new(Gutters)
	err := c.cc.Invoke(ctx, "/model.GutterService/GetAllGutters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GutterServiceServer is the server API for GutterService service.
type GutterServiceServer interface {
	GetGutter(context.Context, *GutterRequest) (*Gutter, error)
	GetAllGutters(context.Context, *empty.Empty) (*Gutters, error)
}

// UnimplementedGutterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGutterServiceServer struct {
}

func (*UnimplementedGutterServiceServer) GetGutter(context.Context, *GutterRequest) (*Gutter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGutter not implemented")
}
func (*UnimplementedGutterServiceServer) GetAllGutters(context.Context, *empty.Empty) (*Gutters, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllGutters not implemented")
}

func RegisterGutterServiceServer(s *grpc.Server, srv GutterServiceServer) {
	s.RegisterService(&_GutterService_serviceDesc, srv)
}

func _GutterService_GetGutter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GutterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GutterServiceServer).GetGutter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.GutterService/GetGutter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GutterServiceServer).GetGutter(ctx, req.(*GutterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GutterService_GetAllGutters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GutterServiceServer).GetAllGutters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.GutterService/GetAllGutters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GutterServiceServer).GetAllGutters(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _GutterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.GutterService",
	HandlerType: (*GutterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGutter",
			Handler:    _GutterService_GetGutter_Handler,
		},
		{
			MethodName: "GetAllGutters",
			Handler:    _GutterService_GetAllGutters_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gutter.proto",
}

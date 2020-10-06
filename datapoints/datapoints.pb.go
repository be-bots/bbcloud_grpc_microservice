// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: datapoints.proto

package datapoints

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

type SubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriptionIDs []string `protobuf:"bytes,1,rep,name=subscriptionIDs,proto3" json:"subscriptionIDs,omitempty"`
}

func (x *SubscriptionRequest) Reset() {
	*x = SubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datapoints_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionRequest) ProtoMessage() {}

func (x *SubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datapoints_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionRequest.ProtoReflect.Descriptor instead.
func (*SubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_datapoints_proto_rawDescGZIP(), []int{0}
}

func (x *SubscriptionRequest) GetSubscriptionIDs() []string {
	if x != nil {
		return x.SubscriptionIDs
	}
	return nil
}

type SubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SuccessfullySubscribedIDs []string `protobuf:"bytes,1,rep,name=successfullySubscribedIDs,proto3" json:"successfullySubscribedIDs,omitempty"`
}

func (x *SubscriptionResponse) Reset() {
	*x = SubscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datapoints_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionResponse) ProtoMessage() {}

func (x *SubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_datapoints_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionResponse.ProtoReflect.Descriptor instead.
func (*SubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_datapoints_proto_rawDescGZIP(), []int{1}
}

func (x *SubscriptionResponse) GetSuccessfullySubscribedIDs() []string {
	if x != nil {
		return x.SuccessfullySubscribedIDs
	}
	return nil
}

var File_datapoints_proto protoreflect.FileDescriptor

var file_datapoints_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x3f,
	0x0a, 0x13, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x73, 0x22,
	0x54, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x19, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x64, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x19, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x64, 0x49, 0x44, 0x73, 0x32, 0x63, 0x0a, 0x11, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x1f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x3b,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_datapoints_proto_rawDescOnce sync.Once
	file_datapoints_proto_rawDescData = file_datapoints_proto_rawDesc
)

func file_datapoints_proto_rawDescGZIP() []byte {
	file_datapoints_proto_rawDescOnce.Do(func() {
		file_datapoints_proto_rawDescData = protoimpl.X.CompressGZIP(file_datapoints_proto_rawDescData)
	})
	return file_datapoints_proto_rawDescData
}

var file_datapoints_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_datapoints_proto_goTypes = []interface{}{
	(*SubscriptionRequest)(nil),  // 0: datapoints.SubscriptionRequest
	(*SubscriptionResponse)(nil), // 1: datapoints.SubscriptionResponse
}
var file_datapoints_proto_depIdxs = []int32{
	0, // 0: datapoints.DatapointsService.Subscribe:input_type -> datapoints.SubscriptionRequest
	1, // 1: datapoints.DatapointsService.Subscribe:output_type -> datapoints.SubscriptionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_datapoints_proto_init() }
func file_datapoints_proto_init() {
	if File_datapoints_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datapoints_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionRequest); i {
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
		file_datapoints_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionResponse); i {
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
			RawDescriptor: file_datapoints_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_datapoints_proto_goTypes,
		DependencyIndexes: file_datapoints_proto_depIdxs,
		MessageInfos:      file_datapoints_proto_msgTypes,
	}.Build()
	File_datapoints_proto = out.File
	file_datapoints_proto_rawDesc = nil
	file_datapoints_proto_goTypes = nil
	file_datapoints_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DatapointsServiceClient is the client API for DatapointsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DatapointsServiceClient interface {
	Subscribe(ctx context.Context, in *SubscriptionRequest, opts ...grpc.CallOption) (*SubscriptionResponse, error)
}

type datapointsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatapointsServiceClient(cc grpc.ClientConnInterface) DatapointsServiceClient {
	return &datapointsServiceClient{cc}
}

func (c *datapointsServiceClient) Subscribe(ctx context.Context, in *SubscriptionRequest, opts ...grpc.CallOption) (*SubscriptionResponse, error) {
	out := new(SubscriptionResponse)
	err := c.cc.Invoke(ctx, "/datapoints.DatapointsService/Subscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatapointsServiceServer is the server API for DatapointsService service.
type DatapointsServiceServer interface {
	Subscribe(context.Context, *SubscriptionRequest) (*SubscriptionResponse, error)
}

// UnimplementedDatapointsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDatapointsServiceServer struct {
}

func (*UnimplementedDatapointsServiceServer) Subscribe(context.Context, *SubscriptionRequest) (*SubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterDatapointsServiceServer(s *grpc.Server, srv DatapointsServiceServer) {
	s.RegisterService(&_DatapointsService_serviceDesc, srv)
}

func _DatapointsService_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatapointsServiceServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datapoints.DatapointsService/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatapointsServiceServer).Subscribe(ctx, req.(*SubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DatapointsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "datapoints.DatapointsService",
	HandlerType: (*DatapointsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Subscribe",
			Handler:    _DatapointsService_Subscribe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "datapoints.proto",
}

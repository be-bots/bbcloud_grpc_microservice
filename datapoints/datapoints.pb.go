// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: datapoints.proto

package datapoints

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type DatapointsRequestedStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data map[string]*any.Any `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DatapointsRequestedStream) Reset() {
	*x = DatapointsRequestedStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datapoints_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatapointsRequestedStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatapointsRequestedStream) ProtoMessage() {}

func (x *DatapointsRequestedStream) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DatapointsRequestedStream.ProtoReflect.Descriptor instead.
func (*DatapointsRequestedStream) Descriptor() ([]byte, []int) {
	return file_datapoints_proto_rawDescGZIP(), []int{0}
}

func (x *DatapointsRequestedStream) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DatapointsRequestedStream) GetData() map[string]*any.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_datapoints_proto protoreflect.FileDescriptor

var file_datapoints_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x19, 0x44, 0x61, 0x74, 0x61, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x43, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x4d, 0x0a, 0x09, 0x44, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x69, 0x0a, 0x11, 0x44, 0x61, 0x74, 0x61,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a,
	0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x25, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x30, 0x01, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	(*DatapointsRequestedStream)(nil), // 0: datapoints.DatapointsRequestedStream
	nil,                               // 1: datapoints.DatapointsRequestedStream.DataEntry
	(*any.Any)(nil),                   // 2: google.protobuf.Any
	(*empty.Empty)(nil),               // 3: google.protobuf.Empty
}
var file_datapoints_proto_depIdxs = []int32{
	1, // 0: datapoints.DatapointsRequestedStream.data:type_name -> datapoints.DatapointsRequestedStream.DataEntry
	2, // 1: datapoints.DatapointsRequestedStream.DataEntry.value:type_name -> google.protobuf.Any
	3, // 2: datapoints.DatapointsService.RequestDatapoints:input_type -> google.protobuf.Empty
	0, // 3: datapoints.DatapointsService.RequestDatapoints:output_type -> datapoints.DatapointsRequestedStream
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_datapoints_proto_init() }
func file_datapoints_proto_init() {
	if File_datapoints_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datapoints_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatapointsRequestedStream); i {
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
	RequestDatapoints(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (DatapointsService_RequestDatapointsClient, error)
}

type datapointsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatapointsServiceClient(cc grpc.ClientConnInterface) DatapointsServiceClient {
	return &datapointsServiceClient{cc}
}

func (c *datapointsServiceClient) RequestDatapoints(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (DatapointsService_RequestDatapointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DatapointsService_serviceDesc.Streams[0], "/datapoints.DatapointsService/RequestDatapoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &datapointsServiceRequestDatapointsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DatapointsService_RequestDatapointsClient interface {
	Recv() (*DatapointsRequestedStream, error)
	grpc.ClientStream
}

type datapointsServiceRequestDatapointsClient struct {
	grpc.ClientStream
}

func (x *datapointsServiceRequestDatapointsClient) Recv() (*DatapointsRequestedStream, error) {
	m := new(DatapointsRequestedStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DatapointsServiceServer is the server API for DatapointsService service.
type DatapointsServiceServer interface {
	RequestDatapoints(*empty.Empty, DatapointsService_RequestDatapointsServer) error
}

// UnimplementedDatapointsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDatapointsServiceServer struct {
}

func (*UnimplementedDatapointsServiceServer) RequestDatapoints(*empty.Empty, DatapointsService_RequestDatapointsServer) error {
	return status.Errorf(codes.Unimplemented, "method RequestDatapoints not implemented")
}

func RegisterDatapointsServiceServer(s *grpc.Server, srv DatapointsServiceServer) {
	s.RegisterService(&_DatapointsService_serviceDesc, srv)
}

func _DatapointsService_RequestDatapoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DatapointsServiceServer).RequestDatapoints(m, &datapointsServiceRequestDatapointsServer{stream})
}

type DatapointsService_RequestDatapointsServer interface {
	Send(*DatapointsRequestedStream) error
	grpc.ServerStream
}

type datapointsServiceRequestDatapointsServer struct {
	grpc.ServerStream
}

func (x *datapointsServiceRequestDatapointsServer) Send(m *DatapointsRequestedStream) error {
	return x.ServerStream.SendMsg(m)
}

var _DatapointsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "datapoints.DatapointsService",
	HandlerType: (*DatapointsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RequestDatapoints",
			Handler:       _DatapointsService_RequestDatapoints_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "datapoints.proto",
}

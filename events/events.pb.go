// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: events.proto

package events

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

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{0}
}

func (x *Meta) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorOcurred bool   `protobuf:"varint,1,opt,name=ErrorOcurred,proto3" json:"ErrorOcurred,omitempty"`
	ErrorType    string `protobuf:"bytes,2,opt,name=ErrorType,proto3" json:"ErrorType,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{1}
}

func (x *Status) GetErrorOcurred() bool {
	if x != nil {
		return x.ErrorOcurred
	}
	return false
}

func (x *Status) GetErrorType() string {
	if x != nil {
		return x.ErrorType
	}
	return ""
}

type Source struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Farm  string `protobuf:"bytes,1,opt,name=Farm,proto3" json:"Farm,omitempty"`
	Cell  int32  `protobuf:"varint,2,opt,name=Cell,proto3" json:"Cell,omitempty"`
	Layer int32  `protobuf:"varint,3,opt,name=Layer,proto3" json:"Layer,omitempty"`
	Slot  int32  `protobuf:"varint,4,opt,name=Slot,proto3" json:"Slot,omitempty"`
	RFID  string `protobuf:"bytes,5,opt,name=RFID,proto3" json:"RFID,omitempty"`
}

func (x *Source) Reset() {
	*x = Source{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Source) ProtoMessage() {}

func (x *Source) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Source.ProtoReflect.Descriptor instead.
func (*Source) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{2}
}

func (x *Source) GetFarm() string {
	if x != nil {
		return x.Farm
	}
	return ""
}

func (x *Source) GetCell() int32 {
	if x != nil {
		return x.Cell
	}
	return 0
}

func (x *Source) GetLayer() int32 {
	if x != nil {
		return x.Layer
	}
	return 0
}

func (x *Source) GetSlot() int32 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *Source) GetRFID() string {
	if x != nil {
		return x.RFID
	}
	return ""
}

type Watering struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreWateringWeight   int32  `protobuf:"varint,1,opt,name=PreWateringWeight,proto3" json:"PreWateringWeight,omitempty"`
	PostWateringWeight  int32  `protobuf:"varint,2,opt,name=PostWateringWeight,proto3" json:"PostWateringWeight,omitempty"`
	TypeOfNossle        string `protobuf:"bytes,3,opt,name=TypeOfNossle,proto3" json:"TypeOfNossle,omitempty"`
	WaterFanCombination bool   `protobuf:"varint,4,opt,name=WaterFanCombination,proto3" json:"WaterFanCombination,omitempty"`
}

func (x *Watering) Reset() {
	*x = Watering{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Watering) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Watering) ProtoMessage() {}

func (x *Watering) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Watering.ProtoReflect.Descriptor instead.
func (*Watering) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{3}
}

func (x *Watering) GetPreWateringWeight() int32 {
	if x != nil {
		return x.PreWateringWeight
	}
	return 0
}

func (x *Watering) GetPostWateringWeight() int32 {
	if x != nil {
		return x.PostWateringWeight
	}
	return 0
}

func (x *Watering) GetTypeOfNossle() string {
	if x != nil {
		return x.TypeOfNossle
	}
	return ""
}

func (x *Watering) GetWaterFanCombination() bool {
	if x != nil {
		return x.WaterFanCombination
	}
	return false
}

type Datapoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Xpos            int32   `protobuf:"varint,1,opt,name=Xpos,proto3" json:"Xpos,omitempty"`
	Temp            float32 `protobuf:"fixed32,2,opt,name=Temp,proto3" json:"Temp,omitempty"`
	Hum             float32 `protobuf:"fixed32,3,opt,name=Hum,proto3" json:"Hum,omitempty"`
	Press           float32 `protobuf:"fixed32,4,opt,name=Press,proto3" json:"Press,omitempty"`
	TopPhotoUri     string  `protobuf:"bytes,5,opt,name=TopPhotoUri,proto3" json:"TopPhotoUri,omitempty"`
	TopPhotoRedUri  string  `protobuf:"bytes,6,opt,name=TopPhotoRedUri,proto3" json:"TopPhotoRedUri,omitempty"`
	TopPhotoIRUri   string  `protobuf:"bytes,7,opt,name=TopPhotoIRUri,proto3" json:"TopPhotoIRUri,omitempty"`
	SidePhotoUri    string  `protobuf:"bytes,8,opt,name=SidePhotoUri,proto3" json:"SidePhotoUri,omitempty"`
	SidePhotoRedUri string  `protobuf:"bytes,9,opt,name=SidePhotoRedUri,proto3" json:"SidePhotoRedUri,omitempty"`
	SidePhotoIRUri  string  `protobuf:"bytes,10,opt,name=SidePhotoIRUri,proto3" json:"SidePhotoIRUri,omitempty"`
}

func (x *Datapoint) Reset() {
	*x = Datapoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Datapoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Datapoint) ProtoMessage() {}

func (x *Datapoint) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Datapoint.ProtoReflect.Descriptor instead.
func (*Datapoint) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{4}
}

func (x *Datapoint) GetXpos() int32 {
	if x != nil {
		return x.Xpos
	}
	return 0
}

func (x *Datapoint) GetTemp() float32 {
	if x != nil {
		return x.Temp
	}
	return 0
}

func (x *Datapoint) GetHum() float32 {
	if x != nil {
		return x.Hum
	}
	return 0
}

func (x *Datapoint) GetPress() float32 {
	if x != nil {
		return x.Press
	}
	return 0
}

func (x *Datapoint) GetTopPhotoUri() string {
	if x != nil {
		return x.TopPhotoUri
	}
	return ""
}

func (x *Datapoint) GetTopPhotoRedUri() string {
	if x != nil {
		return x.TopPhotoRedUri
	}
	return ""
}

func (x *Datapoint) GetTopPhotoIRUri() string {
	if x != nil {
		return x.TopPhotoIRUri
	}
	return ""
}

func (x *Datapoint) GetSidePhotoUri() string {
	if x != nil {
		return x.SidePhotoUri
	}
	return ""
}

func (x *Datapoint) GetSidePhotoRedUri() string {
	if x != nil {
		return x.SidePhotoRedUri
	}
	return ""
}

func (x *Datapoint) GetSidePhotoIRUri() string {
	if x != nil {
		return x.SidePhotoIRUri
	}
	return ""
}

type WateringResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta       *Meta                `protobuf:"bytes,1,opt,name=Meta,proto3" json:"Meta,omitempty"`
	Status     *Status              `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
	MessageId  string               `protobuf:"bytes,3,opt,name=MessageId,proto3" json:"MessageId,omitempty"`
	Type       string               `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	ActionId   string               `protobuf:"bytes,5,opt,name=ActionId,proto3" json:"ActionId,omitempty"`
	Timestamp  *timestamp.Timestamp `protobuf:"bytes,6,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Location   *Source              `protobuf:"bytes,7,opt,name=Location,proto3" json:"Location,omitempty"`
	Watering   *Watering            `protobuf:"bytes,8,opt,name=Watering,proto3" json:"Watering,omitempty"`
	Datapoints []*Datapoint         `protobuf:"bytes,9,rep,name=Datapoints,proto3" json:"Datapoints,omitempty"`
}

func (x *WateringResponse) Reset() {
	*x = WateringResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WateringResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WateringResponse) ProtoMessage() {}

func (x *WateringResponse) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WateringResponse.ProtoReflect.Descriptor instead.
func (*WateringResponse) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{5}
}

func (x *WateringResponse) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *WateringResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *WateringResponse) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *WateringResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *WateringResponse) GetActionId() string {
	if x != nil {
		return x.ActionId
	}
	return ""
}

func (x *WateringResponse) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *WateringResponse) GetLocation() *Source {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *WateringResponse) GetWatering() *Watering {
	if x != nil {
		return x.Watering
	}
	return nil
}

func (x *WateringResponse) GetDatapoints() []*Datapoint {
	if x != nil {
		return x.Datapoints
	}
	return nil
}

var File_events_proto protoreflect.FileDescriptor

var file_events_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x4a, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x4f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x6e, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x61, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x46, 0x61, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x65, 0x6c, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x65, 0x6c, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x61,
	0x79, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x53, 0x6c, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x53, 0x6c, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x46, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x52, 0x46, 0x49, 0x44, 0x22, 0xbe, 0x01, 0x0a, 0x08, 0x57, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x2c, 0x0a, 0x11, 0x50, 0x72, 0x65, 0x57, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x6e, 0x67, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x11, 0x50, 0x72, 0x65, 0x57, 0x61, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x57, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x50, 0x6f, 0x73, 0x74, 0x57, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x12, 0x50, 0x6f, 0x73, 0x74, 0x57, 0x61, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x57, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x79, 0x70, 0x65, 0x4f, 0x66, 0x4e, 0x6f, 0x73,
	0x73, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x54, 0x79, 0x70, 0x65, 0x4f,
	0x66, 0x4e, 0x6f, 0x73, 0x73, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x57, 0x61, 0x74, 0x65, 0x72,
	0x46, 0x61, 0x6e, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x57, 0x61, 0x74, 0x65, 0x72, 0x46, 0x61, 0x6e, 0x43, 0x6f,
	0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc1, 0x02, 0x0a, 0x09, 0x44, 0x61,
	0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x58, 0x70, 0x6f, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x58, 0x70, 0x6f, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x65, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x54, 0x65, 0x6d, 0x70, 0x12,
	0x10, 0x0a, 0x03, 0x48, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x48, 0x75,
	0x6d, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x50, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x6f, 0x70, 0x50, 0x68,
	0x6f, 0x74, 0x6f, 0x55, 0x72, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x6f,
	0x70, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x69, 0x12, 0x26, 0x0a, 0x0e, 0x54, 0x6f, 0x70,
	0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x64, 0x55, 0x72, 0x69, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x54, 0x6f, 0x70, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x64, 0x55, 0x72,
	0x69, 0x12, 0x24, 0x0a, 0x0d, 0x54, 0x6f, 0x70, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x49, 0x52, 0x55,
	0x72, 0x69, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x54, 0x6f, 0x70, 0x50, 0x68, 0x6f,
	0x74, 0x6f, 0x49, 0x52, 0x55, 0x72, 0x69, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x69, 0x64, 0x65, 0x50,
	0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x69, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x53,
	0x69, 0x64, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x69, 0x12, 0x28, 0x0a, 0x0f, 0x53,
	0x69, 0x64, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x64, 0x55, 0x72, 0x69, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x53, 0x69, 0x64, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52,
	0x65, 0x64, 0x55, 0x72, 0x69, 0x12, 0x26, 0x0a, 0x0e, 0x53, 0x69, 0x64, 0x65, 0x50, 0x68, 0x6f,
	0x74, 0x6f, 0x49, 0x52, 0x55, 0x72, 0x69, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x53,
	0x69, 0x64, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x49, 0x52, 0x55, 0x72, 0x69, 0x22, 0xf1, 0x02,
	0x0a, 0x10, 0x57, 0x61, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04,
	0x4d, 0x65, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x2a, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2c, 0x0a, 0x08, 0x57, 0x61, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x57, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x57, 0x61, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x31,
	0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x32, 0x9e, 0x01, 0x0a, 0x0f, 0x57, 0x61, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x57, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x0c, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x1a, 0x18, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x57, 0x61, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x49, 0x0a, 0x15, 0x50, 0x6f, 0x73, 0x74, 0x57,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73,
	0x12, 0x18, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x57, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_events_proto_rawDescOnce sync.Once
	file_events_proto_rawDescData = file_events_proto_rawDesc
)

func file_events_proto_rawDescGZIP() []byte {
	file_events_proto_rawDescOnce.Do(func() {
		file_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_events_proto_rawDescData)
	})
	return file_events_proto_rawDescData
}

var file_events_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_events_proto_goTypes = []interface{}{
	(*Meta)(nil),                // 0: events.Meta
	(*Status)(nil),              // 1: events.Status
	(*Source)(nil),              // 2: events.Source
	(*Watering)(nil),            // 3: events.Watering
	(*Datapoint)(nil),           // 4: events.Datapoint
	(*WateringResponse)(nil),    // 5: events.WateringResponse
	(*timestamp.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*empty.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_events_proto_depIdxs = []int32{
	0, // 0: events.WateringResponse.Meta:type_name -> events.Meta
	1, // 1: events.WateringResponse.Status:type_name -> events.Status
	6, // 2: events.WateringResponse.Timestamp:type_name -> google.protobuf.Timestamp
	2, // 3: events.WateringResponse.Location:type_name -> events.Source
	3, // 4: events.WateringResponse.Watering:type_name -> events.Watering
	4, // 5: events.WateringResponse.Datapoints:type_name -> events.Datapoint
	0, // 6: events.WateringService.GetWateringResponses:input_type -> events.Meta
	5, // 7: events.WateringService.PostWateringResponses:input_type -> events.WateringResponse
	5, // 8: events.WateringService.GetWateringResponses:output_type -> events.WateringResponse
	7, // 9: events.WateringService.PostWateringResponses:output_type -> google.protobuf.Empty
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_events_proto_init() }
func file_events_proto_init() {
	if File_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
		file_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Source); i {
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
		file_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Watering); i {
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
		file_events_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Datapoint); i {
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
		file_events_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WateringResponse); i {
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
			RawDescriptor: file_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_events_proto_goTypes,
		DependencyIndexes: file_events_proto_depIdxs,
		MessageInfos:      file_events_proto_msgTypes,
	}.Build()
	File_events_proto = out.File
	file_events_proto_rawDesc = nil
	file_events_proto_goTypes = nil
	file_events_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WateringServiceClient is the client API for WateringService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WateringServiceClient interface {
	GetWateringResponses(ctx context.Context, in *Meta, opts ...grpc.CallOption) (WateringService_GetWateringResponsesClient, error)
	PostWateringResponses(ctx context.Context, in *WateringResponse, opts ...grpc.CallOption) (*empty.Empty, error)
}

type wateringServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWateringServiceClient(cc grpc.ClientConnInterface) WateringServiceClient {
	return &wateringServiceClient{cc}
}

func (c *wateringServiceClient) GetWateringResponses(ctx context.Context, in *Meta, opts ...grpc.CallOption) (WateringService_GetWateringResponsesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_WateringService_serviceDesc.Streams[0], "/events.WateringService/GetWateringResponses", opts...)
	if err != nil {
		return nil, err
	}
	x := &wateringServiceGetWateringResponsesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WateringService_GetWateringResponsesClient interface {
	Recv() (*WateringResponse, error)
	grpc.ClientStream
}

type wateringServiceGetWateringResponsesClient struct {
	grpc.ClientStream
}

func (x *wateringServiceGetWateringResponsesClient) Recv() (*WateringResponse, error) {
	m := new(WateringResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *wateringServiceClient) PostWateringResponses(ctx context.Context, in *WateringResponse, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/events.WateringService/PostWateringResponses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WateringServiceServer is the server API for WateringService service.
type WateringServiceServer interface {
	GetWateringResponses(*Meta, WateringService_GetWateringResponsesServer) error
	PostWateringResponses(context.Context, *WateringResponse) (*empty.Empty, error)
}

// UnimplementedWateringServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWateringServiceServer struct {
}

func (*UnimplementedWateringServiceServer) GetWateringResponses(*Meta, WateringService_GetWateringResponsesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetWateringResponses not implemented")
}
func (*UnimplementedWateringServiceServer) PostWateringResponses(context.Context, *WateringResponse) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostWateringResponses not implemented")
}

func RegisterWateringServiceServer(s *grpc.Server, srv WateringServiceServer) {
	s.RegisterService(&_WateringService_serviceDesc, srv)
}

func _WateringService_GetWateringResponses_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Meta)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WateringServiceServer).GetWateringResponses(m, &wateringServiceGetWateringResponsesServer{stream})
}

type WateringService_GetWateringResponsesServer interface {
	Send(*WateringResponse) error
	grpc.ServerStream
}

type wateringServiceGetWateringResponsesServer struct {
	grpc.ServerStream
}

func (x *wateringServiceGetWateringResponsesServer) Send(m *WateringResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _WateringService_PostWateringResponses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WateringResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WateringServiceServer).PostWateringResponses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.WateringService/PostWateringResponses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WateringServiceServer).PostWateringResponses(ctx, req.(*WateringResponse))
	}
	return interceptor(ctx, in, info, handler)
}

var _WateringService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "events.WateringService",
	HandlerType: (*WateringServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostWateringResponses",
			Handler:    _WateringService_PostWateringResponses_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetWateringResponses",
			Handler:       _WateringService_GetWateringResponses_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "events.proto",
}

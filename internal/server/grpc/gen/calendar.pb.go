// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.4
// source: calendar.proto

package gen

import (
	reflect "reflect"
	sync "sync"

	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListEventsRequest_Agenda int32

const (
	ListEventsRequest_DAILY   ListEventsRequest_Agenda = 0
	ListEventsRequest_WEEKLY  ListEventsRequest_Agenda = 1
	ListEventsRequest_MONTHLY ListEventsRequest_Agenda = 2
)

// Enum value maps for ListEventsRequest_Agenda.
var (
	ListEventsRequest_Agenda_name = map[int32]string{
		0: "DAILY",
		1: "WEEKLY",
		2: "MONTHLY",
	}
	ListEventsRequest_Agenda_value = map[string]int32{
		"DAILY":   0,
		"WEEKLY":  1,
		"MONTHLY": 2,
	}
)

func (x ListEventsRequest_Agenda) Enum() *ListEventsRequest_Agenda {
	p := new(ListEventsRequest_Agenda)
	*p = x
	return p
}

func (x ListEventsRequest_Agenda) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListEventsRequest_Agenda) Descriptor() protoreflect.EnumDescriptor {
	return file_calendar_proto_enumTypes[0].Descriptor()
}

func (ListEventsRequest_Agenda) Type() protoreflect.EnumType {
	return &file_calendar_proto_enumTypes[0]
}

func (x ListEventsRequest_Agenda) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListEventsRequest_Agenda.Descriptor instead.
func (ListEventsRequest_Agenda) EnumDescriptor() ([]byte, []int) {
	return file_calendar_proto_rawDescGZIP(), []int{2, 0}
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	StartsAt    *timestamp.Timestamp `protobuf:"bytes,3,opt,name=starts_at,json=startsAt,proto3" json:"starts_at,omitempty"`
	EndsAt      *timestamp.Timestamp `protobuf:"bytes,4,opt,name=ends_at,json=endsAt,proto3" json:"ends_at,omitempty"`
	Notes       string               `protobuf:"bytes,5,opt,name=notes,proto3" json:"notes,omitempty"`
	OwnerId     string               `protobuf:"bytes,6,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	AlertBefore int64                `protobuf:"varint,7,opt,name=alert_before,json=alertBefore,proto3" json:"alert_before,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_calendar_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Event) GetStartsAt() *timestamp.Timestamp {
	if x != nil {
		return x.StartsAt
	}
	return nil
}

func (x *Event) GetEndsAt() *timestamp.Timestamp {
	if x != nil {
		return x.EndsAt
	}
	return nil
}

func (x *Event) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *Event) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *Event) GetAlertBefore() int64 {
	if x != nil {
		return x.AlertBefore
	}
	return 0
}

type EventId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EventId) Reset() {
	*x = EventId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventId) ProtoMessage() {}

func (x *EventId) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventId.ProtoReflect.Descriptor instead.
func (*EventId) Descriptor() ([]byte, []int) {
	return file_calendar_proto_rawDescGZIP(), []int{1}
}

func (x *EventId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Agenda    ListEventsRequest_Agenda `protobuf:"varint,1,opt,name=agenda,proto3,enum=calendar.ListEventsRequest_Agenda" json:"agenda,omitempty"`
	StartFrom *timestamp.Timestamp     `protobuf:"bytes,2,opt,name=start_from,json=startFrom,proto3" json:"start_from,omitempty"`
}

func (x *ListEventsRequest) Reset() {
	*x = ListEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventsRequest) ProtoMessage() {}

func (x *ListEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEventsRequest.ProtoReflect.Descriptor instead.
func (*ListEventsRequest) Descriptor() ([]byte, []int) {
	return file_calendar_proto_rawDescGZIP(), []int{2}
}

func (x *ListEventsRequest) GetAgenda() ListEventsRequest_Agenda {
	if x != nil {
		return x.Agenda
	}
	return ListEventsRequest_DAILY
}

func (x *ListEventsRequest) GetStartFrom() *timestamp.Timestamp {
	if x != nil {
		return x.StartFrom
	}
	return nil
}

type ListEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *ListEventsResponse) Reset() {
	*x = ListEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventsResponse) ProtoMessage() {}

func (x *ListEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEventsResponse.ProtoReflect.Descriptor instead.
func (*ListEventsResponse) Descriptor() ([]byte, []int) {
	return file_calendar_proto_rawDescGZIP(), []int{3}
}

func (x *ListEventsResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_calendar_proto protoreflect.FileDescriptor

var file_calendar_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x41,
	0x74, 0x12, 0x33, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06,
	0x65, 0x6e, 0x64, 0x73, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x22, 0x19, 0x0a, 0x07, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb8, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x06, 0x61,
	0x67, 0x65, 0x6e, 0x64, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x61,
	0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x52,
	0x06, 0x61, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x46, 0x72,
	0x6f, 0x6d, 0x22, 0x2c, 0x0a, 0x06, 0x41, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x12, 0x09, 0x0a, 0x05,
	0x44, 0x41, 0x49, 0x4c, 0x59, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x57, 0x45, 0x45, 0x4b, 0x4c,
	0x59, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x4c, 0x59, 0x10, 0x02,
	0x22, 0x3d, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32,
	0xfa, 0x01, 0x0a, 0x0f, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x11, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x08, 0x50, 0x75, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x0f, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x1a, 0x0f, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x49, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x1b, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63,
	0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x4f, 0x5a, 0x4d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x6c, 0x61, 0x64, 0x4e,
	0x46, 0x2f, 0x6f, 0x74, 0x75, 0x73, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x68, 0x77,
	0x31, 0x32, 0x5f, 0x31, 0x33, 0x5f, 0x31, 0x34, 0x5f, 0x31, 0x35, 0x5f, 0x63, 0x61, 0x6c, 0x65,
	0x6e, 0x64, 0x61, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calendar_proto_rawDescOnce sync.Once
	file_calendar_proto_rawDescData = file_calendar_proto_rawDesc
)

func file_calendar_proto_rawDescGZIP() []byte {
	file_calendar_proto_rawDescOnce.Do(func() {
		file_calendar_proto_rawDescData = protoimpl.X.CompressGZIP(file_calendar_proto_rawDescData)
	})
	return file_calendar_proto_rawDescData
}

var (
	file_calendar_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_calendar_proto_msgTypes  = make([]protoimpl.MessageInfo, 4)
	file_calendar_proto_goTypes   = []interface{}{
		(ListEventsRequest_Agenda)(0), // 0: calendar.ListEventsRequest.Agenda
		(*Event)(nil),                 // 1: calendar.Event
		(*EventId)(nil),               // 2: calendar.EventId
		(*ListEventsRequest)(nil),     // 3: calendar.ListEventsRequest
		(*ListEventsResponse)(nil),    // 4: calendar.ListEventsResponse
		(*timestamp.Timestamp)(nil),   // 5: google.protobuf.Timestamp
		(*empty.Empty)(nil),           // 6: google.protobuf.Empty
	}
)

var file_calendar_proto_depIdxs = []int32{
	5, // 0: calendar.Event.starts_at:type_name -> google.protobuf.Timestamp
	5, // 1: calendar.Event.ends_at:type_name -> google.protobuf.Timestamp
	0, // 2: calendar.ListEventsRequest.agenda:type_name -> calendar.ListEventsRequest.Agenda
	5, // 3: calendar.ListEventsRequest.start_from:type_name -> google.protobuf.Timestamp
	1, // 4: calendar.ListEventsResponse.events:type_name -> calendar.Event
	2, // 5: calendar.CalendarService.GetEvent:input_type -> calendar.EventId
	1, // 6: calendar.CalendarService.PutEvent:input_type -> calendar.Event
	2, // 7: calendar.CalendarService.DeleteEvent:input_type -> calendar.EventId
	3, // 8: calendar.CalendarService.ListEvents:input_type -> calendar.ListEventsRequest
	1, // 9: calendar.CalendarService.GetEvent:output_type -> calendar.Event
	1, // 10: calendar.CalendarService.PutEvent:output_type -> calendar.Event
	6, // 11: calendar.CalendarService.DeleteEvent:output_type -> google.protobuf.Empty
	4, // 12: calendar.CalendarService.ListEvents:output_type -> calendar.ListEventsResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_calendar_proto_init() }
func file_calendar_proto_init() {
	if File_calendar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calendar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_calendar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventId); i {
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
		file_calendar_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEventsRequest); i {
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
		file_calendar_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEventsResponse); i {
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
			RawDescriptor: file_calendar_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calendar_proto_goTypes,
		DependencyIndexes: file_calendar_proto_depIdxs,
		EnumInfos:         file_calendar_proto_enumTypes,
		MessageInfos:      file_calendar_proto_msgTypes,
	}.Build()
	File_calendar_proto = out.File
	file_calendar_proto_rawDesc = nil
	file_calendar_proto_goTypes = nil
	file_calendar_proto_depIdxs = nil
}

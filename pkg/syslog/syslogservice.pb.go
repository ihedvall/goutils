// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: syslogservice.proto

package syslog

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Severity int32

const (
	Severity_Emergency     Severity = 0 ///< System is unusable.
	Severity_Alert         Severity = 1 ///< Action must be taken immediately.
	Severity_Critical      Severity = 2 ///< Critical condition.
	Severity_Error         Severity = 3 ///< Error condition.
	Severity_Warning       Severity = 4 ///< Warning condition.
	Severity_Notice        Severity = 5 ///< Normal but significant condition.
	Severity_Informational Severity = 6 ///< Informational message.
	Severity_Debug         Severity = 7 ///< Debug message.
)

// Enum value maps for Severity.
var (
	Severity_name = map[int32]string{
		0: "Emergency",
		1: "Alert",
		2: "Critical",
		3: "Error",
		4: "Warning",
		5: "Notice",
		6: "Informational",
		7: "Debug",
	}
	Severity_value = map[string]int32{
		"Emergency":     0,
		"Alert":         1,
		"Critical":      2,
		"Error":         3,
		"Warning":       4,
		"Notice":        5,
		"Informational": 6,
		"Debug":         7,
	}
)

func (x Severity) Enum() *Severity {
	p := new(Severity)
	*p = x
	return p
}

func (x Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_syslogservice_proto_enumTypes[0].Descriptor()
}

func (Severity) Type() protoreflect.EnumType {
	return &file_syslogservice_proto_enumTypes[0]
}

func (x Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Severity.Descriptor instead.
func (Severity) EnumDescriptor() ([]byte, []int) {
	return file_syslogservice_proto_rawDescGZIP(), []int{0}
}

type SyslogDataDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity    int64  `protobuf:"varint,1,opt,name=identity,proto3" json:"identity,omitempty"`                         ///< Unique ID of the data definition
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                  ///< Internal name of the value. Ident/name
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"` ///< Friendly name of the value.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`                    ///< Descriptive text
	Unit        string `protobuf:"bytes,5,opt,name=unit,proto3" json:"unit,omitempty"`                                  ///< Unit of measure
	Company     string `protobuf:"bytes,6,opt,name=company,proto3" json:"company,omitempty"`                            ///< Company name or enterprise ID, if available.
}

func (x *SyslogDataDefinition) Reset() {
	*x = SyslogDataDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syslogservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyslogDataDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyslogDataDefinition) ProtoMessage() {}

func (x *SyslogDataDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_syslogservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyslogDataDefinition.ProtoReflect.Descriptor instead.
func (*SyslogDataDefinition) Descriptor() ([]byte, []int) {
	return file_syslogservice_proto_rawDescGZIP(), []int{0}
}

func (x *SyslogDataDefinition) GetIdentity() int64 {
	if x != nil {
		return x.Identity
	}
	return 0
}

func (x *SyslogDataDefinition) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SyslogDataDefinition) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *SyslogDataDefinition) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SyslogDataDefinition) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *SyslogDataDefinition) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

type SyslogDataValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity int64  `protobuf:"varint,1,opt,name=identity,proto3" json:"identity,omitempty"` ///< Unique database ID of the structured data definition.
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`          ///< Value name (display name)
	Value    string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`        ///< Data value
	Unit     string `protobuf:"bytes,4,opt,name=unit,proto3" json:"unit,omitempty"`          ///< Unit
}

func (x *SyslogDataValue) Reset() {
	*x = SyslogDataValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syslogservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyslogDataValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyslogDataValue) ProtoMessage() {}

func (x *SyslogDataValue) ProtoReflect() protoreflect.Message {
	mi := &file_syslogservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyslogDataValue.ProtoReflect.Descriptor instead.
func (*SyslogDataValue) Descriptor() ([]byte, []int) {
	return file_syslogservice_proto_rawDescGZIP(), []int{1}
}

func (x *SyslogDataValue) GetIdentity() int64 {
	if x != nil {
		return x.Identity
	}
	return 0
}

func (x *SyslogDataValue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SyslogDataValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *SyslogDataValue) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

type SyslogMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity        int64                  `protobuf:"varint,1,opt,name=identity,proto3" json:"identity,omitempty"`                                     ///< Unique ID (database index) of the event
	Severity        Severity               `protobuf:"varint,2,opt,name=severity,proto3,enum=syslog.Severity" json:"severity,omitempty"`                ///< Severity code.
	Facility        uint32                 `protobuf:"varint,3,opt,name=facility,proto3" json:"facility,omitempty"`                                     ///< Facility code. Normally not used
	Timestamp       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                                    ///< Nano-second since 1970.
	Text            string                 `protobuf:"bytes,5,opt,name=text,proto3" json:"text,omitempty"`                                              ///< Event message text
	Hostname        string                 `protobuf:"bytes,6,opt,name=hostname,proto3" json:"hostname,omitempty"`                                      ///< Host name. Display name or name.
	ApplicationName string                 `protobuf:"bytes,7,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"` ///< Application name
	ProcessId       string                 `protobuf:"bytes,8,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`                   ///< Process ID (PID). Not really used.
	MessageId       string                 `protobuf:"bytes,9,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`                   ///< Message ID
	DataValues      []*SyslogDataValue     `protobuf:"bytes,10,rep,name=data_values,json=dataValues,proto3" json:"data_values,omitempty"`               ///< Extra data/value pair.
}

func (x *SyslogMessage) Reset() {
	*x = SyslogMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syslogservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyslogMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyslogMessage) ProtoMessage() {}

func (x *SyslogMessage) ProtoReflect() protoreflect.Message {
	mi := &file_syslogservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyslogMessage.ProtoReflect.Descriptor instead.
func (*SyslogMessage) Descriptor() ([]byte, []int) {
	return file_syslogservice_proto_rawDescGZIP(), []int{2}
}

func (x *SyslogMessage) GetIdentity() int64 {
	if x != nil {
		return x.Identity
	}
	return 0
}

func (x *SyslogMessage) GetSeverity() Severity {
	if x != nil {
		return x.Severity
	}
	return Severity_Emergency
}

func (x *SyslogMessage) GetFacility() uint32 {
	if x != nil {
		return x.Facility
	}
	return 0
}

func (x *SyslogMessage) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *SyslogMessage) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *SyslogMessage) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *SyslogMessage) GetApplicationName() string {
	if x != nil {
		return x.ApplicationName
	}
	return ""
}

func (x *SyslogMessage) GetProcessId() string {
	if x != nil {
		return x.ProcessId
	}
	return ""
}

func (x *SyslogMessage) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *SyslogMessage) GetDataValues() []*SyslogDataValue {
	if x != nil {
		return x.DataValues
	}
	return nil
}

type EventMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity  int64                  `protobuf:"varint,1,opt,name=identity,proto3" json:"identity,omitempty"`                      ///< Unique ID (database index) of the event
	Severity  Severity               `protobuf:"varint,2,opt,name=severity,proto3,enum=syslog.Severity" json:"severity,omitempty"` ///< Severity code.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                     ///< Nano-second since 1970.
	Text      string                 `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`                               ///< Event message text
}

func (x *EventMessage) Reset() {
	*x = EventMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syslogservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMessage) ProtoMessage() {}

func (x *EventMessage) ProtoReflect() protoreflect.Message {
	mi := &file_syslogservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventMessage.ProtoReflect.Descriptor instead.
func (*EventMessage) Descriptor() ([]byte, []int) {
	return file_syslogservice_proto_rawDescGZIP(), []int{3}
}

func (x *EventMessage) GetIdentity() int64 {
	if x != nil {
		return x.Identity
	}
	return 0
}

func (x *EventMessage) GetSeverity() Severity {
	if x != nil {
		return x.Severity
	}
	return Severity_Emergency
}

func (x *EventMessage) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *EventMessage) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type SyslogFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level      *Severity               `protobuf:"varint,1,opt,name=level,proto3,enum=syslog.Severity,oneof" json:"level,omitempty"` ///< Severity level. Lower severity is filtered out.
	Facility   *uint32                 `protobuf:"varint,2,opt,name=facility,proto3,oneof" json:"facility,omitempty"`                ///< Facility code.
	TextFilter string                  `protobuf:"bytes,3,opt,name=text_filter,json=textFilter,proto3" json:"text_filter,omitempty"` ///< Wild card filter of the text.
	FromTime   *timestamppb.Timestamp  `protobuf:"bytes,4,opt,name=from_time,json=fromTime,proto3,oneof" json:"from_time,omitempty"` ///< From time.
	ToTime     *timestamppb.Timestamp  `protobuf:"bytes,5,opt,name=to_time,json=toTime,proto3,oneof" json:"to_time,omitempty"`       ///< From time.
	Data       []*SyslogDataDefinition `protobuf:"bytes,6,rep,name=data,proto3" json:"data,omitempty"`                               ///< Shows these data definition values
	FromId     *int64                  `protobuf:"varint,7,opt,name=from_id,json=fromId,proto3,oneof" json:"from_id,omitempty"`      ///< If set. Filter out all lower id.
	Offset     uint64                  `protobuf:"varint,8,opt,name=offset,proto3" json:"offset,omitempty"`                          ///< From row offset.
	Count      *uint32                 `protobuf:"varint,9,opt,name=count,proto3,oneof" json:"count,omitempty"`                      ///< Number of rows returned.
}

func (x *SyslogFilter) Reset() {
	*x = SyslogFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syslogservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyslogFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyslogFilter) ProtoMessage() {}

func (x *SyslogFilter) ProtoReflect() protoreflect.Message {
	mi := &file_syslogservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyslogFilter.ProtoReflect.Descriptor instead.
func (*SyslogFilter) Descriptor() ([]byte, []int) {
	return file_syslogservice_proto_rawDescGZIP(), []int{4}
}

func (x *SyslogFilter) GetLevel() Severity {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return Severity_Emergency
}

func (x *SyslogFilter) GetFacility() uint32 {
	if x != nil && x.Facility != nil {
		return *x.Facility
	}
	return 0
}

func (x *SyslogFilter) GetTextFilter() string {
	if x != nil {
		return x.TextFilter
	}
	return ""
}

func (x *SyslogFilter) GetFromTime() *timestamppb.Timestamp {
	if x != nil {
		return x.FromTime
	}
	return nil
}

func (x *SyslogFilter) GetToTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ToTime
	}
	return nil
}

func (x *SyslogFilter) GetData() []*SyslogDataDefinition {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SyslogFilter) GetFromId() int64 {
	if x != nil && x.FromId != nil {
		return *x.FromId
	}
	return 0
}

func (x *SyslogFilter) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *SyslogFilter) GetCount() uint32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

type SyslogCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *SyslogCount) Reset() {
	*x = SyslogCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syslogservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyslogCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyslogCount) ProtoMessage() {}

func (x *SyslogCount) ProtoReflect() protoreflect.Message {
	mi := &file_syslogservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyslogCount.ProtoReflect.Descriptor instead.
func (*SyslogCount) Descriptor() ([]byte, []int) {
	return file_syslogservice_proto_rawDescGZIP(), []int{5}
}

func (x *SyslogCount) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_syslogservice_proto protoreflect.FileDescriptor

var file_syslogservice_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x14,
	0x53, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x44, 0x61, 0x74, 0x61, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x22, 0x6b, 0x0a, 0x0f, 0x53, 0x79, 0x73, 0x6c, 0x6f,
	0x67, 0x44, 0x61, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x6e, 0x69, 0x74, 0x22, 0x82, 0x03, 0x0a, 0x0d, 0x53, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x2c, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x73, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x53, 0x65,
	0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x66, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12,
	0x38, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x53, 0x79,
	0x73, 0x6c, 0x6f, 0x67, 0x44, 0x61, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x64,
	0x61, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0xa6, 0x01, 0x0a, 0x0c, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x73, 0x79, 0x73, 0x6c, 0x6f,
	0x67, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x22, 0xbf, 0x03, 0x0a, 0x0c, 0x53, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x10, 0x2e, 0x73, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x53, 0x65, 0x76, 0x65,
	0x72, 0x69, 0x74, 0x79, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x88, 0x01, 0x01,
	0x12, 0x1f, 0x0a, 0x08, 0x66, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x48, 0x01, 0x52, 0x08, 0x66, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x88, 0x01,
	0x01, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x78, 0x74, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x3c, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x48, 0x02, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x38, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x03, 0x52,
	0x06, 0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x79, 0x73, 0x6c, 0x6f,
	0x67, 0x2e, 0x53, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x44, 0x61, 0x74, 0x61, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x07,
	0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x48, 0x04, 0x52,
	0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0d, 0x48, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x66, 0x61, 0x63, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x74, 0x6f, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x23, 0x0a, 0x0b, 0x53, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2a, 0x74, 0x0a, 0x08, 0x53, 0x65, 0x76,
	0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e,
	0x63, 0x79, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x43, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x10, 0x02, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x61, 0x72, 0x6e,
	0x69, 0x6e, 0x67, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x10,
	0x05, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x65, 0x62, 0x75, 0x67, 0x10, 0x07, 0x32,
	0xd2, 0x02, 0x0a, 0x0d, 0x53, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x73, 0x79, 0x73, 0x6c,
	0x6f, 0x67, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x00, 0x12, 0x37, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x2e,
	0x73, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x53, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x1a, 0x13, 0x2e, 0x73, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x53, 0x79, 0x73,
	0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x2e,
	0x53, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x14, 0x2e, 0x73,
	0x79, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73,
	0x6c, 0x6f, 0x67, 0x12, 0x14, 0x2e, 0x73, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x53, 0x79, 0x73,
	0x6c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x73, 0x79, 0x73, 0x6c,
	0x6f, 0x67, 0x2e, 0x53, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x4e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x44,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x73, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x53, 0x79, 0x73, 0x6c,
	0x6f, 0x67, 0x44, 0x61, 0x74, 0x61, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x00, 0x30, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x73, 0x79, 0x73, 0x6c, 0x6f, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_syslogservice_proto_rawDescOnce sync.Once
	file_syslogservice_proto_rawDescData = file_syslogservice_proto_rawDesc
)

func file_syslogservice_proto_rawDescGZIP() []byte {
	file_syslogservice_proto_rawDescOnce.Do(func() {
		file_syslogservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_syslogservice_proto_rawDescData)
	})
	return file_syslogservice_proto_rawDescData
}

var file_syslogservice_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_syslogservice_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_syslogservice_proto_goTypes = []interface{}{
	(Severity)(0),                 // 0: syslog.Severity
	(*SyslogDataDefinition)(nil),  // 1: syslog.SyslogDataDefinition
	(*SyslogDataValue)(nil),       // 2: syslog.SyslogDataValue
	(*SyslogMessage)(nil),         // 3: syslog.SyslogMessage
	(*EventMessage)(nil),          // 4: syslog.EventMessage
	(*SyslogFilter)(nil),          // 5: syslog.SyslogFilter
	(*SyslogCount)(nil),           // 6: syslog.SyslogCount
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_syslogservice_proto_depIdxs = []int32{
	0,  // 0: syslog.SyslogMessage.severity:type_name -> syslog.Severity
	7,  // 1: syslog.SyslogMessage.timestamp:type_name -> google.protobuf.Timestamp
	2,  // 2: syslog.SyslogMessage.data_values:type_name -> syslog.SyslogDataValue
	0,  // 3: syslog.EventMessage.severity:type_name -> syslog.Severity
	7,  // 4: syslog.EventMessage.timestamp:type_name -> google.protobuf.Timestamp
	0,  // 5: syslog.SyslogFilter.level:type_name -> syslog.Severity
	7,  // 6: syslog.SyslogFilter.from_time:type_name -> google.protobuf.Timestamp
	7,  // 7: syslog.SyslogFilter.to_time:type_name -> google.protobuf.Timestamp
	1,  // 8: syslog.SyslogFilter.data:type_name -> syslog.SyslogDataDefinition
	8,  // 9: syslog.SyslogService.GetLastEvent:input_type -> google.protobuf.Empty
	5,  // 10: syslog.SyslogService.GetCount:input_type -> syslog.SyslogFilter
	5,  // 11: syslog.SyslogService.GetEvent:input_type -> syslog.SyslogFilter
	5,  // 12: syslog.SyslogService.GetSyslog:input_type -> syslog.SyslogFilter
	8,  // 13: syslog.SyslogService.GetDataDefinitions:input_type -> google.protobuf.Empty
	4,  // 14: syslog.SyslogService.GetLastEvent:output_type -> syslog.EventMessage
	6,  // 15: syslog.SyslogService.GetCount:output_type -> syslog.SyslogCount
	4,  // 16: syslog.SyslogService.GetEvent:output_type -> syslog.EventMessage
	3,  // 17: syslog.SyslogService.GetSyslog:output_type -> syslog.SyslogMessage
	1,  // 18: syslog.SyslogService.GetDataDefinitions:output_type -> syslog.SyslogDataDefinition
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_syslogservice_proto_init() }
func file_syslogservice_proto_init() {
	if File_syslogservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_syslogservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyslogDataDefinition); i {
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
		file_syslogservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyslogDataValue); i {
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
		file_syslogservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyslogMessage); i {
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
		file_syslogservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventMessage); i {
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
		file_syslogservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyslogFilter); i {
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
		file_syslogservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyslogCount); i {
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
	file_syslogservice_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_syslogservice_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_syslogservice_proto_goTypes,
		DependencyIndexes: file_syslogservice_proto_depIdxs,
		EnumInfos:         file_syslogservice_proto_enumTypes,
		MessageInfos:      file_syslogservice_proto_msgTypes,
	}.Build()
	File_syslogservice_proto = out.File
	file_syslogservice_proto_rawDesc = nil
	file_syslogservice_proto_goTypes = nil
	file_syslogservice_proto_depIdxs = nil
}

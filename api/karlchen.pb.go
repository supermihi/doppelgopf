// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: karlchen.proto

package api

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

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_karlchen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_karlchen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_karlchen_proto_rawDescGZIP(), []int{0}
}

type OkOrNot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *OkOrNot) Reset() {
	*x = OkOrNot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_karlchen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OkOrNot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OkOrNot) ProtoMessage() {}

func (x *OkOrNot) ProtoReflect() protoreflect.Message {
	mi := &file_karlchen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OkOrNot.ProtoReflect.Descriptor instead.
func (*OkOrNot) Descriptor() ([]byte, []int) {
	return file_karlchen_proto_rawDescGZIP(), []int{1}
}

func (x *OkOrNot) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_karlchen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_karlchen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_karlchen_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RegisterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Secret string `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *RegisterReply) Reset() {
	*x = RegisterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_karlchen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReply) ProtoMessage() {}

func (x *RegisterReply) ProtoReflect() protoreflect.Message {
	mi := &file_karlchen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReply.ProtoReflect.Descriptor instead.
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return file_karlchen_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RegisterReply) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type TableId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TableId) Reset() {
	*x = TableId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_karlchen_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableId) ProtoMessage() {}

func (x *TableId) ProtoReflect() protoreflect.Message {
	mi := &file_karlchen_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableId.ProtoReflect.Descriptor instead.
func (*TableId) Descriptor() ([]byte, []int) {
	return file_karlchen_proto_rawDescGZIP(), []int{4}
}

func (x *TableId) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type TableData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableId    string `protobuf:"bytes,1,opt,name=tableId,proto3" json:"tableId,omitempty"`
	Owner      string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	InviteCode string `protobuf:"bytes,3,opt,name=inviteCode,proto3" json:"inviteCode,omitempty"`
}

func (x *TableData) Reset() {
	*x = TableData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_karlchen_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableData) ProtoMessage() {}

func (x *TableData) ProtoReflect() protoreflect.Message {
	mi := &file_karlchen_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableData.ProtoReflect.Descriptor instead.
func (*TableData) Descriptor() ([]byte, []int) {
	return file_karlchen_proto_rawDescGZIP(), []int{5}
}

func (x *TableData) GetTableId() string {
	if x != nil {
		return x.TableId
	}
	return ""
}

func (x *TableData) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *TableData) GetInviteCode() string {
	if x != nil {
		return x.InviteCode
	}
	return ""
}

type TableList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tables []*TableData `protobuf:"bytes,1,rep,name=tables,proto3" json:"tables,omitempty"`
}

func (x *TableList) Reset() {
	*x = TableList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_karlchen_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableList) ProtoMessage() {}

func (x *TableList) ProtoReflect() protoreflect.Message {
	mi := &file_karlchen_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableList.ProtoReflect.Descriptor instead.
func (*TableList) Descriptor() ([]byte, []int) {
	return file_karlchen_proto_rawDescGZIP(), []int{6}
}

func (x *TableList) GetTables() []*TableData {
	if x != nil {
		return x.Tables
	}
	return nil
}

type JoinTableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableId    string `protobuf:"bytes,1,opt,name=tableId,proto3" json:"tableId,omitempty"`
	InviteCode string `protobuf:"bytes,2,opt,name=inviteCode,proto3" json:"inviteCode,omitempty"`
}

func (x *JoinTableRequest) Reset() {
	*x = JoinTableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_karlchen_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinTableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinTableRequest) ProtoMessage() {}

func (x *JoinTableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_karlchen_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinTableRequest.ProtoReflect.Descriptor instead.
func (*JoinTableRequest) Descriptor() ([]byte, []int) {
	return file_karlchen_proto_rawDescGZIP(), []int{7}
}

func (x *JoinTableRequest) GetTableId() string {
	if x != nil {
		return x.TableId
	}
	return ""
}

func (x *JoinTableRequest) GetInviteCode() string {
	if x != nil {
		return x.InviteCode
	}
	return ""
}

type EmptyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyReply) Reset() {
	*x = EmptyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_karlchen_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyReply) ProtoMessage() {}

func (x *EmptyReply) ProtoReflect() protoreflect.Message {
	mi := &file_karlchen_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyReply.ProtoReflect.Descriptor instead.
func (*EmptyReply) Descriptor() ([]byte, []int) {
	return file_karlchen_proto_rawDescGZIP(), []int{8}
}

var File_karlchen_proto protoreflect.FileDescriptor

var file_karlchen_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6b, 0x61, 0x72, 0x6c, 0x63, 0x68, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1f, 0x0a, 0x07, 0x4f, 0x6b, 0x4f, 0x72, 0x4e, 0x6f, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x25, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a,
	0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x1f, 0x0a, 0x07, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5b, 0x0a, 0x09, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x22, 0x33, 0x0a, 0x09, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x26, 0x0a, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x22, 0x4c, 0x0a, 0x10, 0x4a, 0x6f, 0x69,
	0x6e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xc5, 0x02, 0x0a, 0x0d, 0x4b, 0x61, 0x72, 0x6c, 0x63, 0x68,
	0x65, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x12, 0x36, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x2f, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x11, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x6b, 0x4f, 0x72, 0x4e, 0x6f, 0x74, 0x22, 0x00,
	0x12, 0x31, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x11,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x09, 0x4a, 0x6f, 0x69, 0x6e, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x07, 0x5a,
	0x05, 0x2e, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_karlchen_proto_rawDescOnce sync.Once
	file_karlchen_proto_rawDescData = file_karlchen_proto_rawDesc
)

func file_karlchen_proto_rawDescGZIP() []byte {
	file_karlchen_proto_rawDescOnce.Do(func() {
		file_karlchen_proto_rawDescData = protoimpl.X.CompressGZIP(file_karlchen_proto_rawDescData)
	})
	return file_karlchen_proto_rawDescData
}

var file_karlchen_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_karlchen_proto_goTypes = []interface{}{
	(*EmptyRequest)(nil),     // 0: api.EmptyRequest
	(*OkOrNot)(nil),          // 1: api.OkOrNot
	(*RegisterRequest)(nil),  // 2: api.RegisterRequest
	(*RegisterReply)(nil),    // 3: api.RegisterReply
	(*TableId)(nil),          // 4: api.TableId
	(*TableData)(nil),        // 5: api.TableData
	(*TableList)(nil),        // 6: api.TableList
	(*JoinTableRequest)(nil), // 7: api.JoinTableRequest
	(*EmptyReply)(nil),       // 8: api.EmptyReply
}
var file_karlchen_proto_depIdxs = []int32{
	5, // 0: api.TableList.tables:type_name -> api.TableData
	2, // 1: api.Karlchencloud.Register:input_type -> api.RegisterRequest
	0, // 2: api.Karlchencloud.CheckLogin:input_type -> api.EmptyRequest
	0, // 3: api.Karlchencloud.ListTables:input_type -> api.EmptyRequest
	0, // 4: api.Karlchencloud.CreateTable:input_type -> api.EmptyRequest
	4, // 5: api.Karlchencloud.StartTable:input_type -> api.TableId
	7, // 6: api.Karlchencloud.JoinTable:input_type -> api.JoinTableRequest
	3, // 7: api.Karlchencloud.Register:output_type -> api.RegisterReply
	1, // 8: api.Karlchencloud.CheckLogin:output_type -> api.OkOrNot
	6, // 9: api.Karlchencloud.ListTables:output_type -> api.TableList
	5, // 10: api.Karlchencloud.CreateTable:output_type -> api.TableData
	8, // 11: api.Karlchencloud.StartTable:output_type -> api.EmptyReply
	8, // 12: api.Karlchencloud.JoinTable:output_type -> api.EmptyReply
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_karlchen_proto_init() }
func file_karlchen_proto_init() {
	if File_karlchen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_karlchen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
		file_karlchen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OkOrNot); i {
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
		file_karlchen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_karlchen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReply); i {
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
		file_karlchen_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableId); i {
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
		file_karlchen_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableData); i {
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
		file_karlchen_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableList); i {
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
		file_karlchen_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinTableRequest); i {
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
		file_karlchen_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyReply); i {
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
			RawDescriptor: file_karlchen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_karlchen_proto_goTypes,
		DependencyIndexes: file_karlchen_proto_depIdxs,
		MessageInfos:      file_karlchen_proto_msgTypes,
	}.Build()
	File_karlchen_proto = out.File
	file_karlchen_proto_rawDesc = nil
	file_karlchen_proto_goTypes = nil
	file_karlchen_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KarlchencloudClient is the client API for Karlchencloud service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KarlchencloudClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
	CheckLogin(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*OkOrNot, error)
	ListTables(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*TableList, error)
	CreateTable(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*TableData, error)
	StartTable(ctx context.Context, in *TableId, opts ...grpc.CallOption) (*EmptyReply, error)
	JoinTable(ctx context.Context, in *JoinTableRequest, opts ...grpc.CallOption) (*EmptyReply, error)
}

type karlchencloudClient struct {
	cc grpc.ClientConnInterface
}

func NewKarlchencloudClient(cc grpc.ClientConnInterface) KarlchencloudClient {
	return &karlchencloudClient{cc}
}

func (c *karlchencloudClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, "/api.Karlchencloud/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *karlchencloudClient) CheckLogin(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*OkOrNot, error) {
	out := new(OkOrNot)
	err := c.cc.Invoke(ctx, "/api.Karlchencloud/CheckLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *karlchencloudClient) ListTables(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*TableList, error) {
	out := new(TableList)
	err := c.cc.Invoke(ctx, "/api.Karlchencloud/ListTables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *karlchencloudClient) CreateTable(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*TableData, error) {
	out := new(TableData)
	err := c.cc.Invoke(ctx, "/api.Karlchencloud/CreateTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *karlchencloudClient) StartTable(ctx context.Context, in *TableId, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, "/api.Karlchencloud/StartTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *karlchencloudClient) JoinTable(ctx context.Context, in *JoinTableRequest, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, "/api.Karlchencloud/JoinTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KarlchencloudServer is the server API for Karlchencloud service.
type KarlchencloudServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	CheckLogin(context.Context, *EmptyRequest) (*OkOrNot, error)
	ListTables(context.Context, *EmptyRequest) (*TableList, error)
	CreateTable(context.Context, *EmptyRequest) (*TableData, error)
	StartTable(context.Context, *TableId) (*EmptyReply, error)
	JoinTable(context.Context, *JoinTableRequest) (*EmptyReply, error)
}

// UnimplementedKarlchencloudServer can be embedded to have forward compatible implementations.
type UnimplementedKarlchencloudServer struct {
}

func (*UnimplementedKarlchencloudServer) Register(context.Context, *RegisterRequest) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedKarlchencloudServer) CheckLogin(context.Context, *EmptyRequest) (*OkOrNot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckLogin not implemented")
}
func (*UnimplementedKarlchencloudServer) ListTables(context.Context, *EmptyRequest) (*TableList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTables not implemented")
}
func (*UnimplementedKarlchencloudServer) CreateTable(context.Context, *EmptyRequest) (*TableData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTable not implemented")
}
func (*UnimplementedKarlchencloudServer) StartTable(context.Context, *TableId) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartTable not implemented")
}
func (*UnimplementedKarlchencloudServer) JoinTable(context.Context, *JoinTableRequest) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinTable not implemented")
}

func RegisterKarlchencloudServer(s *grpc.Server, srv KarlchencloudServer) {
	s.RegisterService(&_Karlchencloud_serviceDesc, srv)
}

func _Karlchencloud_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KarlchencloudServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Karlchencloud/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KarlchencloudServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Karlchencloud_CheckLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KarlchencloudServer).CheckLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Karlchencloud/CheckLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KarlchencloudServer).CheckLogin(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Karlchencloud_ListTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KarlchencloudServer).ListTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Karlchencloud/ListTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KarlchencloudServer).ListTables(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Karlchencloud_CreateTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KarlchencloudServer).CreateTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Karlchencloud/CreateTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KarlchencloudServer).CreateTable(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Karlchencloud_StartTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TableId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KarlchencloudServer).StartTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Karlchencloud/StartTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KarlchencloudServer).StartTable(ctx, req.(*TableId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Karlchencloud_JoinTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KarlchencloudServer).JoinTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Karlchencloud/JoinTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KarlchencloudServer).JoinTable(ctx, req.(*JoinTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Karlchencloud_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Karlchencloud",
	HandlerType: (*KarlchencloudServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Karlchencloud_Register_Handler,
		},
		{
			MethodName: "CheckLogin",
			Handler:    _Karlchencloud_CheckLogin_Handler,
		},
		{
			MethodName: "ListTables",
			Handler:    _Karlchencloud_ListTables_Handler,
		},
		{
			MethodName: "CreateTable",
			Handler:    _Karlchencloud_CreateTable_Handler,
		},
		{
			MethodName: "StartTable",
			Handler:    _Karlchencloud_StartTable_Handler,
		},
		{
			MethodName: "JoinTable",
			Handler:    _Karlchencloud_JoinTable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "karlchen.proto",
}
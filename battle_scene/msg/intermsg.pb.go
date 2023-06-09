// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: intermsg.proto

package msg

import (
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

type RequestInterServerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId   uint32 `protobuf:"varint,1,opt,name=ServerId,proto3" json:"ServerId,omitempty"`
	ServerType uint32 `protobuf:"varint,2,opt,name=ServerType,proto3" json:"ServerType,omitempty"`
	OnLineNum  uint32 `protobuf:"varint,3,opt,name=OnLineNum,proto3" json:"OnLineNum,omitempty"`
}

func (x *RequestInterServerInfo) Reset() {
	*x = RequestInterServerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_intermsg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestInterServerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestInterServerInfo) ProtoMessage() {}

func (x *RequestInterServerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_intermsg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestInterServerInfo.ProtoReflect.Descriptor instead.
func (*RequestInterServerInfo) Descriptor() ([]byte, []int) {
	return file_intermsg_proto_rawDescGZIP(), []int{0}
}

func (x *RequestInterServerInfo) GetServerId() uint32 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *RequestInterServerInfo) GetServerType() uint32 {
	if x != nil {
		return x.ServerType
	}
	return 0
}

func (x *RequestInterServerInfo) GetOnLineNum() uint32 {
	if x != nil {
		return x.OnLineNum
	}
	return 0
}

type ResponseInterServerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *ResponseInterServerInfo) Reset() {
	*x = ResponseInterServerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_intermsg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseInterServerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseInterServerInfo) ProtoMessage() {}

func (x *ResponseInterServerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_intermsg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseInterServerInfo.ProtoReflect.Descriptor instead.
func (*ResponseInterServerInfo) Descriptor() ([]byte, []int) {
	return file_intermsg_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseInterServerInfo) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type ReqeustInterGs2BattlHallEnter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId  int64  `protobuf:"varint,1,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	GsServerId uint32 `protobuf:"varint,2,opt,name=GsServerId,proto3" json:"GsServerId,omitempty"`
}

func (x *ReqeustInterGs2BattlHallEnter) Reset() {
	*x = ReqeustInterGs2BattlHallEnter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_intermsg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqeustInterGs2BattlHallEnter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqeustInterGs2BattlHallEnter) ProtoMessage() {}

func (x *ReqeustInterGs2BattlHallEnter) ProtoReflect() protoreflect.Message {
	mi := &file_intermsg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqeustInterGs2BattlHallEnter.ProtoReflect.Descriptor instead.
func (*ReqeustInterGs2BattlHallEnter) Descriptor() ([]byte, []int) {
	return file_intermsg_proto_rawDescGZIP(), []int{2}
}

func (x *ReqeustInterGs2BattlHallEnter) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *ReqeustInterGs2BattlHallEnter) GetGsServerId() uint32 {
	if x != nil {
		return x.GsServerId
	}
	return 0
}

type ResponseInterGs2BattlHallEnter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result        int32  `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
	SceneServerId uint32 `protobuf:"varint,2,opt,name=SceneServerId,proto3" json:"SceneServerId,omitempty"`
	AccountId     int64  `protobuf:"varint,3,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
}

func (x *ResponseInterGs2BattlHallEnter) Reset() {
	*x = ResponseInterGs2BattlHallEnter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_intermsg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseInterGs2BattlHallEnter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseInterGs2BattlHallEnter) ProtoMessage() {}

func (x *ResponseInterGs2BattlHallEnter) ProtoReflect() protoreflect.Message {
	mi := &file_intermsg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseInterGs2BattlHallEnter.ProtoReflect.Descriptor instead.
func (*ResponseInterGs2BattlHallEnter) Descriptor() ([]byte, []int) {
	return file_intermsg_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseInterGs2BattlHallEnter) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *ResponseInterGs2BattlHallEnter) GetSceneServerId() uint32 {
	if x != nil {
		return x.SceneServerId
	}
	return 0
}

func (x *ResponseInterGs2BattlHallEnter) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

type RequestInterBattlHall2SceneEnter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId  int64  `protobuf:"varint,1,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	GsServerId uint32 `protobuf:"varint,2,opt,name=GsServerId,proto3" json:"GsServerId,omitempty"`
}

func (x *RequestInterBattlHall2SceneEnter) Reset() {
	*x = RequestInterBattlHall2SceneEnter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_intermsg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestInterBattlHall2SceneEnter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestInterBattlHall2SceneEnter) ProtoMessage() {}

func (x *RequestInterBattlHall2SceneEnter) ProtoReflect() protoreflect.Message {
	mi := &file_intermsg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestInterBattlHall2SceneEnter.ProtoReflect.Descriptor instead.
func (*RequestInterBattlHall2SceneEnter) Descriptor() ([]byte, []int) {
	return file_intermsg_proto_rawDescGZIP(), []int{4}
}

func (x *RequestInterBattlHall2SceneEnter) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *RequestInterBattlHall2SceneEnter) GetGsServerId() uint32 {
	if x != nil {
		return x.GsServerId
	}
	return 0
}

type ResponseInterBattlHall2SceneEnter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result     int32  `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
	AccountId  int64  `protobuf:"varint,2,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	GsServerId uint32 `protobuf:"varint,3,opt,name=GsServerId,proto3" json:"GsServerId,omitempty"`
	SceneId    uint32 `protobuf:"varint,4,opt,name=SceneId,proto3" json:"SceneId,omitempty"`
}

func (x *ResponseInterBattlHall2SceneEnter) Reset() {
	*x = ResponseInterBattlHall2SceneEnter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_intermsg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseInterBattlHall2SceneEnter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseInterBattlHall2SceneEnter) ProtoMessage() {}

func (x *ResponseInterBattlHall2SceneEnter) ProtoReflect() protoreflect.Message {
	mi := &file_intermsg_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseInterBattlHall2SceneEnter.ProtoReflect.Descriptor instead.
func (*ResponseInterBattlHall2SceneEnter) Descriptor() ([]byte, []int) {
	return file_intermsg_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseInterBattlHall2SceneEnter) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *ResponseInterBattlHall2SceneEnter) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *ResponseInterBattlHall2SceneEnter) GetGsServerId() uint32 {
	if x != nil {
		return x.GsServerId
	}
	return 0
}

func (x *ResponseInterBattlHall2SceneEnter) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

type ReqeustInterGs2BattlHallLeave struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId     int64  `protobuf:"varint,1,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	GsServerId    uint32 `protobuf:"varint,2,opt,name=GsServerId,proto3" json:"GsServerId,omitempty"`
	SceneServerId uint32 `protobuf:"varint,3,opt,name=SceneServerId,proto3" json:"SceneServerId,omitempty"`
}

func (x *ReqeustInterGs2BattlHallLeave) Reset() {
	*x = ReqeustInterGs2BattlHallLeave{}
	if protoimpl.UnsafeEnabled {
		mi := &file_intermsg_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqeustInterGs2BattlHallLeave) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqeustInterGs2BattlHallLeave) ProtoMessage() {}

func (x *ReqeustInterGs2BattlHallLeave) ProtoReflect() protoreflect.Message {
	mi := &file_intermsg_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqeustInterGs2BattlHallLeave.ProtoReflect.Descriptor instead.
func (*ReqeustInterGs2BattlHallLeave) Descriptor() ([]byte, []int) {
	return file_intermsg_proto_rawDescGZIP(), []int{6}
}

func (x *ReqeustInterGs2BattlHallLeave) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *ReqeustInterGs2BattlHallLeave) GetGsServerId() uint32 {
	if x != nil {
		return x.GsServerId
	}
	return 0
}

func (x *ReqeustInterGs2BattlHallLeave) GetSceneServerId() uint32 {
	if x != nil {
		return x.SceneServerId
	}
	return 0
}

type ResponseInterGs2BattlHallLeave struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result        int32  `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
	AccountId     int64  `protobuf:"varint,2,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	SceneServerID uint32 `protobuf:"varint,3,opt,name=SceneServerID,proto3" json:"SceneServerID,omitempty"`
}

func (x *ResponseInterGs2BattlHallLeave) Reset() {
	*x = ResponseInterGs2BattlHallLeave{}
	if protoimpl.UnsafeEnabled {
		mi := &file_intermsg_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseInterGs2BattlHallLeave) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseInterGs2BattlHallLeave) ProtoMessage() {}

func (x *ResponseInterGs2BattlHallLeave) ProtoReflect() protoreflect.Message {
	mi := &file_intermsg_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseInterGs2BattlHallLeave.ProtoReflect.Descriptor instead.
func (*ResponseInterGs2BattlHallLeave) Descriptor() ([]byte, []int) {
	return file_intermsg_proto_rawDescGZIP(), []int{7}
}

func (x *ResponseInterGs2BattlHallLeave) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *ResponseInterGs2BattlHallLeave) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *ResponseInterGs2BattlHallLeave) GetSceneServerID() uint32 {
	if x != nil {
		return x.SceneServerID
	}
	return 0
}

type RequestInterBattlHall2SceneLeave struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId     int64  `protobuf:"varint,1,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	GsServerId    uint32 `protobuf:"varint,2,opt,name=GsServerId,proto3" json:"GsServerId,omitempty"`
	SceneServerId uint32 `protobuf:"varint,3,opt,name=SceneServerId,proto3" json:"SceneServerId,omitempty"`
}

func (x *RequestInterBattlHall2SceneLeave) Reset() {
	*x = RequestInterBattlHall2SceneLeave{}
	if protoimpl.UnsafeEnabled {
		mi := &file_intermsg_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestInterBattlHall2SceneLeave) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestInterBattlHall2SceneLeave) ProtoMessage() {}

func (x *RequestInterBattlHall2SceneLeave) ProtoReflect() protoreflect.Message {
	mi := &file_intermsg_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestInterBattlHall2SceneLeave.ProtoReflect.Descriptor instead.
func (*RequestInterBattlHall2SceneLeave) Descriptor() ([]byte, []int) {
	return file_intermsg_proto_rawDescGZIP(), []int{8}
}

func (x *RequestInterBattlHall2SceneLeave) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *RequestInterBattlHall2SceneLeave) GetGsServerId() uint32 {
	if x != nil {
		return x.GsServerId
	}
	return 0
}

func (x *RequestInterBattlHall2SceneLeave) GetSceneServerId() uint32 {
	if x != nil {
		return x.SceneServerId
	}
	return 0
}

type ResponseInterBattlHall2SceneLeave struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result        int32  `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
	AccountId     int64  `protobuf:"varint,2,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	GsServerId    uint32 `protobuf:"varint,3,opt,name=GsServerId,proto3" json:"GsServerId,omitempty"`
	SceneServerId uint32 `protobuf:"varint,4,opt,name=SceneServerId,proto3" json:"SceneServerId,omitempty"`
}

func (x *ResponseInterBattlHall2SceneLeave) Reset() {
	*x = ResponseInterBattlHall2SceneLeave{}
	if protoimpl.UnsafeEnabled {
		mi := &file_intermsg_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseInterBattlHall2SceneLeave) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseInterBattlHall2SceneLeave) ProtoMessage() {}

func (x *ResponseInterBattlHall2SceneLeave) ProtoReflect() protoreflect.Message {
	mi := &file_intermsg_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseInterBattlHall2SceneLeave.ProtoReflect.Descriptor instead.
func (*ResponseInterBattlHall2SceneLeave) Descriptor() ([]byte, []int) {
	return file_intermsg_proto_rawDescGZIP(), []int{9}
}

func (x *ResponseInterBattlHall2SceneLeave) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *ResponseInterBattlHall2SceneLeave) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *ResponseInterBattlHall2SceneLeave) GetGsServerId() uint32 {
	if x != nil {
		return x.GsServerId
	}
	return 0
}

func (x *ResponseInterBattlHall2SceneLeave) GetSceneServerId() uint32 {
	if x != nil {
		return x.SceneServerId
	}
	return 0
}

var File_intermsg_proto protoreflect.FileDescriptor

var file_intermsg_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x72, 0x0a, 0x16, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1a, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4f,
	0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x4f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x22, 0x31, 0x0a, 0x17, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x5d, 0x0a, 0x1d,
	0x52, 0x65, 0x71, 0x65, 0x75, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x47, 0x73, 0x32, 0x42,
	0x61, 0x74, 0x74, 0x6c, 0x48, 0x61, 0x6c, 0x6c, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x47,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x47, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x7c, 0x0a, 0x1e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x47, 0x73, 0x32, 0x42,
	0x61, 0x74, 0x74, 0x6c, 0x48, 0x61, 0x6c, 0x6c, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x53, 0x63,
	0x65, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x60, 0x0a, 0x20, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x48, 0x61,
	0x6c, 0x6c, 0x32, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x47,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x47, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x93, 0x01, 0x0a, 0x21,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x42, 0x61, 0x74,
	0x74, 0x6c, 0x48, 0x61, 0x6c, 0x6c, 0x32, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x47, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x47, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x63, 0x65, 0x6e, 0x65,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49,
	0x64, 0x22, 0x83, 0x01, 0x0a, 0x1d, 0x52, 0x65, 0x71, 0x65, 0x75, 0x73, 0x74, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x47, 0x73, 0x32, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x48, 0x61, 0x6c, 0x6c, 0x4c, 0x65,
	0x61, 0x76, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x47, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x47, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x7c, 0x0a, 0x1e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x47, 0x73, 0x32, 0x42, 0x61, 0x74, 0x74, 0x6c,
	0x48, 0x61, 0x6c, 0x6c, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x24, 0x0a, 0x0d, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x44, 0x22, 0x86, 0x01, 0x0a, 0x20, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x48, 0x61, 0x6c, 0x6c, 0x32,
	0x53, 0x63, 0x65, 0x6e, 0x65, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x47, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x47, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x63, 0x65, 0x6e,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0d, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x9f,
	0x01, 0x0a, 0x21, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x42, 0x61, 0x74, 0x74, 0x6c, 0x48, 0x61, 0x6c, 0x6c, 0x32, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4c,
	0x65, 0x61, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x47, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x47, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x63,
	0x65, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0d, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_intermsg_proto_rawDescOnce sync.Once
	file_intermsg_proto_rawDescData = file_intermsg_proto_rawDesc
)

func file_intermsg_proto_rawDescGZIP() []byte {
	file_intermsg_proto_rawDescOnce.Do(func() {
		file_intermsg_proto_rawDescData = protoimpl.X.CompressGZIP(file_intermsg_proto_rawDescData)
	})
	return file_intermsg_proto_rawDescData
}

var file_intermsg_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_intermsg_proto_goTypes = []interface{}{
	(*RequestInterServerInfo)(nil),            // 0: msg.RequestInterServerInfo
	(*ResponseInterServerInfo)(nil),           // 1: msg.ResponseInterServerInfo
	(*ReqeustInterGs2BattlHallEnter)(nil),     // 2: msg.ReqeustInterGs2BattlHallEnter
	(*ResponseInterGs2BattlHallEnter)(nil),    // 3: msg.ResponseInterGs2BattlHallEnter
	(*RequestInterBattlHall2SceneEnter)(nil),  // 4: msg.RequestInterBattlHall2SceneEnter
	(*ResponseInterBattlHall2SceneEnter)(nil), // 5: msg.ResponseInterBattlHall2SceneEnter
	(*ReqeustInterGs2BattlHallLeave)(nil),     // 6: msg.ReqeustInterGs2BattlHallLeave
	(*ResponseInterGs2BattlHallLeave)(nil),    // 7: msg.ResponseInterGs2BattlHallLeave
	(*RequestInterBattlHall2SceneLeave)(nil),  // 8: msg.RequestInterBattlHall2SceneLeave
	(*ResponseInterBattlHall2SceneLeave)(nil), // 9: msg.ResponseInterBattlHall2SceneLeave
}
var file_intermsg_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_intermsg_proto_init() }
func file_intermsg_proto_init() {
	if File_intermsg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_intermsg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestInterServerInfo); i {
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
		file_intermsg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseInterServerInfo); i {
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
		file_intermsg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqeustInterGs2BattlHallEnter); i {
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
		file_intermsg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseInterGs2BattlHallEnter); i {
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
		file_intermsg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestInterBattlHall2SceneEnter); i {
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
		file_intermsg_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseInterBattlHall2SceneEnter); i {
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
		file_intermsg_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqeustInterGs2BattlHallLeave); i {
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
		file_intermsg_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseInterGs2BattlHallLeave); i {
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
		file_intermsg_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestInterBattlHall2SceneLeave); i {
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
		file_intermsg_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseInterBattlHall2SceneLeave); i {
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
			RawDescriptor: file_intermsg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_intermsg_proto_goTypes,
		DependencyIndexes: file_intermsg_proto_depIdxs,
		MessageInfos:      file_intermsg_proto_msgTypes,
	}.Build()
	File_intermsg_proto = out.File
	file_intermsg_proto_rawDesc = nil
	file_intermsg_proto_goTypes = nil
	file_intermsg_proto_depIdxs = nil
}

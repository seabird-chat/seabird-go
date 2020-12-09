// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: common.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// When sent by a chat backend, the ID MAY optionally be used as a hint to the
// frontend to get the same external ID on reconnection. When sent by core to a
// plugin, the ID MUST be an ID pointing to an exact backend instance.
type Backend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Backend) Reset() {
	*x = Backend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Backend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Backend) ProtoMessage() {}

func (x *Backend) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Backend.ProtoReflect.Descriptor instead.
func (*Backend) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *Backend) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Backend) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

// When sent by a chat backend, the ID MUST be unique across users in that running
// backend. Failure to follow this will result in incorrect information being
// sent to plugins. When sent by core to a plugin, the ID MUST be an ID pointing
// to an instance of a user on chat backend.
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

// When sent by a chat backend, the ID MUST be unique across channels in that
// running backend. Failure to follow this will result in incorrect information
// being sent to plugins. When sent by core to a plugin, the ID MUST be an ID
// pointing to an instance of a channel on chat backend.
type Channel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Topic       string `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (x *Channel) Reset() {
	*x = Channel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Channel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Channel) ProtoMessage() {}

func (x *Channel) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Channel.ProtoReflect.Descriptor instead.
func (*Channel) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *Channel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Channel) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Channel) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

// ChannelSource represents where a message came from. Note that User is used
// rather than just an ID so any clients can have access to the name as well.
// This may change in the future.
type ChannelSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	User      *User  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ChannelSource) Reset() {
	*x = ChannelSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelSource) ProtoMessage() {}

func (x *ChannelSource) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelSource.ProtoReflect.Descriptor instead.
func (*ChannelSource) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *ChannelSource) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *ChannelSource) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

// MessageEvent will be sent when a user sends a message to a channel.
type MessageEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source *ChannelSource `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Text   string         `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *MessageEvent) Reset() {
	*x = MessageEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageEvent) ProtoMessage() {}

func (x *MessageEvent) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageEvent.ProtoReflect.Descriptor instead.
func (*MessageEvent) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4}
}

func (x *MessageEvent) GetSource() *ChannelSource {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *MessageEvent) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// PrivateMessageEvent will be sent when a user sends a message directly to the
// chat backend.
type PrivateMessageEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source *User  `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Text   string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *PrivateMessageEvent) Reset() {
	*x = PrivateMessageEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateMessageEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateMessageEvent) ProtoMessage() {}

func (x *PrivateMessageEvent) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateMessageEvent.ProtoReflect.Descriptor instead.
func (*PrivateMessageEvent) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{5}
}

func (x *PrivateMessageEvent) GetSource() *User {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *PrivateMessageEvent) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// MentionEvent will be sent when a user mentions the chat backend's user at the
// start of a message in a channel. Note that this may eventually go away when a
// more general message format is developed.
type MentionEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source *ChannelSource `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Text   string         `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *MentionEvent) Reset() {
	*x = MentionEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MentionEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MentionEvent) ProtoMessage() {}

func (x *MentionEvent) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MentionEvent.ProtoReflect.Descriptor instead.
func (*MentionEvent) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{6}
}

func (x *MentionEvent) GetSource() *ChannelSource {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *MentionEvent) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// CommandEvent will be sent when a user issues a command in a channel. Commands
// cannot be issued via private message. If a message is parsed as a command, it
// MUST NOT be also sent as another message type.
type CommandEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source  *ChannelSource `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Command string         `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	Arg     string         `protobuf:"bytes,3,opt,name=arg,proto3" json:"arg,omitempty"`
}

func (x *CommandEvent) Reset() {
	*x = CommandEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandEvent) ProtoMessage() {}

func (x *CommandEvent) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandEvent.ProtoReflect.Descriptor instead.
func (*CommandEvent) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{7}
}

func (x *CommandEvent) GetSource() *ChannelSource {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *CommandEvent) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *CommandEvent) GetArg() string {
	if x != nil {
		return x.Arg
	}
	return ""
}

// ActionEvent will be sent when a user takes an action in a channel. This is
// often triggered with /me.
type ActionEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source *ChannelSource `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Text   string         `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *ActionEvent) Reset() {
	*x = ActionEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionEvent) ProtoMessage() {}

func (x *ActionEvent) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionEvent.ProtoReflect.Descriptor instead.
func (*ActionEvent) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{8}
}

func (x *ActionEvent) GetSource() *ChannelSource {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *ActionEvent) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// PrivateActionEvent will be sent when a user takes an action in a private
// message. This is often triggered with /me.
type PrivateActionEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source *User  `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Text   string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *PrivateActionEvent) Reset() {
	*x = PrivateActionEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateActionEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateActionEvent) ProtoMessage() {}

func (x *PrivateActionEvent) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateActionEvent.ProtoReflect.Descriptor instead.
func (*PrivateActionEvent) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{9}
}

func (x *PrivateActionEvent) GetSource() *User {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *PrivateActionEvent) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x2d, 0x0a, 0x07, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x39, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x52, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x22, 0x50, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x51, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x4f, 0x0a, 0x13, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x24, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x51, 0x0a, 0x0c, 0x4d, 0x65,
	0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x69, 0x0a,
	0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2d, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x72, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x72, 0x67, 0x22, 0x50, 0x0a, 0x0b, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x4e, 0x0a, 0x12, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x24, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x42, 0x1f, 0x0a, 0x17, 0x69, 0x6f,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x2e, 0x73, 0x65, 0x61, 0x62, 0x69, 0x72, 0x64, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_common_proto_goTypes = []interface{}{
	(*Backend)(nil),             // 0: common.Backend
	(*User)(nil),                // 1: common.User
	(*Channel)(nil),             // 2: common.Channel
	(*ChannelSource)(nil),       // 3: common.ChannelSource
	(*MessageEvent)(nil),        // 4: common.MessageEvent
	(*PrivateMessageEvent)(nil), // 5: common.PrivateMessageEvent
	(*MentionEvent)(nil),        // 6: common.MentionEvent
	(*CommandEvent)(nil),        // 7: common.CommandEvent
	(*ActionEvent)(nil),         // 8: common.ActionEvent
	(*PrivateActionEvent)(nil),  // 9: common.PrivateActionEvent
}
var file_common_proto_depIdxs = []int32{
	1, // 0: common.ChannelSource.user:type_name -> common.User
	3, // 1: common.MessageEvent.source:type_name -> common.ChannelSource
	1, // 2: common.PrivateMessageEvent.source:type_name -> common.User
	3, // 3: common.MentionEvent.source:type_name -> common.ChannelSource
	3, // 4: common.CommandEvent.source:type_name -> common.ChannelSource
	3, // 5: common.ActionEvent.source:type_name -> common.ChannelSource
	1, // 6: common.PrivateActionEvent.source:type_name -> common.User
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Backend); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Channel); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelSource); i {
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
		file_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageEvent); i {
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
		file_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateMessageEvent); i {
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
		file_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MentionEvent); i {
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
		file_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandEvent); i {
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
		file_common_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionEvent); i {
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
		file_common_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateActionEvent); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: github.com/aperturerobotics/bifrost/pubsub/grpc/grpc.proto

package pubsub_grpc

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

// SubcribeRequest is a pubsub subscription request message.
type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ChannelId is the channel id to subscribe to.
	// Must be sent before / with publish.
	// Cannot change the channel ID after first transmission.
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// PublishRequest contains a publish message request.
	PublishRequest *PublishRequest `protobuf:"bytes,2,opt,name=publish_request,json=publishRequest,proto3" json:"publish_request,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *SubscribeRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *SubscribeRequest) GetPublishRequest() *PublishRequest {
	if x != nil {
		return x.PublishRequest
	}
	return nil
}

// PublishRequest is a message published via the subscribe channel.
type PublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FromPeerId is the peer identifier of the sender.
	// The peer ID will be used to acquire the peer private key and sign the mesasge.
	// The peer should be loaded with Identify in the peer service.
	FromPeerId string `protobuf:"bytes,1,opt,name=from_peer_id,json=fromPeerId,proto3" json:"from_peer_id,omitempty"`
	// Data is the published data.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// Identifier is a uint32 identifier to use for outgoing status.
	// If zero, no outgoing status response will be sent.
	Identifier uint32 `protobuf:"varint,3,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *PublishRequest) GetFromPeerId() string {
	if x != nil {
		return x.FromPeerId
	}
	return ""
}

func (x *PublishRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PublishRequest) GetIdentifier() uint32 {
	if x != nil {
		return x.Identifier
	}
	return 0
}

// SubcribeResponse is a pubsub subscription response message.
type SubscribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// IncomingMessage is an incoming message.
	IncomingMessage *IncomingMessage `protobuf:"bytes,1,opt,name=incoming_message,json=incomingMessage,proto3" json:"incoming_message,omitempty"`
	// OutgoingStatus is status of an outgoing message.
	// Sent when a Publish request finishes.
	OutgoingStatus *OutgoingStatus `protobuf:"bytes,2,opt,name=outgoing_status,json=outgoingStatus,proto3" json:"outgoing_status,omitempty"`
	// SubscriptionStatus is the status of the subscription
	SubscriptionStatus *SubscriptionStatus `protobuf:"bytes,3,opt,name=subscription_status,json=subscriptionStatus,proto3" json:"subscription_status,omitempty"`
}

func (x *SubscribeResponse) Reset() {
	*x = SubscribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeResponse) ProtoMessage() {}

func (x *SubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeResponse.ProtoReflect.Descriptor instead.
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *SubscribeResponse) GetIncomingMessage() *IncomingMessage {
	if x != nil {
		return x.IncomingMessage
	}
	return nil
}

func (x *SubscribeResponse) GetOutgoingStatus() *OutgoingStatus {
	if x != nil {
		return x.OutgoingStatus
	}
	return nil
}

func (x *SubscribeResponse) GetSubscriptionStatus() *SubscriptionStatus {
	if x != nil {
		return x.SubscriptionStatus
	}
	return nil
}

// SubscripionStatus is the status of the subscription handle.
type SubscriptionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Subscribed indicates the subscription is established.
	Subscribed bool `protobuf:"varint,1,opt,name=subscribed,proto3" json:"subscribed,omitempty"`
}

func (x *SubscriptionStatus) Reset() {
	*x = SubscriptionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionStatus) ProtoMessage() {}

func (x *SubscriptionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionStatus.ProtoReflect.Descriptor instead.
func (*SubscriptionStatus) Descriptor() ([]byte, []int) {
	return file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *SubscriptionStatus) GetSubscribed() bool {
	if x != nil {
		return x.Subscribed
	}
	return false
}

// OutgoingStatus is status of an outgoing message.
type OutgoingStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier is the request-provided identifier for the message.
	Identifier uint32 `protobuf:"varint,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// Sent indicates if the message was sent.
	Sent bool `protobuf:"varint,2,opt,name=sent,proto3" json:"sent,omitempty"`
}

func (x *OutgoingStatus) Reset() {
	*x = OutgoingStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutgoingStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutgoingStatus) ProtoMessage() {}

func (x *OutgoingStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutgoingStatus.ProtoReflect.Descriptor instead.
func (*OutgoingStatus) Descriptor() ([]byte, []int) {
	return file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *OutgoingStatus) GetIdentifier() uint32 {
	if x != nil {
		return x.Identifier
	}
	return 0
}

func (x *OutgoingStatus) GetSent() bool {
	if x != nil {
		return x.Sent
	}
	return false
}

// IncomingMessage implements Message with a proto object.
type IncomingMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FromPeerId is the peer identifier of the sender.
	FromPeerId string `protobuf:"bytes,1,opt,name=from_peer_id,json=fromPeerId,proto3" json:"from_peer_id,omitempty"`
	// Authenticated indicates if the message is verified to be from the sender.
	Authenticated bool `protobuf:"varint,2,opt,name=authenticated,proto3" json:"authenticated,omitempty"`
	// Data is the inner data.
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *IncomingMessage) Reset() {
	*x = IncomingMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncomingMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncomingMessage) ProtoMessage() {}

func (x *IncomingMessage) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncomingMessage.ProtoReflect.Descriptor instead.
func (*IncomingMessage) Descriptor() ([]byte, []int) {
	return file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *IncomingMessage) GetFromPeerId() string {
	if x != nil {
		return x.FromPeerId
	}
	return ""
}

func (x *IncomingMessage) GetAuthenticated() bool {
	if x != nil {
		return x.Authenticated
	}
	return false
}

func (x *IncomingMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto protoreflect.FileDescriptor

var file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x65,
	0x72, 0x74, 0x75, 0x72, 0x65, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x62, 0x69,
	0x66, 0x72, 0x6f, 0x73, 0x74, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x75,
	0x62, 0x73, 0x75, 0x62, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x22, 0x77, 0x0a, 0x10, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x0f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x66, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x70, 0x65, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6d,
	0x50, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0xf4, 0x01, 0x0a, 0x11, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x47, 0x0a, 0x10, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x75, 0x62,
	0x73, 0x75, 0x62, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e,
	0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0f, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x69,
	0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x6f, 0x75, 0x74,
	0x67, 0x6f, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x4f, 0x75, 0x74, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x0e, 0x6f, 0x75, 0x74, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x50, 0x0a, 0x13, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70,
	0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x12, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x34, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x22, 0x44, 0x0a, 0x0e, 0x4f, 0x75, 0x74, 0x67, 0x6f,
	0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x65, 0x6e, 0x74, 0x22, 0x6d, 0x0a,
	0x0f, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x20, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x50, 0x65, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x61, 0x0a, 0x0d,
	0x50, 0x75, 0x62, 0x53, 0x75, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a,
	0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x75, 0x62,
	0x73, 0x75, 0x62, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x75, 0x62, 0x73,
	0x75, 0x62, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_rawDescOnce sync.Once
	file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_rawDescData = file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_rawDesc
)

func file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_rawDescGZIP() []byte {
	file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_rawDescOnce.Do(func() {
		file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_rawDescData)
	})
	return file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_rawDescData
}

var file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_goTypes = []interface{}{
	(*SubscribeRequest)(nil),   // 0: pubsub.grpc.SubscribeRequest
	(*PublishRequest)(nil),     // 1: pubsub.grpc.PublishRequest
	(*SubscribeResponse)(nil),  // 2: pubsub.grpc.SubscribeResponse
	(*SubscriptionStatus)(nil), // 3: pubsub.grpc.SubscriptionStatus
	(*OutgoingStatus)(nil),     // 4: pubsub.grpc.OutgoingStatus
	(*IncomingMessage)(nil),    // 5: pubsub.grpc.IncomingMessage
}
var file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_depIdxs = []int32{
	1, // 0: pubsub.grpc.SubscribeRequest.publish_request:type_name -> pubsub.grpc.PublishRequest
	5, // 1: pubsub.grpc.SubscribeResponse.incoming_message:type_name -> pubsub.grpc.IncomingMessage
	4, // 2: pubsub.grpc.SubscribeResponse.outgoing_status:type_name -> pubsub.grpc.OutgoingStatus
	3, // 3: pubsub.grpc.SubscribeResponse.subscription_status:type_name -> pubsub.grpc.SubscriptionStatus
	0, // 4: pubsub.grpc.PubSubService.Subscribe:input_type -> pubsub.grpc.SubscribeRequest
	2, // 5: pubsub.grpc.PubSubService.Subscribe:output_type -> pubsub.grpc.SubscribeResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_init() }
func file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_init() {
	if File_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRequest); i {
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
		file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeResponse); i {
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
		file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionStatus); i {
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
		file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutgoingStatus); i {
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
		file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncomingMessage); i {
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
			RawDescriptor: file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_goTypes,
		DependencyIndexes: file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_depIdxs,
		MessageInfos:      file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_msgTypes,
	}.Build()
	File_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto = out.File
	file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_rawDesc = nil
	file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_goTypes = nil
	file_github_com_aperturerobotics_bifrost_pubsub_grpc_grpc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PubSubServiceClient is the client API for PubSubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PubSubServiceClient interface {
	// Subscribe subscribes to a channel, allowing the subscriber to publish
	// messages over the same channel.
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (PubSubService_SubscribeClient, error)
}

type pubSubServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPubSubServiceClient(cc grpc.ClientConnInterface) PubSubServiceClient {
	return &pubSubServiceClient{cc}
}

func (c *pubSubServiceClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (PubSubService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PubSubService_serviceDesc.Streams[0], "/pubsub.grpc.PubSubService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &pubSubServiceSubscribeClient{stream}
	return x, nil
}

type PubSubService_SubscribeClient interface {
	Send(*SubscribeRequest) error
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type pubSubServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *pubSubServiceSubscribeClient) Send(m *SubscribeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pubSubServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PubSubServiceServer is the server API for PubSubService service.
type PubSubServiceServer interface {
	// Subscribe subscribes to a channel, allowing the subscriber to publish
	// messages over the same channel.
	Subscribe(PubSubService_SubscribeServer) error
}

// UnimplementedPubSubServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPubSubServiceServer struct {
}

func (*UnimplementedPubSubServiceServer) Subscribe(PubSubService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterPubSubServiceServer(s *grpc.Server, srv PubSubServiceServer) {
	s.RegisterService(&_PubSubService_serviceDesc, srv)
}

func _PubSubService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PubSubServiceServer).Subscribe(&pubSubServiceSubscribeServer{stream})
}

type PubSubService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	Recv() (*SubscribeRequest, error)
	grpc.ServerStream
}

type pubSubServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *pubSubServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pubSubServiceSubscribeServer) Recv() (*SubscribeRequest, error) {
	m := new(SubscribeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PubSubService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pubsub.grpc.PubSubService",
	HandlerType: (*PubSubServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _PubSubService_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/aperturerobotics/bifrost/pubsub/grpc/grpc.proto",
}

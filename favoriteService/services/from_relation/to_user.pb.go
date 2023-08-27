// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.0
// source: from_relation/to_user.proto

package proto

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

type GetRelationStatus_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowingId int64 `protobuf:"varint,1,opt,name=following_id,json=followingId,proto3" json:"following_id,omitempty"` //发起关注的人
	FollowerId  int64 `protobuf:"varint,2,opt,name=follower_id,json=followerId,proto3" json:"follower_id,omitempty"`    //被关注的人
}

func (x *GetRelationStatus_Request) Reset() {
	*x = GetRelationStatus_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_from_relation_to_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationStatus_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationStatus_Request) ProtoMessage() {}

func (x *GetRelationStatus_Request) ProtoReflect() protoreflect.Message {
	mi := &file_from_relation_to_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelationStatus_Request.ProtoReflect.Descriptor instead.
func (*GetRelationStatus_Request) Descriptor() ([]byte, []int) {
	return file_from_relation_to_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetRelationStatus_Request) GetFollowingId() int64 {
	if x != nil {
		return x.FollowingId
	}
	return 0
}

func (x *GetRelationStatus_Request) GetFollowerId() int64 {
	if x != nil {
		return x.FollowerId
	}
	return 0
}

type GetRelationStatus_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	IsFollow   bool  `protobuf:"varint,2,opt,name=isFollow,proto3" json:"isFollow,omitempty"`
}

func (x *GetRelationStatus_Response) Reset() {
	*x = GetRelationStatus_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_from_relation_to_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationStatus_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationStatus_Response) ProtoMessage() {}

func (x *GetRelationStatus_Response) ProtoReflect() protoreflect.Message {
	mi := &file_from_relation_to_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelationStatus_Response.ProtoReflect.Descriptor instead.
func (*GetRelationStatus_Response) Descriptor() ([]byte, []int) {
	return file_from_relation_to_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetRelationStatus_Response) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetRelationStatus_Response) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

type GetRelationsStatus_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RelationStatus []*RelationStatus `protobuf:"bytes,1,rep,name=relationStatus,proto3" json:"relationStatus,omitempty"`
	Token          string            `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetRelationsStatus_Request) Reset() {
	*x = GetRelationsStatus_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_from_relation_to_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationsStatus_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationsStatus_Request) ProtoMessage() {}

func (x *GetRelationsStatus_Request) ProtoReflect() protoreflect.Message {
	mi := &file_from_relation_to_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelationsStatus_Request.ProtoReflect.Descriptor instead.
func (*GetRelationsStatus_Request) Descriptor() ([]byte, []int) {
	return file_from_relation_to_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetRelationsStatus_Request) GetRelationStatus() []*RelationStatus {
	if x != nil {
		return x.RelationStatus
	}
	return nil
}

func (x *GetRelationsStatus_Request) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetRelationsStatus_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64             `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	IsFollow   []*RelationStatus `protobuf:"bytes,2,rep,name=isFollow,proto3" json:"isFollow,omitempty"`
}

func (x *GetRelationsStatus_Response) Reset() {
	*x = GetRelationsStatus_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_from_relation_to_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationsStatus_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationsStatus_Response) ProtoMessage() {}

func (x *GetRelationsStatus_Response) ProtoReflect() protoreflect.Message {
	mi := &file_from_relation_to_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelationsStatus_Response.ProtoReflect.Descriptor instead.
func (*GetRelationsStatus_Response) Descriptor() ([]byte, []int) {
	return file_from_relation_to_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetRelationsStatus_Response) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetRelationsStatus_Response) GetIsFollow() []*RelationStatus {
	if x != nil {
		return x.IsFollow
	}
	return nil
}

type RelationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowingId int64 `protobuf:"varint,1,opt,name=following_id,json=followingId,proto3" json:"following_id,omitempty"` //发起关注的人
	FollowerId  int64 `protobuf:"varint,2,opt,name=follower_id,json=followerId,proto3" json:"follower_id,omitempty"`    //被关注的人
	IsFollow    bool  `protobuf:"varint,3,opt,name=isFollow,proto3" json:"isFollow,omitempty"`                          //是否关注
}

func (x *RelationStatus) Reset() {
	*x = RelationStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_from_relation_to_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationStatus) ProtoMessage() {}

func (x *RelationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_from_relation_to_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationStatus.ProtoReflect.Descriptor instead.
func (*RelationStatus) Descriptor() ([]byte, []int) {
	return file_from_relation_to_user_proto_rawDescGZIP(), []int{4}
}

func (x *RelationStatus) GetFollowingId() int64 {
	if x != nil {
		return x.FollowingId
	}
	return 0
}

func (x *RelationStatus) GetFollowerId() int64 {
	if x != nil {
		return x.FollowerId
	}
	return 0
}

func (x *RelationStatus) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

var File_from_relation_to_user_proto protoreflect.FileDescriptor

var file_from_relation_to_user_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x5f, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x64, 0x22, 0x59, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x22, 0x74, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x40, 0x0a, 0x0e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x0e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x74, 0x0a, 0x1b, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x69, 0x73, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x22,
	0x70, 0x0a, 0x0e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69,
	0x6e, 0x67, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x32, 0xd2, 0x01, 0x0a, 0x0d, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_from_relation_to_user_proto_rawDescOnce sync.Once
	file_from_relation_to_user_proto_rawDescData = file_from_relation_to_user_proto_rawDesc
)

func file_from_relation_to_user_proto_rawDescGZIP() []byte {
	file_from_relation_to_user_proto_rawDescOnce.Do(func() {
		file_from_relation_to_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_from_relation_to_user_proto_rawDescData)
	})
	return file_from_relation_to_user_proto_rawDescData
}

var file_from_relation_to_user_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_from_relation_to_user_proto_goTypes = []interface{}{
	(*GetRelationStatus_Request)(nil),   // 0: services.GetRelationStatus_Request
	(*GetRelationStatus_Response)(nil),  // 1: services.GetRelationStatus_Response
	(*GetRelationsStatus_Request)(nil),  // 2: services.GetRelationsStatus_Request
	(*GetRelationsStatus_Response)(nil), // 3: services.GetRelationsStatus_Response
	(*RelationStatus)(nil),              // 4: services.RelationStatus
}
var file_from_relation_to_user_proto_depIdxs = []int32{
	4, // 0: services.GetRelationsStatus_Request.relationStatus:type_name -> services.RelationStatus
	4, // 1: services.GetRelationsStatus_Response.isFollow:type_name -> services.RelationStatus
	0, // 2: services.ToUserService.GetRelationStatus:input_type -> services.GetRelationStatus_Request
	2, // 3: services.ToUserService.GetRelationsStatus:input_type -> services.GetRelationsStatus_Request
	1, // 4: services.ToUserService.GetRelationStatus:output_type -> services.GetRelationStatus_Response
	3, // 5: services.ToUserService.GetRelationsStatus:output_type -> services.GetRelationsStatus_Response
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_from_relation_to_user_proto_init() }
func file_from_relation_to_user_proto_init() {
	if File_from_relation_to_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_from_relation_to_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelationStatus_Request); i {
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
		file_from_relation_to_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelationStatus_Response); i {
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
		file_from_relation_to_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelationsStatus_Request); i {
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
		file_from_relation_to_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelationsStatus_Response); i {
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
		file_from_relation_to_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationStatus); i {
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
			RawDescriptor: file_from_relation_to_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_from_relation_to_user_proto_goTypes,
		DependencyIndexes: file_from_relation_to_user_proto_depIdxs,
		MessageInfos:      file_from_relation_to_user_proto_msgTypes,
	}.Build()
	File_from_relation_to_user_proto = out.File
	file_from_relation_to_user_proto_rawDesc = nil
	file_from_relation_to_user_proto_goTypes = nil
	file_from_relation_to_user_proto_depIdxs = nil
}

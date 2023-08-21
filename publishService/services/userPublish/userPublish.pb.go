// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.0
// source: userPublish.proto

package userPublish

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

type UpdateWorkCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户id
	Count  int32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`                 // 增加的数量
	Type   int32 `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`                   // 1是增加，2是减少
}

func (x *UpdateWorkCountRequest) Reset() {
	*x = UpdateWorkCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userPublish_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWorkCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWorkCountRequest) ProtoMessage() {}

func (x *UpdateWorkCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userPublish_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWorkCountRequest.ProtoReflect.Descriptor instead.
func (*UpdateWorkCountRequest) Descriptor() ([]byte, []int) {
	return file_userPublish_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateWorkCountRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateWorkCountRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *UpdateWorkCountRequest) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type UpdateWorkCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` //响应，成功是0，失败是其他值
}

func (x *UpdateWorkCountResponse) Reset() {
	*x = UpdateWorkCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userPublish_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWorkCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWorkCountResponse) ProtoMessage() {}

func (x *UpdateWorkCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userPublish_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWorkCountResponse.ProtoReflect.Descriptor instead.
func (*UpdateWorkCountResponse) Descriptor() ([]byte, []int) {
	return file_userPublish_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateWorkCountResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

var File_userPublish_proto protoreflect.FileDescriptor

var file_userPublish_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x22, 0x5c, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3b,
	0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x72, 0x0a, 0x10, 0x54,
	0x6f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x5e, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x24, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userPublish_proto_rawDescOnce sync.Once
	file_userPublish_proto_rawDescData = file_userPublish_proto_rawDesc
)

func file_userPublish_proto_rawDescGZIP() []byte {
	file_userPublish_proto_rawDescOnce.Do(func() {
		file_userPublish_proto_rawDescData = protoimpl.X.CompressGZIP(file_userPublish_proto_rawDescData)
	})
	return file_userPublish_proto_rawDescData
}

var file_userPublish_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_userPublish_proto_goTypes = []interface{}{
	(*UpdateWorkCountRequest)(nil),  // 0: userPublish.UpdateWorkCount_request
	(*UpdateWorkCountResponse)(nil), // 1: userPublish.UpdateWorkCount_response
}
var file_userPublish_proto_depIdxs = []int32{
	0, // 0: userPublish.ToPublishService.UpdateWorkCount:input_type -> userPublish.UpdateWorkCount_request
	1, // 1: userPublish.ToPublishService.UpdateWorkCount:output_type -> userPublish.UpdateWorkCount_response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_userPublish_proto_init() }
func file_userPublish_proto_init() {
	if File_userPublish_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userPublish_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateWorkCountRequest); i {
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
		file_userPublish_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateWorkCountResponse); i {
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
			RawDescriptor: file_userPublish_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userPublish_proto_goTypes,
		DependencyIndexes: file_userPublish_proto_depIdxs,
		MessageInfos:      file_userPublish_proto_msgTypes,
	}.Build()
	File_userPublish_proto = out.File
	file_userPublish_proto_rawDesc = nil
	file_userPublish_proto_goTypes = nil
	file_userPublish_proto_depIdxs = nil
}

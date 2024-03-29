// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: activity/v1/lottery.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type DrawRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId      int32 `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`             // app id
	ChannelId  int32 `protobuf:"varint,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"` // channel id
	LootPoolId int32 `protobuf:"varint,3,opt,name=loot_pool_id,json=lootPoolId,proto3" json:"loot_pool_id,omitempty"`
	ItemId     int64 `protobuf:"varint,4,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	UserId     int64 `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户 id
}

func (x *DrawRequest) Reset() {
	*x = DrawRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_v1_lottery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DrawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DrawRequest) ProtoMessage() {}

func (x *DrawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_activity_v1_lottery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DrawRequest.ProtoReflect.Descriptor instead.
func (*DrawRequest) Descriptor() ([]byte, []int) {
	return file_activity_v1_lottery_proto_rawDescGZIP(), []int{0}
}

func (x *DrawRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *DrawRequest) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *DrawRequest) GetLootPoolId() int32 {
	if x != nil {
		return x.LootPoolId
	}
	return 0
}

func (x *DrawRequest) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *DrawRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DrawReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*DrawItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`    // 列表
	Total int32       `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"` // 数量
}

func (x *DrawReply) Reset() {
	*x = DrawReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_v1_lottery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DrawReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DrawReply) ProtoMessage() {}

func (x *DrawReply) ProtoReflect() protoreflect.Message {
	mi := &file_activity_v1_lottery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DrawReply.ProtoReflect.Descriptor instead.
func (*DrawReply) Descriptor() ([]byte, []int) {
	return file_activity_v1_lottery_proto_rawDescGZIP(), []int{1}
}

func (x *DrawReply) GetList() []*DrawItem {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *DrawReply) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DrawItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId   int32 `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"` // 道具 id
	Quantity int32 `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`           // 购买数量
}

func (x *DrawItem) Reset() {
	*x = DrawItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_v1_lottery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DrawItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DrawItem) ProtoMessage() {}

func (x *DrawItem) ProtoReflect() protoreflect.Message {
	mi := &file_activity_v1_lottery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DrawItem.ProtoReflect.Descriptor instead.
func (*DrawItem) Descriptor() ([]byte, []int) {
	return file_activity_v1_lottery_proto_rawDescGZIP(), []int{2}
}

func (x *DrawItem) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *DrawItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_activity_v1_lottery_proto protoreflect.FileDescriptor

var file_activity_v1_lottery_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f,
	0x74, 0x74, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x0b, 0x44,
	0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0c, 0x6c, 0x6f, 0x6f, 0x74, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x6f, 0x6f, 0x74, 0x50, 0x6f, 0x6f, 0x6c,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x50, 0x0a, 0x09, 0x44, 0x72, 0x61, 0x77, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x2d, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x72, 0x61, 0x77, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x3f, 0x0a, 0x08, 0x44, 0x72, 0x61, 0x77, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x32, 0x7a, 0x0a, 0x07, 0x4c, 0x6f, 0x74, 0x74, 0x65,
	0x72, 0x79, 0x12, 0x6f, 0x0a, 0x04, 0x44, 0x72, 0x61, 0x77, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x72, 0x61,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x72, 0x61, 0x77, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x01, 0x2a, 0x22,
	0x22, 0x2f, 0x73, 0x74, 0x2d, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x6c, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x79, 0x2f, 0x64,
	0x72, 0x61, 0x77, 0x42, 0x41, 0x0a, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x75, 0x6c, 0x75, 0x2d, 0x6c, 0x73, 0x2f, 0x73, 0x74, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_activity_v1_lottery_proto_rawDescOnce sync.Once
	file_activity_v1_lottery_proto_rawDescData = file_activity_v1_lottery_proto_rawDesc
)

func file_activity_v1_lottery_proto_rawDescGZIP() []byte {
	file_activity_v1_lottery_proto_rawDescOnce.Do(func() {
		file_activity_v1_lottery_proto_rawDescData = protoimpl.X.CompressGZIP(file_activity_v1_lottery_proto_rawDescData)
	})
	return file_activity_v1_lottery_proto_rawDescData
}

var file_activity_v1_lottery_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_activity_v1_lottery_proto_goTypes = []interface{}{
	(*DrawRequest)(nil), // 0: api.activity.v1.DrawRequest
	(*DrawReply)(nil),   // 1: api.activity.v1.DrawReply
	(*DrawItem)(nil),    // 2: api.activity.v1.DrawItem
}
var file_activity_v1_lottery_proto_depIdxs = []int32{
	2, // 0: api.activity.v1.DrawReply.list:type_name -> api.activity.v1.DrawItem
	0, // 1: api.activity.v1.Lottery.Draw:input_type -> api.activity.v1.DrawRequest
	1, // 2: api.activity.v1.Lottery.Draw:output_type -> api.activity.v1.DrawReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_activity_v1_lottery_proto_init() }
func file_activity_v1_lottery_proto_init() {
	if File_activity_v1_lottery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_activity_v1_lottery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DrawRequest); i {
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
		file_activity_v1_lottery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DrawReply); i {
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
		file_activity_v1_lottery_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DrawItem); i {
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
			RawDescriptor: file_activity_v1_lottery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_activity_v1_lottery_proto_goTypes,
		DependencyIndexes: file_activity_v1_lottery_proto_depIdxs,
		MessageInfos:      file_activity_v1_lottery_proto_msgTypes,
	}.Build()
	File_activity_v1_lottery_proto = out.File
	file_activity_v1_lottery_proto_rawDesc = nil
	file_activity_v1_lottery_proto_goTypes = nil
	file_activity_v1_lottery_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: free/v1/free.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type FreeTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId int32 `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *FreeTypeRequest) Reset() {
	*x = FreeTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_free_v1_free_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FreeTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FreeTypeRequest) ProtoMessage() {}

func (x *FreeTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_free_v1_free_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FreeTypeRequest.ProtoReflect.Descriptor instead.
func (*FreeTypeRequest) Descriptor() ([]byte, []int) {
	return file_free_v1_free_proto_rawDescGZIP(), []int{0}
}

func (x *FreeTypeRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

type FreeTypeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*FreeTypeItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Total int32           `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *FreeTypeReply) Reset() {
	*x = FreeTypeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_free_v1_free_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FreeTypeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FreeTypeReply) ProtoMessage() {}

func (x *FreeTypeReply) ProtoReflect() protoreflect.Message {
	mi := &file_free_v1_free_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FreeTypeReply.ProtoReflect.Descriptor instead.
func (*FreeTypeReply) Descriptor() ([]byte, []int) {
	return file_free_v1_free_proto_rawDescGZIP(), []int{1}
}

func (x *FreeTypeReply) GetList() []*FreeTypeItem {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *FreeTypeReply) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type FreeTypeItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameSeriesId int32  `protobuf:"varint,1,opt,name=game_series_id,json=gameSeriesId,proto3" json:"game_series_id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FreeTypeItem) Reset() {
	*x = FreeTypeItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_free_v1_free_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FreeTypeItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FreeTypeItem) ProtoMessage() {}

func (x *FreeTypeItem) ProtoReflect() protoreflect.Message {
	mi := &file_free_v1_free_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FreeTypeItem.ProtoReflect.Descriptor instead.
func (*FreeTypeItem) Descriptor() ([]byte, []int) {
	return file_free_v1_free_proto_rawDescGZIP(), []int{2}
}

func (x *FreeTypeItem) GetGameSeriesId() int32 {
	if x != nil {
		return x.GameSeriesId
	}
	return 0
}

func (x *FreeTypeItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FreeListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId        int32 `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`                        // app id  [(validate.rules).int32 = {gt: 0}]
	GameSeriesId int32 `protobuf:"varint,2,opt,name=game_series_id,json=gameSeriesId,proto3" json:"game_series_id,omitempty"` // 侧边栏 id
}

func (x *FreeListRequest) Reset() {
	*x = FreeListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_free_v1_free_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FreeListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FreeListRequest) ProtoMessage() {}

func (x *FreeListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_free_v1_free_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FreeListRequest.ProtoReflect.Descriptor instead.
func (*FreeListRequest) Descriptor() ([]byte, []int) {
	return file_free_v1_free_proto_rawDescGZIP(), []int{3}
}

func (x *FreeListRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *FreeListRequest) GetGameSeriesId() int32 {
	if x != nil {
		return x.GameSeriesId
	}
	return 0
}

type FreeListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*FreeListItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Total int32           `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"` // 条数
}

func (x *FreeListReply) Reset() {
	*x = FreeListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_free_v1_free_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FreeListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FreeListReply) ProtoMessage() {}

func (x *FreeListReply) ProtoReflect() protoreflect.Message {
	mi := &file_free_v1_free_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FreeListReply.ProtoReflect.Descriptor instead.
func (*FreeListReply) Descriptor() ([]byte, []int) {
	return file_free_v1_free_proto_rawDescGZIP(), []int{4}
}

func (x *FreeListReply) GetList() []*FreeListItem {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *FreeListReply) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type FreeListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameFreeConfigId int32  `protobuf:"varint,1,opt,name=game_free_config_id,json=gameFreeConfigId,proto3" json:"game_free_config_id,omitempty"` // 自由场配置 id
	GameTypeId       int32  `protobuf:"varint,2,opt,name=game_type_id,json=gameTypeId,proto3" json:"game_type_id,omitempty"`                     // 赛事配置 id
	GameTypeName     string `protobuf:"bytes,3,opt,name=game_type_name,json=gameTypeName,proto3" json:"game_type_name,omitempty"`                // 赛事类型名称
	// int32 game_category_id = 4; // 赛事类型 id
	// string game_category_name = 5; // 赛事分类名称
	Name        string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`                                    // 配置名称
	Title       string `protobuf:"bytes,7,opt,name=title,proto3" json:"title,omitempty"`                                  // 配置标题
	Image       string `protobuf:"bytes,8,opt,name=image,proto3" json:"image,omitempty"`                                  // 图片
	Desc        string `protobuf:"bytes,9,opt,name=desc,proto3" json:"desc,omitempty"`                                    // 描述
	ItemId      int32  `protobuf:"varint,10,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`                // 道具 id
	ItemName    string `protobuf:"bytes,11,opt,name=item_name,json=itemName,proto3" json:"item_name,omitempty"`           // 道具名称
	ItemImage   string `protobuf:"bytes,12,opt,name=item_image,json=itemImage,proto3" json:"item_image,omitempty"`        // 道具名称
	Quantity    int32  `protobuf:"varint,13,opt,name=quantity,proto3" json:"quantity,omitempty"`                          // 报名道具数量
	MinQuantity int32  `protobuf:"varint,15,opt,name=min_quantity,json=minQuantity,proto3" json:"min_quantity,omitempty"` // 本人拥有最小道具数量
	MaxQuantity int32  `protobuf:"varint,16,opt,name=max_quantity,json=maxQuantity,proto3" json:"max_quantity,omitempty"` // 本人拥有最大道具数量
	MaxMultiple int32  `protobuf:"varint,17,opt,name=max_multiple,json=maxMultiple,proto3" json:"max_multiple,omitempty"` // 封顶倍数
	Robot       int32  `protobuf:"varint,18,opt,name=robot,proto3" json:"robot,omitempty"`                                // 是否是机器人
	BaseScore   int32  `protobuf:"varint,19,opt,name=base_score,json=baseScore,proto3" json:"base_score,omitempty"`       // 底分
	PeopleNo    int32  `protobuf:"varint,20,opt,name=people_no,json=peopleNo,proto3" json:"people_no,omitempty"`          // 当前人数
}

func (x *FreeListItem) Reset() {
	*x = FreeListItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_free_v1_free_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FreeListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FreeListItem) ProtoMessage() {}

func (x *FreeListItem) ProtoReflect() protoreflect.Message {
	mi := &file_free_v1_free_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FreeListItem.ProtoReflect.Descriptor instead.
func (*FreeListItem) Descriptor() ([]byte, []int) {
	return file_free_v1_free_proto_rawDescGZIP(), []int{5}
}

func (x *FreeListItem) GetGameFreeConfigId() int32 {
	if x != nil {
		return x.GameFreeConfigId
	}
	return 0
}

func (x *FreeListItem) GetGameTypeId() int32 {
	if x != nil {
		return x.GameTypeId
	}
	return 0
}

func (x *FreeListItem) GetGameTypeName() string {
	if x != nil {
		return x.GameTypeName
	}
	return ""
}

func (x *FreeListItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FreeListItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *FreeListItem) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *FreeListItem) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *FreeListItem) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *FreeListItem) GetItemName() string {
	if x != nil {
		return x.ItemName
	}
	return ""
}

func (x *FreeListItem) GetItemImage() string {
	if x != nil {
		return x.ItemImage
	}
	return ""
}

func (x *FreeListItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *FreeListItem) GetMinQuantity() int32 {
	if x != nil {
		return x.MinQuantity
	}
	return 0
}

func (x *FreeListItem) GetMaxQuantity() int32 {
	if x != nil {
		return x.MaxQuantity
	}
	return 0
}

func (x *FreeListItem) GetMaxMultiple() int32 {
	if x != nil {
		return x.MaxMultiple
	}
	return 0
}

func (x *FreeListItem) GetRobot() int32 {
	if x != nil {
		return x.Robot
	}
	return 0
}

func (x *FreeListItem) GetBaseScore() int32 {
	if x != nil {
		return x.BaseScore
	}
	return 0
}

func (x *FreeListItem) GetPeopleNo() int32 {
	if x != nil {
		return x.PeopleNo
	}
	return 0
}

type FreeSignupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId           int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                                   // 用户 id
	GameFreeConfigId int32 `protobuf:"varint,2,opt,name=game_free_config_id,json=gameFreeConfigId,proto3" json:"game_free_config_id,omitempty"` // 自由场 id
}

func (x *FreeSignupRequest) Reset() {
	*x = FreeSignupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_free_v1_free_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FreeSignupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FreeSignupRequest) ProtoMessage() {}

func (x *FreeSignupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_free_v1_free_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FreeSignupRequest.ProtoReflect.Descriptor instead.
func (*FreeSignupRequest) Descriptor() ([]byte, []int) {
	return file_free_v1_free_proto_rawDescGZIP(), []int{6}
}

func (x *FreeSignupRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FreeSignupRequest) GetGameFreeConfigId() int32 {
	if x != nil {
		return x.GameFreeConfigId
	}
	return 0
}

type FreeSignupReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FreeSignupReply) Reset() {
	*x = FreeSignupReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_free_v1_free_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FreeSignupReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FreeSignupReply) ProtoMessage() {}

func (x *FreeSignupReply) ProtoReflect() protoreflect.Message {
	mi := &file_free_v1_free_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FreeSignupReply.ProtoReflect.Descriptor instead.
func (*FreeSignupReply) Descriptor() ([]byte, []int) {
	return file_free_v1_free_proto_rawDescGZIP(), []int{7}
}

var File_free_v1_free_proto protoreflect.FileDescriptor

var file_free_v1_free_proto_rawDesc = []byte{
	0x0a, 0x12, 0x66, 0x72, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x72, 0x65, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0f, 0x46, 0x72, 0x65, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x61,
	0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x70, 0x70,
	0x49, 0x64, 0x22, 0x54, 0x0a, 0x0d, 0x46, 0x72, 0x65, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x2d, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x72, 0x65, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x48, 0x0a, 0x0c, 0x46, 0x72, 0x65, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x24, 0x0a, 0x0e, 0x67, 0x61, 0x6d, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x57, 0x0a, 0x0f, 0x46, 0x72, 0x65, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x0e,
	0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x0c, 0x67,
	0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x0d, 0x46,
	0x72, 0x65, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2d, 0x0a, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x72, 0x65, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x22, 0x85, 0x04, 0x0a, 0x0c, 0x46, 0x72, 0x65, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x2d, 0x0a, 0x13, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x66, 0x72, 0x65, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x10, 0x67, 0x61, 0x6d, 0x65, 0x46, 0x72, 0x65, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0c, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x61, 0x6d,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73,
	0x63, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x17, 0x0a,
	0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x21,
	0x0a, 0x0c, 0x6d, 0x69, 0x6e, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x6c, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x62, 0x6f, 0x74,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x4e, 0x6f, 0x22, 0x64, 0x0a, 0x11, 0x46, 0x72, 0x65,
	0x65, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x13, 0x67, 0x61, 0x6d, 0x65, 0x5f,
	0x66, 0x72, 0x65, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x10, 0x67,
	0x61, 0x6d, 0x65, 0x46, 0x72, 0x65, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x22,
	0x11, 0x0a, 0x0f, 0x46, 0x72, 0x65, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x32, 0xcd, 0x02, 0x0a, 0x04, 0x46, 0x72, 0x65, 0x65, 0x12, 0x6b, 0x0a, 0x08, 0x46,
	0x72, 0x65, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x61,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x72, 0x65, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x61, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x72, 0x65, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x73,
	0x74, 0x2d, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x72, 0x65, 0x65, 0x2f,
	0x73, 0x75, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x67, 0x0a, 0x08, 0x46, 0x72, 0x65, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x61, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x72, 0x65, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x72, 0x65, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x73, 0x74, 0x2d, 0x67,
	0x61, 0x6d, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x72, 0x65, 0x65, 0x2f, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x6f, 0x0a, 0x0a, 0x46, 0x72, 0x65, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12,
	0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x72,
	0x65, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x72,
	0x65, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x23, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x73, 0x74, 0x2d, 0x67, 0x61,
	0x6d, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x72, 0x65, 0x65, 0x2f, 0x73, 0x69, 0x67, 0x6e,
	0x75, 0x70, 0x42, 0x28, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x17, 0x73, 0x74, 0x2d, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_free_v1_free_proto_rawDescOnce sync.Once
	file_free_v1_free_proto_rawDescData = file_free_v1_free_proto_rawDesc
)

func file_free_v1_free_proto_rawDescGZIP() []byte {
	file_free_v1_free_proto_rawDescOnce.Do(func() {
		file_free_v1_free_proto_rawDescData = protoimpl.X.CompressGZIP(file_free_v1_free_proto_rawDescData)
	})
	return file_free_v1_free_proto_rawDescData
}

var file_free_v1_free_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_free_v1_free_proto_goTypes = []interface{}{
	(*FreeTypeRequest)(nil),   // 0: api.race.v1.FreeTypeRequest
	(*FreeTypeReply)(nil),     // 1: api.race.v1.FreeTypeReply
	(*FreeTypeItem)(nil),      // 2: api.race.v1.FreeTypeItem
	(*FreeListRequest)(nil),   // 3: api.race.v1.FreeListRequest
	(*FreeListReply)(nil),     // 4: api.race.v1.FreeListReply
	(*FreeListItem)(nil),      // 5: api.race.v1.FreeListItem
	(*FreeSignupRequest)(nil), // 6: api.race.v1.FreeSignupRequest
	(*FreeSignupReply)(nil),   // 7: api.race.v1.FreeSignupReply
}
var file_free_v1_free_proto_depIdxs = []int32{
	2, // 0: api.race.v1.FreeTypeReply.list:type_name -> api.race.v1.FreeTypeItem
	5, // 1: api.race.v1.FreeListReply.list:type_name -> api.race.v1.FreeListItem
	0, // 2: api.race.v1.Free.FreeType:input_type -> api.race.v1.FreeTypeRequest
	3, // 3: api.race.v1.Free.FreeList:input_type -> api.race.v1.FreeListRequest
	6, // 4: api.race.v1.Free.FreeSignup:input_type -> api.race.v1.FreeSignupRequest
	1, // 5: api.race.v1.Free.FreeType:output_type -> api.race.v1.FreeTypeReply
	4, // 6: api.race.v1.Free.FreeList:output_type -> api.race.v1.FreeListReply
	7, // 7: api.race.v1.Free.FreeSignup:output_type -> api.race.v1.FreeSignupReply
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_free_v1_free_proto_init() }
func file_free_v1_free_proto_init() {
	if File_free_v1_free_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_free_v1_free_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FreeTypeRequest); i {
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
		file_free_v1_free_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FreeTypeReply); i {
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
		file_free_v1_free_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FreeTypeItem); i {
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
		file_free_v1_free_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FreeListRequest); i {
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
		file_free_v1_free_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FreeListReply); i {
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
		file_free_v1_free_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FreeListItem); i {
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
		file_free_v1_free_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FreeSignupRequest); i {
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
		file_free_v1_free_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FreeSignupReply); i {
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
			RawDescriptor: file_free_v1_free_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_free_v1_free_proto_goTypes,
		DependencyIndexes: file_free_v1_free_proto_depIdxs,
		MessageInfos:      file_free_v1_free_proto_msgTypes,
	}.Build()
	File_free_v1_free_proto = out.File
	file_free_v1_free_proto_rawDesc = nil
	file_free_v1_free_proto_goTypes = nil
	file_free_v1_free_proto_depIdxs = nil
}
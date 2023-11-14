// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: bag/v1/bag.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	v1 "github.com/lulu-ls/st-api/api/common/v1"
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

type ItemListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   *v1.Paginate `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	AppId  int32        `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`    //  [(validate.rules).int32 = {gt: 0}]
	UserId int64        `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户 id
}

func (x *ItemListRequest) Reset() {
	*x = ItemListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bag_v1_bag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemListRequest) ProtoMessage() {}

func (x *ItemListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bag_v1_bag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemListRequest.ProtoReflect.Descriptor instead.
func (*ItemListRequest) Descriptor() ([]byte, []int) {
	return file_bag_v1_bag_proto_rawDescGZIP(), []int{0}
}

func (x *ItemListRequest) GetPage() *v1.Paginate {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ItemListRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *ItemListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ItemListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*ListItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Total int32       `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ItemListReply) Reset() {
	*x = ItemListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bag_v1_bag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemListReply) ProtoMessage() {}

func (x *ItemListReply) ProtoReflect() protoreflect.Message {
	mi := &file_bag_v1_bag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemListReply.ProtoReflect.Descriptor instead.
func (*ItemListReply) Descriptor() ([]byte, []int) {
	return file_bag_v1_bag_proto_rawDescGZIP(), []int{1}
}

func (x *ItemListReply) GetList() []*ListItem {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *ItemListReply) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type ListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId             int32  `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`                                         // 道具 id
	ItemCategoryId     int32  `protobuf:"varint,2,opt,name=item_category_id,json=itemCategoryId,proto3" json:"item_category_id,omitempty"`               // 分类
	ItemFuncId         int32  `protobuf:"varint,3,opt,name=item_func_id,json=itemFuncId,proto3" json:"item_func_id,omitempty"`                           // 1 虚拟兑换 2 实物兑换
	Name               string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`                                                            // 名称
	Title              string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`                                                          // 标题
	Image              string `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`                                                          // 道具图片
	MaxStack           int32  `protobuf:"varint,7,opt,name=max_stack,json=maxStack,proto3" json:"max_stack,omitempty"`                                   // 最大叠加
	Total              int32  `protobuf:"varint,8,opt,name=total,proto3" json:"total,omitempty"`                                                         // 道具数量
	ItemExchangeWareId int32  `protobuf:"varint,9,opt,name=item_exchange_ware_id,json=itemExchangeWareId,proto3" json:"item_exchange_ware_id,omitempty"` // 兑换配置 id
	ItemQuantity       int32  `protobuf:"varint,10,opt,name=item_quantity,json=itemQuantity,proto3" json:"item_quantity,omitempty"`                      // 兑换所需数量
	WareId             int32  `protobuf:"varint,11,opt,name=ware_id,json=wareId,proto3" json:"ware_id,omitempty"`                                        // 商品 id
	WareCategoryId     int32  `protobuf:"varint,12,opt,name=ware_category_id,json=wareCategoryId,proto3" json:"ware_category_id,omitempty"`              // 兑换商品类型
	WareName           string `protobuf:"bytes,13,opt,name=ware_name,json=wareName,proto3" json:"ware_name,omitempty"`                                   // 商品名称
	WareImage          string `protobuf:"bytes,14,opt,name=ware_image,json=wareImage,proto3" json:"ware_image,omitempty"`                                // 商品图片
	WareExternalUrl    string `protobuf:"bytes,15,opt,name=ware_external_url,json=wareExternalUrl,proto3" json:"ware_external_url,omitempty"`            // 封面图
	Information        string `protobuf:"bytes,16,opt,name=information,proto3" json:"information,omitempty"`                                             // 介绍
	IsPhoneBill        bool   `protobuf:"varint,17,opt,name=is_phone_bill,json=isPhoneBill,proto3" json:"is_phone_bill,omitempty"`                       // 是否是话费
}

func (x *ListItem) Reset() {
	*x = ListItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bag_v1_bag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItem) ProtoMessage() {}

func (x *ListItem) ProtoReflect() protoreflect.Message {
	mi := &file_bag_v1_bag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItem.ProtoReflect.Descriptor instead.
func (*ListItem) Descriptor() ([]byte, []int) {
	return file_bag_v1_bag_proto_rawDescGZIP(), []int{2}
}

func (x *ListItem) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *ListItem) GetItemCategoryId() int32 {
	if x != nil {
		return x.ItemCategoryId
	}
	return 0
}

func (x *ListItem) GetItemFuncId() int32 {
	if x != nil {
		return x.ItemFuncId
	}
	return 0
}

func (x *ListItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ListItem) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *ListItem) GetMaxStack() int32 {
	if x != nil {
		return x.MaxStack
	}
	return 0
}

func (x *ListItem) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListItem) GetItemExchangeWareId() int32 {
	if x != nil {
		return x.ItemExchangeWareId
	}
	return 0
}

func (x *ListItem) GetItemQuantity() int32 {
	if x != nil {
		return x.ItemQuantity
	}
	return 0
}

func (x *ListItem) GetWareId() int32 {
	if x != nil {
		return x.WareId
	}
	return 0
}

func (x *ListItem) GetWareCategoryId() int32 {
	if x != nil {
		return x.WareCategoryId
	}
	return 0
}

func (x *ListItem) GetWareName() string {
	if x != nil {
		return x.WareName
	}
	return ""
}

func (x *ListItem) GetWareImage() string {
	if x != nil {
		return x.WareImage
	}
	return ""
}

func (x *ListItem) GetWareExternalUrl() string {
	if x != nil {
		return x.WareExternalUrl
	}
	return ""
}

func (x *ListItem) GetInformation() string {
	if x != nil {
		return x.Information
	}
	return ""
}

func (x *ListItem) GetIsPhoneBill() bool {
	if x != nil {
		return x.IsPhoneBill
	}
	return false
}

type ListEmojiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId     int32 `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`             // app id
	ChannelId int32 `protobuf:"varint,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"` // channel id
	UserId    int64 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`          // 用户 id
}

func (x *ListEmojiRequest) Reset() {
	*x = ListEmojiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bag_v1_bag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEmojiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEmojiRequest) ProtoMessage() {}

func (x *ListEmojiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bag_v1_bag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEmojiRequest.ProtoReflect.Descriptor instead.
func (*ListEmojiRequest) Descriptor() ([]byte, []int) {
	return file_bag_v1_bag_proto_rawDescGZIP(), []int{3}
}

func (x *ListEmojiRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *ListEmojiRequest) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *ListEmojiRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ListEmojiReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*ListEmojiItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`    // 列表
	Total int32            `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"` // 合计
}

func (x *ListEmojiReply) Reset() {
	*x = ListEmojiReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bag_v1_bag_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEmojiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEmojiReply) ProtoMessage() {}

func (x *ListEmojiReply) ProtoReflect() protoreflect.Message {
	mi := &file_bag_v1_bag_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEmojiReply.ProtoReflect.Descriptor instead.
func (*ListEmojiReply) Descriptor() ([]byte, []int) {
	return file_bag_v1_bag_proto_rawDescGZIP(), []int{4}
}

func (x *ListEmojiReply) GetList() []*ListEmojiItem {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *ListEmojiReply) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type ListEmojiItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId               int32  `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`                                             // 道具 id
	ItemName             string `protobuf:"bytes,2,opt,name=item_name,json=itemName,proto3" json:"item_name,omitempty"`                                        // 道具名称
	ItemImage            string `protobuf:"bytes,3,opt,name=item_image,json=itemImage,proto3" json:"item_image,omitempty"`                                     // 道具图片
	Total                int32  `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`                                                             // 自己拥有的道具数量
	ExchangeItemId       int32  `protobuf:"varint,5,opt,name=exchange_item_id,json=exchangeItemId,proto3" json:"exchange_item_id,omitempty"`                   // 购买所需道具 id
	ExchangeItemName     string `protobuf:"bytes,6,opt,name=exchange_item_name,json=exchangeItemName,proto3" json:"exchange_item_name,omitempty"`              // 购买所需道具名称
	ExchangeItemImage    string `protobuf:"bytes,7,opt,name=exchange_item_image,json=exchangeItemImage,proto3" json:"exchange_item_image,omitempty"`           // 购买所需道具图片
	ExchangeItemQuantity int32  `protobuf:"varint,8,opt,name=exchange_item_quantity,json=exchangeItemQuantity,proto3" json:"exchange_item_quantity,omitempty"` // 购买所需道具数量
}

func (x *ListEmojiItem) Reset() {
	*x = ListEmojiItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bag_v1_bag_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEmojiItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEmojiItem) ProtoMessage() {}

func (x *ListEmojiItem) ProtoReflect() protoreflect.Message {
	mi := &file_bag_v1_bag_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEmojiItem.ProtoReflect.Descriptor instead.
func (*ListEmojiItem) Descriptor() ([]byte, []int) {
	return file_bag_v1_bag_proto_rawDescGZIP(), []int{5}
}

func (x *ListEmojiItem) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *ListEmojiItem) GetItemName() string {
	if x != nil {
		return x.ItemName
	}
	return ""
}

func (x *ListEmojiItem) GetItemImage() string {
	if x != nil {
		return x.ItemImage
	}
	return ""
}

func (x *ListEmojiItem) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListEmojiItem) GetExchangeItemId() int32 {
	if x != nil {
		return x.ExchangeItemId
	}
	return 0
}

func (x *ListEmojiItem) GetExchangeItemName() string {
	if x != nil {
		return x.ExchangeItemName
	}
	return ""
}

func (x *ListEmojiItem) GetExchangeItemImage() string {
	if x != nil {
		return x.ExchangeItemImage
	}
	return ""
}

func (x *ListEmojiItem) GetExchangeItemQuantity() int32 {
	if x != nil {
		return x.ExchangeItemQuantity
	}
	return 0
}

type UseItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId     int32 `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`             // app id
	ChannelId int32 `protobuf:"varint,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"` // channel id
	UserId    int64 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`          // 用户 id
	ItemId    int32 `protobuf:"varint,4,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`          // 道具 id
	Quantity  int32 `protobuf:"varint,6,opt,name=quantity,proto3" json:"quantity,omitempty"`                    // 使用数量
}

func (x *UseItemRequest) Reset() {
	*x = UseItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bag_v1_bag_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UseItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UseItemRequest) ProtoMessage() {}

func (x *UseItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bag_v1_bag_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UseItemRequest.ProtoReflect.Descriptor instead.
func (*UseItemRequest) Descriptor() ([]byte, []int) {
	return file_bag_v1_bag_proto_rawDescGZIP(), []int{6}
}

func (x *UseItemRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *UseItemRequest) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *UseItemRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UseItemRequest) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *UseItemRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type UseItemReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UseItemReply) Reset() {
	*x = UseItemReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bag_v1_bag_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UseItemReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UseItemReply) ProtoMessage() {}

func (x *UseItemReply) ProtoReflect() protoreflect.Message {
	mi := &file_bag_v1_bag_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UseItemReply.ProtoReflect.Descriptor instead.
func (*UseItemReply) Descriptor() ([]byte, []int) {
	return file_bag_v1_bag_proto_rawDescGZIP(), []int{7}
}

var File_bag_v1_bag_proto protoreflect.FileDescriptor

var file_bag_v1_bag_proto_rawDesc = []byte{
	0x0a, 0x10, 0x62, 0x61, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a,
	0x0f, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2b, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x15, 0x0a,
	0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61,
	0x70, 0x70, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4f, 0x0a,
	0x0d, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x28,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x62, 0x61, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xab,
	0x04, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x74,
	0x65, 0x6d, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e,
	0x69, 0x74, 0x65, 0x6d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0c, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x46, 0x75, 0x6e, 0x63, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x31, 0x0a, 0x15, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x5f, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x12, 0x69, 0x74, 0x65, 0x6d, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x57, 0x61, 0x72, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x69,
	0x74, 0x65, 0x6d, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x77,
	0x61, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x61,
	0x72, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e,
	0x77, 0x61, 0x72, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x77, 0x61, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x77,
	0x61, 0x72, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x77, 0x61, 0x72, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x77, 0x61,
	0x72, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x77, 0x61, 0x72, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x62, 0x69, 0x6c, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x69, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x22, 0x61, 0x0a, 0x10,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x55, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x2d, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xb8, 0x02, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x6d, 0x6f, 0x6a, 0x69, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x2c, 0x0a,
	0x12, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x22, 0xa6, 0x01, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x06, 0x69,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00,
	0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x0e, 0x0a, 0x0c, 0x55, 0x73,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xc6, 0x02, 0x0a, 0x03, 0x42,
	0x61, 0x67, 0x12, 0x69, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x62, 0x61, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01,
	0x2a, 0x22, 0x1a, 0x2f, 0x73, 0x74, 0x2d, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x61, 0x67, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x6d, 0x0a,
	0x09, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x62, 0x61, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6d, 0x6f, 0x6a,
	0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62,
	0x61, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22,
	0x1b, 0x2f, 0x73, 0x74, 0x2d, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61,
	0x67, 0x2f, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x65, 0x0a, 0x07,
	0x55, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x73, 0x74, 0x2d, 0x67, 0x61,
	0x6d, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x67, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2f,
	0x75, 0x73, 0x65, 0x42, 0x37, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x67, 0x2e, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x75, 0x6c, 0x75, 0x2d, 0x6c, 0x73, 0x2f, 0x73, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x62, 0x61, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bag_v1_bag_proto_rawDescOnce sync.Once
	file_bag_v1_bag_proto_rawDescData = file_bag_v1_bag_proto_rawDesc
)

func file_bag_v1_bag_proto_rawDescGZIP() []byte {
	file_bag_v1_bag_proto_rawDescOnce.Do(func() {
		file_bag_v1_bag_proto_rawDescData = protoimpl.X.CompressGZIP(file_bag_v1_bag_proto_rawDescData)
	})
	return file_bag_v1_bag_proto_rawDescData
}

var file_bag_v1_bag_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_bag_v1_bag_proto_goTypes = []interface{}{
	(*ItemListRequest)(nil),  // 0: api.bag.v1.ItemListRequest
	(*ItemListReply)(nil),    // 1: api.bag.v1.ItemListReply
	(*ListItem)(nil),         // 2: api.bag.v1.ListItem
	(*ListEmojiRequest)(nil), // 3: api.bag.v1.ListEmojiRequest
	(*ListEmojiReply)(nil),   // 4: api.bag.v1.ListEmojiReply
	(*ListEmojiItem)(nil),    // 5: api.bag.v1.ListEmojiItem
	(*UseItemRequest)(nil),   // 6: api.bag.v1.UseItemRequest
	(*UseItemReply)(nil),     // 7: api.bag.v1.UseItemReply
	(*v1.Paginate)(nil),      // 8: api.common.v1.Paginate
}
var file_bag_v1_bag_proto_depIdxs = []int32{
	8, // 0: api.bag.v1.ItemListRequest.page:type_name -> api.common.v1.Paginate
	2, // 1: api.bag.v1.ItemListReply.list:type_name -> api.bag.v1.ListItem
	5, // 2: api.bag.v1.ListEmojiReply.list:type_name -> api.bag.v1.ListEmojiItem
	0, // 3: api.bag.v1.Bag.ListItem:input_type -> api.bag.v1.ItemListRequest
	3, // 4: api.bag.v1.Bag.ListEmoji:input_type -> api.bag.v1.ListEmojiRequest
	6, // 5: api.bag.v1.Bag.UseItem:input_type -> api.bag.v1.UseItemRequest
	1, // 6: api.bag.v1.Bag.ListItem:output_type -> api.bag.v1.ItemListReply
	4, // 7: api.bag.v1.Bag.ListEmoji:output_type -> api.bag.v1.ListEmojiReply
	7, // 8: api.bag.v1.Bag.UseItem:output_type -> api.bag.v1.UseItemReply
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_bag_v1_bag_proto_init() }
func file_bag_v1_bag_proto_init() {
	if File_bag_v1_bag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bag_v1_bag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemListRequest); i {
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
		file_bag_v1_bag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemListReply); i {
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
		file_bag_v1_bag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListItem); i {
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
		file_bag_v1_bag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEmojiRequest); i {
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
		file_bag_v1_bag_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEmojiReply); i {
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
		file_bag_v1_bag_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEmojiItem); i {
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
		file_bag_v1_bag_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UseItemRequest); i {
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
		file_bag_v1_bag_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UseItemReply); i {
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
			RawDescriptor: file_bag_v1_bag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bag_v1_bag_proto_goTypes,
		DependencyIndexes: file_bag_v1_bag_proto_depIdxs,
		MessageInfos:      file_bag_v1_bag_proto_msgTypes,
	}.Build()
	File_bag_v1_bag_proto = out.File
	file_bag_v1_bag_proto_rawDesc = nil
	file_bag_v1_bag_proto_goTypes = nil
	file_bag_v1_bag_proto_depIdxs = nil
}

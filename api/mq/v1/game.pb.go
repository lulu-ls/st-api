// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: mq/v1/game.proto

package v1

import (
	v1 "github.com/lulu-ls/st-api/api/item/v1"
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

type MqTopic int32

const (
	MqTopic_GameData     MqTopic = 0 // 游戏对局数据通知 飞机 炸弹 联队
	MqTopic_GameComplete MqTopic = 1 // 完成对局
	MqTopic_OrderPayed   MqTopic = 2 // 支付通知
	MqTopic_GetItem      MqTopic = 3 // 获得道具（钻石、金币、奖券）
	MqTopic_RaceSignup   MqTopic = 4 // 赛事报名
	MqTopic_FreeSignup   MqTopic = 5 // 自由场报名
	MqTopic_AssetEdit    MqTopic = 6 // 用户资产变动
	MqTopic_RaceStart    MqTopic = 7 // 比赛开始
)

// Enum value maps for MqTopic.
var (
	MqTopic_name = map[int32]string{
		0: "GameData",
		1: "GameComplete",
		2: "OrderPayed",
		3: "GetItem",
		4: "RaceSignup",
		5: "FreeSignup",
		6: "AssetEdit",
		7: "RaceStart",
	}
	MqTopic_value = map[string]int32{
		"GameData":     0,
		"GameComplete": 1,
		"OrderPayed":   2,
		"GetItem":      3,
		"RaceSignup":   4,
		"FreeSignup":   5,
		"AssetEdit":    6,
		"RaceStart":    7,
	}
)

func (x MqTopic) Enum() *MqTopic {
	p := new(MqTopic)
	*p = x
	return p
}

func (x MqTopic) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MqTopic) Descriptor() protoreflect.EnumDescriptor {
	return file_mq_v1_game_proto_enumTypes[0].Descriptor()
}

func (MqTopic) Type() protoreflect.EnumType {
	return &file_mq_v1_game_proto_enumTypes[0]
}

func (x MqTopic) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MqTopic.Descriptor instead.
func (MqTopic) EnumDescriptor() ([]byte, []int) {
	return file_mq_v1_game_proto_rawDescGZIP(), []int{0}
}

type SignupType int32

const (
	SignupType_CancelSignup SignupType = 0 // 取消报名
	SignupType_Signup       SignupType = 1 // 报名
)

// Enum value maps for SignupType.
var (
	SignupType_name = map[int32]string{
		0: "CancelSignup",
		1: "Signup",
	}
	SignupType_value = map[string]int32{
		"CancelSignup": 0,
		"Signup":       1,
	}
)

func (x SignupType) Enum() *SignupType {
	p := new(SignupType)
	*p = x
	return p
}

func (x SignupType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SignupType) Descriptor() protoreflect.EnumDescriptor {
	return file_mq_v1_game_proto_enumTypes[1].Descriptor()
}

func (SignupType) Type() protoreflect.EnumType {
	return &file_mq_v1_game_proto_enumTypes[1]
}

func (x SignupType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SignupType.Descriptor instead.
func (SignupType) EnumDescriptor() ([]byte, []int) {
	return file_mq_v1_game_proto_rawDescGZIP(), []int{1}
}

type GameDataType int32

const (
	GameDataType_FeiJi   GameDataType = 0 // 飞机
	GameDataType_ZhaDan  GameDataType = 1 // 炸弹
	GameDataType_LianDui GameDataType = 2 // 连对
	GameDataType_HuoJian GameDataType = 3 // 火箭
)

// Enum value maps for GameDataType.
var (
	GameDataType_name = map[int32]string{
		0: "FeiJi",
		1: "ZhaDan",
		2: "LianDui",
		3: "HuoJian",
	}
	GameDataType_value = map[string]int32{
		"FeiJi":   0,
		"ZhaDan":  1,
		"LianDui": 2,
		"HuoJian": 3,
	}
)

func (x GameDataType) Enum() *GameDataType {
	p := new(GameDataType)
	*p = x
	return p
}

func (x GameDataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameDataType) Descriptor() protoreflect.EnumDescriptor {
	return file_mq_v1_game_proto_enumTypes[2].Descriptor()
}

func (GameDataType) Type() protoreflect.EnumType {
	return &file_mq_v1_game_proto_enumTypes[2]
}

func (x GameDataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameDataType.Descriptor instead.
func (GameDataType) EnumDescriptor() ([]byte, []int) {
	return file_mq_v1_game_proto_rawDescGZIP(), []int{2}
}

type GetItemType int32

const (
	GetItemType__       GetItemType = 0
	GetItemType_Gold    GetItemType = 1 // 金币
	GetItemType_Diamond GetItemType = 2 // 钻石
	GetItemType_Ticket  GetItemType = 3 // 奖券
)

// Enum value maps for GetItemType.
var (
	GetItemType_name = map[int32]string{
		0: "_",
		1: "Gold",
		2: "Diamond",
		3: "Ticket",
	}
	GetItemType_value = map[string]int32{
		"_":       0,
		"Gold":    1,
		"Diamond": 2,
		"Ticket":  3,
	}
)

func (x GetItemType) Enum() *GetItemType {
	p := new(GetItemType)
	*p = x
	return p
}

func (x GetItemType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetItemType) Descriptor() protoreflect.EnumDescriptor {
	return file_mq_v1_game_proto_enumTypes[3].Descriptor()
}

func (GetItemType) Type() protoreflect.EnumType {
	return &file_mq_v1_game_proto_enumTypes[3]
}

func (x GetItemType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetItemType.Descriptor instead.
func (GetItemType) EnumDescriptor() ([]byte, []int) {
	return file_mq_v1_game_proto_rawDescGZIP(), []int{3}
}

type GameDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      GameDataType `protobuf:"varint,1,opt,name=type,proto3,enum=api.mq.v1.GameDataType" json:"type,omitempty"` // 数据类型
	UserId    int64        `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`           // 用户 id
	ChannelId int32        `protobuf:"varint,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`  // 渠道 id
	AppId     int32        `protobuf:"varint,4,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`              // app id
	Quantity  int32        `protobuf:"varint,5,opt,name=quantity,proto3" json:"quantity,omitempty"`                     // 数量
}

func (x *GameDataRequest) Reset() {
	*x = GameDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_v1_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameDataRequest) ProtoMessage() {}

func (x *GameDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mq_v1_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameDataRequest.ProtoReflect.Descriptor instead.
func (*GameDataRequest) Descriptor() ([]byte, []int) {
	return file_mq_v1_game_proto_rawDescGZIP(), []int{0}
}

func (x *GameDataRequest) GetType() GameDataType {
	if x != nil {
		return x.Type
	}
	return GameDataType_FeiJi
}

func (x *GameDataRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GameDataRequest) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *GameDataRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *GameDataRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type GameCompleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`          // 用户 id
	ChannelId int32 `protobuf:"varint,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"` // 渠道 id
	AppId     int32 `protobuf:"varint,4,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`             // app id
	IsWz      bool  `protobuf:"varint,5,opt,name=is_wz,json=isWz,proto3" json:"is_wz,omitempty"`                // 是否是挖主
}

func (x *GameCompleteRequest) Reset() {
	*x = GameCompleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_v1_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameCompleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameCompleteRequest) ProtoMessage() {}

func (x *GameCompleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mq_v1_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameCompleteRequest.ProtoReflect.Descriptor instead.
func (*GameCompleteRequest) Descriptor() ([]byte, []int) {
	return file_mq_v1_game_proto_rawDescGZIP(), []int{1}
}

func (x *GameCompleteRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GameCompleteRequest) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *GameCompleteRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *GameCompleteRequest) GetIsWz() bool {
	if x != nil {
		return x.IsWz
	}
	return false
}

type OrderPayedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64   `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`          // 用户 id
	ChannelId int32   `protobuf:"varint,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"` // 渠道 id
	AppId     int32   `protobuf:"varint,4,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`             // app id
	OrderId   int64   `protobuf:"varint,5,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`       // 订单号
	Price     float64 `protobuf:"fixed64,6,opt,name=price,proto3" json:"price,omitempty"`                         // 支付金额
}

func (x *OrderPayedRequest) Reset() {
	*x = OrderPayedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_v1_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPayedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPayedRequest) ProtoMessage() {}

func (x *OrderPayedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mq_v1_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPayedRequest.ProtoReflect.Descriptor instead.
func (*OrderPayedRequest) Descriptor() ([]byte, []int) {
	return file_mq_v1_game_proto_rawDescGZIP(), []int{2}
}

func (x *OrderPayedRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *OrderPayedRequest) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *OrderPayedRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *OrderPayedRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderPayedRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type GetItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`          // 用户 id
	ChannelId int32 `protobuf:"varint,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"` // 渠道 id
	AppId     int32 `protobuf:"varint,4,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`             // app id
	ItemId    int32 `protobuf:"varint,5,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`          // 道具类型
	Quantity  int32 `protobuf:"varint,6,opt,name=quantity,proto3" json:"quantity,omitempty"`                    // 获得道具数量
}

func (x *GetItemRequest) Reset() {
	*x = GetItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_v1_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemRequest) ProtoMessage() {}

func (x *GetItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mq_v1_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemRequest.ProtoReflect.Descriptor instead.
func (*GetItemRequest) Descriptor() ([]byte, []int) {
	return file_mq_v1_game_proto_rawDescGZIP(), []int{3}
}

func (x *GetItemRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetItemRequest) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *GetItemRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *GetItemRequest) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *GetItemRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type RaceSignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int64      `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                     // 用户 id
	GameConfigId int32      `protobuf:"varint,2,opt,name=game_config_id,json=gameConfigId,proto3" json:"game_config_id,omitempty"` // 赛事配置 id
	Type         SignupType `protobuf:"varint,3,opt,name=type,proto3,enum=api.mq.v1.SignupType" json:"type,omitempty"`             // 报名类型 0 取消 1 报名
}

func (x *RaceSignUpRequest) Reset() {
	*x = RaceSignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_v1_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaceSignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaceSignUpRequest) ProtoMessage() {}

func (x *RaceSignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mq_v1_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaceSignUpRequest.ProtoReflect.Descriptor instead.
func (*RaceSignUpRequest) Descriptor() ([]byte, []int) {
	return file_mq_v1_game_proto_rawDescGZIP(), []int{4}
}

func (x *RaceSignUpRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RaceSignUpRequest) GetGameConfigId() int32 {
	if x != nil {
		return x.GameConfigId
	}
	return 0
}

func (x *RaceSignUpRequest) GetType() SignupType {
	if x != nil {
		return x.Type
	}
	return SignupType_CancelSignup
}

type FreeSignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId           int64      `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                                   // 用户 id
	GameFreeConfigId int32      `protobuf:"varint,2,opt,name=game_free_config_id,json=gameFreeConfigId,proto3" json:"game_free_config_id,omitempty"` // 自由场 id
	Type             SignupType `protobuf:"varint,3,opt,name=type,proto3,enum=api.mq.v1.SignupType" json:"type,omitempty"`                           // 报名类型 0 取消 1 报名
}

func (x *FreeSignUpRequest) Reset() {
	*x = FreeSignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_v1_game_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FreeSignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FreeSignUpRequest) ProtoMessage() {}

func (x *FreeSignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mq_v1_game_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FreeSignUpRequest.ProtoReflect.Descriptor instead.
func (*FreeSignUpRequest) Descriptor() ([]byte, []int) {
	return file_mq_v1_game_proto_rawDescGZIP(), []int{5}
}

func (x *FreeSignUpRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FreeSignUpRequest) GetGameFreeConfigId() int32 {
	if x != nil {
		return x.GameFreeConfigId
	}
	return 0
}

func (x *FreeSignUpRequest) GetType() SignupType {
	if x != nil {
		return x.Type
	}
	return SignupType_CancelSignup
}

type AssetEditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64           `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户 id
	SeqId  int64           `protobuf:"varint,2,opt,name=seq_id,json=seqId,proto3" json:"seq_id,omitempty"`    // 唯一 id
	Item   []*v1.AssetItem `protobuf:"bytes,3,rep,name=item,proto3" json:"item,omitempty"`                    // 变动详情
}

func (x *AssetEditRequest) Reset() {
	*x = AssetEditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_v1_game_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetEditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetEditRequest) ProtoMessage() {}

func (x *AssetEditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mq_v1_game_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetEditRequest.ProtoReflect.Descriptor instead.
func (*AssetEditRequest) Descriptor() ([]byte, []int) {
	return file_mq_v1_game_proto_rawDescGZIP(), []int{6}
}

func (x *AssetEditRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AssetEditRequest) GetSeqId() int64 {
	if x != nil {
		return x.SeqId
	}
	return 0
}

func (x *AssetEditRequest) GetItem() []*v1.AssetItem {
	if x != nil {
		return x.Item
	}
	return nil
}

type RaceStartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameConfigId int32                `protobuf:"varint,1,opt,name=game_config_id,json=gameConfigId,proto3" json:"game_config_id,omitempty"` // 赛事配置 id
	GameId       int32                `protobuf:"varint,2,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`                     // 比赛 id
	UserList     []*RaceStartUserItem `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`                // 用户
}

func (x *RaceStartRequest) Reset() {
	*x = RaceStartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_v1_game_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaceStartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaceStartRequest) ProtoMessage() {}

func (x *RaceStartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mq_v1_game_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaceStartRequest.ProtoReflect.Descriptor instead.
func (*RaceStartRequest) Descriptor() ([]byte, []int) {
	return file_mq_v1_game_proto_rawDescGZIP(), []int{7}
}

func (x *RaceStartRequest) GetGameConfigId() int32 {
	if x != nil {
		return x.GameConfigId
	}
	return 0
}

func (x *RaceStartRequest) GetGameId() int32 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *RaceStartRequest) GetUserList() []*RaceStartUserItem {
	if x != nil {
		return x.UserList
	}
	return nil
}

type RaceStartUserItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户 id
}

func (x *RaceStartUserItem) Reset() {
	*x = RaceStartUserItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_v1_game_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaceStartUserItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaceStartUserItem) ProtoMessage() {}

func (x *RaceStartUserItem) ProtoReflect() protoreflect.Message {
	mi := &file_mq_v1_game_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaceStartUserItem.ProtoReflect.Descriptor instead.
func (*RaceStartUserItem) Descriptor() ([]byte, []int) {
	return file_mq_v1_game_proto_rawDescGZIP(), []int{8}
}

func (x *RaceStartUserItem) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_mq_v1_game_proto protoreflect.FileDescriptor

var file_mq_v1_game_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x71, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x71, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x69,
	0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x0f, 0x47, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x71, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x61,
	0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x70, 0x70,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x79,
	0x0a, 0x13, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x15, 0x0a,
	0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61,
	0x70, 0x70, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x69, 0x73, 0x5f, 0x77, 0x7a, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x57, 0x7a, 0x22, 0x93, 0x01, 0x0a, 0x11, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22,
	0x94, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x7d, 0x0a, 0x11, 0x52, 0x61, 0x63, 0x65, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x67, 0x61,
	0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x71, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x11, 0x46, 0x72, 0x65, 0x65, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x13, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x66, 0x72, 0x65,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x10, 0x67, 0x61, 0x6d, 0x65, 0x46, 0x72, 0x65, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x6e,
	0x0a, 0x10, 0x41, 0x73, 0x73, 0x65, 0x74, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73,
	0x65, 0x71, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x65, 0x71,
	0x49, 0x64, 0x12, 0x2a, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x8c,
	0x01, 0x0a, 0x10, 0x52, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x67, 0x61, 0x6d,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65,
	0x49, 0x64, 0x12, 0x39, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x71, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x2c, 0x0a,
	0x11, 0x52, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x2a, 0x84, 0x01, 0x0a, 0x07,
	0x4d, 0x71, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x0c, 0x0a, 0x08, 0x47, 0x61, 0x6d, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x50, 0x61, 0x79, 0x65, 0x64, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x61, 0x63, 0x65, 0x53, 0x69, 0x67, 0x6e,
	0x75, 0x70, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x72, 0x65, 0x65, 0x53, 0x69, 0x67, 0x6e,
	0x75, 0x70, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x73, 0x73, 0x65, 0x74, 0x45, 0x64, 0x69,
	0x74, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x10, 0x07, 0x2a, 0x2a, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x10, 0x0a, 0x0c, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x10, 0x01, 0x2a, 0x3f,
	0x0a, 0x0c, 0x47, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09,
	0x0a, 0x05, 0x46, 0x65, 0x69, 0x4a, 0x69, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x5a, 0x68, 0x61,
	0x44, 0x61, 0x6e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x69, 0x61, 0x6e, 0x44, 0x75, 0x69,
	0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x75, 0x6f, 0x4a, 0x69, 0x61, 0x6e, 0x10, 0x03, 0x2a,
	0x37, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x05,
	0x0a, 0x01, 0x5f, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x6f, 0x6c, 0x64, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x44, 0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x10, 0x03, 0x42, 0x35, 0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e,
	0x6d, 0x71, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x75, 0x6c, 0x75, 0x2d, 0x6c, 0x73, 0x2f, 0x73, 0x74, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x71, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mq_v1_game_proto_rawDescOnce sync.Once
	file_mq_v1_game_proto_rawDescData = file_mq_v1_game_proto_rawDesc
)

func file_mq_v1_game_proto_rawDescGZIP() []byte {
	file_mq_v1_game_proto_rawDescOnce.Do(func() {
		file_mq_v1_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_mq_v1_game_proto_rawDescData)
	})
	return file_mq_v1_game_proto_rawDescData
}

var file_mq_v1_game_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_mq_v1_game_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_mq_v1_game_proto_goTypes = []interface{}{
	(MqTopic)(0),                // 0: api.mq.v1.MqTopic
	(SignupType)(0),             // 1: api.mq.v1.SignupType
	(GameDataType)(0),           // 2: api.mq.v1.GameDataType
	(GetItemType)(0),            // 3: api.mq.v1.GetItemType
	(*GameDataRequest)(nil),     // 4: api.mq.v1.GameDataRequest
	(*GameCompleteRequest)(nil), // 5: api.mq.v1.GameCompleteRequest
	(*OrderPayedRequest)(nil),   // 6: api.mq.v1.OrderPayedRequest
	(*GetItemRequest)(nil),      // 7: api.mq.v1.GetItemRequest
	(*RaceSignUpRequest)(nil),   // 8: api.mq.v1.RaceSignUpRequest
	(*FreeSignUpRequest)(nil),   // 9: api.mq.v1.FreeSignUpRequest
	(*AssetEditRequest)(nil),    // 10: api.mq.v1.AssetEditRequest
	(*RaceStartRequest)(nil),    // 11: api.mq.v1.RaceStartRequest
	(*RaceStartUserItem)(nil),   // 12: api.mq.v1.RaceStartUserItem
	(*v1.AssetItem)(nil),        // 13: api.item.v1.AssetItem
}
var file_mq_v1_game_proto_depIdxs = []int32{
	2,  // 0: api.mq.v1.GameDataRequest.type:type_name -> api.mq.v1.GameDataType
	1,  // 1: api.mq.v1.RaceSignUpRequest.type:type_name -> api.mq.v1.SignupType
	1,  // 2: api.mq.v1.FreeSignUpRequest.type:type_name -> api.mq.v1.SignupType
	13, // 3: api.mq.v1.AssetEditRequest.item:type_name -> api.item.v1.AssetItem
	12, // 4: api.mq.v1.RaceStartRequest.user_list:type_name -> api.mq.v1.RaceStartUserItem
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_mq_v1_game_proto_init() }
func file_mq_v1_game_proto_init() {
	if File_mq_v1_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mq_v1_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameDataRequest); i {
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
		file_mq_v1_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameCompleteRequest); i {
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
		file_mq_v1_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPayedRequest); i {
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
		file_mq_v1_game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemRequest); i {
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
		file_mq_v1_game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaceSignUpRequest); i {
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
		file_mq_v1_game_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FreeSignUpRequest); i {
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
		file_mq_v1_game_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetEditRequest); i {
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
		file_mq_v1_game_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaceStartRequest); i {
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
		file_mq_v1_game_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaceStartUserItem); i {
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
			RawDescriptor: file_mq_v1_game_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mq_v1_game_proto_goTypes,
		DependencyIndexes: file_mq_v1_game_proto_depIdxs,
		EnumInfos:         file_mq_v1_game_proto_enumTypes,
		MessageInfos:      file_mq_v1_game_proto_msgTypes,
	}.Build()
	File_mq_v1_game_proto = out.File
	file_mq_v1_game_proto_rawDesc = nil
	file_mq_v1_game_proto_goTypes = nil
	file_mq_v1_game_proto_depIdxs = nil
}

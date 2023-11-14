// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: user/v1/user.proto

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

type RaceRecordListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   *v1.Paginate `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`                    // 分页数据
	UserId int64        `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户 id
	AppId  int32        `protobuf:"varint,3,opt,name=appId,proto3" json:"appId,omitempty"`                 // 渠道 id
}

func (x *RaceRecordListRequest) Reset() {
	*x = RaceRecordListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaceRecordListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaceRecordListRequest) ProtoMessage() {}

func (x *RaceRecordListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaceRecordListRequest.ProtoReflect.Descriptor instead.
func (*RaceRecordListRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *RaceRecordListRequest) GetPage() *v1.Paginate {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *RaceRecordListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RaceRecordListRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

type RaceRecordListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*RaceRecordListItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Total int32                 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"` // 合计
}

func (x *RaceRecordListReply) Reset() {
	*x = RaceRecordListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaceRecordListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaceRecordListReply) ProtoMessage() {}

func (x *RaceRecordListReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaceRecordListReply.ProtoReflect.Descriptor instead.
func (*RaceRecordListReply) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *RaceRecordListReply) GetList() []*RaceRecordListItem {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *RaceRecordListReply) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type RaceRecordListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                      // 记录 id
	GameId           int32  `protobuf:"varint,2,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`                                // 游戏id
	GameCategoryId   int32  `protobuf:"varint,3,opt,name=game_category_id,json=gameCategoryId,proto3" json:"game_category_id,omitempty"`      // 比赛类型
	GameCategoryName string `protobuf:"bytes,4,opt,name=game_category_name,json=gameCategoryName,proto3" json:"game_category_name,omitempty"` // 比赛类型名称
	GameConfigId     int32  `protobuf:"varint,5,opt,name=game_config_id,json=gameConfigId,proto3" json:"game_config_id,omitempty"`            // 比赛配置 id
	GameConfigName   string `protobuf:"bytes,6,opt,name=game_config_name,json=gameConfigName,proto3" json:"game_config_name,omitempty"`       // 比赛配置名称
	GameConfigTitle  string `protobuf:"bytes,7,opt,name=game_config_title,json=gameConfigTitle,proto3" json:"game_config_title,omitempty"`    // 比赛配置标题
	Rounds           int32  `protobuf:"varint,8,opt,name=rounds,proto3" json:"rounds,omitempty"`                                              // 进入轮次
	Rank             int32  `protobuf:"varint,9,opt,name=rank,proto3" json:"rank,omitempty"`                                                  // 排名
	Win              int32  `protobuf:"varint,10,opt,name=win,proto3" json:"win,omitempty"`                                                   // 胜利场次
	Lose             int32  `protobuf:"varint,11,opt,name=lose,proto3" json:"lose,omitempty"`                                                 // 失败场次
	EnterTime        int64  `protobuf:"varint,12,opt,name=enter_time,json=enterTime,proto3" json:"enter_time,omitempty"`                      // 开始时间
	LeaveTime        int64  `protobuf:"varint,13,opt,name=leave_time,json=leaveTime,proto3" json:"leave_time,omitempty"`                      // 离开时间
}

func (x *RaceRecordListItem) Reset() {
	*x = RaceRecordListItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaceRecordListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaceRecordListItem) ProtoMessage() {}

func (x *RaceRecordListItem) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaceRecordListItem.ProtoReflect.Descriptor instead.
func (*RaceRecordListItem) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *RaceRecordListItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RaceRecordListItem) GetGameId() int32 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *RaceRecordListItem) GetGameCategoryId() int32 {
	if x != nil {
		return x.GameCategoryId
	}
	return 0
}

func (x *RaceRecordListItem) GetGameCategoryName() string {
	if x != nil {
		return x.GameCategoryName
	}
	return ""
}

func (x *RaceRecordListItem) GetGameConfigId() int32 {
	if x != nil {
		return x.GameConfigId
	}
	return 0
}

func (x *RaceRecordListItem) GetGameConfigName() string {
	if x != nil {
		return x.GameConfigName
	}
	return ""
}

func (x *RaceRecordListItem) GetGameConfigTitle() string {
	if x != nil {
		return x.GameConfigTitle
	}
	return ""
}

func (x *RaceRecordListItem) GetRounds() int32 {
	if x != nil {
		return x.Rounds
	}
	return 0
}

func (x *RaceRecordListItem) GetRank() int32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *RaceRecordListItem) GetWin() int32 {
	if x != nil {
		return x.Win
	}
	return 0
}

func (x *RaceRecordListItem) GetLose() int32 {
	if x != nil {
		return x.Lose
	}
	return 0
}

func (x *RaceRecordListItem) GetEnterTime() int64 {
	if x != nil {
		return x.EnterTime
	}
	return 0
}

func (x *RaceRecordListItem) GetLeaveTime() int64 {
	if x != nil {
		return x.LeaveTime
	}
	return 0
}

type FigureEditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户 id
	AppId     int32 `protobuf:"varint,2,opt,name=appId,proto3" json:"appId,omitempty"`                 // 渠道 id
	ChannelId int32 `protobuf:"varint,3,opt,name=channelId,proto3" json:"channelId,omitempty"`         // 渠道 id
	Figure    int32 `protobuf:"varint,4,opt,name=figure,proto3" json:"figure,omitempty"`               // 形象
}

func (x *FigureEditRequest) Reset() {
	*x = FigureEditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FigureEditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FigureEditRequest) ProtoMessage() {}

func (x *FigureEditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FigureEditRequest.ProtoReflect.Descriptor instead.
func (*FigureEditRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *FigureEditRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FigureEditRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *FigureEditRequest) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *FigureEditRequest) GetFigure() int32 {
	if x != nil {
		return x.Figure
	}
	return 0
}

type FigureEditReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FigureEditReply) Reset() {
	*x = FigureEditReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FigureEditReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FigureEditReply) ProtoMessage() {}

func (x *FigureEditReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FigureEditReply.ProtoReflect.Descriptor instead.
func (*FigureEditReply) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{4}
}

type AssetGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户 id
	AppId     int32 `protobuf:"varint,2,opt,name=appId,proto3" json:"appId,omitempty"`                 // 渠道 id
	ChannelId int32 `protobuf:"varint,3,opt,name=channelId,proto3" json:"channelId,omitempty"`         // 渠道 id
}

func (x *AssetGetRequest) Reset() {
	*x = AssetGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetGetRequest) ProtoMessage() {}

func (x *AssetGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetGetRequest.ProtoReflect.Descriptor instead.
func (*AssetGetRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{5}
}

func (x *AssetGetRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AssetGetRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *AssetGetRequest) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

type AssetGetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gold        int32 `protobuf:"varint,1,opt,name=gold,proto3" json:"gold,omitempty"`                                  // 金币数量
	Diamond     int32 `protobuf:"varint,2,opt,name=diamond,proto3" json:"diamond,omitempty"`                            // 钻石数量
	Ticket      int32 `protobuf:"varint,3,opt,name=ticket,proto3" json:"ticket,omitempty"`                              // 唐卡数量
	CardCounter int32 `protobuf:"varint,4,opt,name=card_counter,json=cardCounter,proto3" json:"card_counter,omitempty"` // 记牌器数量
}

func (x *AssetGetReply) Reset() {
	*x = AssetGetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetGetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetGetReply) ProtoMessage() {}

func (x *AssetGetReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetGetReply.ProtoReflect.Descriptor instead.
func (*AssetGetReply) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{6}
}

func (x *AssetGetReply) GetGold() int32 {
	if x != nil {
		return x.Gold
	}
	return 0
}

func (x *AssetGetReply) GetDiamond() int32 {
	if x != nil {
		return x.Diamond
	}
	return 0
}

func (x *AssetGetReply) GetTicket() int32 {
	if x != nil {
		return x.Ticket
	}
	return 0
}

func (x *AssetGetReply) GetCardCounter() int32 {
	if x != nil {
		return x.CardCounter
	}
	return 0
}

var File_user_v1_user_proto protoreflect.FileDescriptor

var file_user_v1_user_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x73, 0x0a, 0x15, 0x52, 0x61, 0x63, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0x60, 0x0a, 0x13, 0x52, 0x61, 0x63, 0x65, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x33, 0x0a, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x63, 0x65, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xa1, 0x03, 0x0a, 0x12, 0x52, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x67, 0x61, 0x6d, 0x65, 0x5f,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0e, 0x67, 0x61, 0x6d, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49,
	0x64, 0x12, 0x2c, 0x0a, 0x12, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x67,
	0x61, 0x6d, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x24, 0x0a, 0x0e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x67, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x67, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x2a, 0x0a, 0x11, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x67, 0x61, 0x6d, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x69, 0x6e, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x77, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x73,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x6c, 0x65, 0x61, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x11,
	0x46, 0x69, 0x67, 0x75, 0x72, 0x65, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x06, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x06, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x22,
	0x11, 0x0a, 0x0f, 0x46, 0x69, 0x67, 0x75, 0x72, 0x65, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x5e, 0x0a, 0x0f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61,
	0x70, 0x70, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x49, 0x64, 0x22, 0x78, 0x0a, 0x0d, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x67, 0x6f, 0x6c, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x61, 0x6d, 0x6f,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x64, 0x69, 0x61, 0x6d, 0x6f, 0x6e,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x72,
	0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x63, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x32, 0xed, 0x02, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x80, 0x01, 0x0a, 0x0e, 0x52, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x63, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x28,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f, 0x73, 0x74, 0x2d, 0x67,
	0x61, 0x6d, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x61, 0x63,
	0x65, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x74, 0x0a, 0x0a, 0x46, 0x69, 0x67, 0x75,
	0x72, 0x65, 0x45, 0x64, 0x69, 0x74, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x67, 0x75, 0x72, 0x65, 0x45, 0x64, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x67, 0x75, 0x72, 0x65, 0x45, 0x64, 0x69, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22,
	0x1d, 0x2f, 0x73, 0x74, 0x2d, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x2f, 0x65, 0x64, 0x69, 0x74, 0x12, 0x6c,
	0x0a, 0x08, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22,
	0x1b, 0x2f, 0x73, 0x74, 0x2d, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x42, 0x39, 0x0a, 0x0b,
	0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x28, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x75, 0x6c, 0x75, 0x2d, 0x6c,
	0x73, 0x2f, 0x73, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_v1_user_proto_rawDescOnce sync.Once
	file_user_v1_user_proto_rawDescData = file_user_v1_user_proto_rawDesc
)

func file_user_v1_user_proto_rawDescGZIP() []byte {
	file_user_v1_user_proto_rawDescOnce.Do(func() {
		file_user_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_v1_user_proto_rawDescData)
	})
	return file_user_v1_user_proto_rawDescData
}

var file_user_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_user_v1_user_proto_goTypes = []interface{}{
	(*RaceRecordListRequest)(nil), // 0: api.user.v1.RaceRecordListRequest
	(*RaceRecordListReply)(nil),   // 1: api.user.v1.RaceRecordListReply
	(*RaceRecordListItem)(nil),    // 2: api.user.v1.RaceRecordListItem
	(*FigureEditRequest)(nil),     // 3: api.user.v1.FigureEditRequest
	(*FigureEditReply)(nil),       // 4: api.user.v1.FigureEditReply
	(*AssetGetRequest)(nil),       // 5: api.user.v1.AssetGetRequest
	(*AssetGetReply)(nil),         // 6: api.user.v1.AssetGetReply
	(*v1.Paginate)(nil),           // 7: api.common.v1.Paginate
}
var file_user_v1_user_proto_depIdxs = []int32{
	7, // 0: api.user.v1.RaceRecordListRequest.page:type_name -> api.common.v1.Paginate
	2, // 1: api.user.v1.RaceRecordListReply.list:type_name -> api.user.v1.RaceRecordListItem
	0, // 2: api.user.v1.User.RaceRecordList:input_type -> api.user.v1.RaceRecordListRequest
	3, // 3: api.user.v1.User.FigureEdit:input_type -> api.user.v1.FigureEditRequest
	5, // 4: api.user.v1.User.AssetGet:input_type -> api.user.v1.AssetGetRequest
	1, // 5: api.user.v1.User.RaceRecordList:output_type -> api.user.v1.RaceRecordListReply
	4, // 6: api.user.v1.User.FigureEdit:output_type -> api.user.v1.FigureEditReply
	6, // 7: api.user.v1.User.AssetGet:output_type -> api.user.v1.AssetGetReply
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_v1_user_proto_init() }
func file_user_v1_user_proto_init() {
	if File_user_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaceRecordListRequest); i {
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
		file_user_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaceRecordListReply); i {
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
		file_user_v1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaceRecordListItem); i {
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
		file_user_v1_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FigureEditRequest); i {
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
		file_user_v1_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FigureEditReply); i {
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
		file_user_v1_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetGetRequest); i {
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
		file_user_v1_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetGetReply); i {
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
			RawDescriptor: file_user_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_v1_user_proto_goTypes,
		DependencyIndexes: file_user_v1_user_proto_depIdxs,
		MessageInfos:      file_user_v1_user_proto_msgTypes,
	}.Build()
	File_user_v1_user_proto = out.File
	file_user_v1_user_proto_rawDesc = nil
	file_user_v1_user_proto_goTypes = nil
	file_user_v1_user_proto_depIdxs = nil
}

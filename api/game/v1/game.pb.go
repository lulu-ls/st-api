// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: game/v1/game.proto

package v1

import (
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

// 任务进度状态
type TaskStatus int32

const (
	TaskStatus_undone   TaskStatus = 0 // 未完成
	TaskStatus_done     TaskStatus = 1 // 已完成
	TaskStatus_received TaskStatus = 2 // 已发放奖励
)

// Enum value maps for TaskStatus.
var (
	TaskStatus_name = map[int32]string{
		0: "undone",
		1: "done",
		2: "received",
	}
	TaskStatus_value = map[string]int32{
		"undone":   0,
		"done":     1,
		"received": 2,
	}
)

func (x TaskStatus) Enum() *TaskStatus {
	p := new(TaskStatus)
	*p = x
	return p
}

func (x TaskStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_game_v1_game_proto_enumTypes[0].Descriptor()
}

func (TaskStatus) Type() protoreflect.EnumType {
	return &file_game_v1_game_proto_enumTypes[0]
}

func (x TaskStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskStatus.Descriptor instead.
func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return file_game_v1_game_proto_rawDescGZIP(), []int{0}
}

// 任务刷新周期 0无 1 天 2周 3 月
type RefreshCycle int32

const (
	RefreshCycle_none  RefreshCycle = 0
	RefreshCycle_day   RefreshCycle = 1
	RefreshCycle_week  RefreshCycle = 2
	RefreshCycle_month RefreshCycle = 3
)

// Enum value maps for RefreshCycle.
var (
	RefreshCycle_name = map[int32]string{
		0: "none",
		1: "day",
		2: "week",
		3: "month",
	}
	RefreshCycle_value = map[string]int32{
		"none":  0,
		"day":   1,
		"week":  2,
		"month": 3,
	}
)

func (x RefreshCycle) Enum() *RefreshCycle {
	p := new(RefreshCycle)
	*p = x
	return p
}

func (x RefreshCycle) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RefreshCycle) Descriptor() protoreflect.EnumDescriptor {
	return file_game_v1_game_proto_enumTypes[1].Descriptor()
}

func (RefreshCycle) Type() protoreflect.EnumType {
	return &file_game_v1_game_proto_enumTypes[1]
}

func (x RefreshCycle) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RefreshCycle.Descriptor instead.
func (RefreshCycle) EnumDescriptor() ([]byte, []int) {
	return file_game_v1_game_proto_rawDescGZIP(), []int{1}
}

type AnnouncementListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      *v1.Paginate `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	AppId     int32        `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`             // app id
	ChannelId int32        `protobuf:"varint,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"` // channel id
}

func (x *AnnouncementListRequest) Reset() {
	*x = AnnouncementListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_v1_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnouncementListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnouncementListRequest) ProtoMessage() {}

func (x *AnnouncementListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnouncementListRequest.ProtoReflect.Descriptor instead.
func (*AnnouncementListRequest) Descriptor() ([]byte, []int) {
	return file_game_v1_game_proto_rawDescGZIP(), []int{0}
}

func (x *AnnouncementListRequest) GetPage() *v1.Paginate {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *AnnouncementListRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *AnnouncementListRequest) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

type AnnouncementListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*AnnouncementListItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`    // 列表
	Total int32                   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"` // 合计
}

func (x *AnnouncementListReply) Reset() {
	*x = AnnouncementListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_v1_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnouncementListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnouncementListReply) ProtoMessage() {}

func (x *AnnouncementListReply) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnouncementListReply.ProtoReflect.Descriptor instead.
func (*AnnouncementListReply) Descriptor() ([]byte, []int) {
	return file_game_v1_game_proto_rawDescGZIP(), []int{1}
}

func (x *AnnouncementListReply) GetList() []*AnnouncementListItem {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *AnnouncementListReply) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type AnnouncementListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                  // 通知 id
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`             // 标题
	Information string `protobuf:"bytes,3,opt,name=information,proto3" json:"information,omitempty"` // 内容
	Image       string `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`             // 图片
}

func (x *AnnouncementListItem) Reset() {
	*x = AnnouncementListItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_v1_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnouncementListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnouncementListItem) ProtoMessage() {}

func (x *AnnouncementListItem) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnouncementListItem.ProtoReflect.Descriptor instead.
func (*AnnouncementListItem) Descriptor() ([]byte, []int) {
	return file_game_v1_game_proto_rawDescGZIP(), []int{2}
}

func (x *AnnouncementListItem) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AnnouncementListItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AnnouncementListItem) GetInformation() string {
	if x != nil {
		return x.Information
	}
	return ""
}

func (x *AnnouncementListItem) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type TaskListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      *v1.Paginate `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	AppId     int32        `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`             // app id
	ChannelId int32        `protobuf:"varint,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"` // channel id
	UserId    int64        `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`          // 用户 id
}

func (x *TaskListRequest) Reset() {
	*x = TaskListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_v1_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListRequest) ProtoMessage() {}

func (x *TaskListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListRequest.ProtoReflect.Descriptor instead.
func (*TaskListRequest) Descriptor() ([]byte, []int) {
	return file_game_v1_game_proto_rawDescGZIP(), []int{3}
}

func (x *TaskListRequest) GetPage() *v1.Paginate {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *TaskListRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *TaskListRequest) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *TaskListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type TaskListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*TaskListItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`    // 列表
	Total int32           `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"` // 数量
}

func (x *TaskListReply) Reset() {
	*x = TaskListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_v1_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListReply) ProtoMessage() {}

func (x *TaskListReply) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListReply.ProtoReflect.Descriptor instead.
func (*TaskListReply) Descriptor() ([]byte, []int) {
	return file_game_v1_game_proto_rawDescGZIP(), []int{4}
}

func (x *TaskListReply) GetList() []*TaskListItem {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *TaskListReply) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type TaskListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId       int32        `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`                                                  // 任务 id
	Title        string       `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`                                                                   // 任务标题
	ItemId       int32        `protobuf:"varint,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`                                                  // 道具 id
	ItemName     string       `protobuf:"bytes,4,opt,name=item_name,json=itemName,proto3" json:"item_name,omitempty"`                                             // 奖励道具名称
	ItemImage    string       `protobuf:"bytes,5,opt,name=item_image,json=itemImage,proto3" json:"item_image,omitempty"`                                          // 奖励道具图片
	ItemQuantity int32        `protobuf:"varint,6,opt,name=item_quantity,json=itemQuantity,proto3" json:"item_quantity,omitempty"`                                // 奖励道具数量
	Status       TaskStatus   `protobuf:"varint,7,opt,name=status,proto3,enum=api.game.v1.TaskStatus" json:"status,omitempty"`                                    // 状态
	Number       int32        `protobuf:"varint,8,opt,name=number,proto3" json:"number,omitempty"`                                                                // 任务需要完成数量
	Process      int32        `protobuf:"varint,9,opt,name=process,proto3" json:"process,omitempty"`                                                              // 当前完成进度
	RefreshCycle RefreshCycle `protobuf:"varint,10,opt,name=refresh_cycle,json=refreshCycle,proto3,enum=api.game.v1.RefreshCycle" json:"refresh_cycle,omitempty"` // 刷新周期
	GameType     int32        `protobuf:"varint,11,opt,name=game_type,json=gameType,proto3" json:"game_type,omitempty"`                                           // 游戏类型：1=竞技场，2=自由场
	Action       int32        `protobuf:"varint,12,opt,name=action,proto3" json:"action,omitempty"`                                                               // 跳转类型：1:url 2:game_config_id 3:game_type_id 4:shop 5:task 6:share 7:task
	Url          string       `protobuf:"bytes,13,opt,name=url,proto3" json:"url,omitempty"`                                                                      // 跳转地址 action = 1 可用
}

func (x *TaskListItem) Reset() {
	*x = TaskListItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_v1_game_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListItem) ProtoMessage() {}

func (x *TaskListItem) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1_game_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListItem.ProtoReflect.Descriptor instead.
func (*TaskListItem) Descriptor() ([]byte, []int) {
	return file_game_v1_game_proto_rawDescGZIP(), []int{5}
}

func (x *TaskListItem) GetTaskId() int32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *TaskListItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TaskListItem) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *TaskListItem) GetItemName() string {
	if x != nil {
		return x.ItemName
	}
	return ""
}

func (x *TaskListItem) GetItemImage() string {
	if x != nil {
		return x.ItemImage
	}
	return ""
}

func (x *TaskListItem) GetItemQuantity() int32 {
	if x != nil {
		return x.ItemQuantity
	}
	return 0
}

func (x *TaskListItem) GetStatus() TaskStatus {
	if x != nil {
		return x.Status
	}
	return TaskStatus_undone
}

func (x *TaskListItem) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *TaskListItem) GetProcess() int32 {
	if x != nil {
		return x.Process
	}
	return 0
}

func (x *TaskListItem) GetRefreshCycle() RefreshCycle {
	if x != nil {
		return x.RefreshCycle
	}
	return RefreshCycle_none
}

func (x *TaskListItem) GetGameType() int32 {
	if x != nil {
		return x.GameType
	}
	return 0
}

func (x *TaskListItem) GetAction() int32 {
	if x != nil {
		return x.Action
	}
	return 0
}

func (x *TaskListItem) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type TaskRewardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户 id
	TaskId int32 `protobuf:"varint,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"` // 任务 id
}

func (x *TaskRewardRequest) Reset() {
	*x = TaskRewardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_v1_game_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRewardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRewardRequest) ProtoMessage() {}

func (x *TaskRewardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1_game_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRewardRequest.ProtoReflect.Descriptor instead.
func (*TaskRewardRequest) Descriptor() ([]byte, []int) {
	return file_game_v1_game_proto_rawDescGZIP(), []int{6}
}

func (x *TaskRewardRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *TaskRewardRequest) GetTaskId() int32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type TaskRewardReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TaskRewardReply) Reset() {
	*x = TaskRewardReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_v1_game_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRewardReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRewardReply) ProtoMessage() {}

func (x *TaskRewardReply) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1_game_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRewardReply.ProtoReflect.Descriptor instead.
func (*TaskRewardReply) Descriptor() ([]byte, []int) {
	return file_game_v1_game_proto_rawDescGZIP(), []int{7}
}

var File_game_v1_game_proto protoreflect.FileDescriptor

var file_game_v1_game_proto_rawDesc = []byte{
	0x0a, 0x12, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x17, 0x41, 0x6e, 0x6e, 0x6f, 0x75,
	0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x49, 0x64, 0x22, 0x64, 0x0a, 0x15, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x35,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75,
	0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x74, 0x0a, 0x14, 0x41,
	0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x22, 0x8d, 0x01, 0x0a, 0x0f, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x54, 0x0a, 0x0d, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x2d, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xa1, 0x03, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x69, 0x74, 0x65, 0x6d, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x3e, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43,
	0x79, 0x63, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x45, 0x0a, 0x11, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b,
	0x49, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x2a, 0x30, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x75, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x64, 0x10, 0x02, 0x2a, 0x36, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x6e, 0x6f, 0x6e, 0x65, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x77, 0x65,
	0x65, 0x6b, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x10, 0x03, 0x32,
	0xf9, 0x02, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x8c, 0x01, 0x0a, 0x10, 0x41, 0x6e, 0x6e,
	0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x6e, 0x6f,
	0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a,
	0x01, 0x2a, 0x22, 0x23, 0x2f, 0x73, 0x74, 0x2d, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x6c, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x26, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x73, 0x74, 0x2d, 0x67, 0x61,
	0x6d, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x74, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f, 0x73,
	0x74, 0x2d, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f,
	0x74, 0x61, 0x73, 0x6b, 0x2f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x42, 0x39, 0x0a, 0x0b, 0x61,
	0x70, 0x69, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x75, 0x6c, 0x75, 0x2d, 0x6c, 0x73,
	0x2f, 0x73, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x6d, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_v1_game_proto_rawDescOnce sync.Once
	file_game_v1_game_proto_rawDescData = file_game_v1_game_proto_rawDesc
)

func file_game_v1_game_proto_rawDescGZIP() []byte {
	file_game_v1_game_proto_rawDescOnce.Do(func() {
		file_game_v1_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_v1_game_proto_rawDescData)
	})
	return file_game_v1_game_proto_rawDescData
}

var file_game_v1_game_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_game_v1_game_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_game_v1_game_proto_goTypes = []interface{}{
	(TaskStatus)(0),                 // 0: api.game.v1.TaskStatus
	(RefreshCycle)(0),               // 1: api.game.v1.RefreshCycle
	(*AnnouncementListRequest)(nil), // 2: api.game.v1.AnnouncementListRequest
	(*AnnouncementListReply)(nil),   // 3: api.game.v1.AnnouncementListReply
	(*AnnouncementListItem)(nil),    // 4: api.game.v1.AnnouncementListItem
	(*TaskListRequest)(nil),         // 5: api.game.v1.TaskListRequest
	(*TaskListReply)(nil),           // 6: api.game.v1.TaskListReply
	(*TaskListItem)(nil),            // 7: api.game.v1.TaskListItem
	(*TaskRewardRequest)(nil),       // 8: api.game.v1.TaskRewardRequest
	(*TaskRewardReply)(nil),         // 9: api.game.v1.TaskRewardReply
	(*v1.Paginate)(nil),             // 10: api.common.v1.Paginate
}
var file_game_v1_game_proto_depIdxs = []int32{
	10, // 0: api.game.v1.AnnouncementListRequest.page:type_name -> api.common.v1.Paginate
	4,  // 1: api.game.v1.AnnouncementListReply.list:type_name -> api.game.v1.AnnouncementListItem
	10, // 2: api.game.v1.TaskListRequest.page:type_name -> api.common.v1.Paginate
	7,  // 3: api.game.v1.TaskListReply.list:type_name -> api.game.v1.TaskListItem
	0,  // 4: api.game.v1.TaskListItem.status:type_name -> api.game.v1.TaskStatus
	1,  // 5: api.game.v1.TaskListItem.refresh_cycle:type_name -> api.game.v1.RefreshCycle
	2,  // 6: api.game.v1.Game.AnnouncementList:input_type -> api.game.v1.AnnouncementListRequest
	5,  // 7: api.game.v1.Game.TaskList:input_type -> api.game.v1.TaskListRequest
	8,  // 8: api.game.v1.Game.TaskReward:input_type -> api.game.v1.TaskRewardRequest
	3,  // 9: api.game.v1.Game.AnnouncementList:output_type -> api.game.v1.AnnouncementListReply
	6,  // 10: api.game.v1.Game.TaskList:output_type -> api.game.v1.TaskListReply
	9,  // 11: api.game.v1.Game.TaskReward:output_type -> api.game.v1.TaskRewardReply
	9,  // [9:12] is the sub-list for method output_type
	6,  // [6:9] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_game_v1_game_proto_init() }
func file_game_v1_game_proto_init() {
	if File_game_v1_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_v1_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnouncementListRequest); i {
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
		file_game_v1_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnouncementListReply); i {
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
		file_game_v1_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnouncementListItem); i {
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
		file_game_v1_game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskListRequest); i {
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
		file_game_v1_game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskListReply); i {
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
		file_game_v1_game_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskListItem); i {
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
		file_game_v1_game_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskRewardRequest); i {
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
		file_game_v1_game_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskRewardReply); i {
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
			RawDescriptor: file_game_v1_game_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_game_v1_game_proto_goTypes,
		DependencyIndexes: file_game_v1_game_proto_depIdxs,
		EnumInfos:         file_game_v1_game_proto_enumTypes,
		MessageInfos:      file_game_v1_game_proto_msgTypes,
	}.Build()
	File_game_v1_game_proto = out.File
	file_game_v1_game_proto_rawDesc = nil
	file_game_v1_game_proto_goTypes = nil
	file_game_v1_game_proto_depIdxs = nil
}

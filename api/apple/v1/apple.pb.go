// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: apple/v1/apple.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	v1 "github.com/lulu-ls/st-api/api/order/v1"
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

type VerifyReceiptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId      int32  `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`                // app id
	ChannelId  int32  `protobuf:"varint,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`    // channel id
	PurchaseId int32  `protobuf:"varint,3,opt,name=purchase_id,json=purchaseId,proto3" json:"purchase_id,omitempty"` // 道具配置 id
	UserId     int64  `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`             // 用户 id
	OrderId    int64  `protobuf:"varint,6,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`          // order id
	Receipt    string `protobuf:"bytes,7,opt,name=receipt,proto3" json:"receipt,omitempty"`                          //支付凭证
	Env        int32  `protobuf:"varint,8,opt,name=env,proto3" json:"env,omitempty"`                                 // 环境配置0：现网环境（也叫正式环境）1：沙箱环境
}

func (x *VerifyReceiptRequest) Reset() {
	*x = VerifyReceiptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apple_v1_apple_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyReceiptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyReceiptRequest) ProtoMessage() {}

func (x *VerifyReceiptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apple_v1_apple_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyReceiptRequest.ProtoReflect.Descriptor instead.
func (*VerifyReceiptRequest) Descriptor() ([]byte, []int) {
	return file_apple_v1_apple_proto_rawDescGZIP(), []int{0}
}

func (x *VerifyReceiptRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *VerifyReceiptRequest) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *VerifyReceiptRequest) GetPurchaseId() int32 {
	if x != nil {
		return x.PurchaseId
	}
	return 0
}

func (x *VerifyReceiptRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *VerifyReceiptRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *VerifyReceiptRequest) GetReceipt() string {
	if x != nil {
		return x.Receipt
	}
	return ""
}

func (x *VerifyReceiptRequest) GetEnv() int32 {
	if x != nil {
		return x.Env
	}
	return 0
}

type VerifyReceiptReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *VerifyReceiptReply) Reset() {
	*x = VerifyReceiptReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apple_v1_apple_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyReceiptReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyReceiptReply) ProtoMessage() {}

func (x *VerifyReceiptReply) ProtoReflect() protoreflect.Message {
	mi := &file_apple_v1_apple_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyReceiptReply.ProtoReflect.Descriptor instead.
func (*VerifyReceiptReply) Descriptor() ([]byte, []int) {
	return file_apple_v1_apple_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyReceiptReply) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *VerifyReceiptReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PayNoticeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PayNoticeRequest) Reset() {
	*x = PayNoticeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apple_v1_apple_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayNoticeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayNoticeRequest) ProtoMessage() {}

func (x *PayNoticeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apple_v1_apple_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayNoticeRequest.ProtoReflect.Descriptor instead.
func (*PayNoticeRequest) Descriptor() ([]byte, []int) {
	return file_apple_v1_apple_proto_rawDescGZIP(), []int{2}
}

type PayNoticeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PayNoticeReply) Reset() {
	*x = PayNoticeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apple_v1_apple_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayNoticeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayNoticeReply) ProtoMessage() {}

func (x *PayNoticeReply) ProtoReflect() protoreflect.Message {
	mi := &file_apple_v1_apple_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayNoticeReply.ProtoReflect.Descriptor instead.
func (*PayNoticeReply) Descriptor() ([]byte, []int) {
	return file_apple_v1_apple_proto_rawDescGZIP(), []int{3}
}

type ApplePayBuyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId      int32 `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`                // app id
	ChannelId  int32 `protobuf:"varint,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`    // channel id
	PurchaseId int32 `protobuf:"varint,3,opt,name=purchase_id,json=purchaseId,proto3" json:"purchase_id,omitempty"` // 道具配置 id
	UserId     int64 `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`             // 用户 id
	OrderId    int64 `protobuf:"varint,6,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`          // order id
}

func (x *ApplePayBuyRequest) Reset() {
	*x = ApplePayBuyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apple_v1_apple_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplePayBuyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplePayBuyRequest) ProtoMessage() {}

func (x *ApplePayBuyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apple_v1_apple_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplePayBuyRequest.ProtoReflect.Descriptor instead.
func (*ApplePayBuyRequest) Descriptor() ([]byte, []int) {
	return file_apple_v1_apple_proto_rawDescGZIP(), []int{4}
}

func (x *ApplePayBuyRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *ApplePayBuyRequest) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *ApplePayBuyRequest) GetPurchaseId() int32 {
	if x != nil {
		return x.PurchaseId
	}
	return 0
}

func (x *ApplePayBuyRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ApplePayBuyRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type ApplePayBuyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId     int64   `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`              // 订单号
	OrderAmount float64 `protobuf:"fixed64,3,opt,name=order_amount,json=orderAmount,proto3" json:"order_amount,omitempty"` // 订单金额
	Quantity    int32   `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`                           // 购买数量
	Title       string  `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`                                  // 标题
	Env         string  `protobuf:"bytes,6,opt,name=env,proto3" json:"env,omitempty"`                                      // 环境
}

func (x *ApplePayBuyReply) Reset() {
	*x = ApplePayBuyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apple_v1_apple_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplePayBuyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplePayBuyReply) ProtoMessage() {}

func (x *ApplePayBuyReply) ProtoReflect() protoreflect.Message {
	mi := &file_apple_v1_apple_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplePayBuyReply.ProtoReflect.Descriptor instead.
func (*ApplePayBuyReply) Descriptor() ([]byte, []int) {
	return file_apple_v1_apple_proto_rawDescGZIP(), []int{5}
}

func (x *ApplePayBuyReply) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *ApplePayBuyReply) GetOrderAmount() float64 {
	if x != nil {
		return x.OrderAmount
	}
	return 0
}

func (x *ApplePayBuyReply) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *ApplePayBuyReply) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ApplePayBuyReply) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"` // 验证码
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apple_v1_apple_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apple_v1_apple_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_apple_v1_apple_proto_rawDescGZIP(), []int{6}
}

func (x *LoginRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type LoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UniqueId string `protobuf:"bytes,1,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"` // 唯一 id
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`                       // 电子邮箱
}

func (x *LoginReply) Reset() {
	*x = LoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apple_v1_apple_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply) ProtoMessage() {}

func (x *LoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_apple_v1_apple_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReply.ProtoReflect.Descriptor instead.
func (*LoginReply) Descriptor() ([]byte, []int) {
	return file_apple_v1_apple_proto_rawDescGZIP(), []int{7}
}

func (x *LoginReply) GetUniqueId() string {
	if x != nil {
		return x.UniqueId
	}
	return ""
}

func (x *LoginReply) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_apple_v1_apple_proto protoreflect.FileDescriptor

var file_apple_v1_apple_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x14, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x65, 0x6e,
	0x76, 0x22, 0x42, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x50, 0x61, 0x79, 0x4e, 0x6f, 0x74, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x50, 0x61, 0x79,
	0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x9f, 0x01, 0x0a, 0x12,
	0x41, 0x70, 0x70, 0x6c, 0x65, 0x50, 0x61, 0x79, 0x42, 0x75, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x75, 0x72, 0x63,
	0x68, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70,
	0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x94, 0x01,
	0x0a, 0x10, 0x41, 0x70, 0x70, 0x6c, 0x65, 0x50, 0x61, 0x79, 0x42, 0x75, 0x79, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x65, 0x6e, 0x76, 0x22, 0x2b, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0x3f, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x32, 0xac, 0x03, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x6c, 0x65, 0x12, 0x83, 0x01, 0x0a,
	0x0d, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x22,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x3a, 0x01, 0x2a, 0x22,
	0x21, 0x2f, 0x73, 0x74, 0x2d, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70,
	0x70, 0x6c, 0x65, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x12, 0x73, 0x0a, 0x09, 0x50, 0x61, 0x79, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x12,
	0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x61, 0x79, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x61, 0x79, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x28, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f, 0x73, 0x74, 0x2d, 0x67, 0x61,
	0x6d, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x61, 0x79,
	0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x03, 0x50, 0x61, 0x79, 0x12, 0x20,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70,
	0x70, 0x6c, 0x65, 0x50, 0x61, 0x79, 0x42, 0x75, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x61, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x73, 0x74, 0x2d,
	0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x2f, 0x70,
	0x61, 0x79, 0x12, 0x3d, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x42, 0x3b, 0x0a, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x75, 0x6c, 0x75, 0x2d, 0x6c, 0x73, 0x2f, 0x73, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apple_v1_apple_proto_rawDescOnce sync.Once
	file_apple_v1_apple_proto_rawDescData = file_apple_v1_apple_proto_rawDesc
)

func file_apple_v1_apple_proto_rawDescGZIP() []byte {
	file_apple_v1_apple_proto_rawDescOnce.Do(func() {
		file_apple_v1_apple_proto_rawDescData = protoimpl.X.CompressGZIP(file_apple_v1_apple_proto_rawDescData)
	})
	return file_apple_v1_apple_proto_rawDescData
}

var file_apple_v1_apple_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_apple_v1_apple_proto_goTypes = []interface{}{
	(*VerifyReceiptRequest)(nil), // 0: api.apple.v1.VerifyReceiptRequest
	(*VerifyReceiptReply)(nil),   // 1: api.apple.v1.VerifyReceiptReply
	(*PayNoticeRequest)(nil),     // 2: api.apple.v1.PayNoticeRequest
	(*PayNoticeReply)(nil),       // 3: api.apple.v1.PayNoticeReply
	(*ApplePayBuyRequest)(nil),   // 4: api.apple.v1.ApplePayBuyRequest
	(*ApplePayBuyReply)(nil),     // 5: api.apple.v1.ApplePayBuyReply
	(*LoginRequest)(nil),         // 6: api.apple.v1.LoginRequest
	(*LoginReply)(nil),           // 7: api.apple.v1.LoginReply
	(*v1.ConfirmPayReply)(nil),   // 8: api.order.v1.ConfirmPayReply
}
var file_apple_v1_apple_proto_depIdxs = []int32{
	0, // 0: api.apple.v1.Apple.VerifyReceipt:input_type -> api.apple.v1.VerifyReceiptRequest
	2, // 1: api.apple.v1.Apple.PayNotice:input_type -> api.apple.v1.PayNoticeRequest
	4, // 2: api.apple.v1.Apple.Pay:input_type -> api.apple.v1.ApplePayBuyRequest
	6, // 3: api.apple.v1.Apple.Login:input_type -> api.apple.v1.LoginRequest
	1, // 4: api.apple.v1.Apple.VerifyReceipt:output_type -> api.apple.v1.VerifyReceiptReply
	3, // 5: api.apple.v1.Apple.PayNotice:output_type -> api.apple.v1.PayNoticeReply
	8, // 6: api.apple.v1.Apple.Pay:output_type -> api.order.v1.ConfirmPayReply
	7, // 7: api.apple.v1.Apple.Login:output_type -> api.apple.v1.LoginReply
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apple_v1_apple_proto_init() }
func file_apple_v1_apple_proto_init() {
	if File_apple_v1_apple_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apple_v1_apple_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyReceiptRequest); i {
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
		file_apple_v1_apple_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyReceiptReply); i {
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
		file_apple_v1_apple_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayNoticeRequest); i {
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
		file_apple_v1_apple_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayNoticeReply); i {
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
		file_apple_v1_apple_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplePayBuyRequest); i {
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
		file_apple_v1_apple_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplePayBuyReply); i {
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
		file_apple_v1_apple_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_apple_v1_apple_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReply); i {
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
			RawDescriptor: file_apple_v1_apple_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apple_v1_apple_proto_goTypes,
		DependencyIndexes: file_apple_v1_apple_proto_depIdxs,
		MessageInfos:      file_apple_v1_apple_proto_msgTypes,
	}.Build()
	File_apple_v1_apple_proto = out.File
	file_apple_v1_apple_proto_rawDesc = nil
	file_apple_v1_apple_proto_goTypes = nil
	file_apple_v1_apple_proto_depIdxs = nil
}

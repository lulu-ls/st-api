// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: ticket/v1/ticket.proto

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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"` // 微信 code 码， 可获取手机号
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_v1_ticket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_v1_ticket_proto_msgTypes[0]
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
	return file_ticket_v1_ticket_proto_rawDescGZIP(), []int{0}
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

	AccessToken  string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`    // token
	MerchantId   int32    `protobuf:"varint,2,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`      // 商户 id
	MerchantName string   `protobuf:"bytes,3,opt,name=merchant_name,json=merchantName,proto3" json:"merchant_name,omitempty"` // 商户名称
	StoreId      int32    `protobuf:"varint,4,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`               // 门店 id
	StoreName    string   `protobuf:"bytes,5,opt,name=store_name,json=storeName,proto3" json:"store_name,omitempty"`          // 门店名称
	HasStores    bool     `protobuf:"varint,6,opt,name=has_stores,json=hasStores,proto3" json:"has_stores,omitempty"`         // 是否有多个门店
	Stores       []*Store `protobuf:"bytes,7,rep,name=stores,proto3" json:"stores,omitempty"`                                 // 门店列表
}

func (x *LoginReply) Reset() {
	*x = LoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_v1_ticket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply) ProtoMessage() {}

func (x *LoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_v1_ticket_proto_msgTypes[1]
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
	return file_ticket_v1_ticket_proto_rawDescGZIP(), []int{1}
}

func (x *LoginReply) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginReply) GetMerchantId() int32 {
	if x != nil {
		return x.MerchantId
	}
	return 0
}

func (x *LoginReply) GetMerchantName() string {
	if x != nil {
		return x.MerchantName
	}
	return ""
}

func (x *LoginReply) GetStoreId() int32 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *LoginReply) GetStoreName() string {
	if x != nil {
		return x.StoreName
	}
	return ""
}

func (x *LoginReply) GetHasStores() bool {
	if x != nil {
		return x.HasStores
	}
	return false
}

func (x *LoginReply) GetStores() []*Store {
	if x != nil {
		return x.Stores
	}
	return nil
}

type Store struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreId   int32  `protobuf:"varint,4,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`      // 门店 id
	StoreName string `protobuf:"bytes,5,opt,name=store_name,json=storeName,proto3" json:"store_name,omitempty"` // 门店名称
}

func (x *Store) Reset() {
	*x = Store{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_v1_ticket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Store) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Store) ProtoMessage() {}

func (x *Store) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_v1_ticket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Store.ProtoReflect.Descriptor instead.
func (*Store) Descriptor() ([]byte, []int) {
	return file_ticket_v1_ticket_proto_rawDescGZIP(), []int{2}
}

func (x *Store) GetStoreId() int32 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *Store) GetStoreName() string {
	if x != nil {
		return x.StoreName
	}
	return ""
}

type StatisticsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmployeeId int64 `protobuf:"varint,1,opt,name=employee_id,json=employeeId,proto3" json:"employee_id,omitempty"` // 用户 id
	StoreId    int32 `protobuf:"varint,4,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`          // 门店 id
}

func (x *StatisticsRequest) Reset() {
	*x = StatisticsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_v1_ticket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatisticsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatisticsRequest) ProtoMessage() {}

func (x *StatisticsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_v1_ticket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatisticsRequest.ProtoReflect.Descriptor instead.
func (*StatisticsRequest) Descriptor() ([]byte, []int) {
	return file_ticket_v1_ticket_proto_rawDescGZIP(), []int{3}
}

func (x *StatisticsRequest) GetEmployeeId() int64 {
	if x != nil {
		return x.EmployeeId
	}
	return 0
}

func (x *StatisticsRequest) GetStoreId() int32 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

type StatisticsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quantity int32   `protobuf:"varint,1,opt,name=quantity,proto3" json:"quantity,omitempty"` // 核销数量
	Amount   float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`    // 核销价值
}

func (x *StatisticsReply) Reset() {
	*x = StatisticsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_v1_ticket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatisticsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatisticsReply) ProtoMessage() {}

func (x *StatisticsReply) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_v1_ticket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatisticsReply.ProtoReflect.Descriptor instead.
func (*StatisticsReply) Descriptor() ([]byte, []int) {
	return file_ticket_v1_ticket_proto_rawDescGZIP(), []int{4}
}

func (x *StatisticsReply) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *StatisticsReply) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_ticket_v1_ticket_proto protoreflect.FileDescriptor

var file_ticket_v1_ticket_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b,
	0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xfc, 0x01, 0x0a, 0x0a,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x0a,
	0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x68, 0x61, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x68, 0x61, 0x73, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x06,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x52, 0x06, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x22, 0x41, 0x0a, 0x05, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4f, 0x0a,
	0x11, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x22, 0x45,
	0x0a, 0x0f, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xe8, 0x01, 0x0a, 0x06, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x64, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x73,
	0x74, 0x2d, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x78, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01,
	0x2a, 0x22, 0x1d, 0x2f, 0x73, 0x74, 0x2d, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x42, 0x3d, 0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x75, 0x6c, 0x75, 0x2d, 0x6c, 0x73, 0x2f, 0x73, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ticket_v1_ticket_proto_rawDescOnce sync.Once
	file_ticket_v1_ticket_proto_rawDescData = file_ticket_v1_ticket_proto_rawDesc
)

func file_ticket_v1_ticket_proto_rawDescGZIP() []byte {
	file_ticket_v1_ticket_proto_rawDescOnce.Do(func() {
		file_ticket_v1_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_ticket_v1_ticket_proto_rawDescData)
	})
	return file_ticket_v1_ticket_proto_rawDescData
}

var file_ticket_v1_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ticket_v1_ticket_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),      // 0: api.ticket.v1.LoginRequest
	(*LoginReply)(nil),        // 1: api.ticket.v1.LoginReply
	(*Store)(nil),             // 2: api.ticket.v1.Store
	(*StatisticsRequest)(nil), // 3: api.ticket.v1.StatisticsRequest
	(*StatisticsReply)(nil),   // 4: api.ticket.v1.StatisticsReply
}
var file_ticket_v1_ticket_proto_depIdxs = []int32{
	2, // 0: api.ticket.v1.LoginReply.stores:type_name -> api.ticket.v1.Store
	0, // 1: api.ticket.v1.Ticket.Login:input_type -> api.ticket.v1.LoginRequest
	3, // 2: api.ticket.v1.Ticket.Statistics:input_type -> api.ticket.v1.StatisticsRequest
	1, // 3: api.ticket.v1.Ticket.Login:output_type -> api.ticket.v1.LoginReply
	4, // 4: api.ticket.v1.Ticket.Statistics:output_type -> api.ticket.v1.StatisticsReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ticket_v1_ticket_proto_init() }
func file_ticket_v1_ticket_proto_init() {
	if File_ticket_v1_ticket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ticket_v1_ticket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_ticket_v1_ticket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_ticket_v1_ticket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Store); i {
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
		file_ticket_v1_ticket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatisticsRequest); i {
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
		file_ticket_v1_ticket_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatisticsReply); i {
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
			RawDescriptor: file_ticket_v1_ticket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ticket_v1_ticket_proto_goTypes,
		DependencyIndexes: file_ticket_v1_ticket_proto_depIdxs,
		MessageInfos:      file_ticket_v1_ticket_proto_msgTypes,
	}.Build()
	File_ticket_v1_ticket_proto = out.File
	file_ticket_v1_ticket_proto_rawDesc = nil
	file_ticket_v1_ticket_proto_goTypes = nil
	file_ticket_v1_ticket_proto_depIdxs = nil
}
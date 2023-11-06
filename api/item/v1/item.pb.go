// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: item/v1/item.proto

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

type ListItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  *v1.Paginate `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	AppId int32        `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"` //  [(validate.rules).int32 = {gt: 0}]
}

func (x *ListItemRequest) Reset() {
	*x = ListItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_v1_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItemRequest) ProtoMessage() {}

func (x *ListItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_item_v1_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItemRequest.ProtoReflect.Descriptor instead.
func (*ListItemRequest) Descriptor() ([]byte, []int) {
	return file_item_v1_item_proto_rawDescGZIP(), []int{0}
}

func (x *ListItemRequest) GetPage() *v1.Paginate {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListItemRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

type ListItemReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*ListItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Total int32       `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListItemReply) Reset() {
	*x = ListItemReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_v1_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItemReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItemReply) ProtoMessage() {}

func (x *ListItemReply) ProtoReflect() protoreflect.Message {
	mi := &file_item_v1_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItemReply.ProtoReflect.Descriptor instead.
func (*ListItemReply) Descriptor() ([]byte, []int) {
	return file_item_v1_item_proto_rawDescGZIP(), []int{1}
}

func (x *ListItemReply) GetList() []*ListItem {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *ListItemReply) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type ListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId         int32  `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	ItemCategoryId int32  `protobuf:"varint,2,opt,name=item_category_id,json=itemCategoryId,proto3" json:"item_category_id,omitempty"`
	ItemFuncId     int32  `protobuf:"varint,3,opt,name=item_func_id,json=itemFuncId,proto3" json:"item_func_id,omitempty"`
	Name           string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Title          string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Image          string `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
	MaxStack       int32  `protobuf:"varint,7,opt,name=max_stack,json=maxStack,proto3" json:"max_stack,omitempty"` //	int32 loot_pool_id = 8;
}

func (x *ListItem) Reset() {
	*x = ListItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_v1_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItem) ProtoMessage() {}

func (x *ListItem) ProtoReflect() protoreflect.Message {
	mi := &file_item_v1_item_proto_msgTypes[2]
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
	return file_item_v1_item_proto_rawDescGZIP(), []int{2}
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

var File_item_v1_item_proto protoreflect.FileDescriptor

var file_item_v1_item_proto_rawDesc = []byte{
	0x0a, 0x12, 0x69, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0x50,
	0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x29, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0xcc, 0x01, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a,
	0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x69, 0x74, 0x65, 0x6d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0c, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x46, 0x75, 0x6e, 0x63,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x32,
	0x6f, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x67, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x73, 0x74, 0x2d, 0x67, 0x61,
	0x6d, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x42, 0x39, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x50,
	0x01, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x75,
	0x6c, 0x75, 0x2d, 0x6c, 0x73, 0x2f, 0x73, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_item_v1_item_proto_rawDescOnce sync.Once
	file_item_v1_item_proto_rawDescData = file_item_v1_item_proto_rawDesc
)

func file_item_v1_item_proto_rawDescGZIP() []byte {
	file_item_v1_item_proto_rawDescOnce.Do(func() {
		file_item_v1_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_item_v1_item_proto_rawDescData)
	})
	return file_item_v1_item_proto_rawDescData
}

var file_item_v1_item_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_item_v1_item_proto_goTypes = []interface{}{
	(*ListItemRequest)(nil), // 0: api.item.v1.ListItemRequest
	(*ListItemReply)(nil),   // 1: api.item.v1.ListItemReply
	(*ListItem)(nil),        // 2: api.item.v1.ListItem
	(*v1.Paginate)(nil),     // 3: api.common.v1.Paginate
}
var file_item_v1_item_proto_depIdxs = []int32{
	3, // 0: api.item.v1.ListItemRequest.page:type_name -> api.common.v1.Paginate
	2, // 1: api.item.v1.ListItemReply.list:type_name -> api.item.v1.ListItem
	0, // 2: api.item.v1.Item.ListItem:input_type -> api.item.v1.ListItemRequest
	1, // 3: api.item.v1.Item.ListItem:output_type -> api.item.v1.ListItemReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_item_v1_item_proto_init() }
func file_item_v1_item_proto_init() {
	if File_item_v1_item_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_item_v1_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListItemRequest); i {
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
		file_item_v1_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListItemReply); i {
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
		file_item_v1_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_item_v1_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_item_v1_item_proto_goTypes,
		DependencyIndexes: file_item_v1_item_proto_depIdxs,
		MessageInfos:      file_item_v1_item_proto_msgTypes,
	}.Build()
	File_item_v1_item_proto = out.File
	file_item_v1_item_proto_rawDesc = nil
	file_item_v1_item_proto_goTypes = nil
	file_item_v1_item_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: common.proto

package commonpb

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

type SendContentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room   string `protobuf:"bytes,1,opt,name=Room,proto3" json:"Room,omitempty"`
	Object string `protobuf:"bytes,2,opt,name=Object,proto3" json:"Object,omitempty"`
	Key    string `protobuf:"bytes,3,opt,name=Key,proto3" json:"Key,omitempty"`
	Value  string `protobuf:"bytes,4,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *SendContentRequest) Reset() {
	*x = SendContentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendContentRequest) ProtoMessage() {}

func (x *SendContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendContentRequest.ProtoReflect.Descriptor instead.
func (*SendContentRequest) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *SendContentRequest) GetRoom() string {
	if x != nil {
		return x.Room
	}
	return ""
}

func (x *SendContentRequest) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *SendContentRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SendContentRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SendContentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendContentResponse) Reset() {
	*x = SendContentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendContentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendContentResponse) ProtoMessage() {}

func (x *SendContentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendContentResponse.ProtoReflect.Descriptor instead.
func (*SendContentResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *SendContentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ReceiveContentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ReceiveContentRequest) Reset() {
	*x = ReceiveContentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveContentRequest) ProtoMessage() {}

func (x *ReceiveContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveContentRequest.ProtoReflect.Descriptor instead.
func (*ReceiveContentRequest) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *ReceiveContentRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ReceiveContentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room   string `protobuf:"bytes,1,opt,name=Room,proto3" json:"Room,omitempty"`
	Object string `protobuf:"bytes,2,opt,name=Object,proto3" json:"Object,omitempty"`
	Key    string `protobuf:"bytes,3,opt,name=Key,proto3" json:"Key,omitempty"`
	Value  string `protobuf:"bytes,4,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *ReceiveContentResponse) Reset() {
	*x = ReceiveContentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveContentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveContentResponse) ProtoMessage() {}

func (x *ReceiveContentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveContentResponse.ProtoReflect.Descriptor instead.
func (*ReceiveContentResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *ReceiveContentResponse) GetRoom() string {
	if x != nil {
		return x.Room
	}
	return ""
}

func (x *ReceiveContentResponse) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *ReceiveContentResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ReceiveContentResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ResponsePostFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResponsePostFile) Reset() {
	*x = ResponsePostFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponsePostFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponsePostFile) ProtoMessage() {}

func (x *ResponsePostFile) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponsePostFile.ProtoReflect.Descriptor instead.
func (*ResponsePostFile) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4}
}

func (x *ResponsePostFile) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RequestPostFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*RequestPostFile_Data
	//	*RequestPostFile_Name
	Value isRequestPostFile_Value `protobuf_oneof:"value"`
}

func (x *RequestPostFile) Reset() {
	*x = RequestPostFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestPostFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestPostFile) ProtoMessage() {}

func (x *RequestPostFile) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestPostFile.ProtoReflect.Descriptor instead.
func (*RequestPostFile) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{5}
}

func (m *RequestPostFile) GetValue() isRequestPostFile_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *RequestPostFile) GetData() []byte {
	if x, ok := x.GetValue().(*RequestPostFile_Data); ok {
		return x.Data
	}
	return nil
}

func (x *RequestPostFile) GetName() string {
	if x, ok := x.GetValue().(*RequestPostFile_Name); ok {
		return x.Name
	}
	return ""
}

type isRequestPostFile_Value interface {
	isRequestPostFile_Value()
}

type RequestPostFile_Data struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3,oneof"`
}

type RequestPostFile_Name struct {
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*RequestPostFile_Data) isRequestPostFile_Value() {}

func (*RequestPostFile_Name) isRequestPostFile_Value() {}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x22, 0x68, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x6f,
	0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x2f, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x31, 0x0a, 0x15, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x6c, 0x0a, 0x16, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x2c, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x46, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x65, 0x0a, 0x12, 0x53, 0x65,
	0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4f, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52,
	0x50, 0x43, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x32, 0x71, 0x0a, 0x15, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x11, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x50, 0x43, 0x12,
	0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x32, 0x54, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x48, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x50, 0x43, 0x12,
	0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x1a, 0x1a, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x75, 0x6a, 0x69, 0x59, 0x61, 0x62,
	0x65, 0x2f, 0x67, 0x6f, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_common_proto_goTypes = []interface{}{
	(*SendContentRequest)(nil),     // 0: commonpb.SendContentRequest
	(*SendContentResponse)(nil),    // 1: commonpb.SendContentResponse
	(*ReceiveContentRequest)(nil),  // 2: commonpb.ReceiveContentRequest
	(*ReceiveContentResponse)(nil), // 3: commonpb.ReceiveContentResponse
	(*ResponsePostFile)(nil),       // 4: commonpb.ResponsePostFile
	(*RequestPostFile)(nil),        // 5: commonpb.RequestPostFile
}
var file_common_proto_depIdxs = []int32{
	0, // 0: commonpb.SendContentService.SendContentRPC:input_type -> commonpb.SendContentRequest
	2, // 1: commonpb.ReceiveContentService.ReceiveContentRPC:input_type -> commonpb.ReceiveContentRequest
	5, // 2: commonpb.PostFile.PostFileRPC:input_type -> commonpb.RequestPostFile
	1, // 3: commonpb.SendContentService.SendContentRPC:output_type -> commonpb.SendContentResponse
	3, // 4: commonpb.ReceiveContentService.ReceiveContentRPC:output_type -> commonpb.ReceiveContentResponse
	4, // 5: commonpb.PostFile.PostFileRPC:output_type -> commonpb.ResponsePostFile
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Isporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendContentRequest); i {
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
		file_common_proto_msgTypes[1].Isporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendContentResponse); i {
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
		file_common_proto_msgTypes[2].Isporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveContentRequest); i {
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
		file_common_proto_msgTypes[3].Isporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveContentResponse); i {
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
		file_common_proto_msgTypes[4].Isporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponsePostFile); i {
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
		file_common_proto_msgTypes[5].Isporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestPostFile); i {
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
	file_common_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*RequestPostFile_Data)(nil),
		(*RequestPostFile_Name)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}

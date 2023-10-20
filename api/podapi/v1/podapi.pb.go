// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: podapi/v1/podapi.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	v1 "github.com/zbnzl/paas-pod/api/pod/v1"
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

type Pair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Pair) Reset() {
	*x = Pair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_podapi_v1_podapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pair) ProtoMessage() {}

func (x *Pair) ProtoReflect() protoreflect.Message {
	mi := &file_podapi_v1_podapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pair.ProtoReflect.Descriptor instead.
func (*Pair) Descriptor() ([]byte, []int) {
	return file_podapi_v1_podapi_proto_rawDescGZIP(), []int{0}
}

func (x *Pair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Pair) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method string           `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Path   string           `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Header map[string]*Pair `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Get    map[string]*Pair `protobuf:"bytes,4,rep,name=get,proto3" json:"get,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Post   map[string]*Pair `protobuf:"bytes,5,rep,name=post,proto3" json:"post,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body   string           `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
	Url    string           `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_podapi_v1_podapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_podapi_v1_podapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_podapi_v1_podapi_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Request) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Request) GetHeader() map[string]*Pair {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Request) GetGet() map[string]*Pair {
	if x != nil {
		return x.Get
	}
	return nil
}

func (x *Request) GetPost() map[string]*Pair {
	if x != nil {
		return x.Post
	}
	return nil
}

func (x *Request) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Request) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32            `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Header     map[string]*Pair `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body       string           `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_podapi_v1_podapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_podapi_v1_podapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_podapi_v1_podapi_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *Response) GetHeader() map[string]*Pair {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Response) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type DeletePodApiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePodApiRequest) Reset() {
	*x = DeletePodApiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_podapi_v1_podapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePodApiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePodApiRequest) ProtoMessage() {}

func (x *DeletePodApiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_podapi_v1_podapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePodApiRequest.ProtoReflect.Descriptor instead.
func (*DeletePodApiRequest) Descriptor() ([]byte, []int) {
	return file_podapi_v1_podapi_proto_rawDescGZIP(), []int{3}
}

type DeletePodApiReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePodApiReply) Reset() {
	*x = DeletePodApiReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_podapi_v1_podapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePodApiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePodApiReply) ProtoMessage() {}

func (x *DeletePodApiReply) ProtoReflect() protoreflect.Message {
	mi := &file_podapi_v1_podapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePodApiReply.ProtoReflect.Descriptor instead.
func (*DeletePodApiReply) Descriptor() ([]byte, []int) {
	return file_podapi_v1_podapi_proto_rawDescGZIP(), []int{4}
}

type GetPodApiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPodApiRequest) Reset() {
	*x = GetPodApiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_podapi_v1_podapi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPodApiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPodApiRequest) ProtoMessage() {}

func (x *GetPodApiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_podapi_v1_podapi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPodApiRequest.ProtoReflect.Descriptor instead.
func (*GetPodApiRequest) Descriptor() ([]byte, []int) {
	return file_podapi_v1_podapi_proto_rawDescGZIP(), []int{5}
}

type GetPodApiReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPodApiReply) Reset() {
	*x = GetPodApiReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_podapi_v1_podapi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPodApiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPodApiReply) ProtoMessage() {}

func (x *GetPodApiReply) ProtoReflect() protoreflect.Message {
	mi := &file_podapi_v1_podapi_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPodApiReply.ProtoReflect.Descriptor instead.
func (*GetPodApiReply) Descriptor() ([]byte, []int) {
	return file_podapi_v1_podapi_proto_rawDescGZIP(), []int{6}
}

type ListPodApiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPodApiRequest) Reset() {
	*x = ListPodApiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_podapi_v1_podapi_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPodApiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPodApiRequest) ProtoMessage() {}

func (x *ListPodApiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_podapi_v1_podapi_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPodApiRequest.ProtoReflect.Descriptor instead.
func (*ListPodApiRequest) Descriptor() ([]byte, []int) {
	return file_podapi_v1_podapi_proto_rawDescGZIP(), []int{7}
}

type ListPodApiReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPodApiReply) Reset() {
	*x = ListPodApiReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_podapi_v1_podapi_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPodApiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPodApiReply) ProtoMessage() {}

func (x *ListPodApiReply) ProtoReflect() protoreflect.Message {
	mi := &file_podapi_v1_podapi_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPodApiReply.ProtoReflect.Descriptor instead.
func (*ListPodApiReply) Descriptor() ([]byte, []int) {
	return file_podapi_v1_podapi_proto_rawDescGZIP(), []int{8}
}

var File_podapi_v1_podapi_proto protoreflect.FileDescriptor

var file_podapi_v1_podapi_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6f, 0x64, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x64, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f,
	0x64, 0x41, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10,
	0x70, 0x6f, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x30, 0x0a, 0x04, 0x50, 0x61, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x22, 0xf4, 0x03, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x3a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x64, 0x41, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x31, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x6f, 0x64, 0x41, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x67,
	0x65, 0x74, 0x12, 0x34, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x64, 0x41, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x1a, 0x4e,
	0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x64, 0x41, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x61, 0x69, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4b,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x6f, 0x64, 0x41, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x69, 0x72,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4c, 0x0a, 0x09, 0x50,
	0x6f, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x6f, 0x64, 0x41, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xcb, 0x01, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x64,
	0x41, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x1a, 0x4e, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f,
	0x64, 0x41, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x6f, 0x64, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x41, 0x70, 0x69, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x64, 0x41, 0x70, 0x69,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x64, 0x41, 0x70, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x6f, 0x64, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11,
	0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x64, 0x41, 0x70, 0x69, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x32, 0xf6, 0x02, 0x0a, 0x06, 0x50, 0x6f, 0x64, 0x41, 0x70, 0x69, 0x12, 0x46, 0x0a, 0x0b,
	0x46, 0x69, 0x6e, 0x64, 0x50, 0x6f, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x11, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x64, 0x49, 0x64, 0x1a, 0x13,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x6f, 0x64, 0x12, 0x47, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x64, 0x12, 0x13,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0c, 0x3a, 0x01, 0x2a, 0x22, 0x07, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x64, 0x12, 0x49, 0x0a,
	0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x11,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x64, 0x49,
	0x64, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x2a,
	0x07, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x64, 0x12, 0x4a, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x64, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x3a, 0x01, 0x2a, 0x1a, 0x07, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x6f, 0x64, 0x12, 0x44, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x6f, 0x64, 0x41, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x64, 0x41, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x05, 0x12, 0x03, 0x2f, 0x76, 0x31, 0x42, 0x3d, 0x0a, 0x0d, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x6f, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x2a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x62, 0x6e, 0x7a, 0x6c, 0x2f,
	0x70, 0x61, 0x61, 0x73, 0x2d, 0x70, 0x6f, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f, 0x64,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_podapi_v1_podapi_proto_rawDescOnce sync.Once
	file_podapi_v1_podapi_proto_rawDescData = file_podapi_v1_podapi_proto_rawDesc
)

func file_podapi_v1_podapi_proto_rawDescGZIP() []byte {
	file_podapi_v1_podapi_proto_rawDescOnce.Do(func() {
		file_podapi_v1_podapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_podapi_v1_podapi_proto_rawDescData)
	})
	return file_podapi_v1_podapi_proto_rawDescData
}

var file_podapi_v1_podapi_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_podapi_v1_podapi_proto_goTypes = []interface{}{
	(*Pair)(nil),                // 0: api.podApi.v1.Pair
	(*Request)(nil),             // 1: api.podApi.v1.Request
	(*Response)(nil),            // 2: api.podApi.v1.Response
	(*DeletePodApiRequest)(nil), // 3: api.podApi.v1.DeletePodApiRequest
	(*DeletePodApiReply)(nil),   // 4: api.podApi.v1.DeletePodApiReply
	(*GetPodApiRequest)(nil),    // 5: api.podApi.v1.GetPodApiRequest
	(*GetPodApiReply)(nil),      // 6: api.podApi.v1.GetPodApiReply
	(*ListPodApiRequest)(nil),   // 7: api.podApi.v1.ListPodApiRequest
	(*ListPodApiReply)(nil),     // 8: api.podApi.v1.ListPodApiReply
	nil,                         // 9: api.podApi.v1.Request.HeaderEntry
	nil,                         // 10: api.podApi.v1.Request.GetEntry
	nil,                         // 11: api.podApi.v1.Request.PostEntry
	nil,                         // 12: api.podApi.v1.Response.HeaderEntry
	(*v1.PodId)(nil),            // 13: api.pod.v1.PodId
	(*v1.PodInfo)(nil),          // 14: api.pod.v1.PodInfo
	(*v1.Response)(nil),         // 15: api.pod.v1.Response
}
var file_podapi_v1_podapi_proto_depIdxs = []int32{
	9,  // 0: api.podApi.v1.Request.header:type_name -> api.podApi.v1.Request.HeaderEntry
	10, // 1: api.podApi.v1.Request.get:type_name -> api.podApi.v1.Request.GetEntry
	11, // 2: api.podApi.v1.Request.post:type_name -> api.podApi.v1.Request.PostEntry
	12, // 3: api.podApi.v1.Response.header:type_name -> api.podApi.v1.Response.HeaderEntry
	0,  // 4: api.podApi.v1.Request.HeaderEntry.value:type_name -> api.podApi.v1.Pair
	0,  // 5: api.podApi.v1.Request.GetEntry.value:type_name -> api.podApi.v1.Pair
	0,  // 6: api.podApi.v1.Request.PostEntry.value:type_name -> api.podApi.v1.Pair
	0,  // 7: api.podApi.v1.Response.HeaderEntry.value:type_name -> api.podApi.v1.Pair
	13, // 8: api.podApi.v1.PodApi.FindPodById:input_type -> api.pod.v1.PodId
	14, // 9: api.podApi.v1.PodApi.AddPod:input_type -> api.pod.v1.PodInfo
	13, // 10: api.podApi.v1.PodApi.DeletePodById:input_type -> api.pod.v1.PodId
	14, // 11: api.podApi.v1.PodApi.UpdatePod:input_type -> api.pod.v1.PodInfo
	1,  // 12: api.podApi.v1.PodApi.Call:input_type -> api.podApi.v1.Request
	14, // 13: api.podApi.v1.PodApi.FindPodById:output_type -> api.pod.v1.PodInfo
	15, // 14: api.podApi.v1.PodApi.AddPod:output_type -> api.pod.v1.Response
	15, // 15: api.podApi.v1.PodApi.DeletePodById:output_type -> api.pod.v1.Response
	15, // 16: api.podApi.v1.PodApi.UpdatePod:output_type -> api.pod.v1.Response
	2,  // 17: api.podApi.v1.PodApi.Call:output_type -> api.podApi.v1.Response
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_podapi_v1_podapi_proto_init() }
func file_podapi_v1_podapi_proto_init() {
	if File_podapi_v1_podapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_podapi_v1_podapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pair); i {
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
		file_podapi_v1_podapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_podapi_v1_podapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_podapi_v1_podapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePodApiRequest); i {
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
		file_podapi_v1_podapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePodApiReply); i {
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
		file_podapi_v1_podapi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPodApiRequest); i {
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
		file_podapi_v1_podapi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPodApiReply); i {
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
		file_podapi_v1_podapi_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPodApiRequest); i {
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
		file_podapi_v1_podapi_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPodApiReply); i {
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
			RawDescriptor: file_podapi_v1_podapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_podapi_v1_podapi_proto_goTypes,
		DependencyIndexes: file_podapi_v1_podapi_proto_depIdxs,
		MessageInfos:      file_podapi_v1_podapi_proto_msgTypes,
	}.Build()
	File_podapi_v1_podapi_proto = out.File
	file_podapi_v1_podapi_proto_rawDesc = nil
	file_podapi_v1_podapi_proto_goTypes = nil
	file_podapi_v1_podapi_proto_depIdxs = nil
}

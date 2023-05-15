// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: screen_recorder.proto

package ypb

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

type ScreenRecorder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Filename  string `protobuf:"bytes,2,opt,name=Filename,proto3" json:"Filename,omitempty"`
	NoteInfo  string `protobuf:"bytes,3,opt,name=NoteInfo,proto3" json:"NoteInfo,omitempty"`
	Project   string `protobuf:"bytes,4,opt,name=Project,proto3" json:"Project,omitempty"`
	CreatedAt int64  `protobuf:"varint,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt int64  `protobuf:"varint,6,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *ScreenRecorder) Reset() {
	*x = ScreenRecorder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_screen_recorder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScreenRecorder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScreenRecorder) ProtoMessage() {}

func (x *ScreenRecorder) ProtoReflect() protoreflect.Message {
	mi := &file_screen_recorder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScreenRecorder.ProtoReflect.Descriptor instead.
func (*ScreenRecorder) Descriptor() ([]byte, []int) {
	return file_screen_recorder_proto_rawDescGZIP(), []int{0}
}

func (x *ScreenRecorder) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ScreenRecorder) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *ScreenRecorder) GetNoteInfo() string {
	if x != nil {
		return x.NoteInfo
	}
	return ""
}

func (x *ScreenRecorder) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *ScreenRecorder) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ScreenRecorder) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type QueryScreenRecorderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       []*ScreenRecorder `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
	Pagination *Paging           `protobuf:"bytes,2,opt,name=Pagination,proto3" json:"Pagination,omitempty"`
	Total      int64             `protobuf:"varint,3,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *QueryScreenRecorderResponse) Reset() {
	*x = QueryScreenRecorderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_screen_recorder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryScreenRecorderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryScreenRecorderResponse) ProtoMessage() {}

func (x *QueryScreenRecorderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_screen_recorder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryScreenRecorderResponse.ProtoReflect.Descriptor instead.
func (*QueryScreenRecorderResponse) Descriptor() ([]byte, []int) {
	return file_screen_recorder_proto_rawDescGZIP(), []int{1}
}

func (x *QueryScreenRecorderResponse) GetData() []*ScreenRecorder {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryScreenRecorderResponse) GetPagination() *Paging {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *QueryScreenRecorderResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type IsScrecorderReadyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok     bool   `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=Reason,proto3" json:"Reason,omitempty"`
}

func (x *IsScrecorderReadyResponse) Reset() {
	*x = IsScrecorderReadyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_screen_recorder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsScrecorderReadyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsScrecorderReadyResponse) ProtoMessage() {}

func (x *IsScrecorderReadyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_screen_recorder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsScrecorderReadyResponse.ProtoReflect.Descriptor instead.
func (*IsScrecorderReadyResponse) Descriptor() ([]byte, []int) {
	return file_screen_recorder_proto_rawDescGZIP(), []int{2}
}

func (x *IsScrecorderReadyResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *IsScrecorderReadyResponse) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type QueryScreenRecorderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project    string  `protobuf:"bytes,1,opt,name=Project,proto3" json:"Project,omitempty"`
	Pagination *Paging `protobuf:"bytes,2,opt,name=Pagination,proto3" json:"Pagination,omitempty"`
}

func (x *QueryScreenRecorderRequest) Reset() {
	*x = QueryScreenRecorderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_screen_recorder_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryScreenRecorderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryScreenRecorderRequest) ProtoMessage() {}

func (x *QueryScreenRecorderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_screen_recorder_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryScreenRecorderRequest.ProtoReflect.Descriptor instead.
func (*QueryScreenRecorderRequest) Descriptor() ([]byte, []int) {
	return file_screen_recorder_proto_rawDescGZIP(), []int{3}
}

func (x *QueryScreenRecorderRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *QueryScreenRecorderRequest) GetPagination() *Paging {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type StartScrecorderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Framerate      int64   `protobuf:"varint,1,opt,name=Framerate,proto3" json:"Framerate,omitempty"`
	ResolutionSize string  `protobuf:"bytes,2,opt,name=ResolutionSize,proto3" json:"ResolutionSize,omitempty"`
	CoefficientPTS float64 `protobuf:"fixed64,3,opt,name=CoefficientPTS,proto3" json:"CoefficientPTS,omitempty"`
	DisableMouse   bool    `protobuf:"varint,4,opt,name=DisableMouse,proto3" json:"DisableMouse,omitempty"`
}

func (x *StartScrecorderRequest) Reset() {
	*x = StartScrecorderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_screen_recorder_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartScrecorderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartScrecorderRequest) ProtoMessage() {}

func (x *StartScrecorderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_screen_recorder_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartScrecorderRequest.ProtoReflect.Descriptor instead.
func (*StartScrecorderRequest) Descriptor() ([]byte, []int) {
	return file_screen_recorder_proto_rawDescGZIP(), []int{4}
}

func (x *StartScrecorderRequest) GetFramerate() int64 {
	if x != nil {
		return x.Framerate
	}
	return 0
}

func (x *StartScrecorderRequest) GetResolutionSize() string {
	if x != nil {
		return x.ResolutionSize
	}
	return ""
}

func (x *StartScrecorderRequest) GetCoefficientPTS() float64 {
	if x != nil {
		return x.CoefficientPTS
	}
	return 0
}

func (x *StartScrecorderRequest) GetDisableMouse() bool {
	if x != nil {
		return x.DisableMouse
	}
	return false
}

type InstallScrecorderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proxy string `protobuf:"bytes,1,opt,name=Proxy,proto3" json:"Proxy,omitempty"`
}

func (x *InstallScrecorderRequest) Reset() {
	*x = InstallScrecorderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_screen_recorder_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallScrecorderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallScrecorderRequest) ProtoMessage() {}

func (x *InstallScrecorderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_screen_recorder_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallScrecorderRequest.ProtoReflect.Descriptor instead.
func (*InstallScrecorderRequest) Descriptor() ([]byte, []int) {
	return file_screen_recorder_proto_rawDescGZIP(), []int{5}
}

func (x *InstallScrecorderRequest) GetProxy() string {
	if x != nil {
		return x.Proxy
	}
	return ""
}

type IsScrecorderReadyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IsScrecorderReadyRequest) Reset() {
	*x = IsScrecorderReadyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_screen_recorder_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsScrecorderReadyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsScrecorderReadyRequest) ProtoMessage() {}

func (x *IsScrecorderReadyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_screen_recorder_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsScrecorderReadyRequest.ProtoReflect.Descriptor instead.
func (*IsScrecorderReadyRequest) Descriptor() ([]byte, []int) {
	return file_screen_recorder_proto_rawDescGZIP(), []int{6}
}

var File_screen_recorder_proto protoreflect.FileDescriptor

var file_screen_recorder_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x79, 0x70, 0x62, 0x1a, 0x0d, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x0e,
	0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x6f,
	0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x6f,
	0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x89, 0x01, 0x0a,
	0x1b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x79, 0x70, 0x62,
	0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2b, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x79, 0x70, 0x62, 0x2e,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x43, 0x0a, 0x19, 0x49, 0x73, 0x53, 0x63,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x02, 0x4f, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x63, 0x0a,
	0x1a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2b, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x79, 0x70, 0x62, 0x2e,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xaa, 0x01, 0x0a, 0x16, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x63, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x52,
	0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x43, 0x6f, 0x65, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65,
	0x6e, 0x74, 0x50, 0x54, 0x53, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x43, 0x6f, 0x65,
	0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x54, 0x53, 0x12, 0x22, 0x0a, 0x0c, 0x44,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x6f, 0x75, 0x73, 0x65, 0x22,
	0x30, 0x0a, 0x18, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53, 0x63, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x22, 0x1a, 0x0a, 0x18, 0x49, 0x73, 0x53, 0x63, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xd0, 0x02,
	0x0a, 0x15, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x11, 0x49, 0x73, 0x53, 0x63, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x79, 0x12, 0x1d, 0x2e, 0x79,
	0x70, 0x62, 0x2e, 0x49, 0x73, 0x53, 0x63, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x79, 0x70,
	0x62, 0x2e, 0x49, 0x73, 0x53, 0x63, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x61, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x11, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53, 0x63, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x1d, 0x2e, 0x79, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53, 0x63,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0f, 0x2e, 0x79, 0x70, 0x62, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x30, 0x01, 0x12, 0x41, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x63, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x79, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x53, 0x63, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x79, 0x70, 0x62, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x30, 0x01, 0x12, 0x59, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x63,
	0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x2e,
	0x79, 0x70, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x79, 0x70, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x07, 0x5a, 0x05, 0x2f, 0x3b, 0x79, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_screen_recorder_proto_rawDescOnce sync.Once
	file_screen_recorder_proto_rawDescData = file_screen_recorder_proto_rawDesc
)

func file_screen_recorder_proto_rawDescGZIP() []byte {
	file_screen_recorder_proto_rawDescOnce.Do(func() {
		file_screen_recorder_proto_rawDescData = protoimpl.X.CompressGZIP(file_screen_recorder_proto_rawDescData)
	})
	return file_screen_recorder_proto_rawDescData
}

var file_screen_recorder_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_screen_recorder_proto_goTypes = []interface{}{
	(*ScreenRecorder)(nil),              // 0: ypb.ScreenRecorder
	(*QueryScreenRecorderResponse)(nil), // 1: ypb.QueryScreenRecorderResponse
	(*IsScrecorderReadyResponse)(nil),   // 2: ypb.IsScrecorderReadyResponse
	(*QueryScreenRecorderRequest)(nil),  // 3: ypb.QueryScreenRecorderRequest
	(*StartScrecorderRequest)(nil),      // 4: ypb.StartScrecorderRequest
	(*InstallScrecorderRequest)(nil),    // 5: ypb.InstallScrecorderRequest
	(*IsScrecorderReadyRequest)(nil),    // 6: ypb.IsScrecorderReadyRequest
	(*Paging)(nil),                      // 7: ypb.Paging
	(*ExecResult)(nil),                  // 8: ypb.ExecResult
}
var file_screen_recorder_proto_depIdxs = []int32{
	0, // 0: ypb.QueryScreenRecorderResponse.Data:type_name -> ypb.ScreenRecorder
	7, // 1: ypb.QueryScreenRecorderResponse.Pagination:type_name -> ypb.Paging
	7, // 2: ypb.QueryScreenRecorderRequest.Pagination:type_name -> ypb.Paging
	6, // 3: ypb.ScreenRecorderService.IsScrecorderReady:input_type -> ypb.IsScrecorderReadyRequest
	5, // 4: ypb.ScreenRecorderService.InstallScrecorder:input_type -> ypb.InstallScrecorderRequest
	4, // 5: ypb.ScreenRecorderService.StartScrecorder:input_type -> ypb.StartScrecorderRequest
	3, // 6: ypb.ScreenRecorderService.QueryScreenRecorders:input_type -> ypb.QueryScreenRecorderRequest
	2, // 7: ypb.ScreenRecorderService.IsScrecorderReady:output_type -> ypb.IsScrecorderReadyResponse
	8, // 8: ypb.ScreenRecorderService.InstallScrecorder:output_type -> ypb.ExecResult
	8, // 9: ypb.ScreenRecorderService.StartScrecorder:output_type -> ypb.ExecResult
	1, // 10: ypb.ScreenRecorderService.QueryScreenRecorders:output_type -> ypb.QueryScreenRecorderResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_screen_recorder_proto_init() }
func file_screen_recorder_proto_init() {
	if File_screen_recorder_proto != nil {
		return
	}
	file_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_screen_recorder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScreenRecorder); i {
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
		file_screen_recorder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryScreenRecorderResponse); i {
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
		file_screen_recorder_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsScrecorderReadyResponse); i {
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
		file_screen_recorder_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryScreenRecorderRequest); i {
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
		file_screen_recorder_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartScrecorderRequest); i {
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
		file_screen_recorder_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallScrecorderRequest); i {
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
		file_screen_recorder_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsScrecorderReadyRequest); i {
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
			RawDescriptor: file_screen_recorder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_screen_recorder_proto_goTypes,
		DependencyIndexes: file_screen_recorder_proto_depIdxs,
		MessageInfos:      file_screen_recorder_proto_msgTypes,
	}.Build()
	File_screen_recorder_proto = out.File
	file_screen_recorder_proto_rawDesc = nil
	file_screen_recorder_proto_goTypes = nil
	file_screen_recorder_proto_depIdxs = nil
}

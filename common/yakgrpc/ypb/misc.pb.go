// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: misc.proto

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

type ResetAndInvalidUserDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResetAndInvalidUserDataRequest) Reset() {
	*x = ResetAndInvalidUserDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_misc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetAndInvalidUserDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetAndInvalidUserDataRequest) ProtoMessage() {}

func (x *ResetAndInvalidUserDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_misc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetAndInvalidUserDataRequest.ProtoReflect.Descriptor instead.
func (*ResetAndInvalidUserDataRequest) Descriptor() ([]byte, []int) {
	return file_misc_proto_rawDescGZIP(), []int{0}
}

type IsPrivilegedForNetRawResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsPrivileged  bool   `protobuf:"varint,1,opt,name=IsPrivileged,proto3" json:"IsPrivileged,omitempty"`
	Advice        string `protobuf:"bytes,2,opt,name=Advice,proto3" json:"Advice,omitempty"`
	AdviceVerbose string `protobuf:"bytes,3,opt,name=AdviceVerbose,proto3" json:"AdviceVerbose,omitempty"`
}

func (x *IsPrivilegedForNetRawResponse) Reset() {
	*x = IsPrivilegedForNetRawResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_misc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsPrivilegedForNetRawResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsPrivilegedForNetRawResponse) ProtoMessage() {}

func (x *IsPrivilegedForNetRawResponse) ProtoReflect() protoreflect.Message {
	mi := &file_misc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsPrivilegedForNetRawResponse.ProtoReflect.Descriptor instead.
func (*IsPrivilegedForNetRawResponse) Descriptor() ([]byte, []int) {
	return file_misc_proto_rawDescGZIP(), []int{1}
}

func (x *IsPrivilegedForNetRawResponse) GetIsPrivileged() bool {
	if x != nil {
		return x.IsPrivileged
	}
	return false
}

func (x *IsPrivilegedForNetRawResponse) GetAdvice() string {
	if x != nil {
		return x.Advice
	}
	return ""
}

func (x *IsPrivilegedForNetRawResponse) GetAdviceVerbose() string {
	if x != nil {
		return x.AdviceVerbose
	}
	return ""
}

var File_misc_proto protoreflect.FileDescriptor

var file_misc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x69, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x79, 0x70,
	0x62, 0x1a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x20, 0x0a, 0x1e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x41, 0x6e, 0x64, 0x49, 0x6e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x81, 0x01, 0x0a, 0x1d, 0x49, 0x73, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65,
	0x67, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x4e, 0x65, 0x74, 0x52, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x49, 0x73, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c,
	0x65, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x49, 0x73, 0x50, 0x72,
	0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x64, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x64, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x24, 0x0a, 0x0d, 0x41, 0x64, 0x76, 0x69, 0x63, 0x65, 0x56, 0x65, 0x72, 0x62, 0x6f, 0x73,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x41, 0x64, 0x76, 0x69, 0x63, 0x65, 0x56,
	0x65, 0x72, 0x62, 0x6f, 0x73, 0x65, 0x32, 0xda, 0x01, 0x0a, 0x0b, 0x4d, 0x69, 0x73, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x17, 0x52, 0x65, 0x73, 0x65, 0x74, 0x41,
	0x6e, 0x64, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x23, 0x2e, 0x79, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x41, 0x6e, 0x64,
	0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x79, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x47, 0x0a, 0x15, 0x49, 0x73, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67,
	0x65, 0x64, 0x46, 0x6f, 0x72, 0x4e, 0x65, 0x74, 0x52, 0x61, 0x77, 0x12, 0x0a, 0x2e, 0x79, 0x70,
	0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x22, 0x2e, 0x79, 0x70, 0x62, 0x2e, 0x49, 0x73,
	0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x4e, 0x65, 0x74,
	0x52, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x1c, 0x50,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x50, 0x63, 0x61, 0x70, 0x12, 0x0a, 0x2e, 0x79, 0x70,
	0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x79, 0x70, 0x62, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x3b, 0x79, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_misc_proto_rawDescOnce sync.Once
	file_misc_proto_rawDescData = file_misc_proto_rawDesc
)

func file_misc_proto_rawDescGZIP() []byte {
	file_misc_proto_rawDescOnce.Do(func() {
		file_misc_proto_rawDescData = protoimpl.X.CompressGZIP(file_misc_proto_rawDescData)
	})
	return file_misc_proto_rawDescData
}

var file_misc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_misc_proto_goTypes = []interface{}{
	(*ResetAndInvalidUserDataRequest)(nil), // 0: ypb.ResetAndInvalidUserDataRequest
	(*IsPrivilegedForNetRawResponse)(nil),  // 1: ypb.IsPrivilegedForNetRawResponse
	(*Empty)(nil),                          // 2: ypb.Empty
}
var file_misc_proto_depIdxs = []int32{
	0, // 0: ypb.MiscService.ResetAndInvalidUserData:input_type -> ypb.ResetAndInvalidUserDataRequest
	2, // 1: ypb.MiscService.IsPrivilegedForNetRaw:input_type -> ypb.Empty
	2, // 2: ypb.MiscService.PromotePermissionForUserPcap:input_type -> ypb.Empty
	2, // 3: ypb.MiscService.ResetAndInvalidUserData:output_type -> ypb.Empty
	1, // 4: ypb.MiscService.IsPrivilegedForNetRaw:output_type -> ypb.IsPrivilegedForNetRawResponse
	2, // 5: ypb.MiscService.PromotePermissionForUserPcap:output_type -> ypb.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_misc_proto_init() }
func file_misc_proto_init() {
	if File_misc_proto != nil {
		return
	}
	file_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_misc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetAndInvalidUserDataRequest); i {
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
		file_misc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsPrivilegedForNetRawResponse); i {
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
			RawDescriptor: file_misc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_misc_proto_goTypes,
		DependencyIndexes: file_misc_proto_depIdxs,
		MessageInfos:      file_misc_proto_msgTypes,
	}.Build()
	File_misc_proto = out.File
	file_misc_proto_rawDesc = nil
	file_misc_proto_goTypes = nil
	file_misc_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: pixelfs/v1/meta.proto

package v1

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

type GetVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetVersionRequest) Reset() {
	*x = GetVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pixelfs_v1_meta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVersionRequest) ProtoMessage() {}

func (x *GetVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pixelfs_v1_meta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVersionRequest.ProtoReflect.Descriptor instead.
func (*GetVersionRequest) Descriptor() ([]byte, []int) {
	return file_pixelfs_v1_meta_proto_rawDescGZIP(), []int{0}
}

type GetVersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *GetVersionResponse) Reset() {
	*x = GetVersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pixelfs_v1_meta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVersionResponse) ProtoMessage() {}

func (x *GetVersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pixelfs_v1_meta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVersionResponse.ProtoReflect.Descriptor instead.
func (*GetVersionResponse) Descriptor() ([]byte, []int) {
	return file_pixelfs_v1_meta_proto_rawDescGZIP(), []int{1}
}

func (x *GetVersionResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetVersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_pixelfs_v1_meta_proto protoreflect.FileDescriptor

var file_pixelfs_v1_meta_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x66, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x66, 0x73,
	0x2e, 0x76, 0x31, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x5c, 0x0a, 0x0b,
	0x4d, 0x65, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x70, 0x69, 0x78, 0x65,
	0x6c, 0x66, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x69, 0x78, 0x65, 0x6c,
	0x66, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x66, 0x73,
	0x2f, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x66, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x69, 0x78,
	0x65, 0x6c, 0x66, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pixelfs_v1_meta_proto_rawDescOnce sync.Once
	file_pixelfs_v1_meta_proto_rawDescData = file_pixelfs_v1_meta_proto_rawDesc
)

func file_pixelfs_v1_meta_proto_rawDescGZIP() []byte {
	file_pixelfs_v1_meta_proto_rawDescOnce.Do(func() {
		file_pixelfs_v1_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_pixelfs_v1_meta_proto_rawDescData)
	})
	return file_pixelfs_v1_meta_proto_rawDescData
}

var file_pixelfs_v1_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pixelfs_v1_meta_proto_goTypes = []interface{}{
	(*GetVersionRequest)(nil),  // 0: pixelfs.v1.GetVersionRequest
	(*GetVersionResponse)(nil), // 1: pixelfs.v1.GetVersionResponse
}
var file_pixelfs_v1_meta_proto_depIdxs = []int32{
	0, // 0: pixelfs.v1.MetaService.GetVersion:input_type -> pixelfs.v1.GetVersionRequest
	1, // 1: pixelfs.v1.MetaService.GetVersion:output_type -> pixelfs.v1.GetVersionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pixelfs_v1_meta_proto_init() }
func file_pixelfs_v1_meta_proto_init() {
	if File_pixelfs_v1_meta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pixelfs_v1_meta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVersionRequest); i {
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
		file_pixelfs_v1_meta_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVersionResponse); i {
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
			RawDescriptor: file_pixelfs_v1_meta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pixelfs_v1_meta_proto_goTypes,
		DependencyIndexes: file_pixelfs_v1_meta_proto_depIdxs,
		MessageInfos:      file_pixelfs_v1_meta_proto_msgTypes,
	}.Build()
	File_pixelfs_v1_meta_proto = out.File
	file_pixelfs_v1_meta_proto_rawDesc = nil
	file_pixelfs_v1_meta_proto_goTypes = nil
	file_pixelfs_v1_meta_proto_depIdxs = nil
}

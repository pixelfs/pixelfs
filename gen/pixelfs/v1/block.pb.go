// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: pixelfs/v1/block.proto

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

type BlockType int32

const (
	BlockType_SIZE     BlockType = 0
	BlockType_DURATION BlockType = 1
)

// Enum value maps for BlockType.
var (
	BlockType_name = map[int32]string{
		0: "SIZE",
		1: "DURATION",
	}
	BlockType_value = map[string]int32{
		"SIZE":     0,
		"DURATION": 1,
	}
)

func (x BlockType) Enum() *BlockType {
	p := new(BlockType)
	*p = x
	return p
}

func (x BlockType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlockType) Descriptor() protoreflect.EnumDescriptor {
	return file_pixelfs_v1_block_proto_enumTypes[0].Descriptor()
}

func (BlockType) Type() protoreflect.EnumType {
	return &file_pixelfs_v1_block_proto_enumTypes[0]
}

func (x BlockType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlockType.Descriptor instead.
func (BlockType) EnumDescriptor() ([]byte, []int) {
	return file_pixelfs_v1_block_proto_rawDescGZIP(), []int{0}
}

type BlockStatus int32

const (
	BlockStatus_PENDING BlockStatus = 0
	BlockStatus_READY   BlockStatus = 1
)

// Enum value maps for BlockStatus.
var (
	BlockStatus_name = map[int32]string{
		0: "PENDING",
		1: "READY",
	}
	BlockStatus_value = map[string]int32{
		"PENDING": 0,
		"READY":   1,
	}
)

func (x BlockStatus) Enum() *BlockStatus {
	p := new(BlockStatus)
	*p = x
	return p
}

func (x BlockStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlockStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_pixelfs_v1_block_proto_enumTypes[1].Descriptor()
}

func (BlockStatus) Type() protoreflect.EnumType {
	return &file_pixelfs_v1_block_proto_enumTypes[1]
}

func (x BlockStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlockStatus.Descriptor instead.
func (BlockStatus) EnumDescriptor() ([]byte, []int) {
	return file_pixelfs_v1_block_proto_rawDescGZIP(), []int{1}
}

type BlockSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width   int32 `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height  int32 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Bitrate int32 `protobuf:"varint,3,opt,name=bitrate,proto3" json:"bitrate,omitempty"` // kbps
}

func (x *BlockSettings) Reset() {
	*x = BlockSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pixelfs_v1_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockSettings) ProtoMessage() {}

func (x *BlockSettings) ProtoReflect() protoreflect.Message {
	mi := &file_pixelfs_v1_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockSettings.ProtoReflect.Descriptor instead.
func (*BlockSettings) Descriptor() ([]byte, []int) {
	return file_pixelfs_v1_block_proto_rawDescGZIP(), []int{0}
}

func (x *BlockSettings) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *BlockSettings) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *BlockSettings) GetBitrate() int32 {
	if x != nil {
		return x.Bitrate
	}
	return 0
}

type GetBlockDurationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId        string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Hash          string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	BlockDuration int64  `protobuf:"varint,3,opt,name=block_duration,json=blockDuration,proto3" json:"block_duration,omitempty"`
}

func (x *GetBlockDurationRequest) Reset() {
	*x = GetBlockDurationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pixelfs_v1_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockDurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockDurationRequest) ProtoMessage() {}

func (x *GetBlockDurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pixelfs_v1_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockDurationRequest.ProtoReflect.Descriptor instead.
func (*GetBlockDurationRequest) Descriptor() ([]byte, []int) {
	return file_pixelfs_v1_block_proto_rawDescGZIP(), []int{1}
}

func (x *GetBlockDurationRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *GetBlockDurationRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *GetBlockDurationRequest) GetBlockDuration() int64 {
	if x != nil {
		return x.BlockDuration
	}
	return 0
}

type GetBlockDurationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data map[int64]float64 `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *GetBlockDurationResponse) Reset() {
	*x = GetBlockDurationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pixelfs_v1_block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockDurationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockDurationResponse) ProtoMessage() {}

func (x *GetBlockDurationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pixelfs_v1_block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockDurationResponse.ProtoReflect.Descriptor instead.
func (*GetBlockDurationResponse) Descriptor() ([]byte, []int) {
	return file_pixelfs_v1_block_proto_rawDescGZIP(), []int{2}
}

func (x *GetBlockDurationResponse) GetData() map[int64]float64 {
	if x != nil {
		return x.Data
	}
	return nil
}

type SetBlockDurationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId   string            `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Location string            `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Path     string            `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Data     map[int64]float64 `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *SetBlockDurationRequest) Reset() {
	*x = SetBlockDurationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pixelfs_v1_block_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetBlockDurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetBlockDurationRequest) ProtoMessage() {}

func (x *SetBlockDurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pixelfs_v1_block_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetBlockDurationRequest.ProtoReflect.Descriptor instead.
func (*SetBlockDurationRequest) Descriptor() ([]byte, []int) {
	return file_pixelfs_v1_block_proto_rawDescGZIP(), []int{3}
}

func (x *SetBlockDurationRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *SetBlockDurationRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *SetBlockDurationRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *SetBlockDurationRequest) GetData() map[int64]float64 {
	if x != nil {
		return x.Data
	}
	return nil
}

type SetBlockDurationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetBlockDurationResponse) Reset() {
	*x = SetBlockDurationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pixelfs_v1_block_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetBlockDurationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetBlockDurationResponse) ProtoMessage() {}

func (x *SetBlockDurationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pixelfs_v1_block_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetBlockDurationResponse.ProtoReflect.Descriptor instead.
func (*SetBlockDurationResponse) Descriptor() ([]byte, []int) {
	return file_pixelfs_v1_block_proto_rawDescGZIP(), []int{4}
}

var File_pixelfs_v1_block_proto protoreflect.FileDescriptor

var file_pixelfs_v1_block_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x66, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x66,
	0x73, 0x2e, 0x76, 0x31, 0x22, 0x57, 0x0a, 0x0d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x22, 0x6d, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x97, 0x01, 0x0a,
	0x18, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x66,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a,
	0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xde, 0x01, 0x0a, 0x17, 0x53, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x41, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x69, 0x78, 0x65,
	0x6c, 0x66, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37,
	0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1a, 0x0a, 0x18, 0x53, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2a, 0x23, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x53, 0x49, 0x5a, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x55,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x2a, 0x25, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49,
	0x4e, 0x47, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x01, 0x32,
	0xd0, 0x01, 0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x5f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x66, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x69, 0x78, 0x65,
	0x6c, 0x66, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5f, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x66, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x69, 0x78,
	0x65, 0x6c, 0x66, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x66, 0x73, 0x2f, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x66, 0x73,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x66, 0x73, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pixelfs_v1_block_proto_rawDescOnce sync.Once
	file_pixelfs_v1_block_proto_rawDescData = file_pixelfs_v1_block_proto_rawDesc
)

func file_pixelfs_v1_block_proto_rawDescGZIP() []byte {
	file_pixelfs_v1_block_proto_rawDescOnce.Do(func() {
		file_pixelfs_v1_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_pixelfs_v1_block_proto_rawDescData)
	})
	return file_pixelfs_v1_block_proto_rawDescData
}

var file_pixelfs_v1_block_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pixelfs_v1_block_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pixelfs_v1_block_proto_goTypes = []interface{}{
	(BlockType)(0),                   // 0: pixelfs.v1.BlockType
	(BlockStatus)(0),                 // 1: pixelfs.v1.BlockStatus
	(*BlockSettings)(nil),            // 2: pixelfs.v1.BlockSettings
	(*GetBlockDurationRequest)(nil),  // 3: pixelfs.v1.GetBlockDurationRequest
	(*GetBlockDurationResponse)(nil), // 4: pixelfs.v1.GetBlockDurationResponse
	(*SetBlockDurationRequest)(nil),  // 5: pixelfs.v1.SetBlockDurationRequest
	(*SetBlockDurationResponse)(nil), // 6: pixelfs.v1.SetBlockDurationResponse
	nil,                              // 7: pixelfs.v1.GetBlockDurationResponse.DataEntry
	nil,                              // 8: pixelfs.v1.SetBlockDurationRequest.DataEntry
}
var file_pixelfs_v1_block_proto_depIdxs = []int32{
	7, // 0: pixelfs.v1.GetBlockDurationResponse.data:type_name -> pixelfs.v1.GetBlockDurationResponse.DataEntry
	8, // 1: pixelfs.v1.SetBlockDurationRequest.data:type_name -> pixelfs.v1.SetBlockDurationRequest.DataEntry
	3, // 2: pixelfs.v1.BlockService.GetBlockDuration:input_type -> pixelfs.v1.GetBlockDurationRequest
	5, // 3: pixelfs.v1.BlockService.SetBlockDuration:input_type -> pixelfs.v1.SetBlockDurationRequest
	4, // 4: pixelfs.v1.BlockService.GetBlockDuration:output_type -> pixelfs.v1.GetBlockDurationResponse
	6, // 5: pixelfs.v1.BlockService.SetBlockDuration:output_type -> pixelfs.v1.SetBlockDurationResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pixelfs_v1_block_proto_init() }
func file_pixelfs_v1_block_proto_init() {
	if File_pixelfs_v1_block_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pixelfs_v1_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockSettings); i {
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
		file_pixelfs_v1_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockDurationRequest); i {
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
		file_pixelfs_v1_block_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockDurationResponse); i {
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
		file_pixelfs_v1_block_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetBlockDurationRequest); i {
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
		file_pixelfs_v1_block_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetBlockDurationResponse); i {
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
			RawDescriptor: file_pixelfs_v1_block_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pixelfs_v1_block_proto_goTypes,
		DependencyIndexes: file_pixelfs_v1_block_proto_depIdxs,
		EnumInfos:         file_pixelfs_v1_block_proto_enumTypes,
		MessageInfos:      file_pixelfs_v1_block_proto_msgTypes,
	}.Build()
	File_pixelfs_v1_block_proto = out.File
	file_pixelfs_v1_block_proto_rawDesc = nil
	file_pixelfs_v1_block_proto_goTypes = nil
	file_pixelfs_v1_block_proto_depIdxs = nil
}

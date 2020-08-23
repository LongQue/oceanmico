// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: proto/rand/rand.proto

package rand

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Max int64 `protobuf:"varint,1,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *RandRequest) Reset() {
	*x = RandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rand_rand_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandRequest) ProtoMessage() {}

func (x *RandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rand_rand_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandRequest.ProtoReflect.Descriptor instead.
func (*RandRequest) Descriptor() ([]byte, []int) {
	return file_proto_rand_rand_proto_rawDescGZIP(), []int{0}
}

func (x *RandRequest) GetMax() int64 {
	if x != nil {
		return x.Max
	}
	return 0
}

type RandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *RandResponse) Reset() {
	*x = RandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rand_rand_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandResponse) ProtoMessage() {}

func (x *RandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rand_rand_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandResponse.ProtoReflect.Descriptor instead.
func (*RandResponse) Descriptor() ([]byte, []int) {
	return file_proto_rand_rand_proto_rawDescGZIP(), []int{1}
}

func (x *RandResponse) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_proto_rand_rand_proto protoreflect.FileDescriptor

var file_proto_rand_rand_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x6e, 0x64, 0x2f, 0x72, 0x61, 0x6e,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x0b, 0x52, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x22, 0x26, 0x0a, 0x0c, 0x52, 0x61, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x32, 0x30, 0x0a, 0x04, 0x52, 0x61, 0x6e, 0x64, 0x12, 0x28, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52,
	0x61, 0x6e, 0x64, 0x12, 0x0c, 0x2e, 0x52, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0d, 0x2e, 0x52, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x6e, 0x64,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rand_rand_proto_rawDescOnce sync.Once
	file_proto_rand_rand_proto_rawDescData = file_proto_rand_rand_proto_rawDesc
)

func file_proto_rand_rand_proto_rawDescGZIP() []byte {
	file_proto_rand_rand_proto_rawDescOnce.Do(func() {
		file_proto_rand_rand_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rand_rand_proto_rawDescData)
	})
	return file_proto_rand_rand_proto_rawDescData
}

var file_proto_rand_rand_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_rand_rand_proto_goTypes = []interface{}{
	(*RandRequest)(nil),  // 0: RandRequest
	(*RandResponse)(nil), // 1: RandResponse
}
var file_proto_rand_rand_proto_depIdxs = []int32{
	0, // 0: Rand.GetRand:input_type -> RandRequest
	1, // 1: Rand.GetRand:output_type -> RandResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_rand_rand_proto_init() }
func file_proto_rand_rand_proto_init() {
	if File_proto_rand_rand_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rand_rand_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandRequest); i {
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
		file_proto_rand_rand_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandResponse); i {
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
			RawDescriptor: file_proto_rand_rand_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rand_rand_proto_goTypes,
		DependencyIndexes: file_proto_rand_rand_proto_depIdxs,
		MessageInfos:      file_proto_rand_rand_proto_msgTypes,
	}.Build()
	File_proto_rand_rand_proto = out.File
	file_proto_rand_rand_proto_rawDesc = nil
	file_proto_rand_rand_proto_goTypes = nil
	file_proto_rand_rand_proto_depIdxs = nil
}
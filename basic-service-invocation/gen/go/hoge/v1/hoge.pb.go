// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: hoge/v1/hoge.proto

package hoge

import (
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

type FugaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FugaRequest) Reset() {
	*x = FugaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hoge_v1_hoge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FugaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FugaRequest) ProtoMessage() {}

func (x *FugaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hoge_v1_hoge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FugaRequest.ProtoReflect.Descriptor instead.
func (*FugaRequest) Descriptor() ([]byte, []int) {
	return file_hoge_v1_hoge_proto_rawDescGZIP(), []int{0}
}

type FugaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *FugaResponse) Reset() {
	*x = FugaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hoge_v1_hoge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FugaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FugaResponse) ProtoMessage() {}

func (x *FugaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hoge_v1_hoge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FugaResponse.ProtoReflect.Descriptor instead.
func (*FugaResponse) Descriptor() ([]byte, []int) {
	return file_hoge_v1_hoge_proto_rawDescGZIP(), []int{1}
}

func (x *FugaResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_hoge_v1_hoge_proto protoreflect.FileDescriptor

var file_hoge_v1_hoge_proto_rawDesc = []byte{
	0x0a, 0x12, 0x68, 0x6f, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x68, 0x6f, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x46, 0x75, 0x67, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x20, 0x0a, 0x0c, 0x46, 0x75, 0x67, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x32, 0x5b, 0x0a, 0x04, 0x48, 0x6f, 0x67, 0x65, 0x12, 0x53, 0x0a, 0x04, 0x46, 0x75,
	0x67, 0x61, 0x12, 0x1a, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x68, 0x6f, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x75, 0x67, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x68, 0x6f, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x75, 0x67, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x68, 0x6f, 0x67, 0x65, 0x2f, 0x66, 0x75, 0x67, 0x61, 0x42,
	0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x7a,
	0x6d, 0x61, 0x6b, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x69, 0x6e, 0x76, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x6f, 0x67, 0x65, 0x2f,
	0x76, 0x31, 0x3b, 0x68, 0x6f, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hoge_v1_hoge_proto_rawDescOnce sync.Once
	file_hoge_v1_hoge_proto_rawDescData = file_hoge_v1_hoge_proto_rawDesc
)

func file_hoge_v1_hoge_proto_rawDescGZIP() []byte {
	file_hoge_v1_hoge_proto_rawDescOnce.Do(func() {
		file_hoge_v1_hoge_proto_rawDescData = protoimpl.X.CompressGZIP(file_hoge_v1_hoge_proto_rawDescData)
	})
	return file_hoge_v1_hoge_proto_rawDescData
}

var file_hoge_v1_hoge_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_hoge_v1_hoge_proto_goTypes = []interface{}{
	(*FugaRequest)(nil),  // 0: basic.hoge.v1.FugaRequest
	(*FugaResponse)(nil), // 1: basic.hoge.v1.FugaResponse
}
var file_hoge_v1_hoge_proto_depIdxs = []int32{
	0, // 0: basic.hoge.v1.Hoge.Fuga:input_type -> basic.hoge.v1.FugaRequest
	1, // 1: basic.hoge.v1.Hoge.Fuga:output_type -> basic.hoge.v1.FugaResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hoge_v1_hoge_proto_init() }
func file_hoge_v1_hoge_proto_init() {
	if File_hoge_v1_hoge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hoge_v1_hoge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FugaRequest); i {
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
		file_hoge_v1_hoge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FugaResponse); i {
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
			RawDescriptor: file_hoge_v1_hoge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hoge_v1_hoge_proto_goTypes,
		DependencyIndexes: file_hoge_v1_hoge_proto_depIdxs,
		MessageInfos:      file_hoge_v1_hoge_proto_msgTypes,
	}.Build()
	File_hoge_v1_hoge_proto = out.File
	file_hoge_v1_hoge_proto_rawDesc = nil
	file_hoge_v1_hoge_proto_goTypes = nil
	file_hoge_v1_hoge_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: hourhand/v1/hourhand.proto

package hourhand

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

type NowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NowRequest) Reset() {
	*x = NowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hourhand_v1_hourhand_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NowRequest) ProtoMessage() {}

func (x *NowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hourhand_v1_hourhand_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NowRequest.ProtoReflect.Descriptor instead.
func (*NowRequest) Descriptor() ([]byte, []int) {
	return file_hourhand_v1_hourhand_proto_rawDescGZIP(), []int{0}
}

type NowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hour string `protobuf:"bytes,1,opt,name=hour,proto3" json:"hour,omitempty"`
}

func (x *NowResponse) Reset() {
	*x = NowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hourhand_v1_hourhand_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NowResponse) ProtoMessage() {}

func (x *NowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hourhand_v1_hourhand_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NowResponse.ProtoReflect.Descriptor instead.
func (*NowResponse) Descriptor() ([]byte, []int) {
	return file_hourhand_v1_hourhand_proto_rawDescGZIP(), []int{1}
}

func (x *NowResponse) GetHour() string {
	if x != nil {
		return x.Hour
	}
	return ""
}

var File_hourhand_v1_hourhand_proto protoreflect.FileDescriptor

var file_hourhand_v1_hourhand_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x68, 0x6f, 0x75, 0x72, 0x68, 0x61, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f,
	0x75, 0x72, 0x68, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x64, 0x61,
	0x70, 0x72, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x68, 0x6f, 0x75, 0x72, 0x68, 0x61, 0x6e, 0x64,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x0c, 0x0a, 0x0a, 0x4e, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x21, 0x0a, 0x0b, 0x4e, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f,
	0x75, 0x72, 0x32, 0x58, 0x0a, 0x08, 0x48, 0x6f, 0x75, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x12, 0x4c,
	0x0a, 0x03, 0x4e, 0x6f, 0x77, 0x12, 0x21, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x63, 0x6c, 0x6f, 0x63,
	0x6b, 0x2e, 0x68, 0x6f, 0x75, 0x72, 0x68, 0x61, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x63,
	0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x68, 0x6f, 0x75, 0x72, 0x68, 0x61, 0x6e, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3f, 0x5a, 0x3d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x7a, 0x6d, 0x61, 0x6b,
	0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x64, 0x61, 0x70, 0x72, 0x63,
	0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x6f, 0x75, 0x72, 0x68, 0x61, 0x6e,
	0x64, 0x2f, 0x76, 0x31, 0x3b, 0x68, 0x6f, 0x75, 0x72, 0x68, 0x61, 0x6e, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hourhand_v1_hourhand_proto_rawDescOnce sync.Once
	file_hourhand_v1_hourhand_proto_rawDescData = file_hourhand_v1_hourhand_proto_rawDesc
)

func file_hourhand_v1_hourhand_proto_rawDescGZIP() []byte {
	file_hourhand_v1_hourhand_proto_rawDescOnce.Do(func() {
		file_hourhand_v1_hourhand_proto_rawDescData = protoimpl.X.CompressGZIP(file_hourhand_v1_hourhand_proto_rawDescData)
	})
	return file_hourhand_v1_hourhand_proto_rawDescData
}

var file_hourhand_v1_hourhand_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_hourhand_v1_hourhand_proto_goTypes = []interface{}{
	(*NowRequest)(nil),  // 0: daprclock.hourhand.v1.NowRequest
	(*NowResponse)(nil), // 1: daprclock.hourhand.v1.NowResponse
}
var file_hourhand_v1_hourhand_proto_depIdxs = []int32{
	0, // 0: daprclock.hourhand.v1.HourHand.Now:input_type -> daprclock.hourhand.v1.NowRequest
	1, // 1: daprclock.hourhand.v1.HourHand.Now:output_type -> daprclock.hourhand.v1.NowResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hourhand_v1_hourhand_proto_init() }
func file_hourhand_v1_hourhand_proto_init() {
	if File_hourhand_v1_hourhand_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hourhand_v1_hourhand_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NowRequest); i {
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
		file_hourhand_v1_hourhand_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NowResponse); i {
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
			RawDescriptor: file_hourhand_v1_hourhand_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hourhand_v1_hourhand_proto_goTypes,
		DependencyIndexes: file_hourhand_v1_hourhand_proto_depIdxs,
		MessageInfos:      file_hourhand_v1_hourhand_proto_msgTypes,
	}.Build()
	File_hourhand_v1_hourhand_proto = out.File
	file_hourhand_v1_hourhand_proto_rawDesc = nil
	file_hourhand_v1_hourhand_proto_goTypes = nil
	file_hourhand_v1_hourhand_proto_depIdxs = nil
}

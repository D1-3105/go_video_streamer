// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.6.1
// source: ClientStreamCreator.proto

package client_stream_creator

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClientNewStreamRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	StreamT       uint32                 `protobuf:"varint,2,opt,name=streamT,proto3" json:"streamT,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClientNewStreamRequest) Reset() {
	*x = ClientNewStreamRequest{}
	mi := &file_ClientStreamCreator_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientNewStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientNewStreamRequest) ProtoMessage() {}

func (x *ClientNewStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ClientStreamCreator_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientNewStreamRequest.ProtoReflect.Descriptor instead.
func (*ClientNewStreamRequest) Descriptor() ([]byte, []int) {
	return file_ClientStreamCreator_proto_rawDescGZIP(), []int{0}
}

func (x *ClientNewStreamRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ClientNewStreamRequest) GetStreamT() uint32 {
	if x != nil {
		return x.StreamT
	}
	return 0
}

var File_ClientStreamCreator_proto protoreflect.FileDescriptor

var file_ClientStreamCreator_proto_rawDesc = string([]byte{
	0x0a, 0x19, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x6f, 0x72, 0x22, 0x44, 0x0a, 0x16, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x65, 0x77, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x6f, 0x5f, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x3b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_ClientStreamCreator_proto_rawDescOnce sync.Once
	file_ClientStreamCreator_proto_rawDescData []byte
)

func file_ClientStreamCreator_proto_rawDescGZIP() []byte {
	file_ClientStreamCreator_proto_rawDescOnce.Do(func() {
		file_ClientStreamCreator_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_ClientStreamCreator_proto_rawDesc), len(file_ClientStreamCreator_proto_rawDesc)))
	})
	return file_ClientStreamCreator_proto_rawDescData
}

var file_ClientStreamCreator_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ClientStreamCreator_proto_goTypes = []any{
	(*ClientNewStreamRequest)(nil), // 0: client_stream_creator.ClientNewStreamRequest
}
var file_ClientStreamCreator_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ClientStreamCreator_proto_init() }
func file_ClientStreamCreator_proto_init() {
	if File_ClientStreamCreator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_ClientStreamCreator_proto_rawDesc), len(file_ClientStreamCreator_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ClientStreamCreator_proto_goTypes,
		DependencyIndexes: file_ClientStreamCreator_proto_depIdxs,
		MessageInfos:      file_ClientStreamCreator_proto_msgTypes,
	}.Build()
	File_ClientStreamCreator_proto = out.File
	file_ClientStreamCreator_proto_goTypes = nil
	file_ClientStreamCreator_proto_depIdxs = nil
}

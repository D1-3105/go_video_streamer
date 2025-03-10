// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.6.1
// source: InputStreamShard.proto

package InputStreamShard

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

type StreamShard struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ImageData     []byte                 `protobuf:"bytes,1,opt,name=image_data,json=imageData,proto3" json:"image_data,omitempty"`
	Width         uint32                 `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height        uint32                 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Fps           float32                `protobuf:"fixed32,4,opt,name=fps,proto3" json:"fps,omitempty"`
	Gzipped       bool                   `protobuf:"varint,5,opt,name=gzipped,proto3" json:"gzipped,omitempty"`
	MatType       uint32                 `protobuf:"varint,6,opt,name=matType,proto3" json:"matType,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamShard) Reset() {
	*x = StreamShard{}
	mi := &file_InputStreamShard_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamShard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamShard) ProtoMessage() {}

func (x *StreamShard) ProtoReflect() protoreflect.Message {
	mi := &file_InputStreamShard_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamShard.ProtoReflect.Descriptor instead.
func (*StreamShard) Descriptor() ([]byte, []int) {
	return file_InputStreamShard_proto_rawDescGZIP(), []int{0}
}

func (x *StreamShard) GetImageData() []byte {
	if x != nil {
		return x.ImageData
	}
	return nil
}

func (x *StreamShard) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *StreamShard) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *StreamShard) GetFps() float32 {
	if x != nil {
		return x.Fps
	}
	return 0
}

func (x *StreamShard) GetGzipped() bool {
	if x != nil {
		return x.Gzipped
	}
	return false
}

func (x *StreamShard) GetMatType() uint32 {
	if x != nil {
		return x.MatType
	}
	return 0
}

var File_InputStreamShard_proto protoreflect.FileDescriptor

var file_InputStreamShard_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x68, 0x61,
	0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x68, 0x61, 0x72, 0x64, 0x22, 0xa0, 0x01, 0x0a, 0x0b, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x68, 0x61, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x70, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x66, 0x70, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x7a, 0x69,
	0x70, 0x70, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x67, 0x7a, 0x69, 0x70,
	0x70, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x3e, 0x5a,
	0x3c, 0x67, 0x6f, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x68, 0x61, 0x72, 0x64, 0x3b, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x68, 0x61, 0x72, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_InputStreamShard_proto_rawDescOnce sync.Once
	file_InputStreamShard_proto_rawDescData []byte
)

func file_InputStreamShard_proto_rawDescGZIP() []byte {
	file_InputStreamShard_proto_rawDescOnce.Do(func() {
		file_InputStreamShard_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_InputStreamShard_proto_rawDesc), len(file_InputStreamShard_proto_rawDesc)))
	})
	return file_InputStreamShard_proto_rawDescData
}

var file_InputStreamShard_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_InputStreamShard_proto_goTypes = []any{
	(*StreamShard)(nil), // 0: InputStreamShard.StreamShard
}
var file_InputStreamShard_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_InputStreamShard_proto_init() }
func file_InputStreamShard_proto_init() {
	if File_InputStreamShard_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_InputStreamShard_proto_rawDesc), len(file_InputStreamShard_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_InputStreamShard_proto_goTypes,
		DependencyIndexes: file_InputStreamShard_proto_depIdxs,
		MessageInfos:      file_InputStreamShard_proto_msgTypes,
	}.Build()
	File_InputStreamShard_proto = out.File
	file_InputStreamShard_proto_goTypes = nil
	file_InputStreamShard_proto_depIdxs = nil
}

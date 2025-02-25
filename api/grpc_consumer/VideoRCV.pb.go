// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.6.1
// source: VideoRCV.proto

package grpc_consumer

import (
	InputStreamShard "go_video_streamer/internal/InputStreamShard"
	base_rpc "go_video_streamer/internal/base_rpc"
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

type NamedFrame struct {
	state         protoimpl.MessageState        `protogen:"open.v1"`
	Stream        string                        `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty"`
	Shard         *InputStreamShard.StreamShard `protobuf:"bytes,2,opt,name=shard,proto3" json:"shard,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NamedFrame) Reset() {
	*x = NamedFrame{}
	mi := &file_VideoRCV_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NamedFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamedFrame) ProtoMessage() {}

func (x *NamedFrame) ProtoReflect() protoreflect.Message {
	mi := &file_VideoRCV_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamedFrame.ProtoReflect.Descriptor instead.
func (*NamedFrame) Descriptor() ([]byte, []int) {
	return file_VideoRCV_proto_rawDescGZIP(), []int{0}
}

func (x *NamedFrame) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

func (x *NamedFrame) GetShard() *InputStreamShard.StreamShard {
	if x != nil {
		return x.Shard
	}
	return nil
}

type EditStreamOperationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        int32                  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EditStreamOperationResponse) Reset() {
	*x = EditStreamOperationResponse{}
	mi := &file_VideoRCV_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EditStreamOperationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditStreamOperationResponse) ProtoMessage() {}

func (x *EditStreamOperationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_VideoRCV_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditStreamOperationResponse.ProtoReflect.Descriptor instead.
func (*EditStreamOperationResponse) Descriptor() ([]byte, []int) {
	return file_VideoRCV_proto_rawDescGZIP(), []int{1}
}

func (x *EditStreamOperationResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *EditStreamOperationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type VideoStreamResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        uint32                 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VideoStreamResponse) Reset() {
	*x = VideoStreamResponse{}
	mi := &file_VideoRCV_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VideoStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoStreamResponse) ProtoMessage() {}

func (x *VideoStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_VideoRCV_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoStreamResponse.ProtoReflect.Descriptor instead.
func (*VideoStreamResponse) Descriptor() ([]byte, []int) {
	return file_VideoRCV_proto_rawDescGZIP(), []int{2}
}

func (x *VideoStreamResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_VideoRCV_proto protoreflect.FileDescriptor

var file_VideoRCV_proto_rawDesc = string([]byte{
	0x0a, 0x0e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x43, 0x56, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x1a,
	0x16, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x68, 0x61, 0x72,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x52, 0x50, 0x43, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x0a, 0x4e, 0x61, 0x6d,
	0x65, 0x64, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x33, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x68, 0x61, 0x72,
	0x64, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x68, 0x61, 0x72, 0x64, 0x52, 0x05, 0x73,
	0x68, 0x61, 0x72, 0x64, 0x22, 0x4f, 0x0a, 0x1b, 0x45, 0x64, 0x69, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2d, 0x0a, 0x13, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x32, 0xf7, 0x01, 0x0a, 0x08, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x43,
	0x56, 0x12, 0x4d, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x73, 0x12, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x1a, 0x22, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4c, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x13, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x1a, 0x2a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e,
	0x0a, 0x08, 0x52, 0x4d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x1a, 0x2a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33,
	0x5a, 0x31, 0x67, 0x6f, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_VideoRCV_proto_rawDescOnce sync.Once
	file_VideoRCV_proto_rawDescData []byte
)

func file_VideoRCV_proto_rawDescGZIP() []byte {
	file_VideoRCV_proto_rawDescOnce.Do(func() {
		file_VideoRCV_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_VideoRCV_proto_rawDesc), len(file_VideoRCV_proto_rawDesc)))
	})
	return file_VideoRCV_proto_rawDescData
}

var file_VideoRCV_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_VideoRCV_proto_goTypes = []any{
	(*NamedFrame)(nil),                   // 0: grpc_consumer.NamedFrame
	(*EditStreamOperationResponse)(nil),  // 1: grpc_consumer.EditStreamOperationResponse
	(*VideoStreamResponse)(nil),          // 2: grpc_consumer.VideoStreamResponse
	(*InputStreamShard.StreamShard)(nil), // 3: InputStreamShard.StreamShard
	(*base_rpc.NewStream)(nil),           // 4: base_rpc.NewStream
	(*base_rpc.RemoveStream)(nil),        // 5: base_rpc.RemoveStream
}
var file_VideoRCV_proto_depIdxs = []int32{
	3, // 0: grpc_consumer.NamedFrame.shard:type_name -> InputStreamShard.StreamShard
	0, // 1: grpc_consumer.VideoRCV.StreamFrames:input_type -> grpc_consumer.NamedFrame
	4, // 2: grpc_consumer.VideoRCV.AddStream:input_type -> base_rpc.NewStream
	5, // 3: grpc_consumer.VideoRCV.RMStream:input_type -> base_rpc.RemoveStream
	2, // 4: grpc_consumer.VideoRCV.StreamFrames:output_type -> grpc_consumer.VideoStreamResponse
	1, // 5: grpc_consumer.VideoRCV.AddStream:output_type -> grpc_consumer.EditStreamOperationResponse
	1, // 6: grpc_consumer.VideoRCV.RMStream:output_type -> grpc_consumer.EditStreamOperationResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_VideoRCV_proto_init() }
func file_VideoRCV_proto_init() {
	if File_VideoRCV_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_VideoRCV_proto_rawDesc), len(file_VideoRCV_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_VideoRCV_proto_goTypes,
		DependencyIndexes: file_VideoRCV_proto_depIdxs,
		MessageInfos:      file_VideoRCV_proto_msgTypes,
	}.Build()
	File_VideoRCV_proto = out.File
	file_VideoRCV_proto_goTypes = nil
	file_VideoRCV_proto_depIdxs = nil
}

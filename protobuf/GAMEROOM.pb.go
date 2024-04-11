// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: GAMEROOM.proto

package protobuf

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

type GAMEROOM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roomid   *uint32 `protobuf:"varint,1,req,name=roomid" json:"roomid,omitempty"`
	MaxScore *uint32 `protobuf:"varint,2,req,name=max_score,json=maxScore" json:"max_score,omitempty"`
}

func (x *GAMEROOM) Reset() {
	*x = GAMEROOM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GAMEROOM_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GAMEROOM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GAMEROOM) ProtoMessage() {}

func (x *GAMEROOM) ProtoReflect() protoreflect.Message {
	mi := &file_GAMEROOM_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GAMEROOM.ProtoReflect.Descriptor instead.
func (*GAMEROOM) Descriptor() ([]byte, []int) {
	return file_GAMEROOM_proto_rawDescGZIP(), []int{0}
}

func (x *GAMEROOM) GetRoomid() uint32 {
	if x != nil && x.Roomid != nil {
		return *x.Roomid
	}
	return 0
}

func (x *GAMEROOM) GetMaxScore() uint32 {
	if x != nil && x.MaxScore != nil {
		return *x.MaxScore
	}
	return 0
}

var File_GAMEROOM_proto protoreflect.FileDescriptor

var file_GAMEROOM_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x47, 0x41, 0x4d, 0x45, 0x52, 0x4f, 0x4f, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x08, 0x47, 0x41, 0x4d,
	0x45, 0x52, 0x4f, 0x4f, 0x4d, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x08, 0x6d, 0x61, 0x78, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_GAMEROOM_proto_rawDescOnce sync.Once
	file_GAMEROOM_proto_rawDescData = file_GAMEROOM_proto_rawDesc
)

func file_GAMEROOM_proto_rawDescGZIP() []byte {
	file_GAMEROOM_proto_rawDescOnce.Do(func() {
		file_GAMEROOM_proto_rawDescData = protoimpl.X.CompressGZIP(file_GAMEROOM_proto_rawDescData)
	})
	return file_GAMEROOM_proto_rawDescData
}

var file_GAMEROOM_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GAMEROOM_proto_goTypes = []interface{}{
	(*GAMEROOM)(nil), // 0: belfast.GAMEROOM
}
var file_GAMEROOM_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GAMEROOM_proto_init() }
func file_GAMEROOM_proto_init() {
	if File_GAMEROOM_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GAMEROOM_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GAMEROOM); i {
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
			RawDescriptor: file_GAMEROOM_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GAMEROOM_proto_goTypes,
		DependencyIndexes: file_GAMEROOM_proto_depIdxs,
		MessageInfos:      file_GAMEROOM_proto_msgTypes,
	}.Build()
	File_GAMEROOM_proto = out.File
	file_GAMEROOM_proto_rawDesc = nil
	file_GAMEROOM_proto_goTypes = nil
	file_GAMEROOM_proto_depIdxs = nil
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.1
// source: char-vs-rune/char_vs_rune.proto

package char_vs_rune

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ToRuneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
}

func (x *ToRuneRequest) Reset() {
	*x = ToRuneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_char_vs_rune_char_vs_rune_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToRuneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToRuneRequest) ProtoMessage() {}

func (x *ToRuneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_char_vs_rune_char_vs_rune_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToRuneRequest.ProtoReflect.Descriptor instead.
func (*ToRuneRequest) Descriptor() ([]byte, []int) {
	return file_char_vs_rune_char_vs_rune_proto_rawDescGZIP(), []int{0}
}

func (x *ToRuneRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

type ToRuneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InRunes []uint32          `protobuf:"varint,1,rep,packed,name=in_runes,json=inRunes,proto3" json:"in_runes,omitempty"`
	InBytes []uint32          `protobuf:"varint,2,rep,packed,name=in_bytes,json=inBytes,proto3" json:"in_bytes,omitempty"`
	Mapped  map[string]uint32 `protobuf:"bytes,3,rep,name=mapped,proto3" json:"mapped,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Bytes   map[string]*Bytes `protobuf:"bytes,4,rep,name=bytes,proto3" json:"bytes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ToRuneResponse) Reset() {
	*x = ToRuneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_char_vs_rune_char_vs_rune_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToRuneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToRuneResponse) ProtoMessage() {}

func (x *ToRuneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_char_vs_rune_char_vs_rune_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToRuneResponse.ProtoReflect.Descriptor instead.
func (*ToRuneResponse) Descriptor() ([]byte, []int) {
	return file_char_vs_rune_char_vs_rune_proto_rawDescGZIP(), []int{1}
}

func (x *ToRuneResponse) GetInRunes() []uint32 {
	if x != nil {
		return x.InRunes
	}
	return nil
}

func (x *ToRuneResponse) GetInBytes() []uint32 {
	if x != nil {
		return x.InBytes
	}
	return nil
}

func (x *ToRuneResponse) GetMapped() map[string]uint32 {
	if x != nil {
		return x.Mapped
	}
	return nil
}

func (x *ToRuneResponse) GetBytes() map[string]*Bytes {
	if x != nil {
		return x.Bytes
	}
	return nil
}

type Bytes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []uint32 `protobuf:"varint,1,rep,packed,name=values,proto3" json:"values,omitempty"`
}

func (x *Bytes) Reset() {
	*x = Bytes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_char_vs_rune_char_vs_rune_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bytes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bytes) ProtoMessage() {}

func (x *Bytes) ProtoReflect() protoreflect.Message {
	mi := &file_char_vs_rune_char_vs_rune_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bytes.ProtoReflect.Descriptor instead.
func (*Bytes) Descriptor() ([]byte, []int) {
	return file_char_vs_rune_char_vs_rune_proto_rawDescGZIP(), []int{2}
}

func (x *Bytes) GetValues() []uint32 {
	if x != nil {
		return x.Values
	}
	return nil
}

type ToCharRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Runes []uint32 `protobuf:"varint,1,rep,packed,name=runes,proto3" json:"runes,omitempty"`
}

func (x *ToCharRequest) Reset() {
	*x = ToCharRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_char_vs_rune_char_vs_rune_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToCharRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToCharRequest) ProtoMessage() {}

func (x *ToCharRequest) ProtoReflect() protoreflect.Message {
	mi := &file_char_vs_rune_char_vs_rune_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToCharRequest.ProtoReflect.Descriptor instead.
func (*ToCharRequest) Descriptor() ([]byte, []int) {
	return file_char_vs_rune_char_vs_rune_proto_rawDescGZIP(), []int{3}
}

func (x *ToCharRequest) GetRunes() []uint32 {
	if x != nil {
		return x.Runes
	}
	return nil
}

type ToCharResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To string `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *ToCharResponse) Reset() {
	*x = ToCharResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_char_vs_rune_char_vs_rune_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToCharResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToCharResponse) ProtoMessage() {}

func (x *ToCharResponse) ProtoReflect() protoreflect.Message {
	mi := &file_char_vs_rune_char_vs_rune_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToCharResponse.ProtoReflect.Descriptor instead.
func (*ToCharResponse) Descriptor() ([]byte, []int) {
	return file_char_vs_rune_char_vs_rune_proto_rawDescGZIP(), []int{4}
}

func (x *ToCharResponse) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

var File_char_vs_rune_char_vs_rune_proto protoreflect.FileDescriptor

var file_char_vs_rune_char_vs_rune_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x68, 0x61, 0x72, 0x2d, 0x76, 0x73, 0x2d, 0x72, 0x75, 0x6e, 0x65, 0x2f, 0x63,
	0x68, 0x61, 0x72, 0x5f, 0x76, 0x73, 0x5f, 0x72, 0x75, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x23, 0x0a, 0x0d, 0x54, 0x6f, 0x52, 0x75, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x22, 0xaa, 0x02, 0x0a, 0x0e, 0x54, 0x6f, 0x52, 0x75, 0x6e,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6e, 0x5f,
	0x72, 0x75, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x69, 0x6e, 0x52,
	0x75, 0x6e, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6e, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x69, 0x6e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12,
	0x33, 0x0a, 0x06, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x54, 0x6f, 0x52, 0x75, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6d, 0x61,
	0x70, 0x70, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x54, 0x6f, 0x52, 0x75, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x64,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x40, 0x0a, 0x0a, 0x42, 0x79, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x1c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x06, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x1f, 0x0a, 0x05, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x22, 0x25, 0x0a, 0x0d, 0x54, 0x6f, 0x43, 0x68, 0x61, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x75, 0x6e, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x05, 0x72, 0x75, 0x6e, 0x65, 0x73, 0x22, 0x20, 0x0a, 0x0e, 0x54,
	0x6f, 0x43, 0x68, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x32, 0x62, 0x0a,
	0x0a, 0x43, 0x68, 0x61, 0x72, 0x56, 0x73, 0x52, 0x75, 0x6e, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x54,
	0x6f, 0x52, 0x75, 0x6e, 0x65, 0x12, 0x0e, 0x2e, 0x54, 0x6f, 0x52, 0x75, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x54, 0x6f, 0x52, 0x75, 0x6e, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x54, 0x6f, 0x43, 0x68, 0x61, 0x72,
	0x12, 0x0e, 0x2e, 0x54, 0x6f, 0x43, 0x68, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0f, 0x2e, 0x54, 0x6f, 0x43, 0x68, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x61, 0x6d, 0x61, 0x72, 0x61, 0x6b, 0x61, 0x75, 0x66, 0x6c, 0x65, 0x72, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x63, 0x68, 0x61, 0x72, 0x2d, 0x76, 0x73, 0x2d, 0x72, 0x75, 0x6e, 0x65, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x5f, 0x76, 0x73, 0x5f, 0x72,
	0x75, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_char_vs_rune_char_vs_rune_proto_rawDescOnce sync.Once
	file_char_vs_rune_char_vs_rune_proto_rawDescData = file_char_vs_rune_char_vs_rune_proto_rawDesc
)

func file_char_vs_rune_char_vs_rune_proto_rawDescGZIP() []byte {
	file_char_vs_rune_char_vs_rune_proto_rawDescOnce.Do(func() {
		file_char_vs_rune_char_vs_rune_proto_rawDescData = protoimpl.X.CompressGZIP(file_char_vs_rune_char_vs_rune_proto_rawDescData)
	})
	return file_char_vs_rune_char_vs_rune_proto_rawDescData
}

var file_char_vs_rune_char_vs_rune_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_char_vs_rune_char_vs_rune_proto_goTypes = []interface{}{
	(*ToRuneRequest)(nil),  // 0: ToRuneRequest
	(*ToRuneResponse)(nil), // 1: ToRuneResponse
	(*Bytes)(nil),          // 2: Bytes
	(*ToCharRequest)(nil),  // 3: ToCharRequest
	(*ToCharResponse)(nil), // 4: ToCharResponse
	nil,                    // 5: ToRuneResponse.MappedEntry
	nil,                    // 6: ToRuneResponse.BytesEntry
}
var file_char_vs_rune_char_vs_rune_proto_depIdxs = []int32{
	5, // 0: ToRuneResponse.mapped:type_name -> ToRuneResponse.MappedEntry
	6, // 1: ToRuneResponse.bytes:type_name -> ToRuneResponse.BytesEntry
	2, // 2: ToRuneResponse.BytesEntry.value:type_name -> Bytes
	0, // 3: CharVsRune.ToRune:input_type -> ToRuneRequest
	3, // 4: CharVsRune.ToChar:input_type -> ToCharRequest
	1, // 5: CharVsRune.ToRune:output_type -> ToRuneResponse
	4, // 6: CharVsRune.ToChar:output_type -> ToCharResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_char_vs_rune_char_vs_rune_proto_init() }
func file_char_vs_rune_char_vs_rune_proto_init() {
	if File_char_vs_rune_char_vs_rune_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_char_vs_rune_char_vs_rune_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToRuneRequest); i {
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
		file_char_vs_rune_char_vs_rune_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToRuneResponse); i {
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
		file_char_vs_rune_char_vs_rune_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bytes); i {
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
		file_char_vs_rune_char_vs_rune_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToCharRequest); i {
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
		file_char_vs_rune_char_vs_rune_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToCharResponse); i {
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
			RawDescriptor: file_char_vs_rune_char_vs_rune_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_char_vs_rune_char_vs_rune_proto_goTypes,
		DependencyIndexes: file_char_vs_rune_char_vs_rune_proto_depIdxs,
		MessageInfos:      file_char_vs_rune_char_vs_rune_proto_msgTypes,
	}.Build()
	File_char_vs_rune_char_vs_rune_proto = out.File
	file_char_vs_rune_char_vs_rune_proto_rawDesc = nil
	file_char_vs_rune_char_vs_rune_proto_goTypes = nil
	file_char_vs_rune_char_vs_rune_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CharVsRuneClient is the client API for CharVsRune service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CharVsRuneClient interface {
	ToRune(ctx context.Context, in *ToRuneRequest, opts ...grpc.CallOption) (*ToRuneResponse, error)
	ToChar(ctx context.Context, in *ToCharRequest, opts ...grpc.CallOption) (*ToCharResponse, error)
}

type charVsRuneClient struct {
	cc grpc.ClientConnInterface
}

func NewCharVsRuneClient(cc grpc.ClientConnInterface) CharVsRuneClient {
	return &charVsRuneClient{cc}
}

func (c *charVsRuneClient) ToRune(ctx context.Context, in *ToRuneRequest, opts ...grpc.CallOption) (*ToRuneResponse, error) {
	out := new(ToRuneResponse)
	err := c.cc.Invoke(ctx, "/CharVsRune/ToRune", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charVsRuneClient) ToChar(ctx context.Context, in *ToCharRequest, opts ...grpc.CallOption) (*ToCharResponse, error) {
	out := new(ToCharResponse)
	err := c.cc.Invoke(ctx, "/CharVsRune/ToChar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CharVsRuneServer is the server API for CharVsRune service.
type CharVsRuneServer interface {
	ToRune(context.Context, *ToRuneRequest) (*ToRuneResponse, error)
	ToChar(context.Context, *ToCharRequest) (*ToCharResponse, error)
}

// UnimplementedCharVsRuneServer can be embedded to have forward compatible implementations.
type UnimplementedCharVsRuneServer struct {
}

func (*UnimplementedCharVsRuneServer) ToRune(context.Context, *ToRuneRequest) (*ToRuneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToRune not implemented")
}
func (*UnimplementedCharVsRuneServer) ToChar(context.Context, *ToCharRequest) (*ToCharResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToChar not implemented")
}

func RegisterCharVsRuneServer(s *grpc.Server, srv CharVsRuneServer) {
	s.RegisterService(&_CharVsRune_serviceDesc, srv)
}

func _CharVsRune_ToRune_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToRuneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharVsRuneServer).ToRune(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CharVsRune/ToRune",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharVsRuneServer).ToRune(ctx, req.(*ToRuneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharVsRune_ToChar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToCharRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharVsRuneServer).ToChar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CharVsRune/ToChar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharVsRuneServer).ToChar(ctx, req.(*ToCharRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CharVsRune_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CharVsRune",
	HandlerType: (*CharVsRuneServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ToRune",
			Handler:    _CharVsRune_ToRune_Handler,
		},
		{
			MethodName: "ToChar",
			Handler:    _CharVsRune_ToChar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "char-vs-rune/char_vs_rune.proto",
}
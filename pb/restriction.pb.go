// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: restriction.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Restriction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RestrictionId     int32                  `protobuf:"varint,1,opt,name=restriction_id,json=restrictionId,proto3" json:"restriction_id,omitempty"`
	RestrictionNameSr string                 `protobuf:"bytes,2,opt,name=restriction_name_sr,json=restrictionNameSr,proto3" json:"restriction_name_sr,omitempty"`
	RestrictionNameEn string                 `protobuf:"bytes,3,opt,name=restriction_name_en,json=restrictionNameEn,proto3" json:"restriction_name_en,omitempty"`
	RestrictionNameBg string                 `protobuf:"bytes,4,opt,name=restriction_name_bg,json=restrictionNameBg,proto3" json:"restriction_name_bg,omitempty"`
	CreatedAt         *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt         *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Restriction) Reset() {
	*x = Restriction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restriction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Restriction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Restriction) ProtoMessage() {}

func (x *Restriction) ProtoReflect() protoreflect.Message {
	mi := &file_restriction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Restriction.ProtoReflect.Descriptor instead.
func (*Restriction) Descriptor() ([]byte, []int) {
	return file_restriction_proto_rawDescGZIP(), []int{0}
}

func (x *Restriction) GetRestrictionId() int32 {
	if x != nil {
		return x.RestrictionId
	}
	return 0
}

func (x *Restriction) GetRestrictionNameSr() string {
	if x != nil {
		return x.RestrictionNameSr
	}
	return ""
}

func (x *Restriction) GetRestrictionNameEn() string {
	if x != nil {
		return x.RestrictionNameEn
	}
	return ""
}

func (x *Restriction) GetRestrictionNameBg() string {
	if x != nil {
		return x.RestrictionNameBg
	}
	return ""
}

func (x *Restriction) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Restriction) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_restriction_proto protoreflect.FileDescriptor

var file_restriction_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x02, 0x0a, 0x0b, 0x52, 0x65, 0x73,
	0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x2e, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x5f, 0x73, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x72, 0x12,
	0x2e, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x5f, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x12,
	0x2e, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x5f, 0x62, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x67, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6a, 0x61, 0x6e, 0x61, 0x64, 0x6d, 0x69, 0x2f, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_restriction_proto_rawDescOnce sync.Once
	file_restriction_proto_rawDescData = file_restriction_proto_rawDesc
)

func file_restriction_proto_rawDescGZIP() []byte {
	file_restriction_proto_rawDescOnce.Do(func() {
		file_restriction_proto_rawDescData = protoimpl.X.CompressGZIP(file_restriction_proto_rawDescData)
	})
	return file_restriction_proto_rawDescData
}

var file_restriction_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_restriction_proto_goTypes = []any{
	(*Restriction)(nil),           // 0: pb.Restriction
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_restriction_proto_depIdxs = []int32{
	1, // 0: pb.Restriction.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: pb.Restriction.updated_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_restriction_proto_init() }
func file_restriction_proto_init() {
	if File_restriction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_restriction_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Restriction); i {
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
			RawDescriptor: file_restriction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_restriction_proto_goTypes,
		DependencyIndexes: file_restriction_proto_depIdxs,
		MessageInfos:      file_restriction_proto_msgTypes,
	}.Build()
	File_restriction_proto = out.File
	file_restriction_proto_rawDesc = nil
	file_restriction_proto_goTypes = nil
	file_restriction_proto_depIdxs = nil
}

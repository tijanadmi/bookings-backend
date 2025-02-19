// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: rpc_list_restrictions.proto

package pb

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

type ListRestrictionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`   // Limit restrikcija koje vraćamo
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"` // Offset za paginaciju
}

func (x *ListRestrictionsRequest) Reset() {
	*x = ListRestrictionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_list_restrictions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRestrictionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRestrictionsRequest) ProtoMessage() {}

func (x *ListRestrictionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_list_restrictions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRestrictionsRequest.ProtoReflect.Descriptor instead.
func (*ListRestrictionsRequest) Descriptor() ([]byte, []int) {
	return file_rpc_list_restrictions_proto_rawDescGZIP(), []int{0}
}

func (x *ListRestrictionsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListRestrictionsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListRestrictionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Restrictions []*Restriction `protobuf:"bytes,1,rep,name=restrictions,proto3" json:"restrictions,omitempty"` // Lista restrikcija
}

func (x *ListRestrictionsResponse) Reset() {
	*x = ListRestrictionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_list_restrictions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRestrictionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRestrictionsResponse) ProtoMessage() {}

func (x *ListRestrictionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_list_restrictions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRestrictionsResponse.ProtoReflect.Descriptor instead.
func (*ListRestrictionsResponse) Descriptor() ([]byte, []int) {
	return file_rpc_list_restrictions_proto_rawDescGZIP(), []int{1}
}

func (x *ListRestrictionsResponse) GetRestrictions() []*Restriction {
	if x != nil {
		return x.Restrictions
	}
	return nil
}

var File_rpc_list_restrictions_proto protoreflect.FileDescriptor

var file_rpc_list_restrictions_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x72,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x11, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x4f, 0x0a,
	0x18, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0c, 0x72, 0x65, 0x73,
	0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x2a,
	0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6a,
	0x61, 0x6e, 0x61, 0x64, 0x6d, 0x69, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x5f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_rpc_list_restrictions_proto_rawDescOnce sync.Once
	file_rpc_list_restrictions_proto_rawDescData = file_rpc_list_restrictions_proto_rawDesc
)

func file_rpc_list_restrictions_proto_rawDescGZIP() []byte {
	file_rpc_list_restrictions_proto_rawDescOnce.Do(func() {
		file_rpc_list_restrictions_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_list_restrictions_proto_rawDescData)
	})
	return file_rpc_list_restrictions_proto_rawDescData
}

var file_rpc_list_restrictions_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_list_restrictions_proto_goTypes = []any{
	(*ListRestrictionsRequest)(nil),  // 0: pb.ListRestrictionsRequest
	(*ListRestrictionsResponse)(nil), // 1: pb.ListRestrictionsResponse
	(*Restriction)(nil),              // 2: pb.Restriction
}
var file_rpc_list_restrictions_proto_depIdxs = []int32{
	2, // 0: pb.ListRestrictionsResponse.restrictions:type_name -> pb.Restriction
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_list_restrictions_proto_init() }
func file_rpc_list_restrictions_proto_init() {
	if File_rpc_list_restrictions_proto != nil {
		return
	}
	file_restriction_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_list_restrictions_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListRestrictionsRequest); i {
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
		file_rpc_list_restrictions_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListRestrictionsResponse); i {
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
			RawDescriptor: file_rpc_list_restrictions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_list_restrictions_proto_goTypes,
		DependencyIndexes: file_rpc_list_restrictions_proto_depIdxs,
		MessageInfos:      file_rpc_list_restrictions_proto_msgTypes,
	}.Build()
	File_rpc_list_restrictions_proto = out.File
	file_rpc_list_restrictions_proto_rawDesc = nil
	file_rpc_list_restrictions_proto_goTypes = nil
	file_rpc_list_restrictions_proto_depIdxs = nil
}

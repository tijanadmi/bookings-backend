// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: rpc_get_reservation.proto

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

type GetReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReservationId int32 `protobuf:"varint,1,opt,name=reservation_id,json=reservationId,proto3" json:"reservation_id,omitempty"`
}

func (x *GetReservationRequest) Reset() {
	*x = GetReservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_get_reservation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReservationRequest) ProtoMessage() {}

func (x *GetReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_get_reservation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReservationRequest.ProtoReflect.Descriptor instead.
func (*GetReservationRequest) Descriptor() ([]byte, []int) {
	return file_rpc_get_reservation_proto_rawDescGZIP(), []int{0}
}

func (x *GetReservationRequest) GetReservationId() int32 {
	if x != nil {
		return x.ReservationId
	}
	return 0
}

type GetReservationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reservation *ReservationAll `protobuf:"bytes,1,opt,name=reservation,proto3" json:"reservation,omitempty"`
}

func (x *GetReservationResponse) Reset() {
	*x = GetReservationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_get_reservation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReservationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReservationResponse) ProtoMessage() {}

func (x *GetReservationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_get_reservation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReservationResponse.ProtoReflect.Descriptor instead.
func (*GetReservationResponse) Descriptor() ([]byte, []int) {
	return file_rpc_get_reservation_proto_rawDescGZIP(), []int{1}
}

func (x *GetReservationResponse) GetReservation() *ReservationAll {
	if x != nil {
		return x.Reservation
	}
	return nil
}

var File_rpc_get_reservation_proto protoreflect.FileDescriptor

var file_rpc_get_reservation_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x15, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6c, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x6c, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6a, 0x61, 0x6e, 0x61, 0x64, 0x6d, 0x69, 0x2f, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_get_reservation_proto_rawDescOnce sync.Once
	file_rpc_get_reservation_proto_rawDescData = file_rpc_get_reservation_proto_rawDesc
)

func file_rpc_get_reservation_proto_rawDescGZIP() []byte {
	file_rpc_get_reservation_proto_rawDescOnce.Do(func() {
		file_rpc_get_reservation_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_get_reservation_proto_rawDescData)
	})
	return file_rpc_get_reservation_proto_rawDescData
}

var file_rpc_get_reservation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_get_reservation_proto_goTypes = []any{
	(*GetReservationRequest)(nil),  // 0: pb.GetReservationRequest
	(*GetReservationResponse)(nil), // 1: pb.GetReservationResponse
	(*ReservationAll)(nil),         // 2: pb.ReservationAll
}
var file_rpc_get_reservation_proto_depIdxs = []int32{
	2, // 0: pb.GetReservationResponse.reservation:type_name -> pb.ReservationAll
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_get_reservation_proto_init() }
func file_rpc_get_reservation_proto_init() {
	if File_rpc_get_reservation_proto != nil {
		return
	}
	file_reservation_all_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_get_reservation_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetReservationRequest); i {
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
		file_rpc_get_reservation_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetReservationResponse); i {
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
			RawDescriptor: file_rpc_get_reservation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_get_reservation_proto_goTypes,
		DependencyIndexes: file_rpc_get_reservation_proto_depIdxs,
		MessageInfos:      file_rpc_get_reservation_proto_msgTypes,
	}.Build()
	File_rpc_get_reservation_proto = out.File
	file_rpc_get_reservation_proto_rawDesc = nil
	file_rpc_get_reservation_proto_goTypes = nil
	file_rpc_get_reservation_proto_depIdxs = nil
}

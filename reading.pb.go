// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: reading.proto

package main

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type TemperatureReading struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	MachineId   string                 `protobuf:"bytes,3,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	Temperature float64                `protobuf:"fixed64,4,opt,name=temperature,proto3" json:"temperature,omitempty"`
}

func (x *TemperatureReading) Reset() {
	*x = TemperatureReading{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reading_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemperatureReading) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemperatureReading) ProtoMessage() {}

func (x *TemperatureReading) ProtoReflect() protoreflect.Message {
	mi := &file_reading_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemperatureReading.ProtoReflect.Descriptor instead.
func (*TemperatureReading) Descriptor() ([]byte, []int) {
	return file_reading_proto_rawDescGZIP(), []int{0}
}

func (x *TemperatureReading) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *TemperatureReading) GetMachineId() string {
	if x != nil {
		return x.MachineId
	}
	return ""
}

func (x *TemperatureReading) GetTemperature() float64 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

var File_reading_proto protoreflect.FileDescriptor

var file_reading_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8f, 0x01, 0x0a, 0x12, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reading_proto_rawDescOnce sync.Once
	file_reading_proto_rawDescData = file_reading_proto_rawDesc
)

func file_reading_proto_rawDescGZIP() []byte {
	file_reading_proto_rawDescOnce.Do(func() {
		file_reading_proto_rawDescData = protoimpl.X.CompressGZIP(file_reading_proto_rawDescData)
	})
	return file_reading_proto_rawDescData
}

var file_reading_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_reading_proto_goTypes = []interface{}{
	(*TemperatureReading)(nil),    // 0: TemperatureReading
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_reading_proto_depIdxs = []int32{
	1, // 0: TemperatureReading.timestamp:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_reading_proto_init() }
func file_reading_proto_init() {
	if File_reading_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reading_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemperatureReading); i {
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
			RawDescriptor: file_reading_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_reading_proto_goTypes,
		DependencyIndexes: file_reading_proto_depIdxs,
		MessageInfos:      file_reading_proto_msgTypes,
	}.Build()
	File_reading_proto = out.File
	file_reading_proto_rawDesc = nil
	file_reading_proto_goTypes = nil
	file_reading_proto_depIdxs = nil
}

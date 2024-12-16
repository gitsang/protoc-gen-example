// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.1
// source: protoc_gen_example/options/v1/annotations.proto

package options

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_protoc_gen_example_options_v1_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         63550,
		Name:          "protoc_gen_example.options.v1.message_options",
		Tag:           "bytes,63550,opt,name=message_options",
		Filename:      "protoc_gen_example/options/v1/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         63550,
		Name:          "protoc_gen_example.options.v1.file_options",
		Tag:           "bytes,63550,opt,name=file_options",
		Filename:      "protoc_gen_example/options/v1/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         63550,
		Name:          "protoc_gen_example.options.v1.service_options",
		Tag:           "bytes,63550,opt,name=service_options",
		Filename:      "protoc_gen_example/options/v1/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         63550,
		Name:          "protoc_gen_example.options.v1.field_options",
		Tag:           "bytes,63550,opt,name=field_options",
		Filename:      "protoc_gen_example/options/v1/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         63550,
		Name:          "protoc_gen_example.options.v1.method_options",
		Tag:           "bytes,63550,opt,name=method_options",
		Filename:      "protoc_gen_example/options/v1/annotations.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional string message_options = 63550;
	E_MessageOptions = &file_protoc_gen_example_options_v1_annotations_proto_extTypes[0]
)

// Extension fields to descriptorpb.FileOptions.
var (
	// optional string file_options = 63550;
	E_FileOptions = &file_protoc_gen_example_options_v1_annotations_proto_extTypes[1]
)

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional string service_options = 63550;
	E_ServiceOptions = &file_protoc_gen_example_options_v1_annotations_proto_extTypes[2]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional string field_options = 63550;
	E_FieldOptions = &file_protoc_gen_example_options_v1_annotations_proto_extTypes[3]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional string method_options = 63550;
	E_MethodOptions = &file_protoc_gen_example_options_v1_annotations_proto_extTypes[4]
)

var File_protoc_gen_example_options_v1_annotations_proto protoreflect.FileDescriptor

var file_protoc_gen_example_options_v1_annotations_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3a, 0x4a, 0x0a, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbe, 0xf0, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x41,
	0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbe, 0xf0, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x3a, 0x4a, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbe, 0xf0, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x44, 0x0a,
	0x0d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbe, 0xf0,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x3a, 0x47, 0x0a, 0x0e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbe, 0xf0, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x53, 0x5a, 0x51,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x73, 0x61,
	0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_protoc_gen_example_options_v1_annotations_proto_goTypes = []interface{}{
	(*descriptorpb.MessageOptions)(nil), // 0: google.protobuf.MessageOptions
	(*descriptorpb.FileOptions)(nil),    // 1: google.protobuf.FileOptions
	(*descriptorpb.ServiceOptions)(nil), // 2: google.protobuf.ServiceOptions
	(*descriptorpb.FieldOptions)(nil),   // 3: google.protobuf.FieldOptions
	(*descriptorpb.MethodOptions)(nil),  // 4: google.protobuf.MethodOptions
}
var file_protoc_gen_example_options_v1_annotations_proto_depIdxs = []int32{
	0, // 0: protoc_gen_example.options.v1.message_options:extendee -> google.protobuf.MessageOptions
	1, // 1: protoc_gen_example.options.v1.file_options:extendee -> google.protobuf.FileOptions
	2, // 2: protoc_gen_example.options.v1.service_options:extendee -> google.protobuf.ServiceOptions
	3, // 3: protoc_gen_example.options.v1.field_options:extendee -> google.protobuf.FieldOptions
	4, // 4: protoc_gen_example.options.v1.method_options:extendee -> google.protobuf.MethodOptions
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	0, // [0:5] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protoc_gen_example_options_v1_annotations_proto_init() }
func file_protoc_gen_example_options_v1_annotations_proto_init() {
	if File_protoc_gen_example_options_v1_annotations_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protoc_gen_example_options_v1_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 5,
			NumServices:   0,
		},
		GoTypes:           file_protoc_gen_example_options_v1_annotations_proto_goTypes,
		DependencyIndexes: file_protoc_gen_example_options_v1_annotations_proto_depIdxs,
		ExtensionInfos:    file_protoc_gen_example_options_v1_annotations_proto_extTypes,
	}.Build()
	File_protoc_gen_example_options_v1_annotations_proto = out.File
	file_protoc_gen_example_options_v1_annotations_proto_rawDesc = nil
	file_protoc_gen_example_options_v1_annotations_proto_goTypes = nil
	file_protoc_gen_example_options_v1_annotations_proto_depIdxs = nil
}
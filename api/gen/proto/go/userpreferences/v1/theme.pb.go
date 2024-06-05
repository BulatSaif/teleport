// Copyright 2023 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: teleport/userpreferences/v1/theme.proto

package userpreferencesv1

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

// Theme is a frontend theme.
type Theme int32

const (
	Theme_THEME_UNSPECIFIED Theme = 0
	// THEME_LIGHT is the light theme.
	Theme_THEME_LIGHT Theme = 1
	// THEME_DARK is the dark theme.
	Theme_THEME_DARK Theme = 2
)

// Enum value maps for Theme.
var (
	Theme_name = map[int32]string{
		0: "THEME_UNSPECIFIED",
		1: "THEME_LIGHT",
		2: "THEME_DARK",
	}
	Theme_value = map[string]int32{
		"THEME_UNSPECIFIED": 0,
		"THEME_LIGHT":       1,
		"THEME_DARK":        2,
	}
)

func (x Theme) Enum() *Theme {
	p := new(Theme)
	*p = x
	return p
}

func (x Theme) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Theme) Descriptor() protoreflect.EnumDescriptor {
	return file_teleport_userpreferences_v1_theme_proto_enumTypes[0].Descriptor()
}

func (Theme) Type() protoreflect.EnumType {
	return &file_teleport_userpreferences_v1_theme_proto_enumTypes[0]
}

func (x Theme) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Theme.Descriptor instead.
func (Theme) EnumDescriptor() ([]byte, []int) {
	return file_teleport_userpreferences_v1_theme_proto_rawDescGZIP(), []int{0}
}

var File_teleport_userpreferences_v1_theme_proto protoreflect.FileDescriptor

var file_teleport_userpreferences_v1_theme_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x68,
	0x65, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2a, 0x3f, 0x0a, 0x05, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x12,
	0x15, 0x0a, 0x11, 0x54, 0x48, 0x45, 0x4d, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x48, 0x45, 0x4d, 0x45, 0x5f,
	0x4c, 0x49, 0x47, 0x48, 0x54, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x48, 0x45, 0x4d, 0x45,
	0x5f, 0x44, 0x41, 0x52, 0x4b, 0x10, 0x02, 0x42, 0x59, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_userpreferences_v1_theme_proto_rawDescOnce sync.Once
	file_teleport_userpreferences_v1_theme_proto_rawDescData = file_teleport_userpreferences_v1_theme_proto_rawDesc
)

func file_teleport_userpreferences_v1_theme_proto_rawDescGZIP() []byte {
	file_teleport_userpreferences_v1_theme_proto_rawDescOnce.Do(func() {
		file_teleport_userpreferences_v1_theme_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_userpreferences_v1_theme_proto_rawDescData)
	})
	return file_teleport_userpreferences_v1_theme_proto_rawDescData
}

var file_teleport_userpreferences_v1_theme_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_teleport_userpreferences_v1_theme_proto_goTypes = []interface{}{
	(Theme)(0), // 0: teleport.userpreferences.v1.Theme
}
var file_teleport_userpreferences_v1_theme_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_teleport_userpreferences_v1_theme_proto_init() }
func file_teleport_userpreferences_v1_theme_proto_init() {
	if File_teleport_userpreferences_v1_theme_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_userpreferences_v1_theme_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_userpreferences_v1_theme_proto_goTypes,
		DependencyIndexes: file_teleport_userpreferences_v1_theme_proto_depIdxs,
		EnumInfos:         file_teleport_userpreferences_v1_theme_proto_enumTypes,
	}.Build()
	File_teleport_userpreferences_v1_theme_proto = out.File
	file_teleport_userpreferences_v1_theme_proto_rawDesc = nil
	file_teleport_userpreferences_v1_theme_proto_goTypes = nil
	file_teleport_userpreferences_v1_theme_proto_depIdxs = nil
}

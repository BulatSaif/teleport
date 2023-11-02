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
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: teleport/secreports/v1/secreports.proto

package secreportsv1

import (
	v1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/header/v1"
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

// AuditQuery is audit query resource.
type AuditQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// header is the header for //the resource.
	Header *v1.ResourceHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// spec is audit query spec.
	Spec *AuditQuerySpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *AuditQuery) Reset() {
	*x = AuditQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_secreports_v1_secreports_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditQuery) ProtoMessage() {}

func (x *AuditQuery) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_secreports_v1_secreports_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditQuery.ProtoReflect.Descriptor instead.
func (*AuditQuery) Descriptor() ([]byte, []int) {
	return file_teleport_secreports_v1_secreports_proto_rawDescGZIP(), []int{0}
}

func (x *AuditQuery) GetHeader() *v1.ResourceHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *AuditQuery) GetSpec() *AuditQuerySpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

// AuditQuerySpec is audit query spec.
type AuditQuerySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the name of the audit query.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// title is the title of the audit query.
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// query is the SQL Query for the audit query.
	Query string `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
	// description is the description of the audit query.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *AuditQuerySpec) Reset() {
	*x = AuditQuerySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_secreports_v1_secreports_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditQuerySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditQuerySpec) ProtoMessage() {}

func (x *AuditQuerySpec) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_secreports_v1_secreports_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditQuerySpec.ProtoReflect.Descriptor instead.
func (*AuditQuerySpec) Descriptor() ([]byte, []int) {
	return file_teleport_secreports_v1_secreports_proto_rawDescGZIP(), []int{1}
}

func (x *AuditQuerySpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuditQuerySpec) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AuditQuerySpec) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *AuditQuerySpec) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// Report is security report resource.
type Report struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// header is the header for the resource.
	Header *v1.ResourceHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// spec is the security report spec.
	Spec *ReportSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *Report) Reset() {
	*x = Report{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_secreports_v1_secreports_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_secreports_v1_secreports_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report.ProtoReflect.Descriptor instead.
func (*Report) Descriptor() ([]byte, []int) {
	return file_teleport_secreports_v1_secreports_proto_rawDescGZIP(), []int{2}
}

func (x *Report) GetHeader() *v1.ResourceHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Report) GetSpec() *ReportSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

// ReportSpec is security report spec.
type ReportSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the name of the security report.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// title is the title of the security report.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// description is the description of the security report
	AuditQueries []*AuditQuerySpec `protobuf:"bytes,3,rep,name=audit_queries,json=auditQueries,proto3" json:"audit_queries,omitempty"`
	// title is the title of the security report.
	Title string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	// version is the version of the security report.
	Version string `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ReportSpec) Reset() {
	*x = ReportSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_secreports_v1_secreports_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportSpec) ProtoMessage() {}

func (x *ReportSpec) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_secreports_v1_secreports_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportSpec.ProtoReflect.Descriptor instead.
func (*ReportSpec) Descriptor() ([]byte, []int) {
	return file_teleport_secreports_v1_secreports_proto_rawDescGZIP(), []int{3}
}

func (x *ReportSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReportSpec) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ReportSpec) GetAuditQueries() []*AuditQuerySpec {
	if x != nil {
		return x.AuditQueries
	}
	return nil
}

func (x *ReportSpec) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ReportSpec) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// ReportState is security report state resource.
type ReportState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// header is the header for the resource.
	Header *v1.ResourceHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// spec is the security report state spec.
	Spec *ReportStateSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *ReportState) Reset() {
	*x = ReportState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_secreports_v1_secreports_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportState) ProtoMessage() {}

func (x *ReportState) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_secreports_v1_secreports_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportState.ProtoReflect.Descriptor instead.
func (*ReportState) Descriptor() ([]byte, []int) {
	return file_teleport_secreports_v1_secreports_proto_rawDescGZIP(), []int{4}
}

func (x *ReportState) GetHeader() *v1.ResourceHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ReportState) GetSpec() *ReportStateSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

// ReportStateSpec is security report state spec.
type ReportStateSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// state is the state of the security report.
	State string `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	// updated_at is the time when the security report state was updated.
	UpdatedAt string `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *ReportStateSpec) Reset() {
	*x = ReportStateSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_secreports_v1_secreports_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportStateSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportStateSpec) ProtoMessage() {}

func (x *ReportStateSpec) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_secreports_v1_secreports_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportStateSpec.ProtoReflect.Descriptor instead.
func (*ReportStateSpec) Descriptor() ([]byte, []int) {
	return file_teleport_secreports_v1_secreports_proto_rawDescGZIP(), []int{5}
}

func (x *ReportStateSpec) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *ReportStateSpec) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_teleport_secreports_v1_secreports_proto protoreflect.FileDescriptor

var file_teleport_secreports_v1_secreports_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x27, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x0a, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3a, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x22, 0x72, 0x0a, 0x0e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7c, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x3a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x22, 0xbf, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x51, 0x75,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x86, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x3b, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x46,
	0x0a, 0x0f, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x58, 0x5a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_secreports_v1_secreports_proto_rawDescOnce sync.Once
	file_teleport_secreports_v1_secreports_proto_rawDescData = file_teleport_secreports_v1_secreports_proto_rawDesc
)

func file_teleport_secreports_v1_secreports_proto_rawDescGZIP() []byte {
	file_teleport_secreports_v1_secreports_proto_rawDescOnce.Do(func() {
		file_teleport_secreports_v1_secreports_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_secreports_v1_secreports_proto_rawDescData)
	})
	return file_teleport_secreports_v1_secreports_proto_rawDescData
}

var file_teleport_secreports_v1_secreports_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_teleport_secreports_v1_secreports_proto_goTypes = []interface{}{
	(*AuditQuery)(nil),        // 0: teleport.secreports.v1.AuditQuery
	(*AuditQuerySpec)(nil),    // 1: teleport.secreports.v1.AuditQuerySpec
	(*Report)(nil),            // 2: teleport.secreports.v1.Report
	(*ReportSpec)(nil),        // 3: teleport.secreports.v1.ReportSpec
	(*ReportState)(nil),       // 4: teleport.secreports.v1.ReportState
	(*ReportStateSpec)(nil),   // 5: teleport.secreports.v1.ReportStateSpec
	(*v1.ResourceHeader)(nil), // 6: teleport.header.v1.ResourceHeader
}
var file_teleport_secreports_v1_secreports_proto_depIdxs = []int32{
	6, // 0: teleport.secreports.v1.AuditQuery.header:type_name -> teleport.header.v1.ResourceHeader
	1, // 1: teleport.secreports.v1.AuditQuery.spec:type_name -> teleport.secreports.v1.AuditQuerySpec
	6, // 2: teleport.secreports.v1.Report.header:type_name -> teleport.header.v1.ResourceHeader
	3, // 3: teleport.secreports.v1.Report.spec:type_name -> teleport.secreports.v1.ReportSpec
	1, // 4: teleport.secreports.v1.ReportSpec.audit_queries:type_name -> teleport.secreports.v1.AuditQuerySpec
	6, // 5: teleport.secreports.v1.ReportState.header:type_name -> teleport.header.v1.ResourceHeader
	5, // 6: teleport.secreports.v1.ReportState.spec:type_name -> teleport.secreports.v1.ReportStateSpec
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_teleport_secreports_v1_secreports_proto_init() }
func file_teleport_secreports_v1_secreports_proto_init() {
	if File_teleport_secreports_v1_secreports_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teleport_secreports_v1_secreports_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditQuery); i {
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
		file_teleport_secreports_v1_secreports_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditQuerySpec); i {
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
		file_teleport_secreports_v1_secreports_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Report); i {
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
		file_teleport_secreports_v1_secreports_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportSpec); i {
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
		file_teleport_secreports_v1_secreports_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportState); i {
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
		file_teleport_secreports_v1_secreports_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportStateSpec); i {
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
			RawDescriptor: file_teleport_secreports_v1_secreports_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_secreports_v1_secreports_proto_goTypes,
		DependencyIndexes: file_teleport_secreports_v1_secreports_proto_depIdxs,
		MessageInfos:      file_teleport_secreports_v1_secreports_proto_msgTypes,
	}.Build()
	File_teleport_secreports_v1_secreports_proto = out.File
	file_teleport_secreports_v1_secreports_proto_rawDesc = nil
	file_teleport_secreports_v1_secreports_proto_goTypes = nil
	file_teleport_secreports_v1_secreports_proto_depIdxs = nil
}

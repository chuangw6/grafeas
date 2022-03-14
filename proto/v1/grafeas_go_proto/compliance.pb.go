// Copyright 2021 The Grafeas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.13.0
// source: proto/v1/compliance.proto

package grafeas_go_proto

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

type ComplianceNote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The title that identifies this compliance check.
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// A description about this compliance check.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The OS and config versions the benchmark applies to.
	Version []*ComplianceVersion `protobuf:"bytes,3,rep,name=version,proto3" json:"version,omitempty"`
	// A rationale for the existence of this compliance check.
	Rationale string `protobuf:"bytes,4,opt,name=rationale,proto3" json:"rationale,omitempty"`
	// A description of remediation steps if the compliance check fails.
	Remediation string `protobuf:"bytes,5,opt,name=remediation,proto3" json:"remediation,omitempty"`
	// Types that are assignable to ComplianceType:
	//	*ComplianceNote_CisBenchmark_
	ComplianceType isComplianceNote_ComplianceType `protobuf_oneof:"compliance_type"`
	// Serialized scan instructions with a predefined format.
	ScanInstructions []byte `protobuf:"bytes,7,opt,name=scan_instructions,json=scanInstructions,proto3" json:"scan_instructions,omitempty"`
}

func (x *ComplianceNote) Reset() {
	*x = ComplianceNote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_compliance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplianceNote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplianceNote) ProtoMessage() {}

func (x *ComplianceNote) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_compliance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplianceNote.ProtoReflect.Descriptor instead.
func (*ComplianceNote) Descriptor() ([]byte, []int) {
	return file_proto_v1_compliance_proto_rawDescGZIP(), []int{0}
}

func (x *ComplianceNote) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ComplianceNote) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ComplianceNote) GetVersion() []*ComplianceVersion {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *ComplianceNote) GetRationale() string {
	if x != nil {
		return x.Rationale
	}
	return ""
}

func (x *ComplianceNote) GetRemediation() string {
	if x != nil {
		return x.Remediation
	}
	return ""
}

func (m *ComplianceNote) GetComplianceType() isComplianceNote_ComplianceType {
	if m != nil {
		return m.ComplianceType
	}
	return nil
}

func (x *ComplianceNote) GetCisBenchmark() *ComplianceNote_CisBenchmark {
	if x, ok := x.GetComplianceType().(*ComplianceNote_CisBenchmark_); ok {
		return x.CisBenchmark
	}
	return nil
}

func (x *ComplianceNote) GetScanInstructions() []byte {
	if x != nil {
		return x.ScanInstructions
	}
	return nil
}

type isComplianceNote_ComplianceType interface {
	isComplianceNote_ComplianceType()
}

type ComplianceNote_CisBenchmark_ struct {
	CisBenchmark *ComplianceNote_CisBenchmark `protobuf:"bytes,6,opt,name=cis_benchmark,json=cisBenchmark,proto3,oneof"`
}

func (*ComplianceNote_CisBenchmark_) isComplianceNote_ComplianceType() {}

// Describes the CIS benchmark version that is applicable to a given OS and
// os version.
type ComplianceVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The CPE URI (https://cpe.mitre.org/specification/) this benchmark is
	// applicable to.
	CpeUri string `protobuf:"bytes,1,opt,name=cpe_uri,json=cpeUri,proto3" json:"cpe_uri,omitempty"`
	// The version of the benchmark. This is set to the version of the OS-specific
	// CIS document the benchmark is defined in.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ComplianceVersion) Reset() {
	*x = ComplianceVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_compliance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplianceVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplianceVersion) ProtoMessage() {}

func (x *ComplianceVersion) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_compliance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplianceVersion.ProtoReflect.Descriptor instead.
func (*ComplianceVersion) Descriptor() ([]byte, []int) {
	return file_proto_v1_compliance_proto_rawDescGZIP(), []int{1}
}

func (x *ComplianceVersion) GetCpeUri() string {
	if x != nil {
		return x.CpeUri
	}
	return ""
}

func (x *ComplianceVersion) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// An indication that the compliance checks in the associated ComplianceNote
// were not satisfied for particular resources or a specified reason.
type ComplianceOccurrence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NonCompliantFiles   []*NonCompliantFile `protobuf:"bytes,2,rep,name=non_compliant_files,json=nonCompliantFiles,proto3" json:"non_compliant_files,omitempty"`
	NonComplianceReason string              `protobuf:"bytes,3,opt,name=non_compliance_reason,json=nonComplianceReason,proto3" json:"non_compliance_reason,omitempty"`
}

func (x *ComplianceOccurrence) Reset() {
	*x = ComplianceOccurrence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_compliance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplianceOccurrence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplianceOccurrence) ProtoMessage() {}

func (x *ComplianceOccurrence) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_compliance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplianceOccurrence.ProtoReflect.Descriptor instead.
func (*ComplianceOccurrence) Descriptor() ([]byte, []int) {
	return file_proto_v1_compliance_proto_rawDescGZIP(), []int{2}
}

func (x *ComplianceOccurrence) GetNonCompliantFiles() []*NonCompliantFile {
	if x != nil {
		return x.NonCompliantFiles
	}
	return nil
}

func (x *ComplianceOccurrence) GetNonComplianceReason() string {
	if x != nil {
		return x.NonComplianceReason
	}
	return ""
}

// Details about files that caused a compliance check to fail.
type NonCompliantFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Empty if `display_command` is set.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Command to display the non-compliant files.
	DisplayCommand string `protobuf:"bytes,2,opt,name=display_command,json=displayCommand,proto3" json:"display_command,omitempty"`
	// Explains why a file is non compliant for a CIS check.
	Reason string `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *NonCompliantFile) Reset() {
	*x = NonCompliantFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_compliance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NonCompliantFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NonCompliantFile) ProtoMessage() {}

func (x *NonCompliantFile) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_compliance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NonCompliantFile.ProtoReflect.Descriptor instead.
func (*NonCompliantFile) Descriptor() ([]byte, []int) {
	return file_proto_v1_compliance_proto_rawDescGZIP(), []int{3}
}

func (x *NonCompliantFile) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *NonCompliantFile) GetDisplayCommand() string {
	if x != nil {
		return x.DisplayCommand
	}
	return ""
}

func (x *NonCompliantFile) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

// A compliance check that is a CIS benchmark.
type ComplianceNote_CisBenchmark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfileLevel int32    `protobuf:"varint,1,opt,name=profile_level,json=profileLevel,proto3" json:"profile_level,omitempty"`
	Severity     Severity `protobuf:"varint,2,opt,name=severity,proto3,enum=grafeas.v1.Severity" json:"severity,omitempty"`
}

func (x *ComplianceNote_CisBenchmark) Reset() {
	*x = ComplianceNote_CisBenchmark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_compliance_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplianceNote_CisBenchmark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplianceNote_CisBenchmark) ProtoMessage() {}

func (x *ComplianceNote_CisBenchmark) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_compliance_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplianceNote_CisBenchmark.ProtoReflect.Descriptor instead.
func (*ComplianceNote_CisBenchmark) Descriptor() ([]byte, []int) {
	return file_proto_v1_compliance_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ComplianceNote_CisBenchmark) GetProfileLevel() int32 {
	if x != nil {
		return x.ProfileLevel
	}
	return 0
}

func (x *ComplianceNote_CisBenchmark) GetSeverity() Severity {
	if x != nil {
		return x.Severity
	}
	return Severity_SEVERITY_UNSPECIFIED
}

var File_proto_v1_compliance_proto protoreflect.FileDescriptor

var file_proto_v1_compliance_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x72, 0x61,
	0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb8, 0x03, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x4e,
	0x6f, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67,
	0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69,
	0x61, 0x6e, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4e, 0x0a, 0x0d, 0x63, 0x69, 0x73, 0x5f, 0x62, 0x65, 0x6e, 0x63,
	0x68, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x72,
	0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61,
	0x6e, 0x63, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x2e, 0x43, 0x69, 0x73, 0x42, 0x65, 0x6e, 0x63, 0x68,
	0x6d, 0x61, 0x72, 0x6b, 0x48, 0x00, 0x52, 0x0c, 0x63, 0x69, 0x73, 0x42, 0x65, 0x6e, 0x63, 0x68,
	0x6d, 0x61, 0x72, 0x6b, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x69, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x10, 0x73, 0x63, 0x61, 0x6e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x65, 0x0a, 0x0c, 0x43, 0x69, 0x73, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72,
	0x6b, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65,
	0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08,
	0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x42, 0x11, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x46, 0x0a, 0x11, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x17, 0x0a, 0x07, 0x63, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x70, 0x65, 0x55, 0x72, 0x69, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x98, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e,
	0x63, 0x65, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x13,
	0x6e, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x72, 0x61, 0x66,
	0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69,
	0x61, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x11, 0x6e, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x69, 0x61, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x6e, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6e, 0x6f, 0x6e, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x67,
	0x0a, 0x10, 0x4e, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x4d, 0x0a, 0x0d, 0x69, 0x6f, 0x2e, 0x67, 0x72,
	0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2f, 0x67,
	0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xa2, 0x02, 0x03, 0x47, 0x52, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v1_compliance_proto_rawDescOnce sync.Once
	file_proto_v1_compliance_proto_rawDescData = file_proto_v1_compliance_proto_rawDesc
)

func file_proto_v1_compliance_proto_rawDescGZIP() []byte {
	file_proto_v1_compliance_proto_rawDescOnce.Do(func() {
		file_proto_v1_compliance_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_compliance_proto_rawDescData)
	})
	return file_proto_v1_compliance_proto_rawDescData
}

var file_proto_v1_compliance_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_v1_compliance_proto_goTypes = []interface{}{
	(*ComplianceNote)(nil),              // 0: grafeas.v1.ComplianceNote
	(*ComplianceVersion)(nil),           // 1: grafeas.v1.ComplianceVersion
	(*ComplianceOccurrence)(nil),        // 2: grafeas.v1.ComplianceOccurrence
	(*NonCompliantFile)(nil),            // 3: grafeas.v1.NonCompliantFile
	(*ComplianceNote_CisBenchmark)(nil), // 4: grafeas.v1.ComplianceNote.CisBenchmark
	(Severity)(0),                       // 5: grafeas.v1.Severity
}
var file_proto_v1_compliance_proto_depIdxs = []int32{
	1, // 0: grafeas.v1.ComplianceNote.version:type_name -> grafeas.v1.ComplianceVersion
	4, // 1: grafeas.v1.ComplianceNote.cis_benchmark:type_name -> grafeas.v1.ComplianceNote.CisBenchmark
	3, // 2: grafeas.v1.ComplianceOccurrence.non_compliant_files:type_name -> grafeas.v1.NonCompliantFile
	5, // 3: grafeas.v1.ComplianceNote.CisBenchmark.severity:type_name -> grafeas.v1.Severity
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_v1_compliance_proto_init() }
func file_proto_v1_compliance_proto_init() {
	if File_proto_v1_compliance_proto != nil {
		return
	}
	file_proto_v1_severity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_v1_compliance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplianceNote); i {
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
		file_proto_v1_compliance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplianceVersion); i {
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
		file_proto_v1_compliance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplianceOccurrence); i {
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
		file_proto_v1_compliance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NonCompliantFile); i {
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
		file_proto_v1_compliance_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplianceNote_CisBenchmark); i {
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
	file_proto_v1_compliance_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ComplianceNote_CisBenchmark_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_v1_compliance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_v1_compliance_proto_goTypes,
		DependencyIndexes: file_proto_v1_compliance_proto_depIdxs,
		MessageInfos:      file_proto_v1_compliance_proto_msgTypes,
	}.Build()
	File_proto_v1_compliance_proto = out.File
	file_proto_v1_compliance_proto_rawDesc = nil
	file_proto_v1_compliance_proto_goTypes = nil
	file_proto_v1_compliance_proto_depIdxs = nil
}

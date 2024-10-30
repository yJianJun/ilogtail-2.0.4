//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: common/trace-common.proto

package agent

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

type SpanType int32

const (
	SpanType_Entry SpanType = 0
	SpanType_Exit  SpanType = 1
	SpanType_Local SpanType = 2
)

// Enum value maps for SpanType.
var (
	SpanType_name = map[int32]string{
		0: "Entry",
		1: "Exit",
		2: "Local",
	}
	SpanType_value = map[string]int32{
		"Entry": 0,
		"Exit":  1,
		"Local": 2,
	}
)

func (x SpanType) Enum() *SpanType {
	p := new(SpanType)
	*p = x
	return p
}

func (x SpanType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SpanType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_trace_common_proto_enumTypes[0].Descriptor()
}

func (SpanType) Type() protoreflect.EnumType {
	return &file_common_trace_common_proto_enumTypes[0]
}

func (x SpanType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SpanType.Descriptor instead.
func (SpanType) EnumDescriptor() ([]byte, []int) {
	return file_common_trace_common_proto_rawDescGZIP(), []int{0}
}

type RefType int32

const (
	RefType_CrossProcess RefType = 0
	RefType_CrossThread  RefType = 1
)

// Enum value maps for RefType.
var (
	RefType_name = map[int32]string{
		0: "CrossProcess",
		1: "CrossThread",
	}
	RefType_value = map[string]int32{
		"CrossProcess": 0,
		"CrossThread":  1,
	}
)

func (x RefType) Enum() *RefType {
	p := new(RefType)
	*p = x
	return p
}

func (x RefType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RefType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_trace_common_proto_enumTypes[1].Descriptor()
}

func (RefType) Type() protoreflect.EnumType {
	return &file_common_trace_common_proto_enumTypes[1]
}

func (x RefType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RefType.Descriptor instead.
func (RefType) EnumDescriptor() ([]byte, []int) {
	return file_common_trace_common_proto_rawDescGZIP(), []int{1}
}

type SpanLayer int32

const (
	SpanLayer_Unknown      SpanLayer = 0
	SpanLayer_Database     SpanLayer = 1
	SpanLayer_RPCFramework SpanLayer = 2
	SpanLayer_Http         SpanLayer = 3
	SpanLayer_MQ           SpanLayer = 4
	SpanLayer_Cache        SpanLayer = 5
)

// Enum value maps for SpanLayer.
var (
	SpanLayer_name = map[int32]string{
		0: "Unknown",
		1: "Database",
		2: "RPCFramework",
		3: "Http",
		4: "MQ",
		5: "Cache",
	}
	SpanLayer_value = map[string]int32{
		"Unknown":      0,
		"Database":     1,
		"RPCFramework": 2,
		"Http":         3,
		"MQ":           4,
		"Cache":        5,
	}
)

func (x SpanLayer) Enum() *SpanLayer {
	p := new(SpanLayer)
	*p = x
	return p
}

func (x SpanLayer) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SpanLayer) Descriptor() protoreflect.EnumDescriptor {
	return file_common_trace_common_proto_enumTypes[2].Descriptor()
}

func (SpanLayer) Type() protoreflect.EnumType {
	return &file_common_trace_common_proto_enumTypes[2]
}

func (x SpanLayer) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SpanLayer.Descriptor instead.
func (SpanLayer) EnumDescriptor() ([]byte, []int) {
	return file_common_trace_common_proto_rawDescGZIP(), []int{2}
}

type UpstreamSegment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GlobalTraceIds []*UniqueId `protobuf:"bytes,1,rep,name=globalTraceIds,proto3" json:"globalTraceIds,omitempty"`
	Segment        []byte      `protobuf:"bytes,2,opt,name=segment,proto3" json:"segment,omitempty"` // the byte array of TraceSegmentObject
}

func (x *UpstreamSegment) Reset() {
	*x = UpstreamSegment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_trace_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpstreamSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpstreamSegment) ProtoMessage() {}

func (x *UpstreamSegment) ProtoReflect() protoreflect.Message {
	mi := &file_common_trace_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpstreamSegment.ProtoReflect.Descriptor instead.
func (*UpstreamSegment) Descriptor() ([]byte, []int) {
	return file_common_trace_common_proto_rawDescGZIP(), []int{0}
}

func (x *UpstreamSegment) GetGlobalTraceIds() []*UniqueId {
	if x != nil {
		return x.GlobalTraceIds
	}
	return nil
}

func (x *UpstreamSegment) GetSegment() []byte {
	if x != nil {
		return x.Segment
	}
	return nil
}

type UniqueId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdParts []int64 `protobuf:"varint,1,rep,packed,name=idParts,proto3" json:"idParts,omitempty"`
}

func (x *UniqueId) Reset() {
	*x = UniqueId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_trace_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniqueId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniqueId) ProtoMessage() {}

func (x *UniqueId) ProtoReflect() protoreflect.Message {
	mi := &file_common_trace_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniqueId.ProtoReflect.Descriptor instead.
func (*UniqueId) Descriptor() ([]byte, []int) {
	return file_common_trace_common_proto_rawDescGZIP(), []int{1}
}

func (x *UniqueId) GetIdParts() []int64 {
	if x != nil {
		return x.IdParts
	}
	return nil
}

var File_common_trace_common_proto protoreflect.FileDescriptor

var file_common_trace_common_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2d, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x0f, 0x55,
	0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x31,
	0x0a, 0x0e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49,
	0x64, 0x52, 0x0e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x24, 0x0a, 0x08, 0x55,
	0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x64, 0x50, 0x61, 0x72,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x69, 0x64, 0x50, 0x61, 0x72, 0x74,
	0x73, 0x2a, 0x2a, 0x0a, 0x08, 0x53, 0x70, 0x61, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x78, 0x69, 0x74,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x10, 0x02, 0x2a, 0x2c, 0x0a,
	0x07, 0x52, 0x65, 0x66, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x72, 0x6f, 0x73,
	0x73, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x72,
	0x6f, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x10, 0x01, 0x2a, 0x55, 0x0a, 0x09, 0x53,
	0x70, 0x61, 0x6e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x50, 0x43, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x74, 0x74, 0x70, 0x10, 0x03, 0x12,
	0x06, 0x0a, 0x02, 0x4d, 0x51, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x10, 0x05, 0x42, 0xa1, 0x01, 0x0a, 0x30, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68,
	0x65, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x6d,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x01, 0x5a, 0x4e, 0x6c, 0x6f, 0x67, 0x74, 0x61,
	0x69, 0x6c, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2f, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e,
	0x67, 0x76, 0x32, 0x2f, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x61,
	0x70, 0x6d, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0xaa, 0x02, 0x1a, 0x53, 0x6b, 0x79, 0x57,
	0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_trace_common_proto_rawDescOnce sync.Once
	file_common_trace_common_proto_rawDescData = file_common_trace_common_proto_rawDesc
)

func file_common_trace_common_proto_rawDescGZIP() []byte {
	file_common_trace_common_proto_rawDescOnce.Do(func() {
		file_common_trace_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_trace_common_proto_rawDescData)
	})
	return file_common_trace_common_proto_rawDescData
}

var file_common_trace_common_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_common_trace_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_trace_common_proto_goTypes = []interface{}{
	(SpanType)(0),           // 0: SpanType
	(RefType)(0),            // 1: RefType
	(SpanLayer)(0),          // 2: SpanLayer
	(*UpstreamSegment)(nil), // 3: UpstreamSegment
	(*UniqueId)(nil),        // 4: UniqueId
}
var file_common_trace_common_proto_depIdxs = []int32{
	4, // 0: UpstreamSegment.globalTraceIds:type_name -> UniqueId
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_trace_common_proto_init() }
func file_common_trace_common_proto_init() {
	if File_common_trace_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_trace_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpstreamSegment); i {
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
		file_common_trace_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniqueId); i {
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
			RawDescriptor: file_common_trace_common_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_trace_common_proto_goTypes,
		DependencyIndexes: file_common_trace_common_proto_depIdxs,
		EnumInfos:         file_common_trace_common_proto_enumTypes,
		MessageInfos:      file_common_trace_common_proto_msgTypes,
	}.Build()
	File_common_trace_common_proto = out.File
	file_common_trace_common_proto_rawDesc = nil
	file_common_trace_common_proto_goTypes = nil
	file_common_trace_common_proto_depIdxs = nil
}

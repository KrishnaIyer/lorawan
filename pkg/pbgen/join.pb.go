// Copyright © 2021 Krishna Iyer Easwaran
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: join.proto

package pbgen

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JoinEui  []byte `protobuf:"bytes,1,opt,name=join_eui,json=joinEui,proto3" json:"join_eui,omitempty"`
	DevEui   []byte `protobuf:"bytes,2,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	DevNonce []byte `protobuf:"bytes,3,opt,name=dev_nonce,json=devNonce,proto3" json:"dev_nonce,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_join_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_join_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_join_proto_rawDescGZIP(), []int{0}
}

func (x *JoinRequest) GetJoinEui() []byte {
	if x != nil {
		return x.JoinEui
	}
	return nil
}

func (x *JoinRequest) GetDevEui() []byte {
	if x != nil {
		return x.DevEui
	}
	return nil
}

func (x *JoinRequest) GetDevNonce() []byte {
	if x != nil {
		return x.DevNonce
	}
	return nil
}

var File_join_proto protoreflect.FileDescriptor

var file_join_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6a, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6c, 0x6f,
	0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x0b, 0x4a,
	0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x08, 0x6a, 0x6f,
	0x69, 0x6e, 0x5f, 0x65, 0x75, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x7a, 0x02, 0x68, 0x08, 0x52, 0x07, 0x6a, 0x6f, 0x69, 0x6e, 0x45, 0x75, 0x69, 0x12, 0x20,
	0x0a, 0x07, 0x64, 0x65, 0x76, 0x5f, 0x65, 0x75, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x7a, 0x02, 0x68, 0x08, 0x52, 0x06, 0x64, 0x65, 0x76, 0x45, 0x75, 0x69,
	0x12, 0x24, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x7a, 0x02, 0x68, 0x02, 0x52, 0x08, 0x64, 0x65,
	0x76, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x6f, 0x2e, 0x6b, 0x72, 0x69,
	0x73, 0x68, 0x6e, 0x61, 0x69, 0x79, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6c, 0x6f, 0x72,
	0x61, 0x77, 0x61, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_join_proto_rawDescOnce sync.Once
	file_join_proto_rawDescData = file_join_proto_rawDesc
)

func file_join_proto_rawDescGZIP() []byte {
	file_join_proto_rawDescOnce.Do(func() {
		file_join_proto_rawDescData = protoimpl.X.CompressGZIP(file_join_proto_rawDescData)
	})
	return file_join_proto_rawDescData
}

var file_join_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_join_proto_goTypes = []interface{}{
	(*JoinRequest)(nil), // 0: lorawan.v1.JoinRequest
}
var file_join_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_join_proto_init() }
func file_join_proto_init() {
	if File_join_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_join_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRequest); i {
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
			RawDescriptor: file_join_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_join_proto_goTypes,
		DependencyIndexes: file_join_proto_depIdxs,
		MessageInfos:      file_join_proto_msgTypes,
	}.Build()
	File_join_proto = out.File
	file_join_proto_rawDesc = nil
	file_join_proto_goTypes = nil
	file_join_proto_depIdxs = nil
}

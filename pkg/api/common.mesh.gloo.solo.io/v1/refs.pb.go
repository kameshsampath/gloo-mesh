// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo-mesh/api/common/v1/refs.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TODO: migrate to skv2 core apis
// Object providing a list of object refs.
// Used to store lists of refs inside a map.
type ObjectRefList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Refs []*v1.ObjectRef `protobuf:"bytes,1,rep,name=refs,proto3" json:"refs,omitempty"`
}

func (x *ObjectRefList) Reset() {
	*x = ObjectRefList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectRefList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectRefList) ProtoMessage() {}

func (x *ObjectRefList) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectRefList.ProtoReflect.Descriptor instead.
func (*ObjectRefList) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_rawDescGZIP(), []int{0}
}

func (x *ObjectRefList) GetRefs() []*v1.ObjectRef {
	if x != nil {
		return x.Refs
	}
	return nil
}

var File_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x66,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x0d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x72, 0x65, 0x66, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x66, 0x52, 0x04, 0x72, 0x65, 0x66, 0x73, 0x42, 0x46, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0xc0, 0xf5, 0x04, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_rawDescData = file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_goTypes = []interface{}{
	(*ObjectRefList)(nil), // 0: common.mesh.gloo.solo.io.ObjectRefList
	(*v1.ObjectRef)(nil),  // 1: core.skv2.solo.io.ObjectRef
}
var file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_depIdxs = []int32{
	1, // 0: common.mesh.gloo.solo.io.ObjectRefList.refs:type_name -> core.skv2.solo.io.ObjectRef
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_init() }
func file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_init() {
	if File_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectRefList); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto = out.File
	file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_api_common_v1_refs_proto_depIdxs = nil
}

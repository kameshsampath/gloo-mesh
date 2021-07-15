// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo-mesh/api/enterprise/istio/v1/istio_installation.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The state of a IstioOperator installation
type IstioInstallationStatus_IstioOperatorStatus_State int32

const (
	// Waiting for resources to be reconciled
	IstioInstallationStatus_IstioOperatorStatus_PENDING IstioInstallationStatus_IstioOperatorStatus_State = 0
	// In the process of reconciling Istio resources on to the managed cluster
	IstioInstallationStatus_IstioOperatorStatus_RECONCILING IstioInstallationStatus_IstioOperatorStatus_State = 1
	// All Istio components were installed successfully and they are healthy
	IstioInstallationStatus_IstioOperatorStatus_HEALTHY IstioInstallationStatus_IstioOperatorStatus_State = 2
	// One or more installed Istio component(s) in an error state
	IstioInstallationStatus_IstioOperatorStatus_ERROR IstioInstallationStatus_IstioOperatorStatus_State = 3
)

// Enum value maps for IstioInstallationStatus_IstioOperatorStatus_State.
var (
	IstioInstallationStatus_IstioOperatorStatus_State_name = map[int32]string{
		0: "PENDING",
		1: "RECONCILING",
		2: "HEALTHY",
		3: "ERROR",
	}
	IstioInstallationStatus_IstioOperatorStatus_State_value = map[string]int32{
		"PENDING":     0,
		"RECONCILING": 1,
		"HEALTHY":     2,
		"ERROR":       3,
	}
)

func (x IstioInstallationStatus_IstioOperatorStatus_State) Enum() *IstioInstallationStatus_IstioOperatorStatus_State {
	p := new(IstioInstallationStatus_IstioOperatorStatus_State)
	*p = x
	return p
}

func (x IstioInstallationStatus_IstioOperatorStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IstioInstallationStatus_IstioOperatorStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_enumTypes[0].Descriptor()
}

func (IstioInstallationStatus_IstioOperatorStatus_State) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_enumTypes[0]
}

func (x IstioInstallationStatus_IstioOperatorStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IstioInstallationStatus_IstioOperatorStatus_State.Descriptor instead.
func (IstioInstallationStatus_IstioOperatorStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_rawDescGZIP(), []int{1, 0, 0}
}

type IstioInstallationSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The cluster where the IstioOperators should be installed
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// References for the IstioOperators that should be installed in the managed cluster
	IstioOperatorRefs []*v1.ObjectRef `protobuf:"bytes,2,rep,name=istio_operator_refs,json=istioOperatorRefs,proto3" json:"istio_operator_refs,omitempty"`
}

func (x *IstioInstallationSpec) Reset() {
	*x = IstioInstallationSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IstioInstallationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IstioInstallationSpec) ProtoMessage() {}

func (x *IstioInstallationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IstioInstallationSpec.ProtoReflect.Descriptor instead.
func (*IstioInstallationSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_rawDescGZIP(), []int{0}
}

func (x *IstioInstallationSpec) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *IstioInstallationSpec) GetIstioOperatorRefs() []*v1.ObjectRef {
	if x != nil {
		return x.IstioOperatorRefs
	}
	return nil
}

type IstioInstallationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The most recent generation observed in the the IstioInstallation metadata.
	// If the `observedGeneration` does not match `metadata.generation`,
	// Gloo Mesh has not processed the most recent version of this resource.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// The status of the installation for each IstioOperator that should be applied.
	Statuses map[string]*IstioInstallationStatus_IstioOperatorStatus `protobuf:"bytes,2,rep,name=statuses,proto3" json:"statuses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *IstioInstallationStatus) Reset() {
	*x = IstioInstallationStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IstioInstallationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IstioInstallationStatus) ProtoMessage() {}

func (x *IstioInstallationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IstioInstallationStatus.ProtoReflect.Descriptor instead.
func (*IstioInstallationStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_rawDescGZIP(), []int{1}
}

func (x *IstioInstallationStatus) GetObservedGeneration() int64 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *IstioInstallationStatus) GetStatuses() map[string]*IstioInstallationStatus_IstioOperatorStatus {
	if x != nil {
		return x.Statuses
	}
	return nil
}

type IstioInstallationStatus_IstioOperatorStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The most recent generation observed in the the IstioOperator metadata.
	// If the `observedGeneration` does not match `metadata.generation`,
	// Gloo Mesh has not processed the most recent version of this resource.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// The current state of the IstioOperator
	State IstioInstallationStatus_IstioOperatorStatus_State `protobuf:"varint,2,opt,name=state,proto3,enum=istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus_IstioOperatorStatus_State" json:"state,omitempty"`
	// A human readable message about the current state of the IstioOperator
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *IstioInstallationStatus_IstioOperatorStatus) Reset() {
	*x = IstioInstallationStatus_IstioOperatorStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IstioInstallationStatus_IstioOperatorStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IstioInstallationStatus_IstioOperatorStatus) ProtoMessage() {}

func (x *IstioInstallationStatus_IstioOperatorStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IstioInstallationStatus_IstioOperatorStatus.ProtoReflect.Descriptor instead.
func (*IstioInstallationStatus_IstioOperatorStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_rawDescGZIP(), []int{1, 0}
}

func (x *IstioInstallationStatus_IstioOperatorStatus) GetObservedGeneration() int64 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *IstioInstallationStatus_IstioOperatorStatus) GetState() IstioInstallationStatus_IstioOperatorStatus_State {
	if x != nil {
		return x.State
	}
	return IstioInstallationStatus_IstioOperatorStatus_PENDING
}

func (x *IstioInstallationStatus_IstioOperatorStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x69, 0x73,
	0x74, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x5f, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x22, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73,
	0x65, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x15, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x4c, 0x0a, 0x13, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x72, 0x65, 0x66, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x52, 0x11, 0x69, 0x73, 0x74,
	0x69, 0x6f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x66, 0x73, 0x22, 0xcf,
	0x04, 0x0a, 0x17, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x6f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x65, 0x0a, 0x08, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x49, 0x2e,
	0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65,
	0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x65, 0x73, 0x1a, 0x8c, 0x02, 0x0a, 0x13, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x6f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6b, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x55, 0x2e, 0x69, 0x73, 0x74,
	0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x49, 0x73, 0x74, 0x69, 0x6f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x3d, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x43, 0x4f,
	0x4e, 0x43, 0x49, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x45, 0x41,
	0x4c, 0x54, 0x48, 0x59, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x03, 0x1a, 0x8c, 0x01, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x65, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x4f, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_rawDescData = file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_goTypes = []interface{}{
	(IstioInstallationStatus_IstioOperatorStatus_State)(0), // 0: istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus.IstioOperatorStatus.State
	(*IstioInstallationSpec)(nil),                          // 1: istio.enterprise.mesh.gloo.solo.io.IstioInstallationSpec
	(*IstioInstallationStatus)(nil),                        // 2: istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus
	(*IstioInstallationStatus_IstioOperatorStatus)(nil),    // 3: istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus.IstioOperatorStatus
	nil,                  // 4: istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus.StatusesEntry
	(*v1.ObjectRef)(nil), // 5: core.skv2.solo.io.ObjectRef
}
var file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_depIdxs = []int32{
	5, // 0: istio.enterprise.mesh.gloo.solo.io.IstioInstallationSpec.istio_operator_refs:type_name -> core.skv2.solo.io.ObjectRef
	4, // 1: istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus.statuses:type_name -> istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus.StatusesEntry
	0, // 2: istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus.IstioOperatorStatus.state:type_name -> istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus.IstioOperatorStatus.State
	3, // 3: istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus.StatusesEntry.value:type_name -> istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus.IstioOperatorStatus
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_init()
}
func file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_init() {
	if File_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IstioInstallationSpec); i {
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
		file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IstioInstallationStatus); i {
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
		file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IstioInstallationStatus_IstioOperatorStatus); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto = out.File
	file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_api_enterprise_istio_v1_istio_installation_proto_depIdxs = nil
}

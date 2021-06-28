// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo-mesh/api/certificates/v1/ca_options.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// State of Certificate Rotation
// Possible states in which a CertificateRotation can exist.
type CertificateRotationState int32

const (
	// No Certificate rotation is currently happening
	CertificateRotationState_NOT_APPLICABLE CertificateRotationState = 0
	// The CertificateRotation has yet to be picked up by the management-plane.
	CertificateRotationState_PENDING CertificateRotationState = 1
	// The CertificateRotation is underway, both roots are set, and the new root is being propogated
	CertificateRotationState_ADDING_NEW_ROOT CertificateRotationState = 2
	// The CertificateRotation is underway again.
	// The initial verification is over, the traffic continues to work with both roots present.
	// Now the old root is being removed, and the new root is being propgated alone to the data-plane clusters
	CertificateRotationState_PROPOGATING_NEW_INTERMEDIATE CertificateRotationState = 3
	// The CertificateRotation is underway again.
	// Removing the old-root from all data-plane clusters
	CertificateRotationState_DELETING_OLD_ROOT CertificateRotationState = 4
	// Verifying connectivity between workloads, the workflow will not progress until connectivity has been verified.
	// This can either be manual or in the future automated
	CertificateRotationState_VERIFYING CertificateRotationState = 5
	// The connectivity has been verified.
	CertificateRotationState_VERIFIED CertificateRotationState = 6
	// The rotation has finished, the new root has been propgated to all data-plane clusters, and traffic has
	// been verified successfully.
	CertificateRotationState_FINISHED CertificateRotationState = 7
	// Processing the certificate rotation workflow failed.
	CertificateRotationState_FAILED CertificateRotationState = 8
)

// Enum value maps for CertificateRotationState.
var (
	CertificateRotationState_name = map[int32]string{
		0: "NOT_APPLICABLE",
		1: "PENDING",
		2: "ADDING_NEW_ROOT",
		3: "PROPOGATING_NEW_INTERMEDIATE",
		4: "DELETING_OLD_ROOT",
		5: "VERIFYING",
		6: "VERIFIED",
		7: "FINISHED",
		8: "FAILED",
	}
	CertificateRotationState_value = map[string]int32{
		"NOT_APPLICABLE":               0,
		"PENDING":                      1,
		"ADDING_NEW_ROOT":              2,
		"PROPOGATING_NEW_INTERMEDIATE": 3,
		"DELETING_OLD_ROOT":            4,
		"VERIFYING":                    5,
		"VERIFIED":                     6,
		"FINISHED":                     7,
		"FAILED":                       8,
	}
)

func (x CertificateRotationState) Enum() *CertificateRotationState {
	p := new(CertificateRotationState)
	*p = x
	return p
}

func (x CertificateRotationState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CertificateRotationState) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_enumTypes[0].Descriptor()
}

func (CertificateRotationState) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_enumTypes[0]
}

func (x CertificateRotationState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CertificateRotationState.Descriptor instead.
func (CertificateRotationState) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDescGZIP(), []int{0}
}

// Configuration for generating a self-signed root certificate.
// Uses the X.509 format, RFC5280.
type CommonCertOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of days before root cert expires. Defaults to 365.
	TtlDays uint32 `protobuf:"varint,1,opt,name=ttl_days,json=ttlDays,proto3" json:"ttl_days,omitempty"`
	// Size in bytes of the root cert's private key. Defaults to 4096.
	RsaKeySizeBytes uint32 `protobuf:"varint,2,opt,name=rsa_key_size_bytes,json=rsaKeySizeBytes,proto3" json:"rsa_key_size_bytes,omitempty"`
	// Root cert organization name. Defaults to "gloo-mesh".
	OrgName string `protobuf:"bytes,3,opt,name=org_name,json=orgName,proto3" json:"org_name,omitempty"`
	// The ratio of cert lifetime to refresh a cert. For example, at 0.10 and 1 hour TTL,
	// we would refresh 6 minutes before expiration
	SecretRotationGracePeriodRatio float32 `protobuf:"fixed32,4,opt,name=secret_rotation_grace_period_ratio,json=secretRotationGracePeriodRatio,proto3" json:"secret_rotation_grace_period_ratio,omitempty"`
}

func (x *CommonCertOptions) Reset() {
	*x = CommonCertOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonCertOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonCertOptions) ProtoMessage() {}

func (x *CommonCertOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonCertOptions.ProtoReflect.Descriptor instead.
func (*CommonCertOptions) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDescGZIP(), []int{0}
}

func (x *CommonCertOptions) GetTtlDays() uint32 {
	if x != nil {
		return x.TtlDays
	}
	return 0
}

func (x *CommonCertOptions) GetRsaKeySizeBytes() uint32 {
	if x != nil {
		return x.RsaKeySizeBytes
	}
	return 0
}

func (x *CommonCertOptions) GetOrgName() string {
	if x != nil {
		return x.OrgName
	}
	return ""
}

func (x *CommonCertOptions) GetSecretRotationGracePeriodRatio() float32 {
	if x != nil {
		return x.SecretRotationGracePeriodRatio
	}
	return 0
}

// Specify parameters for configuring the root certificate authority for a VirtualMesh.
type IntermediateCertificateAuthority struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specify the source of the Root CA data which Gloo Mesh will use for the VirtualMesh.
	//
	// Types that are assignable to CaSource:
	//	*IntermediateCertificateAuthority_Vault
	CaSource isIntermediateCertificateAuthority_CaSource `protobuf_oneof:"ca_source"`
}

func (x *IntermediateCertificateAuthority) Reset() {
	*x = IntermediateCertificateAuthority{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntermediateCertificateAuthority) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntermediateCertificateAuthority) ProtoMessage() {}

func (x *IntermediateCertificateAuthority) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntermediateCertificateAuthority.ProtoReflect.Descriptor instead.
func (*IntermediateCertificateAuthority) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDescGZIP(), []int{1}
}

func (m *IntermediateCertificateAuthority) GetCaSource() isIntermediateCertificateAuthority_CaSource {
	if m != nil {
		return m.CaSource
	}
	return nil
}

func (x *IntermediateCertificateAuthority) GetVault() *VaultCA {
	if x, ok := x.GetCaSource().(*IntermediateCertificateAuthority_Vault); ok {
		return x.Vault
	}
	return nil
}

type isIntermediateCertificateAuthority_CaSource interface {
	isIntermediateCertificateAuthority_CaSource()
}

type IntermediateCertificateAuthority_Vault struct {
	// Use vault as the intermediate CA source
	Vault *VaultCA `protobuf:"bytes,1,opt,name=vault,proto3,oneof"`
}

func (*IntermediateCertificateAuthority_Vault) isIntermediateCertificateAuthority_CaSource() {}

// CertificateRotationCondition represents a timesptamped snapshot of the certificate
// rotation workflow. This is used to keep track of the steps which have been completed
// thus far.
type CertificateRotationCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The time at which this condition was recorded
	Timestamp string `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The current state of the cert rotation
	State CertificateRotationState `protobuf:"varint,2,opt,name=state,proto3,enum=certificates.mesh.gloo.solo.io.CertificateRotationState" json:"state,omitempty"`
	// A human readable message related to the current condition
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// Any errors which occured during the current rotation stage
	Errors []string `protobuf:"bytes,4,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *CertificateRotationCondition) Reset() {
	*x = CertificateRotationCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateRotationCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateRotationCondition) ProtoMessage() {}

func (x *CertificateRotationCondition) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateRotationCondition.ProtoReflect.Descriptor instead.
func (*CertificateRotationCondition) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDescGZIP(), []int{2}
}

func (x *CertificateRotationCondition) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *CertificateRotationCondition) GetState() CertificateRotationState {
	if x != nil {
		return x.State
	}
	return CertificateRotationState_NOT_APPLICABLE
}

func (x *CertificateRotationCondition) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CertificateRotationCondition) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDesc = []byte{
	0x0a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x61, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65,
	0x73, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x43, 0x65, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x74,
	0x74, 0x6c, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x74,
	0x74, 0x6c, 0x44, 0x61, 0x79, 0x73, 0x12, 0x2b, 0x0a, 0x12, 0x72, 0x73, 0x61, 0x5f, 0x6b, 0x65,
	0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0f, 0x72, 0x73, 0x61, 0x4b, 0x65, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4a,
	0x0a, 0x22, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x67, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x1e, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x61, 0x63, 0x65, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x22, 0x70, 0x0a, 0x20, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x3f,
	0x0a, 0x05, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x56,
	0x61, 0x75, 0x6c, 0x74, 0x43, 0x41, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x42,
	0x0b, 0x0a, 0x09, 0x63, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0xbe, 0x01, 0x0a,
	0x1c, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x4e, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x63, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2a, 0xc0, 0x01,
	0x0a, 0x18, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x4f,
	0x54, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x41,
	0x44, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x52, 0x4f, 0x4f, 0x54, 0x10, 0x02,
	0x12, 0x20, 0x0a, 0x1c, 0x50, 0x52, 0x4f, 0x50, 0x4f, 0x47, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f,
	0x4e, 0x45, 0x57, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x54, 0x45,
	0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4f,
	0x4c, 0x44, 0x5f, 0x52, 0x4f, 0x4f, 0x54, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x56, 0x45, 0x52,
	0x49, 0x46, 0x59, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x56, 0x45, 0x52, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48,
	0x45, 0x44, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x08,
	0x42, 0x4c, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDescData = file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_goTypes = []interface{}{
	(CertificateRotationState)(0),            // 0: certificates.mesh.gloo.solo.io.CertificateRotationState
	(*CommonCertOptions)(nil),                // 1: certificates.mesh.gloo.solo.io.CommonCertOptions
	(*IntermediateCertificateAuthority)(nil), // 2: certificates.mesh.gloo.solo.io.IntermediateCertificateAuthority
	(*CertificateRotationCondition)(nil),     // 3: certificates.mesh.gloo.solo.io.CertificateRotationCondition
	(*VaultCA)(nil),                          // 4: certificates.mesh.gloo.solo.io.VaultCA
}
var file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_depIdxs = []int32{
	4, // 0: certificates.mesh.gloo.solo.io.IntermediateCertificateAuthority.vault:type_name -> certificates.mesh.gloo.solo.io.VaultCA
	0, // 1: certificates.mesh.gloo.solo.io.CertificateRotationCondition.state:type_name -> certificates.mesh.gloo.solo.io.CertificateRotationState
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_init() }
func file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_init() {
	if File_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto != nil {
		return
	}
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonCertOptions); i {
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
		file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntermediateCertificateAuthority); i {
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
		file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateRotationCondition); i {
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
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*IntermediateCertificateAuthority_Vault)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto = out.File
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_depIdxs = nil
}

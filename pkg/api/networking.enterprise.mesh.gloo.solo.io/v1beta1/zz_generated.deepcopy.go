// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for networking.enterprise.mesh.gloo.solo.io/v1beta1 resources

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// Generated Deepcopy methods for WasmDeployment

func (in *WasmDeployment) DeepCopyInto(out *WasmDeployment) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *WasmDeployment) DeepCopy() *WasmDeployment {
	if in == nil {
		return nil
	}
	out := new(WasmDeployment)
	in.DeepCopyInto(out)
	return out
}

func (in *WasmDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *WasmDeploymentList) DeepCopyInto(out *WasmDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WasmDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *WasmDeploymentList) DeepCopy() *WasmDeploymentList {
	if in == nil {
		return nil
	}
	out := new(WasmDeploymentList)
	in.DeepCopyInto(out)
	return out
}

func (in *WasmDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for RateLimiterServerConfig

func (in *RateLimiterServerConfig) DeepCopyInto(out *RateLimiterServerConfig) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *RateLimiterServerConfig) DeepCopy() *RateLimiterServerConfig {
	if in == nil {
		return nil
	}
	out := new(RateLimiterServerConfig)
	in.DeepCopyInto(out)
	return out
}

func (in *RateLimiterServerConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *RateLimiterServerConfigList) DeepCopyInto(out *RateLimiterServerConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RateLimiterServerConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *RateLimiterServerConfigList) DeepCopy() *RateLimiterServerConfigList {
	if in == nil {
		return nil
	}
	out := new(RateLimiterServerConfigList)
	in.DeepCopyInto(out)
	return out
}

func (in *RateLimiterServerConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for VirtualDestination

func (in *VirtualDestination) DeepCopyInto(out *VirtualDestination) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *VirtualDestination) DeepCopy() *VirtualDestination {
	if in == nil {
		return nil
	}
	out := new(VirtualDestination)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualDestination) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *VirtualDestinationList) DeepCopyInto(out *VirtualDestinationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualDestination, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *VirtualDestinationList) DeepCopy() *VirtualDestinationList {
	if in == nil {
		return nil
	}
	out := new(VirtualDestinationList)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualDestinationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for VirtualGateway

func (in *VirtualGateway) DeepCopyInto(out *VirtualGateway) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *VirtualGateway) DeepCopy() *VirtualGateway {
	if in == nil {
		return nil
	}
	out := new(VirtualGateway)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualGateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *VirtualGatewayList) DeepCopyInto(out *VirtualGatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualGateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *VirtualGatewayList) DeepCopy() *VirtualGatewayList {
	if in == nil {
		return nil
	}
	out := new(VirtualGatewayList)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualGatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for VirtualHost

func (in *VirtualHost) DeepCopyInto(out *VirtualHost) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *VirtualHost) DeepCopy() *VirtualHost {
	if in == nil {
		return nil
	}
	out := new(VirtualHost)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualHost) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *VirtualHostList) DeepCopyInto(out *VirtualHostList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualHost, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *VirtualHostList) DeepCopy() *VirtualHostList {
	if in == nil {
		return nil
	}
	out := new(VirtualHostList)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualHostList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for RouteTable

func (in *RouteTable) DeepCopyInto(out *RouteTable) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *RouteTable) DeepCopy() *RouteTable {
	if in == nil {
		return nil
	}
	out := new(RouteTable)
	in.DeepCopyInto(out)
	return out
}

func (in *RouteTable) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *RouteTableList) DeepCopyInto(out *RouteTableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RouteTable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *RouteTableList) DeepCopy() *RouteTableList {
	if in == nil {
		return nil
	}
	out := new(RouteTableList)
	in.DeepCopyInto(out)
	return out
}

func (in *RouteTableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for ServiceDependency

func (in *ServiceDependency) DeepCopyInto(out *ServiceDependency) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *ServiceDependency) DeepCopy() *ServiceDependency {
	if in == nil {
		return nil
	}
	out := new(ServiceDependency)
	in.DeepCopyInto(out)
	return out
}

func (in *ServiceDependency) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *ServiceDependencyList) DeepCopyInto(out *ServiceDependencyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceDependency, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *ServiceDependencyList) DeepCopy() *ServiceDependencyList {
	if in == nil {
		return nil
	}
	out := new(ServiceDependencyList)
	in.DeepCopyInto(out)
	return out
}

func (in *ServiceDependencyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for CertificateVerification

func (in *CertificateVerification) DeepCopyInto(out *CertificateVerification) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *CertificateVerification) DeepCopy() *CertificateVerification {
	if in == nil {
		return nil
	}
	out := new(CertificateVerification)
	in.DeepCopyInto(out)
	return out
}

func (in *CertificateVerification) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *CertificateVerificationList) DeepCopyInto(out *CertificateVerificationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CertificateVerification, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *CertificateVerificationList) DeepCopy() *CertificateVerificationList {
	if in == nil {
		return nil
	}
	out := new(CertificateVerificationList)
	in.DeepCopyInto(out)
	return out
}

func (in *CertificateVerificationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

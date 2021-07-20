// Code generated by skv2. DO NOT EDIT.

/*
	Utility for manually building input snapshots. Used primarily in tests.
*/
package input

import (
	certificates_mesh_gloo_solo_io_v1 "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1"
	certificates_mesh_gloo_solo_io_v1_sets "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1/sets"

	xds_agent_enterprise_mesh_gloo_solo_io_v1beta1 "github.com/solo-io/gloo-mesh/pkg/api/xds.agent.enterprise.mesh.gloo.solo.io/v1beta1"
	xds_agent_enterprise_mesh_gloo_solo_io_v1beta1_sets "github.com/solo-io/gloo-mesh/pkg/api/xds.agent.enterprise.mesh.gloo.solo.io/v1beta1/sets"

	networking_istio_io_v1alpha3_sets "github.com/solo-io/external-apis/pkg/api/istio/networking.istio.io/v1alpha3/sets"
	networking_istio_io_v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"

	security_istio_io_v1beta1_sets "github.com/solo-io/external-apis/pkg/api/istio/security.istio.io/v1beta1/sets"
	security_istio_io_v1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"

	install_istio_io_v1alpha1_sets "github.com/solo-io/external-apis/pkg/api/istio/install.istio.io/v1alpha1/sets"
	install_istio_io_v1alpha1 "istio.io/istio/operator/pkg/apis/istio/v1alpha1"
)

type InputRemoteSnapshotManualBuilder struct {
	name string

	issuedCertificates  certificates_mesh_gloo_solo_io_v1_sets.IssuedCertificateSet
	podBounceDirectives certificates_mesh_gloo_solo_io_v1_sets.PodBounceDirectiveSet

	xdsConfigs xds_agent_enterprise_mesh_gloo_solo_io_v1beta1_sets.XdsConfigSet

	destinationRules networking_istio_io_v1alpha3_sets.DestinationRuleSet
	envoyFilters     networking_istio_io_v1alpha3_sets.EnvoyFilterSet
	gateways         networking_istio_io_v1alpha3_sets.GatewaySet
	serviceEntries   networking_istio_io_v1alpha3_sets.ServiceEntrySet
	virtualServices  networking_istio_io_v1alpha3_sets.VirtualServiceSet
	sidecars         networking_istio_io_v1alpha3_sets.SidecarSet

	authorizationPolicies security_istio_io_v1beta1_sets.AuthorizationPolicySet

	istioOperators install_istio_io_v1alpha1_sets.IstioOperatorSet
}

func NewInputRemoteSnapshotManualBuilder(name string) *InputRemoteSnapshotManualBuilder {
	return &InputRemoteSnapshotManualBuilder{
		name: name,

		issuedCertificates:  certificates_mesh_gloo_solo_io_v1_sets.NewIssuedCertificateSet(),
		podBounceDirectives: certificates_mesh_gloo_solo_io_v1_sets.NewPodBounceDirectiveSet(),

		xdsConfigs: xds_agent_enterprise_mesh_gloo_solo_io_v1beta1_sets.NewXdsConfigSet(),

		destinationRules: networking_istio_io_v1alpha3_sets.NewDestinationRuleSet(),
		envoyFilters:     networking_istio_io_v1alpha3_sets.NewEnvoyFilterSet(),
		gateways:         networking_istio_io_v1alpha3_sets.NewGatewaySet(),
		serviceEntries:   networking_istio_io_v1alpha3_sets.NewServiceEntrySet(),
		virtualServices:  networking_istio_io_v1alpha3_sets.NewVirtualServiceSet(),
		sidecars:         networking_istio_io_v1alpha3_sets.NewSidecarSet(),

		authorizationPolicies: security_istio_io_v1beta1_sets.NewAuthorizationPolicySet(),

		istioOperators: install_istio_io_v1alpha1_sets.NewIstioOperatorSet(),
	}
}

func (i *InputRemoteSnapshotManualBuilder) Build() RemoteSnapshot {
	return NewRemoteSnapshot(
		i.name,

		i.issuedCertificates,
		i.podBounceDirectives,

		i.xdsConfigs,

		i.destinationRules,
		i.envoyFilters,
		i.gateways,
		i.serviceEntries,
		i.virtualServices,
		i.sidecars,

		i.authorizationPolicies,

		i.istioOperators,
	)
}
func (i *InputRemoteSnapshotManualBuilder) AddIssuedCertificates(issuedCertificates []*certificates_mesh_gloo_solo_io_v1.IssuedCertificate) *InputRemoteSnapshotManualBuilder {
	i.issuedCertificates.Insert(issuedCertificates...)
	return i
}
func (i *InputRemoteSnapshotManualBuilder) AddPodBounceDirectives(podBounceDirectives []*certificates_mesh_gloo_solo_io_v1.PodBounceDirective) *InputRemoteSnapshotManualBuilder {
	i.podBounceDirectives.Insert(podBounceDirectives...)
	return i
}
func (i *InputRemoteSnapshotManualBuilder) AddXdsConfigs(xdsConfigs []*xds_agent_enterprise_mesh_gloo_solo_io_v1beta1.XdsConfig) *InputRemoteSnapshotManualBuilder {
	i.xdsConfigs.Insert(xdsConfigs...)
	return i
}
func (i *InputRemoteSnapshotManualBuilder) AddDestinationRules(destinationRules []*networking_istio_io_v1alpha3.DestinationRule) *InputRemoteSnapshotManualBuilder {
	i.destinationRules.Insert(destinationRules...)
	return i
}
func (i *InputRemoteSnapshotManualBuilder) AddEnvoyFilters(envoyFilters []*networking_istio_io_v1alpha3.EnvoyFilter) *InputRemoteSnapshotManualBuilder {
	i.envoyFilters.Insert(envoyFilters...)
	return i
}
func (i *InputRemoteSnapshotManualBuilder) AddGateways(gateways []*networking_istio_io_v1alpha3.Gateway) *InputRemoteSnapshotManualBuilder {
	i.gateways.Insert(gateways...)
	return i
}
func (i *InputRemoteSnapshotManualBuilder) AddServiceEntries(serviceEntries []*networking_istio_io_v1alpha3.ServiceEntry) *InputRemoteSnapshotManualBuilder {
	i.serviceEntries.Insert(serviceEntries...)
	return i
}
func (i *InputRemoteSnapshotManualBuilder) AddVirtualServices(virtualServices []*networking_istio_io_v1alpha3.VirtualService) *InputRemoteSnapshotManualBuilder {
	i.virtualServices.Insert(virtualServices...)
	return i
}
func (i *InputRemoteSnapshotManualBuilder) AddSidecars(sidecars []*networking_istio_io_v1alpha3.Sidecar) *InputRemoteSnapshotManualBuilder {
	i.sidecars.Insert(sidecars...)
	return i
}
func (i *InputRemoteSnapshotManualBuilder) AddAuthorizationPolicies(authorizationPolicies []*security_istio_io_v1beta1.AuthorizationPolicy) *InputRemoteSnapshotManualBuilder {
	i.authorizationPolicies.Insert(authorizationPolicies...)
	return i
}
func (i *InputRemoteSnapshotManualBuilder) AddIstioOperators(istioOperators []*install_istio_io_v1alpha1.IstioOperator) *InputRemoteSnapshotManualBuilder {
	i.istioOperators.Insert(istioOperators...)
	return i
}

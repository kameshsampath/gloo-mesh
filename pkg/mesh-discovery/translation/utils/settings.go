package utils

import (
	"context"

	"github.com/rotisserie/eris"
	"github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/input"
	settingsv1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1alpha2"
	"github.com/solo-io/gloo-mesh/pkg/common/defaults"
	"github.com/solo-io/go-utils/contextutils"
)

// TODO this is copied from networking snapshot_utils. remove this when functionality is added to skv2
// Safely fetch the single SettingsMeshGlooSoloIov1Alpha2Settings object from snapshot. Log and error if not singleton.
func GetSingletonSettings(ctx context.Context, inLocal input.LocalSnapshot) (*settingsv1alpha2.Settings, error) {
	settings := inLocal.SettingsMeshGlooSoloIov1Alpha2Settings().List()
	n := len(settings)
	if n != 1 {
		err := eris.Errorf("Snapshot does not contain single SettingsMeshGlooSoloIov1Alpha2Settings object, %d found.", n)
		contextutils.LoggerFrom(ctx).Errorf("%+v", err)
		return nil, err
	}
	return settings[0], nil
}

// Get the workload labels and TLS port name used to detect ingress gateways in the given cluster.
func GetIngressGatewayDetector(settings *settingsv1alpha2.DiscoverySettings, clusterName string) (*settingsv1alpha2.DiscoverySettings_Istio_IngressGatewayDetector, error) {
	var labels map[string]string
	var portName string

	gatewayDetectors := settings.GetIstio().GetIngressGatewayDetectors()
	if gatewayDetectors != nil {
		// First, check if cluster-specific values are set
		if gatewayDetectors[clusterName] != nil {
			labels = gatewayDetectors[clusterName].GetGatewayWorkloadLabels()
			portName = gatewayDetectors[clusterName].GetGatewayTlsPortName()
		}

		// Check the wildcard (all clusters) entry
		if (labels == nil || portName == "") && gatewayDetectors["*"] != nil {
			if labels == nil {
				labels = gatewayDetectors["*"].GetGatewayWorkloadLabels()
			}
			if portName == "" {
				portName = gatewayDetectors["*"].GetGatewayTlsPortName()
			}
		}
	}

	// Fall back to default values
	if labels == nil {
		labels = defaults.DefaultGatewayWorkloadLabels
	}
	if portName == "" {
		portName = defaults.DefaultGatewayPortName
	}

	return &settingsv1alpha2.DiscoverySettings_Istio_IngressGatewayDetector{
		GatewayWorkloadLabels: labels,
		GatewayTlsPortName:    portName,
	}, nil
}

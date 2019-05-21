// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"fmt"

	gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"

	"github.com/solo-io/go-utils/hashutils"
	"go.uber.org/zap"
)

type ConfigSnapshot struct {
	Meshes        MeshList
	Meshingresses MeshIngressList
	Meshgroups    MeshGroupList
	Routingrules  RoutingRuleList
	Securityrules SecurityRuleList
	Tlssecrets    TlsSecretList
	Upstreams     gloo_solo_io.UpstreamList
	Pods          github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodList
	Services      github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.ServiceList
}

func (s ConfigSnapshot) Clone() ConfigSnapshot {
	return ConfigSnapshot{
		Meshes:        s.Meshes.Clone(),
		Meshingresses: s.Meshingresses.Clone(),
		Meshgroups:    s.Meshgroups.Clone(),
		Routingrules:  s.Routingrules.Clone(),
		Securityrules: s.Securityrules.Clone(),
		Tlssecrets:    s.Tlssecrets.Clone(),
		Upstreams:     s.Upstreams.Clone(),
		Pods:          s.Pods.Clone(),
		Services:      s.Services.Clone(),
	}
}

func (s ConfigSnapshot) Hash() uint64 {
	return hashutils.HashAll(
		s.hashMeshes(),
		s.hashMeshingresses(),
		s.hashMeshgroups(),
		s.hashRoutingrules(),
		s.hashSecurityrules(),
		s.hashTlssecrets(),
		s.hashUpstreams(),
		s.hashPods(),
		s.hashServices(),
	)
}

func (s ConfigSnapshot) hashMeshes() uint64 {
	return hashutils.HashAll(s.Meshes.AsInterfaces()...)
}

func (s ConfigSnapshot) hashMeshingresses() uint64 {
	return hashutils.HashAll(s.Meshingresses.AsInterfaces()...)
}

func (s ConfigSnapshot) hashMeshgroups() uint64 {
	return hashutils.HashAll(s.Meshgroups.AsInterfaces()...)
}

func (s ConfigSnapshot) hashRoutingrules() uint64 {
	return hashutils.HashAll(s.Routingrules.AsInterfaces()...)
}

func (s ConfigSnapshot) hashSecurityrules() uint64 {
	return hashutils.HashAll(s.Securityrules.AsInterfaces()...)
}

func (s ConfigSnapshot) hashTlssecrets() uint64 {
	return hashutils.HashAll(s.Tlssecrets.AsInterfaces()...)
}

func (s ConfigSnapshot) hashUpstreams() uint64 {
	return hashutils.HashAll(s.Upstreams.AsInterfaces()...)
}

func (s ConfigSnapshot) hashPods() uint64 {
	return hashutils.HashAll(s.Pods.AsInterfaces()...)
}

func (s ConfigSnapshot) hashServices() uint64 {
	return hashutils.HashAll(s.Services.AsInterfaces()...)
}

func (s ConfigSnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	fields = append(fields, zap.Uint64("meshes", s.hashMeshes()))
	fields = append(fields, zap.Uint64("meshingresses", s.hashMeshingresses()))
	fields = append(fields, zap.Uint64("meshgroups", s.hashMeshgroups()))
	fields = append(fields, zap.Uint64("routingrules", s.hashRoutingrules()))
	fields = append(fields, zap.Uint64("securityrules", s.hashSecurityrules()))
	fields = append(fields, zap.Uint64("tlssecrets", s.hashTlssecrets()))
	fields = append(fields, zap.Uint64("upstreams", s.hashUpstreams()))
	fields = append(fields, zap.Uint64("pods", s.hashPods()))
	fields = append(fields, zap.Uint64("services", s.hashServices()))

	return append(fields, zap.Uint64("snapshotHash", s.Hash()))
}

type ConfigSnapshotStringer struct {
	Version       uint64
	Meshes        []string
	Meshingresses []string
	Meshgroups    []string
	Routingrules  []string
	Securityrules []string
	Tlssecrets    []string
	Upstreams     []string
	Pods          []string
	Services      []string
}

func (ss ConfigSnapshotStringer) String() string {
	s := fmt.Sprintf("ConfigSnapshot %v\n", ss.Version)

	s += fmt.Sprintf("  Meshes %v\n", len(ss.Meshes))
	for _, name := range ss.Meshes {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Meshingresses %v\n", len(ss.Meshingresses))
	for _, name := range ss.Meshingresses {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Meshgroups %v\n", len(ss.Meshgroups))
	for _, name := range ss.Meshgroups {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Routingrules %v\n", len(ss.Routingrules))
	for _, name := range ss.Routingrules {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Securityrules %v\n", len(ss.Securityrules))
	for _, name := range ss.Securityrules {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Tlssecrets %v\n", len(ss.Tlssecrets))
	for _, name := range ss.Tlssecrets {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Upstreams %v\n", len(ss.Upstreams))
	for _, name := range ss.Upstreams {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Pods %v\n", len(ss.Pods))
	for _, name := range ss.Pods {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Services %v\n", len(ss.Services))
	for _, name := range ss.Services {
		s += fmt.Sprintf("    %v\n", name)
	}

	return s
}

func (s ConfigSnapshot) Stringer() ConfigSnapshotStringer {
	return ConfigSnapshotStringer{
		Version:       s.Hash(),
		Meshes:        s.Meshes.NamespacesDotNames(),
		Meshingresses: s.Meshingresses.NamespacesDotNames(),
		Meshgroups:    s.Meshgroups.NamespacesDotNames(),
		Routingrules:  s.Routingrules.NamespacesDotNames(),
		Securityrules: s.Securityrules.NamespacesDotNames(),
		Tlssecrets:    s.Tlssecrets.NamespacesDotNames(),
		Upstreams:     s.Upstreams.NamespacesDotNames(),
		Pods:          s.Pods.NamespacesDotNames(),
		Services:      s.Services.NamespacesDotNames(),
	}
}

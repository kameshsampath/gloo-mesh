// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./snapshot.go -destination mocks/snapshot.go

// Definitions for Output Snapshots
package certagent

import (
	"context"
	"encoding/json"
	"sort"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/pkg/multicluster"

	"github.com/rotisserie/eris"
	"github.com/solo-io/skv2/contrib/pkg/output"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"sigs.k8s.io/controller-runtime/pkg/client"

	certificates_mesh_gloo_solo_io_v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1alpha2"
	certificates_mesh_gloo_solo_io_v1alpha2_sets "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1alpha2/sets"

	v1_sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	v1 "k8s.io/api/core/v1"
)

// this error can occur if constructing a Partitioned Snapshot from a resource
// that is missing the partition label
var MissingRequiredLabelError = func(labelKey, resourceKind string, obj ezkube.ResourceId) error {
	return eris.Errorf("expected label %v not on labels of %v %v", labelKey, resourceKind, sets.Key(obj))
}

// the snapshot of output resources produced by a translation
type Snapshot interface {

	// return the set of CertificatesMeshGlooSoloIov1Alpha2CertificateRequests with a given set of labels
	CertificatesMeshGlooSoloIov1Alpha2CertificateRequests() []LabeledCertificatesMeshGlooSoloIov1Alpha2CertificateRequestSet
	// return the set of V1Secrets with a given set of labels
	V1Secrets() []LabeledV1SecretSet

	// apply the snapshot to the local cluster, garbage collecting stale resources
	ApplyLocalCluster(ctx context.Context, clusterClient client.Client, errHandler output.ErrorHandler)

	// apply resources from the snapshot across multiple clusters, garbage collecting stale resources
	ApplyMultiCluster(ctx context.Context, multiClusterClient multicluster.Client, errHandler output.ErrorHandler)

	// serialize the entire snapshot as JSON
	MarshalJSON() ([]byte, error)
}

type snapshot struct {
	name string

	certificatesMeshGlooSoloIov1Alpha2CertificateRequests []LabeledCertificatesMeshGlooSoloIov1Alpha2CertificateRequestSet
	v1Secrets                                             []LabeledV1SecretSet
	clusters                                              []string
}

func NewSnapshot(
	name string,

	certificatesMeshGlooSoloIov1Alpha2CertificateRequests []LabeledCertificatesMeshGlooSoloIov1Alpha2CertificateRequestSet,
	v1Secrets []LabeledV1SecretSet,
	clusters ...string, // the set of clusters to apply the snapshot to. only required for multicluster snapshots.
) Snapshot {
	return &snapshot{
		name: name,

		certificatesMeshGlooSoloIov1Alpha2CertificateRequests: certificatesMeshGlooSoloIov1Alpha2CertificateRequests,
		v1Secrets: v1Secrets,
		clusters:  clusters,
	}
}

// automatically partitions the input resources
// by the presence of the provided label.
func NewLabelPartitionedSnapshot(
	name,
	labelKey string, // the key by which to partition the resources

	certificatesMeshGlooSoloIov1Alpha2CertificateRequests certificates_mesh_gloo_solo_io_v1alpha2_sets.CertificateRequestSet,

	v1Secrets v1_sets.SecretSet,
	clusters ...string, // the set of clusters to apply the snapshot to. only required for multicluster snapshots.
) (Snapshot, error) {

	partitionedCertificatesMeshGlooSoloIov1Alpha2CertificateRequests, err := partitionCertificatesMeshGlooSoloIov1Alpha2CertificateRequestsByLabel(labelKey, certificatesMeshGlooSoloIov1Alpha2CertificateRequests)
	if err != nil {
		return nil, err
	}
	partitionedV1Secrets, err := partitionV1SecretsByLabel(labelKey, v1Secrets)
	if err != nil {
		return nil, err
	}

	return NewSnapshot(
		name,

		partitionedCertificatesMeshGlooSoloIov1Alpha2CertificateRequests,
		partitionedV1Secrets,
		clusters...,
	), nil
}

// simplified constructor for a snapshot
// with a single label partition (i.e. all resources share a single set of labels).
func NewSinglePartitionedSnapshot(
	name string,
	snapshotLabels map[string]string, // a single set of labels shared by all resources

	certificatesMeshGlooSoloIov1Alpha2CertificateRequests certificates_mesh_gloo_solo_io_v1alpha2_sets.CertificateRequestSet,

	v1Secrets v1_sets.SecretSet,
	clusters ...string, // the set of clusters to apply the snapshot to. only required for multicluster snapshots.
) (Snapshot, error) {

	labeledCertificatesMeshGlooSoloIov1Alpha2CertificateRequest, err := NewLabeledCertificatesMeshGlooSoloIov1Alpha2CertificateRequestSet(certificatesMeshGlooSoloIov1Alpha2CertificateRequests, snapshotLabels)
	if err != nil {
		return nil, err
	}
	labeledV1Secret, err := NewLabeledV1SecretSet(v1Secrets, snapshotLabels)
	if err != nil {
		return nil, err
	}

	return NewSnapshot(
		name,

		[]LabeledCertificatesMeshGlooSoloIov1Alpha2CertificateRequestSet{labeledCertificatesMeshGlooSoloIov1Alpha2CertificateRequest},
		[]LabeledV1SecretSet{labeledV1Secret},
		clusters...,
	), nil
}

// apply the desired resources to the cluster state; remove stale resources where necessary
func (s *snapshot) ApplyLocalCluster(ctx context.Context, cli client.Client, errHandler output.ErrorHandler) {
	var genericLists []output.ResourceList

	for _, outputSet := range s.certificatesMeshGlooSoloIov1Alpha2CertificateRequests {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.v1Secrets {
		genericLists = append(genericLists, outputSet.Generic())
	}

	output.Snapshot{
		Name:        s.name,
		ListsToSync: genericLists,
	}.SyncLocalCluster(ctx, cli, errHandler)
}

// apply the desired resources to multiple cluster states; remove stale resources where necessary
func (s *snapshot) ApplyMultiCluster(ctx context.Context, multiClusterClient multicluster.Client, errHandler output.ErrorHandler) {
	var genericLists []output.ResourceList

	for _, outputSet := range s.certificatesMeshGlooSoloIov1Alpha2CertificateRequests {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.v1Secrets {
		genericLists = append(genericLists, outputSet.Generic())
	}

	output.Snapshot{
		Name:        s.name,
		Clusters:    s.clusters,
		ListsToSync: genericLists,
	}.SyncMultiCluster(ctx, multiClusterClient, errHandler)
}

func partitionCertificatesMeshGlooSoloIov1Alpha2CertificateRequestsByLabel(labelKey string, set certificates_mesh_gloo_solo_io_v1alpha2_sets.CertificateRequestSet) ([]LabeledCertificatesMeshGlooSoloIov1Alpha2CertificateRequestSet, error) {
	setsByLabel := map[string]certificates_mesh_gloo_solo_io_v1alpha2_sets.CertificateRequestSet{}

	for _, obj := range set.List() {
		if obj.Labels == nil {
			return nil, MissingRequiredLabelError(labelKey, "CertificatesMeshGlooSoloIov1Alpha2CertificateRequest", obj)
		}
		labelValue := obj.Labels[labelKey]
		if labelValue == "" {
			return nil, MissingRequiredLabelError(labelKey, "CertificatesMeshGlooSoloIov1Alpha2CertificateRequest", obj)
		}

		setForValue, ok := setsByLabel[labelValue]
		if !ok {
			setForValue = certificates_mesh_gloo_solo_io_v1alpha2_sets.NewCertificateRequestSet()
			setsByLabel[labelValue] = setForValue
		}
		setForValue.Insert(obj)
	}

	// partition by label key
	var partitionedCertificatesMeshGlooSoloIov1Alpha2CertificateRequests []LabeledCertificatesMeshGlooSoloIov1Alpha2CertificateRequestSet

	for labelValue, setForValue := range setsByLabel {
		labels := map[string]string{labelKey: labelValue}

		partitionedSet, err := NewLabeledCertificatesMeshGlooSoloIov1Alpha2CertificateRequestSet(setForValue, labels)
		if err != nil {
			return nil, err
		}

		partitionedCertificatesMeshGlooSoloIov1Alpha2CertificateRequests = append(partitionedCertificatesMeshGlooSoloIov1Alpha2CertificateRequests, partitionedSet)
	}

	// sort for idempotency
	sort.SliceStable(partitionedCertificatesMeshGlooSoloIov1Alpha2CertificateRequests, func(i, j int) bool {
		leftLabelValue := partitionedCertificatesMeshGlooSoloIov1Alpha2CertificateRequests[i].Labels()[labelKey]
		rightLabelValue := partitionedCertificatesMeshGlooSoloIov1Alpha2CertificateRequests[j].Labels()[labelKey]
		return leftLabelValue < rightLabelValue
	})

	return partitionedCertificatesMeshGlooSoloIov1Alpha2CertificateRequests, nil
}

func partitionV1SecretsByLabel(labelKey string, set v1_sets.SecretSet) ([]LabeledV1SecretSet, error) {
	setsByLabel := map[string]v1_sets.SecretSet{}

	for _, obj := range set.List() {
		if obj.Labels == nil {
			return nil, MissingRequiredLabelError(labelKey, "V1Secret", obj)
		}
		labelValue := obj.Labels[labelKey]
		if labelValue == "" {
			return nil, MissingRequiredLabelError(labelKey, "V1Secret", obj)
		}

		setForValue, ok := setsByLabel[labelValue]
		if !ok {
			setForValue = v1_sets.NewSecretSet()
			setsByLabel[labelValue] = setForValue
		}
		setForValue.Insert(obj)
	}

	// partition by label key
	var partitionedV1Secrets []LabeledV1SecretSet

	for labelValue, setForValue := range setsByLabel {
		labels := map[string]string{labelKey: labelValue}

		partitionedSet, err := NewLabeledV1SecretSet(setForValue, labels)
		if err != nil {
			return nil, err
		}

		partitionedV1Secrets = append(partitionedV1Secrets, partitionedSet)
	}

	// sort for idempotency
	sort.SliceStable(partitionedV1Secrets, func(i, j int) bool {
		leftLabelValue := partitionedV1Secrets[i].Labels()[labelKey]
		rightLabelValue := partitionedV1Secrets[j].Labels()[labelKey]
		return leftLabelValue < rightLabelValue
	})

	return partitionedV1Secrets, nil
}

func (s snapshot) CertificatesMeshGlooSoloIov1Alpha2CertificateRequests() []LabeledCertificatesMeshGlooSoloIov1Alpha2CertificateRequestSet {
	return s.certificatesMeshGlooSoloIov1Alpha2CertificateRequests
}

func (s snapshot) V1Secrets() []LabeledV1SecretSet {
	return s.v1Secrets
}

func (s snapshot) MarshalJSON() ([]byte, error) {
	snapshotMap := map[string]interface{}{"name": s.name}

	certificatesMeshGlooSoloIov1Alpha2CertificateRequestSet := certificates_mesh_gloo_solo_io_v1alpha2_sets.NewCertificateRequestSet()
	for _, set := range s.certificatesMeshGlooSoloIov1Alpha2CertificateRequests {
		certificatesMeshGlooSoloIov1Alpha2CertificateRequestSet = certificatesMeshGlooSoloIov1Alpha2CertificateRequestSet.Union(set.Set())
	}
	snapshotMap["certificatesMeshGlooSoloIov1Alpha2CertificateRequests"] = certificatesMeshGlooSoloIov1Alpha2CertificateRequestSet.List()

	v1SecretSet := v1_sets.NewSecretSet()
	for _, set := range s.v1Secrets {
		v1SecretSet = v1SecretSet.Union(set.Set())
	}
	snapshotMap["v1Secrets"] = v1SecretSet.List()

	snapshotMap["clusters"] = s.clusters

	return json.Marshal(snapshotMap)
}

// LabeledCertificatesMeshGlooSoloIov1Alpha2CertificateRequestSet represents a set of certificatesMeshGlooSoloIov1Alpha2CertificateRequests
// which share a common set of labels.
// These labels are used to find diffs between CertificatesMeshGlooSoloIov1Alpha2CertificateRequestSets.
type LabeledCertificatesMeshGlooSoloIov1Alpha2CertificateRequestSet interface {
	// returns the set of Labels shared by this CertificatesMeshGlooSoloIov1Alpha2CertificateRequestSet
	Labels() map[string]string

	// returns the set of CertificateRequestes with the given labels
	Set() certificates_mesh_gloo_solo_io_v1alpha2_sets.CertificateRequestSet

	// converts the set to a generic format which can be applied by the Snapshot.Apply functions
	Generic() output.ResourceList
}

type labeledCertificatesMeshGlooSoloIov1Alpha2CertificateRequestSet struct {
	set    certificates_mesh_gloo_solo_io_v1alpha2_sets.CertificateRequestSet
	labels map[string]string
}

func NewLabeledCertificatesMeshGlooSoloIov1Alpha2CertificateRequestSet(set certificates_mesh_gloo_solo_io_v1alpha2_sets.CertificateRequestSet, labels map[string]string) (LabeledCertificatesMeshGlooSoloIov1Alpha2CertificateRequestSet, error) {
	// validate that each CertificateRequest contains the labels, else this is not a valid LabeledCertificatesMeshGlooSoloIov1Alpha2CertificateRequestSet
	for _, item := range set.List() {
		for k, v := range labels {
			// k=v must be present in the item
			if item.Labels[k] != v {
				return nil, eris.Errorf("internal error: %v=%v missing on CertificatesMeshGlooSoloIov1Alpha2CertificateRequest %v", k, v, item.Name)
			}
		}
	}

	return &labeledCertificatesMeshGlooSoloIov1Alpha2CertificateRequestSet{set: set, labels: labels}, nil
}

func (l *labeledCertificatesMeshGlooSoloIov1Alpha2CertificateRequestSet) Labels() map[string]string {
	return l.labels
}

func (l *labeledCertificatesMeshGlooSoloIov1Alpha2CertificateRequestSet) Set() certificates_mesh_gloo_solo_io_v1alpha2_sets.CertificateRequestSet {
	return l.set
}

func (l labeledCertificatesMeshGlooSoloIov1Alpha2CertificateRequestSet) Generic() output.ResourceList {
	var desiredResources []ezkube.Object
	for _, desired := range l.set.List() {
		desiredResources = append(desiredResources, desired)
	}

	// enable list func for garbage collection
	listFunc := func(ctx context.Context, cli client.Client) ([]ezkube.Object, error) {
		var list certificates_mesh_gloo_solo_io_v1alpha2.CertificateRequestList
		if err := cli.List(ctx, &list, client.MatchingLabels(l.labels)); err != nil {
			return nil, err
		}
		var items []ezkube.Object
		for _, item := range list.Items {
			item := item // pike
			items = append(items, &item)
		}
		return items, nil
	}

	return output.ResourceList{
		Resources:    desiredResources,
		ListFunc:     listFunc,
		ResourceKind: "CertificateRequest",
	}
}

// LabeledV1SecretSet represents a set of v1Secrets
// which share a common set of labels.
// These labels are used to find diffs between V1SecretSets.
type LabeledV1SecretSet interface {
	// returns the set of Labels shared by this V1SecretSet
	Labels() map[string]string

	// returns the set of Secretes with the given labels
	Set() v1_sets.SecretSet

	// converts the set to a generic format which can be applied by the Snapshot.Apply functions
	Generic() output.ResourceList
}

type labeledV1SecretSet struct {
	set    v1_sets.SecretSet
	labels map[string]string
}

func NewLabeledV1SecretSet(set v1_sets.SecretSet, labels map[string]string) (LabeledV1SecretSet, error) {
	// validate that each Secret contains the labels, else this is not a valid LabeledV1SecretSet
	for _, item := range set.List() {
		for k, v := range labels {
			// k=v must be present in the item
			if item.Labels[k] != v {
				return nil, eris.Errorf("internal error: %v=%v missing on V1Secret %v", k, v, item.Name)
			}
		}
	}

	return &labeledV1SecretSet{set: set, labels: labels}, nil
}

func (l *labeledV1SecretSet) Labels() map[string]string {
	return l.labels
}

func (l *labeledV1SecretSet) Set() v1_sets.SecretSet {
	return l.set
}

func (l labeledV1SecretSet) Generic() output.ResourceList {
	var desiredResources []ezkube.Object
	for _, desired := range l.set.List() {
		desiredResources = append(desiredResources, desired)
	}

	// enable list func for garbage collection
	listFunc := func(ctx context.Context, cli client.Client) ([]ezkube.Object, error) {
		var list v1.SecretList
		if err := cli.List(ctx, &list, client.MatchingLabels(l.labels)); err != nil {
			return nil, err
		}
		var items []ezkube.Object
		for _, item := range list.Items {
			item := item // pike
			items = append(items, &item)
		}
		return items, nil
	}

	return output.ResourceList{
		Resources:    desiredResources,
		ListFunc:     listFunc,
		ResourceKind: "Secret",
	}
}

type builder struct {
	ctx      context.Context
	name     string
	clusters []string

	certificatesMeshGlooSoloIov1Alpha2CertificateRequests certificates_mesh_gloo_solo_io_v1alpha2_sets.CertificateRequestSet

	v1Secrets v1_sets.SecretSet
}

func NewBuilder(ctx context.Context, name string) *builder {
	return &builder{
		ctx:  ctx,
		name: name,

		certificatesMeshGlooSoloIov1Alpha2CertificateRequests: certificates_mesh_gloo_solo_io_v1alpha2_sets.NewCertificateRequestSet(),

		v1Secrets: v1_sets.NewSecretSet(),
	}
}

// the output Builder uses a builder pattern to allow
// iteratively collecting outputs before producing a final snapshot
type Builder interface {

	// add CertificatesMeshGlooSoloIov1Alpha2CertificateRequests to the collected outputs
	AddCertificatesMeshGlooSoloIov1Alpha2CertificateRequests(certificatesMeshGlooSoloIov1Alpha2CertificateRequests ...*certificates_mesh_gloo_solo_io_v1alpha2.CertificateRequest)

	// get the collected CertificatesMeshGlooSoloIov1Alpha2CertificateRequests
	GetCertificatesMeshGlooSoloIov1Alpha2CertificateRequests() certificates_mesh_gloo_solo_io_v1alpha2_sets.CertificateRequestSet

	// add V1Secrets to the collected outputs
	AddV1Secrets(v1Secrets ...*v1.Secret)

	// get the collected V1Secrets
	GetV1Secrets() v1_sets.SecretSet

	// build the collected outputs into a label-partitioned snapshot
	BuildLabelPartitionedSnapshot(labelKey string) (Snapshot, error)

	// build the collected outputs into a snapshot with a single partition
	BuildSinglePartitionedSnapshot(snapshotLabels map[string]string) (Snapshot, error)

	// add a cluster to the collected clusters.
	// this can be used to collect clusters for use with MultiCluster snapshots.
	AddCluster(cluster string)

	// returns the set of clusters currently stored in this builder
	Clusters() []string

	// merge all the resources from another Builder into this one
	Merge(other Builder)

	// create a clone of this builder (deepcopying all resources)
	Clone() Builder

	// return the difference between the snapshot in this builder's and another
	Delta(newSnap Builder) output.SnapshotDelta
}

func (b *builder) AddCertificatesMeshGlooSoloIov1Alpha2CertificateRequests(certificatesMeshGlooSoloIov1Alpha2CertificateRequests ...*certificates_mesh_gloo_solo_io_v1alpha2.CertificateRequest) {
	for _, obj := range certificatesMeshGlooSoloIov1Alpha2CertificateRequests {
		if obj == nil {
			continue
		}
		contextutils.LoggerFrom(b.ctx).Debugf("added output CertificatesMeshGlooSoloIov1Alpha2CertificateRequest %v", sets.Key(obj))
		b.certificatesMeshGlooSoloIov1Alpha2CertificateRequests.Insert(obj)
	}
}
func (b *builder) AddV1Secrets(v1Secrets ...*v1.Secret) {
	for _, obj := range v1Secrets {
		if obj == nil {
			continue
		}
		contextutils.LoggerFrom(b.ctx).Debugf("added output V1Secret %v", sets.Key(obj))
		b.v1Secrets.Insert(obj)
	}
}

func (b *builder) GetCertificatesMeshGlooSoloIov1Alpha2CertificateRequests() certificates_mesh_gloo_solo_io_v1alpha2_sets.CertificateRequestSet {
	return b.certificatesMeshGlooSoloIov1Alpha2CertificateRequests
}

func (b *builder) GetV1Secrets() v1_sets.SecretSet {
	return b.v1Secrets
}

func (b *builder) BuildLabelPartitionedSnapshot(labelKey string) (Snapshot, error) {
	return NewLabelPartitionedSnapshot(
		b.name,
		labelKey,

		b.certificatesMeshGlooSoloIov1Alpha2CertificateRequests,

		b.v1Secrets,
		b.clusters...,
	)
}

func (b *builder) BuildSinglePartitionedSnapshot(snapshotLabels map[string]string) (Snapshot, error) {
	return NewSinglePartitionedSnapshot(
		b.name,
		snapshotLabels,

		b.certificatesMeshGlooSoloIov1Alpha2CertificateRequests,

		b.v1Secrets,
		b.clusters...,
	)
}

func (b *builder) AddCluster(cluster string) {
	b.clusters = append(b.clusters, cluster)
}

func (b *builder) Clusters() []string {
	return b.clusters
}

func (b *builder) Merge(other Builder) {
	if other == nil {
		return
	}

	b.AddCertificatesMeshGlooSoloIov1Alpha2CertificateRequests(other.GetCertificatesMeshGlooSoloIov1Alpha2CertificateRequests().List()...)

	b.AddV1Secrets(other.GetV1Secrets().List()...)
	for _, cluster := range other.Clusters() {
		b.AddCluster(cluster)
	}
}

func (b *builder) Clone() Builder {
	if b == nil {
		return nil
	}
	clone := NewBuilder(b.ctx, b.name)

	for _, certificatesMeshGlooSoloIov1Alpha2CertificateRequest := range b.GetCertificatesMeshGlooSoloIov1Alpha2CertificateRequests().List() {
		clone.AddCertificatesMeshGlooSoloIov1Alpha2CertificateRequests(certificatesMeshGlooSoloIov1Alpha2CertificateRequest.DeepCopy())
	}

	for _, v1Secret := range b.GetV1Secrets().List() {
		clone.AddV1Secrets(v1Secret.DeepCopy())
	}
	for _, cluster := range b.Clusters() {
		clone.AddCluster(cluster)
	}
	return clone
}

func (b *builder) Delta(other Builder) output.SnapshotDelta {
	delta := output.SnapshotDelta{}
	if b == nil {
		return delta
	}

	// calcualte delta between CertificateRequests
	certificateRequestDelta := b.GetCertificateRequests().Delta(other.GetCertificateRequests())
	certificateRequestGvk := schema.GroupVersionKind{
		Group:   "certificates.mesh.gloo.solo.io",
		Version: "v1alpha2",
		Kind:    "CertificateRequest",
	}
	delta.AddInserted(certificateRequestGvk, certificateRequestDelta.Inserted)
	delta.AddRemoved(certificateRequestGvk, certificateRequestDelta.Removed)

	// calcualte delta between Secrets
	secretDelta := b.GetSecrets().Delta(other.GetSecrets())
	secretGvk := schema.GroupVersionKind{
		Group:   "",
		Version: "v1",
		Kind:    "Secret",
	}
	delta.AddInserted(secretGvk, secretDelta.Inserted)
	delta.AddRemoved(secretGvk, secretDelta.Removed)
	return delta
}

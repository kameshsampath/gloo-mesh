// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./snapshot.go -destination mocks/snapshot.go

// Definitions for Output Snapshots
package output

import (
	"context"
	"encoding/json"
	"sort"

	"github.com/solo-io/go-utils/contextutils"

	"github.com/rotisserie/eris"
	"github.com/solo-io/skv2/contrib/pkg/output"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	"sigs.k8s.io/controller-runtime/pkg/client"

	discovery_smh_solo_io_v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2"
	discovery_smh_solo_io_v1alpha2_sets "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2/sets"
)

// this error can occur if constructing a Partitioned Snapshot from a resource
// that is missing the partition label
var MissingRequiredLabelError = func(labelKey, resourceKind string, obj ezkube.ResourceId) error {
	return eris.Errorf("expected label %v not on labels of %v %v", labelKey, resourceKind, sets.Key(obj))
}

// the snapshot of output resources produced by a translation
type Snapshot interface {

	// return the set of MeshServices with a given set of labels
	MeshServices() []LabeledMeshServiceSet
	// return the set of MeshWorkloads with a given set of labels
	MeshWorkloads() []LabeledMeshWorkloadSet
	// return the set of Meshes with a given set of labels
	Meshes() []LabeledMeshSet

	// apply the snapshot to the local cluster, garbage collecting stale resources
	ApplyLocalCluster(ctx context.Context, clusterClient client.Client, errHandler output.ErrorHandler)

	// apply resources from the snapshot across multiple clusters, garbage collecting stale resources
	ApplyMultiCluster(ctx context.Context, multiClusterClient multicluster.Client, errHandler output.ErrorHandler)

	// serialize the entire snapshot as JSON
	MarshalJSON() ([]byte, error)
}

type snapshot struct {
	name string

	meshServices  []LabeledMeshServiceSet
	meshWorkloads []LabeledMeshWorkloadSet
	meshes        []LabeledMeshSet
}

func NewSnapshot(
	name string,

	meshServices []LabeledMeshServiceSet,
	meshWorkloads []LabeledMeshWorkloadSet,
	meshes []LabeledMeshSet,

) Snapshot {
	return &snapshot{
		name: name,

		meshServices:  meshServices,
		meshWorkloads: meshWorkloads,
		meshes:        meshes,
	}
}

// automatically partitions the input resources
// by the presence of the provided label.
func NewLabelPartitionedSnapshot(
	name,
	labelKey string, // the key by which to partition the resources

	meshServices discovery_smh_solo_io_v1alpha2_sets.MeshServiceSet,
	meshWorkloads discovery_smh_solo_io_v1alpha2_sets.MeshWorkloadSet,
	meshes discovery_smh_solo_io_v1alpha2_sets.MeshSet,

) (Snapshot, error) {

	partitionedMeshServices, err := partitionMeshServicesByLabel(labelKey, meshServices)
	if err != nil {
		return nil, err
	}
	partitionedMeshWorkloads, err := partitionMeshWorkloadsByLabel(labelKey, meshWorkloads)
	if err != nil {
		return nil, err
	}
	partitionedMeshes, err := partitionMeshesByLabel(labelKey, meshes)
	if err != nil {
		return nil, err
	}

	return NewSnapshot(
		name,

		partitionedMeshServices,
		partitionedMeshWorkloads,
		partitionedMeshes,
	), nil
}

// simplified constructor for a snapshot
// with a single label partition (i.e. all resources share a single set of labels).
func NewSinglePartitionedSnapshot(
	name string,
	snapshotLabels map[string]string, // a single set of labels shared by all resources

	meshServices discovery_smh_solo_io_v1alpha2_sets.MeshServiceSet,
	meshWorkloads discovery_smh_solo_io_v1alpha2_sets.MeshWorkloadSet,
	meshes discovery_smh_solo_io_v1alpha2_sets.MeshSet,

) (Snapshot, error) {

	labeledMeshServices, err := NewLabeledMeshServiceSet(meshServices, snapshotLabels)
	if err != nil {
		return nil, err
	}
	labeledMeshWorkloads, err := NewLabeledMeshWorkloadSet(meshWorkloads, snapshotLabels)
	if err != nil {
		return nil, err
	}
	labeledMeshes, err := NewLabeledMeshSet(meshes, snapshotLabels)
	if err != nil {
		return nil, err
	}

	return NewSnapshot(
		name,

		[]LabeledMeshServiceSet{labeledMeshServices},
		[]LabeledMeshWorkloadSet{labeledMeshWorkloads},
		[]LabeledMeshSet{labeledMeshes},
	), nil
}

// apply the desired resources to the cluster state; remove stale resources where necessary
func (s *snapshot) ApplyLocalCluster(ctx context.Context, cli client.Client, errHandler output.ErrorHandler) {
	var genericLists []output.ResourceList

	for _, outputSet := range s.meshServices {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.meshWorkloads {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.meshes {
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

	for _, outputSet := range s.meshServices {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.meshWorkloads {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.meshes {
		genericLists = append(genericLists, outputSet.Generic())
	}

	output.Snapshot{
		Name:        s.name,
		ListsToSync: genericLists,
	}.SyncMultiCluster(ctx, multiClusterClient, errHandler)
}

func partitionMeshServicesByLabel(labelKey string, set discovery_smh_solo_io_v1alpha2_sets.MeshServiceSet) ([]LabeledMeshServiceSet, error) {
	setsByLabel := map[string]discovery_smh_solo_io_v1alpha2_sets.MeshServiceSet{}

	for _, obj := range set.List() {
		if obj.Labels == nil {
			return nil, MissingRequiredLabelError(labelKey, "MeshService", obj)
		}
		labelValue := obj.Labels[labelKey]
		if labelValue == "" {
			return nil, MissingRequiredLabelError(labelKey, "MeshService", obj)
		}

		setForValue, ok := setsByLabel[labelValue]
		if !ok {
			setForValue = discovery_smh_solo_io_v1alpha2_sets.NewMeshServiceSet()
			setsByLabel[labelValue] = setForValue
		}
		setForValue.Insert(obj)
	}

	// partition by label key
	var partitionedMeshServices []LabeledMeshServiceSet

	for labelValue, setForValue := range setsByLabel {
		labels := map[string]string{labelKey: labelValue}

		partitionedSet, err := NewLabeledMeshServiceSet(setForValue, labels)
		if err != nil {
			return nil, err
		}

		partitionedMeshServices = append(partitionedMeshServices, partitionedSet)
	}

	// sort for idempotency
	sort.SliceStable(partitionedMeshServices, func(i, j int) bool {
		leftLabelValue := partitionedMeshServices[i].Labels()[labelKey]
		rightLabelValue := partitionedMeshServices[j].Labels()[labelKey]
		return leftLabelValue < rightLabelValue
	})

	return partitionedMeshServices, nil
}

func partitionMeshWorkloadsByLabel(labelKey string, set discovery_smh_solo_io_v1alpha2_sets.MeshWorkloadSet) ([]LabeledMeshWorkloadSet, error) {
	setsByLabel := map[string]discovery_smh_solo_io_v1alpha2_sets.MeshWorkloadSet{}

	for _, obj := range set.List() {
		if obj.Labels == nil {
			return nil, MissingRequiredLabelError(labelKey, "MeshWorkload", obj)
		}
		labelValue := obj.Labels[labelKey]
		if labelValue == "" {
			return nil, MissingRequiredLabelError(labelKey, "MeshWorkload", obj)
		}

		setForValue, ok := setsByLabel[labelValue]
		if !ok {
			setForValue = discovery_smh_solo_io_v1alpha2_sets.NewMeshWorkloadSet()
			setsByLabel[labelValue] = setForValue
		}
		setForValue.Insert(obj)
	}

	// partition by label key
	var partitionedMeshWorkloads []LabeledMeshWorkloadSet

	for labelValue, setForValue := range setsByLabel {
		labels := map[string]string{labelKey: labelValue}

		partitionedSet, err := NewLabeledMeshWorkloadSet(setForValue, labels)
		if err != nil {
			return nil, err
		}

		partitionedMeshWorkloads = append(partitionedMeshWorkloads, partitionedSet)
	}

	// sort for idempotency
	sort.SliceStable(partitionedMeshWorkloads, func(i, j int) bool {
		leftLabelValue := partitionedMeshWorkloads[i].Labels()[labelKey]
		rightLabelValue := partitionedMeshWorkloads[j].Labels()[labelKey]
		return leftLabelValue < rightLabelValue
	})

	return partitionedMeshWorkloads, nil
}

func partitionMeshesByLabel(labelKey string, set discovery_smh_solo_io_v1alpha2_sets.MeshSet) ([]LabeledMeshSet, error) {
	setsByLabel := map[string]discovery_smh_solo_io_v1alpha2_sets.MeshSet{}

	for _, obj := range set.List() {
		if obj.Labels == nil {
			return nil, MissingRequiredLabelError(labelKey, "Mesh", obj)
		}
		labelValue := obj.Labels[labelKey]
		if labelValue == "" {
			return nil, MissingRequiredLabelError(labelKey, "Mesh", obj)
		}

		setForValue, ok := setsByLabel[labelValue]
		if !ok {
			setForValue = discovery_smh_solo_io_v1alpha2_sets.NewMeshSet()
			setsByLabel[labelValue] = setForValue
		}
		setForValue.Insert(obj)
	}

	// partition by label key
	var partitionedMeshes []LabeledMeshSet

	for labelValue, setForValue := range setsByLabel {
		labels := map[string]string{labelKey: labelValue}

		partitionedSet, err := NewLabeledMeshSet(setForValue, labels)
		if err != nil {
			return nil, err
		}

		partitionedMeshes = append(partitionedMeshes, partitionedSet)
	}

	// sort for idempotency
	sort.SliceStable(partitionedMeshes, func(i, j int) bool {
		leftLabelValue := partitionedMeshes[i].Labels()[labelKey]
		rightLabelValue := partitionedMeshes[j].Labels()[labelKey]
		return leftLabelValue < rightLabelValue
	})

	return partitionedMeshes, nil
}

func (s snapshot) MeshServices() []LabeledMeshServiceSet {
	return s.meshServices
}

func (s snapshot) MeshWorkloads() []LabeledMeshWorkloadSet {
	return s.meshWorkloads
}

func (s snapshot) Meshes() []LabeledMeshSet {
	return s.meshes
}

func (s snapshot) MarshalJSON() ([]byte, error) {
	snapshotMap := map[string]interface{}{"name": s.name}

	meshServiceSet := discovery_smh_solo_io_v1alpha2_sets.NewMeshServiceSet()
	for _, set := range s.meshServices {
		meshServiceSet = meshServiceSet.Union(set.Set())
	}
	snapshotMap["meshServices"] = meshServiceSet.List()
	meshWorkloadSet := discovery_smh_solo_io_v1alpha2_sets.NewMeshWorkloadSet()
	for _, set := range s.meshWorkloads {
		meshWorkloadSet = meshWorkloadSet.Union(set.Set())
	}
	snapshotMap["meshWorkloads"] = meshWorkloadSet.List()
	meshSet := discovery_smh_solo_io_v1alpha2_sets.NewMeshSet()
	for _, set := range s.meshes {
		meshSet = meshSet.Union(set.Set())
	}
	snapshotMap["meshes"] = meshSet.List()

	return json.Marshal(snapshotMap)
}

// LabeledMeshServiceSet represents a set of meshServices
// which share a common set of labels.
// These labels are used to find diffs between MeshServiceSets.
type LabeledMeshServiceSet interface {
	// returns the set of Labels shared by this MeshServiceSet
	Labels() map[string]string

	// returns the set of MeshServicees with the given labels
	Set() discovery_smh_solo_io_v1alpha2_sets.MeshServiceSet

	// converts the set to a generic format which can be applied by the Snapshot.Apply functions
	Generic() output.ResourceList
}

type labeledMeshServiceSet struct {
	set    discovery_smh_solo_io_v1alpha2_sets.MeshServiceSet
	labels map[string]string
}

func NewLabeledMeshServiceSet(set discovery_smh_solo_io_v1alpha2_sets.MeshServiceSet, labels map[string]string) (LabeledMeshServiceSet, error) {
	// validate that each MeshService contains the labels, else this is not a valid LabeledMeshServiceSet
	for _, item := range set.List() {
		for k, v := range labels {
			// k=v must be present in the item
			if item.Labels[k] != v {
				return nil, eris.Errorf("internal error: %v=%v missing on MeshService %v", k, v, item.Name)
			}
		}
	}

	return &labeledMeshServiceSet{set: set, labels: labels}, nil
}

func (l *labeledMeshServiceSet) Labels() map[string]string {
	return l.labels
}

func (l *labeledMeshServiceSet) Set() discovery_smh_solo_io_v1alpha2_sets.MeshServiceSet {
	return l.set
}

func (l labeledMeshServiceSet) Generic() output.ResourceList {
	var desiredResources []ezkube.Object
	for _, desired := range l.set.List() {
		desiredResources = append(desiredResources, desired)
	}

	// enable list func for garbage collection
	listFunc := func(ctx context.Context, cli client.Client) ([]ezkube.Object, error) {
		var list discovery_smh_solo_io_v1alpha2.MeshServiceList
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
		ResourceKind: "MeshService",
	}
}

// LabeledMeshWorkloadSet represents a set of meshWorkloads
// which share a common set of labels.
// These labels are used to find diffs between MeshWorkloadSets.
type LabeledMeshWorkloadSet interface {
	// returns the set of Labels shared by this MeshWorkloadSet
	Labels() map[string]string

	// returns the set of MeshWorkloades with the given labels
	Set() discovery_smh_solo_io_v1alpha2_sets.MeshWorkloadSet

	// converts the set to a generic format which can be applied by the Snapshot.Apply functions
	Generic() output.ResourceList
}

type labeledMeshWorkloadSet struct {
	set    discovery_smh_solo_io_v1alpha2_sets.MeshWorkloadSet
	labels map[string]string
}

func NewLabeledMeshWorkloadSet(set discovery_smh_solo_io_v1alpha2_sets.MeshWorkloadSet, labels map[string]string) (LabeledMeshWorkloadSet, error) {
	// validate that each MeshWorkload contains the labels, else this is not a valid LabeledMeshWorkloadSet
	for _, item := range set.List() {
		for k, v := range labels {
			// k=v must be present in the item
			if item.Labels[k] != v {
				return nil, eris.Errorf("internal error: %v=%v missing on MeshWorkload %v", k, v, item.Name)
			}
		}
	}

	return &labeledMeshWorkloadSet{set: set, labels: labels}, nil
}

func (l *labeledMeshWorkloadSet) Labels() map[string]string {
	return l.labels
}

func (l *labeledMeshWorkloadSet) Set() discovery_smh_solo_io_v1alpha2_sets.MeshWorkloadSet {
	return l.set
}

func (l labeledMeshWorkloadSet) Generic() output.ResourceList {
	var desiredResources []ezkube.Object
	for _, desired := range l.set.List() {
		desiredResources = append(desiredResources, desired)
	}

	// enable list func for garbage collection
	listFunc := func(ctx context.Context, cli client.Client) ([]ezkube.Object, error) {
		var list discovery_smh_solo_io_v1alpha2.MeshWorkloadList
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
		ResourceKind: "MeshWorkload",
	}
}

// LabeledMeshSet represents a set of meshes
// which share a common set of labels.
// These labels are used to find diffs between MeshSets.
type LabeledMeshSet interface {
	// returns the set of Labels shared by this MeshSet
	Labels() map[string]string

	// returns the set of Meshes with the given labels
	Set() discovery_smh_solo_io_v1alpha2_sets.MeshSet

	// converts the set to a generic format which can be applied by the Snapshot.Apply functions
	Generic() output.ResourceList
}

type labeledMeshSet struct {
	set    discovery_smh_solo_io_v1alpha2_sets.MeshSet
	labels map[string]string
}

func NewLabeledMeshSet(set discovery_smh_solo_io_v1alpha2_sets.MeshSet, labels map[string]string) (LabeledMeshSet, error) {
	// validate that each Mesh contains the labels, else this is not a valid LabeledMeshSet
	for _, item := range set.List() {
		for k, v := range labels {
			// k=v must be present in the item
			if item.Labels[k] != v {
				return nil, eris.Errorf("internal error: %v=%v missing on Mesh %v", k, v, item.Name)
			}
		}
	}

	return &labeledMeshSet{set: set, labels: labels}, nil
}

func (l *labeledMeshSet) Labels() map[string]string {
	return l.labels
}

func (l *labeledMeshSet) Set() discovery_smh_solo_io_v1alpha2_sets.MeshSet {
	return l.set
}

func (l labeledMeshSet) Generic() output.ResourceList {
	var desiredResources []ezkube.Object
	for _, desired := range l.set.List() {
		desiredResources = append(desiredResources, desired)
	}

	// enable list func for garbage collection
	listFunc := func(ctx context.Context, cli client.Client) ([]ezkube.Object, error) {
		var list discovery_smh_solo_io_v1alpha2.MeshList
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
		ResourceKind: "Mesh",
	}
}

type builder struct {
	ctx  context.Context
	name string

	meshServices  discovery_smh_solo_io_v1alpha2_sets.MeshServiceSet
	meshWorkloads discovery_smh_solo_io_v1alpha2_sets.MeshWorkloadSet
	meshes        discovery_smh_solo_io_v1alpha2_sets.MeshSet
}

func NewBuilder(ctx context.Context, name string) *builder {
	return &builder{
		ctx:  ctx,
		name: name,

		meshServices:  discovery_smh_solo_io_v1alpha2_sets.NewMeshServiceSet(),
		meshWorkloads: discovery_smh_solo_io_v1alpha2_sets.NewMeshWorkloadSet(),
		meshes:        discovery_smh_solo_io_v1alpha2_sets.NewMeshSet(),
	}
}

// the output Builder uses a builder pattern to allow
// iteratively collecting outputs before producing a final snapshot
type Builder interface {

	// add MeshServices to the collected outputs
	AddMeshServices(meshServices ...*discovery_smh_solo_io_v1alpha2.MeshService)

	// get the collected MeshServices
	GetMeshServices() discovery_smh_solo_io_v1alpha2_sets.MeshServiceSet

	// add MeshWorkloads to the collected outputs
	AddMeshWorkloads(meshWorkloads ...*discovery_smh_solo_io_v1alpha2.MeshWorkload)

	// get the collected MeshWorkloads
	GetMeshWorkloads() discovery_smh_solo_io_v1alpha2_sets.MeshWorkloadSet

	// add Meshes to the collected outputs
	AddMeshes(meshes ...*discovery_smh_solo_io_v1alpha2.Mesh)

	// get the collected Meshes
	GetMeshes() discovery_smh_solo_io_v1alpha2_sets.MeshSet

	// build the collected outputs into a label-partitioned snapshot
	BuildLabelPartitionedSnapshot(labelKey string) (Snapshot, error)

	// build the collected outputs into a snapshot with a single partition
	BuildSinglePartitionedSnapshot(snapshotLabels map[string]string) (Snapshot, error)
}

func (b *builder) AddMeshServices(meshServices ...*discovery_smh_solo_io_v1alpha2.MeshService) {
	for _, obj := range meshServices {
		contextutils.LoggerFrom(b.ctx).Debugf("added output MeshService %v", sets.Key(obj))
	}

	b.meshServices.Insert(meshServices...)
}
func (b *builder) AddMeshWorkloads(meshWorkloads ...*discovery_smh_solo_io_v1alpha2.MeshWorkload) {
	for _, obj := range meshWorkloads {
		contextutils.LoggerFrom(b.ctx).Debugf("added output MeshWorkload %v", sets.Key(obj))
	}

	b.meshWorkloads.Insert(meshWorkloads...)
}
func (b *builder) AddMeshes(meshes ...*discovery_smh_solo_io_v1alpha2.Mesh) {
	for _, obj := range meshes {
		contextutils.LoggerFrom(b.ctx).Debugf("added output Mesh %v", sets.Key(obj))
	}

	b.meshes.Insert(meshes...)
}

func (b *builder) GetMeshServices() discovery_smh_solo_io_v1alpha2_sets.MeshServiceSet {
	return b.meshServices
}
func (b *builder) GetMeshWorkloads() discovery_smh_solo_io_v1alpha2_sets.MeshWorkloadSet {
	return b.meshWorkloads
}
func (b *builder) GetMeshes() discovery_smh_solo_io_v1alpha2_sets.MeshSet {
	return b.meshes
}

func (b *builder) BuildLabelPartitionedSnapshot(labelKey string) (Snapshot, error) {
	return NewLabelPartitionedSnapshot(
		b.name,
		labelKey,

		b.meshServices,
		b.meshWorkloads,
		b.meshes,
	)
}

func (b *builder) BuildSinglePartitionedSnapshot(snapshotLabels map[string]string) (Snapshot, error) {
	return NewSinglePartitionedSnapshot(
		b.name,
		snapshotLabels,

		b.meshServices,
		b.meshWorkloads,
		b.meshes,
	)
}

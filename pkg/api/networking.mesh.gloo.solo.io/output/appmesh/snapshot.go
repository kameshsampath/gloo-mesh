// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./snapshot.go -destination mocks/snapshot.go

// Definitions for Output Snapshots
package appmesh

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

	appmesh_k8s_aws_v1beta2 "github.com/aws/aws-app-mesh-controller-for-k8s/apis/appmesh/v1beta2"
	appmesh_k8s_aws_v1beta2_sets "github.com/solo-io/external-apis/pkg/api/appmesh/appmesh.k8s.aws/v1beta2/sets"
)

// this error can occur if constructing a Partitioned Snapshot from a resource
// that is missing the partition label
var MissingRequiredLabelError = func(labelKey, resourceKind string, obj ezkube.ResourceId) error {
	return eris.Errorf("expected label %v not on labels of %v %v", labelKey, resourceKind, sets.Key(obj))
}

// the snapshot of output resources produced by a translation
type Snapshot interface {

	// return the set of AppmeshK8SAwsv1Beta2VirtualServices with a given set of labels
	AppmeshK8SAwsv1Beta2VirtualServices() []LabeledAppmeshK8SAwsv1Beta2VirtualServiceSet
	// return the set of AppmeshK8SAwsv1Beta2VirtualNodes with a given set of labels
	AppmeshK8SAwsv1Beta2VirtualNodes() []LabeledAppmeshK8SAwsv1Beta2VirtualNodeSet
	// return the set of AppmeshK8SAwsv1Beta2VirtualRouters with a given set of labels
	AppmeshK8SAwsv1Beta2VirtualRouters() []LabeledAppmeshK8SAwsv1Beta2VirtualRouterSet

	// apply the snapshot to the local cluster, garbage collecting stale resources
	ApplyLocalCluster(ctx context.Context, clusterClient client.Client, errHandler output.ErrorHandler)

	// apply resources from the snapshot across multiple clusters, garbage collecting stale resources
	ApplyMultiCluster(ctx context.Context, multiClusterClient multicluster.Client, errHandler output.ErrorHandler)

	// serialize the entire snapshot as JSON
	MarshalJSON() ([]byte, error)
}

type snapshot struct {
	name string

	appmeshK8SAwsv1Beta2VirtualServices []LabeledAppmeshK8SAwsv1Beta2VirtualServiceSet
	appmeshK8SAwsv1Beta2VirtualNodes    []LabeledAppmeshK8SAwsv1Beta2VirtualNodeSet
	appmeshK8SAwsv1Beta2VirtualRouters  []LabeledAppmeshK8SAwsv1Beta2VirtualRouterSet
	clusters                            []string
}

func NewSnapshot(
	name string,

	appmeshK8SAwsv1Beta2VirtualServices []LabeledAppmeshK8SAwsv1Beta2VirtualServiceSet,
	appmeshK8SAwsv1Beta2VirtualNodes []LabeledAppmeshK8SAwsv1Beta2VirtualNodeSet,
	appmeshK8SAwsv1Beta2VirtualRouters []LabeledAppmeshK8SAwsv1Beta2VirtualRouterSet,
	clusters ...string, // the set of clusters to apply the snapshot to. only required for multicluster snapshots.
) Snapshot {
	return &snapshot{
		name: name,

		appmeshK8SAwsv1Beta2VirtualServices: appmeshK8SAwsv1Beta2VirtualServices,
		appmeshK8SAwsv1Beta2VirtualNodes:    appmeshK8SAwsv1Beta2VirtualNodes,
		appmeshK8SAwsv1Beta2VirtualRouters:  appmeshK8SAwsv1Beta2VirtualRouters,
		clusters:                            clusters,
	}
}

// automatically partitions the input resources
// by the presence of the provided label.
func NewLabelPartitionedSnapshot(
	name,
	labelKey string, // the key by which to partition the resources

	appmeshK8SAwsv1Beta2VirtualServices appmesh_k8s_aws_v1beta2_sets.VirtualServiceSet,
	appmeshK8SAwsv1Beta2VirtualNodes appmesh_k8s_aws_v1beta2_sets.VirtualNodeSet,
	appmeshK8SAwsv1Beta2VirtualRouters appmesh_k8s_aws_v1beta2_sets.VirtualRouterSet,
	clusters ...string, // the set of clusters to apply the snapshot to. only required for multicluster snapshots.
) (Snapshot, error) {

	partitionedAppmeshK8SAwsv1Beta2VirtualServices, err := partitionAppmeshK8SAwsv1Beta2VirtualServicesByLabel(labelKey, appmeshK8SAwsv1Beta2VirtualServices)
	if err != nil {
		return nil, err
	}
	partitionedAppmeshK8SAwsv1Beta2VirtualNodes, err := partitionAppmeshK8SAwsv1Beta2VirtualNodesByLabel(labelKey, appmeshK8SAwsv1Beta2VirtualNodes)
	if err != nil {
		return nil, err
	}
	partitionedAppmeshK8SAwsv1Beta2VirtualRouters, err := partitionAppmeshK8SAwsv1Beta2VirtualRoutersByLabel(labelKey, appmeshK8SAwsv1Beta2VirtualRouters)
	if err != nil {
		return nil, err
	}

	return NewSnapshot(
		name,

		partitionedAppmeshK8SAwsv1Beta2VirtualServices,
		partitionedAppmeshK8SAwsv1Beta2VirtualNodes,
		partitionedAppmeshK8SAwsv1Beta2VirtualRouters,
		clusters...,
	), nil
}

// simplified constructor for a snapshot
// with a single label partition (i.e. all resources share a single set of labels).
func NewSinglePartitionedSnapshot(
	name string,
	snapshotLabels map[string]string, // a single set of labels shared by all resources

	appmeshK8SAwsv1Beta2VirtualServices appmesh_k8s_aws_v1beta2_sets.VirtualServiceSet,
	appmeshK8SAwsv1Beta2VirtualNodes appmesh_k8s_aws_v1beta2_sets.VirtualNodeSet,
	appmeshK8SAwsv1Beta2VirtualRouters appmesh_k8s_aws_v1beta2_sets.VirtualRouterSet,
	clusters ...string, // the set of clusters to apply the snapshot to. only required for multicluster snapshots.
) (Snapshot, error) {

	labeledAppmeshK8SAwsv1Beta2VirtualService, err := NewLabeledAppmeshK8SAwsv1Beta2VirtualServiceSet(appmeshK8SAwsv1Beta2VirtualServices, snapshotLabels)
	if err != nil {
		return nil, err
	}
	labeledAppmeshK8SAwsv1Beta2VirtualNode, err := NewLabeledAppmeshK8SAwsv1Beta2VirtualNodeSet(appmeshK8SAwsv1Beta2VirtualNodes, snapshotLabels)
	if err != nil {
		return nil, err
	}
	labeledAppmeshK8SAwsv1Beta2VirtualRouter, err := NewLabeledAppmeshK8SAwsv1Beta2VirtualRouterSet(appmeshK8SAwsv1Beta2VirtualRouters, snapshotLabels)
	if err != nil {
		return nil, err
	}

	return NewSnapshot(
		name,

		[]LabeledAppmeshK8SAwsv1Beta2VirtualServiceSet{labeledAppmeshK8SAwsv1Beta2VirtualService},
		[]LabeledAppmeshK8SAwsv1Beta2VirtualNodeSet{labeledAppmeshK8SAwsv1Beta2VirtualNode},
		[]LabeledAppmeshK8SAwsv1Beta2VirtualRouterSet{labeledAppmeshK8SAwsv1Beta2VirtualRouter},
		clusters...,
	), nil
}

// apply the desired resources to the cluster state; remove stale resources where necessary
func (s *snapshot) ApplyLocalCluster(ctx context.Context, cli client.Client, errHandler output.ErrorHandler) {
	var genericLists []output.ResourceList

	for _, outputSet := range s.appmeshK8SAwsv1Beta2VirtualServices {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.appmeshK8SAwsv1Beta2VirtualNodes {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.appmeshK8SAwsv1Beta2VirtualRouters {
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

	for _, outputSet := range s.appmeshK8SAwsv1Beta2VirtualServices {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.appmeshK8SAwsv1Beta2VirtualNodes {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.appmeshK8SAwsv1Beta2VirtualRouters {
		genericLists = append(genericLists, outputSet.Generic())
	}

	output.Snapshot{
		Name:        s.name,
		Clusters:    s.clusters,
		ListsToSync: genericLists,
	}.SyncMultiCluster(ctx, multiClusterClient, errHandler)
}

func partitionAppmeshK8SAwsv1Beta2VirtualServicesByLabel(labelKey string, set appmesh_k8s_aws_v1beta2_sets.VirtualServiceSet) ([]LabeledAppmeshK8SAwsv1Beta2VirtualServiceSet, error) {
	setsByLabel := map[string]appmesh_k8s_aws_v1beta2_sets.VirtualServiceSet{}

	for _, obj := range set.List() {
		if obj.Labels == nil {
			return nil, MissingRequiredLabelError(labelKey, "AppmeshK8SAwsv1Beta2VirtualService", obj)
		}
		labelValue := obj.Labels[labelKey]
		if labelValue == "" {
			return nil, MissingRequiredLabelError(labelKey, "AppmeshK8SAwsv1Beta2VirtualService", obj)
		}

		setForValue, ok := setsByLabel[labelValue]
		if !ok {
			setForValue = appmesh_k8s_aws_v1beta2_sets.NewVirtualServiceSet()
			setsByLabel[labelValue] = setForValue
		}
		setForValue.Insert(obj)
	}

	// partition by label key
	var partitionedAppmeshK8SAwsv1Beta2VirtualServices []LabeledAppmeshK8SAwsv1Beta2VirtualServiceSet

	for labelValue, setForValue := range setsByLabel {
		labels := map[string]string{labelKey: labelValue}

		partitionedSet, err := NewLabeledAppmeshK8SAwsv1Beta2VirtualServiceSet(setForValue, labels)
		if err != nil {
			return nil, err
		}

		partitionedAppmeshK8SAwsv1Beta2VirtualServices = append(partitionedAppmeshK8SAwsv1Beta2VirtualServices, partitionedSet)
	}

	// sort for idempotency
	sort.SliceStable(partitionedAppmeshK8SAwsv1Beta2VirtualServices, func(i, j int) bool {
		leftLabelValue := partitionedAppmeshK8SAwsv1Beta2VirtualServices[i].Labels()[labelKey]
		rightLabelValue := partitionedAppmeshK8SAwsv1Beta2VirtualServices[j].Labels()[labelKey]
		return leftLabelValue < rightLabelValue
	})

	return partitionedAppmeshK8SAwsv1Beta2VirtualServices, nil
}

func partitionAppmeshK8SAwsv1Beta2VirtualNodesByLabel(labelKey string, set appmesh_k8s_aws_v1beta2_sets.VirtualNodeSet) ([]LabeledAppmeshK8SAwsv1Beta2VirtualNodeSet, error) {
	setsByLabel := map[string]appmesh_k8s_aws_v1beta2_sets.VirtualNodeSet{}

	for _, obj := range set.List() {
		if obj.Labels == nil {
			return nil, MissingRequiredLabelError(labelKey, "AppmeshK8SAwsv1Beta2VirtualNode", obj)
		}
		labelValue := obj.Labels[labelKey]
		if labelValue == "" {
			return nil, MissingRequiredLabelError(labelKey, "AppmeshK8SAwsv1Beta2VirtualNode", obj)
		}

		setForValue, ok := setsByLabel[labelValue]
		if !ok {
			setForValue = appmesh_k8s_aws_v1beta2_sets.NewVirtualNodeSet()
			setsByLabel[labelValue] = setForValue
		}
		setForValue.Insert(obj)
	}

	// partition by label key
	var partitionedAppmeshK8SAwsv1Beta2VirtualNodes []LabeledAppmeshK8SAwsv1Beta2VirtualNodeSet

	for labelValue, setForValue := range setsByLabel {
		labels := map[string]string{labelKey: labelValue}

		partitionedSet, err := NewLabeledAppmeshK8SAwsv1Beta2VirtualNodeSet(setForValue, labels)
		if err != nil {
			return nil, err
		}

		partitionedAppmeshK8SAwsv1Beta2VirtualNodes = append(partitionedAppmeshK8SAwsv1Beta2VirtualNodes, partitionedSet)
	}

	// sort for idempotency
	sort.SliceStable(partitionedAppmeshK8SAwsv1Beta2VirtualNodes, func(i, j int) bool {
		leftLabelValue := partitionedAppmeshK8SAwsv1Beta2VirtualNodes[i].Labels()[labelKey]
		rightLabelValue := partitionedAppmeshK8SAwsv1Beta2VirtualNodes[j].Labels()[labelKey]
		return leftLabelValue < rightLabelValue
	})

	return partitionedAppmeshK8SAwsv1Beta2VirtualNodes, nil
}

func partitionAppmeshK8SAwsv1Beta2VirtualRoutersByLabel(labelKey string, set appmesh_k8s_aws_v1beta2_sets.VirtualRouterSet) ([]LabeledAppmeshK8SAwsv1Beta2VirtualRouterSet, error) {
	setsByLabel := map[string]appmesh_k8s_aws_v1beta2_sets.VirtualRouterSet{}

	for _, obj := range set.List() {
		if obj.Labels == nil {
			return nil, MissingRequiredLabelError(labelKey, "AppmeshK8SAwsv1Beta2VirtualRouter", obj)
		}
		labelValue := obj.Labels[labelKey]
		if labelValue == "" {
			return nil, MissingRequiredLabelError(labelKey, "AppmeshK8SAwsv1Beta2VirtualRouter", obj)
		}

		setForValue, ok := setsByLabel[labelValue]
		if !ok {
			setForValue = appmesh_k8s_aws_v1beta2_sets.NewVirtualRouterSet()
			setsByLabel[labelValue] = setForValue
		}
		setForValue.Insert(obj)
	}

	// partition by label key
	var partitionedAppmeshK8SAwsv1Beta2VirtualRouters []LabeledAppmeshK8SAwsv1Beta2VirtualRouterSet

	for labelValue, setForValue := range setsByLabel {
		labels := map[string]string{labelKey: labelValue}

		partitionedSet, err := NewLabeledAppmeshK8SAwsv1Beta2VirtualRouterSet(setForValue, labels)
		if err != nil {
			return nil, err
		}

		partitionedAppmeshK8SAwsv1Beta2VirtualRouters = append(partitionedAppmeshK8SAwsv1Beta2VirtualRouters, partitionedSet)
	}

	// sort for idempotency
	sort.SliceStable(partitionedAppmeshK8SAwsv1Beta2VirtualRouters, func(i, j int) bool {
		leftLabelValue := partitionedAppmeshK8SAwsv1Beta2VirtualRouters[i].Labels()[labelKey]
		rightLabelValue := partitionedAppmeshK8SAwsv1Beta2VirtualRouters[j].Labels()[labelKey]
		return leftLabelValue < rightLabelValue
	})

	return partitionedAppmeshK8SAwsv1Beta2VirtualRouters, nil
}

func (s snapshot) AppmeshK8SAwsv1Beta2VirtualServices() []LabeledAppmeshK8SAwsv1Beta2VirtualServiceSet {
	return s.appmeshK8SAwsv1Beta2VirtualServices
}

func (s snapshot) AppmeshK8SAwsv1Beta2VirtualNodes() []LabeledAppmeshK8SAwsv1Beta2VirtualNodeSet {
	return s.appmeshK8SAwsv1Beta2VirtualNodes
}

func (s snapshot) AppmeshK8SAwsv1Beta2VirtualRouters() []LabeledAppmeshK8SAwsv1Beta2VirtualRouterSet {
	return s.appmeshK8SAwsv1Beta2VirtualRouters
}

func (s snapshot) MarshalJSON() ([]byte, error) {
	snapshotMap := map[string]interface{}{"name": s.name}

	appmeshK8SAwsv1Beta2VirtualServiceSet := appmesh_k8s_aws_v1beta2_sets.NewVirtualServiceSet()
	for _, set := range s.appmeshK8SAwsv1Beta2VirtualServices {
		appmeshK8SAwsv1Beta2VirtualServiceSet = appmeshK8SAwsv1Beta2VirtualServiceSet.Union(set.Set())
	}
	snapshotMap["appmeshK8SAwsv1Beta2VirtualServices"] = appmeshK8SAwsv1Beta2VirtualServiceSet.List()
	appmeshK8SAwsv1Beta2VirtualNodeSet := appmesh_k8s_aws_v1beta2_sets.NewVirtualNodeSet()
	for _, set := range s.appmeshK8SAwsv1Beta2VirtualNodes {
		appmeshK8SAwsv1Beta2VirtualNodeSet = appmeshK8SAwsv1Beta2VirtualNodeSet.Union(set.Set())
	}
	snapshotMap["appmeshK8SAwsv1Beta2VirtualNodes"] = appmeshK8SAwsv1Beta2VirtualNodeSet.List()
	appmeshK8SAwsv1Beta2VirtualRouterSet := appmesh_k8s_aws_v1beta2_sets.NewVirtualRouterSet()
	for _, set := range s.appmeshK8SAwsv1Beta2VirtualRouters {
		appmeshK8SAwsv1Beta2VirtualRouterSet = appmeshK8SAwsv1Beta2VirtualRouterSet.Union(set.Set())
	}
	snapshotMap["appmeshK8SAwsv1Beta2VirtualRouters"] = appmeshK8SAwsv1Beta2VirtualRouterSet.List()

	snapshotMap["clusters"] = s.clusters

	return json.Marshal(snapshotMap)
}

// LabeledAppmeshK8SAwsv1Beta2VirtualServiceSet represents a set of appmeshK8SAwsv1Beta2VirtualServices
// which share a common set of labels.
// These labels are used to find diffs between AppmeshK8SAwsv1Beta2VirtualServiceSets.
type LabeledAppmeshK8SAwsv1Beta2VirtualServiceSet interface {
	// returns the set of Labels shared by this AppmeshK8SAwsv1Beta2VirtualServiceSet
	Labels() map[string]string

	// returns the set of VirtualServicees with the given labels
	Set() appmesh_k8s_aws_v1beta2_sets.VirtualServiceSet

	// converts the set to a generic format which can be applied by the Snapshot.Apply functions
	Generic() output.ResourceList
}

type labeledAppmeshK8SAwsv1Beta2VirtualServiceSet struct {
	set    appmesh_k8s_aws_v1beta2_sets.VirtualServiceSet
	labels map[string]string
}

func NewLabeledAppmeshK8SAwsv1Beta2VirtualServiceSet(set appmesh_k8s_aws_v1beta2_sets.VirtualServiceSet, labels map[string]string) (LabeledAppmeshK8SAwsv1Beta2VirtualServiceSet, error) {
	// validate that each VirtualService contains the labels, else this is not a valid LabeledAppmeshK8SAwsv1Beta2VirtualServiceSet
	for _, item := range set.List() {
		for k, v := range labels {
			// k=v must be present in the item
			if item.Labels[k] != v {
				return nil, eris.Errorf("internal error: %v=%v missing on AppmeshK8SAwsv1Beta2VirtualService %v", k, v, item.Name)
			}
		}
	}

	return &labeledAppmeshK8SAwsv1Beta2VirtualServiceSet{set: set, labels: labels}, nil
}

func (l *labeledAppmeshK8SAwsv1Beta2VirtualServiceSet) Labels() map[string]string {
	return l.labels
}

func (l *labeledAppmeshK8SAwsv1Beta2VirtualServiceSet) Set() appmesh_k8s_aws_v1beta2_sets.VirtualServiceSet {
	return l.set
}

func (l labeledAppmeshK8SAwsv1Beta2VirtualServiceSet) Generic() output.ResourceList {
	var desiredResources []ezkube.Object
	for _, desired := range l.set.List() {
		desiredResources = append(desiredResources, desired)
	}

	// enable list func for garbage collection
	listFunc := func(ctx context.Context, cli client.Client) ([]ezkube.Object, error) {
		var list appmesh_k8s_aws_v1beta2.VirtualServiceList
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
		ResourceKind: "VirtualService",
	}
}

// LabeledAppmeshK8SAwsv1Beta2VirtualNodeSet represents a set of appmeshK8SAwsv1Beta2VirtualNodes
// which share a common set of labels.
// These labels are used to find diffs between AppmeshK8SAwsv1Beta2VirtualNodeSets.
type LabeledAppmeshK8SAwsv1Beta2VirtualNodeSet interface {
	// returns the set of Labels shared by this AppmeshK8SAwsv1Beta2VirtualNodeSet
	Labels() map[string]string

	// returns the set of VirtualNodees with the given labels
	Set() appmesh_k8s_aws_v1beta2_sets.VirtualNodeSet

	// converts the set to a generic format which can be applied by the Snapshot.Apply functions
	Generic() output.ResourceList
}

type labeledAppmeshK8SAwsv1Beta2VirtualNodeSet struct {
	set    appmesh_k8s_aws_v1beta2_sets.VirtualNodeSet
	labels map[string]string
}

func NewLabeledAppmeshK8SAwsv1Beta2VirtualNodeSet(set appmesh_k8s_aws_v1beta2_sets.VirtualNodeSet, labels map[string]string) (LabeledAppmeshK8SAwsv1Beta2VirtualNodeSet, error) {
	// validate that each VirtualNode contains the labels, else this is not a valid LabeledAppmeshK8SAwsv1Beta2VirtualNodeSet
	for _, item := range set.List() {
		for k, v := range labels {
			// k=v must be present in the item
			if item.Labels[k] != v {
				return nil, eris.Errorf("internal error: %v=%v missing on AppmeshK8SAwsv1Beta2VirtualNode %v", k, v, item.Name)
			}
		}
	}

	return &labeledAppmeshK8SAwsv1Beta2VirtualNodeSet{set: set, labels: labels}, nil
}

func (l *labeledAppmeshK8SAwsv1Beta2VirtualNodeSet) Labels() map[string]string {
	return l.labels
}

func (l *labeledAppmeshK8SAwsv1Beta2VirtualNodeSet) Set() appmesh_k8s_aws_v1beta2_sets.VirtualNodeSet {
	return l.set
}

func (l labeledAppmeshK8SAwsv1Beta2VirtualNodeSet) Generic() output.ResourceList {
	var desiredResources []ezkube.Object
	for _, desired := range l.set.List() {
		desiredResources = append(desiredResources, desired)
	}

	// enable list func for garbage collection
	listFunc := func(ctx context.Context, cli client.Client) ([]ezkube.Object, error) {
		var list appmesh_k8s_aws_v1beta2.VirtualNodeList
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
		ResourceKind: "VirtualNode",
	}
}

// LabeledAppmeshK8SAwsv1Beta2VirtualRouterSet represents a set of appmeshK8SAwsv1Beta2VirtualRouters
// which share a common set of labels.
// These labels are used to find diffs between AppmeshK8SAwsv1Beta2VirtualRouterSets.
type LabeledAppmeshK8SAwsv1Beta2VirtualRouterSet interface {
	// returns the set of Labels shared by this AppmeshK8SAwsv1Beta2VirtualRouterSet
	Labels() map[string]string

	// returns the set of VirtualRouteres with the given labels
	Set() appmesh_k8s_aws_v1beta2_sets.VirtualRouterSet

	// converts the set to a generic format which can be applied by the Snapshot.Apply functions
	Generic() output.ResourceList
}

type labeledAppmeshK8SAwsv1Beta2VirtualRouterSet struct {
	set    appmesh_k8s_aws_v1beta2_sets.VirtualRouterSet
	labels map[string]string
}

func NewLabeledAppmeshK8SAwsv1Beta2VirtualRouterSet(set appmesh_k8s_aws_v1beta2_sets.VirtualRouterSet, labels map[string]string) (LabeledAppmeshK8SAwsv1Beta2VirtualRouterSet, error) {
	// validate that each VirtualRouter contains the labels, else this is not a valid LabeledAppmeshK8SAwsv1Beta2VirtualRouterSet
	for _, item := range set.List() {
		for k, v := range labels {
			// k=v must be present in the item
			if item.Labels[k] != v {
				return nil, eris.Errorf("internal error: %v=%v missing on AppmeshK8SAwsv1Beta2VirtualRouter %v", k, v, item.Name)
			}
		}
	}

	return &labeledAppmeshK8SAwsv1Beta2VirtualRouterSet{set: set, labels: labels}, nil
}

func (l *labeledAppmeshK8SAwsv1Beta2VirtualRouterSet) Labels() map[string]string {
	return l.labels
}

func (l *labeledAppmeshK8SAwsv1Beta2VirtualRouterSet) Set() appmesh_k8s_aws_v1beta2_sets.VirtualRouterSet {
	return l.set
}

func (l labeledAppmeshK8SAwsv1Beta2VirtualRouterSet) Generic() output.ResourceList {
	var desiredResources []ezkube.Object
	for _, desired := range l.set.List() {
		desiredResources = append(desiredResources, desired)
	}

	// enable list func for garbage collection
	listFunc := func(ctx context.Context, cli client.Client) ([]ezkube.Object, error) {
		var list appmesh_k8s_aws_v1beta2.VirtualRouterList
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
		ResourceKind: "VirtualRouter",
	}
}

type builder struct {
	ctx      context.Context
	name     string
	clusters []string

	appmeshK8SAwsv1Beta2VirtualServices appmesh_k8s_aws_v1beta2_sets.VirtualServiceSet
	appmeshK8SAwsv1Beta2VirtualNodes    appmesh_k8s_aws_v1beta2_sets.VirtualNodeSet
	appmeshK8SAwsv1Beta2VirtualRouters  appmesh_k8s_aws_v1beta2_sets.VirtualRouterSet
}

func NewBuilder(ctx context.Context, name string) *builder {
	return &builder{
		ctx:  ctx,
		name: name,

		appmeshK8SAwsv1Beta2VirtualServices: appmesh_k8s_aws_v1beta2_sets.NewVirtualServiceSet(),
		appmeshK8SAwsv1Beta2VirtualNodes:    appmesh_k8s_aws_v1beta2_sets.NewVirtualNodeSet(),
		appmeshK8SAwsv1Beta2VirtualRouters:  appmesh_k8s_aws_v1beta2_sets.NewVirtualRouterSet(),
	}
}

// the output Builder uses a builder pattern to allow
// iteratively collecting outputs before producing a final snapshot
type Builder interface {

	// add AppmeshK8SAwsv1Beta2VirtualServices to the collected outputs
	AddAppmeshK8SAwsv1Beta2VirtualServices(appmeshK8SAwsv1Beta2VirtualServices ...*appmesh_k8s_aws_v1beta2.VirtualService)

	// get the collected AppmeshK8SAwsv1Beta2VirtualServices
	GetAppmeshK8SAwsv1Beta2VirtualServices() appmesh_k8s_aws_v1beta2_sets.VirtualServiceSet

	// add AppmeshK8SAwsv1Beta2VirtualNodes to the collected outputs
	AddAppmeshK8SAwsv1Beta2VirtualNodes(appmeshK8SAwsv1Beta2VirtualNodes ...*appmesh_k8s_aws_v1beta2.VirtualNode)

	// get the collected AppmeshK8SAwsv1Beta2VirtualNodes
	GetAppmeshK8SAwsv1Beta2VirtualNodes() appmesh_k8s_aws_v1beta2_sets.VirtualNodeSet

	// add AppmeshK8SAwsv1Beta2VirtualRouters to the collected outputs
	AddAppmeshK8SAwsv1Beta2VirtualRouters(appmeshK8SAwsv1Beta2VirtualRouters ...*appmesh_k8s_aws_v1beta2.VirtualRouter)

	// get the collected AppmeshK8SAwsv1Beta2VirtualRouters
	GetAppmeshK8SAwsv1Beta2VirtualRouters() appmesh_k8s_aws_v1beta2_sets.VirtualRouterSet

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

func (b *builder) AddAppmeshK8SAwsv1Beta2VirtualServices(appmeshK8SAwsv1Beta2VirtualServices ...*appmesh_k8s_aws_v1beta2.VirtualService) {
	for _, obj := range appmeshK8SAwsv1Beta2VirtualServices {
		if obj == nil {
			continue
		}
		contextutils.LoggerFrom(b.ctx).Debugf("added output AppmeshK8SAwsv1Beta2VirtualService %v", sets.Key(obj))
		b.appmeshK8SAwsv1Beta2VirtualServices.Insert(obj)
	}
}
func (b *builder) AddAppmeshK8SAwsv1Beta2VirtualNodes(appmeshK8SAwsv1Beta2VirtualNodes ...*appmesh_k8s_aws_v1beta2.VirtualNode) {
	for _, obj := range appmeshK8SAwsv1Beta2VirtualNodes {
		if obj == nil {
			continue
		}
		contextutils.LoggerFrom(b.ctx).Debugf("added output AppmeshK8SAwsv1Beta2VirtualNode %v", sets.Key(obj))
		b.appmeshK8SAwsv1Beta2VirtualNodes.Insert(obj)
	}
}
func (b *builder) AddAppmeshK8SAwsv1Beta2VirtualRouters(appmeshK8SAwsv1Beta2VirtualRouters ...*appmesh_k8s_aws_v1beta2.VirtualRouter) {
	for _, obj := range appmeshK8SAwsv1Beta2VirtualRouters {
		if obj == nil {
			continue
		}
		contextutils.LoggerFrom(b.ctx).Debugf("added output AppmeshK8SAwsv1Beta2VirtualRouter %v", sets.Key(obj))
		b.appmeshK8SAwsv1Beta2VirtualRouters.Insert(obj)
	}
}

func (b *builder) GetAppmeshK8SAwsv1Beta2VirtualServices() appmesh_k8s_aws_v1beta2_sets.VirtualServiceSet {
	return b.appmeshK8SAwsv1Beta2VirtualServices
}
func (b *builder) GetAppmeshK8SAwsv1Beta2VirtualNodes() appmesh_k8s_aws_v1beta2_sets.VirtualNodeSet {
	return b.appmeshK8SAwsv1Beta2VirtualNodes
}
func (b *builder) GetAppmeshK8SAwsv1Beta2VirtualRouters() appmesh_k8s_aws_v1beta2_sets.VirtualRouterSet {
	return b.appmeshK8SAwsv1Beta2VirtualRouters
}

func (b *builder) BuildLabelPartitionedSnapshot(labelKey string) (Snapshot, error) {
	return NewLabelPartitionedSnapshot(
		b.name,
		labelKey,

		b.appmeshK8SAwsv1Beta2VirtualServices,
		b.appmeshK8SAwsv1Beta2VirtualNodes,
		b.appmeshK8SAwsv1Beta2VirtualRouters,
		b.clusters...,
	)
}

func (b *builder) BuildSinglePartitionedSnapshot(snapshotLabels map[string]string) (Snapshot, error) {
	return NewSinglePartitionedSnapshot(
		b.name,
		snapshotLabels,

		b.appmeshK8SAwsv1Beta2VirtualServices,
		b.appmeshK8SAwsv1Beta2VirtualNodes,
		b.appmeshK8SAwsv1Beta2VirtualRouters,
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

	b.AddAppmeshK8SAwsv1Beta2VirtualServices(other.GetAppmeshK8SAwsv1Beta2VirtualServices().List()...)
	b.AddAppmeshK8SAwsv1Beta2VirtualNodes(other.GetAppmeshK8SAwsv1Beta2VirtualNodes().List()...)
	b.AddAppmeshK8SAwsv1Beta2VirtualRouters(other.GetAppmeshK8SAwsv1Beta2VirtualRouters().List()...)
	for _, cluster := range other.Clusters() {
		b.AddCluster(cluster)
	}
}

func (b *builder) Clone() Builder {
	if b == nil {
		return nil
	}
	clone := NewBuilder(b.ctx, b.name)

	for _, appmeshK8SAwsv1Beta2VirtualService := range b.GetAppmeshK8SAwsv1Beta2VirtualServices().List() {
		clone.AddAppmeshK8SAwsv1Beta2VirtualServices(appmeshK8SAwsv1Beta2VirtualService.DeepCopy())
	}
	for _, appmeshK8SAwsv1Beta2VirtualNode := range b.GetAppmeshK8SAwsv1Beta2VirtualNodes().List() {
		clone.AddAppmeshK8SAwsv1Beta2VirtualNodes(appmeshK8SAwsv1Beta2VirtualNode.DeepCopy())
	}
	for _, appmeshK8SAwsv1Beta2VirtualRouter := range b.GetAppmeshK8SAwsv1Beta2VirtualRouters().List() {
		clone.AddAppmeshK8SAwsv1Beta2VirtualRouters(appmeshK8SAwsv1Beta2VirtualRouter.DeepCopy())
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

	// calcualte delta between VirtualServices
	virtualServiceDelta := b.GetVirtualServices().Delta(other.GetVirtualServices())
	virtualServiceGvk := schema.GroupVersionKind{
		Group:   "appmesh.k8s.aws",
		Version: "v1beta2",
		Kind:    "VirtualService",
	}
	delta.AddInserted(virtualServiceGvk, virtualServiceDelta.Inserted)
	delta.AddRemoved(virtualServiceGvk, virtualServiceDelta.Removed)
	// calcualte delta between VirtualNodes
	virtualNodeDelta := b.GetVirtualNodes().Delta(other.GetVirtualNodes())
	virtualNodeGvk := schema.GroupVersionKind{
		Group:   "appmesh.k8s.aws",
		Version: "v1beta2",
		Kind:    "VirtualNode",
	}
	delta.AddInserted(virtualNodeGvk, virtualNodeDelta.Inserted)
	delta.AddRemoved(virtualNodeGvk, virtualNodeDelta.Removed)
	// calcualte delta between VirtualRouters
	virtualRouterDelta := b.GetVirtualRouters().Delta(other.GetVirtualRouters())
	virtualRouterGvk := schema.GroupVersionKind{
		Group:   "appmesh.k8s.aws",
		Version: "v1beta2",
		Kind:    "VirtualRouter",
	}
	delta.AddInserted(virtualRouterGvk, virtualRouterDelta.Inserted)
	delta.AddRemoved(virtualRouterGvk, virtualRouterDelta.Removed)
	return delta
}

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type GroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GroupParameters struct {

	// **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Required
	GroupID *string `json:"groupId" tf:"group_id,omitempty"`

	// Defaults to `on`.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// +kubebuilder:validation:Optional
	PackageID *string `json:"packageId,omitempty" tf:"package_id,omitempty"`

	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**
	// +crossplane:generate:reference:type=github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1.Zone
	// +crossplane:generate:reference:refFieldName=ZoneIDRefs
	// +crossplane:generate:reference:selectorFieldName=ZoneIDSelector
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a Zone in zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDRefs *v1.Reference `json:"zoneIdRefs,omitempty" tf:"-"`

	// Selector for a Zone in zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

// GroupSpec defines the desired state of Group
type GroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupParameters `json:"forProvider"`
}

// GroupStatus defines the observed state of Group.
type GroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Group is the Schema for the Groups API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Group struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupSpec   `json:"spec"`
	Status            GroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupList contains a list of Groups
type GroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Group `json:"items"`
}

// Repository type metadata.
var (
	Group_Kind             = "Group"
	Group_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Group_Kind}.String()
	Group_KindAPIVersion   = Group_Kind + "." + CRDGroupVersion.String()
	Group_GroupVersionKind = CRDGroupVersion.WithKind(Group_Kind)
)

func init() {
	SchemeBuilder.Register(&Group{}, &GroupList{})
}
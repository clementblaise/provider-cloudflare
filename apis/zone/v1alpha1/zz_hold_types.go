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

type HoldInitParameters struct {

	// Enablement status of the zone hold.
	Hold *bool `json:"hold,omitempty" tf:"hold,omitempty"`

	// The RFC3339 compatible timestamp when to automatically re-enable the zone hold.
	HoldAfter *string `json:"holdAfter,omitempty" tf:"hold_after,omitempty"`

	// Whether to extend to block any subdomain of the given zone.
	IncludeSubdomains *bool `json:"includeSubdomains,omitempty" tf:"include_subdomains,omitempty"`
}

type HoldObservation struct {

	// Enablement status of the zone hold.
	Hold *bool `json:"hold,omitempty" tf:"hold,omitempty"`

	// The RFC3339 compatible timestamp when to automatically re-enable the zone hold.
	HoldAfter *string `json:"holdAfter,omitempty" tf:"hold_after,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether to extend to block any subdomain of the given zone.
	IncludeSubdomains *bool `json:"includeSubdomains,omitempty" tf:"include_subdomains,omitempty"`

	// The zone identifier to target for the resource.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type HoldParameters struct {

	// Enablement status of the zone hold.
	// +kubebuilder:validation:Optional
	Hold *bool `json:"hold,omitempty" tf:"hold,omitempty"`

	// The RFC3339 compatible timestamp when to automatically re-enable the zone hold.
	// +kubebuilder:validation:Optional
	HoldAfter *string `json:"holdAfter,omitempty" tf:"hold_after,omitempty"`

	// Whether to extend to block any subdomain of the given zone.
	// +kubebuilder:validation:Optional
	IncludeSubdomains *bool `json:"includeSubdomains,omitempty" tf:"include_subdomains,omitempty"`

	// The zone identifier to target for the resource.
	// +crossplane:generate:reference:type=Zone
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a Zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// Selector for a Zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

// HoldSpec defines the desired state of Hold
type HoldSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HoldParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider HoldInitParameters `json:"initProvider,omitempty"`
}

// HoldStatus defines the observed state of Hold.
type HoldStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HoldObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Hold is the Schema for the Holds API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Hold struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hold) || has(self.initProvider.hold)",message="hold is a required parameter"
	Spec   HoldSpec   `json:"spec"`
	Status HoldStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HoldList contains a list of Holds
type HoldList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Hold `json:"items"`
}

// Repository type metadata.
var (
	Hold_Kind             = "Hold"
	Hold_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Hold_Kind}.String()
	Hold_KindAPIVersion   = Hold_Kind + "." + CRDGroupVersion.String()
	Hold_GroupVersionKind = CRDGroupVersion.WithKind(Hold_Kind)
)

func init() {
	SchemeBuilder.Register(&Hold{}, &HoldList{})
}

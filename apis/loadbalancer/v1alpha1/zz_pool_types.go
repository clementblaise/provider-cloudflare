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

type LoadSheddingObservation struct {
}

type LoadSheddingParameters struct {

	// Percent of traffic to shed 0 - 100. Defaults to `0`.
	// +kubebuilder:validation:Optional
	DefaultPercent *float64 `json:"defaultPercent,omitempty" tf:"default_percent,omitempty"`

	// Method of shedding traffic. Available values: “, `hash`, `random`. Defaults to `""`.
	// +kubebuilder:validation:Optional
	DefaultPolicy *string `json:"defaultPolicy,omitempty" tf:"default_policy,omitempty"`

	// Percent of session traffic to shed 0 - 100. Defaults to `0`.
	// +kubebuilder:validation:Optional
	SessionPercent *float64 `json:"sessionPercent,omitempty" tf:"session_percent,omitempty"`

	// Method of shedding traffic. Available values: “, `hash`. Defaults to `""`.
	// +kubebuilder:validation:Optional
	SessionPolicy *string `json:"sessionPolicy,omitempty" tf:"session_policy,omitempty"`
}

type OriginSteeringObservation struct {
}

type OriginSteeringParameters struct {

	// Origin steering policy to be used. Available values: “, `hash`, `random`. Defaults to `random`.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type OriginsHeaderObservation struct {
}

type OriginsHeaderParameters struct {

	// HTTP Header name.
	// +kubebuilder:validation:Required
	Header *string `json:"header" tf:"header,omitempty"`

	// Values for the HTTP headers.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type OriginsObservation struct {
}

type OriginsParameters struct {

	// The IP address (IPv4 or IPv6) of the origin, or the publicly addressable hostname.
	// +kubebuilder:validation:Required
	Address *string `json:"address" tf:"address,omitempty"`

	// Whether this origin is enabled. Disabled origins will not receive traffic and are excluded from health checks. Defaults to `true`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// HTTP request headers.
	// +kubebuilder:validation:Optional
	Header []OriginsHeaderParameters `json:"header,omitempty" tf:"header,omitempty"`

	// A human-identifiable name for the origin.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The weight (0.01 - 1.00) of this origin, relative to other origins in the pool. Equal values mean equal weighting. A weight of 0 means traffic will not be sent to this origin, but health is still checked. Defaults to `1`.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type PoolObservation struct {
	CreatedOn *string `json:"createdOn,omitempty" tf:"created_on,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ModifiedOn *string `json:"modifiedOn,omitempty" tf:"modified_on,omitempty"`
}

type PoolParameters struct {

	// The account identifier to target for the resource.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// +kubebuilder:validation:Optional
	CheckRegions []*string `json:"checkRegions,omitempty" tf:"check_regions,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Defaults to `true`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Latitude *float64 `json:"latitude,omitempty" tf:"latitude,omitempty"`

	// +kubebuilder:validation:Optional
	LoadShedding []LoadSheddingParameters `json:"loadShedding,omitempty" tf:"load_shedding,omitempty"`

	// +kubebuilder:validation:Optional
	Longitude *float64 `json:"longitude,omitempty" tf:"longitude,omitempty"`

	// Defaults to `1`.
	// +kubebuilder:validation:Optional
	MinimumOrigins *float64 `json:"minimumOrigins,omitempty" tf:"minimum_origins,omitempty"`

	// +crossplane:generate:reference:type=Monitor
	// +kubebuilder:validation:Optional
	Monitor *string `json:"monitor,omitempty" tf:"monitor,omitempty"`

	// Reference to a Monitor to populate monitor.
	// +kubebuilder:validation:Optional
	MonitorRef *v1.Reference `json:"monitorRef,omitempty" tf:"-"`

	// Selector for a Monitor to populate monitor.
	// +kubebuilder:validation:Optional
	MonitorSelector *v1.Selector `json:"monitorSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NotificationEmail *string `json:"notificationEmail,omitempty" tf:"notification_email,omitempty"`

	// +kubebuilder:validation:Optional
	OriginSteering []OriginSteeringParameters `json:"originSteering,omitempty" tf:"origin_steering,omitempty"`

	// +kubebuilder:validation:Required
	Origins []OriginsParameters `json:"origins" tf:"origins,omitempty"`
}

// PoolSpec defines the desired state of Pool
type PoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PoolParameters `json:"forProvider"`
}

// PoolStatus defines the observed state of Pool.
type PoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Pool is the Schema for the Pools API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Pool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PoolSpec   `json:"spec"`
	Status            PoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PoolList contains a list of Pools
type PoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pool `json:"items"`
}

// Repository type metadata.
var (
	Pool_Kind             = "Pool"
	Pool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Pool_Kind}.String()
	Pool_KindAPIVersion   = Pool_Kind + "." + CRDGroupVersion.String()
	Pool_GroupVersionKind = CRDGroupVersion.WithKind(Pool_Kind)
)

func init() {
	SchemeBuilder.Register(&Pool{}, &PoolList{})
}

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

type SplitTunnelObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SplitTunnelParameters struct {

	// The account identifier to target for the resource.
	// +kubebuilder:validation:Required
	AccountID *string `json:"accountId" tf:"account_id,omitempty"`

	// The mode of the split tunnel policy. Available values: `include`, `exclude`.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The settings policy for which to configure this split tunnel policy.
	// +crossplane:generate:reference:type=DeviceSettingsPolicy
	// +kubebuilder:validation:Optional
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Reference to a DeviceSettingsPolicy to populate policyId.
	// +kubebuilder:validation:Optional
	PolicyIDRef *v1.Reference `json:"policyIdRef,omitempty" tf:"-"`

	// Selector for a DeviceSettingsPolicy to populate policyId.
	// +kubebuilder:validation:Optional
	PolicyIDSelector *v1.Selector `json:"policyIdSelector,omitempty" tf:"-"`

	// The value of the tunnel attributes.
	// +kubebuilder:validation:Required
	Tunnels []TunnelsParameters `json:"tunnels" tf:"tunnels,omitempty"`
}

type TunnelsObservation struct {
}

type TunnelsParameters struct {

	// The address for the tunnel.
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// A description for the tunnel.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The domain name for the tunnel.
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`
}

// SplitTunnelSpec defines the desired state of SplitTunnel
type SplitTunnelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SplitTunnelParameters `json:"forProvider"`
}

// SplitTunnelStatus defines the observed state of SplitTunnel.
type SplitTunnelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SplitTunnelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SplitTunnel is the Schema for the SplitTunnels API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type SplitTunnel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SplitTunnelSpec   `json:"spec"`
	Status            SplitTunnelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SplitTunnelList contains a list of SplitTunnels
type SplitTunnelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SplitTunnel `json:"items"`
}

// Repository type metadata.
var (
	SplitTunnel_Kind             = "SplitTunnel"
	SplitTunnel_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SplitTunnel_Kind}.String()
	SplitTunnel_KindAPIVersion   = SplitTunnel_Kind + "." + CRDGroupVersion.String()
	SplitTunnel_GroupVersionKind = CRDGroupVersion.WithKind(SplitTunnel_Kind)
)

func init() {
	SchemeBuilder.Register(&SplitTunnel{}, &SplitTunnelList{})
}
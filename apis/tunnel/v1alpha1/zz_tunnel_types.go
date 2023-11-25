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

type TunnelInitParameters struct {

	// Indicates if this is a locally or remotely configured tunnel. If `local`, manage the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the tunnel on the Zero Trust dashboard or using tunnel_config, tunnel_route or tunnel_virtual_network resources. Available values: `local`, `cloudflare`. **Modifying this attribute will force creation of a new resource.**
	ConfigSrc *string `json:"configSrc,omitempty" tf:"config_src,omitempty"`

	// A user-friendly name chosen when the tunnel is created. **Modifying this attribute will force creation of a new resource.**
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type TunnelObservation struct {

	// The account identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Usable CNAME for accessing the Tunnel.
	Cname *string `json:"cname,omitempty" tf:"cname,omitempty"`

	// Indicates if this is a locally or remotely configured tunnel. If `local`, manage the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the tunnel on the Zero Trust dashboard or using tunnel_config, tunnel_route or tunnel_virtual_network resources. Available values: `local`, `cloudflare`. **Modifying this attribute will force creation of a new resource.**
	ConfigSrc *string `json:"configSrc,omitempty" tf:"config_src,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A user-friendly name chosen when the tunnel is created. **Modifying this attribute will force creation of a new resource.**
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type TunnelParameters struct {

	// The account identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**
	// +crossplane:generate:reference:type=github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Reference to a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDRef *v1.Reference `json:"accountIdRef,omitempty" tf:"-"`

	// Selector for a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDSelector *v1.Selector `json:"accountIdSelector,omitempty" tf:"-"`

	// Indicates if this is a locally or remotely configured tunnel. If `local`, manage the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the tunnel on the Zero Trust dashboard or using tunnel_config, tunnel_route or tunnel_virtual_network resources. Available values: `local`, `cloudflare`. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	ConfigSrc *string `json:"configSrc,omitempty" tf:"config_src,omitempty"`

	// A user-friendly name chosen when the tunnel is created. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// 32 or more bytes, encoded as a base64 string. The Create Argo Tunnel endpoint sets this as the tunnel's password. Anyone wishing to run the tunnel needs this password. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	SecretSecretRef v1.SecretKeySelector `json:"secretSecretRef" tf:"-"`
}

// TunnelSpec defines the desired state of Tunnel
type TunnelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TunnelParameters `json:"forProvider"`
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
	InitProvider TunnelInitParameters `json:"initProvider,omitempty"`
}

// TunnelStatus defines the observed state of Tunnel.
type TunnelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TunnelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Tunnel is the Schema for the Tunnels API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Tunnel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.secretSecretRef)",message="secretSecretRef is a required parameter"
	Spec   TunnelSpec   `json:"spec"`
	Status TunnelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TunnelList contains a list of Tunnels
type TunnelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tunnel `json:"items"`
}

// Repository type metadata.
var (
	Tunnel_Kind             = "Tunnel"
	Tunnel_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Tunnel_Kind}.String()
	Tunnel_KindAPIVersion   = Tunnel_Kind + "." + CRDGroupVersion.String()
	Tunnel_GroupVersionKind = CRDGroupVersion.WithKind(Tunnel_Kind)
)

func init() {
	SchemeBuilder.Register(&Tunnel{}, &TunnelList{})
}

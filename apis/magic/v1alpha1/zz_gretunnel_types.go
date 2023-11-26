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

type GRETunnelInitParameters struct {

	// The IP address assigned to the Cloudflare side of the GRE tunnel.
	// The IP address assigned to the Cloudflare side of the GRE tunnel.
	CloudflareGreEndpoint *string `json:"cloudflareGreEndpoint,omitempty" tf:"cloudflare_gre_endpoint,omitempty"`

	// The IP address assigned to the customer side of the GRE tunnel.
	// The IP address assigned to the customer side of the GRE tunnel.
	CustomerGreEndpoint *string `json:"customerGreEndpoint,omitempty" tf:"customer_gre_endpoint,omitempty"`

	// An optional description of the GRE tunnel.
	// Description of the GRE tunnel intent.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies if ICMP tunnel health checks are enabled Default: true.
	// Specifies if ICMP tunnel health checks are enabled.
	HealthCheckEnabled *bool `json:"healthCheckEnabled,omitempty" tf:"health_check_enabled,omitempty"`

	// The IP address of the customer endpoint that will receive tunnel health checks. Default: <customer_gre_endpoint>.
	// The IP address of the customer endpoint that will receive tunnel health checks.
	HealthCheckTarget *string `json:"healthCheckTarget,omitempty" tf:"health_check_target,omitempty"`

	// Specifies the ICMP echo type for the health check (request or reply) Default: reply.
	// Specifies the ICMP echo type for the health check. Available values: `request`, `reply`.
	HealthCheckType *string `json:"healthCheckType,omitempty" tf:"health_check_type,omitempty"`

	// 31-bit prefix (/31 in CIDR notation) supporting 2 hosts, one for each side of the tunnel.
	// 31-bit prefix (/31 in CIDR notation) supporting 2 hosts, one for each side of the tunnel.
	InterfaceAddress *string `json:"interfaceAddress,omitempty" tf:"interface_address,omitempty"`

	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. Maximum value 1476 and minimum value 576. Default: 1476.
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel.
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// Name of the GRE tunnel.
	// Name of the GRE tunnel.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Time To Live (TTL) in number of hops of the GRE tunnel. Minimum value 64. Default: 64.
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type GRETunnelObservation struct {

	// The ID of the account where the tunnel is being created.
	// The account identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The IP address assigned to the Cloudflare side of the GRE tunnel.
	// The IP address assigned to the Cloudflare side of the GRE tunnel.
	CloudflareGreEndpoint *string `json:"cloudflareGreEndpoint,omitempty" tf:"cloudflare_gre_endpoint,omitempty"`

	// The IP address assigned to the customer side of the GRE tunnel.
	// The IP address assigned to the customer side of the GRE tunnel.
	CustomerGreEndpoint *string `json:"customerGreEndpoint,omitempty" tf:"customer_gre_endpoint,omitempty"`

	// An optional description of the GRE tunnel.
	// Description of the GRE tunnel intent.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies if ICMP tunnel health checks are enabled Default: true.
	// Specifies if ICMP tunnel health checks are enabled.
	HealthCheckEnabled *bool `json:"healthCheckEnabled,omitempty" tf:"health_check_enabled,omitempty"`

	// The IP address of the customer endpoint that will receive tunnel health checks. Default: <customer_gre_endpoint>.
	// The IP address of the customer endpoint that will receive tunnel health checks.
	HealthCheckTarget *string `json:"healthCheckTarget,omitempty" tf:"health_check_target,omitempty"`

	// Specifies the ICMP echo type for the health check (request or reply) Default: reply.
	// Specifies the ICMP echo type for the health check. Available values: `request`, `reply`.
	HealthCheckType *string `json:"healthCheckType,omitempty" tf:"health_check_type,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// 31-bit prefix (/31 in CIDR notation) supporting 2 hosts, one for each side of the tunnel.
	// 31-bit prefix (/31 in CIDR notation) supporting 2 hosts, one for each side of the tunnel.
	InterfaceAddress *string `json:"interfaceAddress,omitempty" tf:"interface_address,omitempty"`

	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. Maximum value 1476 and minimum value 576. Default: 1476.
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel.
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// Name of the GRE tunnel.
	// Name of the GRE tunnel.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Time To Live (TTL) in number of hops of the GRE tunnel. Minimum value 64. Default: 64.
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type GRETunnelParameters struct {

	// The ID of the account where the tunnel is being created.
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

	// The IP address assigned to the Cloudflare side of the GRE tunnel.
	// The IP address assigned to the Cloudflare side of the GRE tunnel.
	// +kubebuilder:validation:Optional
	CloudflareGreEndpoint *string `json:"cloudflareGreEndpoint,omitempty" tf:"cloudflare_gre_endpoint,omitempty"`

	// The IP address assigned to the customer side of the GRE tunnel.
	// The IP address assigned to the customer side of the GRE tunnel.
	// +kubebuilder:validation:Optional
	CustomerGreEndpoint *string `json:"customerGreEndpoint,omitempty" tf:"customer_gre_endpoint,omitempty"`

	// An optional description of the GRE tunnel.
	// Description of the GRE tunnel intent.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies if ICMP tunnel health checks are enabled Default: true.
	// Specifies if ICMP tunnel health checks are enabled.
	// +kubebuilder:validation:Optional
	HealthCheckEnabled *bool `json:"healthCheckEnabled,omitempty" tf:"health_check_enabled,omitempty"`

	// The IP address of the customer endpoint that will receive tunnel health checks. Default: <customer_gre_endpoint>.
	// The IP address of the customer endpoint that will receive tunnel health checks.
	// +kubebuilder:validation:Optional
	HealthCheckTarget *string `json:"healthCheckTarget,omitempty" tf:"health_check_target,omitempty"`

	// Specifies the ICMP echo type for the health check (request or reply) Default: reply.
	// Specifies the ICMP echo type for the health check. Available values: `request`, `reply`.
	// +kubebuilder:validation:Optional
	HealthCheckType *string `json:"healthCheckType,omitempty" tf:"health_check_type,omitempty"`

	// 31-bit prefix (/31 in CIDR notation) supporting 2 hosts, one for each side of the tunnel.
	// 31-bit prefix (/31 in CIDR notation) supporting 2 hosts, one for each side of the tunnel.
	// +kubebuilder:validation:Optional
	InterfaceAddress *string `json:"interfaceAddress,omitempty" tf:"interface_address,omitempty"`

	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. Maximum value 1476 and minimum value 576. Default: 1476.
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel.
	// +kubebuilder:validation:Optional
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// Name of the GRE tunnel.
	// Name of the GRE tunnel.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Time To Live (TTL) in number of hops of the GRE tunnel. Minimum value 64. Default: 64.
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

// GRETunnelSpec defines the desired state of GRETunnel
type GRETunnelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GRETunnelParameters `json:"forProvider"`
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
	InitProvider GRETunnelInitParameters `json:"initProvider,omitempty"`
}

// GRETunnelStatus defines the observed state of GRETunnel.
type GRETunnelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GRETunnelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GRETunnel is the Schema for the GRETunnels API. Provides a resource which manages GRE tunnels for Magic Transit.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type GRETunnel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cloudflareGreEndpoint) || has(self.initProvider.cloudflareGreEndpoint)",message="cloudflareGreEndpoint is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.customerGreEndpoint) || has(self.initProvider.customerGreEndpoint)",message="customerGreEndpoint is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.interfaceAddress) || has(self.initProvider.interfaceAddress)",message="interfaceAddress is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   GRETunnelSpec   `json:"spec"`
	Status GRETunnelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GRETunnelList contains a list of GRETunnels
type GRETunnelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GRETunnel `json:"items"`
}

// Repository type metadata.
var (
	GRETunnel_Kind             = "GRETunnel"
	GRETunnel_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GRETunnel_Kind}.String()
	GRETunnel_KindAPIVersion   = GRETunnel_Kind + "." + CRDGroupVersion.String()
	GRETunnel_GroupVersionKind = CRDGroupVersion.WithKind(GRETunnel_Kind)
)

func init() {
	SchemeBuilder.Register(&GRETunnel{}, &GRETunnelList{})
}

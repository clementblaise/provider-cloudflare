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

type DomainsInitParameters struct {

	// A list of IP addresses to handle domain resolution.
	DNSServer []*string `json:"dnsServer,omitempty" tf:"dns_server,omitempty"`

	// A description of the fallback domain, displayed in the client UI.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The domain suffix to match when resolving locally.
	Suffix *string `json:"suffix,omitempty" tf:"suffix,omitempty"`
}

type DomainsObservation struct {

	// A list of IP addresses to handle domain resolution.
	DNSServer []*string `json:"dnsServer,omitempty" tf:"dns_server,omitempty"`

	// A description of the fallback domain, displayed in the client UI.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The domain suffix to match when resolving locally.
	Suffix *string `json:"suffix,omitempty" tf:"suffix,omitempty"`
}

type DomainsParameters struct {

	// A list of IP addresses to handle domain resolution.
	// +kubebuilder:validation:Optional
	DNSServer []*string `json:"dnsServer,omitempty" tf:"dns_server,omitempty"`

	// A description of the fallback domain, displayed in the client UI.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The domain suffix to match when resolving locally.
	// +kubebuilder:validation:Optional
	Suffix *string `json:"suffix,omitempty" tf:"suffix,omitempty"`
}

type FallbackDomainInitParameters struct {
	Domains []DomainsInitParameters `json:"domains,omitempty" tf:"domains,omitempty"`
}

type FallbackDomainObservation struct {

	// The account identifier to target for the resource.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	Domains []DomainsObservation `json:"domains,omitempty" tf:"domains,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The settings policy for which to configure this fallback domain policy.
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`
}

type FallbackDomainParameters struct {

	// The account identifier to target for the resource.
	// +crossplane:generate:reference:type=github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1.Account
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Reference to a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDRef *v1.Reference `json:"accountIdRef,omitempty" tf:"-"`

	// Selector for a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDSelector *v1.Selector `json:"accountIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Domains []DomainsParameters `json:"domains,omitempty" tf:"domains,omitempty"`

	// The settings policy for which to configure this fallback domain policy.
	// +crossplane:generate:reference:type=DeviceSettingsPolicy
	// +kubebuilder:validation:Optional
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Reference to a DeviceSettingsPolicy to populate policyId.
	// +kubebuilder:validation:Optional
	PolicyIDRef *v1.Reference `json:"policyIdRef,omitempty" tf:"-"`

	// Selector for a DeviceSettingsPolicy to populate policyId.
	// +kubebuilder:validation:Optional
	PolicyIDSelector *v1.Selector `json:"policyIdSelector,omitempty" tf:"-"`
}

// FallbackDomainSpec defines the desired state of FallbackDomain
type FallbackDomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FallbackDomainParameters `json:"forProvider"`
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
	InitProvider FallbackDomainInitParameters `json:"initProvider,omitempty"`
}

// FallbackDomainStatus defines the observed state of FallbackDomain.
type FallbackDomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FallbackDomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FallbackDomain is the Schema for the FallbackDomains API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type FallbackDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.domains) || has(self.initProvider.domains)",message="domains is a required parameter"
	Spec   FallbackDomainSpec   `json:"spec"`
	Status FallbackDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FallbackDomainList contains a list of FallbackDomains
type FallbackDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FallbackDomain `json:"items"`
}

// Repository type metadata.
var (
	FallbackDomain_Kind             = "FallbackDomain"
	FallbackDomain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FallbackDomain_Kind}.String()
	FallbackDomain_KindAPIVersion   = FallbackDomain_Kind + "." + CRDGroupVersion.String()
	FallbackDomain_GroupVersionKind = CRDGroupVersion.WithKind(FallbackDomain_Kind)
)

func init() {
	SchemeBuilder.Register(&FallbackDomain{}, &FallbackDomainList{})
}

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

type KeylessCertificateInitParameters struct {

	// A ubiquitous bundle has the highest probability of being verified everywhere, even by clients using outdated or unusual trust stores. An optimal bundle uses the shortest chain and newest intermediates. And the force bundle verifies the chain, but does not otherwise modify it. Available values: `ubiquitous`, `optimal`, `force`. Defaults to `ubiquitous`. **Modifying this attribute will force creation of a new resource.**
	BundleMethod *string `json:"bundleMethod,omitempty" tf:"bundle_method,omitempty"`

	// The zone's SSL certificate or SSL certificate and intermediate(s). **Modifying this attribute will force creation of a new resource.**
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// Whether the KeyLess SSL is on.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The KeyLess SSL host.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The KeyLess SSL name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The KeyLess SSL port used to communicate between Cloudflare and the client's KeyLess SSL server. Defaults to `24008`.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}

type KeylessCertificateObservation struct {

	// A ubiquitous bundle has the highest probability of being verified everywhere, even by clients using outdated or unusual trust stores. An optimal bundle uses the shortest chain and newest intermediates. And the force bundle verifies the chain, but does not otherwise modify it. Available values: `ubiquitous`, `optimal`, `force`. Defaults to `ubiquitous`. **Modifying this attribute will force creation of a new resource.**
	BundleMethod *string `json:"bundleMethod,omitempty" tf:"bundle_method,omitempty"`

	// The zone's SSL certificate or SSL certificate and intermediate(s). **Modifying this attribute will force creation of a new resource.**
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// Whether the KeyLess SSL is on.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The KeyLess SSL host.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The KeyLess SSL name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The KeyLess SSL port used to communicate between Cloudflare and the client's KeyLess SSL server. Defaults to `24008`.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Status of the KeyLess SSL.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The zone identifier to target for the resource.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type KeylessCertificateParameters struct {

	// A ubiquitous bundle has the highest probability of being verified everywhere, even by clients using outdated or unusual trust stores. An optimal bundle uses the shortest chain and newest intermediates. And the force bundle verifies the chain, but does not otherwise modify it. Available values: `ubiquitous`, `optimal`, `force`. Defaults to `ubiquitous`. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	BundleMethod *string `json:"bundleMethod,omitempty" tf:"bundle_method,omitempty"`

	// The zone's SSL certificate or SSL certificate and intermediate(s). **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// Whether the KeyLess SSL is on.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The KeyLess SSL host.
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The KeyLess SSL name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The KeyLess SSL port used to communicate between Cloudflare and the client's KeyLess SSL server. Defaults to `24008`.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

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

// KeylessCertificateSpec defines the desired state of KeylessCertificate
type KeylessCertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KeylessCertificateParameters `json:"forProvider"`
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
	InitProvider KeylessCertificateInitParameters `json:"initProvider,omitempty"`
}

// KeylessCertificateStatus defines the observed state of KeylessCertificate.
type KeylessCertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KeylessCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KeylessCertificate is the Schema for the KeylessCertificates API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type KeylessCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.certificate) || has(self.initProvider.certificate)",message="certificate is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.host) || has(self.initProvider.host)",message="host is a required parameter"
	Spec   KeylessCertificateSpec   `json:"spec"`
	Status KeylessCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeylessCertificateList contains a list of KeylessCertificates
type KeylessCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeylessCertificate `json:"items"`
}

// Repository type metadata.
var (
	KeylessCertificate_Kind             = "KeylessCertificate"
	KeylessCertificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KeylessCertificate_Kind}.String()
	KeylessCertificate_KindAPIVersion   = KeylessCertificate_Kind + "." + CRDGroupVersion.String()
	KeylessCertificate_GroupVersionKind = CRDGroupVersion.WithKind(KeylessCertificate_Kind)
)

func init() {
	SchemeBuilder.Register(&KeylessCertificate{}, &KeylessCertificateList{})
}

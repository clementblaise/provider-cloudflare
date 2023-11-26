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

type SiteInitParameters struct {

	// Whether Cloudflare will automatically inject the JavaScript snippet for orange-clouded sites. **Modifying this attribute will force creation of a new resource.**
	AutoInstall *bool `json:"autoInstall,omitempty" tf:"auto_install,omitempty"`

	// The hostname to use for gray-clouded sites. Must provide only one of `zone_tag`. **Modifying this attribute will force creation of a new resource.**
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The zone identifier for orange-clouded sites. Must provide only one of `host`. **Modifying this attribute will force creation of a new resource.**
	ZoneTag *string `json:"zoneTag,omitempty" tf:"zone_tag,omitempty"`
}

type SiteObservation struct {

	// The account identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Whether Cloudflare will automatically inject the JavaScript snippet for orange-clouded sites. **Modifying this attribute will force creation of a new resource.**
	AutoInstall *bool `json:"autoInstall,omitempty" tf:"auto_install,omitempty"`

	// The hostname to use for gray-clouded sites. Must provide only one of `zone_tag`. **Modifying this attribute will force creation of a new resource.**
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID for the ruleset associated to this Web Analytics Site.
	RulesetID *string `json:"rulesetId,omitempty" tf:"ruleset_id,omitempty"`

	// The Web Analytics site tag.
	SiteTag *string `json:"siteTag,omitempty" tf:"site_tag,omitempty"`

	// The zone identifier for orange-clouded sites. Must provide only one of `host`. **Modifying this attribute will force creation of a new resource.**
	ZoneTag *string `json:"zoneTag,omitempty" tf:"zone_tag,omitempty"`
}

type SiteParameters struct {

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

	// Whether Cloudflare will automatically inject the JavaScript snippet for orange-clouded sites. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	AutoInstall *bool `json:"autoInstall,omitempty" tf:"auto_install,omitempty"`

	// The hostname to use for gray-clouded sites. Must provide only one of `zone_tag`. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The zone identifier for orange-clouded sites. Must provide only one of `host`. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	ZoneTag *string `json:"zoneTag,omitempty" tf:"zone_tag,omitempty"`
}

// SiteSpec defines the desired state of Site
type SiteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SiteParameters `json:"forProvider"`
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
	InitProvider SiteInitParameters `json:"initProvider,omitempty"`
}

// SiteStatus defines the observed state of Site.
type SiteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SiteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Site is the Schema for the Sites API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Site struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.autoInstall) || has(self.initProvider.autoInstall)",message="autoInstall is a required parameter"
	Spec   SiteSpec   `json:"spec"`
	Status SiteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiteList contains a list of Sites
type SiteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Site `json:"items"`
}

// Repository type metadata.
var (
	Site_Kind             = "Site"
	Site_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Site_Kind}.String()
	Site_KindAPIVersion   = Site_Kind + "." + CRDGroupVersion.String()
	Site_GroupVersionKind = CRDGroupVersion.WithKind(Site_Kind)
)

func init() {
	SchemeBuilder.Register(&Site{}, &SiteList{})
}

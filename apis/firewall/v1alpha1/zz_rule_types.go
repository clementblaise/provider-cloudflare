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

type RuleInitParameters struct {

	// The action to apply to a matched request. Available values: `block`, `challenge`, `allow`, `js_challenge`, `managed_challenge`, `log`, `bypass`.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// A description of the rule to help identify it.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether this filter based firewall rule is currently paused.
	Paused *bool `json:"paused,omitempty" tf:"paused,omitempty"`

	// The priority of the rule to allow control of processing order. A lower number indicates high priority. If not provided, any rules with a priority will be sequenced before those without.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// List of products to bypass for a request when the bypass action is used. Available values: `zoneLockdown`, `uaBlock`, `bic`, `hot`, `securityLevel`, `rateLimit`, `waf`.
	Products []*string `json:"products,omitempty" tf:"products,omitempty"`
}

type RuleObservation struct {

	// The action to apply to a matched request. Available values: `block`, `challenge`, `allow`, `js_challenge`, `managed_challenge`, `log`, `bypass`.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// A description of the rule to help identify it.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The identifier of the Filter to use for determining if the Firewall Rule should be triggered.
	FilterID *string `json:"filterId,omitempty" tf:"filter_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether this filter based firewall rule is currently paused.
	Paused *bool `json:"paused,omitempty" tf:"paused,omitempty"`

	// The priority of the rule to allow control of processing order. A lower number indicates high priority. If not provided, any rules with a priority will be sequenced before those without.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// List of products to bypass for a request when the bypass action is used. Available values: `zoneLockdown`, `uaBlock`, `bic`, `hot`, `securityLevel`, `rateLimit`, `waf`.
	Products []*string `json:"products,omitempty" tf:"products,omitempty"`

	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type RuleParameters struct {

	// The action to apply to a matched request. Available values: `block`, `challenge`, `allow`, `js_challenge`, `managed_challenge`, `log`, `bypass`.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// A description of the rule to help identify it.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The identifier of the Filter to use for determining if the Firewall Rule should be triggered.
	// +crossplane:generate:reference:type=github.com/clementblaise/provider-cloudflare/apis/filters/v1alpha1.Filter
	// +kubebuilder:validation:Optional
	FilterID *string `json:"filterId,omitempty" tf:"filter_id,omitempty"`

	// Reference to a Filter in filters to populate filterId.
	// +kubebuilder:validation:Optional
	FilterIDRef *v1.Reference `json:"filterIdRef,omitempty" tf:"-"`

	// Selector for a Filter in filters to populate filterId.
	// +kubebuilder:validation:Optional
	FilterIDSelector *v1.Selector `json:"filterIdSelector,omitempty" tf:"-"`

	// Whether this filter based firewall rule is currently paused.
	// +kubebuilder:validation:Optional
	Paused *bool `json:"paused,omitempty" tf:"paused,omitempty"`

	// The priority of the rule to allow control of processing order. A lower number indicates high priority. If not provided, any rules with a priority will be sequenced before those without.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// List of products to bypass for a request when the bypass action is used. Available values: `zoneLockdown`, `uaBlock`, `bic`, `hot`, `securityLevel`, `rateLimit`, `waf`.
	// +kubebuilder:validation:Optional
	Products []*string `json:"products,omitempty" tf:"products,omitempty"`

	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**
	// +crossplane:generate:reference:type=github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1.Zone
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a Zone in zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// Selector for a Zone in zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

// RuleSpec defines the desired state of Rule
type RuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RuleParameters `json:"forProvider"`
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
	InitProvider RuleInitParameters `json:"initProvider,omitempty"`
}

// RuleStatus defines the observed state of Rule.
type RuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Rule is the Schema for the Rules API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Rule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.action) || has(self.initProvider.action)",message="action is a required parameter"
	Spec   RuleSpec   `json:"spec"`
	Status RuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuleList contains a list of Rules
type RuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rule `json:"items"`
}

// Repository type metadata.
var (
	Rule_Kind             = "Rule"
	Rule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Rule_Kind}.String()
	Rule_KindAPIVersion   = Rule_Kind + "." + CRDGroupVersion.String()
	Rule_GroupVersionKind = CRDGroupVersion.WithKind(Rule_Kind)
)

func init() {
	SchemeBuilder.Register(&Rule{}, &RuleList{})
}

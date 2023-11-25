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

type ManagementInitParameters struct {

	// Automatically update to the newest bot detection models created by Cloudflare as they are released. [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes).
	AutoUpdateModel *bool `json:"autoUpdateModel,omitempty" tf:"auto_update_model,omitempty"`

	// Use lightweight, invisible JavaScript detections to improve Bot Management. [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs *bool `json:"enableJs,omitempty" tf:"enable_js,omitempty"`

	// Whether to enable Bot Fight Mode.
	FightMode *bool `json:"fightMode,omitempty" tf:"fight_mode,omitempty"`

	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress *bool `json:"optimizeWordpress,omitempty" tf:"optimize_wordpress,omitempty"`

	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated *string `json:"sbfmDefinitelyAutomated,omitempty" tf:"sbfm_definitely_automated,omitempty"`

	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SbfmLikelyAutomated *string `json:"sbfmLikelyAutomated,omitempty" tf:"sbfm_likely_automated,omitempty"`

	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if static resources on your application need bot protection. Note: Static resource protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection *bool `json:"sbfmStaticResourceProtection,omitempty" tf:"sbfm_static_resource_protection,omitempty"`

	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots *string `json:"sbfmVerifiedBots,omitempty" tf:"sbfm_verified_bots,omitempty"`

	// Whether to disable tracking the highest bot score for a session in the Bot Management cookie.
	SuppressSessionScore *bool `json:"suppressSessionScore,omitempty" tf:"suppress_session_score,omitempty"`
}

type ManagementObservation struct {

	// Automatically update to the newest bot detection models created by Cloudflare as they are released. [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes).
	AutoUpdateModel *bool `json:"autoUpdateModel,omitempty" tf:"auto_update_model,omitempty"`

	// Use lightweight, invisible JavaScript detections to improve Bot Management. [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs *bool `json:"enableJs,omitempty" tf:"enable_js,omitempty"`

	// Whether to enable Bot Fight Mode.
	FightMode *bool `json:"fightMode,omitempty" tf:"fight_mode,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress *bool `json:"optimizeWordpress,omitempty" tf:"optimize_wordpress,omitempty"`

	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated *string `json:"sbfmDefinitelyAutomated,omitempty" tf:"sbfm_definitely_automated,omitempty"`

	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SbfmLikelyAutomated *string `json:"sbfmLikelyAutomated,omitempty" tf:"sbfm_likely_automated,omitempty"`

	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if static resources on your application need bot protection. Note: Static resource protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection *bool `json:"sbfmStaticResourceProtection,omitempty" tf:"sbfm_static_resource_protection,omitempty"`

	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots *string `json:"sbfmVerifiedBots,omitempty" tf:"sbfm_verified_bots,omitempty"`

	// Whether to disable tracking the highest bot score for a session in the Bot Management cookie.
	SuppressSessionScore *bool `json:"suppressSessionScore,omitempty" tf:"suppress_session_score,omitempty"`

	// A read-only field that indicates whether the zone currently is running the latest ML model.
	UsingLatestModel *bool `json:"usingLatestModel,omitempty" tf:"using_latest_model,omitempty"`

	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type ManagementParameters struct {

	// Automatically update to the newest bot detection models created by Cloudflare as they are released. [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes).
	// +kubebuilder:validation:Optional
	AutoUpdateModel *bool `json:"autoUpdateModel,omitempty" tf:"auto_update_model,omitempty"`

	// Use lightweight, invisible JavaScript detections to improve Bot Management. [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	// +kubebuilder:validation:Optional
	EnableJs *bool `json:"enableJs,omitempty" tf:"enable_js,omitempty"`

	// Whether to enable Bot Fight Mode.
	// +kubebuilder:validation:Optional
	FightMode *bool `json:"fightMode,omitempty" tf:"fight_mode,omitempty"`

	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	// +kubebuilder:validation:Optional
	OptimizeWordpress *bool `json:"optimizeWordpress,omitempty" tf:"optimize_wordpress,omitempty"`

	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	// +kubebuilder:validation:Optional
	SbfmDefinitelyAutomated *string `json:"sbfmDefinitelyAutomated,omitempty" tf:"sbfm_definitely_automated,omitempty"`

	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	// +kubebuilder:validation:Optional
	SbfmLikelyAutomated *string `json:"sbfmLikelyAutomated,omitempty" tf:"sbfm_likely_automated,omitempty"`

	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if static resources on your application need bot protection. Note: Static resource protection can also result in legitimate traffic being blocked.
	// +kubebuilder:validation:Optional
	SbfmStaticResourceProtection *bool `json:"sbfmStaticResourceProtection,omitempty" tf:"sbfm_static_resource_protection,omitempty"`

	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	// +kubebuilder:validation:Optional
	SbfmVerifiedBots *string `json:"sbfmVerifiedBots,omitempty" tf:"sbfm_verified_bots,omitempty"`

	// Whether to disable tracking the highest bot score for a session in the Bot Management cookie.
	// +kubebuilder:validation:Optional
	SuppressSessionScore *bool `json:"suppressSessionScore,omitempty" tf:"suppress_session_score,omitempty"`

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

// ManagementSpec defines the desired state of Management
type ManagementSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagementParameters `json:"forProvider"`
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
	InitProvider ManagementInitParameters `json:"initProvider,omitempty"`
}

// ManagementStatus defines the observed state of Management.
type ManagementStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagementObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Management is the Schema for the Managements API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Management struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagementSpec   `json:"spec"`
	Status            ManagementStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementList contains a list of Managements
type ManagementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Management `json:"items"`
}

// Repository type metadata.
var (
	Management_Kind             = "Management"
	Management_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Management_Kind}.String()
	Management_KindAPIVersion   = Management_Kind + "." + CRDGroupVersion.String()
	Management_GroupVersionKind = CRDGroupVersion.WithKind(Management_Kind)
)

func init() {
	SchemeBuilder.Register(&Management{}, &ManagementList{})
}

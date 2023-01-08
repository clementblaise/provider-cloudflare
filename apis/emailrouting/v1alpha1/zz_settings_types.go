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

type SettingsObservation struct {

	// The date and time the settings have been created.
	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The date and time the settings have been modified.
	Modified *string `json:"modified,omitempty" tf:"modified,omitempty"`

	// Domain of your zone.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Show the state of your account, and the type or configuration error.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Email Routing settings identifier.
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

type SettingsParameters struct {

	// State of the zone settings for Email Routing. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// Flag to check if the user skipped the configuration wizard.
	// +kubebuilder:validation:Optional
	SkipWizard *bool `json:"skipWizard,omitempty" tf:"skip_wizard,omitempty"`

	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**
	// +crossplane:generate:reference:type=github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1.Zone
	// +crossplane:generate:reference:refFieldName=ZoneIDRefs
	// +crossplane:generate:reference:selectorFieldName=ZoneIDSelector
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a Zone in zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDRefs *v1.Reference `json:"zoneIdRefs,omitempty" tf:"-"`

	// Selector for a Zone in zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

// SettingsSpec defines the desired state of Settings
type SettingsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SettingsParameters `json:"forProvider"`
}

// SettingsStatus defines the observed state of Settings.
type SettingsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SettingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Settings is the Schema for the Settingss API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Settings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SettingsSpec   `json:"spec"`
	Status            SettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SettingsList contains a list of Settingss
type SettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Settings `json:"items"`
}

// Repository type metadata.
var (
	Settings_Kind             = "Settings"
	Settings_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Settings_Kind}.String()
	Settings_KindAPIVersion   = Settings_Kind + "." + CRDGroupVersion.String()
	Settings_GroupVersionKind = CRDGroupVersion.WithKind(Settings_Kind)
)

func init() {
	SchemeBuilder.Register(&Settings{}, &SettingsList{})
}
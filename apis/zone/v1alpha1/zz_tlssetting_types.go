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

type TLSSettingInitParameters struct {

	// Hostname that belongs to this zone name. **Modifying this attribute will force creation of a new resource.**
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// TLS setting name. **Modifying this attribute will force creation of a new resource.**
	Setting *string `json:"setting,omitempty" tf:"setting,omitempty"`

	// TLS setting value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TLSSettingObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Hostname that belongs to this zone name. **Modifying this attribute will force creation of a new resource.**
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// TLS setting name. **Modifying this attribute will force creation of a new resource.**
	Setting *string `json:"setting,omitempty" tf:"setting,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	// TLS setting value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type TLSSettingParameters struct {

	// Hostname that belongs to this zone name. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// TLS setting name. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	Setting *string `json:"setting,omitempty" tf:"setting,omitempty"`

	// TLS setting value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**
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

// TLSSettingSpec defines the desired state of TLSSetting
type TLSSettingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TLSSettingParameters `json:"forProvider"`
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
	InitProvider TLSSettingInitParameters `json:"initProvider,omitempty"`
}

// TLSSettingStatus defines the observed state of TLSSetting.
type TLSSettingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TLSSettingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TLSSetting is the Schema for the TLSSettings API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type TLSSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hostname) || has(self.initProvider.hostname)",message="hostname is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.setting) || has(self.initProvider.setting)",message="setting is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.value) || has(self.initProvider.value)",message="value is a required parameter"
	Spec   TLSSettingSpec   `json:"spec"`
	Status TLSSettingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TLSSettingList contains a list of TLSSettings
type TLSSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TLSSetting `json:"items"`
}

// Repository type metadata.
var (
	TLSSetting_Kind             = "TLSSetting"
	TLSSetting_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TLSSetting_Kind}.String()
	TLSSetting_KindAPIVersion   = TLSSetting_Kind + "." + CRDGroupVersion.String()
	TLSSetting_GroupVersionKind = CRDGroupVersion.WithKind(TLSSetting_Kind)
)

func init() {
	SchemeBuilder.Register(&TLSSetting{}, &TLSSettingList{})
}

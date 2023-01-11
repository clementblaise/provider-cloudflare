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

type DevicePostureRuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DevicePostureRuleParameters struct {

	// The account identifier to target for the resource.
	// +crossplane:generate:reference:type=github.com/cdloh/provider-cloudflare/apis/account/v1alpha1.Account
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Reference to a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDRef *v1.Reference `json:"accountIdRef,omitempty" tf:"-"`

	// Selector for a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDSelector *v1.Selector `json:"accountIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Expire posture results after the specified amount of time. Must be in the format `1h` or `30m`. Valid units are `h` and `m`.
	// +kubebuilder:validation:Optional
	Expiration *string `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// +kubebuilder:validation:Optional
	Input []InputParameters `json:"input,omitempty" tf:"input,omitempty"`

	// The conditions that the client must match to run the rule.
	// +kubebuilder:validation:Optional
	Match []MatchParameters `json:"match,omitempty" tf:"match,omitempty"`

	// Name of the device posture rule.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Tells the client when to run the device posture check. Must be in the format `1h` or `30m`. Valid units are `h` and `m`.
	// +kubebuilder:validation:Optional
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// The device posture rule type. Available values: `serial_number`, `file`, `application`, `gateway`, `warp`, `domain_joined`, `os_version`, `disk_encryption`, `firewall`, `workspace_one`, `unique_client_id`.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type InputObservation struct {
}

type InputParameters struct {

	// The workspace one device compliance status.
	// +kubebuilder:validation:Optional
	ComplianceStatus *string `json:"complianceStatus,omitempty" tf:"compliance_status,omitempty"`

	// The workspace one connection id.
	// +kubebuilder:validation:Optional
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// The domain that the client must join.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// True if the firewall must be enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Checks if the file should exist.
	// +kubebuilder:validation:Optional
	Exists *bool `json:"exists,omitempty" tf:"exists,omitempty"`

	// The Teams List id.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The version comparison operator.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// The operating system excluding version information.
	// +kubebuilder:validation:Optional
	OsDistroName *string `json:"osDistroName,omitempty" tf:"os_distro_name,omitempty"`

	// The operating system version excluding OS name information or release name.
	// +kubebuilder:validation:Optional
	OsDistroRevision *string `json:"osDistroRevision,omitempty" tf:"os_distro_revision,omitempty"`

	// The path to the file.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// True if all drives must be encrypted.
	// +kubebuilder:validation:Optional
	RequireAll *bool `json:"requireAll,omitempty" tf:"require_all,omitempty"`

	// Checks if the application should be running.
	// +kubebuilder:validation:Optional
	Running *bool `json:"running,omitempty" tf:"running,omitempty"`

	// The sha256 hash of the file.
	// +kubebuilder:validation:Optional
	Sha256 *string `json:"sha256,omitempty" tf:"sha256,omitempty"`

	// The thumbprint of the file certificate.
	// +kubebuilder:validation:Optional
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`

	// The operating system semantic version.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type MatchObservation struct {
}

type MatchParameters struct {

	// The platform of the device. Available values: `windows`, `mac`, `linux`, `android`, `ios`, `chromeos`.
	// +kubebuilder:validation:Optional
	Platform *string `json:"platform,omitempty" tf:"platform,omitempty"`
}

// DevicePostureRuleSpec defines the desired state of DevicePostureRule
type DevicePostureRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DevicePostureRuleParameters `json:"forProvider"`
}

// DevicePostureRuleStatus defines the observed state of DevicePostureRule.
type DevicePostureRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DevicePostureRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DevicePostureRule is the Schema for the DevicePostureRules API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type DevicePostureRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DevicePostureRuleSpec   `json:"spec"`
	Status            DevicePostureRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DevicePostureRuleList contains a list of DevicePostureRules
type DevicePostureRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevicePostureRule `json:"items"`
}

// Repository type metadata.
var (
	DevicePostureRule_Kind             = "DevicePostureRule"
	DevicePostureRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DevicePostureRule_Kind}.String()
	DevicePostureRule_KindAPIVersion   = DevicePostureRule_Kind + "." + CRDGroupVersion.String()
	DevicePostureRule_GroupVersionKind = CRDGroupVersion.WithKind(DevicePostureRule_Kind)
)

func init() {
	SchemeBuilder.Register(&DevicePostureRule{}, &DevicePostureRuleList{})
}

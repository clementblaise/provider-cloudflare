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

type HeaderObservation struct {
}

type HeaderParameters struct {

	// The header name.
	// +kubebuilder:validation:Required
	Header *string `json:"header" tf:"header,omitempty"`

	// A list of string values for the header.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type HealthcheckObservation struct {

	// Creation time.
	CreatedOn *string `json:"createdOn,omitempty" tf:"created_on,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Last modified time.
	ModifiedOn *string `json:"modifiedOn,omitempty" tf:"modified_on,omitempty"`
}

type HealthcheckParameters struct {

	// The hostname or IP address of the origin server to run health checks on.
	// +kubebuilder:validation:Required
	Address *string `json:"address" tf:"address,omitempty"`

	// Do not validate the certificate when the health check uses HTTPS. Defaults to `false`.
	// +kubebuilder:validation:Optional
	AllowInsecure *bool `json:"allowInsecure,omitempty" tf:"allow_insecure,omitempty"`

	// A list of regions from which to run health checks. If not set, Cloudflare will pick a default region. Available values: `WNAM`, `ENAM`, `WEU`, `EEU`, `NSAM`, `SSAM`, `OC`, `ME`, `NAF`, `SAF`, `IN`, `SEAS`, `NEAS`, `ALL_REGIONS`.
	// +kubebuilder:validation:Optional
	CheckRegions []*string `json:"checkRegions,omitempty" tf:"check_regions,omitempty"`

	// The number of consecutive fails required from a health check before changing the health to unhealthy. Defaults to `1`.
	// +kubebuilder:validation:Optional
	ConsecutiveFails *float64 `json:"consecutiveFails,omitempty" tf:"consecutive_fails,omitempty"`

	// The number of consecutive successes required from a health check before changing the health to healthy. Defaults to `1`.
	// +kubebuilder:validation:Optional
	ConsecutiveSuccesses *float64 `json:"consecutiveSuccesses,omitempty" tf:"consecutive_successes,omitempty"`

	// A human-readable description of the health check.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A case-insensitive sub-string to look for in the response body. If this string is not found the origin will be marked as unhealthy.
	// +kubebuilder:validation:Optional
	ExpectedBody *string `json:"expectedBody,omitempty" tf:"expected_body,omitempty"`

	// The expected HTTP response codes (e.g. '200') or code ranges (e.g. '2xx' for all codes starting with 2) of the health check.
	// +kubebuilder:validation:Optional
	ExpectedCodes []*string `json:"expectedCodes,omitempty" tf:"expected_codes,omitempty"`

	// Follow redirects if the origin returns a 3xx status code. Defaults to `false`.
	// +kubebuilder:validation:Optional
	FollowRedirects *bool `json:"followRedirects,omitempty" tf:"follow_redirects,omitempty"`

	// The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The User-Agent header cannot be overridden.
	// +kubebuilder:validation:Optional
	Header []HeaderParameters `json:"header,omitempty" tf:"header,omitempty"`

	// The interval between each health check. Shorter intervals may give quicker notifications if the origin status changes, but will increase the load on the origin as we check from multiple locations. Defaults to `60`.
	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// The HTTP method to use for the health check. Available values: `connection_established`, `GET`, `HEAD`.
	// +kubebuilder:validation:Optional
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// A short name to identify the health check. Only alphanumeric characters, hyphens, and underscores are allowed.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The endpoint path to health check against. Defaults to `/`.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Port number to connect to for the health check. Defaults to `80`.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted immediately. Defaults to `2`.
	// +kubebuilder:validation:Optional
	Retries *float64 `json:"retries,omitempty" tf:"retries,omitempty"`

	// If suspended, no health checks are sent to the origin. Defaults to `false`.
	// +kubebuilder:validation:Optional
	Suspended *bool `json:"suspended,omitempty" tf:"suspended,omitempty"`

	// The timeout (in seconds) before marking the health check as failed. Defaults to `5`.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// The protocol to use for the health check. Available values: `TCP`, `HTTP`, `HTTPS`.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

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

// HealthcheckSpec defines the desired state of Healthcheck
type HealthcheckSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HealthcheckParameters `json:"forProvider"`
}

// HealthcheckStatus defines the observed state of Healthcheck.
type HealthcheckStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HealthcheckObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Healthcheck is the Schema for the Healthchecks API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Healthcheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HealthcheckSpec   `json:"spec"`
	Status            HealthcheckStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HealthcheckList contains a list of Healthchecks
type HealthcheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Healthcheck `json:"items"`
}

// Repository type metadata.
var (
	Healthcheck_Kind             = "Healthcheck"
	Healthcheck_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Healthcheck_Kind}.String()
	Healthcheck_KindAPIVersion   = Healthcheck_Kind + "." + CRDGroupVersion.String()
	Healthcheck_GroupVersionKind = CRDGroupVersion.WithKind(Healthcheck_Kind)
)

func init() {
	SchemeBuilder.Register(&Healthcheck{}, &HealthcheckList{})
}

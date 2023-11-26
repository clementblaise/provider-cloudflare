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

type ActionInitParameters struct {

	// The type of action to perform. Allowable values are 'simulate', 'ban', 'challenge', 'js_challenge' and 'managed_challenge'.
	// The type of action to perform. Available values: `simulate`, `ban`, `challenge`, `js_challenge`, `managed_challenge`.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Matches HTTP responses before they are returned to the client from Cloudflare. If this is defined, then the entire counting of traffic occurs at this stage. This field is not required.
	// Custom content-type and body to return, this overrides the custom error for the zone. This field is not required. Omission will result in default HTML error page.
	Response []ResponseInitParameters `json:"response,omitempty" tf:"response,omitempty"`

	// The time in seconds as an integer to perform the mitigation action. This field is required if the mode is either simulate or ban. Must be the same or greater than the period (min: 1, max: 86400).
	// The time in seconds as an integer to perform the mitigation action. This field is required if the `mode` is either `simulate` or `ban`. Must be the same or greater than the period.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type ActionObservation struct {

	// The type of action to perform. Allowable values are 'simulate', 'ban', 'challenge', 'js_challenge' and 'managed_challenge'.
	// The type of action to perform. Available values: `simulate`, `ban`, `challenge`, `js_challenge`, `managed_challenge`.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Matches HTTP responses before they are returned to the client from Cloudflare. If this is defined, then the entire counting of traffic occurs at this stage. This field is not required.
	// Custom content-type and body to return, this overrides the custom error for the zone. This field is not required. Omission will result in default HTML error page.
	Response []ResponseObservation `json:"response,omitempty" tf:"response,omitempty"`

	// The time in seconds as an integer to perform the mitigation action. This field is required if the mode is either simulate or ban. Must be the same or greater than the period (min: 1, max: 86400).
	// The time in seconds as an integer to perform the mitigation action. This field is required if the `mode` is either `simulate` or `ban`. Must be the same or greater than the period.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type ActionParameters struct {

	// The type of action to perform. Allowable values are 'simulate', 'ban', 'challenge', 'js_challenge' and 'managed_challenge'.
	// The type of action to perform. Available values: `simulate`, `ban`, `challenge`, `js_challenge`, `managed_challenge`.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Matches HTTP responses before they are returned to the client from Cloudflare. If this is defined, then the entire counting of traffic occurs at this stage. This field is not required.
	// Custom content-type and body to return, this overrides the custom error for the zone. This field is not required. Omission will result in default HTML error page.
	// +kubebuilder:validation:Optional
	Response []ResponseParameters `json:"response,omitempty" tf:"response,omitempty"`

	// The time in seconds as an integer to perform the mitigation action. This field is required if the mode is either simulate or ban. Must be the same or greater than the period (min: 1, max: 86400).
	// The time in seconds as an integer to perform the mitigation action. This field is required if the `mode` is either `simulate` or `ban`. Must be the same or greater than the period.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type CorrelateInitParameters struct {

	// If set to 'nat', NAT support will be enabled for rate limiting.
	// If set to 'nat', NAT support will be enabled for rate limiting. Available values: `nat`.
	By *string `json:"by,omitempty" tf:"by,omitempty"`
}

type CorrelateObservation struct {

	// If set to 'nat', NAT support will be enabled for rate limiting.
	// If set to 'nat', NAT support will be enabled for rate limiting. Available values: `nat`.
	By *string `json:"by,omitempty" tf:"by,omitempty"`
}

type CorrelateParameters struct {

	// If set to 'nat', NAT support will be enabled for rate limiting.
	// If set to 'nat', NAT support will be enabled for rate limiting. Available values: `nat`.
	// +kubebuilder:validation:Optional
	By *string `json:"by,omitempty" tf:"by,omitempty"`
}

type MatchInitParameters struct {

	// Matches HTTP requests (from the client to Cloudflare). See definition below.
	// Matches HTTP requests (from the client to Cloudflare).
	Request []RequestInitParameters `json:"request,omitempty" tf:"request,omitempty"`

	// Matches HTTP responses before they are returned to the client from Cloudflare. If this is defined, then the entire counting of traffic occurs at this stage. This field is not required.
	// Matches HTTP responses before they are returned to the client from Cloudflare. If this is defined, then the entire counting of traffic occurs at this stage.
	Response []MatchResponseInitParameters `json:"response,omitempty" tf:"response,omitempty"`
}

type MatchObservation struct {

	// Matches HTTP requests (from the client to Cloudflare). See definition below.
	// Matches HTTP requests (from the client to Cloudflare).
	Request []RequestObservation `json:"request,omitempty" tf:"request,omitempty"`

	// Matches HTTP responses before they are returned to the client from Cloudflare. If this is defined, then the entire counting of traffic occurs at this stage. This field is not required.
	// Matches HTTP responses before they are returned to the client from Cloudflare. If this is defined, then the entire counting of traffic occurs at this stage.
	Response []MatchResponseObservation `json:"response,omitempty" tf:"response,omitempty"`
}

type MatchParameters struct {

	// Matches HTTP requests (from the client to Cloudflare). See definition below.
	// Matches HTTP requests (from the client to Cloudflare).
	// +kubebuilder:validation:Optional
	Request []RequestParameters `json:"request,omitempty" tf:"request,omitempty"`

	// Matches HTTP responses before they are returned to the client from Cloudflare. If this is defined, then the entire counting of traffic occurs at this stage. This field is not required.
	// Matches HTTP responses before they are returned to the client from Cloudflare. If this is defined, then the entire counting of traffic occurs at this stage.
	// +kubebuilder:validation:Optional
	Response []MatchResponseParameters `json:"response,omitempty" tf:"response,omitempty"`
}

type MatchResponseInitParameters struct {

	// block is a list of maps with the following attributes:
	// List of HTTP headers maps to match the origin response on.
	Headers []map[string]*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// Only count traffic that has come from your origin servers. If true, cached items that Cloudflare serve will not count towards rate limiting. Default: true.
	// Only count traffic that has come from your origin servers. If true, cached items that Cloudflare serve will not count towards rate limiting.
	OriginTraffic *bool `json:"originTraffic,omitempty" tf:"origin_traffic,omitempty"`

	// HTTP Status codes, can be one [403], many [401,403] or indicate all by not providing this value.
	// HTTP Status codes, can be one, many or indicate all by not providing this value.
	Statuses []*float64 `json:"statuses,omitempty" tf:"statuses,omitempty"`
}

type MatchResponseObservation struct {

	// block is a list of maps with the following attributes:
	// List of HTTP headers maps to match the origin response on.
	Headers []map[string]*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// Only count traffic that has come from your origin servers. If true, cached items that Cloudflare serve will not count towards rate limiting. Default: true.
	// Only count traffic that has come from your origin servers. If true, cached items that Cloudflare serve will not count towards rate limiting.
	OriginTraffic *bool `json:"originTraffic,omitempty" tf:"origin_traffic,omitempty"`

	// HTTP Status codes, can be one [403], many [401,403] or indicate all by not providing this value.
	// HTTP Status codes, can be one, many or indicate all by not providing this value.
	Statuses []*float64 `json:"statuses,omitempty" tf:"statuses,omitempty"`
}

type MatchResponseParameters struct {

	// block is a list of maps with the following attributes:
	// List of HTTP headers maps to match the origin response on.
	// +kubebuilder:validation:Optional
	Headers []map[string]*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// Only count traffic that has come from your origin servers. If true, cached items that Cloudflare serve will not count towards rate limiting. Default: true.
	// Only count traffic that has come from your origin servers. If true, cached items that Cloudflare serve will not count towards rate limiting.
	// +kubebuilder:validation:Optional
	OriginTraffic *bool `json:"originTraffic,omitempty" tf:"origin_traffic,omitempty"`

	// HTTP Status codes, can be one [403], many [401,403] or indicate all by not providing this value.
	// HTTP Status codes, can be one, many or indicate all by not providing this value.
	// +kubebuilder:validation:Optional
	Statuses []*float64 `json:"statuses,omitempty" tf:"statuses,omitempty"`
}

type RateLimitInitParameters struct {

	// The action to be performed when the threshold of matched traffic within the period defined is exceeded.
	// The action to be performed when the threshold of matched traffic within the period defined is exceeded.
	Action []ActionInitParameters `json:"action,omitempty" tf:"action,omitempty"`

	// URLs matching the patterns specified here will be excluded from rate limiting.
	BypassURLPatterns []*string `json:"bypassUrlPatterns,omitempty" tf:"bypass_url_patterns,omitempty"`

	// Determines how rate limiting is applied. By default if not specified, rate limiting applies to the clients IP address.
	// Determines how rate limiting is applied. By default if not specified, rate limiting applies to the clients IP address.
	Correlate []CorrelateInitParameters `json:"correlate,omitempty" tf:"correlate,omitempty"`

	// A note that you can use to describe the reason for a rate limit. This value is sanitized and all tags are removed.
	// A note that you can use to describe the reason for a rate limit. This value is sanitized and all tags are removed.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether this ratelimit is currently disabled. Default: false.
	// Whether this ratelimit is currently disabled. Defaults to `false`.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// Determines which traffic the rate limit counts towards the threshold. By default matches all traffic in the zone. See definition below.
	// Determines which traffic the rate limit counts towards the threshold. By default matches all traffic in the zone.
	Match []MatchInitParameters `json:"match,omitempty" tf:"match,omitempty"`

	// The time in seconds to count matching traffic. If the count exceeds threshold within this period the action will be performed (min: 1, max: 86,400).
	// The time in seconds to count matching traffic. If the count exceeds threshold within this period the action will be performed.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// The threshold that triggers the rate limit mitigations, combine with period. i.e. threshold per period (min: 2, max: 1,000,000).
	// The threshold that triggers the rate limit mitigations, combine with period.
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`
}

type RateLimitObservation struct {

	// The action to be performed when the threshold of matched traffic within the period defined is exceeded.
	// The action to be performed when the threshold of matched traffic within the period defined is exceeded.
	Action []ActionObservation `json:"action,omitempty" tf:"action,omitempty"`

	// URLs matching the patterns specified here will be excluded from rate limiting.
	BypassURLPatterns []*string `json:"bypassUrlPatterns,omitempty" tf:"bypass_url_patterns,omitempty"`

	// Determines how rate limiting is applied. By default if not specified, rate limiting applies to the clients IP address.
	// Determines how rate limiting is applied. By default if not specified, rate limiting applies to the clients IP address.
	Correlate []CorrelateObservation `json:"correlate,omitempty" tf:"correlate,omitempty"`

	// A note that you can use to describe the reason for a rate limit. This value is sanitized and all tags are removed.
	// A note that you can use to describe the reason for a rate limit. This value is sanitized and all tags are removed.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether this ratelimit is currently disabled. Default: false.
	// Whether this ratelimit is currently disabled. Defaults to `false`.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The Rate limit ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Determines which traffic the rate limit counts towards the threshold. By default matches all traffic in the zone. See definition below.
	// Determines which traffic the rate limit counts towards the threshold. By default matches all traffic in the zone.
	Match []MatchObservation `json:"match,omitempty" tf:"match,omitempty"`

	// The time in seconds to count matching traffic. If the count exceeds threshold within this period the action will be performed (min: 1, max: 86,400).
	// The time in seconds to count matching traffic. If the count exceeds threshold within this period the action will be performed.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// The threshold that triggers the rate limit mitigations, combine with period. i.e. threshold per period (min: 2, max: 1,000,000).
	// The threshold that triggers the rate limit mitigations, combine with period.
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// The DNS zone ID to apply rate limiting to.
	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type RateLimitParameters struct {

	// The action to be performed when the threshold of matched traffic within the period defined is exceeded.
	// The action to be performed when the threshold of matched traffic within the period defined is exceeded.
	// +kubebuilder:validation:Optional
	Action []ActionParameters `json:"action,omitempty" tf:"action,omitempty"`

	// URLs matching the patterns specified here will be excluded from rate limiting.
	// +kubebuilder:validation:Optional
	BypassURLPatterns []*string `json:"bypassUrlPatterns,omitempty" tf:"bypass_url_patterns,omitempty"`

	// Determines how rate limiting is applied. By default if not specified, rate limiting applies to the clients IP address.
	// Determines how rate limiting is applied. By default if not specified, rate limiting applies to the clients IP address.
	// +kubebuilder:validation:Optional
	Correlate []CorrelateParameters `json:"correlate,omitempty" tf:"correlate,omitempty"`

	// A note that you can use to describe the reason for a rate limit. This value is sanitized and all tags are removed.
	// A note that you can use to describe the reason for a rate limit. This value is sanitized and all tags are removed.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether this ratelimit is currently disabled. Default: false.
	// Whether this ratelimit is currently disabled. Defaults to `false`.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// Determines which traffic the rate limit counts towards the threshold. By default matches all traffic in the zone. See definition below.
	// Determines which traffic the rate limit counts towards the threshold. By default matches all traffic in the zone.
	// +kubebuilder:validation:Optional
	Match []MatchParameters `json:"match,omitempty" tf:"match,omitempty"`

	// The time in seconds to count matching traffic. If the count exceeds threshold within this period the action will be performed (min: 1, max: 86,400).
	// The time in seconds to count matching traffic. If the count exceeds threshold within this period the action will be performed.
	// +kubebuilder:validation:Optional
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// The threshold that triggers the rate limit mitigations, combine with period. i.e. threshold per period (min: 2, max: 1,000,000).
	// The threshold that triggers the rate limit mitigations, combine with period.
	// +kubebuilder:validation:Optional
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// The DNS zone ID to apply rate limiting to.
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

type RequestInitParameters struct {

	// HTTP Methods, can be a subset ['POST','PUT'] or all ['_ALL_']. Default: ['_ALL_'].
	// HTTP Methods to match traffic on. Available values: `GET`, `POST`, `PUT`, `DELETE`, `PATCH`, `HEAD`, `_ALL_`.
	Methods []*string `json:"methods,omitempty" tf:"methods,omitempty"`

	// HTTP Schemes, can be one ['HTTPS'], both ['HTTP','HTTPS'] or all ['_ALL_']. Default: ['_ALL_'].
	// HTTP schemes to match traffic on. Available values: `HTTP`, `HTTPS`, `_ALL_`.
	Schemes []*string `json:"schemes,omitempty" tf:"schemes,omitempty"`

	// The URL pattern to match comprised of the host and path, i.e. example.org/path. Wildcard are expanded to match applicable traffic, query strings are not matched. Use _ for all traffic to your zone. Default: '_'.
	// The URL pattern to match comprised of the host and path, i.e. example.org/path. Wildcard are expanded to match applicable traffic, query strings are not matched. Use _ for all traffic to your zone.
	URLPattern *string `json:"urlPattern,omitempty" tf:"url_pattern,omitempty"`
}

type RequestObservation struct {

	// HTTP Methods, can be a subset ['POST','PUT'] or all ['_ALL_']. Default: ['_ALL_'].
	// HTTP Methods to match traffic on. Available values: `GET`, `POST`, `PUT`, `DELETE`, `PATCH`, `HEAD`, `_ALL_`.
	Methods []*string `json:"methods,omitempty" tf:"methods,omitempty"`

	// HTTP Schemes, can be one ['HTTPS'], both ['HTTP','HTTPS'] or all ['_ALL_']. Default: ['_ALL_'].
	// HTTP schemes to match traffic on. Available values: `HTTP`, `HTTPS`, `_ALL_`.
	Schemes []*string `json:"schemes,omitempty" tf:"schemes,omitempty"`

	// The URL pattern to match comprised of the host and path, i.e. example.org/path. Wildcard are expanded to match applicable traffic, query strings are not matched. Use _ for all traffic to your zone. Default: '_'.
	// The URL pattern to match comprised of the host and path, i.e. example.org/path. Wildcard are expanded to match applicable traffic, query strings are not matched. Use _ for all traffic to your zone.
	URLPattern *string `json:"urlPattern,omitempty" tf:"url_pattern,omitempty"`
}

type RequestParameters struct {

	// HTTP Methods, can be a subset ['POST','PUT'] or all ['_ALL_']. Default: ['_ALL_'].
	// HTTP Methods to match traffic on. Available values: `GET`, `POST`, `PUT`, `DELETE`, `PATCH`, `HEAD`, `_ALL_`.
	// +kubebuilder:validation:Optional
	Methods []*string `json:"methods,omitempty" tf:"methods,omitempty"`

	// HTTP Schemes, can be one ['HTTPS'], both ['HTTP','HTTPS'] or all ['_ALL_']. Default: ['_ALL_'].
	// HTTP schemes to match traffic on. Available values: `HTTP`, `HTTPS`, `_ALL_`.
	// +kubebuilder:validation:Optional
	Schemes []*string `json:"schemes,omitempty" tf:"schemes,omitempty"`

	// The URL pattern to match comprised of the host and path, i.e. example.org/path. Wildcard are expanded to match applicable traffic, query strings are not matched. Use _ for all traffic to your zone. Default: '_'.
	// The URL pattern to match comprised of the host and path, i.e. example.org/path. Wildcard are expanded to match applicable traffic, query strings are not matched. Use _ for all traffic to your zone.
	// +kubebuilder:validation:Optional
	URLPattern *string `json:"urlPattern,omitempty" tf:"url_pattern,omitempty"`
}

type ResponseInitParameters struct {

	// The body to return, the content here should conform to the content_type.
	// The body to return, the content here should conform to the `content_type`.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// The content-type of the body, must be one of: 'text/plain', 'text/xml', 'application/json'.
	// The content-type of the body. Available values: `text/plain`, `text/xml`, `application/json`.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`
}

type ResponseObservation struct {

	// The body to return, the content here should conform to the content_type.
	// The body to return, the content here should conform to the `content_type`.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// The content-type of the body, must be one of: 'text/plain', 'text/xml', 'application/json'.
	// The content-type of the body. Available values: `text/plain`, `text/xml`, `application/json`.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`
}

type ResponseParameters struct {

	// The body to return, the content here should conform to the content_type.
	// The body to return, the content here should conform to the `content_type`.
	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// The content-type of the body, must be one of: 'text/plain', 'text/xml', 'application/json'.
	// The content-type of the body. Available values: `text/plain`, `text/xml`, `application/json`.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`
}

// RateLimitSpec defines the desired state of RateLimit
type RateLimitSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RateLimitParameters `json:"forProvider"`
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
	InitProvider RateLimitInitParameters `json:"initProvider,omitempty"`
}

// RateLimitStatus defines the observed state of RateLimit.
type RateLimitStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RateLimitObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RateLimit is the Schema for the RateLimits API. Provides a Cloudflare rate limit resource for a particular zone.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type RateLimit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.action) || has(self.initProvider.action)",message="action is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.period) || has(self.initProvider.period)",message="period is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.threshold) || has(self.initProvider.threshold)",message="threshold is a required parameter"
	Spec   RateLimitSpec   `json:"spec"`
	Status RateLimitStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RateLimitList contains a list of RateLimits
type RateLimitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RateLimit `json:"items"`
}

// Repository type metadata.
var (
	RateLimit_Kind             = "RateLimit"
	RateLimit_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RateLimit_Kind}.String()
	RateLimit_KindAPIVersion   = RateLimit_Kind + "." + CRDGroupVersion.String()
	RateLimit_GroupVersionKind = CRDGroupVersion.WithKind(RateLimit_Kind)
)

func init() {
	SchemeBuilder.Register(&RateLimit{}, &RateLimitList{})
}

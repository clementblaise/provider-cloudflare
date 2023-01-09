/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/cdloh/provider-cloudflare/apis/account/v1alpha1"
	v1alpha1apishield "github.com/cdloh/provider-cloudflare/apis/apishield/v1alpha1"
	v1alpha1argo "github.com/cdloh/provider-cloudflare/apis/argo/v1alpha1"
	v1alpha1authenticatedoriginpulls "github.com/cdloh/provider-cloudflare/apis/authenticatedoriginpulls/v1alpha1"
	v1alpha1byoip "github.com/cdloh/provider-cloudflare/apis/byoip/v1alpha1"
	v1alpha1certificate "github.com/cdloh/provider-cloudflare/apis/certificate/v1alpha1"
	v1alpha1custom "github.com/cdloh/provider-cloudflare/apis/custom/v1alpha1"
	v1alpha1customhostname "github.com/cdloh/provider-cloudflare/apis/customhostname/v1alpha1"
	v1alpha1dlp "github.com/cdloh/provider-cloudflare/apis/dlp/v1alpha1"
	v1alpha1dns "github.com/cdloh/provider-cloudflare/apis/dns/v1alpha1"
	v1alpha1emailrouting "github.com/cdloh/provider-cloudflare/apis/emailrouting/v1alpha1"
	v1alpha1loadbalancer "github.com/cdloh/provider-cloudflare/apis/loadbalancer/v1alpha1"
	v1alpha1logpush "github.com/cdloh/provider-cloudflare/apis/logpush/v1alpha1"
	v1alpha1page "github.com/cdloh/provider-cloudflare/apis/page/v1alpha1"
	v1alpha1teams "github.com/cdloh/provider-cloudflare/apis/teams/v1alpha1"
	v1alpha1apis "github.com/cdloh/provider-cloudflare/apis/v1alpha1"
	v1beta1 "github.com/cdloh/provider-cloudflare/apis/v1beta1"
	v1alpha1waf "github.com/cdloh/provider-cloudflare/apis/waf/v1alpha1"
	v1alpha1waitingroom "github.com/cdloh/provider-cloudflare/apis/waitingroom/v1alpha1"
	v1alpha1warp "github.com/cdloh/provider-cloudflare/apis/warp/v1alpha1"
	v1alpha1worker "github.com/cdloh/provider-cloudflare/apis/worker/v1alpha1"
	v1alpha1zone "github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1apishield.SchemeBuilder.AddToScheme,
		v1alpha1argo.SchemeBuilder.AddToScheme,
		v1alpha1authenticatedoriginpulls.SchemeBuilder.AddToScheme,
		v1alpha1byoip.SchemeBuilder.AddToScheme,
		v1alpha1certificate.SchemeBuilder.AddToScheme,
		v1alpha1custom.SchemeBuilder.AddToScheme,
		v1alpha1customhostname.SchemeBuilder.AddToScheme,
		v1alpha1dlp.SchemeBuilder.AddToScheme,
		v1alpha1dns.SchemeBuilder.AddToScheme,
		v1alpha1emailrouting.SchemeBuilder.AddToScheme,
		v1alpha1loadbalancer.SchemeBuilder.AddToScheme,
		v1alpha1logpush.SchemeBuilder.AddToScheme,
		v1alpha1page.SchemeBuilder.AddToScheme,
		v1alpha1teams.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
		v1alpha1waf.SchemeBuilder.AddToScheme,
		v1alpha1waitingroom.SchemeBuilder.AddToScheme,
		v1alpha1warp.SchemeBuilder.AddToScheme,
		v1alpha1worker.SchemeBuilder.AddToScheme,
		v1alpha1zone.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}

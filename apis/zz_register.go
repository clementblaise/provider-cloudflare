/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/cdloh/provider-cloudflare/apis/v1alpha1"
	v1beta1 "github.com/cdloh/provider-cloudflare/apis/v1beta1"
	v1alpha1waf "github.com/cdloh/provider-cloudflare/apis/waf/v1alpha1"
	v1alpha1worker "github.com/cdloh/provider-cloudflare/apis/worker/v1alpha1"
	v1alpha1zone "github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
		v1alpha1waf.SchemeBuilder.AddToScheme,
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

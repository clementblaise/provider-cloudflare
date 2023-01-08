//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigObservation) DeepCopyInto(out *ConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigObservation.
func (in *ConfigObservation) DeepCopy() *ConfigObservation {
	if in == nil {
		return nil
	}
	out := new(ConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigParameters) DeepCopyInto(out *ConfigParameters) {
	*out = *in
	if in.APIURL != nil {
		in, out := &in.APIURL, &out.APIURL
		*out = new(string)
		**out = **in
	}
	if in.AuthURL != nil {
		in, out := &in.AuthURL, &out.AuthURL
		*out = new(string)
		**out = **in
	}
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.ClientKeySecretRef != nil {
		in, out := &in.ClientKeySecretRef, &out.ClientKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ClientSecretSecretRef != nil {
		in, out := &in.ClientSecretSecretRef, &out.ClientSecretSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.CustomerID != nil {
		in, out := &in.CustomerID, &out.CustomerID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigParameters.
func (in *ConfigParameters) DeepCopy() *ConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePolicyCertificates) DeepCopyInto(out *DevicePolicyCertificates) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePolicyCertificates.
func (in *DevicePolicyCertificates) DeepCopy() *DevicePolicyCertificates {
	if in == nil {
		return nil
	}
	out := new(DevicePolicyCertificates)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevicePolicyCertificates) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePolicyCertificatesList) DeepCopyInto(out *DevicePolicyCertificatesList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DevicePolicyCertificates, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePolicyCertificatesList.
func (in *DevicePolicyCertificatesList) DeepCopy() *DevicePolicyCertificatesList {
	if in == nil {
		return nil
	}
	out := new(DevicePolicyCertificatesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevicePolicyCertificatesList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePolicyCertificatesObservation) DeepCopyInto(out *DevicePolicyCertificatesObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePolicyCertificatesObservation.
func (in *DevicePolicyCertificatesObservation) DeepCopy() *DevicePolicyCertificatesObservation {
	if in == nil {
		return nil
	}
	out := new(DevicePolicyCertificatesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePolicyCertificatesParameters) DeepCopyInto(out *DevicePolicyCertificatesParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
	if in.ZoneIDRefs != nil {
		in, out := &in.ZoneIDRefs, &out.ZoneIDRefs
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ZoneIDSelector != nil {
		in, out := &in.ZoneIDSelector, &out.ZoneIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePolicyCertificatesParameters.
func (in *DevicePolicyCertificatesParameters) DeepCopy() *DevicePolicyCertificatesParameters {
	if in == nil {
		return nil
	}
	out := new(DevicePolicyCertificatesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePolicyCertificatesSpec) DeepCopyInto(out *DevicePolicyCertificatesSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePolicyCertificatesSpec.
func (in *DevicePolicyCertificatesSpec) DeepCopy() *DevicePolicyCertificatesSpec {
	if in == nil {
		return nil
	}
	out := new(DevicePolicyCertificatesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePolicyCertificatesStatus) DeepCopyInto(out *DevicePolicyCertificatesStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePolicyCertificatesStatus.
func (in *DevicePolicyCertificatesStatus) DeepCopy() *DevicePolicyCertificatesStatus {
	if in == nil {
		return nil
	}
	out := new(DevicePolicyCertificatesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePostureIntegration) DeepCopyInto(out *DevicePostureIntegration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePostureIntegration.
func (in *DevicePostureIntegration) DeepCopy() *DevicePostureIntegration {
	if in == nil {
		return nil
	}
	out := new(DevicePostureIntegration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevicePostureIntegration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePostureIntegrationList) DeepCopyInto(out *DevicePostureIntegrationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DevicePostureIntegration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePostureIntegrationList.
func (in *DevicePostureIntegrationList) DeepCopy() *DevicePostureIntegrationList {
	if in == nil {
		return nil
	}
	out := new(DevicePostureIntegrationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevicePostureIntegrationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePostureIntegrationObservation) DeepCopyInto(out *DevicePostureIntegrationObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePostureIntegrationObservation.
func (in *DevicePostureIntegrationObservation) DeepCopy() *DevicePostureIntegrationObservation {
	if in == nil {
		return nil
	}
	out := new(DevicePostureIntegrationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePostureIntegrationParameters) DeepCopyInto(out *DevicePostureIntegrationParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make([]ConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Identifier != nil {
		in, out := &in.Identifier, &out.Identifier
		*out = new(string)
		**out = **in
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePostureIntegrationParameters.
func (in *DevicePostureIntegrationParameters) DeepCopy() *DevicePostureIntegrationParameters {
	if in == nil {
		return nil
	}
	out := new(DevicePostureIntegrationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePostureIntegrationSpec) DeepCopyInto(out *DevicePostureIntegrationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePostureIntegrationSpec.
func (in *DevicePostureIntegrationSpec) DeepCopy() *DevicePostureIntegrationSpec {
	if in == nil {
		return nil
	}
	out := new(DevicePostureIntegrationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePostureIntegrationStatus) DeepCopyInto(out *DevicePostureIntegrationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePostureIntegrationStatus.
func (in *DevicePostureIntegrationStatus) DeepCopy() *DevicePostureIntegrationStatus {
	if in == nil {
		return nil
	}
	out := new(DevicePostureIntegrationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePostureRule) DeepCopyInto(out *DevicePostureRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePostureRule.
func (in *DevicePostureRule) DeepCopy() *DevicePostureRule {
	if in == nil {
		return nil
	}
	out := new(DevicePostureRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevicePostureRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePostureRuleList) DeepCopyInto(out *DevicePostureRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DevicePostureRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePostureRuleList.
func (in *DevicePostureRuleList) DeepCopy() *DevicePostureRuleList {
	if in == nil {
		return nil
	}
	out := new(DevicePostureRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevicePostureRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePostureRuleObservation) DeepCopyInto(out *DevicePostureRuleObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePostureRuleObservation.
func (in *DevicePostureRuleObservation) DeepCopy() *DevicePostureRuleObservation {
	if in == nil {
		return nil
	}
	out := new(DevicePostureRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePostureRuleParameters) DeepCopyInto(out *DevicePostureRuleParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expiration != nil {
		in, out := &in.Expiration, &out.Expiration
		*out = new(string)
		**out = **in
	}
	if in.Input != nil {
		in, out := &in.Input, &out.Input
		*out = make([]InputParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = make([]MatchParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePostureRuleParameters.
func (in *DevicePostureRuleParameters) DeepCopy() *DevicePostureRuleParameters {
	if in == nil {
		return nil
	}
	out := new(DevicePostureRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePostureRuleSpec) DeepCopyInto(out *DevicePostureRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePostureRuleSpec.
func (in *DevicePostureRuleSpec) DeepCopy() *DevicePostureRuleSpec {
	if in == nil {
		return nil
	}
	out := new(DevicePostureRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePostureRuleStatus) DeepCopyInto(out *DevicePostureRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePostureRuleStatus.
func (in *DevicePostureRuleStatus) DeepCopy() *DevicePostureRuleStatus {
	if in == nil {
		return nil
	}
	out := new(DevicePostureRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceSettingsPolicy) DeepCopyInto(out *DeviceSettingsPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceSettingsPolicy.
func (in *DeviceSettingsPolicy) DeepCopy() *DeviceSettingsPolicy {
	if in == nil {
		return nil
	}
	out := new(DeviceSettingsPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeviceSettingsPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceSettingsPolicyList) DeepCopyInto(out *DeviceSettingsPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeviceSettingsPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceSettingsPolicyList.
func (in *DeviceSettingsPolicyList) DeepCopy() *DeviceSettingsPolicyList {
	if in == nil {
		return nil
	}
	out := new(DeviceSettingsPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeviceSettingsPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceSettingsPolicyObservation) DeepCopyInto(out *DeviceSettingsPolicyObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceSettingsPolicyObservation.
func (in *DeviceSettingsPolicyObservation) DeepCopy() *DeviceSettingsPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(DeviceSettingsPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceSettingsPolicyParameters) DeepCopyInto(out *DeviceSettingsPolicyParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.AllowModeSwitch != nil {
		in, out := &in.AllowModeSwitch, &out.AllowModeSwitch
		*out = new(bool)
		**out = **in
	}
	if in.AllowUpdates != nil {
		in, out := &in.AllowUpdates, &out.AllowUpdates
		*out = new(bool)
		**out = **in
	}
	if in.AllowedToLeave != nil {
		in, out := &in.AllowedToLeave, &out.AllowedToLeave
		*out = new(bool)
		**out = **in
	}
	if in.AutoConnect != nil {
		in, out := &in.AutoConnect, &out.AutoConnect
		*out = new(float64)
		**out = **in
	}
	if in.CaptivePortal != nil {
		in, out := &in.CaptivePortal, &out.CaptivePortal
		*out = new(float64)
		**out = **in
	}
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(bool)
		**out = **in
	}
	if in.DisableAutoFallback != nil {
		in, out := &in.DisableAutoFallback, &out.DisableAutoFallback
		*out = new(bool)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Precedence != nil {
		in, out := &in.Precedence, &out.Precedence
		*out = new(float64)
		**out = **in
	}
	if in.ServiceModeV2Mode != nil {
		in, out := &in.ServiceModeV2Mode, &out.ServiceModeV2Mode
		*out = new(string)
		**out = **in
	}
	if in.ServiceModeV2Port != nil {
		in, out := &in.ServiceModeV2Port, &out.ServiceModeV2Port
		*out = new(float64)
		**out = **in
	}
	if in.SupportURL != nil {
		in, out := &in.SupportURL, &out.SupportURL
		*out = new(string)
		**out = **in
	}
	if in.SwitchLocked != nil {
		in, out := &in.SwitchLocked, &out.SwitchLocked
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceSettingsPolicyParameters.
func (in *DeviceSettingsPolicyParameters) DeepCopy() *DeviceSettingsPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(DeviceSettingsPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceSettingsPolicySpec) DeepCopyInto(out *DeviceSettingsPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceSettingsPolicySpec.
func (in *DeviceSettingsPolicySpec) DeepCopy() *DeviceSettingsPolicySpec {
	if in == nil {
		return nil
	}
	out := new(DeviceSettingsPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceSettingsPolicyStatus) DeepCopyInto(out *DeviceSettingsPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceSettingsPolicyStatus.
func (in *DeviceSettingsPolicyStatus) DeepCopy() *DeviceSettingsPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(DeviceSettingsPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InputObservation) DeepCopyInto(out *InputObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InputObservation.
func (in *InputObservation) DeepCopy() *InputObservation {
	if in == nil {
		return nil
	}
	out := new(InputObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InputParameters) DeepCopyInto(out *InputParameters) {
	*out = *in
	if in.ComplianceStatus != nil {
		in, out := &in.ComplianceStatus, &out.ComplianceStatus
		*out = new(string)
		**out = **in
	}
	if in.ConnectionID != nil {
		in, out := &in.ConnectionID, &out.ConnectionID
		*out = new(string)
		**out = **in
	}
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Exists != nil {
		in, out := &in.Exists, &out.Exists
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.OsDistroName != nil {
		in, out := &in.OsDistroName, &out.OsDistroName
		*out = new(string)
		**out = **in
	}
	if in.OsDistroRevision != nil {
		in, out := &in.OsDistroRevision, &out.OsDistroRevision
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.RequireAll != nil {
		in, out := &in.RequireAll, &out.RequireAll
		*out = new(bool)
		**out = **in
	}
	if in.Running != nil {
		in, out := &in.Running, &out.Running
		*out = new(bool)
		**out = **in
	}
	if in.Sha256 != nil {
		in, out := &in.Sha256, &out.Sha256
		*out = new(string)
		**out = **in
	}
	if in.Thumbprint != nil {
		in, out := &in.Thumbprint, &out.Thumbprint
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InputParameters.
func (in *InputParameters) DeepCopy() *InputParameters {
	if in == nil {
		return nil
	}
	out := new(InputParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchObservation) DeepCopyInto(out *MatchObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchObservation.
func (in *MatchObservation) DeepCopy() *MatchObservation {
	if in == nil {
		return nil
	}
	out := new(MatchObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchParameters) DeepCopyInto(out *MatchParameters) {
	*out = *in
	if in.Platform != nil {
		in, out := &in.Platform, &out.Platform
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchParameters.
func (in *MatchParameters) DeepCopy() *MatchParameters {
	if in == nil {
		return nil
	}
	out := new(MatchParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SplitTunnel) DeepCopyInto(out *SplitTunnel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SplitTunnel.
func (in *SplitTunnel) DeepCopy() *SplitTunnel {
	if in == nil {
		return nil
	}
	out := new(SplitTunnel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SplitTunnel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SplitTunnelList) DeepCopyInto(out *SplitTunnelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SplitTunnel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SplitTunnelList.
func (in *SplitTunnelList) DeepCopy() *SplitTunnelList {
	if in == nil {
		return nil
	}
	out := new(SplitTunnelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SplitTunnelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SplitTunnelObservation) DeepCopyInto(out *SplitTunnelObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SplitTunnelObservation.
func (in *SplitTunnelObservation) DeepCopy() *SplitTunnelObservation {
	if in == nil {
		return nil
	}
	out := new(SplitTunnelObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SplitTunnelParameters) DeepCopyInto(out *SplitTunnelParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.PolicyID != nil {
		in, out := &in.PolicyID, &out.PolicyID
		*out = new(string)
		**out = **in
	}
	if in.PolicyIDRef != nil {
		in, out := &in.PolicyIDRef, &out.PolicyIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.PolicyIDSelector != nil {
		in, out := &in.PolicyIDSelector, &out.PolicyIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Tunnels != nil {
		in, out := &in.Tunnels, &out.Tunnels
		*out = make([]TunnelsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SplitTunnelParameters.
func (in *SplitTunnelParameters) DeepCopy() *SplitTunnelParameters {
	if in == nil {
		return nil
	}
	out := new(SplitTunnelParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SplitTunnelSpec) DeepCopyInto(out *SplitTunnelSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SplitTunnelSpec.
func (in *SplitTunnelSpec) DeepCopy() *SplitTunnelSpec {
	if in == nil {
		return nil
	}
	out := new(SplitTunnelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SplitTunnelStatus) DeepCopyInto(out *SplitTunnelStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SplitTunnelStatus.
func (in *SplitTunnelStatus) DeepCopy() *SplitTunnelStatus {
	if in == nil {
		return nil
	}
	out := new(SplitTunnelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TunnelsObservation) DeepCopyInto(out *TunnelsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TunnelsObservation.
func (in *TunnelsObservation) DeepCopy() *TunnelsObservation {
	if in == nil {
		return nil
	}
	out := new(TunnelsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TunnelsParameters) DeepCopyInto(out *TunnelsParameters) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TunnelsParameters.
func (in *TunnelsParameters) DeepCopy() *TunnelsParameters {
	if in == nil {
		return nil
	}
	out := new(TunnelsParameters)
	in.DeepCopyInto(out)
	return out
}
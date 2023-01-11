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
func (in *AnalyticsEngineBindingObservation) DeepCopyInto(out *AnalyticsEngineBindingObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsEngineBindingObservation.
func (in *AnalyticsEngineBindingObservation) DeepCopy() *AnalyticsEngineBindingObservation {
	if in == nil {
		return nil
	}
	out := new(AnalyticsEngineBindingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsEngineBindingParameters) DeepCopyInto(out *AnalyticsEngineBindingParameters) {
	*out = *in
	if in.Dataset != nil {
		in, out := &in.Dataset, &out.Dataset
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsEngineBindingParameters.
func (in *AnalyticsEngineBindingParameters) DeepCopy() *AnalyticsEngineBindingParameters {
	if in == nil {
		return nil
	}
	out := new(AnalyticsEngineBindingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronTrigger) DeepCopyInto(out *CronTrigger) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronTrigger.
func (in *CronTrigger) DeepCopy() *CronTrigger {
	if in == nil {
		return nil
	}
	out := new(CronTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CronTrigger) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronTriggerList) DeepCopyInto(out *CronTriggerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CronTrigger, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronTriggerList.
func (in *CronTriggerList) DeepCopy() *CronTriggerList {
	if in == nil {
		return nil
	}
	out := new(CronTriggerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CronTriggerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronTriggerObservation) DeepCopyInto(out *CronTriggerObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronTriggerObservation.
func (in *CronTriggerObservation) DeepCopy() *CronTriggerObservation {
	if in == nil {
		return nil
	}
	out := new(CronTriggerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronTriggerParameters) DeepCopyInto(out *CronTriggerParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.AccountIDRef != nil {
		in, out := &in.AccountIDRef, &out.AccountIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.AccountIDSelector != nil {
		in, out := &in.AccountIDSelector, &out.AccountIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Schedules != nil {
		in, out := &in.Schedules, &out.Schedules
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ScriptName != nil {
		in, out := &in.ScriptName, &out.ScriptName
		*out = new(string)
		**out = **in
	}
	if in.ScriptNameRef != nil {
		in, out := &in.ScriptNameRef, &out.ScriptNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ScriptNameSelector != nil {
		in, out := &in.ScriptNameSelector, &out.ScriptNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronTriggerParameters.
func (in *CronTriggerParameters) DeepCopy() *CronTriggerParameters {
	if in == nil {
		return nil
	}
	out := new(CronTriggerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronTriggerSpec) DeepCopyInto(out *CronTriggerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronTriggerSpec.
func (in *CronTriggerSpec) DeepCopy() *CronTriggerSpec {
	if in == nil {
		return nil
	}
	out := new(CronTriggerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronTriggerStatus) DeepCopyInto(out *CronTriggerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronTriggerStatus.
func (in *CronTriggerStatus) DeepCopy() *CronTriggerStatus {
	if in == nil {
		return nil
	}
	out := new(CronTriggerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVNamespace) DeepCopyInto(out *KVNamespace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVNamespace.
func (in *KVNamespace) DeepCopy() *KVNamespace {
	if in == nil {
		return nil
	}
	out := new(KVNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KVNamespace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVNamespaceList) DeepCopyInto(out *KVNamespaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KVNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVNamespaceList.
func (in *KVNamespaceList) DeepCopy() *KVNamespaceList {
	if in == nil {
		return nil
	}
	out := new(KVNamespaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KVNamespaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVNamespaceObservation) DeepCopyInto(out *KVNamespaceObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVNamespaceObservation.
func (in *KVNamespaceObservation) DeepCopy() *KVNamespaceObservation {
	if in == nil {
		return nil
	}
	out := new(KVNamespaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVNamespaceParameters) DeepCopyInto(out *KVNamespaceParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.AccountIDRef != nil {
		in, out := &in.AccountIDRef, &out.AccountIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.AccountIDSelector != nil {
		in, out := &in.AccountIDSelector, &out.AccountIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVNamespaceParameters.
func (in *KVNamespaceParameters) DeepCopy() *KVNamespaceParameters {
	if in == nil {
		return nil
	}
	out := new(KVNamespaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVNamespaceSpec) DeepCopyInto(out *KVNamespaceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVNamespaceSpec.
func (in *KVNamespaceSpec) DeepCopy() *KVNamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(KVNamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVNamespaceStatus) DeepCopyInto(out *KVNamespaceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVNamespaceStatus.
func (in *KVNamespaceStatus) DeepCopy() *KVNamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(KVNamespaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kv) DeepCopyInto(out *Kv) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kv.
func (in *Kv) DeepCopy() *Kv {
	if in == nil {
		return nil
	}
	out := new(Kv)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Kv) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KvList) DeepCopyInto(out *KvList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Kv, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KvList.
func (in *KvList) DeepCopy() *KvList {
	if in == nil {
		return nil
	}
	out := new(KvList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KvList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KvNamespaceBindingObservation) DeepCopyInto(out *KvNamespaceBindingObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KvNamespaceBindingObservation.
func (in *KvNamespaceBindingObservation) DeepCopy() *KvNamespaceBindingObservation {
	if in == nil {
		return nil
	}
	out := new(KvNamespaceBindingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KvNamespaceBindingParameters) DeepCopyInto(out *KvNamespaceBindingParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamespaceID != nil {
		in, out := &in.NamespaceID, &out.NamespaceID
		*out = new(string)
		**out = **in
	}
	if in.NamespaceIDRef != nil {
		in, out := &in.NamespaceIDRef, &out.NamespaceIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceIDSelector != nil {
		in, out := &in.NamespaceIDSelector, &out.NamespaceIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KvNamespaceBindingParameters.
func (in *KvNamespaceBindingParameters) DeepCopy() *KvNamespaceBindingParameters {
	if in == nil {
		return nil
	}
	out := new(KvNamespaceBindingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KvObservation) DeepCopyInto(out *KvObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KvObservation.
func (in *KvObservation) DeepCopy() *KvObservation {
	if in == nil {
		return nil
	}
	out := new(KvObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KvParameters) DeepCopyInto(out *KvParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.AccountIDRef != nil {
		in, out := &in.AccountIDRef, &out.AccountIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.AccountIDSelector != nil {
		in, out := &in.AccountIDSelector, &out.AccountIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.NamespaceID != nil {
		in, out := &in.NamespaceID, &out.NamespaceID
		*out = new(string)
		**out = **in
	}
	if in.NamespaceIDRef != nil {
		in, out := &in.NamespaceIDRef, &out.NamespaceIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceIDSelector != nil {
		in, out := &in.NamespaceIDSelector, &out.NamespaceIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KvParameters.
func (in *KvParameters) DeepCopy() *KvParameters {
	if in == nil {
		return nil
	}
	out := new(KvParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KvSpec) DeepCopyInto(out *KvSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KvSpec.
func (in *KvSpec) DeepCopy() *KvSpec {
	if in == nil {
		return nil
	}
	out := new(KvSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KvStatus) DeepCopyInto(out *KvStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KvStatus.
func (in *KvStatus) DeepCopy() *KvStatus {
	if in == nil {
		return nil
	}
	out := new(KvStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlainTextBindingObservation) DeepCopyInto(out *PlainTextBindingObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlainTextBindingObservation.
func (in *PlainTextBindingObservation) DeepCopy() *PlainTextBindingObservation {
	if in == nil {
		return nil
	}
	out := new(PlainTextBindingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlainTextBindingParameters) DeepCopyInto(out *PlainTextBindingParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Text != nil {
		in, out := &in.Text, &out.Text
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlainTextBindingParameters.
func (in *PlainTextBindingParameters) DeepCopy() *PlainTextBindingParameters {
	if in == nil {
		return nil
	}
	out := new(PlainTextBindingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *R2BucketBindingObservation) DeepCopyInto(out *R2BucketBindingObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new R2BucketBindingObservation.
func (in *R2BucketBindingObservation) DeepCopy() *R2BucketBindingObservation {
	if in == nil {
		return nil
	}
	out := new(R2BucketBindingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *R2BucketBindingParameters) DeepCopyInto(out *R2BucketBindingParameters) {
	*out = *in
	if in.BucketName != nil {
		in, out := &in.BucketName, &out.BucketName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new R2BucketBindingParameters.
func (in *R2BucketBindingParameters) DeepCopy() *R2BucketBindingParameters {
	if in == nil {
		return nil
	}
	out := new(R2BucketBindingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route) DeepCopyInto(out *Route) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route.
func (in *Route) DeepCopy() *Route {
	if in == nil {
		return nil
	}
	out := new(Route)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Route) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteList) DeepCopyInto(out *RouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Route, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteList.
func (in *RouteList) DeepCopy() *RouteList {
	if in == nil {
		return nil
	}
	out := new(RouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteObservation) DeepCopyInto(out *RouteObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteObservation.
func (in *RouteObservation) DeepCopy() *RouteObservation {
	if in == nil {
		return nil
	}
	out := new(RouteObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteParameters) DeepCopyInto(out *RouteParameters) {
	*out = *in
	if in.Pattern != nil {
		in, out := &in.Pattern, &out.Pattern
		*out = new(string)
		**out = **in
	}
	if in.ScriptName != nil {
		in, out := &in.ScriptName, &out.ScriptName
		*out = new(string)
		**out = **in
	}
	if in.ScriptNameRef != nil {
		in, out := &in.ScriptNameRef, &out.ScriptNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ScriptNameSelector != nil {
		in, out := &in.ScriptNameSelector, &out.ScriptNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
	if in.ZoneIDRef != nil {
		in, out := &in.ZoneIDRef, &out.ZoneIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ZoneIDSelector != nil {
		in, out := &in.ZoneIDSelector, &out.ZoneIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteParameters.
func (in *RouteParameters) DeepCopy() *RouteParameters {
	if in == nil {
		return nil
	}
	out := new(RouteParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteSpec) DeepCopyInto(out *RouteSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteSpec.
func (in *RouteSpec) DeepCopy() *RouteSpec {
	if in == nil {
		return nil
	}
	out := new(RouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteStatus) DeepCopyInto(out *RouteStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteStatus.
func (in *RouteStatus) DeepCopy() *RouteStatus {
	if in == nil {
		return nil
	}
	out := new(RouteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Script) DeepCopyInto(out *Script) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Script.
func (in *Script) DeepCopy() *Script {
	if in == nil {
		return nil
	}
	out := new(Script)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Script) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScriptList) DeepCopyInto(out *ScriptList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Script, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScriptList.
func (in *ScriptList) DeepCopy() *ScriptList {
	if in == nil {
		return nil
	}
	out := new(ScriptList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScriptList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScriptObservation) DeepCopyInto(out *ScriptObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScriptObservation.
func (in *ScriptObservation) DeepCopy() *ScriptObservation {
	if in == nil {
		return nil
	}
	out := new(ScriptObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScriptParameters) DeepCopyInto(out *ScriptParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.AnalyticsEngineBinding != nil {
		in, out := &in.AnalyticsEngineBinding, &out.AnalyticsEngineBinding
		*out = make([]AnalyticsEngineBindingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.KvNamespaceBinding != nil {
		in, out := &in.KvNamespaceBinding, &out.KvNamespaceBinding
		*out = make([]KvNamespaceBindingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Module != nil {
		in, out := &in.Module, &out.Module
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PlainTextBinding != nil {
		in, out := &in.PlainTextBinding, &out.PlainTextBinding
		*out = make([]PlainTextBindingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.R2BucketBinding != nil {
		in, out := &in.R2BucketBinding, &out.R2BucketBinding
		*out = make([]R2BucketBindingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecretTextBinding != nil {
		in, out := &in.SecretTextBinding, &out.SecretTextBinding
		*out = make([]SecretTextBindingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceBinding != nil {
		in, out := &in.ServiceBinding, &out.ServiceBinding
		*out = make([]ServiceBindingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WebassemblyBinding != nil {
		in, out := &in.WebassemblyBinding, &out.WebassemblyBinding
		*out = make([]WebassemblyBindingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScriptParameters.
func (in *ScriptParameters) DeepCopy() *ScriptParameters {
	if in == nil {
		return nil
	}
	out := new(ScriptParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScriptSpec) DeepCopyInto(out *ScriptSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScriptSpec.
func (in *ScriptSpec) DeepCopy() *ScriptSpec {
	if in == nil {
		return nil
	}
	out := new(ScriptSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScriptStatus) DeepCopyInto(out *ScriptStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScriptStatus.
func (in *ScriptStatus) DeepCopy() *ScriptStatus {
	if in == nil {
		return nil
	}
	out := new(ScriptStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretTextBindingObservation) DeepCopyInto(out *SecretTextBindingObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretTextBindingObservation.
func (in *SecretTextBindingObservation) DeepCopy() *SecretTextBindingObservation {
	if in == nil {
		return nil
	}
	out := new(SecretTextBindingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretTextBindingParameters) DeepCopyInto(out *SecretTextBindingParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	out.TextSecretRef = in.TextSecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretTextBindingParameters.
func (in *SecretTextBindingParameters) DeepCopy() *SecretTextBindingParameters {
	if in == nil {
		return nil
	}
	out := new(SecretTextBindingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingObservation) DeepCopyInto(out *ServiceBindingObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingObservation.
func (in *ServiceBindingObservation) DeepCopy() *ServiceBindingObservation {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingParameters) DeepCopyInto(out *ServiceBindingParameters) {
	*out = *in
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingParameters.
func (in *ServiceBindingParameters) DeepCopy() *ServiceBindingParameters {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebassemblyBindingObservation) DeepCopyInto(out *WebassemblyBindingObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebassemblyBindingObservation.
func (in *WebassemblyBindingObservation) DeepCopy() *WebassemblyBindingObservation {
	if in == nil {
		return nil
	}
	out := new(WebassemblyBindingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebassemblyBindingParameters) DeepCopyInto(out *WebassemblyBindingParameters) {
	*out = *in
	if in.Module != nil {
		in, out := &in.Module, &out.Module
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebassemblyBindingParameters.
func (in *WebassemblyBindingParameters) DeepCopy() *WebassemblyBindingParameters {
	if in == nil {
		return nil
	}
	out := new(WebassemblyBindingParameters)
	in.DeepCopyInto(out)
	return out
}

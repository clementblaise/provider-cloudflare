/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/clementblaise/provider-cloudflare/apis/account/v1alpha1"
	v1alpha11 "github.com/clementblaise/provider-cloudflare/apis/zone/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this CronTrigger.
func (mg *CronTrigger) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AccountIDRef,
		Selector:     mg.Spec.ForProvider.AccountIDSelector,
		To: reference.To{
			List:    &v1alpha1.AccountList{},
			Managed: &v1alpha1.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AccountID")
	}
	mg.Spec.ForProvider.AccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccountIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ScriptName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ScriptNameRef,
		Selector:     mg.Spec.ForProvider.ScriptNameSelector,
		To: reference.To{
			List:    &ScriptList{},
			Managed: &Script{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ScriptName")
	}
	mg.Spec.ForProvider.ScriptName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ScriptNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Domain.
func (mg *Domain) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AccountIDRef,
		Selector:     mg.Spec.ForProvider.AccountIDSelector,
		To: reference.To{
			List:    &v1alpha1.AccountList{},
			Managed: &v1alpha1.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AccountID")
	}
	mg.Spec.ForProvider.AccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccountIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &v1alpha11.ZoneList{},
			Managed: &v1alpha11.Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this KV.
func (mg *KV) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AccountIDRef,
		Selector:     mg.Spec.ForProvider.AccountIDSelector,
		To: reference.To{
			List:    &v1alpha1.AccountList{},
			Managed: &v1alpha1.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AccountID")
	}
	mg.Spec.ForProvider.AccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccountIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NamespaceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NamespaceIDRef,
		Selector:     mg.Spec.ForProvider.NamespaceIDSelector,
		To: reference.To{
			List:    &KVNamespaceList{},
			Managed: &KVNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NamespaceID")
	}
	mg.Spec.ForProvider.NamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this KVNamespace.
func (mg *KVNamespace) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AccountIDRef,
		Selector:     mg.Spec.ForProvider.AccountIDSelector,
		To: reference.To{
			List:    &v1alpha1.AccountList{},
			Managed: &v1alpha1.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AccountID")
	}
	mg.Spec.ForProvider.AccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccountIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Route.
func (mg *Route) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ScriptName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ScriptNameRef,
		Selector:     mg.Spec.ForProvider.ScriptNameSelector,
		To: reference.To{
			List:    &ScriptList{},
			Managed: &Script{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ScriptName")
	}
	mg.Spec.ForProvider.ScriptName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ScriptNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &v1alpha11.ZoneList{},
			Managed: &v1alpha11.Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Script.
func (mg *Script) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.KvNamespaceBinding); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KvNamespaceBinding[i3].NamespaceID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.KvNamespaceBinding[i3].NamespaceIDRef,
			Selector:     mg.Spec.ForProvider.KvNamespaceBinding[i3].NamespaceIDSelector,
			To: reference.To{
				List:    &KVNamespaceList{},
				Managed: &KVNamespace{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.KvNamespaceBinding[i3].NamespaceID")
		}
		mg.Spec.ForProvider.KvNamespaceBinding[i3].NamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.KvNamespaceBinding[i3].NamespaceIDRef = rsp.ResolvedReference

	}

	return nil
}

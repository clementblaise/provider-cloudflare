/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/cdloh/provider-cloudflare/apis/zone/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this LoadBalancer.
func (mg *LoadBalancer) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.CountryPools); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.CountryPools[i3].PoolIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.CountryPools[i3].PoolIdsRefs,
			Selector:      mg.Spec.ForProvider.CountryPools[i3].PoolIdsSelector,
			To: reference.To{
				List:    &PoolList{},
				Managed: &Pool{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CountryPools[i3].PoolIds")
		}
		mg.Spec.ForProvider.CountryPools[i3].PoolIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.CountryPools[i3].PoolIdsRefs = mrsp.ResolvedReferences

	}
	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.DefaultPoolIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.DefaultPoolIdsRefs,
		Selector:      mg.Spec.ForProvider.DefaultPoolIdsSelector,
		To: reference.To{
			List:    &PoolList{},
			Managed: &Pool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DefaultPoolIds")
	}
	mg.Spec.ForProvider.DefaultPoolIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.DefaultPoolIdsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FallbackPoolID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FallbackPoolIDRef,
		Selector:     mg.Spec.ForProvider.FallbackPoolIDSelector,
		To: reference.To{
			List:    &PoolList{},
			Managed: &Pool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FallbackPoolID")
	}
	mg.Spec.ForProvider.FallbackPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FallbackPoolIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.PopPools); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.PopPools[i3].PoolIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.PopPools[i3].PoolIdsRefs,
			Selector:      mg.Spec.ForProvider.PopPools[i3].PoolIdsSelector,
			To: reference.To{
				List:    &PoolList{},
				Managed: &Pool{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.PopPools[i3].PoolIds")
		}
		mg.Spec.ForProvider.PopPools[i3].PoolIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.PopPools[i3].PoolIdsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.RegionPools); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.RegionPools[i3].PoolIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.RegionPools[i3].PoolIdsRefs,
			Selector:      mg.Spec.ForProvider.RegionPools[i3].PoolIdsSelector,
			To: reference.To{
				List:    &PoolList{},
				Managed: &Pool{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.RegionPools[i3].PoolIds")
		}
		mg.Spec.ForProvider.RegionPools[i3].PoolIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.RegionPools[i3].PoolIdsRefs = mrsp.ResolvedReferences

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &v1alpha1.ZoneList{},
			Managed: &v1alpha1.Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Pool.
func (mg *Pool) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Monitor),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.MonitorRef,
		Selector:     mg.Spec.ForProvider.MonitorSelector,
		To: reference.To{
			List:    &MonitorList{},
			Managed: &Monitor{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Monitor")
	}
	mg.Spec.ForProvider.Monitor = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MonitorRef = rsp.ResolvedReference

	return nil
}
/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Profile.
func (mg *Profile) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.RoleArns),
		Extract:       common.ARNExtractor(),
		References:    mg.Spec.ForProvider.RoleArnsRefs,
		Selector:      mg.Spec.ForProvider.RoleArnsSelector,
		To: reference.To{
			List:    &v1beta1.RoleList{},
			Managed: &v1beta1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArns")
	}
	mg.Spec.ForProvider.RoleArns = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.RoleArnsRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.RoleArns),
		Extract:       common.ARNExtractor(),
		References:    mg.Spec.InitProvider.RoleArnsRefs,
		Selector:      mg.Spec.InitProvider.RoleArnsSelector,
		To: reference.To{
			List:    &v1beta1.RoleList{},
			Managed: &v1beta1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RoleArns")
	}
	mg.Spec.InitProvider.RoleArns = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.RoleArnsRefs = mrsp.ResolvedReferences

	return nil
}

/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1beta11 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	v1beta1 "github.com/upbound/provider-aws/apis/kms/v1beta1"
	v1beta12 "github.com/upbound/provider-aws/apis/s3/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this EventDataStore.
func (mg *EventDataStore) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyIDRef,
		Selector:     mg.Spec.ForProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1beta1.KeyList{},
			Managed: &v1beta1.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyID")
	}
	mg.Spec.ForProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.KMSKeyIDRef,
		Selector:     mg.Spec.InitProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1beta1.KeyList{},
			Managed: &v1beta1.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KMSKeyID")
	}
	mg.Spec.InitProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KMSKeyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Trail.
func (mg *Trail) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CloudWatchLogsRoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.CloudWatchLogsRoleArnRef,
		Selector:     mg.Spec.ForProvider.CloudWatchLogsRoleArnSelector,
		To: reference.To{
			List:    &v1beta11.RoleList{},
			Managed: &v1beta11.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CloudWatchLogsRoleArn")
	}
	mg.Spec.ForProvider.CloudWatchLogsRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CloudWatchLogsRoleArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyIDRef,
		Selector:     mg.Spec.ForProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1beta1.KeyList{},
			Managed: &v1beta1.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyID")
	}
	mg.Spec.ForProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.S3BucketName),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.S3BucketNameRef,
		Selector:     mg.Spec.ForProvider.S3BucketNameSelector,
		To: reference.To{
			List:    &v1beta12.BucketList{},
			Managed: &v1beta12.Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.S3BucketName")
	}
	mg.Spec.ForProvider.S3BucketName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.S3BucketNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CloudWatchLogsRoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.InitProvider.CloudWatchLogsRoleArnRef,
		Selector:     mg.Spec.InitProvider.CloudWatchLogsRoleArnSelector,
		To: reference.To{
			List:    &v1beta11.RoleList{},
			Managed: &v1beta11.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CloudWatchLogsRoleArn")
	}
	mg.Spec.InitProvider.CloudWatchLogsRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CloudWatchLogsRoleArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.KMSKeyIDRef,
		Selector:     mg.Spec.InitProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1beta1.KeyList{},
			Managed: &v1beta1.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KMSKeyID")
	}
	mg.Spec.InitProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KMSKeyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.S3BucketName),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.S3BucketNameRef,
		Selector:     mg.Spec.InitProvider.S3BucketNameSelector,
		To: reference.To{
			List:    &v1beta12.BucketList{},
			Managed: &v1beta12.Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.S3BucketName")
	}
	mg.Spec.InitProvider.S3BucketName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.S3BucketNameRef = rsp.ResolvedReference

	return nil
}

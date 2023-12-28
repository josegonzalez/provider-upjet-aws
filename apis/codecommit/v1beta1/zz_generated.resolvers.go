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
	v1beta1 "github.com/upbound/provider-aws/apis/sns/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ApprovalRuleTemplateAssociation.
func (mg *ApprovalRuleTemplateAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApprovalRuleTemplateName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ApprovalRuleTemplateNameRef,
		Selector:     mg.Spec.ForProvider.ApprovalRuleTemplateNameSelector,
		To: reference.To{
			List:    &ApprovalRuleTemplateList{},
			Managed: &ApprovalRuleTemplate{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApprovalRuleTemplateName")
	}
	mg.Spec.ForProvider.ApprovalRuleTemplateName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApprovalRuleTemplateNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RepositoryName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryNameRef,
		Selector:     mg.Spec.ForProvider.RepositoryNameSelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RepositoryName")
	}
	mg.Spec.ForProvider.RepositoryName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Trigger.
func (mg *Trigger) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RepositoryName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryNameRef,
		Selector:     mg.Spec.ForProvider.RepositoryNameSelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RepositoryName")
	}
	mg.Spec.ForProvider.RepositoryName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Trigger); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Trigger[i3].DestinationArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.Trigger[i3].DestinationArnRef,
			Selector:     mg.Spec.ForProvider.Trigger[i3].DestinationArnSelector,
			To: reference.To{
				List:    &v1beta1.TopicList{},
				Managed: &v1beta1.Topic{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Trigger[i3].DestinationArn")
		}
		mg.Spec.ForProvider.Trigger[i3].DestinationArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Trigger[i3].DestinationArnRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RepositoryName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RepositoryNameRef,
		Selector:     mg.Spec.InitProvider.RepositoryNameSelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RepositoryName")
	}
	mg.Spec.InitProvider.RepositoryName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Trigger); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Trigger[i3].DestinationArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.Trigger[i3].DestinationArnRef,
			Selector:     mg.Spec.InitProvider.Trigger[i3].DestinationArnSelector,
			To: reference.To{
				List:    &v1beta1.TopicList{},
				Managed: &v1beta1.Topic{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Trigger[i3].DestinationArn")
		}
		mg.Spec.InitProvider.Trigger[i3].DestinationArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Trigger[i3].DestinationArnRef = rsp.ResolvedReference

	}

	return nil
}

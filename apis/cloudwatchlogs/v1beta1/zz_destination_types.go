// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DestinationInitParameters struct {

	// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to put data into the target.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ARN of the target Amazon Kinesis stream resource for the destination.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kinesis/v1beta1.Stream
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.TerraformID()
	TargetArn *string `json:"targetArn,omitempty" tf:"target_arn,omitempty"`

	// Reference to a Stream in kinesis to populate targetArn.
	// +kubebuilder:validation:Optional
	TargetArnRef *v1.Reference `json:"targetArnRef,omitempty" tf:"-"`

	// Selector for a Stream in kinesis to populate targetArn.
	// +kubebuilder:validation:Optional
	TargetArnSelector *v1.Selector `json:"targetArnSelector,omitempty" tf:"-"`
}

type DestinationObservation struct {

	// The Amazon Resource Name (ARN) specifying the log destination.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to put data into the target.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The ARN of the target Amazon Kinesis stream resource for the destination.
	TargetArn *string `json:"targetArn,omitempty" tf:"target_arn,omitempty"`
}

type DestinationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to put data into the target.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ARN of the target Amazon Kinesis stream resource for the destination.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kinesis/v1beta1.Stream
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.TerraformID()
	// +kubebuilder:validation:Optional
	TargetArn *string `json:"targetArn,omitempty" tf:"target_arn,omitempty"`

	// Reference to a Stream in kinesis to populate targetArn.
	// +kubebuilder:validation:Optional
	TargetArnRef *v1.Reference `json:"targetArnRef,omitempty" tf:"-"`

	// Selector for a Stream in kinesis to populate targetArn.
	// +kubebuilder:validation:Optional
	TargetArnSelector *v1.Selector `json:"targetArnSelector,omitempty" tf:"-"`
}

// DestinationSpec defines the desired state of Destination
type DestinationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DestinationParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider DestinationInitParameters `json:"initProvider,omitempty"`
}

// DestinationStatus defines the observed state of Destination.
type DestinationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DestinationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Destination is the Schema for the Destinations API. Provides a CloudWatch Logs destination.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Destination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DestinationSpec   `json:"spec"`
	Status            DestinationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DestinationList contains a list of Destinations
type DestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Destination `json:"items"`
}

// Repository type metadata.
var (
	Destination_Kind             = "Destination"
	Destination_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Destination_Kind}.String()
	Destination_KindAPIVersion   = Destination_Kind + "." + CRDGroupVersion.String()
	Destination_GroupVersionKind = CRDGroupVersion.WithKind(Destination_Kind)
)

func init() {
	SchemeBuilder.Register(&Destination{}, &DestinationList{})
}

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

type QueueRedrivePolicyInitParameters struct {

	// The URL of the SQS Queue to which to attach the policy
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sqs/v1beta1.Queue
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.TerraformID()
	QueueURL *string `json:"queueUrl,omitempty" tf:"queue_url,omitempty"`

	// Reference to a Queue in sqs to populate queueUrl.
	// +kubebuilder:validation:Optional
	QueueURLRef *v1.Reference `json:"queueUrlRef,omitempty" tf:"-"`

	// Selector for a Queue in sqs to populate queueUrl.
	// +kubebuilder:validation:Optional
	QueueURLSelector *v1.Selector `json:"queueUrlSelector,omitempty" tf:"-"`

	// The JSON redrive policy for the SQS queue. Accepts two key/val pairs: deadLetterTargetArn and maxReceiveCount. Learn more in the Amazon SQS dead-letter queues documentation.
	RedrivePolicy *string `json:"redrivePolicy,omitempty" tf:"redrive_policy,omitempty"`
}

type QueueRedrivePolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The URL of the SQS Queue to which to attach the policy
	QueueURL *string `json:"queueUrl,omitempty" tf:"queue_url,omitempty"`

	// The JSON redrive policy for the SQS queue. Accepts two key/val pairs: deadLetterTargetArn and maxReceiveCount. Learn more in the Amazon SQS dead-letter queues documentation.
	RedrivePolicy *string `json:"redrivePolicy,omitempty" tf:"redrive_policy,omitempty"`
}

type QueueRedrivePolicyParameters struct {

	// The URL of the SQS Queue to which to attach the policy
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sqs/v1beta1.Queue
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.TerraformID()
	// +kubebuilder:validation:Optional
	QueueURL *string `json:"queueUrl,omitempty" tf:"queue_url,omitempty"`

	// Reference to a Queue in sqs to populate queueUrl.
	// +kubebuilder:validation:Optional
	QueueURLRef *v1.Reference `json:"queueUrlRef,omitempty" tf:"-"`

	// Selector for a Queue in sqs to populate queueUrl.
	// +kubebuilder:validation:Optional
	QueueURLSelector *v1.Selector `json:"queueUrlSelector,omitempty" tf:"-"`

	// The JSON redrive policy for the SQS queue. Accepts two key/val pairs: deadLetterTargetArn and maxReceiveCount. Learn more in the Amazon SQS dead-letter queues documentation.
	// +kubebuilder:validation:Optional
	RedrivePolicy *string `json:"redrivePolicy,omitempty" tf:"redrive_policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// QueueRedrivePolicySpec defines the desired state of QueueRedrivePolicy
type QueueRedrivePolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QueueRedrivePolicyParameters `json:"forProvider"`
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
	InitProvider QueueRedrivePolicyInitParameters `json:"initProvider,omitempty"`
}

// QueueRedrivePolicyStatus defines the observed state of QueueRedrivePolicy.
type QueueRedrivePolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QueueRedrivePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// QueueRedrivePolicy is the Schema for the QueueRedrivePolicys API. Provides a SQS Queue Redrive Policy resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type QueueRedrivePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.redrivePolicy) || (has(self.initProvider) && has(self.initProvider.redrivePolicy))",message="spec.forProvider.redrivePolicy is a required parameter"
	Spec   QueueRedrivePolicySpec   `json:"spec"`
	Status QueueRedrivePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QueueRedrivePolicyList contains a list of QueueRedrivePolicys
type QueueRedrivePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QueueRedrivePolicy `json:"items"`
}

// Repository type metadata.
var (
	QueueRedrivePolicy_Kind             = "QueueRedrivePolicy"
	QueueRedrivePolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: QueueRedrivePolicy_Kind}.String()
	QueueRedrivePolicy_KindAPIVersion   = QueueRedrivePolicy_Kind + "." + CRDGroupVersion.String()
	QueueRedrivePolicy_GroupVersionKind = CRDGroupVersion.WithKind(QueueRedrivePolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&QueueRedrivePolicy{}, &QueueRedrivePolicyList{})
}

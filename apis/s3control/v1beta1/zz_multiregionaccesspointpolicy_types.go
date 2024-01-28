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

type MultiRegionAccessPointPolicyDetailsInitParameters struct {

	// The name of the Multi-Region Access Point.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A valid JSON document that specifies the policy that you want to associate with this Multi-Region Access Point. Once applied, the policy can be edited, but not deleted. For more information, see the documentation on Multi-Region Access Point Permissions.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type MultiRegionAccessPointPolicyDetailsObservation struct {

	// The name of the Multi-Region Access Point.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A valid JSON document that specifies the policy that you want to associate with this Multi-Region Access Point. Once applied, the policy can be edited, but not deleted. For more information, see the documentation on Multi-Region Access Point Permissions.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type MultiRegionAccessPointPolicyDetailsParameters struct {

	// The name of the Multi-Region Access Point.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// A valid JSON document that specifies the policy that you want to associate with this Multi-Region Access Point. Once applied, the policy can be edited, but not deleted. For more information, see the documentation on Multi-Region Access Point Permissions.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy" tf:"policy,omitempty"`
}

type MultiRegionAccessPointPolicyInitParameters struct {

	// The AWS account ID for the owner of the Multi-Region Access Point.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// A configuration block containing details about the policy for the Multi-Region Access Point. See Details Configuration Block below for more details
	Details []MultiRegionAccessPointPolicyDetailsInitParameters `json:"details,omitempty" tf:"details,omitempty"`
}

type MultiRegionAccessPointPolicyObservation struct {

	// The AWS account ID for the owner of the Multi-Region Access Point.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// A configuration block containing details about the policy for the Multi-Region Access Point. See Details Configuration Block below for more details
	Details []MultiRegionAccessPointPolicyDetailsObservation `json:"details,omitempty" tf:"details,omitempty"`

	// The last established policy for the Multi-Region Access Point.
	Established *string `json:"established,omitempty" tf:"established,omitempty"`

	// The AWS account ID and access point name separated by a colon (:).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The proposed policy for the Multi-Region Access Point.
	Proposed *string `json:"proposed,omitempty" tf:"proposed,omitempty"`
}

type MultiRegionAccessPointPolicyParameters struct {

	// The AWS account ID for the owner of the Multi-Region Access Point.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// A configuration block containing details about the policy for the Multi-Region Access Point. See Details Configuration Block below for more details
	// +kubebuilder:validation:Optional
	Details []MultiRegionAccessPointPolicyDetailsParameters `json:"details,omitempty" tf:"details,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// MultiRegionAccessPointPolicySpec defines the desired state of MultiRegionAccessPointPolicy
type MultiRegionAccessPointPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MultiRegionAccessPointPolicyParameters `json:"forProvider"`
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
	InitProvider MultiRegionAccessPointPolicyInitParameters `json:"initProvider,omitempty"`
}

// MultiRegionAccessPointPolicyStatus defines the observed state of MultiRegionAccessPointPolicy.
type MultiRegionAccessPointPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MultiRegionAccessPointPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// MultiRegionAccessPointPolicy is the Schema for the MultiRegionAccessPointPolicys API. Provides a resource to manage an S3 Multi-Region Access Point access control policy.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type MultiRegionAccessPointPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.details) || (has(self.initProvider) && has(self.initProvider.details))",message="spec.forProvider.details is a required parameter"
	Spec   MultiRegionAccessPointPolicySpec   `json:"spec"`
	Status MultiRegionAccessPointPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MultiRegionAccessPointPolicyList contains a list of MultiRegionAccessPointPolicys
type MultiRegionAccessPointPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MultiRegionAccessPointPolicy `json:"items"`
}

// Repository type metadata.
var (
	MultiRegionAccessPointPolicy_Kind             = "MultiRegionAccessPointPolicy"
	MultiRegionAccessPointPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MultiRegionAccessPointPolicy_Kind}.String()
	MultiRegionAccessPointPolicy_KindAPIVersion   = MultiRegionAccessPointPolicy_Kind + "." + CRDGroupVersion.String()
	MultiRegionAccessPointPolicy_GroupVersionKind = CRDGroupVersion.WithKind(MultiRegionAccessPointPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&MultiRegionAccessPointPolicy{}, &MultiRegionAccessPointPolicyList{})
}

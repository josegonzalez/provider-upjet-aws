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

type GroupMembershipInitParameters struct {

	// The identifier for a group in the Identity Store.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/identitystore/v1beta1.Group
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("group_id",true)
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Reference to a Group in identitystore to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDRef *v1.Reference `json:"groupIdRef,omitempty" tf:"-"`

	// Selector for a Group in identitystore to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDSelector *v1.Selector `json:"groupIdSelector,omitempty" tf:"-"`

	// The identifier for a user in the Identity Store.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/identitystore/v1beta1.User
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("user_id",true)
	MemberID *string `json:"memberId,omitempty" tf:"member_id,omitempty"`

	// Reference to a User in identitystore to populate memberId.
	// +kubebuilder:validation:Optional
	MemberIDRef *v1.Reference `json:"memberIdRef,omitempty" tf:"-"`

	// Selector for a User in identitystore to populate memberId.
	// +kubebuilder:validation:Optional
	MemberIDSelector *v1.Selector `json:"memberIdSelector,omitempty" tf:"-"`
}

type GroupMembershipObservation struct {

	// The identifier for a group in the Identity Store.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Identity Store ID associated with the Single Sign-On Instance.
	IdentityStoreID *string `json:"identityStoreId,omitempty" tf:"identity_store_id,omitempty"`

	// The identifier for a user in the Identity Store.
	MemberID *string `json:"memberId,omitempty" tf:"member_id,omitempty"`

	// The identifier of the newly created group membership in the Identity Store.
	MembershipID *string `json:"membershipId,omitempty" tf:"membership_id,omitempty"`
}

type GroupMembershipParameters struct {

	// The identifier for a group in the Identity Store.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/identitystore/v1beta1.Group
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("group_id",true)
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Reference to a Group in identitystore to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDRef *v1.Reference `json:"groupIdRef,omitempty" tf:"-"`

	// Selector for a Group in identitystore to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDSelector *v1.Selector `json:"groupIdSelector,omitempty" tf:"-"`

	// Identity Store ID associated with the Single Sign-On Instance.
	// +kubebuilder:validation:Required
	IdentityStoreID *string `json:"identityStoreId" tf:"identity_store_id,omitempty"`

	// The identifier for a user in the Identity Store.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/identitystore/v1beta1.User
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("user_id",true)
	// +kubebuilder:validation:Optional
	MemberID *string `json:"memberId,omitempty" tf:"member_id,omitempty"`

	// Reference to a User in identitystore to populate memberId.
	// +kubebuilder:validation:Optional
	MemberIDRef *v1.Reference `json:"memberIdRef,omitempty" tf:"-"`

	// Selector for a User in identitystore to populate memberId.
	// +kubebuilder:validation:Optional
	MemberIDSelector *v1.Selector `json:"memberIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// GroupMembershipSpec defines the desired state of GroupMembership
type GroupMembershipSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupMembershipParameters `json:"forProvider"`
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
	InitProvider GroupMembershipInitParameters `json:"initProvider,omitempty"`
}

// GroupMembershipStatus defines the observed state of GroupMembership.
type GroupMembershipStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupMembershipObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// GroupMembership is the Schema for the GroupMemberships API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type GroupMembership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupMembershipSpec   `json:"spec"`
	Status            GroupMembershipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupMembershipList contains a list of GroupMemberships
type GroupMembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupMembership `json:"items"`
}

// Repository type metadata.
var (
	GroupMembership_Kind             = "GroupMembership"
	GroupMembership_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GroupMembership_Kind}.String()
	GroupMembership_KindAPIVersion   = GroupMembership_Kind + "." + CRDGroupVersion.String()
	GroupMembership_GroupVersionKind = CRDGroupVersion.WithKind(GroupMembership_Kind)
)

func init() {
	SchemeBuilder.Register(&GroupMembership{}, &GroupMembershipList{})
}

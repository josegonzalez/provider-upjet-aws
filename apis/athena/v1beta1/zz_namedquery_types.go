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

type NamedQueryInitParameters struct {

	// Database to which the query belongs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/athena/v1beta1.Database
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// Reference to a Database in athena to populate database.
	// +kubebuilder:validation:Optional
	DatabaseRef *v1.Reference `json:"databaseRef,omitempty" tf:"-"`

	// Selector for a Database in athena to populate database.
	// +kubebuilder:validation:Optional
	DatabaseSelector *v1.Selector `json:"databaseSelector,omitempty" tf:"-"`

	// Brief explanation of the query. Maximum length of 1024.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Plain language name for the query. Maximum length of 128.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Text of the query itself. In other words, all query statements. Maximum length of 262144.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// Workgroup to which the query belongs. Defaults to primary
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/athena/v1beta1.Workgroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Workgroup *string `json:"workgroup,omitempty" tf:"workgroup,omitempty"`

	// Reference to a Workgroup in athena to populate workgroup.
	// +kubebuilder:validation:Optional
	WorkgroupRef *v1.Reference `json:"workgroupRef,omitempty" tf:"-"`

	// Selector for a Workgroup in athena to populate workgroup.
	// +kubebuilder:validation:Optional
	WorkgroupSelector *v1.Selector `json:"workgroupSelector,omitempty" tf:"-"`
}

type NamedQueryObservation struct {

	// Database to which the query belongs.
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// Brief explanation of the query. Maximum length of 1024.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Unique ID of the query.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Plain language name for the query. Maximum length of 128.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Text of the query itself. In other words, all query statements. Maximum length of 262144.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// Workgroup to which the query belongs. Defaults to primary
	Workgroup *string `json:"workgroup,omitempty" tf:"workgroup,omitempty"`
}

type NamedQueryParameters struct {

	// Database to which the query belongs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/athena/v1beta1.Database
	// +kubebuilder:validation:Optional
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// Reference to a Database in athena to populate database.
	// +kubebuilder:validation:Optional
	DatabaseRef *v1.Reference `json:"databaseRef,omitempty" tf:"-"`

	// Selector for a Database in athena to populate database.
	// +kubebuilder:validation:Optional
	DatabaseSelector *v1.Selector `json:"databaseSelector,omitempty" tf:"-"`

	// Brief explanation of the query. Maximum length of 1024.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Plain language name for the query. Maximum length of 128.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Text of the query itself. In other words, all query statements. Maximum length of 262144.
	// +kubebuilder:validation:Optional
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Workgroup to which the query belongs. Defaults to primary
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/athena/v1beta1.Workgroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Workgroup *string `json:"workgroup,omitempty" tf:"workgroup,omitempty"`

	// Reference to a Workgroup in athena to populate workgroup.
	// +kubebuilder:validation:Optional
	WorkgroupRef *v1.Reference `json:"workgroupRef,omitempty" tf:"-"`

	// Selector for a Workgroup in athena to populate workgroup.
	// +kubebuilder:validation:Optional
	WorkgroupSelector *v1.Selector `json:"workgroupSelector,omitempty" tf:"-"`
}

// NamedQuerySpec defines the desired state of NamedQuery
type NamedQuerySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NamedQueryParameters `json:"forProvider"`
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
	InitProvider NamedQueryInitParameters `json:"initProvider,omitempty"`
}

// NamedQueryStatus defines the observed state of NamedQuery.
type NamedQueryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NamedQueryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NamedQuery is the Schema for the NamedQuerys API. Provides an Athena Named Query resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type NamedQuery struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.query) || (has(self.initProvider) && has(self.initProvider.query))",message="spec.forProvider.query is a required parameter"
	Spec   NamedQuerySpec   `json:"spec"`
	Status NamedQueryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NamedQueryList contains a list of NamedQuerys
type NamedQueryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamedQuery `json:"items"`
}

// Repository type metadata.
var (
	NamedQuery_Kind             = "NamedQuery"
	NamedQuery_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NamedQuery_Kind}.String()
	NamedQuery_KindAPIVersion   = NamedQuery_Kind + "." + CRDGroupVersion.String()
	NamedQuery_GroupVersionKind = CRDGroupVersion.WithKind(NamedQuery_Kind)
)

func init() {
	SchemeBuilder.Register(&NamedQuery{}, &NamedQueryList{})
}

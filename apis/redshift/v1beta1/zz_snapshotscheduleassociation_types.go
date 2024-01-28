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

type SnapshotScheduleAssociationInitParameters struct {

	// The cluster identifier.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/redshift/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`

	// Reference to a Cluster in redshift to populate clusterIdentifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifierRef *v1.Reference `json:"clusterIdentifierRef,omitempty" tf:"-"`

	// Selector for a Cluster in redshift to populate clusterIdentifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifierSelector *v1.Selector `json:"clusterIdentifierSelector,omitempty" tf:"-"`

	// The snapshot schedule identifier.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/redshift/v1beta1.SnapshotSchedule
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ScheduleIdentifier *string `json:"scheduleIdentifier,omitempty" tf:"schedule_identifier,omitempty"`

	// Reference to a SnapshotSchedule in redshift to populate scheduleIdentifier.
	// +kubebuilder:validation:Optional
	ScheduleIdentifierRef *v1.Reference `json:"scheduleIdentifierRef,omitempty" tf:"-"`

	// Selector for a SnapshotSchedule in redshift to populate scheduleIdentifier.
	// +kubebuilder:validation:Optional
	ScheduleIdentifierSelector *v1.Selector `json:"scheduleIdentifierSelector,omitempty" tf:"-"`
}

type SnapshotScheduleAssociationObservation struct {

	// The cluster identifier.
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The snapshot schedule identifier.
	ScheduleIdentifier *string `json:"scheduleIdentifier,omitempty" tf:"schedule_identifier,omitempty"`
}

type SnapshotScheduleAssociationParameters struct {

	// The cluster identifier.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/redshift/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`

	// Reference to a Cluster in redshift to populate clusterIdentifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifierRef *v1.Reference `json:"clusterIdentifierRef,omitempty" tf:"-"`

	// Selector for a Cluster in redshift to populate clusterIdentifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifierSelector *v1.Selector `json:"clusterIdentifierSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The snapshot schedule identifier.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/redshift/v1beta1.SnapshotSchedule
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ScheduleIdentifier *string `json:"scheduleIdentifier,omitempty" tf:"schedule_identifier,omitempty"`

	// Reference to a SnapshotSchedule in redshift to populate scheduleIdentifier.
	// +kubebuilder:validation:Optional
	ScheduleIdentifierRef *v1.Reference `json:"scheduleIdentifierRef,omitempty" tf:"-"`

	// Selector for a SnapshotSchedule in redshift to populate scheduleIdentifier.
	// +kubebuilder:validation:Optional
	ScheduleIdentifierSelector *v1.Selector `json:"scheduleIdentifierSelector,omitempty" tf:"-"`
}

// SnapshotScheduleAssociationSpec defines the desired state of SnapshotScheduleAssociation
type SnapshotScheduleAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotScheduleAssociationParameters `json:"forProvider"`
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
	InitProvider SnapshotScheduleAssociationInitParameters `json:"initProvider,omitempty"`
}

// SnapshotScheduleAssociationStatus defines the observed state of SnapshotScheduleAssociation.
type SnapshotScheduleAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotScheduleAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SnapshotScheduleAssociation is the Schema for the SnapshotScheduleAssociations API. Provides an Association Redshift Cluster and Snapshot Schedule resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SnapshotScheduleAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnapshotScheduleAssociationSpec   `json:"spec"`
	Status            SnapshotScheduleAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotScheduleAssociationList contains a list of SnapshotScheduleAssociations
type SnapshotScheduleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnapshotScheduleAssociation `json:"items"`
}

// Repository type metadata.
var (
	SnapshotScheduleAssociation_Kind             = "SnapshotScheduleAssociation"
	SnapshotScheduleAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SnapshotScheduleAssociation_Kind}.String()
	SnapshotScheduleAssociation_KindAPIVersion   = SnapshotScheduleAssociation_Kind + "." + CRDGroupVersion.String()
	SnapshotScheduleAssociation_GroupVersionKind = CRDGroupVersion.WithKind(SnapshotScheduleAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&SnapshotScheduleAssociation{}, &SnapshotScheduleAssociationList{})
}

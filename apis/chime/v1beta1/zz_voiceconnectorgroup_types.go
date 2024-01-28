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

type ConnectorInitParameters struct {

	// The priority associated with the Amazon Chime Voice Connector, with 1 being the highest priority. Higher priority Amazon Chime Voice Connectors are attempted first.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The Amazon Chime Voice Connector ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/chime/v1beta1.VoiceConnector
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	VoiceConnectorID *string `json:"voiceConnectorId,omitempty" tf:"voice_connector_id,omitempty"`

	// Reference to a VoiceConnector in chime to populate voiceConnectorId.
	// +kubebuilder:validation:Optional
	VoiceConnectorIDRef *v1.Reference `json:"voiceConnectorIdRef,omitempty" tf:"-"`

	// Selector for a VoiceConnector in chime to populate voiceConnectorId.
	// +kubebuilder:validation:Optional
	VoiceConnectorIDSelector *v1.Selector `json:"voiceConnectorIdSelector,omitempty" tf:"-"`
}

type ConnectorObservation struct {

	// The priority associated with the Amazon Chime Voice Connector, with 1 being the highest priority. Higher priority Amazon Chime Voice Connectors are attempted first.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The Amazon Chime Voice Connector ID.
	VoiceConnectorID *string `json:"voiceConnectorId,omitempty" tf:"voice_connector_id,omitempty"`
}

type ConnectorParameters struct {

	// The priority associated with the Amazon Chime Voice Connector, with 1 being the highest priority. Higher priority Amazon Chime Voice Connectors are attempted first.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority" tf:"priority,omitempty"`

	// The Amazon Chime Voice Connector ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/chime/v1beta1.VoiceConnector
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VoiceConnectorID *string `json:"voiceConnectorId,omitempty" tf:"voice_connector_id,omitempty"`

	// Reference to a VoiceConnector in chime to populate voiceConnectorId.
	// +kubebuilder:validation:Optional
	VoiceConnectorIDRef *v1.Reference `json:"voiceConnectorIdRef,omitempty" tf:"-"`

	// Selector for a VoiceConnector in chime to populate voiceConnectorId.
	// +kubebuilder:validation:Optional
	VoiceConnectorIDSelector *v1.Selector `json:"voiceConnectorIdSelector,omitempty" tf:"-"`
}

type VoiceConnectorGroupInitParameters struct {

	// The Amazon Chime Voice Connectors to route inbound calls to.
	Connector []ConnectorInitParameters `json:"connector,omitempty" tf:"connector,omitempty"`
}

type VoiceConnectorGroupObservation struct {

	// The Amazon Chime Voice Connectors to route inbound calls to.
	Connector []ConnectorObservation `json:"connector,omitempty" tf:"connector,omitempty"`

	// Amazon Chime Voice Connector group ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VoiceConnectorGroupParameters struct {

	// The Amazon Chime Voice Connectors to route inbound calls to.
	// +kubebuilder:validation:Optional
	Connector []ConnectorParameters `json:"connector,omitempty" tf:"connector,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// VoiceConnectorGroupSpec defines the desired state of VoiceConnectorGroup
type VoiceConnectorGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VoiceConnectorGroupParameters `json:"forProvider"`
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
	InitProvider VoiceConnectorGroupInitParameters `json:"initProvider,omitempty"`
}

// VoiceConnectorGroupStatus defines the observed state of VoiceConnectorGroup.
type VoiceConnectorGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VoiceConnectorGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// VoiceConnectorGroup is the Schema for the VoiceConnectorGroups API. Creates an Amazon Chime Voice Connector group under the administrator's AWS account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VoiceConnectorGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VoiceConnectorGroupSpec   `json:"spec"`
	Status            VoiceConnectorGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VoiceConnectorGroupList contains a list of VoiceConnectorGroups
type VoiceConnectorGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VoiceConnectorGroup `json:"items"`
}

// Repository type metadata.
var (
	VoiceConnectorGroup_Kind             = "VoiceConnectorGroup"
	VoiceConnectorGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VoiceConnectorGroup_Kind}.String()
	VoiceConnectorGroup_KindAPIVersion   = VoiceConnectorGroup_Kind + "." + CRDGroupVersion.String()
	VoiceConnectorGroup_GroupVersionKind = CRDGroupVersion.WithKind(VoiceConnectorGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&VoiceConnectorGroup{}, &VoiceConnectorGroupList{})
}

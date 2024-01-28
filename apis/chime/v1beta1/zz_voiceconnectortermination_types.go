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

type VoiceConnectorTerminationInitParameters struct {

	// The countries to which calls are allowed, in ISO 3166-1 alpha-2 format.
	// +listType=set
	CallingRegions []*string `json:"callingRegions,omitempty" tf:"calling_regions,omitempty"`

	// The IP addresses allowed to make calls, in CIDR format.
	// +listType=set
	CidrAllowList []*string `json:"cidrAllowList,omitempty" tf:"cidr_allow_list,omitempty"`

	// The limit on calls per second. Max value based on account service quota. Default value of 1.
	CpsLimit *float64 `json:"cpsLimit,omitempty" tf:"cps_limit,omitempty"`

	// The default caller ID phone number.
	DefaultPhoneNumber *string `json:"defaultPhoneNumber,omitempty" tf:"default_phone_number,omitempty"`

	// When termination settings are disabled, outbound calls can not be made.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

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

type VoiceConnectorTerminationObservation struct {

	// The countries to which calls are allowed, in ISO 3166-1 alpha-2 format.
	// +listType=set
	CallingRegions []*string `json:"callingRegions,omitempty" tf:"calling_regions,omitempty"`

	// The IP addresses allowed to make calls, in CIDR format.
	// +listType=set
	CidrAllowList []*string `json:"cidrAllowList,omitempty" tf:"cidr_allow_list,omitempty"`

	// The limit on calls per second. Max value based on account service quota. Default value of 1.
	CpsLimit *float64 `json:"cpsLimit,omitempty" tf:"cps_limit,omitempty"`

	// The default caller ID phone number.
	DefaultPhoneNumber *string `json:"defaultPhoneNumber,omitempty" tf:"default_phone_number,omitempty"`

	// When termination settings are disabled, outbound calls can not be made.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The Amazon Chime Voice Connector ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Amazon Chime Voice Connector ID.
	VoiceConnectorID *string `json:"voiceConnectorId,omitempty" tf:"voice_connector_id,omitempty"`
}

type VoiceConnectorTerminationParameters struct {

	// The countries to which calls are allowed, in ISO 3166-1 alpha-2 format.
	// +kubebuilder:validation:Optional
	// +listType=set
	CallingRegions []*string `json:"callingRegions,omitempty" tf:"calling_regions,omitempty"`

	// The IP addresses allowed to make calls, in CIDR format.
	// +kubebuilder:validation:Optional
	// +listType=set
	CidrAllowList []*string `json:"cidrAllowList,omitempty" tf:"cidr_allow_list,omitempty"`

	// The limit on calls per second. Max value based on account service quota. Default value of 1.
	// +kubebuilder:validation:Optional
	CpsLimit *float64 `json:"cpsLimit,omitempty" tf:"cps_limit,omitempty"`

	// The default caller ID phone number.
	// +kubebuilder:validation:Optional
	DefaultPhoneNumber *string `json:"defaultPhoneNumber,omitempty" tf:"default_phone_number,omitempty"`

	// When termination settings are disabled, outbound calls can not be made.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

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

// VoiceConnectorTerminationSpec defines the desired state of VoiceConnectorTermination
type VoiceConnectorTerminationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VoiceConnectorTerminationParameters `json:"forProvider"`
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
	InitProvider VoiceConnectorTerminationInitParameters `json:"initProvider,omitempty"`
}

// VoiceConnectorTerminationStatus defines the observed state of VoiceConnectorTermination.
type VoiceConnectorTerminationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VoiceConnectorTerminationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// VoiceConnectorTermination is the Schema for the VoiceConnectorTerminations API. Enable Termination settings to control outbound calling from your SIP infrastructure.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VoiceConnectorTermination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.callingRegions) || (has(self.initProvider) && has(self.initProvider.callingRegions))",message="spec.forProvider.callingRegions is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cidrAllowList) || (has(self.initProvider) && has(self.initProvider.cidrAllowList))",message="spec.forProvider.cidrAllowList is a required parameter"
	Spec   VoiceConnectorTerminationSpec   `json:"spec"`
	Status VoiceConnectorTerminationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VoiceConnectorTerminationList contains a list of VoiceConnectorTerminations
type VoiceConnectorTerminationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VoiceConnectorTermination `json:"items"`
}

// Repository type metadata.
var (
	VoiceConnectorTermination_Kind             = "VoiceConnectorTermination"
	VoiceConnectorTermination_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VoiceConnectorTermination_Kind}.String()
	VoiceConnectorTermination_KindAPIVersion   = VoiceConnectorTermination_Kind + "." + CRDGroupVersion.String()
	VoiceConnectorTermination_GroupVersionKind = CRDGroupVersion.WithKind(VoiceConnectorTermination_Kind)
)

func init() {
	SchemeBuilder.Register(&VoiceConnectorTermination{}, &VoiceConnectorTerminationList{})
}

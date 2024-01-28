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

type VoiceConnectorInitParameters struct {

	// The AWS Region in which the Amazon Chime Voice Connector is created. Default value: us-east-1
	AwsRegion *string `json:"awsRegion,omitempty" tf:"aws_region,omitempty"`

	// When enabled, requires encryption for the Amazon Chime Voice Connector.
	RequireEncryption *bool `json:"requireEncryption,omitempty" tf:"require_encryption,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VoiceConnectorObservation struct {

	// ARN (Amazon Resource Name) of the Amazon Chime Voice Connector.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The AWS Region in which the Amazon Chime Voice Connector is created. Default value: us-east-1
	AwsRegion *string `json:"awsRegion,omitempty" tf:"aws_region,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The outbound host name for the Amazon Chime Voice Connector.
	OutboundHostName *string `json:"outboundHostName,omitempty" tf:"outbound_host_name,omitempty"`

	// When enabled, requires encryption for the Amazon Chime Voice Connector.
	RequireEncryption *bool `json:"requireEncryption,omitempty" tf:"require_encryption,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type VoiceConnectorParameters struct {

	// The AWS Region in which the Amazon Chime Voice Connector is created. Default value: us-east-1
	// +kubebuilder:validation:Optional
	AwsRegion *string `json:"awsRegion,omitempty" tf:"aws_region,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// When enabled, requires encryption for the Amazon Chime Voice Connector.
	// +kubebuilder:validation:Optional
	RequireEncryption *bool `json:"requireEncryption,omitempty" tf:"require_encryption,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VoiceConnectorSpec defines the desired state of VoiceConnector
type VoiceConnectorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VoiceConnectorParameters `json:"forProvider"`
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
	InitProvider VoiceConnectorInitParameters `json:"initProvider,omitempty"`
}

// VoiceConnectorStatus defines the observed state of VoiceConnector.
type VoiceConnectorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VoiceConnectorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// VoiceConnector is the Schema for the VoiceConnectors API. Enables you to connect your phone system to the telephone network at a substantial cost savings by using SIP trunking.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VoiceConnector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.requireEncryption) || (has(self.initProvider) && has(self.initProvider.requireEncryption))",message="spec.forProvider.requireEncryption is a required parameter"
	Spec   VoiceConnectorSpec   `json:"spec"`
	Status VoiceConnectorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VoiceConnectorList contains a list of VoiceConnectors
type VoiceConnectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VoiceConnector `json:"items"`
}

// Repository type metadata.
var (
	VoiceConnector_Kind             = "VoiceConnector"
	VoiceConnector_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VoiceConnector_Kind}.String()
	VoiceConnector_KindAPIVersion   = VoiceConnector_Kind + "." + CRDGroupVersion.String()
	VoiceConnector_GroupVersionKind = CRDGroupVersion.WithKind(VoiceConnector_Kind)
)

func init() {
	SchemeBuilder.Register(&VoiceConnector{}, &VoiceConnectorList{})
}

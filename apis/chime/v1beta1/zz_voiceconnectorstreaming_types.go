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

type MediaInsightsConfigurationInitParameters struct {

	// The media insights configuration that will be invoked by the Voice Connector.
	ConfigurationArn *string `json:"configurationArn,omitempty" tf:"configuration_arn,omitempty"`

	// When true, the media insights configuration is not enabled. Defaults to false.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`
}

type MediaInsightsConfigurationObservation struct {

	// The media insights configuration that will be invoked by the Voice Connector.
	ConfigurationArn *string `json:"configurationArn,omitempty" tf:"configuration_arn,omitempty"`

	// When true, the media insights configuration is not enabled. Defaults to false.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`
}

type MediaInsightsConfigurationParameters struct {

	// The media insights configuration that will be invoked by the Voice Connector.
	// +kubebuilder:validation:Optional
	ConfigurationArn *string `json:"configurationArn,omitempty" tf:"configuration_arn,omitempty"`

	// When true, the media insights configuration is not enabled. Defaults to false.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`
}

type VoiceConnectorStreamingInitParameters struct {

	// The retention period, in hours, for the Amazon Kinesis data.
	DataRetention *float64 `json:"dataRetention,omitempty" tf:"data_retention,omitempty"`

	// When true, media streaming to Amazon Kinesis is turned off. Default: false
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The media insights configuration. See media_insights_configuration.
	MediaInsightsConfiguration []MediaInsightsConfigurationInitParameters `json:"mediaInsightsConfiguration,omitempty" tf:"media_insights_configuration,omitempty"`

	// The streaming notification targets. Valid Values: EventBridge | SNS | SQS
	// +listType=set
	StreamingNotificationTargets []*string `json:"streamingNotificationTargets,omitempty" tf:"streaming_notification_targets,omitempty"`

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

type VoiceConnectorStreamingObservation struct {

	// The retention period, in hours, for the Amazon Kinesis data.
	DataRetention *float64 `json:"dataRetention,omitempty" tf:"data_retention,omitempty"`

	// When true, media streaming to Amazon Kinesis is turned off. Default: false
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The Amazon Chime Voice Connector ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The media insights configuration. See media_insights_configuration.
	MediaInsightsConfiguration []MediaInsightsConfigurationObservation `json:"mediaInsightsConfiguration,omitempty" tf:"media_insights_configuration,omitempty"`

	// The streaming notification targets. Valid Values: EventBridge | SNS | SQS
	// +listType=set
	StreamingNotificationTargets []*string `json:"streamingNotificationTargets,omitempty" tf:"streaming_notification_targets,omitempty"`

	// The Amazon Chime Voice Connector ID.
	VoiceConnectorID *string `json:"voiceConnectorId,omitempty" tf:"voice_connector_id,omitempty"`
}

type VoiceConnectorStreamingParameters struct {

	// The retention period, in hours, for the Amazon Kinesis data.
	// +kubebuilder:validation:Optional
	DataRetention *float64 `json:"dataRetention,omitempty" tf:"data_retention,omitempty"`

	// When true, media streaming to Amazon Kinesis is turned off. Default: false
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The media insights configuration. See media_insights_configuration.
	// +kubebuilder:validation:Optional
	MediaInsightsConfiguration []MediaInsightsConfigurationParameters `json:"mediaInsightsConfiguration,omitempty" tf:"media_insights_configuration,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The streaming notification targets. Valid Values: EventBridge | SNS | SQS
	// +kubebuilder:validation:Optional
	// +listType=set
	StreamingNotificationTargets []*string `json:"streamingNotificationTargets,omitempty" tf:"streaming_notification_targets,omitempty"`

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

// VoiceConnectorStreamingSpec defines the desired state of VoiceConnectorStreaming
type VoiceConnectorStreamingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VoiceConnectorStreamingParameters `json:"forProvider"`
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
	InitProvider VoiceConnectorStreamingInitParameters `json:"initProvider,omitempty"`
}

// VoiceConnectorStreamingStatus defines the observed state of VoiceConnectorStreaming.
type VoiceConnectorStreamingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VoiceConnectorStreamingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// VoiceConnectorStreaming is the Schema for the VoiceConnectorStreamings API. The streaming configuration associated with an Amazon Chime Voice Connector. Specifies whether media streaming is enabled for sending to Amazon Kinesis, and shows the retention period for the Amazon Kinesis data, in hours.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VoiceConnectorStreaming struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dataRetention) || (has(self.initProvider) && has(self.initProvider.dataRetention))",message="spec.forProvider.dataRetention is a required parameter"
	Spec   VoiceConnectorStreamingSpec   `json:"spec"`
	Status VoiceConnectorStreamingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VoiceConnectorStreamingList contains a list of VoiceConnectorStreamings
type VoiceConnectorStreamingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VoiceConnectorStreaming `json:"items"`
}

// Repository type metadata.
var (
	VoiceConnectorStreaming_Kind             = "VoiceConnectorStreaming"
	VoiceConnectorStreaming_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VoiceConnectorStreaming_Kind}.String()
	VoiceConnectorStreaming_KindAPIVersion   = VoiceConnectorStreaming_Kind + "." + CRDGroupVersion.String()
	VoiceConnectorStreaming_GroupVersionKind = CRDGroupVersion.WithKind(VoiceConnectorStreaming_Kind)
)

func init() {
	SchemeBuilder.Register(&VoiceConnectorStreaming{}, &VoiceConnectorStreamingList{})
}

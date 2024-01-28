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

type CloudWatchDestinationInitParameters struct {

	// An array of objects that define the dimensions to use when you send email events to Amazon CloudWatch. See dimension_configuration below.
	DimensionConfiguration []DimensionConfigurationInitParameters `json:"dimensionConfiguration,omitempty" tf:"dimension_configuration,omitempty"`
}

type CloudWatchDestinationObservation struct {

	// An array of objects that define the dimensions to use when you send email events to Amazon CloudWatch. See dimension_configuration below.
	DimensionConfiguration []DimensionConfigurationObservation `json:"dimensionConfiguration,omitempty" tf:"dimension_configuration,omitempty"`
}

type CloudWatchDestinationParameters struct {

	// An array of objects that define the dimensions to use when you send email events to Amazon CloudWatch. See dimension_configuration below.
	// +kubebuilder:validation:Optional
	DimensionConfiguration []DimensionConfigurationParameters `json:"dimensionConfiguration" tf:"dimension_configuration,omitempty"`
}

type ConfigurationSetEventDestinationInitParameters struct {

	// The name of the configuration set.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sesv2/v1beta1.ConfigurationSet
	ConfigurationSetName *string `json:"configurationSetName,omitempty" tf:"configuration_set_name,omitempty"`

	// Reference to a ConfigurationSet in sesv2 to populate configurationSetName.
	// +kubebuilder:validation:Optional
	ConfigurationSetNameRef *v1.Reference `json:"configurationSetNameRef,omitempty" tf:"-"`

	// Selector for a ConfigurationSet in sesv2 to populate configurationSetName.
	// +kubebuilder:validation:Optional
	ConfigurationSetNameSelector *v1.Selector `json:"configurationSetNameSelector,omitempty" tf:"-"`

	// A name that identifies the event destination within the configuration set.
	EventDestination []EventDestinationInitParameters `json:"eventDestination,omitempty" tf:"event_destination,omitempty"`

	// An object that defines the event destination. See event_destination below.
	EventDestinationName *string `json:"eventDestinationName,omitempty" tf:"event_destination_name,omitempty"`
}

type ConfigurationSetEventDestinationObservation struct {

	// The name of the configuration set.
	ConfigurationSetName *string `json:"configurationSetName,omitempty" tf:"configuration_set_name,omitempty"`

	// A name that identifies the event destination within the configuration set.
	EventDestination []EventDestinationObservation `json:"eventDestination,omitempty" tf:"event_destination,omitempty"`

	// An object that defines the event destination. See event_destination below.
	EventDestinationName *string `json:"eventDestinationName,omitempty" tf:"event_destination_name,omitempty"`

	// A pipe-delimited string combining configuration_set_name and event_destination_name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ConfigurationSetEventDestinationParameters struct {

	// The name of the configuration set.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sesv2/v1beta1.ConfigurationSet
	// +kubebuilder:validation:Optional
	ConfigurationSetName *string `json:"configurationSetName,omitempty" tf:"configuration_set_name,omitempty"`

	// Reference to a ConfigurationSet in sesv2 to populate configurationSetName.
	// +kubebuilder:validation:Optional
	ConfigurationSetNameRef *v1.Reference `json:"configurationSetNameRef,omitempty" tf:"-"`

	// Selector for a ConfigurationSet in sesv2 to populate configurationSetName.
	// +kubebuilder:validation:Optional
	ConfigurationSetNameSelector *v1.Selector `json:"configurationSetNameSelector,omitempty" tf:"-"`

	// A name that identifies the event destination within the configuration set.
	// +kubebuilder:validation:Optional
	EventDestination []EventDestinationParameters `json:"eventDestination,omitempty" tf:"event_destination,omitempty"`

	// An object that defines the event destination. See event_destination below.
	// +kubebuilder:validation:Optional
	EventDestinationName *string `json:"eventDestinationName,omitempty" tf:"event_destination_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type DimensionConfigurationInitParameters struct {

	// The default value of the dimension that is published to Amazon CloudWatch if you don't provide the value of the dimension when you send an email.
	DefaultDimensionValue *string `json:"defaultDimensionValue,omitempty" tf:"default_dimension_value,omitempty"`

	// The name of an Amazon CloudWatch dimension associated with an email sending metric.
	DimensionName *string `json:"dimensionName,omitempty" tf:"dimension_name,omitempty"`

	// The location where the Amazon SES API v2 finds the value of a dimension to publish to Amazon CloudWatch. Valid values: MESSAGE_TAG, EMAIL_HEADER, LINK_TAG.
	DimensionValueSource *string `json:"dimensionValueSource,omitempty" tf:"dimension_value_source,omitempty"`
}

type DimensionConfigurationObservation struct {

	// The default value of the dimension that is published to Amazon CloudWatch if you don't provide the value of the dimension when you send an email.
	DefaultDimensionValue *string `json:"defaultDimensionValue,omitempty" tf:"default_dimension_value,omitempty"`

	// The name of an Amazon CloudWatch dimension associated with an email sending metric.
	DimensionName *string `json:"dimensionName,omitempty" tf:"dimension_name,omitempty"`

	// The location where the Amazon SES API v2 finds the value of a dimension to publish to Amazon CloudWatch. Valid values: MESSAGE_TAG, EMAIL_HEADER, LINK_TAG.
	DimensionValueSource *string `json:"dimensionValueSource,omitempty" tf:"dimension_value_source,omitempty"`
}

type DimensionConfigurationParameters struct {

	// The default value of the dimension that is published to Amazon CloudWatch if you don't provide the value of the dimension when you send an email.
	// +kubebuilder:validation:Optional
	DefaultDimensionValue *string `json:"defaultDimensionValue" tf:"default_dimension_value,omitempty"`

	// The name of an Amazon CloudWatch dimension associated with an email sending metric.
	// +kubebuilder:validation:Optional
	DimensionName *string `json:"dimensionName" tf:"dimension_name,omitempty"`

	// The location where the Amazon SES API v2 finds the value of a dimension to publish to Amazon CloudWatch. Valid values: MESSAGE_TAG, EMAIL_HEADER, LINK_TAG.
	// +kubebuilder:validation:Optional
	DimensionValueSource *string `json:"dimensionValueSource" tf:"dimension_value_source,omitempty"`
}

type EventDestinationInitParameters struct {

	// An object that defines an Amazon CloudWatch destination for email events. See cloud_watch_destination below
	CloudWatchDestination []CloudWatchDestinationInitParameters `json:"cloudWatchDestination,omitempty" tf:"cloud_watch_destination,omitempty"`

	// When the event destination is enabled, the specified event types are sent to the destinations. Default: false.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// An object that defines an Amazon Kinesis Data Firehose destination for email events. See kinesis_firehose_destination below.
	KinesisFirehoseDestination []KinesisFirehoseDestinationInitParameters `json:"kinesisFirehoseDestination,omitempty" tf:"kinesis_firehose_destination,omitempty"`

	// - An array that specifies which events the Amazon SES API v2 should send to the destinations. Valid values: SEND, REJECT, BOUNCE, COMPLAINT, DELIVERY, OPEN, CLICK, RENDERING_FAILURE, DELIVERY_DELAY, SUBSCRIPTION.
	MatchingEventTypes []*string `json:"matchingEventTypes,omitempty" tf:"matching_event_types,omitempty"`

	// An object that defines an Amazon Pinpoint project destination for email events. See pinpoint_destination below.
	PinpointDestination []PinpointDestinationInitParameters `json:"pinpointDestination,omitempty" tf:"pinpoint_destination,omitempty"`

	// An object that defines an Amazon SNS destination for email events. See sns_destination below.
	SnsDestination []SnsDestinationInitParameters `json:"snsDestination,omitempty" tf:"sns_destination,omitempty"`
}

type EventDestinationObservation struct {

	// An object that defines an Amazon CloudWatch destination for email events. See cloud_watch_destination below
	CloudWatchDestination []CloudWatchDestinationObservation `json:"cloudWatchDestination,omitempty" tf:"cloud_watch_destination,omitempty"`

	// When the event destination is enabled, the specified event types are sent to the destinations. Default: false.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// An object that defines an Amazon Kinesis Data Firehose destination for email events. See kinesis_firehose_destination below.
	KinesisFirehoseDestination []KinesisFirehoseDestinationObservation `json:"kinesisFirehoseDestination,omitempty" tf:"kinesis_firehose_destination,omitempty"`

	// - An array that specifies which events the Amazon SES API v2 should send to the destinations. Valid values: SEND, REJECT, BOUNCE, COMPLAINT, DELIVERY, OPEN, CLICK, RENDERING_FAILURE, DELIVERY_DELAY, SUBSCRIPTION.
	MatchingEventTypes []*string `json:"matchingEventTypes,omitempty" tf:"matching_event_types,omitempty"`

	// An object that defines an Amazon Pinpoint project destination for email events. See pinpoint_destination below.
	PinpointDestination []PinpointDestinationObservation `json:"pinpointDestination,omitempty" tf:"pinpoint_destination,omitempty"`

	// An object that defines an Amazon SNS destination for email events. See sns_destination below.
	SnsDestination []SnsDestinationObservation `json:"snsDestination,omitempty" tf:"sns_destination,omitempty"`
}

type EventDestinationParameters struct {

	// An object that defines an Amazon CloudWatch destination for email events. See cloud_watch_destination below
	// +kubebuilder:validation:Optional
	CloudWatchDestination []CloudWatchDestinationParameters `json:"cloudWatchDestination,omitempty" tf:"cloud_watch_destination,omitempty"`

	// When the event destination is enabled, the specified event types are sent to the destinations. Default: false.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// An object that defines an Amazon Kinesis Data Firehose destination for email events. See kinesis_firehose_destination below.
	// +kubebuilder:validation:Optional
	KinesisFirehoseDestination []KinesisFirehoseDestinationParameters `json:"kinesisFirehoseDestination,omitempty" tf:"kinesis_firehose_destination,omitempty"`

	// - An array that specifies which events the Amazon SES API v2 should send to the destinations. Valid values: SEND, REJECT, BOUNCE, COMPLAINT, DELIVERY, OPEN, CLICK, RENDERING_FAILURE, DELIVERY_DELAY, SUBSCRIPTION.
	// +kubebuilder:validation:Optional
	MatchingEventTypes []*string `json:"matchingEventTypes" tf:"matching_event_types,omitempty"`

	// An object that defines an Amazon Pinpoint project destination for email events. See pinpoint_destination below.
	// +kubebuilder:validation:Optional
	PinpointDestination []PinpointDestinationParameters `json:"pinpointDestination,omitempty" tf:"pinpoint_destination,omitempty"`

	// An object that defines an Amazon SNS destination for email events. See sns_destination below.
	// +kubebuilder:validation:Optional
	SnsDestination []SnsDestinationParameters `json:"snsDestination,omitempty" tf:"sns_destination,omitempty"`
}

type KinesisFirehoseDestinationInitParameters struct {

	// The Amazon Resource Name (ARN) of the Amazon Kinesis Data Firehose stream that the Amazon SES API v2 sends email events to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/firehose/v1beta1.DeliveryStream
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",false)
	DeliveryStreamArn *string `json:"deliveryStreamArn,omitempty" tf:"delivery_stream_arn,omitempty"`

	// Reference to a DeliveryStream in firehose to populate deliveryStreamArn.
	// +kubebuilder:validation:Optional
	DeliveryStreamArnRef *v1.Reference `json:"deliveryStreamArnRef,omitempty" tf:"-"`

	// Selector for a DeliveryStream in firehose to populate deliveryStreamArn.
	// +kubebuilder:validation:Optional
	DeliveryStreamArnSelector *v1.Selector `json:"deliveryStreamArnSelector,omitempty" tf:"-"`

	// The Amazon Resource Name (ARN) of the IAM role that the Amazon SES API v2 uses to send email events to the Amazon Kinesis Data Firehose stream.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	IAMRoleArn *string `json:"iamRoleArn,omitempty" tf:"iam_role_arn,omitempty"`

	// Reference to a Role in iam to populate iamRoleArn.
	// +kubebuilder:validation:Optional
	IAMRoleArnRef *v1.Reference `json:"iamRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate iamRoleArn.
	// +kubebuilder:validation:Optional
	IAMRoleArnSelector *v1.Selector `json:"iamRoleArnSelector,omitempty" tf:"-"`
}

type KinesisFirehoseDestinationObservation struct {

	// The Amazon Resource Name (ARN) of the Amazon Kinesis Data Firehose stream that the Amazon SES API v2 sends email events to.
	DeliveryStreamArn *string `json:"deliveryStreamArn,omitempty" tf:"delivery_stream_arn,omitempty"`

	// The Amazon Resource Name (ARN) of the IAM role that the Amazon SES API v2 uses to send email events to the Amazon Kinesis Data Firehose stream.
	IAMRoleArn *string `json:"iamRoleArn,omitempty" tf:"iam_role_arn,omitempty"`
}

type KinesisFirehoseDestinationParameters struct {

	// The Amazon Resource Name (ARN) of the Amazon Kinesis Data Firehose stream that the Amazon SES API v2 sends email events to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/firehose/v1beta1.DeliveryStream
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",false)
	// +kubebuilder:validation:Optional
	DeliveryStreamArn *string `json:"deliveryStreamArn,omitempty" tf:"delivery_stream_arn,omitempty"`

	// Reference to a DeliveryStream in firehose to populate deliveryStreamArn.
	// +kubebuilder:validation:Optional
	DeliveryStreamArnRef *v1.Reference `json:"deliveryStreamArnRef,omitempty" tf:"-"`

	// Selector for a DeliveryStream in firehose to populate deliveryStreamArn.
	// +kubebuilder:validation:Optional
	DeliveryStreamArnSelector *v1.Selector `json:"deliveryStreamArnSelector,omitempty" tf:"-"`

	// The Amazon Resource Name (ARN) of the IAM role that the Amazon SES API v2 uses to send email events to the Amazon Kinesis Data Firehose stream.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	IAMRoleArn *string `json:"iamRoleArn,omitempty" tf:"iam_role_arn,omitempty"`

	// Reference to a Role in iam to populate iamRoleArn.
	// +kubebuilder:validation:Optional
	IAMRoleArnRef *v1.Reference `json:"iamRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate iamRoleArn.
	// +kubebuilder:validation:Optional
	IAMRoleArnSelector *v1.Selector `json:"iamRoleArnSelector,omitempty" tf:"-"`
}

type PinpointDestinationInitParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/pinpoint/v1beta1.App
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	ApplicationArn *string `json:"applicationArn,omitempty" tf:"application_arn,omitempty"`

	// Reference to a App in pinpoint to populate applicationArn.
	// +kubebuilder:validation:Optional
	ApplicationArnRef *v1.Reference `json:"applicationArnRef,omitempty" tf:"-"`

	// Selector for a App in pinpoint to populate applicationArn.
	// +kubebuilder:validation:Optional
	ApplicationArnSelector *v1.Selector `json:"applicationArnSelector,omitempty" tf:"-"`
}

type PinpointDestinationObservation struct {
	ApplicationArn *string `json:"applicationArn,omitempty" tf:"application_arn,omitempty"`
}

type PinpointDestinationParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/pinpoint/v1beta1.App
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ApplicationArn *string `json:"applicationArn,omitempty" tf:"application_arn,omitempty"`

	// Reference to a App in pinpoint to populate applicationArn.
	// +kubebuilder:validation:Optional
	ApplicationArnRef *v1.Reference `json:"applicationArnRef,omitempty" tf:"-"`

	// Selector for a App in pinpoint to populate applicationArn.
	// +kubebuilder:validation:Optional
	ApplicationArnSelector *v1.Selector `json:"applicationArnSelector,omitempty" tf:"-"`
}

type SnsDestinationInitParameters struct {

	// The Amazon Resource Name (ARN) of the Amazon SNS topic to publish email events to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`

	// Reference to a Topic in sns to populate topicArn.
	// +kubebuilder:validation:Optional
	TopicArnRef *v1.Reference `json:"topicArnRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate topicArn.
	// +kubebuilder:validation:Optional
	TopicArnSelector *v1.Selector `json:"topicArnSelector,omitempty" tf:"-"`
}

type SnsDestinationObservation struct {

	// The Amazon Resource Name (ARN) of the Amazon SNS topic to publish email events to.
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`
}

type SnsDestinationParameters struct {

	// The Amazon Resource Name (ARN) of the Amazon SNS topic to publish email events to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`

	// Reference to a Topic in sns to populate topicArn.
	// +kubebuilder:validation:Optional
	TopicArnRef *v1.Reference `json:"topicArnRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate topicArn.
	// +kubebuilder:validation:Optional
	TopicArnSelector *v1.Selector `json:"topicArnSelector,omitempty" tf:"-"`
}

// ConfigurationSetEventDestinationSpec defines the desired state of ConfigurationSetEventDestination
type ConfigurationSetEventDestinationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConfigurationSetEventDestinationParameters `json:"forProvider"`
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
	InitProvider ConfigurationSetEventDestinationInitParameters `json:"initProvider,omitempty"`
}

// ConfigurationSetEventDestinationStatus defines the observed state of ConfigurationSetEventDestination.
type ConfigurationSetEventDestinationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConfigurationSetEventDestinationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ConfigurationSetEventDestination is the Schema for the ConfigurationSetEventDestinations API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ConfigurationSetEventDestination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.eventDestination) || (has(self.initProvider) && has(self.initProvider.eventDestination))",message="spec.forProvider.eventDestination is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.eventDestinationName) || (has(self.initProvider) && has(self.initProvider.eventDestinationName))",message="spec.forProvider.eventDestinationName is a required parameter"
	Spec   ConfigurationSetEventDestinationSpec   `json:"spec"`
	Status ConfigurationSetEventDestinationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigurationSetEventDestinationList contains a list of ConfigurationSetEventDestinations
type ConfigurationSetEventDestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigurationSetEventDestination `json:"items"`
}

// Repository type metadata.
var (
	ConfigurationSetEventDestination_Kind             = "ConfigurationSetEventDestination"
	ConfigurationSetEventDestination_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConfigurationSetEventDestination_Kind}.String()
	ConfigurationSetEventDestination_KindAPIVersion   = ConfigurationSetEventDestination_Kind + "." + CRDGroupVersion.String()
	ConfigurationSetEventDestination_GroupVersionKind = CRDGroupVersion.WithKind(ConfigurationSetEventDestination_Kind)
)

func init() {
	SchemeBuilder.Register(&ConfigurationSetEventDestination{}, &ConfigurationSetEventDestinationList{})
}

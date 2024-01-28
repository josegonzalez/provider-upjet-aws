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

type ActionInitParameters struct {

	// Information about the action.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The action name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An Amazon Resource Name (ARN) for an Identity and Access Management assume role.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// The extension URI associated to the action point in the extension definition. The URI can be an Amazon Resource Name (ARN) for one of the following: an Lambda function, an Amazon Simple Queue Service queue, an Amazon Simple Notification Service topic, or the Amazon EventBridge default event bus.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`

	// Reference to a Topic in sns to populate uri.
	// +kubebuilder:validation:Optional
	URIRef *v1.Reference `json:"uriRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate uri.
	// +kubebuilder:validation:Optional
	URISelector *v1.Selector `json:"uriSelector,omitempty" tf:"-"`
}

type ActionObservation struct {

	// Information about the action.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The action name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An Amazon Resource Name (ARN) for an Identity and Access Management assume role.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// The extension URI associated to the action point in the extension definition. The URI can be an Amazon Resource Name (ARN) for one of the following: an Lambda function, an Amazon Simple Queue Service queue, an Amazon Simple Notification Service topic, or the Amazon EventBridge default event bus.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type ActionParameters struct {

	// Information about the action.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The action name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// An Amazon Resource Name (ARN) for an Identity and Access Management assume role.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// The extension URI associated to the action point in the extension definition. The URI can be an Amazon Resource Name (ARN) for one of the following: an Lambda function, an Amazon Simple Queue Service queue, an Amazon Simple Notification Service topic, or the Amazon EventBridge default event bus.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`

	// Reference to a Topic in sns to populate uri.
	// +kubebuilder:validation:Optional
	URIRef *v1.Reference `json:"uriRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate uri.
	// +kubebuilder:validation:Optional
	URISelector *v1.Selector `json:"uriSelector,omitempty" tf:"-"`
}

type ActionPointInitParameters struct {

	// An action defines the tasks the extension performs during the AppConfig workflow. Detailed below.
	Action []ActionInitParameters `json:"action,omitempty" tf:"action,omitempty"`

	// The point at which to perform the defined actions. Valid points are PRE_CREATE_HOSTED_CONFIGURATION_VERSION, PRE_START_DEPLOYMENT, ON_DEPLOYMENT_START, ON_DEPLOYMENT_STEP, ON_DEPLOYMENT_BAKING, ON_DEPLOYMENT_COMPLETE, ON_DEPLOYMENT_ROLLED_BACK.
	Point *string `json:"point,omitempty" tf:"point,omitempty"`
}

type ActionPointObservation struct {

	// An action defines the tasks the extension performs during the AppConfig workflow. Detailed below.
	Action []ActionObservation `json:"action,omitempty" tf:"action,omitempty"`

	// The point at which to perform the defined actions. Valid points are PRE_CREATE_HOSTED_CONFIGURATION_VERSION, PRE_START_DEPLOYMENT, ON_DEPLOYMENT_START, ON_DEPLOYMENT_STEP, ON_DEPLOYMENT_BAKING, ON_DEPLOYMENT_COMPLETE, ON_DEPLOYMENT_ROLLED_BACK.
	Point *string `json:"point,omitempty" tf:"point,omitempty"`
}

type ActionPointParameters struct {

	// An action defines the tasks the extension performs during the AppConfig workflow. Detailed below.
	// +kubebuilder:validation:Optional
	Action []ActionParameters `json:"action" tf:"action,omitempty"`

	// The point at which to perform the defined actions. Valid points are PRE_CREATE_HOSTED_CONFIGURATION_VERSION, PRE_START_DEPLOYMENT, ON_DEPLOYMENT_START, ON_DEPLOYMENT_STEP, ON_DEPLOYMENT_BAKING, ON_DEPLOYMENT_COMPLETE, ON_DEPLOYMENT_ROLLED_BACK.
	// +kubebuilder:validation:Optional
	Point *string `json:"point" tf:"point,omitempty"`
}

type ExtensionInitParameters struct {

	// The action points defined in the extension. Detailed below.
	ActionPoint []ActionPointInitParameters `json:"actionPoint,omitempty" tf:"action_point,omitempty"`

	// Information about the extension.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A name for the extension. Each extension name in your account must be unique. Extension versions use the same name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The parameters accepted by the extension. You specify parameter values when you associate the extension to an AppConfig resource by using the CreateExtensionAssociation API action. For Lambda extension actions, these parameters are included in the Lambda request object. Detailed below.
	Parameter []ParameterInitParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ExtensionObservation struct {

	// The action points defined in the extension. Detailed below.
	ActionPoint []ActionPointObservation `json:"actionPoint,omitempty" tf:"action_point,omitempty"`

	// ARN of the AppConfig Extension.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Information about the extension.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// AppConfig Extension ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A name for the extension. Each extension name in your account must be unique. Extension versions use the same name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The parameters accepted by the extension. You specify parameter values when you associate the extension to an AppConfig resource by using the CreateExtensionAssociation API action. For Lambda extension actions, these parameters are included in the Lambda request object. Detailed below.
	Parameter []ParameterObservation `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The version number for the extension.
	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

type ExtensionParameters struct {

	// The action points defined in the extension. Detailed below.
	// +kubebuilder:validation:Optional
	ActionPoint []ActionPointParameters `json:"actionPoint,omitempty" tf:"action_point,omitempty"`

	// Information about the extension.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A name for the extension. Each extension name in your account must be unique. Extension versions use the same name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The parameters accepted by the extension. You specify parameter values when you associate the extension to an AppConfig resource by using the CreateExtensionAssociation API action. For Lambda extension actions, these parameters are included in the Lambda request object. Detailed below.
	// +kubebuilder:validation:Optional
	Parameter []ParameterParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ParameterInitParameters struct {

	// Information about the parameter.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The parameter name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Determines if a parameter value must be specified in the extension association.
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`
}

type ParameterObservation struct {

	// Information about the parameter.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The parameter name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Determines if a parameter value must be specified in the extension association.
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`
}

type ParameterParameters struct {

	// Information about the parameter.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The parameter name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Determines if a parameter value must be specified in the extension association.
	// +kubebuilder:validation:Optional
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`
}

// ExtensionSpec defines the desired state of Extension
type ExtensionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExtensionParameters `json:"forProvider"`
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
	InitProvider ExtensionInitParameters `json:"initProvider,omitempty"`
}

// ExtensionStatus defines the observed state of Extension.
type ExtensionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExtensionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Extension is the Schema for the Extensions API. Provides an AppConfig Extension resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Extension struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.actionPoint) || (has(self.initProvider) && has(self.initProvider.actionPoint))",message="spec.forProvider.actionPoint is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ExtensionSpec   `json:"spec"`
	Status ExtensionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExtensionList contains a list of Extensions
type ExtensionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Extension `json:"items"`
}

// Repository type metadata.
var (
	Extension_Kind             = "Extension"
	Extension_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Extension_Kind}.String()
	Extension_KindAPIVersion   = Extension_Kind + "." + CRDGroupVersion.String()
	Extension_GroupVersionKind = CRDGroupVersion.WithKind(Extension_Kind)
)

func init() {
	SchemeBuilder.Register(&Extension{}, &ExtensionList{})
}

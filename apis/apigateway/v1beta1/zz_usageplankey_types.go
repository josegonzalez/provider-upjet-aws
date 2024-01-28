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

type UsagePlanKeyInitParameters struct {

	// Identifier of the API key resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.APIKey
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// Reference to a APIKey in apigateway to populate keyId.
	// +kubebuilder:validation:Optional
	KeyIDRef *v1.Reference `json:"keyIdRef,omitempty" tf:"-"`

	// Selector for a APIKey in apigateway to populate keyId.
	// +kubebuilder:validation:Optional
	KeyIDSelector *v1.Selector `json:"keyIdSelector,omitempty" tf:"-"`

	// Type of the API key resource. Currently, the valid key type is API_KEY.
	KeyType *string `json:"keyType,omitempty" tf:"key_type,omitempty"`

	// Id of the usage plan resource representing to associate the key to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.UsagePlan
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	UsagePlanID *string `json:"usagePlanId,omitempty" tf:"usage_plan_id,omitempty"`

	// Reference to a UsagePlan in apigateway to populate usagePlanId.
	// +kubebuilder:validation:Optional
	UsagePlanIDRef *v1.Reference `json:"usagePlanIdRef,omitempty" tf:"-"`

	// Selector for a UsagePlan in apigateway to populate usagePlanId.
	// +kubebuilder:validation:Optional
	UsagePlanIDSelector *v1.Selector `json:"usagePlanIdSelector,omitempty" tf:"-"`
}

type UsagePlanKeyObservation struct {

	// ID of a usage plan key.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Identifier of the API key resource.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// Type of the API key resource. Currently, the valid key type is API_KEY.
	KeyType *string `json:"keyType,omitempty" tf:"key_type,omitempty"`

	// Name of a usage plan key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Id of the usage plan resource representing to associate the key to.
	UsagePlanID *string `json:"usagePlanId,omitempty" tf:"usage_plan_id,omitempty"`

	// Value of a usage plan key.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type UsagePlanKeyParameters struct {

	// Identifier of the API key resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.APIKey
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// Reference to a APIKey in apigateway to populate keyId.
	// +kubebuilder:validation:Optional
	KeyIDRef *v1.Reference `json:"keyIdRef,omitempty" tf:"-"`

	// Selector for a APIKey in apigateway to populate keyId.
	// +kubebuilder:validation:Optional
	KeyIDSelector *v1.Selector `json:"keyIdSelector,omitempty" tf:"-"`

	// Type of the API key resource. Currently, the valid key type is API_KEY.
	// +kubebuilder:validation:Optional
	KeyType *string `json:"keyType,omitempty" tf:"key_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Id of the usage plan resource representing to associate the key to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.UsagePlan
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	UsagePlanID *string `json:"usagePlanId,omitempty" tf:"usage_plan_id,omitempty"`

	// Reference to a UsagePlan in apigateway to populate usagePlanId.
	// +kubebuilder:validation:Optional
	UsagePlanIDRef *v1.Reference `json:"usagePlanIdRef,omitempty" tf:"-"`

	// Selector for a UsagePlan in apigateway to populate usagePlanId.
	// +kubebuilder:validation:Optional
	UsagePlanIDSelector *v1.Selector `json:"usagePlanIdSelector,omitempty" tf:"-"`
}

// UsagePlanKeySpec defines the desired state of UsagePlanKey
type UsagePlanKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UsagePlanKeyParameters `json:"forProvider"`
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
	InitProvider UsagePlanKeyInitParameters `json:"initProvider,omitempty"`
}

// UsagePlanKeyStatus defines the observed state of UsagePlanKey.
type UsagePlanKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UsagePlanKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// UsagePlanKey is the Schema for the UsagePlanKeys API. Provides an API Gateway Usage Plan Key.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type UsagePlanKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.keyType) || (has(self.initProvider) && has(self.initProvider.keyType))",message="spec.forProvider.keyType is a required parameter"
	Spec   UsagePlanKeySpec   `json:"spec"`
	Status UsagePlanKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UsagePlanKeyList contains a list of UsagePlanKeys
type UsagePlanKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UsagePlanKey `json:"items"`
}

// Repository type metadata.
var (
	UsagePlanKey_Kind             = "UsagePlanKey"
	UsagePlanKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UsagePlanKey_Kind}.String()
	UsagePlanKey_KindAPIVersion   = UsagePlanKey_Kind + "." + CRDGroupVersion.String()
	UsagePlanKey_GroupVersionKind = CRDGroupVersion.WithKind(UsagePlanKey_Kind)
)

func init() {
	SchemeBuilder.Register(&UsagePlanKey{}, &UsagePlanKeyList{})
}

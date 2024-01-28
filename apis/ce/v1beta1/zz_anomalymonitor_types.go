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

type AnomalyMonitorInitParameters struct {

	// The dimensions to evaluate. Valid values: SERVICE.
	MonitorDimension *string `json:"monitorDimension,omitempty" tf:"monitor_dimension,omitempty"`

	// A valid JSON representation for the Expression object.
	MonitorSpecification *string `json:"monitorSpecification,omitempty" tf:"monitor_specification,omitempty"`

	// The possible type values. Valid values: DIMENSIONAL | CUSTOM.
	MonitorType *string `json:"monitorType,omitempty" tf:"monitor_type,omitempty"`

	// The name of the monitor.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AnomalyMonitorObservation struct {

	// ARN of the anomaly monitor.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Unique ID of the anomaly monitor. Same as arn.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The dimensions to evaluate. Valid values: SERVICE.
	MonitorDimension *string `json:"monitorDimension,omitempty" tf:"monitor_dimension,omitempty"`

	// A valid JSON representation for the Expression object.
	MonitorSpecification *string `json:"monitorSpecification,omitempty" tf:"monitor_specification,omitempty"`

	// The possible type values. Valid values: DIMENSIONAL | CUSTOM.
	MonitorType *string `json:"monitorType,omitempty" tf:"monitor_type,omitempty"`

	// The name of the monitor.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type AnomalyMonitorParameters struct {

	// The dimensions to evaluate. Valid values: SERVICE.
	// +kubebuilder:validation:Optional
	MonitorDimension *string `json:"monitorDimension,omitempty" tf:"monitor_dimension,omitempty"`

	// A valid JSON representation for the Expression object.
	// +kubebuilder:validation:Optional
	MonitorSpecification *string `json:"monitorSpecification,omitempty" tf:"monitor_specification,omitempty"`

	// The possible type values. Valid values: DIMENSIONAL | CUSTOM.
	// +kubebuilder:validation:Optional
	MonitorType *string `json:"monitorType,omitempty" tf:"monitor_type,omitempty"`

	// The name of the monitor.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// AnomalyMonitorSpec defines the desired state of AnomalyMonitor
type AnomalyMonitorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AnomalyMonitorParameters `json:"forProvider"`
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
	InitProvider AnomalyMonitorInitParameters `json:"initProvider,omitempty"`
}

// AnomalyMonitorStatus defines the observed state of AnomalyMonitor.
type AnomalyMonitorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AnomalyMonitorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AnomalyMonitor is the Schema for the AnomalyMonitors API. Provides a CE Cost Anomaly Monitor
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AnomalyMonitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.monitorType) || (has(self.initProvider) && has(self.initProvider.monitorType))",message="spec.forProvider.monitorType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   AnomalyMonitorSpec   `json:"spec"`
	Status AnomalyMonitorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AnomalyMonitorList contains a list of AnomalyMonitors
type AnomalyMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AnomalyMonitor `json:"items"`
}

// Repository type metadata.
var (
	AnomalyMonitor_Kind             = "AnomalyMonitor"
	AnomalyMonitor_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AnomalyMonitor_Kind}.String()
	AnomalyMonitor_KindAPIVersion   = AnomalyMonitor_Kind + "." + CRDGroupVersion.String()
	AnomalyMonitor_GroupVersionKind = CRDGroupVersion.WithKind(AnomalyMonitor_Kind)
)

func init() {
	SchemeBuilder.Register(&AnomalyMonitor{}, &AnomalyMonitorList{})
}

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

type VPCIpamPoolCidrAllocationInitParameters struct {

	// The CIDR you want to assign to the pool.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// The description for the allocation.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Exclude a particular CIDR range from being returned by the pool.
	// +listType=set
	DisallowedCidrs []*string `json:"disallowedCidrs,omitempty" tf:"disallowed_cidrs,omitempty"`

	// The ID of the pool to which you want to assign a CIDR.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPCIpamPool
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	IpamPoolID *string `json:"ipamPoolId,omitempty" tf:"ipam_pool_id,omitempty"`

	// Reference to a VPCIpamPool in ec2 to populate ipamPoolId.
	// +kubebuilder:validation:Optional
	IpamPoolIDRef *v1.Reference `json:"ipamPoolIdRef,omitempty" tf:"-"`

	// Selector for a VPCIpamPool in ec2 to populate ipamPoolId.
	// +kubebuilder:validation:Optional
	IpamPoolIDSelector *v1.Selector `json:"ipamPoolIdSelector,omitempty" tf:"-"`

	// The netmask length of the CIDR you would like to allocate to the IPAM pool. Valid Values: 0-128.
	NetmaskLength *float64 `json:"netmaskLength,omitempty" tf:"netmask_length,omitempty"`
}

type VPCIpamPoolCidrAllocationObservation struct {

	// The CIDR you want to assign to the pool.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// The description for the allocation.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Exclude a particular CIDR range from being returned by the pool.
	// +listType=set
	DisallowedCidrs []*string `json:"disallowedCidrs,omitempty" tf:"disallowed_cidrs,omitempty"`

	// The ID of the allocation.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the allocation.
	IpamPoolAllocationID *string `json:"ipamPoolAllocationId,omitempty" tf:"ipam_pool_allocation_id,omitempty"`

	// The ID of the pool to which you want to assign a CIDR.
	IpamPoolID *string `json:"ipamPoolId,omitempty" tf:"ipam_pool_id,omitempty"`

	// The netmask length of the CIDR you would like to allocate to the IPAM pool. Valid Values: 0-128.
	NetmaskLength *float64 `json:"netmaskLength,omitempty" tf:"netmask_length,omitempty"`

	// The ID of the resource.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// The owner of the resource.
	ResourceOwner *string `json:"resourceOwner,omitempty" tf:"resource_owner,omitempty"`

	// The type of the resource.
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`
}

type VPCIpamPoolCidrAllocationParameters struct {

	// The CIDR you want to assign to the pool.
	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// The description for the allocation.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Exclude a particular CIDR range from being returned by the pool.
	// +kubebuilder:validation:Optional
	// +listType=set
	DisallowedCidrs []*string `json:"disallowedCidrs,omitempty" tf:"disallowed_cidrs,omitempty"`

	// The ID of the pool to which you want to assign a CIDR.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPCIpamPool
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	IpamPoolID *string `json:"ipamPoolId,omitempty" tf:"ipam_pool_id,omitempty"`

	// Reference to a VPCIpamPool in ec2 to populate ipamPoolId.
	// +kubebuilder:validation:Optional
	IpamPoolIDRef *v1.Reference `json:"ipamPoolIdRef,omitempty" tf:"-"`

	// Selector for a VPCIpamPool in ec2 to populate ipamPoolId.
	// +kubebuilder:validation:Optional
	IpamPoolIDSelector *v1.Selector `json:"ipamPoolIdSelector,omitempty" tf:"-"`

	// The netmask length of the CIDR you would like to allocate to the IPAM pool. Valid Values: 0-128.
	// +kubebuilder:validation:Optional
	NetmaskLength *float64 `json:"netmaskLength,omitempty" tf:"netmask_length,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// VPCIpamPoolCidrAllocationSpec defines the desired state of VPCIpamPoolCidrAllocation
type VPCIpamPoolCidrAllocationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCIpamPoolCidrAllocationParameters `json:"forProvider"`
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
	InitProvider VPCIpamPoolCidrAllocationInitParameters `json:"initProvider,omitempty"`
}

// VPCIpamPoolCidrAllocationStatus defines the observed state of VPCIpamPoolCidrAllocation.
type VPCIpamPoolCidrAllocationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCIpamPoolCidrAllocationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCIpamPoolCidrAllocation is the Schema for the VPCIpamPoolCidrAllocations API. Allocates (reserves) a CIDR from an IPAM address pool, preventing usage by IPAM.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCIpamPoolCidrAllocation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCIpamPoolCidrAllocationSpec   `json:"spec"`
	Status            VPCIpamPoolCidrAllocationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCIpamPoolCidrAllocationList contains a list of VPCIpamPoolCidrAllocations
type VPCIpamPoolCidrAllocationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCIpamPoolCidrAllocation `json:"items"`
}

// Repository type metadata.
var (
	VPCIpamPoolCidrAllocation_Kind             = "VPCIpamPoolCidrAllocation"
	VPCIpamPoolCidrAllocation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCIpamPoolCidrAllocation_Kind}.String()
	VPCIpamPoolCidrAllocation_KindAPIVersion   = VPCIpamPoolCidrAllocation_Kind + "." + CRDGroupVersion.String()
	VPCIpamPoolCidrAllocation_GroupVersionKind = CRDGroupVersion.WithKind(VPCIpamPoolCidrAllocation_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCIpamPoolCidrAllocation{}, &VPCIpamPoolCidrAllocationList{})
}

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

type VPCIpamPoolInitParameters struct {

	// The IP protocol assigned to this pool. You must choose either IPv4 or IPv6 protocol for a pool.
	AddressFamily *string `json:"addressFamily,omitempty" tf:"address_family,omitempty"`

	// A default netmask length for allocations added to this pool. If, for example, the CIDR assigned to this pool is 10.0.0.0/8 and you enter 16 here, new allocations will default to 10.0.0.0/16 (unless you provide a different netmask value when you create the new allocation).
	AllocationDefaultNetmaskLength *float64 `json:"allocationDefaultNetmaskLength,omitempty" tf:"allocation_default_netmask_length,omitempty"`

	// The maximum netmask length that will be required for CIDR allocations in this pool.
	AllocationMaxNetmaskLength *float64 `json:"allocationMaxNetmaskLength,omitempty" tf:"allocation_max_netmask_length,omitempty"`

	// The minimum netmask length that will be required for CIDR allocations in this pool.
	AllocationMinNetmaskLength *float64 `json:"allocationMinNetmaskLength,omitempty" tf:"allocation_min_netmask_length,omitempty"`

	// Tags that are required for resources that use CIDRs from this IPAM pool. Resources that do not have these tags will not be allowed to allocate space from the pool. If the resources have their tags changed after they have allocated space or if the allocation tagging requirements are changed on the pool, the resource may be marked as noncompliant.
	// +mapType=granular
	AllocationResourceTags map[string]*string `json:"allocationResourceTags,omitempty" tf:"allocation_resource_tags,omitempty"`

	// If you include this argument, IPAM automatically imports any VPCs you have in your scope that fall
	// within the CIDR range in the pool.
	AutoImport *bool `json:"autoImport,omitempty" tf:"auto_import,omitempty"`

	// Limits which AWS service the pool can be used in. Only useable on public scopes. Valid Values: ec2.
	AwsService *string `json:"awsService,omitempty" tf:"aws_service,omitempty"`

	// A description for the IPAM pool.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the scope in which you would like to create the IPAM pool.
	// +crossplane:generate:reference:type=VPCIpamScope
	IpamScopeID *string `json:"ipamScopeId,omitempty" tf:"ipam_scope_id,omitempty"`

	// Reference to a VPCIpamScope to populate ipamScopeId.
	// +kubebuilder:validation:Optional
	IpamScopeIDRef *v1.Reference `json:"ipamScopeIdRef,omitempty" tf:"-"`

	// Selector for a VPCIpamScope to populate ipamScopeId.
	// +kubebuilder:validation:Optional
	IpamScopeIDSelector *v1.Selector `json:"ipamScopeIdSelector,omitempty" tf:"-"`

	// The locale in which you would like to create the IPAM pool. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. Possible values: Any AWS region, such as us-east-1.
	Locale *string `json:"locale,omitempty" tf:"locale,omitempty"`

	// The IP address source for pools in the public scope. Only used for provisioning IP address CIDRs to pools in the public scope. Valid values are byoip or amazon. Default is byoip.
	PublicIPSource *string `json:"publicIpSource,omitempty" tf:"public_ip_source,omitempty"`

	// Defines whether or not IPv6 pool space is publicly advertisable over the internet. This argument is required if address_family = "ipv6" and public_ip_source = "byoip", default is false. This option is not available for IPv4 pool space or if public_ip_source = "amazon".
	PubliclyAdvertisable *bool `json:"publiclyAdvertisable,omitempty" tf:"publicly_advertisable,omitempty"`

	// The ID of the source IPAM pool. Use this argument to create a child pool within an existing pool.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPCIpamPool
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SourceIpamPoolID *string `json:"sourceIpamPoolId,omitempty" tf:"source_ipam_pool_id,omitempty"`

	// Reference to a VPCIpamPool in ec2 to populate sourceIpamPoolId.
	// +kubebuilder:validation:Optional
	SourceIpamPoolIDRef *v1.Reference `json:"sourceIpamPoolIdRef,omitempty" tf:"-"`

	// Selector for a VPCIpamPool in ec2 to populate sourceIpamPoolId.
	// +kubebuilder:validation:Optional
	SourceIpamPoolIDSelector *v1.Selector `json:"sourceIpamPoolIdSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VPCIpamPoolObservation struct {

	// The IP protocol assigned to this pool. You must choose either IPv4 or IPv6 protocol for a pool.
	AddressFamily *string `json:"addressFamily,omitempty" tf:"address_family,omitempty"`

	// A default netmask length for allocations added to this pool. If, for example, the CIDR assigned to this pool is 10.0.0.0/8 and you enter 16 here, new allocations will default to 10.0.0.0/16 (unless you provide a different netmask value when you create the new allocation).
	AllocationDefaultNetmaskLength *float64 `json:"allocationDefaultNetmaskLength,omitempty" tf:"allocation_default_netmask_length,omitempty"`

	// The maximum netmask length that will be required for CIDR allocations in this pool.
	AllocationMaxNetmaskLength *float64 `json:"allocationMaxNetmaskLength,omitempty" tf:"allocation_max_netmask_length,omitempty"`

	// The minimum netmask length that will be required for CIDR allocations in this pool.
	AllocationMinNetmaskLength *float64 `json:"allocationMinNetmaskLength,omitempty" tf:"allocation_min_netmask_length,omitempty"`

	// Tags that are required for resources that use CIDRs from this IPAM pool. Resources that do not have these tags will not be allowed to allocate space from the pool. If the resources have their tags changed after they have allocated space or if the allocation tagging requirements are changed on the pool, the resource may be marked as noncompliant.
	// +mapType=granular
	AllocationResourceTags map[string]*string `json:"allocationResourceTags,omitempty" tf:"allocation_resource_tags,omitempty"`

	// Amazon Resource Name (ARN) of IPAM
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// If you include this argument, IPAM automatically imports any VPCs you have in your scope that fall
	// within the CIDR range in the pool.
	AutoImport *bool `json:"autoImport,omitempty" tf:"auto_import,omitempty"`

	// Limits which AWS service the pool can be used in. Only useable on public scopes. Valid Values: ec2.
	AwsService *string `json:"awsService,omitempty" tf:"aws_service,omitempty"`

	// A description for the IPAM pool.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the IPAM
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the scope in which you would like to create the IPAM pool.
	IpamScopeID *string `json:"ipamScopeId,omitempty" tf:"ipam_scope_id,omitempty"`

	IpamScopeType *string `json:"ipamScopeType,omitempty" tf:"ipam_scope_type,omitempty"`

	// The locale in which you would like to create the IPAM pool. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. Possible values: Any AWS region, such as us-east-1.
	Locale *string `json:"locale,omitempty" tf:"locale,omitempty"`

	PoolDepth *float64 `json:"poolDepth,omitempty" tf:"pool_depth,omitempty"`

	// The IP address source for pools in the public scope. Only used for provisioning IP address CIDRs to pools in the public scope. Valid values are byoip or amazon. Default is byoip.
	PublicIPSource *string `json:"publicIpSource,omitempty" tf:"public_ip_source,omitempty"`

	// Defines whether or not IPv6 pool space is publicly advertisable over the internet. This argument is required if address_family = "ipv6" and public_ip_source = "byoip", default is false. This option is not available for IPv4 pool space or if public_ip_source = "amazon".
	PubliclyAdvertisable *bool `json:"publiclyAdvertisable,omitempty" tf:"publicly_advertisable,omitempty"`

	// The ID of the source IPAM pool. Use this argument to create a child pool within an existing pool.
	SourceIpamPoolID *string `json:"sourceIpamPoolId,omitempty" tf:"source_ipam_pool_id,omitempty"`

	// The ID of the IPAM
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type VPCIpamPoolParameters struct {

	// The IP protocol assigned to this pool. You must choose either IPv4 or IPv6 protocol for a pool.
	// +kubebuilder:validation:Optional
	AddressFamily *string `json:"addressFamily,omitempty" tf:"address_family,omitempty"`

	// A default netmask length for allocations added to this pool. If, for example, the CIDR assigned to this pool is 10.0.0.0/8 and you enter 16 here, new allocations will default to 10.0.0.0/16 (unless you provide a different netmask value when you create the new allocation).
	// +kubebuilder:validation:Optional
	AllocationDefaultNetmaskLength *float64 `json:"allocationDefaultNetmaskLength,omitempty" tf:"allocation_default_netmask_length,omitempty"`

	// The maximum netmask length that will be required for CIDR allocations in this pool.
	// +kubebuilder:validation:Optional
	AllocationMaxNetmaskLength *float64 `json:"allocationMaxNetmaskLength,omitempty" tf:"allocation_max_netmask_length,omitempty"`

	// The minimum netmask length that will be required for CIDR allocations in this pool.
	// +kubebuilder:validation:Optional
	AllocationMinNetmaskLength *float64 `json:"allocationMinNetmaskLength,omitempty" tf:"allocation_min_netmask_length,omitempty"`

	// Tags that are required for resources that use CIDRs from this IPAM pool. Resources that do not have these tags will not be allowed to allocate space from the pool. If the resources have their tags changed after they have allocated space or if the allocation tagging requirements are changed on the pool, the resource may be marked as noncompliant.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	AllocationResourceTags map[string]*string `json:"allocationResourceTags,omitempty" tf:"allocation_resource_tags,omitempty"`

	// If you include this argument, IPAM automatically imports any VPCs you have in your scope that fall
	// within the CIDR range in the pool.
	// +kubebuilder:validation:Optional
	AutoImport *bool `json:"autoImport,omitempty" tf:"auto_import,omitempty"`

	// Limits which AWS service the pool can be used in. Only useable on public scopes. Valid Values: ec2.
	// +kubebuilder:validation:Optional
	AwsService *string `json:"awsService,omitempty" tf:"aws_service,omitempty"`

	// A description for the IPAM pool.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the scope in which you would like to create the IPAM pool.
	// +crossplane:generate:reference:type=VPCIpamScope
	// +kubebuilder:validation:Optional
	IpamScopeID *string `json:"ipamScopeId,omitempty" tf:"ipam_scope_id,omitempty"`

	// Reference to a VPCIpamScope to populate ipamScopeId.
	// +kubebuilder:validation:Optional
	IpamScopeIDRef *v1.Reference `json:"ipamScopeIdRef,omitempty" tf:"-"`

	// Selector for a VPCIpamScope to populate ipamScopeId.
	// +kubebuilder:validation:Optional
	IpamScopeIDSelector *v1.Selector `json:"ipamScopeIdSelector,omitempty" tf:"-"`

	// The locale in which you would like to create the IPAM pool. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. Possible values: Any AWS region, such as us-east-1.
	// +kubebuilder:validation:Optional
	Locale *string `json:"locale,omitempty" tf:"locale,omitempty"`

	// The IP address source for pools in the public scope. Only used for provisioning IP address CIDRs to pools in the public scope. Valid values are byoip or amazon. Default is byoip.
	// +kubebuilder:validation:Optional
	PublicIPSource *string `json:"publicIpSource,omitempty" tf:"public_ip_source,omitempty"`

	// Defines whether or not IPv6 pool space is publicly advertisable over the internet. This argument is required if address_family = "ipv6" and public_ip_source = "byoip", default is false. This option is not available for IPv4 pool space or if public_ip_source = "amazon".
	// +kubebuilder:validation:Optional
	PubliclyAdvertisable *bool `json:"publiclyAdvertisable,omitempty" tf:"publicly_advertisable,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the source IPAM pool. Use this argument to create a child pool within an existing pool.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPCIpamPool
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SourceIpamPoolID *string `json:"sourceIpamPoolId,omitempty" tf:"source_ipam_pool_id,omitempty"`

	// Reference to a VPCIpamPool in ec2 to populate sourceIpamPoolId.
	// +kubebuilder:validation:Optional
	SourceIpamPoolIDRef *v1.Reference `json:"sourceIpamPoolIdRef,omitempty" tf:"-"`

	// Selector for a VPCIpamPool in ec2 to populate sourceIpamPoolId.
	// +kubebuilder:validation:Optional
	SourceIpamPoolIDSelector *v1.Selector `json:"sourceIpamPoolIdSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VPCIpamPoolSpec defines the desired state of VPCIpamPool
type VPCIpamPoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCIpamPoolParameters `json:"forProvider"`
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
	InitProvider VPCIpamPoolInitParameters `json:"initProvider,omitempty"`
}

// VPCIpamPoolStatus defines the observed state of VPCIpamPool.
type VPCIpamPoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCIpamPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// VPCIpamPool is the Schema for the VPCIpamPools API. Provides a IP address pool resource for IPAM.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCIpamPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.addressFamily) || (has(self.initProvider) && has(self.initProvider.addressFamily))",message="spec.forProvider.addressFamily is a required parameter"
	Spec   VPCIpamPoolSpec   `json:"spec"`
	Status VPCIpamPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCIpamPoolList contains a list of VPCIpamPools
type VPCIpamPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCIpamPool `json:"items"`
}

// Repository type metadata.
var (
	VPCIpamPool_Kind             = "VPCIpamPool"
	VPCIpamPool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCIpamPool_Kind}.String()
	VPCIpamPool_KindAPIVersion   = VPCIpamPool_Kind + "." + CRDGroupVersion.String()
	VPCIpamPool_GroupVersionKind = CRDGroupVersion.WithKind(VPCIpamPool_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCIpamPool{}, &VPCIpamPoolList{})
}

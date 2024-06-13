// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type InstanceConnectEndpointInitParameters struct {

	// Indicates whether your client's IP address is preserved as the source. Default: true.
	PreserveClientIP *bool `json:"preserveClientIp,omitempty" tf:"preserve_client_ip,omitempty"`

	// References to SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRefs []v1.Reference `json:"securityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// One or more security groups to associate with the endpoint. If you don't specify a security group, the default security group for the VPC will be associated with the endpoint.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupIDSelector
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The ID of the subnet in which to create the EC2 Instance Connect Endpoint.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type InstanceConnectEndpointObservation struct {

	// The Amazon Resource Name (ARN) of the EC2 Instance Connect Endpoint.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The Availability Zone of the EC2 Instance Connect Endpoint.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The DNS name of the EC2 Instance Connect Endpoint.
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// The DNS name of the EC2 Instance Connect FIPS Endpoint.
	FipsDNSName *string `json:"fipsDnsName,omitempty" tf:"fips_dns_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IDs of the ENIs that Amazon EC2 automatically created when creating the EC2 Instance Connect Endpoint.
	NetworkInterfaceIds []*string `json:"networkInterfaceIds,omitempty" tf:"network_interface_ids,omitempty"`

	// The ID of the AWS account that created the EC2 Instance Connect Endpoint.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// Indicates whether your client's IP address is preserved as the source. Default: true.
	PreserveClientIP *bool `json:"preserveClientIp,omitempty" tf:"preserve_client_ip,omitempty"`

	// One or more security groups to associate with the endpoint. If you don't specify a security group, the default security group for the VPC will be associated with the endpoint.
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The ID of the subnet in which to create the EC2 Instance Connect Endpoint.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The ID of the VPC in which the EC2 Instance Connect Endpoint was created.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type InstanceConnectEndpointParameters struct {

	// Indicates whether your client's IP address is preserved as the source. Default: true.
	// +kubebuilder:validation:Optional
	PreserveClientIP *bool `json:"preserveClientIp,omitempty" tf:"preserve_client_ip,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// References to SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRefs []v1.Reference `json:"securityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// One or more security groups to associate with the endpoint. If you don't specify a security group, the default security group for the VPC will be associated with the endpoint.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The ID of the subnet in which to create the EC2 Instance Connect Endpoint.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// InstanceConnectEndpointSpec defines the desired state of InstanceConnectEndpoint
type InstanceConnectEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceConnectEndpointParameters `json:"forProvider"`
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
	InitProvider InstanceConnectEndpointInitParameters `json:"initProvider,omitempty"`
}

// InstanceConnectEndpointStatus defines the observed state of InstanceConnectEndpoint.
type InstanceConnectEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceConnectEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// InstanceConnectEndpoint is the Schema for the InstanceConnectEndpoints API. Provides an EC2 Instance Connect Endpoint resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type InstanceConnectEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceConnectEndpointSpec   `json:"spec"`
	Status            InstanceConnectEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceConnectEndpointList contains a list of InstanceConnectEndpoints
type InstanceConnectEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceConnectEndpoint `json:"items"`
}

// Repository type metadata.
var (
	InstanceConnectEndpoint_Kind             = "InstanceConnectEndpoint"
	InstanceConnectEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceConnectEndpoint_Kind}.String()
	InstanceConnectEndpoint_KindAPIVersion   = InstanceConnectEndpoint_Kind + "." + CRDGroupVersion.String()
	InstanceConnectEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(InstanceConnectEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceConnectEndpoint{}, &InstanceConnectEndpointList{})
}

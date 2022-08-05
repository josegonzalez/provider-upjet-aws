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

type DeadLetterConfigObservation struct {
}

type DeadLetterConfigParameters struct {

	// +kubebuilder:validation:Required
	TargetArn *string `json:"targetArn" tf:"target_arn,omitempty"`
}

type EnvironmentObservation struct {
}

type EnvironmentParameters struct {

	// +kubebuilder:validation:Optional
	Variables map[string]*string `json:"variables,omitempty" tf:"variables,omitempty"`
}

type EphemeralStorageObservation struct {
}

type EphemeralStorageParameters struct {

	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`
}

type FileSystemConfigObservation struct {
}

type FileSystemConfigParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/efs/v1beta1.AccessPoint
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// +kubebuilder:validation:Optional
	ArnRef *v1.Reference `json:"arnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ArnSelector *v1.Selector `json:"arnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	LocalMountPath *string `json:"localMountPath" tf:"local_mount_path,omitempty"`
}

type FunctionObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InvokeArn *string `json:"invokeArn,omitempty" tf:"invoke_arn,omitempty"`

	LastModified *string `json:"lastModified,omitempty" tf:"last_modified,omitempty"`

	QualifiedArn *string `json:"qualifiedArn,omitempty" tf:"qualified_arn,omitempty"`

	SigningJobArn *string `json:"signingJobArn,omitempty" tf:"signing_job_arn,omitempty"`

	SigningProfileVersionArn *string `json:"signingProfileVersionArn,omitempty" tf:"signing_profile_version_arn,omitempty"`

	SourceCodeSize *float64 `json:"sourceCodeSize,omitempty" tf:"source_code_size,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	VPCConfig []VPCConfigObservation `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`

	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type FunctionParameters struct {

	// +kubebuilder:validation:Optional
	Architectures []*string `json:"architectures,omitempty" tf:"architectures,omitempty"`

	// +kubebuilder:validation:Optional
	CodeSigningConfigArn *string `json:"codeSigningConfigArn,omitempty" tf:"code_signing_config_arn,omitempty"`

	// +kubebuilder:validation:Optional
	DeadLetterConfig []DeadLetterConfigParameters `json:"deadLetterConfig,omitempty" tf:"dead_letter_config,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Environment []EnvironmentParameters `json:"environment,omitempty" tf:"environment,omitempty"`

	// +kubebuilder:validation:Optional
	EphemeralStorage []EphemeralStorageParameters `json:"ephemeralStorage,omitempty" tf:"ephemeral_storage,omitempty"`

	// +kubebuilder:validation:Optional
	FileSystemConfig []FileSystemConfigParameters `json:"fileSystemConfig,omitempty" tf:"file_system_config,omitempty"`

	// +kubebuilder:validation:Optional
	Handler *string `json:"handler,omitempty" tf:"handler,omitempty"`

	// +kubebuilder:validation:Optional
	ImageConfig []ImageConfigParameters `json:"imageConfig,omitempty" tf:"image_config,omitempty"`

	// +kubebuilder:validation:Optional
	ImageURI *string `json:"imageUri,omitempty" tf:"image_uri,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyArnRef *v1.Reference `json:"kmsKeyArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KMSKeyArnSelector *v1.Selector `json:"kmsKeyArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Layers []*string `json:"layers,omitempty" tf:"layers,omitempty"`

	// +kubebuilder:validation:Optional
	MemorySize *float64 `json:"memorySize,omitempty" tf:"memory_size,omitempty"`

	// +kubebuilder:validation:Optional
	PackageType *string `json:"packageType,omitempty" tf:"package_type,omitempty"`

	// +kubebuilder:validation:Optional
	Publish *bool `json:"publish,omitempty" tf:"publish,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ReservedConcurrentExecutions *float64 `json:"reservedConcurrentExecutions,omitempty" tf:"reserved_concurrent_executions,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	RoleRef *v1.Reference `json:"roleRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RoleSelector *v1.Selector `json:"roleSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Runtime *string `json:"runtime,omitempty" tf:"runtime,omitempty"`

	// +kubebuilder:validation:Optional
	S3Bucket *string `json:"s3Bucket,omitempty" tf:"s3_bucket,omitempty"`

	// +kubebuilder:validation:Optional
	S3Key *string `json:"s3Key,omitempty" tf:"s3_key,omitempty"`

	// +kubebuilder:validation:Optional
	S3ObjectVersion *string `json:"s3ObjectVersion,omitempty" tf:"s3_object_version,omitempty"`

	// +kubebuilder:validation:Optional
	SourceCodeHash *string `json:"sourceCodeHash,omitempty" tf:"source_code_hash,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// +kubebuilder:validation:Optional
	TracingConfig []TracingConfigParameters `json:"tracingConfig,omitempty" tf:"tracing_config,omitempty"`

	// +kubebuilder:validation:Optional
	VPCConfig []VPCConfigParameters `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`
}

type ImageConfigObservation struct {
}

type ImageConfigParameters struct {

	// +kubebuilder:validation:Optional
	Command []*string `json:"command,omitempty" tf:"command,omitempty"`

	// +kubebuilder:validation:Optional
	EntryPoint []*string `json:"entryPoint,omitempty" tf:"entry_point,omitempty"`

	// +kubebuilder:validation:Optional
	WorkingDirectory *string `json:"workingDirectory,omitempty" tf:"working_directory,omitempty"`
}

type TracingConfigObservation struct {
}

type TracingConfigParameters struct {

	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`
}

type VPCConfigObservation struct {
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type VPCConfigParameters struct {

	// +kubebuilder:validation:Required
	SecurityGroupIds []*string `json:"securityGroupIds" tf:"security_group_ids,omitempty"`

	// +kubebuilder:validation:Required
	SubnetIds []*string `json:"subnetIds" tf:"subnet_ids,omitempty"`
}

// FunctionSpec defines the desired state of Function
type FunctionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FunctionParameters `json:"forProvider"`
}

// FunctionStatus defines the observed state of Function.
type FunctionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FunctionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Function is the Schema for the Functions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Function struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionSpec   `json:"spec"`
	Status            FunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionList contains a list of Functions
type FunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Function `json:"items"`
}

// Repository type metadata.
var (
	Function_Kind             = "Function"
	Function_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Function_Kind}.String()
	Function_KindAPIVersion   = Function_Kind + "." + CRDGroupVersion.String()
	Function_GroupVersionKind = CRDGroupVersion.WithKind(Function_Kind)
)

func init() {
	SchemeBuilder.Register(&Function{}, &FunctionList{})
}

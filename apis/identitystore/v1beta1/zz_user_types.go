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

type AddressesInitParameters struct {

	// The country that this address is in.
	Country *string `json:"country,omitempty" tf:"country,omitempty"`

	// The name that is typically displayed when the address is shown for display.
	Formatted *string `json:"formatted,omitempty" tf:"formatted,omitempty"`

	// The address locality.
	Locality *string `json:"locality,omitempty" tf:"locality,omitempty"`

	// The postal code of the address.
	PostalCode *string `json:"postalCode,omitempty" tf:"postal_code,omitempty"`

	// When true, this is the primary address associated with the user.
	Primary *bool `json:"primary,omitempty" tf:"primary,omitempty"`

	// The street of the address.
	StreetAddress *string `json:"streetAddress,omitempty" tf:"street_address,omitempty"`

	// The type of address.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AddressesObservation struct {

	// The country that this address is in.
	Country *string `json:"country,omitempty" tf:"country,omitempty"`

	// The name that is typically displayed when the address is shown for display.
	Formatted *string `json:"formatted,omitempty" tf:"formatted,omitempty"`

	// The address locality.
	Locality *string `json:"locality,omitempty" tf:"locality,omitempty"`

	// The postal code of the address.
	PostalCode *string `json:"postalCode,omitempty" tf:"postal_code,omitempty"`

	// When true, this is the primary address associated with the user.
	Primary *bool `json:"primary,omitempty" tf:"primary,omitempty"`

	// The region of the address.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The street of the address.
	StreetAddress *string `json:"streetAddress,omitempty" tf:"street_address,omitempty"`

	// The type of address.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AddressesParameters struct {

	// The country that this address is in.
	// +kubebuilder:validation:Optional
	Country *string `json:"country,omitempty" tf:"country,omitempty"`

	// The name that is typically displayed when the address is shown for display.
	// +kubebuilder:validation:Optional
	Formatted *string `json:"formatted,omitempty" tf:"formatted,omitempty"`

	// The address locality.
	// +kubebuilder:validation:Optional
	Locality *string `json:"locality,omitempty" tf:"locality,omitempty"`

	// The postal code of the address.
	// +kubebuilder:validation:Optional
	PostalCode *string `json:"postalCode,omitempty" tf:"postal_code,omitempty"`

	// When true, this is the primary address associated with the user.
	// +kubebuilder:validation:Optional
	Primary *bool `json:"primary,omitempty" tf:"primary,omitempty"`

	// The region of the address.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The street of the address.
	// +kubebuilder:validation:Optional
	StreetAddress *string `json:"streetAddress,omitempty" tf:"street_address,omitempty"`

	// The type of address.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type EmailsInitParameters struct {

	// When true, this is the primary email associated with the user.
	Primary *bool `json:"primary,omitempty" tf:"primary,omitempty"`

	// The type of email.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The email address. This value must be unique across the identity store.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type EmailsObservation struct {

	// When true, this is the primary email associated with the user.
	Primary *bool `json:"primary,omitempty" tf:"primary,omitempty"`

	// The type of email.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The email address. This value must be unique across the identity store.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type EmailsParameters struct {

	// When true, this is the primary email associated with the user.
	// +kubebuilder:validation:Optional
	Primary *bool `json:"primary,omitempty" tf:"primary,omitempty"`

	// The type of email.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The email address. This value must be unique across the identity store.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type NameInitParameters struct {

	// The family name of the user.
	FamilyName *string `json:"familyName,omitempty" tf:"family_name,omitempty"`

	// The name that is typically displayed when the name is shown for display.
	Formatted *string `json:"formatted,omitempty" tf:"formatted,omitempty"`

	// The given name of the user.
	GivenName *string `json:"givenName,omitempty" tf:"given_name,omitempty"`

	// The honorific prefix of the user.
	HonorificPrefix *string `json:"honorificPrefix,omitempty" tf:"honorific_prefix,omitempty"`

	// The honorific suffix of the user.
	HonorificSuffix *string `json:"honorificSuffix,omitempty" tf:"honorific_suffix,omitempty"`

	// The middle name of the user.
	MiddleName *string `json:"middleName,omitempty" tf:"middle_name,omitempty"`
}

type NameObservation struct {

	// The family name of the user.
	FamilyName *string `json:"familyName,omitempty" tf:"family_name,omitempty"`

	// The name that is typically displayed when the name is shown for display.
	Formatted *string `json:"formatted,omitempty" tf:"formatted,omitempty"`

	// The given name of the user.
	GivenName *string `json:"givenName,omitempty" tf:"given_name,omitempty"`

	// The honorific prefix of the user.
	HonorificPrefix *string `json:"honorificPrefix,omitempty" tf:"honorific_prefix,omitempty"`

	// The honorific suffix of the user.
	HonorificSuffix *string `json:"honorificSuffix,omitempty" tf:"honorific_suffix,omitempty"`

	// The middle name of the user.
	MiddleName *string `json:"middleName,omitempty" tf:"middle_name,omitempty"`
}

type NameParameters struct {

	// The family name of the user.
	// +kubebuilder:validation:Optional
	FamilyName *string `json:"familyName" tf:"family_name,omitempty"`

	// The name that is typically displayed when the name is shown for display.
	// +kubebuilder:validation:Optional
	Formatted *string `json:"formatted,omitempty" tf:"formatted,omitempty"`

	// The given name of the user.
	// +kubebuilder:validation:Optional
	GivenName *string `json:"givenName" tf:"given_name,omitempty"`

	// The honorific prefix of the user.
	// +kubebuilder:validation:Optional
	HonorificPrefix *string `json:"honorificPrefix,omitempty" tf:"honorific_prefix,omitempty"`

	// The honorific suffix of the user.
	// +kubebuilder:validation:Optional
	HonorificSuffix *string `json:"honorificSuffix,omitempty" tf:"honorific_suffix,omitempty"`

	// The middle name of the user.
	// +kubebuilder:validation:Optional
	MiddleName *string `json:"middleName,omitempty" tf:"middle_name,omitempty"`
}

type PhoneNumbersInitParameters struct {

	// When true, this is the primary phone number associated with the user.
	Primary *bool `json:"primary,omitempty" tf:"primary,omitempty"`

	// The type of phone number.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The user's phone number.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PhoneNumbersObservation struct {

	// When true, this is the primary phone number associated with the user.
	Primary *bool `json:"primary,omitempty" tf:"primary,omitempty"`

	// The type of phone number.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The user's phone number.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PhoneNumbersParameters struct {

	// When true, this is the primary phone number associated with the user.
	// +kubebuilder:validation:Optional
	Primary *bool `json:"primary,omitempty" tf:"primary,omitempty"`

	// The type of phone number.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The user's phone number.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type UserExternalIdsInitParameters struct {
}

type UserExternalIdsObservation struct {

	// The identifier issued to this resource by an external identity provider.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The issuer for an external identifier.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`
}

type UserExternalIdsParameters struct {
}

type UserInitParameters struct {

	// Details about the user's address. At most 1 address is allowed. Detailed below.
	Addresses []AddressesInitParameters `json:"addresses,omitempty" tf:"addresses,omitempty"`

	// The name that is typically displayed when the user is referenced.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Details about the user's email. At most 1 email is allowed. Detailed below.
	Emails []EmailsInitParameters `json:"emails,omitempty" tf:"emails,omitempty"`

	// The user's geographical region or location.
	Locale *string `json:"locale,omitempty" tf:"locale,omitempty"`

	// Details about the user's full name. Detailed below.
	Name []NameInitParameters `json:"name,omitempty" tf:"name,omitempty"`

	// An alternate name for the user.
	Nickname *string `json:"nickname,omitempty" tf:"nickname,omitempty"`

	// Details about the user's phone number. At most 1 phone number is allowed. Detailed below.
	PhoneNumbers []PhoneNumbersInitParameters `json:"phoneNumbers,omitempty" tf:"phone_numbers,omitempty"`

	// The preferred language of the user.
	PreferredLanguage *string `json:"preferredLanguage,omitempty" tf:"preferred_language,omitempty"`

	// An URL that may be associated with the user.
	ProfileURL *string `json:"profileUrl,omitempty" tf:"profile_url,omitempty"`

	// The user's time zone.
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`

	// The user's title.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`

	// A unique string used to identify the user. This value can consist of letters, accented characters, symbols, numbers, and punctuation. This value is specified at the time the user is created and stored as an attribute of the user object in the identity store. The limit is 128 characters.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`

	// The user type.
	UserType *string `json:"userType,omitempty" tf:"user_type,omitempty"`
}

type UserObservation struct {

	// Details about the user's address. At most 1 address is allowed. Detailed below.
	Addresses []AddressesObservation `json:"addresses,omitempty" tf:"addresses,omitempty"`

	// The name that is typically displayed when the user is referenced.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Details about the user's email. At most 1 email is allowed. Detailed below.
	Emails []EmailsObservation `json:"emails,omitempty" tf:"emails,omitempty"`

	// A list of identifiers issued to this resource by an external identity provider.
	ExternalIds []UserExternalIdsObservation `json:"externalIds,omitempty" tf:"external_ids,omitempty"`

	// The identifier issued to this resource by an external identity provider.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The globally unique identifier for the identity store that this user is in.
	IdentityStoreID *string `json:"identityStoreId,omitempty" tf:"identity_store_id,omitempty"`

	// The user's geographical region or location.
	Locale *string `json:"locale,omitempty" tf:"locale,omitempty"`

	// Details about the user's full name. Detailed below.
	Name []NameObservation `json:"name,omitempty" tf:"name,omitempty"`

	// An alternate name for the user.
	Nickname *string `json:"nickname,omitempty" tf:"nickname,omitempty"`

	// Details about the user's phone number. At most 1 phone number is allowed. Detailed below.
	PhoneNumbers []PhoneNumbersObservation `json:"phoneNumbers,omitempty" tf:"phone_numbers,omitempty"`

	// The preferred language of the user.
	PreferredLanguage *string `json:"preferredLanguage,omitempty" tf:"preferred_language,omitempty"`

	// An URL that may be associated with the user.
	ProfileURL *string `json:"profileUrl,omitempty" tf:"profile_url,omitempty"`

	// The user's time zone.
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`

	// The user's title.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`

	// The identifier for this user in the identity store.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// A unique string used to identify the user. This value can consist of letters, accented characters, symbols, numbers, and punctuation. This value is specified at the time the user is created and stored as an attribute of the user object in the identity store. The limit is 128 characters.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`

	// The user type.
	UserType *string `json:"userType,omitempty" tf:"user_type,omitempty"`
}

type UserParameters struct {

	// Details about the user's address. At most 1 address is allowed. Detailed below.
	// +kubebuilder:validation:Optional
	Addresses []AddressesParameters `json:"addresses,omitempty" tf:"addresses,omitempty"`

	// The name that is typically displayed when the user is referenced.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Details about the user's email. At most 1 email is allowed. Detailed below.
	// +kubebuilder:validation:Optional
	Emails []EmailsParameters `json:"emails,omitempty" tf:"emails,omitempty"`

	// The globally unique identifier for the identity store that this user is in.
	// +kubebuilder:validation:Required
	IdentityStoreID *string `json:"identityStoreId" tf:"identity_store_id,omitempty"`

	// The user's geographical region or location.
	// +kubebuilder:validation:Optional
	Locale *string `json:"locale,omitempty" tf:"locale,omitempty"`

	// Details about the user's full name. Detailed below.
	// +kubebuilder:validation:Optional
	Name []NameParameters `json:"name,omitempty" tf:"name,omitempty"`

	// An alternate name for the user.
	// +kubebuilder:validation:Optional
	Nickname *string `json:"nickname,omitempty" tf:"nickname,omitempty"`

	// Details about the user's phone number. At most 1 phone number is allowed. Detailed below.
	// +kubebuilder:validation:Optional
	PhoneNumbers []PhoneNumbersParameters `json:"phoneNumbers,omitempty" tf:"phone_numbers,omitempty"`

	// The preferred language of the user.
	// +kubebuilder:validation:Optional
	PreferredLanguage *string `json:"preferredLanguage,omitempty" tf:"preferred_language,omitempty"`

	// An URL that may be associated with the user.
	// +kubebuilder:validation:Optional
	ProfileURL *string `json:"profileUrl,omitempty" tf:"profile_url,omitempty"`

	// The region of the address.
	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The user's time zone.
	// +kubebuilder:validation:Optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`

	// The user's title.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`

	// A unique string used to identify the user. This value can consist of letters, accented characters, symbols, numbers, and punctuation. This value is specified at the time the user is created and stored as an attribute of the user object in the identity store. The limit is 128 characters.
	// +kubebuilder:validation:Optional
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`

	// The user type.
	// +kubebuilder:validation:Optional
	UserType *string `json:"userType,omitempty" tf:"user_type,omitempty"`
}

// UserSpec defines the desired state of User
type UserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserParameters `json:"forProvider"`
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
	InitProvider UserInitParameters `json:"initProvider,omitempty"`
}

// UserStatus defines the observed state of User.
type UserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// User is the Schema for the Users API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.userName) || (has(self.initProvider) && has(self.initProvider.userName))",message="spec.forProvider.userName is a required parameter"
	Spec   UserSpec   `json:"spec"`
	Status UserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserList contains a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// Repository type metadata.
var (
	User_Kind             = "User"
	User_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: User_Kind}.String()
	User_KindAPIVersion   = User_Kind + "." + CRDGroupVersion.String()
	User_GroupVersionKind = CRDGroupVersion.WithKind(User_Kind)
)

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}

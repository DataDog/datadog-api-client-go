/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// SyntheticsPrivateLocationMetadata Object containing metadata about the private location.
type SyntheticsPrivateLocationMetadata struct {
	// A list of role identifiers that can be pulled from the Roles API, for restricting read and write access.
	RestrictedRoles []string `json:"restricted_roles,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewSyntheticsPrivateLocationMetadata instantiates a new SyntheticsPrivateLocationMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsPrivateLocationMetadata() *SyntheticsPrivateLocationMetadata {
	this := SyntheticsPrivateLocationMetadata{}
	return &this
}

// NewSyntheticsPrivateLocationMetadataWithDefaults instantiates a new SyntheticsPrivateLocationMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsPrivateLocationMetadataWithDefaults() *SyntheticsPrivateLocationMetadata {
	this := SyntheticsPrivateLocationMetadata{}
	return &this
}

// GetRestrictedRoles returns the RestrictedRoles field value if set, zero value otherwise.
func (o *SyntheticsPrivateLocationMetadata) GetRestrictedRoles() []string {
	if o == nil || o.RestrictedRoles == nil {
		var ret []string
		return ret
	}
	return o.RestrictedRoles
}

// GetRestrictedRolesOk returns a tuple with the RestrictedRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsPrivateLocationMetadata) GetRestrictedRolesOk() (*[]string, bool) {
	if o == nil || o.RestrictedRoles == nil {
		return nil, false
	}
	return &o.RestrictedRoles, true
}

// HasRestrictedRoles returns a boolean if a field has been set.
func (o *SyntheticsPrivateLocationMetadata) HasRestrictedRoles() bool {
	if o != nil && o.RestrictedRoles != nil {
		return true
	}

	return false
}

// SetRestrictedRoles gets a reference to the given []string and assigns it to the RestrictedRoles field.
func (o *SyntheticsPrivateLocationMetadata) SetRestrictedRoles(v []string) {
	o.RestrictedRoles = v
}

func (o SyntheticsPrivateLocationMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.RestrictedRoles != nil {
		toSerialize["restricted_roles"] = o.RestrictedRoles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *SyntheticsPrivateLocationMetadata) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		RestrictedRoles []string `json:"restricted_roles,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.RestrictedRoles = all.RestrictedRoles
	return nil
}

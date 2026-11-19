// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TerraformCloudAuthPersonaMappingCreateData Data for creating an HCP Terraform persona mapping.
type TerraformCloudAuthPersonaMappingCreateData struct {
	// Attributes for creating an HCP Terraform persona mapping.
	Attributes TerraformCloudAuthPersonaMappingCreateAttributes `json:"attributes"`
	// Type identifier for an HCP Terraform persona mapping.
	Type TerraformCloudAuthPersonaMappingType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTerraformCloudAuthPersonaMappingCreateData instantiates a new TerraformCloudAuthPersonaMappingCreateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTerraformCloudAuthPersonaMappingCreateData(attributes TerraformCloudAuthPersonaMappingCreateAttributes, typeVar TerraformCloudAuthPersonaMappingType) *TerraformCloudAuthPersonaMappingCreateData {
	this := TerraformCloudAuthPersonaMappingCreateData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewTerraformCloudAuthPersonaMappingCreateDataWithDefaults instantiates a new TerraformCloudAuthPersonaMappingCreateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTerraformCloudAuthPersonaMappingCreateDataWithDefaults() *TerraformCloudAuthPersonaMappingCreateData {
	this := TerraformCloudAuthPersonaMappingCreateData{}
	var typeVar TerraformCloudAuthPersonaMappingType = TERRAFORMCLOUDAUTHPERSONAMAPPINGTYPE_TERRAFORM_CLOUD_AUTH_CONFIG
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *TerraformCloudAuthPersonaMappingCreateData) GetAttributes() TerraformCloudAuthPersonaMappingCreateAttributes {
	if o == nil {
		var ret TerraformCloudAuthPersonaMappingCreateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TerraformCloudAuthPersonaMappingCreateData) GetAttributesOk() (*TerraformCloudAuthPersonaMappingCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *TerraformCloudAuthPersonaMappingCreateData) SetAttributes(v TerraformCloudAuthPersonaMappingCreateAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *TerraformCloudAuthPersonaMappingCreateData) GetType() TerraformCloudAuthPersonaMappingType {
	if o == nil {
		var ret TerraformCloudAuthPersonaMappingType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TerraformCloudAuthPersonaMappingCreateData) GetTypeOk() (*TerraformCloudAuthPersonaMappingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TerraformCloudAuthPersonaMappingCreateData) SetType(v TerraformCloudAuthPersonaMappingType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TerraformCloudAuthPersonaMappingCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TerraformCloudAuthPersonaMappingCreateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *TerraformCloudAuthPersonaMappingCreateAttributes `json:"attributes"`
		Type       *TerraformCloudAuthPersonaMappingType             `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TerraformCloudAuthPersonaMappingDataResponse HCP Terraform persona mapping resource.
type TerraformCloudAuthPersonaMappingDataResponse struct {
	// Stored HCP Terraform persona mapping attributes.
	Attributes TerraformCloudAuthPersonaMappingAttributesResponse `json:"attributes"`
	// Unique identifier of the persona mapping.
	Id string `json:"id"`
	// Type identifier for an HCP Terraform persona mapping.
	Type TerraformCloudAuthPersonaMappingType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTerraformCloudAuthPersonaMappingDataResponse instantiates a new TerraformCloudAuthPersonaMappingDataResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTerraformCloudAuthPersonaMappingDataResponse(attributes TerraformCloudAuthPersonaMappingAttributesResponse, id string, typeVar TerraformCloudAuthPersonaMappingType) *TerraformCloudAuthPersonaMappingDataResponse {
	this := TerraformCloudAuthPersonaMappingDataResponse{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewTerraformCloudAuthPersonaMappingDataResponseWithDefaults instantiates a new TerraformCloudAuthPersonaMappingDataResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTerraformCloudAuthPersonaMappingDataResponseWithDefaults() *TerraformCloudAuthPersonaMappingDataResponse {
	this := TerraformCloudAuthPersonaMappingDataResponse{}
	var typeVar TerraformCloudAuthPersonaMappingType = TERRAFORMCLOUDAUTHPERSONAMAPPINGTYPE_TERRAFORM_CLOUD_AUTH_CONFIG
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *TerraformCloudAuthPersonaMappingDataResponse) GetAttributes() TerraformCloudAuthPersonaMappingAttributesResponse {
	if o == nil {
		var ret TerraformCloudAuthPersonaMappingAttributesResponse
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TerraformCloudAuthPersonaMappingDataResponse) GetAttributesOk() (*TerraformCloudAuthPersonaMappingAttributesResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *TerraformCloudAuthPersonaMappingDataResponse) SetAttributes(v TerraformCloudAuthPersonaMappingAttributesResponse) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *TerraformCloudAuthPersonaMappingDataResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TerraformCloudAuthPersonaMappingDataResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *TerraformCloudAuthPersonaMappingDataResponse) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *TerraformCloudAuthPersonaMappingDataResponse) GetType() TerraformCloudAuthPersonaMappingType {
	if o == nil {
		var ret TerraformCloudAuthPersonaMappingType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TerraformCloudAuthPersonaMappingDataResponse) GetTypeOk() (*TerraformCloudAuthPersonaMappingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TerraformCloudAuthPersonaMappingDataResponse) SetType(v TerraformCloudAuthPersonaMappingType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TerraformCloudAuthPersonaMappingDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TerraformCloudAuthPersonaMappingDataResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *TerraformCloudAuthPersonaMappingAttributesResponse `json:"attributes"`
		Id         *string                                             `json:"id"`
		Type       *TerraformCloudAuthPersonaMappingType               `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

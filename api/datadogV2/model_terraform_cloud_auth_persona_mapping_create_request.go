// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TerraformCloudAuthPersonaMappingCreateRequest Request to create an HCP Terraform persona mapping.
type TerraformCloudAuthPersonaMappingCreateRequest struct {
	// Data for creating an HCP Terraform persona mapping.
	Data TerraformCloudAuthPersonaMappingCreateData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTerraformCloudAuthPersonaMappingCreateRequest instantiates a new TerraformCloudAuthPersonaMappingCreateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTerraformCloudAuthPersonaMappingCreateRequest(data TerraformCloudAuthPersonaMappingCreateData) *TerraformCloudAuthPersonaMappingCreateRequest {
	this := TerraformCloudAuthPersonaMappingCreateRequest{}
	this.Data = data
	return &this
}

// NewTerraformCloudAuthPersonaMappingCreateRequestWithDefaults instantiates a new TerraformCloudAuthPersonaMappingCreateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTerraformCloudAuthPersonaMappingCreateRequestWithDefaults() *TerraformCloudAuthPersonaMappingCreateRequest {
	this := TerraformCloudAuthPersonaMappingCreateRequest{}
	return &this
}

// GetData returns the Data field value.
func (o *TerraformCloudAuthPersonaMappingCreateRequest) GetData() TerraformCloudAuthPersonaMappingCreateData {
	if o == nil {
		var ret TerraformCloudAuthPersonaMappingCreateData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TerraformCloudAuthPersonaMappingCreateRequest) GetDataOk() (*TerraformCloudAuthPersonaMappingCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *TerraformCloudAuthPersonaMappingCreateRequest) SetData(v TerraformCloudAuthPersonaMappingCreateData) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TerraformCloudAuthPersonaMappingCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TerraformCloudAuthPersonaMappingCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *TerraformCloudAuthPersonaMappingCreateData `json:"data"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}

	hasInvalidField := false
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = *all.Data

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

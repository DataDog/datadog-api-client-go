// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostTagPipelineActiveKeyResponse List of tag keys actively set by tag pipeline rules.
type CostTagPipelineActiveKeyResponse struct {
	// List of active tag keys.
	Data []CostTagPipelineActiveKeyData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostTagPipelineActiveKeyResponse instantiates a new CostTagPipelineActiveKeyResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostTagPipelineActiveKeyResponse(data []CostTagPipelineActiveKeyData) *CostTagPipelineActiveKeyResponse {
	this := CostTagPipelineActiveKeyResponse{}
	this.Data = data
	return &this
}

// NewCostTagPipelineActiveKeyResponseWithDefaults instantiates a new CostTagPipelineActiveKeyResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostTagPipelineActiveKeyResponseWithDefaults() *CostTagPipelineActiveKeyResponse {
	this := CostTagPipelineActiveKeyResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *CostTagPipelineActiveKeyResponse) GetData() []CostTagPipelineActiveKeyData {
	if o == nil {
		var ret []CostTagPipelineActiveKeyData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CostTagPipelineActiveKeyResponse) GetDataOk() (*[]CostTagPipelineActiveKeyData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *CostTagPipelineActiveKeyResponse) SetData(v []CostTagPipelineActiveKeyData) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostTagPipelineActiveKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostTagPipelineActiveKeyResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]CostTagPipelineActiveKeyData `json:"data"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}
	o.Data = *all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExecutionPolicyListResponse Response object that includes a list of execution policies.
type ExecutionPolicyListResponse struct {
	// The execution policies.
	Data []ExecutionPolicyResponseData `json:"data"`
	// Pagination metadata for the list of execution policies.
	Meta ExecutionPolicyListResponseMeta `json:"meta"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExecutionPolicyListResponse instantiates a new ExecutionPolicyListResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExecutionPolicyListResponse(data []ExecutionPolicyResponseData, meta ExecutionPolicyListResponseMeta) *ExecutionPolicyListResponse {
	this := ExecutionPolicyListResponse{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewExecutionPolicyListResponseWithDefaults instantiates a new ExecutionPolicyListResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExecutionPolicyListResponseWithDefaults() *ExecutionPolicyListResponse {
	this := ExecutionPolicyListResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *ExecutionPolicyListResponse) GetData() []ExecutionPolicyResponseData {
	if o == nil {
		var ret []ExecutionPolicyResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyListResponse) GetDataOk() (*[]ExecutionPolicyResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ExecutionPolicyListResponse) SetData(v []ExecutionPolicyResponseData) {
	o.Data = v
}

// GetMeta returns the Meta field value.
func (o *ExecutionPolicyListResponse) GetMeta() ExecutionPolicyListResponseMeta {
	if o == nil {
		var ret ExecutionPolicyListResponseMeta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyListResponse) GetMetaOk() (*ExecutionPolicyListResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value.
func (o *ExecutionPolicyListResponse) SetMeta(v ExecutionPolicyListResponseMeta) {
	o.Meta = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExecutionPolicyListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	toSerialize["meta"] = o.Meta

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExecutionPolicyListResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]ExecutionPolicyResponseData   `json:"data"`
		Meta *ExecutionPolicyListResponseMeta `json:"meta"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	if all.Meta == nil {
		return fmt.Errorf("required field meta missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	if all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = *all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

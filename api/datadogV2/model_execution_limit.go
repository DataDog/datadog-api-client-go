// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExecutionLimit The maximum number of times to execute a workflow for an incident.
type ExecutionLimit struct {
	// The maximum number of workflow executions.
	Count int32 `json:"count"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExecutionLimit instantiates a new ExecutionLimit object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExecutionLimit(count int32) *ExecutionLimit {
	this := ExecutionLimit{}
	this.Count = count
	return &this
}

// NewExecutionLimitWithDefaults instantiates a new ExecutionLimit object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExecutionLimitWithDefaults() *ExecutionLimit {
	this := ExecutionLimit{}
	return &this
}

// GetCount returns the Count field value.
func (o *ExecutionLimit) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ExecutionLimit) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *ExecutionLimit) SetCount(v int32) {
	o.Count = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExecutionLimit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["count"] = o.Count

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExecutionLimit) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Count *int32 `json:"count"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Count == nil {
		return fmt.Errorf("required field count missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"count"})
	} else {
		return err
	}
	o.Count = *all.Count

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListExecutionStepsAttributes Attributes of an execution steps response.
type ListExecutionStepsAttributes struct {
	// The list of steps for the workflow execution.
	Steps []ExecutionStep `json:"steps"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListExecutionStepsAttributes instantiates a new ListExecutionStepsAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListExecutionStepsAttributes(steps []ExecutionStep) *ListExecutionStepsAttributes {
	this := ListExecutionStepsAttributes{}
	this.Steps = steps
	return &this
}

// NewListExecutionStepsAttributesWithDefaults instantiates a new ListExecutionStepsAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListExecutionStepsAttributesWithDefaults() *ListExecutionStepsAttributes {
	this := ListExecutionStepsAttributes{}
	return &this
}

// GetSteps returns the Steps field value.
func (o *ListExecutionStepsAttributes) GetSteps() []ExecutionStep {
	if o == nil {
		var ret []ExecutionStep
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value
// and a boolean to check if the value has been set.
func (o *ListExecutionStepsAttributes) GetStepsOk() (*[]ExecutionStep, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Steps, true
}

// SetSteps sets field value.
func (o *ListExecutionStepsAttributes) SetSteps(v []ExecutionStep) {
	o.Steps = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListExecutionStepsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["steps"] = o.Steps

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListExecutionStepsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Steps *[]ExecutionStep `json:"steps"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Steps == nil {
		return fmt.Errorf("required field steps missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"steps"})
	} else {
		return err
	}
	o.Steps = *all.Steps

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

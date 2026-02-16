// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DataTransformationContext
type DataTransformationContext struct {
	// Available context variables for the transformation.
	ContextVariables string `json:"contextVariables"`
	// The current script to modify or enhance.
	CurrentScript string `json:"currentScript"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataTransformationContext instantiates a new DataTransformationContext object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataTransformationContext(contextVariables string, currentScript string) *DataTransformationContext {
	this := DataTransformationContext{}
	this.ContextVariables = contextVariables
	this.CurrentScript = currentScript
	return &this
}

// NewDataTransformationContextWithDefaults instantiates a new DataTransformationContext object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataTransformationContextWithDefaults() *DataTransformationContext {
	this := DataTransformationContext{}
	return &this
}

// GetContextVariables returns the ContextVariables field value.
func (o *DataTransformationContext) GetContextVariables() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ContextVariables
}

// GetContextVariablesOk returns a tuple with the ContextVariables field value
// and a boolean to check if the value has been set.
func (o *DataTransformationContext) GetContextVariablesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContextVariables, true
}

// SetContextVariables sets field value.
func (o *DataTransformationContext) SetContextVariables(v string) {
	o.ContextVariables = v
}

// GetCurrentScript returns the CurrentScript field value.
func (o *DataTransformationContext) GetCurrentScript() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CurrentScript
}

// GetCurrentScriptOk returns a tuple with the CurrentScript field value
// and a boolean to check if the value has been set.
func (o *DataTransformationContext) GetCurrentScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentScript, true
}

// SetCurrentScript sets field value.
func (o *DataTransformationContext) SetCurrentScript(v string) {
	o.CurrentScript = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataTransformationContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["contextVariables"] = o.ContextVariables
	toSerialize["currentScript"] = o.CurrentScript

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataTransformationContext) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ContextVariables *string `json:"contextVariables"`
		CurrentScript    *string `json:"currentScript"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ContextVariables == nil {
		return fmt.Errorf("required field contextVariables missing")
	}
	if all.CurrentScript == nil {
		return fmt.Errorf("required field currentScript missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"contextVariables", "currentScript"})
	} else {
		return err
	}
	o.ContextVariables = *all.ContextVariables
	o.CurrentScript = *all.CurrentScript

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

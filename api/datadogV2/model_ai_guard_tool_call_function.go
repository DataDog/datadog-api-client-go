// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIGuardToolCallFunction The function definition within a tool call.
type AIGuardToolCallFunction struct {
	// The JSON-encoded arguments passed to the function.
	Arguments string `json:"arguments"`
	// The name of the function being called.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAIGuardToolCallFunction instantiates a new AIGuardToolCallFunction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAIGuardToolCallFunction(arguments string, name string) *AIGuardToolCallFunction {
	this := AIGuardToolCallFunction{}
	this.Arguments = arguments
	this.Name = name
	return &this
}

// NewAIGuardToolCallFunctionWithDefaults instantiates a new AIGuardToolCallFunction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAIGuardToolCallFunctionWithDefaults() *AIGuardToolCallFunction {
	this := AIGuardToolCallFunction{}
	return &this
}

// GetArguments returns the Arguments field value.
func (o *AIGuardToolCallFunction) GetArguments() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value
// and a boolean to check if the value has been set.
func (o *AIGuardToolCallFunction) GetArgumentsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Arguments, true
}

// SetArguments sets field value.
func (o *AIGuardToolCallFunction) SetArguments(v string) {
	o.Arguments = v
}

// GetName returns the Name field value.
func (o *AIGuardToolCallFunction) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AIGuardToolCallFunction) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AIGuardToolCallFunction) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AIGuardToolCallFunction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["arguments"] = o.Arguments
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AIGuardToolCallFunction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Arguments *string `json:"arguments"`
		Name      *string `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Arguments == nil {
		return fmt.Errorf("required field arguments missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"arguments", "name"})
	} else {
		return err
	}
	o.Arguments = *all.Arguments
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

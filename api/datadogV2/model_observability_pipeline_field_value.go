// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineFieldValue Represents a static key-value pair used in various processors.
type ObservabilityPipelineFieldValue struct {
	// The field name.
	Name string `json:"name"`
	// The field value.
	Value string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineFieldValue instantiates a new ObservabilityPipelineFieldValue object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineFieldValue(name string, value string) *ObservabilityPipelineFieldValue {
	this := ObservabilityPipelineFieldValue{}
	this.Name = name
	this.Value = value
	return &this
}

// NewObservabilityPipelineFieldValueWithDefaults instantiates a new ObservabilityPipelineFieldValue object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineFieldValueWithDefaults() *ObservabilityPipelineFieldValue {
	this := ObservabilityPipelineFieldValue{}
	return &this
}

// GetName returns the Name field value.
func (o *ObservabilityPipelineFieldValue) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineFieldValue) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ObservabilityPipelineFieldValue) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value.
func (o *ObservabilityPipelineFieldValue) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineFieldValue) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *ObservabilityPipelineFieldValue) SetValue(v string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineFieldValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineFieldValue) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name  *string `json:"name"`
		Value *string `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "value"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

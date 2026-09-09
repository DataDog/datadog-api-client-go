// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EmbeddedAppWidgetInput An input passed to the embedded app.
type EmbeddedAppWidgetInput struct {
	// Name of the app input.
	Name string `json:"name"`
	// Value of the app input. This can be a string, number, boolean, object, or a non-empty homogeneous array of those types.
	Value EmbeddedAppWidgetInputValue `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEmbeddedAppWidgetInput instantiates a new EmbeddedAppWidgetInput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEmbeddedAppWidgetInput(name string, value EmbeddedAppWidgetInputValue) *EmbeddedAppWidgetInput {
	this := EmbeddedAppWidgetInput{}
	this.Name = name
	this.Value = value
	return &this
}

// NewEmbeddedAppWidgetInputWithDefaults instantiates a new EmbeddedAppWidgetInput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEmbeddedAppWidgetInputWithDefaults() *EmbeddedAppWidgetInput {
	this := EmbeddedAppWidgetInput{}
	return &this
}

// GetName returns the Name field value.
func (o *EmbeddedAppWidgetInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EmbeddedAppWidgetInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *EmbeddedAppWidgetInput) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value.
func (o *EmbeddedAppWidgetInput) GetValue() EmbeddedAppWidgetInputValue {
	if o == nil {
		var ret EmbeddedAppWidgetInputValue
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *EmbeddedAppWidgetInput) GetValueOk() (*EmbeddedAppWidgetInputValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *EmbeddedAppWidgetInput) SetValue(v EmbeddedAppWidgetInputValue) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EmbeddedAppWidgetInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EmbeddedAppWidgetInput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name  *string                      `json:"name"`
		Value *EmbeddedAppWidgetInputValue `json:"value"`
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
	o.Name = *all.Name
	o.Value = *all.Value

	return nil
}

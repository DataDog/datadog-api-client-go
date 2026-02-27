// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookTemplateVariable Template variable for a notebook.
type NotebookTemplateVariable struct {
	// The list of values that the template variable drop-down is limited to.
	AvailableValues datadog.NullableList[string] `json:"available_values,omitempty"`
	// The default value for the template variable.
	Default datadog.NullableString `json:"default,omitempty"`
	// The name of the variable.
	Name string `json:"name"`
	// The tag prefix associated with the variable. Only tags with this prefix appear in the variable drop-down.
	Prefix datadog.NullableString `json:"prefix,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNotebookTemplateVariable instantiates a new NotebookTemplateVariable object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotebookTemplateVariable(name string) *NotebookTemplateVariable {
	this := NotebookTemplateVariable{}
	this.Name = name
	return &this
}

// NewNotebookTemplateVariableWithDefaults instantiates a new NotebookTemplateVariable object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotebookTemplateVariableWithDefaults() *NotebookTemplateVariable {
	this := NotebookTemplateVariable{}
	return &this
}

// GetAvailableValues returns the AvailableValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotebookTemplateVariable) GetAvailableValues() []string {
	if o == nil || o.AvailableValues.Get() == nil {
		var ret []string
		return ret
	}
	return *o.AvailableValues.Get()
}

// GetAvailableValuesOk returns a tuple with the AvailableValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *NotebookTemplateVariable) GetAvailableValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableValues.Get(), o.AvailableValues.IsSet()
}

// HasAvailableValues returns a boolean if a field has been set.
func (o *NotebookTemplateVariable) HasAvailableValues() bool {
	return o != nil && o.AvailableValues.IsSet()
}

// SetAvailableValues gets a reference to the given datadog.NullableList[string] and assigns it to the AvailableValues field.
func (o *NotebookTemplateVariable) SetAvailableValues(v []string) {
	o.AvailableValues.Set(&v)
}

// SetAvailableValuesNil sets the value for AvailableValues to be an explicit nil.
func (o *NotebookTemplateVariable) SetAvailableValuesNil() {
	o.AvailableValues.Set(nil)
}

// UnsetAvailableValues ensures that no value is present for AvailableValues, not even an explicit nil.
func (o *NotebookTemplateVariable) UnsetAvailableValues() {
	o.AvailableValues.Unset()
}

// GetDefault returns the Default field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotebookTemplateVariable) GetDefault() string {
	if o == nil || o.Default.Get() == nil {
		var ret string
		return ret
	}
	return *o.Default.Get()
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *NotebookTemplateVariable) GetDefaultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Default.Get(), o.Default.IsSet()
}

// HasDefault returns a boolean if a field has been set.
func (o *NotebookTemplateVariable) HasDefault() bool {
	return o != nil && o.Default.IsSet()
}

// SetDefault gets a reference to the given datadog.NullableString and assigns it to the Default field.
func (o *NotebookTemplateVariable) SetDefault(v string) {
	o.Default.Set(&v)
}

// SetDefaultNil sets the value for Default to be an explicit nil.
func (o *NotebookTemplateVariable) SetDefaultNil() {
	o.Default.Set(nil)
}

// UnsetDefault ensures that no value is present for Default, not even an explicit nil.
func (o *NotebookTemplateVariable) UnsetDefault() {
	o.Default.Unset()
}

// GetName returns the Name field value.
func (o *NotebookTemplateVariable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NotebookTemplateVariable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *NotebookTemplateVariable) SetName(v string) {
	o.Name = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotebookTemplateVariable) GetPrefix() string {
	if o == nil || o.Prefix.Get() == nil {
		var ret string
		return ret
	}
	return *o.Prefix.Get()
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *NotebookTemplateVariable) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prefix.Get(), o.Prefix.IsSet()
}

// HasPrefix returns a boolean if a field has been set.
func (o *NotebookTemplateVariable) HasPrefix() bool {
	return o != nil && o.Prefix.IsSet()
}

// SetPrefix gets a reference to the given datadog.NullableString and assigns it to the Prefix field.
func (o *NotebookTemplateVariable) SetPrefix(v string) {
	o.Prefix.Set(&v)
}

// SetPrefixNil sets the value for Prefix to be an explicit nil.
func (o *NotebookTemplateVariable) SetPrefixNil() {
	o.Prefix.Set(nil)
}

// UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil.
func (o *NotebookTemplateVariable) UnsetPrefix() {
	o.Prefix.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o NotebookTemplateVariable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AvailableValues.IsSet() {
		toSerialize["available_values"] = o.AvailableValues.Get()
	}
	if o.Default.IsSet() {
		toSerialize["default"] = o.Default.Get()
	}
	toSerialize["name"] = o.Name
	if o.Prefix.IsSet() {
		toSerialize["prefix"] = o.Prefix.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotebookTemplateVariable) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AvailableValues datadog.NullableList[string] `json:"available_values,omitempty"`
		Default         datadog.NullableString       `json:"default,omitempty"`
		Name            *string                      `json:"name"`
		Prefix          datadog.NullableString       `json:"prefix,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"available_values", "default", "name", "prefix"})
	} else {
		return err
	}
	o.AvailableValues = all.AvailableValues
	o.Default = all.Default
	o.Name = *all.Name
	o.Prefix = all.Prefix

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

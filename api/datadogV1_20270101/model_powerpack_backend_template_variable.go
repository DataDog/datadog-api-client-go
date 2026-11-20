// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PowerpackBackendTemplateVariable Test-only generated schema for the dashboard versioning proof of concept.
type PowerpackBackendTemplateVariable struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Prefix *string `json:"prefix,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Values []string `json:"values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewPowerpackBackendTemplateVariable instantiates a new PowerpackBackendTemplateVariable object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPowerpackBackendTemplateVariable(name string, values []string) *PowerpackBackendTemplateVariable {
	this := PowerpackBackendTemplateVariable{}
	this.Name = name
	this.Values = values
	return &this
}

// NewPowerpackBackendTemplateVariableWithDefaults instantiates a new PowerpackBackendTemplateVariable object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPowerpackBackendTemplateVariableWithDefaults() *PowerpackBackendTemplateVariable {
	this := PowerpackBackendTemplateVariable{}
	return &this
}

// GetName returns the Name field value.
func (o *PowerpackBackendTemplateVariable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PowerpackBackendTemplateVariable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *PowerpackBackendTemplateVariable) SetName(v string) {
	o.Name = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *PowerpackBackendTemplateVariable) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackBackendTemplateVariable) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *PowerpackBackendTemplateVariable) HasPrefix() bool {
	return o != nil && o.Prefix != nil
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *PowerpackBackendTemplateVariable) SetPrefix(v string) {
	o.Prefix = &v
}

// GetValues returns the Values field value.
func (o *PowerpackBackendTemplateVariable) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *PowerpackBackendTemplateVariable) GetValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *PowerpackBackendTemplateVariable) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PowerpackBackendTemplateVariable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Prefix != nil {
		toSerialize["prefix"] = o.Prefix
	}
	toSerialize["values"] = o.Values
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PowerpackBackendTemplateVariable) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name   *string   `json:"name"`
		Prefix *string   `json:"prefix,omitempty"`
		Values *[]string `json:"values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Values == nil {
		return fmt.Errorf("required field values missing")
	}
	o.Name = *all.Name
	o.Prefix = all.Prefix
	o.Values = *all.Values

	return nil
}

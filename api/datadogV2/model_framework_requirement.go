// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FrameworkRequirement Framework Requirement.
type FrameworkRequirement struct {
	// Requirement Controls.
	Controls []FrameworkControl `json:"controls"`
	// Requirement Name.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFrameworkRequirement instantiates a new FrameworkRequirement object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFrameworkRequirement(controls []FrameworkControl, name string) *FrameworkRequirement {
	this := FrameworkRequirement{}
	this.Controls = controls
	this.Name = name
	return &this
}

// NewFrameworkRequirementWithDefaults instantiates a new FrameworkRequirement object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFrameworkRequirementWithDefaults() *FrameworkRequirement {
	this := FrameworkRequirement{}
	return &this
}

// GetControls returns the Controls field value.
func (o *FrameworkRequirement) GetControls() []FrameworkControl {
	if o == nil {
		var ret []FrameworkControl
		return ret
	}
	return o.Controls
}

// GetControlsOk returns a tuple with the Controls field value
// and a boolean to check if the value has been set.
func (o *FrameworkRequirement) GetControlsOk() (*[]FrameworkControl, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Controls, true
}

// SetControls sets field value.
func (o *FrameworkRequirement) SetControls(v []FrameworkControl) {
	o.Controls = v
}

// GetName returns the Name field value.
func (o *FrameworkRequirement) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FrameworkRequirement) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FrameworkRequirement) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FrameworkRequirement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["controls"] = o.Controls
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FrameworkRequirement) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Controls *[]FrameworkControl `json:"controls"`
		Name     *string             `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Controls == nil {
		return fmt.Errorf("required field controls missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"controls", "name"})
	} else {
		return err
	}
	o.Controls = *all.Controls
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

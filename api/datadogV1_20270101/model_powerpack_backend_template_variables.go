// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PowerpackBackendTemplateVariables Test-only generated schema for the dashboard versioning proof of concept.
type PowerpackBackendTemplateVariables struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ControlledByPowerpack []PowerpackBackendTemplateVariable `json:"controlled_by_powerpack"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ControlledExternally []PowerpackBackendTemplateVariable `json:"controlled_externally"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewPowerpackBackendTemplateVariables instantiates a new PowerpackBackendTemplateVariables object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPowerpackBackendTemplateVariables(controlledByPowerpack []PowerpackBackendTemplateVariable, controlledExternally []PowerpackBackendTemplateVariable) *PowerpackBackendTemplateVariables {
	this := PowerpackBackendTemplateVariables{}
	this.ControlledByPowerpack = controlledByPowerpack
	this.ControlledExternally = controlledExternally
	return &this
}

// NewPowerpackBackendTemplateVariablesWithDefaults instantiates a new PowerpackBackendTemplateVariables object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPowerpackBackendTemplateVariablesWithDefaults() *PowerpackBackendTemplateVariables {
	this := PowerpackBackendTemplateVariables{}
	return &this
}

// GetControlledByPowerpack returns the ControlledByPowerpack field value.
func (o *PowerpackBackendTemplateVariables) GetControlledByPowerpack() []PowerpackBackendTemplateVariable {
	if o == nil {
		var ret []PowerpackBackendTemplateVariable
		return ret
	}
	return o.ControlledByPowerpack
}

// GetControlledByPowerpackOk returns a tuple with the ControlledByPowerpack field value
// and a boolean to check if the value has been set.
func (o *PowerpackBackendTemplateVariables) GetControlledByPowerpackOk() (*[]PowerpackBackendTemplateVariable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ControlledByPowerpack, true
}

// SetControlledByPowerpack sets field value.
func (o *PowerpackBackendTemplateVariables) SetControlledByPowerpack(v []PowerpackBackendTemplateVariable) {
	o.ControlledByPowerpack = v
}

// GetControlledExternally returns the ControlledExternally field value.
func (o *PowerpackBackendTemplateVariables) GetControlledExternally() []PowerpackBackendTemplateVariable {
	if o == nil {
		var ret []PowerpackBackendTemplateVariable
		return ret
	}
	return o.ControlledExternally
}

// GetControlledExternallyOk returns a tuple with the ControlledExternally field value
// and a boolean to check if the value has been set.
func (o *PowerpackBackendTemplateVariables) GetControlledExternallyOk() (*[]PowerpackBackendTemplateVariable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ControlledExternally, true
}

// SetControlledExternally sets field value.
func (o *PowerpackBackendTemplateVariables) SetControlledExternally(v []PowerpackBackendTemplateVariable) {
	o.ControlledExternally = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PowerpackBackendTemplateVariables) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["controlled_by_powerpack"] = o.ControlledByPowerpack
	toSerialize["controlled_externally"] = o.ControlledExternally
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PowerpackBackendTemplateVariables) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ControlledByPowerpack *[]PowerpackBackendTemplateVariable `json:"controlled_by_powerpack"`
		ControlledExternally  *[]PowerpackBackendTemplateVariable `json:"controlled_externally"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ControlledByPowerpack == nil {
		return fmt.Errorf("required field controlled_by_powerpack missing")
	}
	if all.ControlledExternally == nil {
		return fmt.Errorf("required field controlled_externally missing")
	}
	o.ControlledByPowerpack = *all.ControlledByPowerpack
	o.ControlledExternally = *all.ControlledExternally

	return nil
}

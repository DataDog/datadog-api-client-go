// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSegmentTemplateParameterDef A parameter definition for a segment template.
type RumSegmentTemplateParameterDef struct {
	// The default value for the parameter.
	Default string `json:"default"`
	// A description of the parameter.
	Description string `json:"description"`
	// Validation rules for the parameter.
	Validate string `json:"validate"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumSegmentTemplateParameterDef instantiates a new RumSegmentTemplateParameterDef object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumSegmentTemplateParameterDef(defaultVar string, description string, validate string) *RumSegmentTemplateParameterDef {
	this := RumSegmentTemplateParameterDef{}
	this.Default = defaultVar
	this.Description = description
	this.Validate = validate
	return &this
}

// NewRumSegmentTemplateParameterDefWithDefaults instantiates a new RumSegmentTemplateParameterDef object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumSegmentTemplateParameterDefWithDefaults() *RumSegmentTemplateParameterDef {
	this := RumSegmentTemplateParameterDef{}
	return &this
}

// GetDefault returns the Default field value.
func (o *RumSegmentTemplateParameterDef) GetDefault() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *RumSegmentTemplateParameterDef) GetDefaultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value.
func (o *RumSegmentTemplateParameterDef) SetDefault(v string) {
	o.Default = v
}

// GetDescription returns the Description field value.
func (o *RumSegmentTemplateParameterDef) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RumSegmentTemplateParameterDef) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *RumSegmentTemplateParameterDef) SetDescription(v string) {
	o.Description = v
}

// GetValidate returns the Validate field value.
func (o *RumSegmentTemplateParameterDef) GetValidate() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Validate
}

// GetValidateOk returns a tuple with the Validate field value
// and a boolean to check if the value has been set.
func (o *RumSegmentTemplateParameterDef) GetValidateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Validate, true
}

// SetValidate sets field value.
func (o *RumSegmentTemplateParameterDef) SetValidate(v string) {
	o.Validate = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumSegmentTemplateParameterDef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["default"] = o.Default
	toSerialize["description"] = o.Description
	toSerialize["validate"] = o.Validate

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumSegmentTemplateParameterDef) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Default     *string `json:"default"`
		Description *string `json:"description"`
		Validate    *string `json:"validate"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Default == nil {
		return fmt.Errorf("required field default missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Validate == nil {
		return fmt.Errorf("required field validate missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"default", "description", "validate"})
	} else {
		return err
	}
	o.Default = *all.Default
	o.Description = *all.Description
	o.Validate = *all.Validate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

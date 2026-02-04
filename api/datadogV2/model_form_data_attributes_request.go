// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormDataAttributesRequest Attributes for creating a form.
type FormDataAttributesRequest struct {
	// The data definition for the form.
	DataDefinition interface{} `json:"data_definition"`
	// The description of the form.
	Description string `json:"description"`
	// The name of the form.
	Name string `json:"name"`
	// The UI definition for the form.
	UiDefinition interface{} `json:"ui_definition"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormDataAttributesRequest instantiates a new FormDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormDataAttributesRequest(dataDefinition interface{}, description string, name string, uiDefinition interface{}) *FormDataAttributesRequest {
	this := FormDataAttributesRequest{}
	this.DataDefinition = dataDefinition
	this.Description = description
	this.Name = name
	this.UiDefinition = uiDefinition
	return &this
}

// NewFormDataAttributesRequestWithDefaults instantiates a new FormDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormDataAttributesRequestWithDefaults() *FormDataAttributesRequest {
	this := FormDataAttributesRequest{}
	return &this
}

// GetDataDefinition returns the DataDefinition field value.
func (o *FormDataAttributesRequest) GetDataDefinition() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DataDefinition
}

// GetDataDefinitionOk returns a tuple with the DataDefinition field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributesRequest) GetDataDefinitionOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataDefinition, true
}

// SetDataDefinition sets field value.
func (o *FormDataAttributesRequest) SetDataDefinition(v interface{}) {
	o.DataDefinition = v
}

// GetDescription returns the Description field value.
func (o *FormDataAttributesRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributesRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *FormDataAttributesRequest) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value.
func (o *FormDataAttributesRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributesRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormDataAttributesRequest) SetName(v string) {
	o.Name = v
}

// GetUiDefinition returns the UiDefinition field value.
func (o *FormDataAttributesRequest) GetUiDefinition() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UiDefinition
}

// GetUiDefinitionOk returns a tuple with the UiDefinition field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributesRequest) GetUiDefinitionOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UiDefinition, true
}

// SetUiDefinition sets field value.
func (o *FormDataAttributesRequest) SetUiDefinition(v interface{}) {
	o.UiDefinition = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_definition"] = o.DataDefinition
	toSerialize["description"] = o.Description
	toSerialize["name"] = o.Name
	toSerialize["ui_definition"] = o.UiDefinition

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataDefinition *interface{} `json:"data_definition"`
		Description    *string      `json:"description"`
		Name           *string      `json:"name"`
		UiDefinition   *interface{} `json:"ui_definition"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataDefinition == nil {
		return fmt.Errorf("required field data_definition missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.UiDefinition == nil {
		return fmt.Errorf("required field ui_definition missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_definition", "description", "name", "ui_definition"})
	} else {
		return err
	}
	o.DataDefinition = *all.DataDefinition
	o.Description = *all.Description
	o.Name = *all.Name
	o.UiDefinition = *all.UiDefinition

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

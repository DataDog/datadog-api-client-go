// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormVersionAttributes
type FormVersionAttributes struct {
	// The data definition for the form.
	DataDefinition interface{} `json:"data_definition"`
	// The state of the form version.
	State *FormVersionState `json:"state,omitempty"`
	// The UI definition for the form.
	UiDefinition interface{} `json:"ui_definition"`
	// Parameters for upserting a form version.
	UpsertParams FormVersionUpsertParams `json:"upsert_params"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormVersionAttributes instantiates a new FormVersionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormVersionAttributes(dataDefinition interface{}, uiDefinition interface{}, upsertParams FormVersionUpsertParams) *FormVersionAttributes {
	this := FormVersionAttributes{}
	this.DataDefinition = dataDefinition
	this.UiDefinition = uiDefinition
	this.UpsertParams = upsertParams
	return &this
}

// NewFormVersionAttributesWithDefaults instantiates a new FormVersionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormVersionAttributesWithDefaults() *FormVersionAttributes {
	this := FormVersionAttributes{}
	return &this
}

// GetDataDefinition returns the DataDefinition field value.
func (o *FormVersionAttributes) GetDataDefinition() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DataDefinition
}

// GetDataDefinitionOk returns a tuple with the DataDefinition field value
// and a boolean to check if the value has been set.
func (o *FormVersionAttributes) GetDataDefinitionOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataDefinition, true
}

// SetDataDefinition sets field value.
func (o *FormVersionAttributes) SetDataDefinition(v interface{}) {
	o.DataDefinition = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *FormVersionAttributes) GetState() FormVersionState {
	if o == nil || o.State == nil {
		var ret FormVersionState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormVersionAttributes) GetStateOk() (*FormVersionState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *FormVersionAttributes) HasState() bool {
	return o != nil && o.State != nil
}

// SetState gets a reference to the given FormVersionState and assigns it to the State field.
func (o *FormVersionAttributes) SetState(v FormVersionState) {
	o.State = &v
}

// GetUiDefinition returns the UiDefinition field value.
func (o *FormVersionAttributes) GetUiDefinition() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UiDefinition
}

// GetUiDefinitionOk returns a tuple with the UiDefinition field value
// and a boolean to check if the value has been set.
func (o *FormVersionAttributes) GetUiDefinitionOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UiDefinition, true
}

// SetUiDefinition sets field value.
func (o *FormVersionAttributes) SetUiDefinition(v interface{}) {
	o.UiDefinition = v
}

// GetUpsertParams returns the UpsertParams field value.
func (o *FormVersionAttributes) GetUpsertParams() FormVersionUpsertParams {
	if o == nil {
		var ret FormVersionUpsertParams
		return ret
	}
	return o.UpsertParams
}

// GetUpsertParamsOk returns a tuple with the UpsertParams field value
// and a boolean to check if the value has been set.
func (o *FormVersionAttributes) GetUpsertParamsOk() (*FormVersionUpsertParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpsertParams, true
}

// SetUpsertParams sets field value.
func (o *FormVersionAttributes) SetUpsertParams(v FormVersionUpsertParams) {
	o.UpsertParams = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormVersionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_definition"] = o.DataDefinition
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	toSerialize["ui_definition"] = o.UiDefinition
	toSerialize["upsert_params"] = o.UpsertParams

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormVersionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataDefinition *interface{}             `json:"data_definition"`
		State          *FormVersionState        `json:"state,omitempty"`
		UiDefinition   *interface{}             `json:"ui_definition"`
		UpsertParams   *FormVersionUpsertParams `json:"upsert_params"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataDefinition == nil {
		return fmt.Errorf("required field data_definition missing")
	}
	if all.UiDefinition == nil {
		return fmt.Errorf("required field ui_definition missing")
	}
	if all.UpsertParams == nil {
		return fmt.Errorf("required field upsert_params missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_definition", "state", "ui_definition", "upsert_params"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DataDefinition = *all.DataDefinition
	if all.State != nil && !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = all.State
	}
	o.UiDefinition = *all.UiDefinition
	if all.UpsertParams.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UpsertParams = *all.UpsertParams

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

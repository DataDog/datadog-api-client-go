/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// NotebookLogStreamCellAttributes The attributes of a notebook `log_stream` cell.
type NotebookLogStreamCellAttributes struct {
	Definition LogStreamWidgetDefinition `json:"definition"`
	GraphSize  *NotebookGraphSize        `json:"graph_size,omitempty"`
	Time       NullableNotebookCellTime  `json:"time,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewNotebookLogStreamCellAttributes instantiates a new NotebookLogStreamCellAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotebookLogStreamCellAttributes(definition LogStreamWidgetDefinition) *NotebookLogStreamCellAttributes {
	this := NotebookLogStreamCellAttributes{}
	this.Definition = definition
	return &this
}

// NewNotebookLogStreamCellAttributesWithDefaults instantiates a new NotebookLogStreamCellAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotebookLogStreamCellAttributesWithDefaults() *NotebookLogStreamCellAttributes {
	this := NotebookLogStreamCellAttributes{}
	return &this
}

// GetDefinition returns the Definition field value
func (o *NotebookLogStreamCellAttributes) GetDefinition() LogStreamWidgetDefinition {
	if o == nil {
		var ret LogStreamWidgetDefinition
		return ret
	}

	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *NotebookLogStreamCellAttributes) GetDefinitionOk() (*LogStreamWidgetDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value
func (o *NotebookLogStreamCellAttributes) SetDefinition(v LogStreamWidgetDefinition) {
	o.Definition = v
}

// GetGraphSize returns the GraphSize field value if set, zero value otherwise.
func (o *NotebookLogStreamCellAttributes) GetGraphSize() NotebookGraphSize {
	if o == nil || o.GraphSize == nil {
		var ret NotebookGraphSize
		return ret
	}
	return *o.GraphSize
}

// GetGraphSizeOk returns a tuple with the GraphSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookLogStreamCellAttributes) GetGraphSizeOk() (*NotebookGraphSize, bool) {
	if o == nil || o.GraphSize == nil {
		return nil, false
	}
	return o.GraphSize, true
}

// HasGraphSize returns a boolean if a field has been set.
func (o *NotebookLogStreamCellAttributes) HasGraphSize() bool {
	if o != nil && o.GraphSize != nil {
		return true
	}

	return false
}

// SetGraphSize gets a reference to the given NotebookGraphSize and assigns it to the GraphSize field.
func (o *NotebookLogStreamCellAttributes) SetGraphSize(v NotebookGraphSize) {
	o.GraphSize = &v
}

// GetTime returns the Time field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotebookLogStreamCellAttributes) GetTime() NotebookCellTime {
	if o == nil || o.Time.Get() == nil {
		var ret NotebookCellTime
		return ret
	}
	return *o.Time.Get()
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotebookLogStreamCellAttributes) GetTimeOk() (*NotebookCellTime, bool) {
	if o == nil {
		return nil, false
	}
	return o.Time.Get(), o.Time.IsSet()
}

// HasTime returns a boolean if a field has been set.
func (o *NotebookLogStreamCellAttributes) HasTime() bool {
	if o != nil && o.Time.IsSet() {
		return true
	}

	return false
}

// SetTime gets a reference to the given NullableNotebookCellTime and assigns it to the Time field.
func (o *NotebookLogStreamCellAttributes) SetTime(v NotebookCellTime) {
	o.Time.Set(&v)
}

// SetTimeNil sets the value for Time to be an explicit nil
func (o *NotebookLogStreamCellAttributes) SetTimeNil() {
	o.Time.Set(nil)
}

// UnsetTime ensures that no value is present for Time, not even an explicit nil
func (o *NotebookLogStreamCellAttributes) UnsetTime() {
	o.Time.Unset()
}

func (o NotebookLogStreamCellAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if true {
		toSerialize["definition"] = o.Definition
	}
	if o.GraphSize != nil {
		toSerialize["graph_size"] = o.GraphSize
	}
	if o.Time.IsSet() {
		toSerialize["time"] = o.Time.Get()
	}
	return json.Marshal(toSerialize)
}

func (o *NotebookLogStreamCellAttributes) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Definition *LogStreamWidgetDefinition `json:"definition"`
	}{}
	all := struct {
		Definition LogStreamWidgetDefinition `json:"definition"`
		GraphSize  *NotebookGraphSize        `json:"graph_size,omitempty"`
		Time       NullableNotebookCellTime  `json:"time,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Definition == nil {
		return fmt.Errorf("Required field definition missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.GraphSize; v != nil && !v.IsValid() {
		errr := json.Unmarshal(bytes, &raw)
		if errr != nil {
			return errr
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Definition = all.Definition
	o.GraphSize = all.GraphSize
	o.Time = all.Time
	return nil
}

type NullableNotebookLogStreamCellAttributes struct {
	value *NotebookLogStreamCellAttributes
	isSet bool
}

func (v NullableNotebookLogStreamCellAttributes) Get() *NotebookLogStreamCellAttributes {
	return v.value
}

func (v *NullableNotebookLogStreamCellAttributes) Set(val *NotebookLogStreamCellAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableNotebookLogStreamCellAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableNotebookLogStreamCellAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotebookLogStreamCellAttributes(val *NotebookLogStreamCellAttributes) *NullableNotebookLogStreamCellAttributes {
	return &NullableNotebookLogStreamCellAttributes{value: val, isSet: true}
}

func (v NullableNotebookLogStreamCellAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotebookLogStreamCellAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

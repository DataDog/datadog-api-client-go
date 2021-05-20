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

// NotebookMarkdownCellAttributes The attributes of a notebook `markdown` cell.
type NotebookMarkdownCellAttributes struct {
	Definition NotebookMarkdownCellDefinition `json:"definition"`
}

// NewNotebookMarkdownCellAttributes instantiates a new NotebookMarkdownCellAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotebookMarkdownCellAttributes(definition NotebookMarkdownCellDefinition) *NotebookMarkdownCellAttributes {
	this := NotebookMarkdownCellAttributes{}
	this.Definition = definition
	return &this
}

// NewNotebookMarkdownCellAttributesWithDefaults instantiates a new NotebookMarkdownCellAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotebookMarkdownCellAttributesWithDefaults() *NotebookMarkdownCellAttributes {
	this := NotebookMarkdownCellAttributes{}
	return &this
}

// GetDefinition returns the Definition field value
func (o *NotebookMarkdownCellAttributes) GetDefinition() NotebookMarkdownCellDefinition {
	if o == nil {
		var ret NotebookMarkdownCellDefinition
		return ret
	}

	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *NotebookMarkdownCellAttributes) GetDefinitionOk() (*NotebookMarkdownCellDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value
func (o *NotebookMarkdownCellAttributes) SetDefinition(v NotebookMarkdownCellDefinition) {
	o.Definition = v
}

func (o NotebookMarkdownCellAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["definition"] = o.Definition
	}
	return json.Marshal(toSerialize)
}

func (o *NotebookMarkdownCellAttributes) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Definition *NotebookMarkdownCellDefinition `json:"definition"`
	}{}
	all := struct {
		Definition NotebookMarkdownCellDefinition `json:"definition"}`
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
		return err
	}
	o.Definition = all.Definition
	return nil
}

type NullableNotebookMarkdownCellAttributes struct {
	value *NotebookMarkdownCellAttributes
	isSet bool
}

func (v NullableNotebookMarkdownCellAttributes) Get() *NotebookMarkdownCellAttributes {
	return v.value
}

func (v *NullableNotebookMarkdownCellAttributes) Set(val *NotebookMarkdownCellAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableNotebookMarkdownCellAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableNotebookMarkdownCellAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotebookMarkdownCellAttributes(val *NotebookMarkdownCellAttributes) *NullableNotebookMarkdownCellAttributes {
	return &NullableNotebookMarkdownCellAttributes{value: val, isSet: true}
}

func (v NullableNotebookMarkdownCellAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotebookMarkdownCellAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

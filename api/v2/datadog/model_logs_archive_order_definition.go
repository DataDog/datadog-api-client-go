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

// LogsArchiveOrderDefinition The definition of an archive order.
type LogsArchiveOrderDefinition struct {
	Attributes LogsArchiveOrderAttributes     `json:"attributes"`
	Type       LogsArchiveOrderDefinitionType `json:"type"`
}

// NewLogsArchiveOrderDefinition instantiates a new LogsArchiveOrderDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsArchiveOrderDefinition(attributes LogsArchiveOrderAttributes, type_ LogsArchiveOrderDefinitionType) *LogsArchiveOrderDefinition {
	this := LogsArchiveOrderDefinition{}
	this.Attributes = attributes
	this.Type = type_
	return &this
}

// NewLogsArchiveOrderDefinitionWithDefaults instantiates a new LogsArchiveOrderDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsArchiveOrderDefinitionWithDefaults() *LogsArchiveOrderDefinition {
	this := LogsArchiveOrderDefinition{}
	var type_ LogsArchiveOrderDefinitionType = LOGSARCHIVEORDERDEFINITIONTYPE_ARCHIVE_ORDER
	this.Type = type_
	return &this
}

// GetAttributes returns the Attributes field value
func (o *LogsArchiveOrderDefinition) GetAttributes() LogsArchiveOrderAttributes {
	if o == nil {
		var ret LogsArchiveOrderAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *LogsArchiveOrderDefinition) GetAttributesOk() (*LogsArchiveOrderAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *LogsArchiveOrderDefinition) SetAttributes(v LogsArchiveOrderAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value
func (o *LogsArchiveOrderDefinition) GetType() LogsArchiveOrderDefinitionType {
	if o == nil {
		var ret LogsArchiveOrderDefinitionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsArchiveOrderDefinition) GetTypeOk() (*LogsArchiveOrderDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LogsArchiveOrderDefinition) SetType(v LogsArchiveOrderDefinitionType) {
	o.Type = v
}

func (o LogsArchiveOrderDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

func (o *LogsArchiveOrderDefinition) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Attributes *LogsArchiveOrderAttributes     `json:"attributes"`
		Type       *LogsArchiveOrderDefinitionType `json:"type"`
	}{}
	all := struct {
		Attributes LogsArchiveOrderAttributes     `json:"attributes"}`
		Type       LogsArchiveOrderDefinitionType `json:"type"}`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Attributes == nil {
		return fmt.Errorf("Required field attributes missing")
	}
	if required.Type == nil {
		return fmt.Errorf("Required field type missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		return err
	}
	o.Attributes = all.Attributes
	o.Type = all.Type
	return nil
}

type NullableLogsArchiveOrderDefinition struct {
	value *LogsArchiveOrderDefinition
	isSet bool
}

func (v NullableLogsArchiveOrderDefinition) Get() *LogsArchiveOrderDefinition {
	return v.value
}

func (v *NullableLogsArchiveOrderDefinition) Set(val *LogsArchiveOrderDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsArchiveOrderDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsArchiveOrderDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsArchiveOrderDefinition(val *LogsArchiveOrderDefinition) *NullableLogsArchiveOrderDefinition {
	return &NullableLogsArchiveOrderDefinition{value: val, isSet: true}
}

func (v NullableLogsArchiveOrderDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsArchiveOrderDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

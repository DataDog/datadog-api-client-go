/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// SyntheticsGlobalVariableParseTestOptions Parser options to use for retrieving a Synthetics global variable from a Synthetics Test. Used in conjunction with `parse_test_public_id`.
type SyntheticsGlobalVariableParseTestOptions struct {
	// When type is `http_header`, name of the header to use to extract the value.
	Field  *string                                        `json:"field,omitempty"`
	Parser SyntheticsGlobalVariableParseTestOptionsParser `json:"parser"`
	Type   SyntheticsGlobalVariableParseTestOptionsType   `json:"type"`
}

// NewSyntheticsGlobalVariableParseTestOptions instantiates a new SyntheticsGlobalVariableParseTestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsGlobalVariableParseTestOptions(parser SyntheticsGlobalVariableParseTestOptionsParser, type_ SyntheticsGlobalVariableParseTestOptionsType) *SyntheticsGlobalVariableParseTestOptions {
	this := SyntheticsGlobalVariableParseTestOptions{}
	this.Parser = parser
	this.Type = type_
	return &this
}

// NewSyntheticsGlobalVariableParseTestOptionsWithDefaults instantiates a new SyntheticsGlobalVariableParseTestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsGlobalVariableParseTestOptionsWithDefaults() *SyntheticsGlobalVariableParseTestOptions {
	this := SyntheticsGlobalVariableParseTestOptions{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *SyntheticsGlobalVariableParseTestOptions) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsGlobalVariableParseTestOptions) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *SyntheticsGlobalVariableParseTestOptions) HasField() bool {
	if o != nil && o.Field != nil {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *SyntheticsGlobalVariableParseTestOptions) SetField(v string) {
	o.Field = &v
}

// GetParser returns the Parser field value
func (o *SyntheticsGlobalVariableParseTestOptions) GetParser() SyntheticsGlobalVariableParseTestOptionsParser {
	if o == nil {
		var ret SyntheticsGlobalVariableParseTestOptionsParser
		return ret
	}

	return o.Parser
}

// GetParserOk returns a tuple with the Parser field value
// and a boolean to check if the value has been set.
func (o *SyntheticsGlobalVariableParseTestOptions) GetParserOk() (*SyntheticsGlobalVariableParseTestOptionsParser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parser, true
}

// SetParser sets field value
func (o *SyntheticsGlobalVariableParseTestOptions) SetParser(v SyntheticsGlobalVariableParseTestOptionsParser) {
	o.Parser = v
}

// GetType returns the Type field value
func (o *SyntheticsGlobalVariableParseTestOptions) GetType() SyntheticsGlobalVariableParseTestOptionsType {
	if o == nil {
		var ret SyntheticsGlobalVariableParseTestOptionsType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticsGlobalVariableParseTestOptions) GetTypeOk() (*SyntheticsGlobalVariableParseTestOptionsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SyntheticsGlobalVariableParseTestOptions) SetType(v SyntheticsGlobalVariableParseTestOptionsType) {
	o.Type = v
}

func (o SyntheticsGlobalVariableParseTestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	if true {
		toSerialize["parser"] = o.Parser
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticsGlobalVariableParseTestOptions struct {
	value *SyntheticsGlobalVariableParseTestOptions
	isSet bool
}

func (v NullableSyntheticsGlobalVariableParseTestOptions) Get() *SyntheticsGlobalVariableParseTestOptions {
	return v.value
}

func (v *NullableSyntheticsGlobalVariableParseTestOptions) Set(val *SyntheticsGlobalVariableParseTestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsGlobalVariableParseTestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsGlobalVariableParseTestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsGlobalVariableParseTestOptions(val *SyntheticsGlobalVariableParseTestOptions) *NullableSyntheticsGlobalVariableParseTestOptions {
	return &NullableSyntheticsGlobalVariableParseTestOptions{value: val, isSet: true}
}

func (v NullableSyntheticsGlobalVariableParseTestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsGlobalVariableParseTestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

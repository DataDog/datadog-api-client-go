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

// LogsStringBuilderProcessor Use the string builder processor to add a new attribute (without spaces or special characters) to a log with the result of the provided template. This enables aggregation of different attributes or raw strings into a single attribute.  The template is defined by both raw text and blocks with the syntax `%{attribute_path}`.  **Notes**:  - The processor only accepts attributes with values or an array of values in the blocks. - If an attribute cannot be used (object or array of object),   it is replaced by an empty string or the entire operation is skipped depending on your selection. - If the target attribute already exists, it is overwritten by the result of the template. - Results of the template cannot exceed 256 characters.
type LogsStringBuilderProcessor struct {
	// If true, it replaces all missing attributes of `template` by an empty string. If `false` (default), skips the operation for missing attributes.
	IsReplaceMissing *bool `json:"is_replace_missing,omitempty"`
	// The name of the attribute that contains the result of the template.
	Target string `json:"target"`
	// A formula with one or more attributes and raw text.
	Template string `json:"template"`
	// Type of processor.
	Type *string `json:"type,omitempty"`
	// Whether or not the processor is enabled.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Name of the processor.
	Name *string `json:"name,omitempty"`
}

// NewLogsStringBuilderProcessor instantiates a new LogsStringBuilderProcessor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsStringBuilderProcessor(target string, template string) *LogsStringBuilderProcessor {
	this := LogsStringBuilderProcessor{}
	var isReplaceMissing bool = false
	this.IsReplaceMissing = &isReplaceMissing
	this.Target = target
	this.Template = template
	var type_ string = "string-builder-processor"
	this.Type = &type_
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	return &this
}

// NewLogsStringBuilderProcessorWithDefaults instantiates a new LogsStringBuilderProcessor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsStringBuilderProcessorWithDefaults() *LogsStringBuilderProcessor {
	this := LogsStringBuilderProcessor{}
	var isReplaceMissing bool = false
	this.IsReplaceMissing = &isReplaceMissing
	var type_ string = "string-builder-processor"
	this.Type = &type_
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	return &this
}

// GetIsReplaceMissing returns the IsReplaceMissing field value if set, zero value otherwise.
func (o *LogsStringBuilderProcessor) GetIsReplaceMissing() bool {
	if o == nil || o.IsReplaceMissing == nil {
		var ret bool
		return ret
	}
	return *o.IsReplaceMissing
}

// GetIsReplaceMissingOk returns a tuple with the IsReplaceMissing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsStringBuilderProcessor) GetIsReplaceMissingOk() (*bool, bool) {
	if o == nil || o.IsReplaceMissing == nil {
		return nil, false
	}
	return o.IsReplaceMissing, true
}

// HasIsReplaceMissing returns a boolean if a field has been set.
func (o *LogsStringBuilderProcessor) HasIsReplaceMissing() bool {
	if o != nil && o.IsReplaceMissing != nil {
		return true
	}

	return false
}

// SetIsReplaceMissing gets a reference to the given bool and assigns it to the IsReplaceMissing field.
func (o *LogsStringBuilderProcessor) SetIsReplaceMissing(v bool) {
	o.IsReplaceMissing = &v
}

// GetTarget returns the Target field value
func (o *LogsStringBuilderProcessor) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *LogsStringBuilderProcessor) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *LogsStringBuilderProcessor) SetTarget(v string) {
	o.Target = v
}

// GetTemplate returns the Template field value
func (o *LogsStringBuilderProcessor) GetTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *LogsStringBuilderProcessor) GetTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *LogsStringBuilderProcessor) SetTemplate(v string) {
	o.Template = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LogsStringBuilderProcessor) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsStringBuilderProcessor) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LogsStringBuilderProcessor) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LogsStringBuilderProcessor) SetType(v string) {
	o.Type = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *LogsStringBuilderProcessor) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsStringBuilderProcessor) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *LogsStringBuilderProcessor) HasIsEnabled() bool {
	if o != nil && o.IsEnabled != nil {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *LogsStringBuilderProcessor) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogsStringBuilderProcessor) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsStringBuilderProcessor) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogsStringBuilderProcessor) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogsStringBuilderProcessor) SetName(v string) {
	o.Name = &v
}

func (o LogsStringBuilderProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsReplaceMissing != nil {
		toSerialize["is_replace_missing"] = o.IsReplaceMissing
	}
	if true {
		toSerialize["target"] = o.Target
	}
	if true {
		toSerialize["template"] = o.Template
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.IsEnabled != nil {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

// AsLogsProcessor wraps this instance of LogsStringBuilderProcessor in LogsProcessor
func (s *LogsStringBuilderProcessor) AsLogsProcessor() LogsProcessor {
	return LogsProcessor{LogsProcessorInterface: s}
}

type NullableLogsStringBuilderProcessor struct {
	value *LogsStringBuilderProcessor
	isSet bool
}

func (v NullableLogsStringBuilderProcessor) Get() *LogsStringBuilderProcessor {
	return v.value
}

func (v *NullableLogsStringBuilderProcessor) Set(val *LogsStringBuilderProcessor) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsStringBuilderProcessor) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsStringBuilderProcessor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsStringBuilderProcessor(val *LogsStringBuilderProcessor) *NullableLogsStringBuilderProcessor {
	return &NullableLogsStringBuilderProcessor{value: val, isSet: true}
}

func (v NullableLogsStringBuilderProcessor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsStringBuilderProcessor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

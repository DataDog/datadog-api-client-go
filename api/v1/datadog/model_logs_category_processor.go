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

// LogsCategoryProcessor Use the Category Processor to add a new attribute (without spaces or special characters in the new attribute name) to a log matching a provided search query. Use categories to create groups for an analytical view. For example, URL groups, machine groups, environments, and response time buckets.  **Notes**:  - The syntax of the query is the one of Logs Explorer search bar.   The query can be done on any log attribute or tag, whether it is a facet or not.   Wildcards can also be used inside your query. - Once the log has matched one of the Processor queries, it stops.   Make sure they are properly ordered in case a log could match several queries. - The names of the categories must be unique. - Once defined in the Category Processor, you can map categories to log status using the Log Status Remapper.
type LogsCategoryProcessor struct {
	// Array of filters to match or not a log and their corresponding `name`to assign a custom value to the log.
	Categories []LogsCategoryProcessorCategories `json:"categories"`
	// Name of the target attribute which value is defined by the matching category.
	Target string `json:"target"`
	// Type of processor.
	Type *string `json:"type,omitempty"`
	// Whether or not the processor is enabled.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Name of the processor.
	Name *string `json:"name,omitempty"`
}

// NewLogsCategoryProcessor instantiates a new LogsCategoryProcessor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsCategoryProcessor(categories []LogsCategoryProcessorCategories, target string) *LogsCategoryProcessor {
	this := LogsCategoryProcessor{}
	this.Categories = categories
	this.Target = target
	var type_ string = "category-processor"
	this.Type = &type_
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	return &this
}

// NewLogsCategoryProcessorWithDefaults instantiates a new LogsCategoryProcessor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsCategoryProcessorWithDefaults() *LogsCategoryProcessor {
	this := LogsCategoryProcessor{}
	var type_ string = "category-processor"
	this.Type = &type_
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	return &this
}

// GetCategories returns the Categories field value
func (o *LogsCategoryProcessor) GetCategories() []LogsCategoryProcessorCategories {
	if o == nil {
		var ret []LogsCategoryProcessorCategories
		return ret
	}

	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value
// and a boolean to check if the value has been set.
func (o *LogsCategoryProcessor) GetCategoriesOk() (*[]LogsCategoryProcessorCategories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Categories, true
}

// SetCategories sets field value
func (o *LogsCategoryProcessor) SetCategories(v []LogsCategoryProcessorCategories) {
	o.Categories = v
}

// GetTarget returns the Target field value
func (o *LogsCategoryProcessor) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *LogsCategoryProcessor) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *LogsCategoryProcessor) SetTarget(v string) {
	o.Target = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LogsCategoryProcessor) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsCategoryProcessor) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LogsCategoryProcessor) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LogsCategoryProcessor) SetType(v string) {
	o.Type = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *LogsCategoryProcessor) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsCategoryProcessor) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *LogsCategoryProcessor) HasIsEnabled() bool {
	if o != nil && o.IsEnabled != nil {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *LogsCategoryProcessor) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogsCategoryProcessor) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsCategoryProcessor) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogsCategoryProcessor) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogsCategoryProcessor) SetName(v string) {
	o.Name = &v
}

func (o LogsCategoryProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["categories"] = o.Categories
	}
	if true {
		toSerialize["target"] = o.Target
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

// AsLogsProcessor wraps this instance of LogsCategoryProcessor in LogsProcessor
func (s *LogsCategoryProcessor) AsLogsProcessor() LogsProcessor {
	return LogsProcessor{LogsProcessorInterface: s}
}

type NullableLogsCategoryProcessor struct {
	value *LogsCategoryProcessor
	isSet bool
}

func (v NullableLogsCategoryProcessor) Get() *LogsCategoryProcessor {
	return v.value
}

func (v *NullableLogsCategoryProcessor) Set(val *LogsCategoryProcessor) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsCategoryProcessor) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsCategoryProcessor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsCategoryProcessor(val *LogsCategoryProcessor) *NullableLogsCategoryProcessor {
	return &NullableLogsCategoryProcessor{value: val, isSet: true}
}

func (v NullableLogsCategoryProcessor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsCategoryProcessor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

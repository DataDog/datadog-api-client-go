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

// ChangeWidgetDefinition The Change graph shows you the change in a value over the time period chosen
type ChangeWidgetDefinition struct {
	Requests []ChangeWidgetRequest `json:"requests"`
	Time     *WidgetTime           `json:"time,omitempty"`
	// Title of the widget
	Title      *string          `json:"title,omitempty"`
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Size of the title
	TitleSize *string `json:"title_size,omitempty"`
	// Type of the widget
	Type string `json:"type"`
}

// NewChangeWidgetDefinition instantiates a new ChangeWidgetDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeWidgetDefinition(requests []ChangeWidgetRequest, type_ string) *ChangeWidgetDefinition {
	this := ChangeWidgetDefinition{}
	this.Requests = requests
	this.Type = type_
	return &this
}

// NewChangeWidgetDefinitionWithDefaults instantiates a new ChangeWidgetDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeWidgetDefinitionWithDefaults() *ChangeWidgetDefinition {
	this := ChangeWidgetDefinition{}
	var type_ string = "change"
	this.Type = type_
	return &this
}

// GetRequests returns the Requests field value
func (o *ChangeWidgetDefinition) GetRequests() []ChangeWidgetRequest {
	if o == nil {
		var ret []ChangeWidgetRequest
		return ret
	}

	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *ChangeWidgetDefinition) GetRequestsOk() (*[]ChangeWidgetRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value
func (o *ChangeWidgetDefinition) SetRequests(v []ChangeWidgetRequest) {
	o.Requests = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ChangeWidgetDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeWidgetDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ChangeWidgetDefinition) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *ChangeWidgetDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ChangeWidgetDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeWidgetDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ChangeWidgetDefinition) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ChangeWidgetDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *ChangeWidgetDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *ChangeWidgetDefinition) HasTitleAlign() bool {
	if o != nil && o.TitleAlign != nil {
		return true
	}

	return false
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *ChangeWidgetDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *ChangeWidgetDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeWidgetDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *ChangeWidgetDefinition) HasTitleSize() bool {
	if o != nil && o.TitleSize != nil {
		return true
	}

	return false
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *ChangeWidgetDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value
func (o *ChangeWidgetDefinition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ChangeWidgetDefinition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ChangeWidgetDefinition) SetType(v string) {
	o.Type = v
}

func (o ChangeWidgetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["requests"] = o.Requests
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.TitleAlign != nil {
		toSerialize["title_align"] = o.TitleAlign
	}
	if o.TitleSize != nil {
		toSerialize["title_size"] = o.TitleSize
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

// AsWidgetDefinition wraps this instance of ChangeWidgetDefinition in WidgetDefinition
func (s *ChangeWidgetDefinition) AsWidgetDefinition() WidgetDefinition {
	return WidgetDefinition{WidgetDefinitionInterface: s}
}

type NullableChangeWidgetDefinition struct {
	value *ChangeWidgetDefinition
	isSet bool
}

func (v NullableChangeWidgetDefinition) Get() *ChangeWidgetDefinition {
	return v.value
}

func (v *NullableChangeWidgetDefinition) Set(val *ChangeWidgetDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeWidgetDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeWidgetDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeWidgetDefinition(val *ChangeWidgetDefinition) *NullableChangeWidgetDefinition {
	return &NullableChangeWidgetDefinition{value: val, isSet: true}
}

func (v NullableChangeWidgetDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeWidgetDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

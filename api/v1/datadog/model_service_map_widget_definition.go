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

// ServiceMapWidgetDefinition This widget displays a map of a service to all of the services that call it, and all of the services that it calls.
type ServiceMapWidgetDefinition struct {
	// Your env and primary tag (or * if enabled for your account).
	Filters []string `json:"filters"`
	// The ID of the service you want to map.
	Service string `json:"service"`
	// The title of your widget.
	Title      *string          `json:"title,omitempty"`
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Size of the title
	TitleSize *string `json:"title_size,omitempty"`
	// Type of the widget
	Type string `json:"type"`
}

// NewServiceMapWidgetDefinition instantiates a new ServiceMapWidgetDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceMapWidgetDefinition(filters []string, service string, type_ string) *ServiceMapWidgetDefinition {
	this := ServiceMapWidgetDefinition{}
	this.Filters = filters
	this.Service = service
	this.Type = type_
	return &this
}

// NewServiceMapWidgetDefinitionWithDefaults instantiates a new ServiceMapWidgetDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceMapWidgetDefinitionWithDefaults() *ServiceMapWidgetDefinition {
	this := ServiceMapWidgetDefinition{}
	var type_ string = "servicemap"
	this.Type = type_
	return &this
}

// GetFilters returns the Filters field value
func (o *ServiceMapWidgetDefinition) GetFilters() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *ServiceMapWidgetDefinition) GetFiltersOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filters, true
}

// SetFilters sets field value
func (o *ServiceMapWidgetDefinition) SetFilters(v []string) {
	o.Filters = v
}

// GetService returns the Service field value
func (o *ServiceMapWidgetDefinition) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *ServiceMapWidgetDefinition) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *ServiceMapWidgetDefinition) SetService(v string) {
	o.Service = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ServiceMapWidgetDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceMapWidgetDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ServiceMapWidgetDefinition) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ServiceMapWidgetDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *ServiceMapWidgetDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceMapWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *ServiceMapWidgetDefinition) HasTitleAlign() bool {
	if o != nil && o.TitleAlign != nil {
		return true
	}

	return false
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *ServiceMapWidgetDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *ServiceMapWidgetDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceMapWidgetDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *ServiceMapWidgetDefinition) HasTitleSize() bool {
	if o != nil && o.TitleSize != nil {
		return true
	}

	return false
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *ServiceMapWidgetDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value
func (o *ServiceMapWidgetDefinition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceMapWidgetDefinition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServiceMapWidgetDefinition) SetType(v string) {
	o.Type = v
}

func (o ServiceMapWidgetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["filters"] = o.Filters
	}
	if true {
		toSerialize["service"] = o.Service
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

// AsWidgetDefinition wraps this instance of ServiceMapWidgetDefinition in WidgetDefinition
func (s *ServiceMapWidgetDefinition) AsWidgetDefinition() WidgetDefinition {
	return WidgetDefinition{WidgetDefinitionInterface: s}
}

type NullableServiceMapWidgetDefinition struct {
	value *ServiceMapWidgetDefinition
	isSet bool
}

func (v NullableServiceMapWidgetDefinition) Get() *ServiceMapWidgetDefinition {
	return v.value
}

func (v *NullableServiceMapWidgetDefinition) Set(val *ServiceMapWidgetDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceMapWidgetDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceMapWidgetDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceMapWidgetDefinition(val *ServiceMapWidgetDefinition) *NullableServiceMapWidgetDefinition {
	return &NullableServiceMapWidgetDefinition{value: val, isSet: true}
}

func (v NullableServiceMapWidgetDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceMapWidgetDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

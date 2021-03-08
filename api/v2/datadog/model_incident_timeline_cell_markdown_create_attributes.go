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

// IncidentTimelineCellMarkdownCreateAttributes Timeline cell data for Markdown timeline cells for a create request.
type IncidentTimelineCellMarkdownCreateAttributes struct {
	CellType IncidentTimelineCellMarkdownContentType             `json:"cell_type"`
	Content  IncidentTimelineCellMarkdownCreateAttributesContent `json:"content"`
	// A flag indicating whether the timeline cell is important and should be highlighted.
	Important *bool `json:"important,omitempty"`
}

// NewIncidentTimelineCellMarkdownCreateAttributes instantiates a new IncidentTimelineCellMarkdownCreateAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentTimelineCellMarkdownCreateAttributes(cellType IncidentTimelineCellMarkdownContentType, content IncidentTimelineCellMarkdownCreateAttributesContent) *IncidentTimelineCellMarkdownCreateAttributes {
	this := IncidentTimelineCellMarkdownCreateAttributes{}
	this.CellType = cellType
	this.Content = content
	var important bool = false
	this.Important = &important
	return &this
}

// NewIncidentTimelineCellMarkdownCreateAttributesWithDefaults instantiates a new IncidentTimelineCellMarkdownCreateAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentTimelineCellMarkdownCreateAttributesWithDefaults() *IncidentTimelineCellMarkdownCreateAttributes {
	this := IncidentTimelineCellMarkdownCreateAttributes{}
	var cellType IncidentTimelineCellMarkdownContentType = INCIDENTTIMELINECELLMARKDOWNCONTENTTYPE_MARKDOWN
	this.CellType = cellType
	var important bool = false
	this.Important = &important
	return &this
}

// GetCellType returns the CellType field value
func (o *IncidentTimelineCellMarkdownCreateAttributes) GetCellType() IncidentTimelineCellMarkdownContentType {
	if o == nil {
		var ret IncidentTimelineCellMarkdownContentType
		return ret
	}

	return o.CellType
}

// GetCellTypeOk returns a tuple with the CellType field value
// and a boolean to check if the value has been set.
func (o *IncidentTimelineCellMarkdownCreateAttributes) GetCellTypeOk() (*IncidentTimelineCellMarkdownContentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CellType, true
}

// SetCellType sets field value
func (o *IncidentTimelineCellMarkdownCreateAttributes) SetCellType(v IncidentTimelineCellMarkdownContentType) {
	o.CellType = v
}

// GetContent returns the Content field value
func (o *IncidentTimelineCellMarkdownCreateAttributes) GetContent() IncidentTimelineCellMarkdownCreateAttributesContent {
	if o == nil {
		var ret IncidentTimelineCellMarkdownCreateAttributesContent
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *IncidentTimelineCellMarkdownCreateAttributes) GetContentOk() (*IncidentTimelineCellMarkdownCreateAttributesContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *IncidentTimelineCellMarkdownCreateAttributes) SetContent(v IncidentTimelineCellMarkdownCreateAttributesContent) {
	o.Content = v
}

// GetImportant returns the Important field value if set, zero value otherwise.
func (o *IncidentTimelineCellMarkdownCreateAttributes) GetImportant() bool {
	if o == nil || o.Important == nil {
		var ret bool
		return ret
	}
	return *o.Important
}

// GetImportantOk returns a tuple with the Important field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTimelineCellMarkdownCreateAttributes) GetImportantOk() (*bool, bool) {
	if o == nil || o.Important == nil {
		return nil, false
	}
	return o.Important, true
}

// HasImportant returns a boolean if a field has been set.
func (o *IncidentTimelineCellMarkdownCreateAttributes) HasImportant() bool {
	if o != nil && o.Important != nil {
		return true
	}

	return false
}

// SetImportant gets a reference to the given bool and assigns it to the Important field.
func (o *IncidentTimelineCellMarkdownCreateAttributes) SetImportant(v bool) {
	o.Important = &v
}

func (o IncidentTimelineCellMarkdownCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cell_type"] = o.CellType
	}
	if true {
		toSerialize["content"] = o.Content
	}
	if o.Important != nil {
		toSerialize["important"] = o.Important
	}
	return json.Marshal(toSerialize)
}

type NullableIncidentTimelineCellMarkdownCreateAttributes struct {
	value *IncidentTimelineCellMarkdownCreateAttributes
	isSet bool
}

func (v NullableIncidentTimelineCellMarkdownCreateAttributes) Get() *IncidentTimelineCellMarkdownCreateAttributes {
	return v.value
}

func (v *NullableIncidentTimelineCellMarkdownCreateAttributes) Set(val *IncidentTimelineCellMarkdownCreateAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentTimelineCellMarkdownCreateAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentTimelineCellMarkdownCreateAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentTimelineCellMarkdownCreateAttributes(val *IncidentTimelineCellMarkdownCreateAttributes) *NullableIncidentTimelineCellMarkdownCreateAttributes {
	return &NullableIncidentTimelineCellMarkdownCreateAttributes{value: val, isSet: true}
}

func (v NullableIncidentTimelineCellMarkdownCreateAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentTimelineCellMarkdownCreateAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

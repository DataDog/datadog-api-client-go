// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapBackendRegionsRequest20270101OneOf4457645216 Test-only generated schema for the dashboard versioning proof of concept.
type GeomapBackendRegionsRequest20270101OneOf4457645216 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ProfilemetricsQuery LogQueryDefinition `json:"profilemetrics_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *GeomapRequestStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGeomapBackendRegionsRequest20270101OneOf4457645216 instantiates a new GeomapBackendRegionsRequest20270101OneOf4457645216 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeomapBackendRegionsRequest20270101OneOf4457645216(profilemetricsQuery LogQueryDefinition) *GeomapBackendRegionsRequest20270101OneOf4457645216 {
	this := GeomapBackendRegionsRequest20270101OneOf4457645216{}
	this.ProfilemetricsQuery = profilemetricsQuery
	return &this
}

// NewGeomapBackendRegionsRequest20270101OneOf4457645216WithDefaults instantiates a new GeomapBackendRegionsRequest20270101OneOf4457645216 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeomapBackendRegionsRequest20270101OneOf4457645216WithDefaults() *GeomapBackendRegionsRequest20270101OneOf4457645216 {
	this := GeomapBackendRegionsRequest20270101OneOf4457645216{}
	return &this
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *GeomapBackendRegionsRequest20270101OneOf4457645216) GetConditionalFormats() []ThresholdFormattingRule {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457645216) GetConditionalFormatsOk() (*[]ThresholdFormattingRule, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457645216) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule and assigns it to the ConditionalFormats field.
func (o *GeomapBackendRegionsRequest20270101OneOf4457645216) SetConditionalFormats(v []ThresholdFormattingRule) {
	o.ConditionalFormats = v
}

// GetProfilemetricsQuery returns the ProfilemetricsQuery field value.
func (o *GeomapBackendRegionsRequest20270101OneOf4457645216) GetProfilemetricsQuery() LogQueryDefinition {
	if o == nil {
		var ret LogQueryDefinition
		return ret
	}
	return o.ProfilemetricsQuery
}

// GetProfilemetricsQueryOk returns a tuple with the ProfilemetricsQuery field value
// and a boolean to check if the value has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457645216) GetProfilemetricsQueryOk() (*LogQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfilemetricsQuery, true
}

// SetProfilemetricsQuery sets field value.
func (o *GeomapBackendRegionsRequest20270101OneOf4457645216) SetProfilemetricsQuery(v LogQueryDefinition) {
	o.ProfilemetricsQuery = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *GeomapBackendRegionsRequest20270101OneOf4457645216) GetStyle() GeomapRequestStyle {
	if o == nil || o.Style == nil {
		var ret GeomapRequestStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457645216) GetStyleOk() (*GeomapRequestStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457645216) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given GeomapRequestStyle and assigns it to the Style field.
func (o *GeomapBackendRegionsRequest20270101OneOf4457645216) SetStyle(v GeomapRequestStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GeomapBackendRegionsRequest20270101OneOf4457645216) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	toSerialize["profilemetrics_query"] = o.ProfilemetricsQuery
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GeomapBackendRegionsRequest20270101OneOf4457645216) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConditionalFormats  []ThresholdFormattingRule `json:"conditional_formats,omitempty"`
		ProfilemetricsQuery *LogQueryDefinition       `json:"profilemetrics_query"`
		Style               *GeomapRequestStyle       `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ProfilemetricsQuery == nil {
		return fmt.Errorf("required field profilemetrics_query missing")
	}

	hasInvalidField := false
	o.ConditionalFormats = all.ConditionalFormats
	if all.ProfilemetricsQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ProfilemetricsQuery = *all.ProfilemetricsQuery
	if all.Style != nil && all.Style.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Style = all.Style

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

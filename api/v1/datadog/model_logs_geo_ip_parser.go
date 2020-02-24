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

// LogsGeoIPParser The GeoIP parser takes an IP address attribute and extracts if available the Continent, Country, Subdivision, and City information in the target attribute path.
type LogsGeoIPParser struct {
	// Array of source attributes
	Sources []string `json:"sources"`
	// Name of the parent attribute that contains all the extracted details from the `sources`
	Target string `json:"target"`
	// Type of processor
	Type *string `json:"type,omitempty"`
	// Whether or not the processor is enabled
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Name of the processor
	Name *string `json:"name,omitempty"`
}

// NewLogsGeoIPParser instantiates a new LogsGeoIPParser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsGeoIPParser(sources []string, target string) *LogsGeoIPParser {
	this := LogsGeoIPParser{}
	this.Sources = sources
	this.Target = target
	var type_ string = "geo-ip-parser"
	this.Type = &type_
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	return &this
}

// NewLogsGeoIPParserWithDefaults instantiates a new LogsGeoIPParser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsGeoIPParserWithDefaults() *LogsGeoIPParser {
	this := LogsGeoIPParser{}
	var target string = "network.client.geoip"
	this.Target = target
	var type_ string = "geo-ip-parser"
	this.Type = &type_
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	return &this
}

// GetSources returns the Sources field value
func (o *LogsGeoIPParser) GetSources() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Sources
}

// SetSources sets field value
func (o *LogsGeoIPParser) SetSources(v []string) {
	o.Sources = v
}

// GetTarget returns the Target field value
func (o *LogsGeoIPParser) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// SetTarget sets field value
func (o *LogsGeoIPParser) SetTarget(v string) {
	o.Target = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LogsGeoIPParser) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LogsGeoIPParser) GetTypeOk() (string, bool) {
	if o == nil || o.Type == nil {
		var ret string
		return ret, false
	}
	return *o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LogsGeoIPParser) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LogsGeoIPParser) SetType(v string) {
	o.Type = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *LogsGeoIPParser) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LogsGeoIPParser) GetIsEnabledOk() (bool, bool) {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret, false
	}
	return *o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *LogsGeoIPParser) HasIsEnabled() bool {
	if o != nil && o.IsEnabled != nil {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *LogsGeoIPParser) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogsGeoIPParser) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LogsGeoIPParser) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogsGeoIPParser) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogsGeoIPParser) SetName(v string) {
	o.Name = &v
}

func (o LogsGeoIPParser) MarshalJSON() ([]byte, error) {
	//TODO: serialize parents?
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sources"] = o.Sources
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

// AsLogsProcessor wraps this instance of LogsGeoIPParser in LogsProcessor
func (s *LogsGeoIPParser) AsLogsProcessor() LogsProcessor {
	return LogsProcessor{LogsProcessorInterface: s}
}

type NullableLogsGeoIPParser struct {
	value *LogsGeoIPParser
	isSet bool
}

func (v NullableLogsGeoIPParser) Get() *LogsGeoIPParser {
	return v.value
}

func (v NullableLogsGeoIPParser) Set(val *LogsGeoIPParser) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsGeoIPParser) IsSet() bool {
	return v.isSet
}

func (v NullableLogsGeoIPParser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsGeoIPParser(val *LogsGeoIPParser) *NullableLogsGeoIPParser {
	return &NullableLogsGeoIPParser{value: val, isSet: true}
}

func (v NullableLogsGeoIPParser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsGeoIPParser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

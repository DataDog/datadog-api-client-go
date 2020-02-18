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

// MetricsQueryResponseUnit struct for MetricsQueryResponseUnit
type MetricsQueryResponseUnit struct {
	// Unit family, allows for conversion between units of the same family, for scaling.
	Family *string `json:"family,omitempty"`
	// Unit name
	Name *string `json:"name,omitempty"`
	// Plural form of the unit name
	Plural *string `json:"plural,omitempty"`
	// Factor for scaling between units of the same family.
	ScaleFactor *float64 `json:"scale_factor,omitempty"`
	// Abbreviation of the unit
	ShortName *string `json:"short_name,omitempty"`
}

// NewMetricsQueryResponseUnit instantiates a new MetricsQueryResponseUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsQueryResponseUnit() *MetricsQueryResponseUnit {
	this := MetricsQueryResponseUnit{}
	return &this
}

// NewMetricsQueryResponseUnitWithDefaults instantiates a new MetricsQueryResponseUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsQueryResponseUnitWithDefaults() *MetricsQueryResponseUnit {
	this := MetricsQueryResponseUnit{}
	return &this
}

// GetFamily returns the Family field value if set, zero value otherwise.
func (o *MetricsQueryResponseUnit) GetFamily() string {
	if o == nil || o.Family == nil {
		var ret string
		return ret
	}
	return *o.Family
}

// GetFamilyOk returns a tuple with the Family field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MetricsQueryResponseUnit) GetFamilyOk() (string, bool) {
	if o == nil || o.Family == nil {
		var ret string
		return ret, false
	}
	return *o.Family, true
}

// HasFamily returns a boolean if a field has been set.
func (o *MetricsQueryResponseUnit) HasFamily() bool {
	if o != nil && o.Family != nil {
		return true
	}

	return false
}

// SetFamily gets a reference to the given string and assigns it to the Family field.
func (o *MetricsQueryResponseUnit) SetFamily(v string) {
	o.Family = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MetricsQueryResponseUnit) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MetricsQueryResponseUnit) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MetricsQueryResponseUnit) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MetricsQueryResponseUnit) SetName(v string) {
	o.Name = &v
}

// GetPlural returns the Plural field value if set, zero value otherwise.
func (o *MetricsQueryResponseUnit) GetPlural() string {
	if o == nil || o.Plural == nil {
		var ret string
		return ret
	}
	return *o.Plural
}

// GetPluralOk returns a tuple with the Plural field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MetricsQueryResponseUnit) GetPluralOk() (string, bool) {
	if o == nil || o.Plural == nil {
		var ret string
		return ret, false
	}
	return *o.Plural, true
}

// HasPlural returns a boolean if a field has been set.
func (o *MetricsQueryResponseUnit) HasPlural() bool {
	if o != nil && o.Plural != nil {
		return true
	}

	return false
}

// SetPlural gets a reference to the given string and assigns it to the Plural field.
func (o *MetricsQueryResponseUnit) SetPlural(v string) {
	o.Plural = &v
}

// GetScaleFactor returns the ScaleFactor field value if set, zero value otherwise.
func (o *MetricsQueryResponseUnit) GetScaleFactor() float64 {
	if o == nil || o.ScaleFactor == nil {
		var ret float64
		return ret
	}
	return *o.ScaleFactor
}

// GetScaleFactorOk returns a tuple with the ScaleFactor field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MetricsQueryResponseUnit) GetScaleFactorOk() (float64, bool) {
	if o == nil || o.ScaleFactor == nil {
		var ret float64
		return ret, false
	}
	return *o.ScaleFactor, true
}

// HasScaleFactor returns a boolean if a field has been set.
func (o *MetricsQueryResponseUnit) HasScaleFactor() bool {
	if o != nil && o.ScaleFactor != nil {
		return true
	}

	return false
}

// SetScaleFactor gets a reference to the given float64 and assigns it to the ScaleFactor field.
func (o *MetricsQueryResponseUnit) SetScaleFactor(v float64) {
	o.ScaleFactor = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *MetricsQueryResponseUnit) GetShortName() string {
	if o == nil || o.ShortName == nil {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MetricsQueryResponseUnit) GetShortNameOk() (string, bool) {
	if o == nil || o.ShortName == nil {
		var ret string
		return ret, false
	}
	return *o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *MetricsQueryResponseUnit) HasShortName() bool {
	if o != nil && o.ShortName != nil {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *MetricsQueryResponseUnit) SetShortName(v string) {
	o.ShortName = &v
}

func (o MetricsQueryResponseUnit) MarshalJSON() ([]byte, error) {
	//TODO: serialize parents?
	toSerialize := map[string]interface{}{}
	if o.Family != nil {
		toSerialize["family"] = o.Family
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Plural != nil {
		toSerialize["plural"] = o.Plural
	}
	if o.ScaleFactor != nil {
		toSerialize["scale_factor"] = o.ScaleFactor
	}
	if o.ShortName != nil {
		toSerialize["short_name"] = o.ShortName
	}
	return json.Marshal(toSerialize)
}

type NullableMetricsQueryResponseUnit struct {
	value *MetricsQueryResponseUnit
	isSet bool
}

func (v NullableMetricsQueryResponseUnit) Get() *MetricsQueryResponseUnit {
	return v.value
}

func (v NullableMetricsQueryResponseUnit) Set(val *MetricsQueryResponseUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsQueryResponseUnit) IsSet() bool {
	return v.isSet
}

func (v NullableMetricsQueryResponseUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsQueryResponseUnit(val *MetricsQueryResponseUnit) *NullableMetricsQueryResponseUnit {
	return &NullableMetricsQueryResponseUnit{value: val, isSet: true}
}

func (v NullableMetricsQueryResponseUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsQueryResponseUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

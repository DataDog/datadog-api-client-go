// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstLegend20270101OneOf Test-only generated schema for the dashboard versioning proof of concept.
type SunburstLegend20270101OneOf struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	HidePercent *bool `json:"hide_percent,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	HideValue *bool `json:"hide_value,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type SunburstLegendAnyOfAnyOfItemType_20270101 `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSunburstLegend20270101OneOf instantiates a new SunburstLegend20270101OneOf object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSunburstLegend20270101OneOf(typeVar SunburstLegendAnyOfAnyOfItemType_20270101) *SunburstLegend20270101OneOf {
	this := SunburstLegend20270101OneOf{}
	this.Type = typeVar
	return &this
}

// NewSunburstLegend20270101OneOfWithDefaults instantiates a new SunburstLegend20270101OneOf object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSunburstLegend20270101OneOfWithDefaults() *SunburstLegend20270101OneOf {
	this := SunburstLegend20270101OneOf{}
	return &this
}

// GetHidePercent returns the HidePercent field value if set, zero value otherwise.
func (o *SunburstLegend20270101OneOf) GetHidePercent() bool {
	if o == nil || o.HidePercent == nil {
		var ret bool
		return ret
	}
	return *o.HidePercent
}

// GetHidePercentOk returns a tuple with the HidePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SunburstLegend20270101OneOf) GetHidePercentOk() (*bool, bool) {
	if o == nil || o.HidePercent == nil {
		return nil, false
	}
	return o.HidePercent, true
}

// HasHidePercent returns a boolean if a field has been set.
func (o *SunburstLegend20270101OneOf) HasHidePercent() bool {
	return o != nil && o.HidePercent != nil
}

// SetHidePercent gets a reference to the given bool and assigns it to the HidePercent field.
func (o *SunburstLegend20270101OneOf) SetHidePercent(v bool) {
	o.HidePercent = &v
}

// GetHideValue returns the HideValue field value if set, zero value otherwise.
func (o *SunburstLegend20270101OneOf) GetHideValue() bool {
	if o == nil || o.HideValue == nil {
		var ret bool
		return ret
	}
	return *o.HideValue
}

// GetHideValueOk returns a tuple with the HideValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SunburstLegend20270101OneOf) GetHideValueOk() (*bool, bool) {
	if o == nil || o.HideValue == nil {
		return nil, false
	}
	return o.HideValue, true
}

// HasHideValue returns a boolean if a field has been set.
func (o *SunburstLegend20270101OneOf) HasHideValue() bool {
	return o != nil && o.HideValue != nil
}

// SetHideValue gets a reference to the given bool and assigns it to the HideValue field.
func (o *SunburstLegend20270101OneOf) SetHideValue(v bool) {
	o.HideValue = &v
}

// GetType returns the Type field value.
func (o *SunburstLegend20270101OneOf) GetType() SunburstLegendAnyOfAnyOfItemType_20270101 {
	if o == nil {
		var ret SunburstLegendAnyOfAnyOfItemType_20270101
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SunburstLegend20270101OneOf) GetTypeOk() (*SunburstLegendAnyOfAnyOfItemType_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SunburstLegend20270101OneOf) SetType(v SunburstLegendAnyOfAnyOfItemType_20270101) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SunburstLegend20270101OneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.HidePercent != nil {
		toSerialize["hide_percent"] = o.HidePercent
	}
	if o.HideValue != nil {
		toSerialize["hide_value"] = o.HideValue
	}
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SunburstLegend20270101OneOf) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HidePercent *bool                                      `json:"hide_percent,omitempty"`
		HideValue   *bool                                      `json:"hide_value,omitempty"`
		Type        *SunburstLegendAnyOfAnyOfItemType_20270101 `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.HidePercent = all.HidePercent
	o.HideValue = all.HideValue
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstKnownLegend20270101OneOf4462657920 Test-only generated schema for the dashboard versioning proof of concept.
type SunburstKnownLegend20270101OneOf4462657920 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type SunburstKnownLegendAnyOfAnyOfItemType4452236352_20270101 `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSunburstKnownLegend20270101OneOf4462657920 instantiates a new SunburstKnownLegend20270101OneOf4462657920 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSunburstKnownLegend20270101OneOf4462657920(typeVar SunburstKnownLegendAnyOfAnyOfItemType4452236352_20270101) *SunburstKnownLegend20270101OneOf4462657920 {
	this := SunburstKnownLegend20270101OneOf4462657920{}
	this.Type = typeVar
	return &this
}

// NewSunburstKnownLegend20270101OneOf4462657920WithDefaults instantiates a new SunburstKnownLegend20270101OneOf4462657920 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSunburstKnownLegend20270101OneOf4462657920WithDefaults() *SunburstKnownLegend20270101OneOf4462657920 {
	this := SunburstKnownLegend20270101OneOf4462657920{}
	return &this
}

// GetType returns the Type field value.
func (o *SunburstKnownLegend20270101OneOf4462657920) GetType() SunburstKnownLegendAnyOfAnyOfItemType4452236352_20270101 {
	if o == nil {
		var ret SunburstKnownLegendAnyOfAnyOfItemType4452236352_20270101
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SunburstKnownLegend20270101OneOf4462657920) GetTypeOk() (*SunburstKnownLegendAnyOfAnyOfItemType4452236352_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SunburstKnownLegend20270101OneOf4462657920) SetType(v SunburstKnownLegendAnyOfAnyOfItemType4452236352_20270101) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SunburstKnownLegend20270101OneOf4462657920) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SunburstKnownLegend20270101OneOf4462657920) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type *SunburstKnownLegendAnyOfAnyOfItemType4452236352_20270101 `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
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

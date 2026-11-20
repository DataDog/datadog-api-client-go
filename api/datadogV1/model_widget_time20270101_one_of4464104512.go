// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetTime20270101OneOf4464104512 Test-only generated schema for the dashboard versioning proof of concept.
type WidgetTime20270101OneOf4464104512 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	HideIncompleteCostData *bool `json:"hide_incomplete_cost_data,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Offset float64 `json:"offset"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type WidgetCalendarAlignedSpanType_20270101 `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewWidgetTime20270101OneOf4464104512 instantiates a new WidgetTime20270101OneOf4464104512 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetTime20270101OneOf4464104512(offset float64, typeVar WidgetCalendarAlignedSpanType_20270101) *WidgetTime20270101OneOf4464104512 {
	this := WidgetTime20270101OneOf4464104512{}
	this.Offset = offset
	this.Type = typeVar
	return &this
}

// NewWidgetTime20270101OneOf4464104512WithDefaults instantiates a new WidgetTime20270101OneOf4464104512 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWidgetTime20270101OneOf4464104512WithDefaults() *WidgetTime20270101OneOf4464104512 {
	this := WidgetTime20270101OneOf4464104512{}
	return &this
}

// GetHideIncompleteCostData returns the HideIncompleteCostData field value if set, zero value otherwise.
func (o *WidgetTime20270101OneOf4464104512) GetHideIncompleteCostData() bool {
	if o == nil || o.HideIncompleteCostData == nil {
		var ret bool
		return ret
	}
	return *o.HideIncompleteCostData
}

// GetHideIncompleteCostDataOk returns a tuple with the HideIncompleteCostData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetTime20270101OneOf4464104512) GetHideIncompleteCostDataOk() (*bool, bool) {
	if o == nil || o.HideIncompleteCostData == nil {
		return nil, false
	}
	return o.HideIncompleteCostData, true
}

// HasHideIncompleteCostData returns a boolean if a field has been set.
func (o *WidgetTime20270101OneOf4464104512) HasHideIncompleteCostData() bool {
	return o != nil && o.HideIncompleteCostData != nil
}

// SetHideIncompleteCostData gets a reference to the given bool and assigns it to the HideIncompleteCostData field.
func (o *WidgetTime20270101OneOf4464104512) SetHideIncompleteCostData(v bool) {
	o.HideIncompleteCostData = &v
}

// GetOffset returns the Offset field value.
func (o *WidgetTime20270101OneOf4464104512) GetOffset() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *WidgetTime20270101OneOf4464104512) GetOffsetOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value.
func (o *WidgetTime20270101OneOf4464104512) SetOffset(v float64) {
	o.Offset = v
}

// GetType returns the Type field value.
func (o *WidgetTime20270101OneOf4464104512) GetType() WidgetCalendarAlignedSpanType_20270101 {
	if o == nil {
		var ret WidgetCalendarAlignedSpanType_20270101
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WidgetTime20270101OneOf4464104512) GetTypeOk() (*WidgetCalendarAlignedSpanType_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *WidgetTime20270101OneOf4464104512) SetType(v WidgetCalendarAlignedSpanType_20270101) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WidgetTime20270101OneOf4464104512) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.HideIncompleteCostData != nil {
		toSerialize["hide_incomplete_cost_data"] = o.HideIncompleteCostData
	}
	toSerialize["offset"] = o.Offset
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WidgetTime20270101OneOf4464104512) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HideIncompleteCostData *bool                                   `json:"hide_incomplete_cost_data,omitempty"`
		Offset                 *float64                                `json:"offset"`
		Type                   *WidgetCalendarAlignedSpanType_20270101 `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Offset == nil {
		return fmt.Errorf("required field offset missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.HideIncompleteCostData = all.HideIncompleteCostData
	o.Offset = *all.Offset
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

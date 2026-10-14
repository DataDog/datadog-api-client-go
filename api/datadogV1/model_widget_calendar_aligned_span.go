// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetCalendarAlignedSpan Used for calendar-aligned time spans, such as the current month or previous year.
type WidgetCalendarAlignedSpan struct {
	// Whether to hide incomplete cost data in the widget.
	HideIncompleteCostData *bool `json:"hide_incomplete_cost_data,omitempty"`
	// Number of completed periods before the current period. 0 represents the current period.
	Offset int64 `json:"offset"`
	// Time zone used to align the calendar period.
	Timezone *string `json:"timezone,omitempty"`
	// Calendar-aligned time span type.
	Type WidgetCalendarAlignedSpanType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWidgetCalendarAlignedSpan instantiates a new WidgetCalendarAlignedSpan object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetCalendarAlignedSpan(offset int64, typeVar WidgetCalendarAlignedSpanType) *WidgetCalendarAlignedSpan {
	this := WidgetCalendarAlignedSpan{}
	this.Offset = offset
	this.Type = typeVar
	return &this
}

// NewWidgetCalendarAlignedSpanWithDefaults instantiates a new WidgetCalendarAlignedSpan object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWidgetCalendarAlignedSpanWithDefaults() *WidgetCalendarAlignedSpan {
	this := WidgetCalendarAlignedSpan{}
	return &this
}

// GetHideIncompleteCostData returns the HideIncompleteCostData field value if set, zero value otherwise.
func (o *WidgetCalendarAlignedSpan) GetHideIncompleteCostData() bool {
	if o == nil || o.HideIncompleteCostData == nil {
		var ret bool
		return ret
	}
	return *o.HideIncompleteCostData
}

// GetHideIncompleteCostDataOk returns a tuple with the HideIncompleteCostData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetCalendarAlignedSpan) GetHideIncompleteCostDataOk() (*bool, bool) {
	if o == nil || o.HideIncompleteCostData == nil {
		return nil, false
	}
	return o.HideIncompleteCostData, true
}

// HasHideIncompleteCostData returns a boolean if a field has been set.
func (o *WidgetCalendarAlignedSpan) HasHideIncompleteCostData() bool {
	return o != nil && o.HideIncompleteCostData != nil
}

// SetHideIncompleteCostData gets a reference to the given bool and assigns it to the HideIncompleteCostData field.
func (o *WidgetCalendarAlignedSpan) SetHideIncompleteCostData(v bool) {
	o.HideIncompleteCostData = &v
}

// GetOffset returns the Offset field value.
func (o *WidgetCalendarAlignedSpan) GetOffset() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *WidgetCalendarAlignedSpan) GetOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value.
func (o *WidgetCalendarAlignedSpan) SetOffset(v int64) {
	o.Offset = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *WidgetCalendarAlignedSpan) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetCalendarAlignedSpan) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *WidgetCalendarAlignedSpan) HasTimezone() bool {
	return o != nil && o.Timezone != nil
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *WidgetCalendarAlignedSpan) SetTimezone(v string) {
	o.Timezone = &v
}

// GetType returns the Type field value.
func (o *WidgetCalendarAlignedSpan) GetType() WidgetCalendarAlignedSpanType {
	if o == nil {
		var ret WidgetCalendarAlignedSpanType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WidgetCalendarAlignedSpan) GetTypeOk() (*WidgetCalendarAlignedSpanType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *WidgetCalendarAlignedSpan) SetType(v WidgetCalendarAlignedSpanType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WidgetCalendarAlignedSpan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.HideIncompleteCostData != nil {
		toSerialize["hide_incomplete_cost_data"] = o.HideIncompleteCostData
	}
	toSerialize["offset"] = o.Offset
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WidgetCalendarAlignedSpan) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HideIncompleteCostData *bool                          `json:"hide_incomplete_cost_data,omitempty"`
		Offset                 *int64                         `json:"offset"`
		Timezone               *string                        `json:"timezone,omitempty"`
		Type                   *WidgetCalendarAlignedSpanType `json:"type"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"hide_incomplete_cost_data", "offset", "timezone", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.HideIncompleteCostData = all.HideIncompleteCostData
	o.Offset = *all.Offset
	o.Timezone = all.Timezone
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

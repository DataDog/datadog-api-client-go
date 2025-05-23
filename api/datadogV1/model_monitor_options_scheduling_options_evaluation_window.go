// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorOptionsSchedulingOptionsEvaluationWindow Configuration options for the evaluation window. If `hour_starts` is set, no other fields may be set. Otherwise, `day_starts` and `month_starts` must be set together.
type MonitorOptionsSchedulingOptionsEvaluationWindow struct {
	// The time of the day at which a one day cumulative evaluation window starts.
	DayStarts *string `json:"day_starts,omitempty"`
	// The minute of the hour at which a one hour cumulative evaluation window starts.
	HourStarts *int32 `json:"hour_starts,omitempty"`
	// The day of the month at which a one month cumulative evaluation window starts.
	MonthStarts *int32 `json:"month_starts,omitempty"`
	// The timezone of the time of the day of the cumulative evaluation window start.
	Timezone *string `json:"timezone,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorOptionsSchedulingOptionsEvaluationWindow instantiates a new MonitorOptionsSchedulingOptionsEvaluationWindow object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorOptionsSchedulingOptionsEvaluationWindow() *MonitorOptionsSchedulingOptionsEvaluationWindow {
	this := MonitorOptionsSchedulingOptionsEvaluationWindow{}
	return &this
}

// NewMonitorOptionsSchedulingOptionsEvaluationWindowWithDefaults instantiates a new MonitorOptionsSchedulingOptionsEvaluationWindow object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorOptionsSchedulingOptionsEvaluationWindowWithDefaults() *MonitorOptionsSchedulingOptionsEvaluationWindow {
	this := MonitorOptionsSchedulingOptionsEvaluationWindow{}
	return &this
}

// GetDayStarts returns the DayStarts field value if set, zero value otherwise.
func (o *MonitorOptionsSchedulingOptionsEvaluationWindow) GetDayStarts() string {
	if o == nil || o.DayStarts == nil {
		var ret string
		return ret
	}
	return *o.DayStarts
}

// GetDayStartsOk returns a tuple with the DayStarts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorOptionsSchedulingOptionsEvaluationWindow) GetDayStartsOk() (*string, bool) {
	if o == nil || o.DayStarts == nil {
		return nil, false
	}
	return o.DayStarts, true
}

// HasDayStarts returns a boolean if a field has been set.
func (o *MonitorOptionsSchedulingOptionsEvaluationWindow) HasDayStarts() bool {
	return o != nil && o.DayStarts != nil
}

// SetDayStarts gets a reference to the given string and assigns it to the DayStarts field.
func (o *MonitorOptionsSchedulingOptionsEvaluationWindow) SetDayStarts(v string) {
	o.DayStarts = &v
}

// GetHourStarts returns the HourStarts field value if set, zero value otherwise.
func (o *MonitorOptionsSchedulingOptionsEvaluationWindow) GetHourStarts() int32 {
	if o == nil || o.HourStarts == nil {
		var ret int32
		return ret
	}
	return *o.HourStarts
}

// GetHourStartsOk returns a tuple with the HourStarts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorOptionsSchedulingOptionsEvaluationWindow) GetHourStartsOk() (*int32, bool) {
	if o == nil || o.HourStarts == nil {
		return nil, false
	}
	return o.HourStarts, true
}

// HasHourStarts returns a boolean if a field has been set.
func (o *MonitorOptionsSchedulingOptionsEvaluationWindow) HasHourStarts() bool {
	return o != nil && o.HourStarts != nil
}

// SetHourStarts gets a reference to the given int32 and assigns it to the HourStarts field.
func (o *MonitorOptionsSchedulingOptionsEvaluationWindow) SetHourStarts(v int32) {
	o.HourStarts = &v
}

// GetMonthStarts returns the MonthStarts field value if set, zero value otherwise.
func (o *MonitorOptionsSchedulingOptionsEvaluationWindow) GetMonthStarts() int32 {
	if o == nil || o.MonthStarts == nil {
		var ret int32
		return ret
	}
	return *o.MonthStarts
}

// GetMonthStartsOk returns a tuple with the MonthStarts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorOptionsSchedulingOptionsEvaluationWindow) GetMonthStartsOk() (*int32, bool) {
	if o == nil || o.MonthStarts == nil {
		return nil, false
	}
	return o.MonthStarts, true
}

// HasMonthStarts returns a boolean if a field has been set.
func (o *MonitorOptionsSchedulingOptionsEvaluationWindow) HasMonthStarts() bool {
	return o != nil && o.MonthStarts != nil
}

// SetMonthStarts gets a reference to the given int32 and assigns it to the MonthStarts field.
func (o *MonitorOptionsSchedulingOptionsEvaluationWindow) SetMonthStarts(v int32) {
	o.MonthStarts = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *MonitorOptionsSchedulingOptionsEvaluationWindow) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorOptionsSchedulingOptionsEvaluationWindow) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *MonitorOptionsSchedulingOptionsEvaluationWindow) HasTimezone() bool {
	return o != nil && o.Timezone != nil
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *MonitorOptionsSchedulingOptionsEvaluationWindow) SetTimezone(v string) {
	o.Timezone = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorOptionsSchedulingOptionsEvaluationWindow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DayStarts != nil {
		toSerialize["day_starts"] = o.DayStarts
	}
	if o.HourStarts != nil {
		toSerialize["hour_starts"] = o.HourStarts
	}
	if o.MonthStarts != nil {
		toSerialize["month_starts"] = o.MonthStarts
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorOptionsSchedulingOptionsEvaluationWindow) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DayStarts   *string `json:"day_starts,omitempty"`
		HourStarts  *int32  `json:"hour_starts,omitempty"`
		MonthStarts *int32  `json:"month_starts,omitempty"`
		Timezone    *string `json:"timezone,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"day_starts", "hour_starts", "month_starts", "timezone"})
	} else {
		return err
	}
	o.DayStarts = all.DayStarts
	o.HourStarts = all.HourStarts
	o.MonthStarts = all.MonthStarts
	o.Timezone = all.Timezone

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

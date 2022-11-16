// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
)

// DashboardReportSchedule Report schedule-specific configuration that defines when and how often a report is sent.
type DashboardReportSchedule struct {
	// Enables or disables automatic report sending as configured in the report's schedule. If set to `false`, schedule is paused and reports are not sent.
	Active *bool `json:"active,omitempty"`
	// Frequency with which to send the report.
	Frequency *DashboardReportScheduleFrequency `json:"frequency,omitempty"`
	// ISO8601 formatted time (HH:MM) to send the report, using the time zone specified in the `timezone` field.
	RepeatAt *DashboardReportScheduleRepeatAt `json:"repeat_at,omitempty"`
	// Defines which day of the month to send the monthly report. If the frequency field uses months (for example, `1mo`), this field is required. If the frequency field does not use months (for example `1d`, `1w`), then this field must be `null` (if provided). If this field is not provided in an update request, and the update request changes the frequency to use days or weeks (for example, `1d`, `1w`), then this field is automatically set to `null`. This field is mutually exclusive with `repeat_on_day_of_week`, and cannot be defined in the same request that includes a non-null value for `repeat_on_day_of_week`.
	RepeatOnDayOfMonth NullableDashboardReportScheduleRepeatOnDayOfMonth `json:"repeat_on_day_of_month,omitempty"`
	// Defines which day of the week to send the weekly report. If the frequency field uses weeks (such as `1w`), this field is required and defines which day of the week to send the report. If the frequency field does not use weeks (for example, `1d`, `1mo`), then this field must be `null` (if provided). If this field is not provided in an update request, and the update request changes the frequency to use days or months (for example, `1d`, `1mo`), then this field is automatically set to `null`.  This field is mutually exclusive with `repeat_on_day_of_month` and cannot be defined in the same request that includes a non-null value for `repeat_on_day_of_month`.
	RepeatOnDayOfWeek NullableDashboardReportScheduleRepeatOnDayOfWeek `json:"repeat_on_day_of_week,omitempty"`
	// Time zone to use for repeat_at. Must be a valid Time Zone Database (https://www.iana.org/time-zones) name
	Timezone *string `json:"timezone,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewDashboardReportSchedule instantiates a new DashboardReportSchedule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardReportSchedule() *DashboardReportSchedule {
	this := DashboardReportSchedule{}
	return &this
}

// NewDashboardReportScheduleWithDefaults instantiates a new DashboardReportSchedule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardReportScheduleWithDefaults() *DashboardReportSchedule {
	this := DashboardReportSchedule{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *DashboardReportSchedule) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardReportSchedule) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *DashboardReportSchedule) HasActive() bool {
	return o != nil && o.Active != nil
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *DashboardReportSchedule) SetActive(v bool) {
	o.Active = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *DashboardReportSchedule) GetFrequency() DashboardReportScheduleFrequency {
	if o == nil || o.Frequency == nil {
		var ret DashboardReportScheduleFrequency
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardReportSchedule) GetFrequencyOk() (*DashboardReportScheduleFrequency, bool) {
	if o == nil || o.Frequency == nil {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *DashboardReportSchedule) HasFrequency() bool {
	return o != nil && o.Frequency != nil
}

// SetFrequency gets a reference to the given DashboardReportScheduleFrequency and assigns it to the Frequency field.
func (o *DashboardReportSchedule) SetFrequency(v DashboardReportScheduleFrequency) {
	o.Frequency = &v
}

// GetRepeatAt returns the RepeatAt field value if set, zero value otherwise.
func (o *DashboardReportSchedule) GetRepeatAt() DashboardReportScheduleRepeatAt {
	if o == nil || o.RepeatAt == nil {
		var ret DashboardReportScheduleRepeatAt
		return ret
	}
	return *o.RepeatAt
}

// GetRepeatAtOk returns a tuple with the RepeatAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardReportSchedule) GetRepeatAtOk() (*DashboardReportScheduleRepeatAt, bool) {
	if o == nil || o.RepeatAt == nil {
		return nil, false
	}
	return o.RepeatAt, true
}

// HasRepeatAt returns a boolean if a field has been set.
func (o *DashboardReportSchedule) HasRepeatAt() bool {
	return o != nil && o.RepeatAt != nil
}

// SetRepeatAt gets a reference to the given DashboardReportScheduleRepeatAt and assigns it to the RepeatAt field.
func (o *DashboardReportSchedule) SetRepeatAt(v DashboardReportScheduleRepeatAt) {
	o.RepeatAt = &v
}

// GetRepeatOnDayOfMonth returns the RepeatOnDayOfMonth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardReportSchedule) GetRepeatOnDayOfMonth() DashboardReportScheduleRepeatOnDayOfMonth {
	if o == nil || o.RepeatOnDayOfMonth.Get() == nil {
		var ret DashboardReportScheduleRepeatOnDayOfMonth
		return ret
	}
	return *o.RepeatOnDayOfMonth.Get()
}

// GetRepeatOnDayOfMonthOk returns a tuple with the RepeatOnDayOfMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardReportSchedule) GetRepeatOnDayOfMonthOk() (*DashboardReportScheduleRepeatOnDayOfMonth, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepeatOnDayOfMonth.Get(), o.RepeatOnDayOfMonth.IsSet()
}

// HasRepeatOnDayOfMonth returns a boolean if a field has been set.
func (o *DashboardReportSchedule) HasRepeatOnDayOfMonth() bool {
	return o != nil && o.RepeatOnDayOfMonth.IsSet()
}

// SetRepeatOnDayOfMonth gets a reference to the given NullableDashboardReportScheduleRepeatOnDayOfMonth and assigns it to the RepeatOnDayOfMonth field.
func (o *DashboardReportSchedule) SetRepeatOnDayOfMonth(v DashboardReportScheduleRepeatOnDayOfMonth) {
	o.RepeatOnDayOfMonth.Set(&v)
}

// SetRepeatOnDayOfMonthNil sets the value for RepeatOnDayOfMonth to be an explicit nil.
func (o *DashboardReportSchedule) SetRepeatOnDayOfMonthNil() {
	o.RepeatOnDayOfMonth.Set(nil)
}

// UnsetRepeatOnDayOfMonth ensures that no value is present for RepeatOnDayOfMonth, not even an explicit nil.
func (o *DashboardReportSchedule) UnsetRepeatOnDayOfMonth() {
	o.RepeatOnDayOfMonth.Unset()
}

// GetRepeatOnDayOfWeek returns the RepeatOnDayOfWeek field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardReportSchedule) GetRepeatOnDayOfWeek() DashboardReportScheduleRepeatOnDayOfWeek {
	if o == nil || o.RepeatOnDayOfWeek.Get() == nil {
		var ret DashboardReportScheduleRepeatOnDayOfWeek
		return ret
	}
	return *o.RepeatOnDayOfWeek.Get()
}

// GetRepeatOnDayOfWeekOk returns a tuple with the RepeatOnDayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardReportSchedule) GetRepeatOnDayOfWeekOk() (*DashboardReportScheduleRepeatOnDayOfWeek, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepeatOnDayOfWeek.Get(), o.RepeatOnDayOfWeek.IsSet()
}

// HasRepeatOnDayOfWeek returns a boolean if a field has been set.
func (o *DashboardReportSchedule) HasRepeatOnDayOfWeek() bool {
	return o != nil && o.RepeatOnDayOfWeek.IsSet()
}

// SetRepeatOnDayOfWeek gets a reference to the given NullableDashboardReportScheduleRepeatOnDayOfWeek and assigns it to the RepeatOnDayOfWeek field.
func (o *DashboardReportSchedule) SetRepeatOnDayOfWeek(v DashboardReportScheduleRepeatOnDayOfWeek) {
	o.RepeatOnDayOfWeek.Set(&v)
}

// SetRepeatOnDayOfWeekNil sets the value for RepeatOnDayOfWeek to be an explicit nil.
func (o *DashboardReportSchedule) SetRepeatOnDayOfWeekNil() {
	o.RepeatOnDayOfWeek.Set(nil)
}

// UnsetRepeatOnDayOfWeek ensures that no value is present for RepeatOnDayOfWeek, not even an explicit nil.
func (o *DashboardReportSchedule) UnsetRepeatOnDayOfWeek() {
	o.RepeatOnDayOfWeek.Unset()
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *DashboardReportSchedule) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardReportSchedule) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *DashboardReportSchedule) HasTimezone() bool {
	return o != nil && o.Timezone != nil
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *DashboardReportSchedule) SetTimezone(v string) {
	o.Timezone = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardReportSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Frequency != nil {
		toSerialize["frequency"] = o.Frequency
	}
	if o.RepeatAt != nil {
		toSerialize["repeat_at"] = o.RepeatAt
	}
	if o.RepeatOnDayOfMonth.IsSet() {
		toSerialize["repeat_on_day_of_month"] = o.RepeatOnDayOfMonth.Get()
	}
	if o.RepeatOnDayOfWeek.IsSet() {
		toSerialize["repeat_on_day_of_week"] = o.RepeatOnDayOfWeek.Get()
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardReportSchedule) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Active             *bool                                             `json:"active,omitempty"`
		Frequency          *DashboardReportScheduleFrequency                 `json:"frequency,omitempty"`
		RepeatAt           *DashboardReportScheduleRepeatAt                  `json:"repeat_at,omitempty"`
		RepeatOnDayOfMonth NullableDashboardReportScheduleRepeatOnDayOfMonth `json:"repeat_on_day_of_month,omitempty"`
		RepeatOnDayOfWeek  NullableDashboardReportScheduleRepeatOnDayOfWeek  `json:"repeat_on_day_of_week,omitempty"`
		Timezone           *string                                           `json:"timezone,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.Frequency; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.RepeatAt; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.RepeatOnDayOfMonth; v.Get() != nil && !v.Get().IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.RepeatOnDayOfWeek; v.Get() != nil && !v.Get().IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Active = all.Active
	o.Frequency = all.Frequency
	o.RepeatAt = all.RepeatAt
	o.RepeatOnDayOfMonth = all.RepeatOnDayOfMonth
	o.RepeatOnDayOfWeek = all.RepeatOnDayOfWeek
	o.Timezone = all.Timezone
	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetScheduleV2RecurrenceRule Defines the recurrence pattern for the schedule.
type FleetScheduleV2RecurrenceRule struct {
	// Days of the week when the schedule triggers. Valid values are
	// "Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun".
	DaysOfWeek []string `json:"days_of_week,omitempty"`
	// Interval between schedule runs in weeks. 1 means the schedule runs every week
	// on the specified days. Higher values repeat every N weeks.
	Interval *int64 `json:"interval,omitempty"`
	// Duration of the maintenance window in minutes.
	MaintenanceWindowDuration *int64 `json:"maintenance_window_duration,omitempty"`
	// Start time of the maintenance window in 24-hour clock format (HHMM).
	// Deployments are triggered at this time on the specified days.
	StartMaintenanceWindow *string `json:"start_maintenance_window,omitempty"`
	// Timezone in IANA Time Zone Database format.
	Timezone *string `json:"timezone,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetScheduleV2RecurrenceRule instantiates a new FleetScheduleV2RecurrenceRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetScheduleV2RecurrenceRule() *FleetScheduleV2RecurrenceRule {
	this := FleetScheduleV2RecurrenceRule{}
	return &this
}

// NewFleetScheduleV2RecurrenceRuleWithDefaults instantiates a new FleetScheduleV2RecurrenceRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetScheduleV2RecurrenceRuleWithDefaults() *FleetScheduleV2RecurrenceRule {
	this := FleetScheduleV2RecurrenceRule{}
	return &this
}

// GetDaysOfWeek returns the DaysOfWeek field value if set, zero value otherwise.
func (o *FleetScheduleV2RecurrenceRule) GetDaysOfWeek() []string {
	if o == nil || o.DaysOfWeek == nil {
		var ret []string
		return ret
	}
	return o.DaysOfWeek
}

// GetDaysOfWeekOk returns a tuple with the DaysOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleV2RecurrenceRule) GetDaysOfWeekOk() (*[]string, bool) {
	if o == nil || o.DaysOfWeek == nil {
		return nil, false
	}
	return &o.DaysOfWeek, true
}

// HasDaysOfWeek returns a boolean if a field has been set.
func (o *FleetScheduleV2RecurrenceRule) HasDaysOfWeek() bool {
	return o != nil && o.DaysOfWeek != nil
}

// SetDaysOfWeek gets a reference to the given []string and assigns it to the DaysOfWeek field.
func (o *FleetScheduleV2RecurrenceRule) SetDaysOfWeek(v []string) {
	o.DaysOfWeek = v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *FleetScheduleV2RecurrenceRule) GetInterval() int64 {
	if o == nil || o.Interval == nil {
		var ret int64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleV2RecurrenceRule) GetIntervalOk() (*int64, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *FleetScheduleV2RecurrenceRule) HasInterval() bool {
	return o != nil && o.Interval != nil
}

// SetInterval gets a reference to the given int64 and assigns it to the Interval field.
func (o *FleetScheduleV2RecurrenceRule) SetInterval(v int64) {
	o.Interval = &v
}

// GetMaintenanceWindowDuration returns the MaintenanceWindowDuration field value if set, zero value otherwise.
func (o *FleetScheduleV2RecurrenceRule) GetMaintenanceWindowDuration() int64 {
	if o == nil || o.MaintenanceWindowDuration == nil {
		var ret int64
		return ret
	}
	return *o.MaintenanceWindowDuration
}

// GetMaintenanceWindowDurationOk returns a tuple with the MaintenanceWindowDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleV2RecurrenceRule) GetMaintenanceWindowDurationOk() (*int64, bool) {
	if o == nil || o.MaintenanceWindowDuration == nil {
		return nil, false
	}
	return o.MaintenanceWindowDuration, true
}

// HasMaintenanceWindowDuration returns a boolean if a field has been set.
func (o *FleetScheduleV2RecurrenceRule) HasMaintenanceWindowDuration() bool {
	return o != nil && o.MaintenanceWindowDuration != nil
}

// SetMaintenanceWindowDuration gets a reference to the given int64 and assigns it to the MaintenanceWindowDuration field.
func (o *FleetScheduleV2RecurrenceRule) SetMaintenanceWindowDuration(v int64) {
	o.MaintenanceWindowDuration = &v
}

// GetStartMaintenanceWindow returns the StartMaintenanceWindow field value if set, zero value otherwise.
func (o *FleetScheduleV2RecurrenceRule) GetStartMaintenanceWindow() string {
	if o == nil || o.StartMaintenanceWindow == nil {
		var ret string
		return ret
	}
	return *o.StartMaintenanceWindow
}

// GetStartMaintenanceWindowOk returns a tuple with the StartMaintenanceWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleV2RecurrenceRule) GetStartMaintenanceWindowOk() (*string, bool) {
	if o == nil || o.StartMaintenanceWindow == nil {
		return nil, false
	}
	return o.StartMaintenanceWindow, true
}

// HasStartMaintenanceWindow returns a boolean if a field has been set.
func (o *FleetScheduleV2RecurrenceRule) HasStartMaintenanceWindow() bool {
	return o != nil && o.StartMaintenanceWindow != nil
}

// SetStartMaintenanceWindow gets a reference to the given string and assigns it to the StartMaintenanceWindow field.
func (o *FleetScheduleV2RecurrenceRule) SetStartMaintenanceWindow(v string) {
	o.StartMaintenanceWindow = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *FleetScheduleV2RecurrenceRule) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleV2RecurrenceRule) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *FleetScheduleV2RecurrenceRule) HasTimezone() bool {
	return o != nil && o.Timezone != nil
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *FleetScheduleV2RecurrenceRule) SetTimezone(v string) {
	o.Timezone = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetScheduleV2RecurrenceRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DaysOfWeek != nil {
		toSerialize["days_of_week"] = o.DaysOfWeek
	}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	if o.MaintenanceWindowDuration != nil {
		toSerialize["maintenance_window_duration"] = o.MaintenanceWindowDuration
	}
	if o.StartMaintenanceWindow != nil {
		toSerialize["start_maintenance_window"] = o.StartMaintenanceWindow
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
func (o *FleetScheduleV2RecurrenceRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DaysOfWeek                []string `json:"days_of_week,omitempty"`
		Interval                  *int64   `json:"interval,omitempty"`
		MaintenanceWindowDuration *int64   `json:"maintenance_window_duration,omitempty"`
		StartMaintenanceWindow    *string  `json:"start_maintenance_window,omitempty"`
		Timezone                  *string  `json:"timezone,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"days_of_week", "interval", "maintenance_window_duration", "start_maintenance_window", "timezone"})
	} else {
		return err
	}
	o.DaysOfWeek = all.DaysOfWeek
	o.Interval = all.Interval
	o.MaintenanceWindowDuration = all.MaintenanceWindowDuration
	o.StartMaintenanceWindow = all.StartMaintenanceWindow
	o.Timezone = all.Timezone

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

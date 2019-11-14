/*
 * (C) Datadog, Inc. 2019
 * All rights reserved
 * Licensed under a 3-clause BSD style license (see LICENSE)
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog_v1

import (
	"bytes"
	"encoding/json"
)

// Downtime struct for Downtime
type Downtime struct {
	Active       *bool                       `json:"active,omitempty"`
	Canceled     *NullableInt64              `json:"canceled,omitempty"`
	CreatorId    *int32                      `json:"creator_id,omitempty"`
	Disabled     *bool                       `json:"disabled,omitempty"`
	DowntimeType *int32                      `json:"downtime_type,omitempty"`
	End          *NullableInt64              `json:"end,omitempty"`
	Id           *int64                      `json:"id,omitempty"`
	Message      *string                     `json:"message,omitempty"`
	MonitorId    *NullableInt64              `json:"monitor_id,omitempty"`
	MonitorTags  *[]string                   `json:"monitor_tags,omitempty"`
	ParentId     *NullableInt32              `json:"parent_id,omitempty"`
	Recurrence   *NullableDowntimeRecurrence `json:"recurrence,omitempty"`
	Scope        *[]string                   `json:"scope,omitempty"`
	Start        *int64                      `json:"start,omitempty"`
	Timezone     *string                     `json:"timezone,omitempty"`
	UpdaterId    *int32                      `json:"updater_id,omitempty"`
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Downtime) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Downtime) GetActiveOk() (bool, bool) {
	if o == nil || o.Active == nil {
		var ret bool
		return ret, false
	}
	return *o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Downtime) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Downtime) SetActive(v bool) {
	o.Active = &v
}

// GetCanceled returns the Canceled field value if set, zero value otherwise.
func (o *Downtime) GetCanceled() NullableInt64 {
	if o == nil || o.Canceled == nil {
		var ret NullableInt64
		return ret
	}
	return *o.Canceled
}

// GetCanceledOk returns a tuple with the Canceled field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Downtime) GetCanceledOk() (NullableInt64, bool) {
	if o == nil || o.Canceled == nil {
		var ret NullableInt64
		return ret, false
	}
	return *o.Canceled, true
}

// HasCanceled returns a boolean if a field has been set.
func (o *Downtime) HasCanceled() bool {
	if o != nil && o.Canceled != nil {
		return true
	}

	return false
}

// SetCanceled gets a reference to the given NullableInt64 and assigns it to the Canceled field.
func (o *Downtime) SetCanceled(v NullableInt64) {
	o.Canceled = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *Downtime) GetCreatorId() int32 {
	if o == nil || o.CreatorId == nil {
		var ret int32
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Downtime) GetCreatorIdOk() (int32, bool) {
	if o == nil || o.CreatorId == nil {
		var ret int32
		return ret, false
	}
	return *o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *Downtime) HasCreatorId() bool {
	if o != nil && o.CreatorId != nil {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given int32 and assigns it to the CreatorId field.
func (o *Downtime) SetCreatorId(v int32) {
	o.CreatorId = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *Downtime) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Downtime) GetDisabledOk() (bool, bool) {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret, false
	}
	return *o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *Downtime) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *Downtime) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetDowntimeType returns the DowntimeType field value if set, zero value otherwise.
func (o *Downtime) GetDowntimeType() int32 {
	if o == nil || o.DowntimeType == nil {
		var ret int32
		return ret
	}
	return *o.DowntimeType
}

// GetDowntimeTypeOk returns a tuple with the DowntimeType field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Downtime) GetDowntimeTypeOk() (int32, bool) {
	if o == nil || o.DowntimeType == nil {
		var ret int32
		return ret, false
	}
	return *o.DowntimeType, true
}

// HasDowntimeType returns a boolean if a field has been set.
func (o *Downtime) HasDowntimeType() bool {
	if o != nil && o.DowntimeType != nil {
		return true
	}

	return false
}

// SetDowntimeType gets a reference to the given int32 and assigns it to the DowntimeType field.
func (o *Downtime) SetDowntimeType(v int32) {
	o.DowntimeType = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *Downtime) GetEnd() NullableInt64 {
	if o == nil || o.End == nil {
		var ret NullableInt64
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Downtime) GetEndOk() (NullableInt64, bool) {
	if o == nil || o.End == nil {
		var ret NullableInt64
		return ret, false
	}
	return *o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *Downtime) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given NullableInt64 and assigns it to the End field.
func (o *Downtime) SetEnd(v NullableInt64) {
	o.End = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Downtime) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Downtime) GetIdOk() (int64, bool) {
	if o == nil || o.Id == nil {
		var ret int64
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Downtime) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Downtime) SetId(v int64) {
	o.Id = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Downtime) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Downtime) GetMessageOk() (string, bool) {
	if o == nil || o.Message == nil {
		var ret string
		return ret, false
	}
	return *o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Downtime) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Downtime) SetMessage(v string) {
	o.Message = &v
}

// GetMonitorId returns the MonitorId field value if set, zero value otherwise.
func (o *Downtime) GetMonitorId() NullableInt64 {
	if o == nil || o.MonitorId == nil {
		var ret NullableInt64
		return ret
	}
	return *o.MonitorId
}

// GetMonitorIdOk returns a tuple with the MonitorId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Downtime) GetMonitorIdOk() (NullableInt64, bool) {
	if o == nil || o.MonitorId == nil {
		var ret NullableInt64
		return ret, false
	}
	return *o.MonitorId, true
}

// HasMonitorId returns a boolean if a field has been set.
func (o *Downtime) HasMonitorId() bool {
	if o != nil && o.MonitorId != nil {
		return true
	}

	return false
}

// SetMonitorId gets a reference to the given NullableInt64 and assigns it to the MonitorId field.
func (o *Downtime) SetMonitorId(v NullableInt64) {
	o.MonitorId = &v
}

// GetMonitorTags returns the MonitorTags field value if set, zero value otherwise.
func (o *Downtime) GetMonitorTags() []string {
	if o == nil || o.MonitorTags == nil {
		var ret []string
		return ret
	}
	return *o.MonitorTags
}

// GetMonitorTagsOk returns a tuple with the MonitorTags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Downtime) GetMonitorTagsOk() ([]string, bool) {
	if o == nil || o.MonitorTags == nil {
		var ret []string
		return ret, false
	}
	return *o.MonitorTags, true
}

// HasMonitorTags returns a boolean if a field has been set.
func (o *Downtime) HasMonitorTags() bool {
	if o != nil && o.MonitorTags != nil {
		return true
	}

	return false
}

// SetMonitorTags gets a reference to the given []string and assigns it to the MonitorTags field.
func (o *Downtime) SetMonitorTags(v []string) {
	o.MonitorTags = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *Downtime) GetParentId() NullableInt32 {
	if o == nil || o.ParentId == nil {
		var ret NullableInt32
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Downtime) GetParentIdOk() (NullableInt32, bool) {
	if o == nil || o.ParentId == nil {
		var ret NullableInt32
		return ret, false
	}
	return *o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *Downtime) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableInt32 and assigns it to the ParentId field.
func (o *Downtime) SetParentId(v NullableInt32) {
	o.ParentId = &v
}

// GetRecurrence returns the Recurrence field value if set, zero value otherwise.
func (o *Downtime) GetRecurrence() NullableDowntimeRecurrence {
	if o == nil || o.Recurrence == nil {
		var ret NullableDowntimeRecurrence
		return ret
	}
	return *o.Recurrence
}

// GetRecurrenceOk returns a tuple with the Recurrence field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Downtime) GetRecurrenceOk() (NullableDowntimeRecurrence, bool) {
	if o == nil || o.Recurrence == nil {
		var ret NullableDowntimeRecurrence
		return ret, false
	}
	return *o.Recurrence, true
}

// HasRecurrence returns a boolean if a field has been set.
func (o *Downtime) HasRecurrence() bool {
	if o != nil && o.Recurrence != nil {
		return true
	}

	return false
}

// SetRecurrence gets a reference to the given NullableDowntimeRecurrence and assigns it to the Recurrence field.
func (o *Downtime) SetRecurrence(v NullableDowntimeRecurrence) {
	o.Recurrence = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *Downtime) GetScope() []string {
	if o == nil || o.Scope == nil {
		var ret []string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Downtime) GetScopeOk() ([]string, bool) {
	if o == nil || o.Scope == nil {
		var ret []string
		return ret, false
	}
	return *o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *Downtime) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given []string and assigns it to the Scope field.
func (o *Downtime) SetScope(v []string) {
	o.Scope = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *Downtime) GetStart() int64 {
	if o == nil || o.Start == nil {
		var ret int64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Downtime) GetStartOk() (int64, bool) {
	if o == nil || o.Start == nil {
		var ret int64
		return ret, false
	}
	return *o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *Downtime) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int64 and assigns it to the Start field.
func (o *Downtime) SetStart(v int64) {
	o.Start = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *Downtime) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Downtime) GetTimezoneOk() (string, bool) {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret, false
	}
	return *o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *Downtime) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *Downtime) SetTimezone(v string) {
	o.Timezone = &v
}

// GetUpdaterId returns the UpdaterId field value if set, zero value otherwise.
func (o *Downtime) GetUpdaterId() int32 {
	if o == nil || o.UpdaterId == nil {
		var ret int32
		return ret
	}
	return *o.UpdaterId
}

// GetUpdaterIdOk returns a tuple with the UpdaterId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Downtime) GetUpdaterIdOk() (int32, bool) {
	if o == nil || o.UpdaterId == nil {
		var ret int32
		return ret, false
	}
	return *o.UpdaterId, true
}

// HasUpdaterId returns a boolean if a field has been set.
func (o *Downtime) HasUpdaterId() bool {
	if o != nil && o.UpdaterId != nil {
		return true
	}

	return false
}

// SetUpdaterId gets a reference to the given int32 and assigns it to the UpdaterId field.
func (o *Downtime) SetUpdaterId(v int32) {
	o.UpdaterId = &v
}

type NullableDowntime struct {
	Value        Downtime
	ExplicitNull bool
}

func (v NullableDowntime) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableDowntime) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

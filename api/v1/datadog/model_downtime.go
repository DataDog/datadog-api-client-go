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

// Downtime Downtiming gives you greater control over monitor notifications by allowing you to globally exclude scopes from alerting. Downtime settings, which can be scheduled with start and end times, prevent all alerting related to specified Datadog tags.
type Downtime struct {
	// If a scheduled downtime currently exists.
	Active *bool `json:"active,omitempty"`
	// If a scheduled downtime is cancelled.
	Canceled NullableInt64 `json:"canceled,omitempty"`
	// User ID of the downtime creator.
	CreatorId *int32 `json:"creator_id,omitempty"`
	// If a downtime has been disabled.
	Disabled *bool `json:"disabled,omitempty"`
	// 0 for a downtime applied on `*` or all; 1 when the downtime is only scoped to hosts; 2 when the downtime is scoped to anything but hosts.
	DowntimeType *int32 `json:"downtime_type,omitempty"`
	// POSIX timestamp to end the downtime. If not provided, the downtime is in effect indefinitely until you cancel it.
	End NullableInt64 `json:"end,omitempty"`
	// The downtime ID.
	Id *int64 `json:"id,omitempty"`
	// A message to include with notifications for this downtime. Email notifications can be sent to specific users by using the same ‘@username’ notation as events.
	Message *string `json:"message,omitempty"`
	// A single monitor to which the downtime applies. If not provided, the downtime applies to all monitors.
	MonitorId NullableInt64 `json:"monitor_id,omitempty"`
	// A comma-separated list of monitor tags. For example, tags that are applied directly to monitors, not tags that are used in monitor queries (which are filtered by the scope parameter), to which the downtime applies. The resulting downtime applies to monitors that match ALL provided monitor tags. For example, `service:postgres` **AND** `team:frontend`.
	MonitorTags *[]string `json:"monitor_tags,omitempty"`
	// ID of the parent Downtime.
	ParentId   NullableInt64              `json:"parent_id,omitempty"`
	Recurrence NullableDowntimeRecurrence `json:"recurrence,omitempty"`
	// The scope(s) to which the downtime applies. For example, `host:app2`. Provide multiple scopes as a comma-separated list like `env:dev,env:prod`. The resulting downtime applies to sources that matches ALL provided scopes (`env:dev` **AND** `env:prod`).
	Scope *[]string `json:"scope,omitempty"`
	// POSIX timestamp to start the downtime. If not provided, the downtime starts the moment it is created.
	Start *int64 `json:"start,omitempty"`
	// The timezone for the downtime.
	Timezone *string `json:"timezone,omitempty"`
	// ID of the last user that updated the downtime.
	UpdaterId NullableInt32 `json:"updater_id,omitempty"`
}

// NewDowntime instantiates a new Downtime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDowntime() *Downtime {
	this := Downtime{}
	return &this
}

// NewDowntimeWithDefaults instantiates a new Downtime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDowntimeWithDefaults() *Downtime {
	this := Downtime{}
	return &this
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
	if o == nil {
		var ret NullableInt64
		return ret
	}
	return o.Canceled
}

// GetCanceledOk returns a tuple with the Canceled field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Downtime) GetCanceledOk() (NullableInt64, bool) {
	if o == nil {
		var ret NullableInt64
		return ret, false
	}
	return o.Canceled, o.Canceled.IsSet()
}

// HasCanceled returns a boolean if a field has been set.
func (o *Downtime) HasCanceled() bool {
	if o != nil && o.Canceled.IsSet() {
		return true
	}

	return false
}

// SetCanceled gets a reference to the given NullableInt64 and assigns it to the Canceled field.
func (o *Downtime) SetCanceled(v NullableInt64) {
	o.Canceled = v
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
	if o == nil {
		var ret NullableInt64
		return ret
	}
	return o.End
}

// GetEndOk returns a tuple with the End field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Downtime) GetEndOk() (NullableInt64, bool) {
	if o == nil {
		var ret NullableInt64
		return ret, false
	}
	return o.End, o.End.IsSet()
}

// HasEnd returns a boolean if a field has been set.
func (o *Downtime) HasEnd() bool {
	if o != nil && o.End.IsSet() {
		return true
	}

	return false
}

// SetEnd gets a reference to the given NullableInt64 and assigns it to the End field.
func (o *Downtime) SetEnd(v NullableInt64) {
	o.End = v
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
	if o == nil {
		var ret NullableInt64
		return ret
	}
	return o.MonitorId
}

// GetMonitorIdOk returns a tuple with the MonitorId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Downtime) GetMonitorIdOk() (NullableInt64, bool) {
	if o == nil {
		var ret NullableInt64
		return ret, false
	}
	return o.MonitorId, o.MonitorId.IsSet()
}

// HasMonitorId returns a boolean if a field has been set.
func (o *Downtime) HasMonitorId() bool {
	if o != nil && o.MonitorId.IsSet() {
		return true
	}

	return false
}

// SetMonitorId gets a reference to the given NullableInt64 and assigns it to the MonitorId field.
func (o *Downtime) SetMonitorId(v NullableInt64) {
	o.MonitorId = v
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
func (o *Downtime) GetParentId() NullableInt64 {
	if o == nil {
		var ret NullableInt64
		return ret
	}
	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Downtime) GetParentIdOk() (NullableInt64, bool) {
	if o == nil {
		var ret NullableInt64
		return ret, false
	}
	return o.ParentId, o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *Downtime) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableInt64 and assigns it to the ParentId field.
func (o *Downtime) SetParentId(v NullableInt64) {
	o.ParentId = v
}

// GetRecurrence returns the Recurrence field value if set, zero value otherwise.
func (o *Downtime) GetRecurrence() NullableDowntimeRecurrence {
	if o == nil {
		var ret NullableDowntimeRecurrence
		return ret
	}
	return o.Recurrence
}

// GetRecurrenceOk returns a tuple with the Recurrence field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Downtime) GetRecurrenceOk() (NullableDowntimeRecurrence, bool) {
	if o == nil {
		var ret NullableDowntimeRecurrence
		return ret, false
	}
	return o.Recurrence, o.Recurrence.IsSet()
}

// HasRecurrence returns a boolean if a field has been set.
func (o *Downtime) HasRecurrence() bool {
	if o != nil && o.Recurrence.IsSet() {
		return true
	}

	return false
}

// SetRecurrence gets a reference to the given NullableDowntimeRecurrence and assigns it to the Recurrence field.
func (o *Downtime) SetRecurrence(v NullableDowntimeRecurrence) {
	o.Recurrence = v
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
func (o *Downtime) GetUpdaterId() NullableInt32 {
	if o == nil {
		var ret NullableInt32
		return ret
	}
	return o.UpdaterId
}

// GetUpdaterIdOk returns a tuple with the UpdaterId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Downtime) GetUpdaterIdOk() (NullableInt32, bool) {
	if o == nil {
		var ret NullableInt32
		return ret, false
	}
	return o.UpdaterId, o.UpdaterId.IsSet()
}

// HasUpdaterId returns a boolean if a field has been set.
func (o *Downtime) HasUpdaterId() bool {
	if o != nil && o.UpdaterId.IsSet() {
		return true
	}

	return false
}

// SetUpdaterId gets a reference to the given NullableInt32 and assigns it to the UpdaterId field.
func (o *Downtime) SetUpdaterId(v NullableInt32) {
	o.UpdaterId = v
}

func (o Downtime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Canceled.IsSet() {
		toSerialize["canceled"] = o.Canceled.Get()
	}
	if o.CreatorId != nil {
		toSerialize["creator_id"] = o.CreatorId
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.DowntimeType != nil {
		toSerialize["downtime_type"] = o.DowntimeType
	}
	if o.End.IsSet() {
		toSerialize["end"] = o.End.Get()
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.MonitorId.IsSet() {
		toSerialize["monitor_id"] = o.MonitorId.Get()
	}
	if o.MonitorTags != nil {
		toSerialize["monitor_tags"] = o.MonitorTags
	}
	if o.ParentId.IsSet() {
		toSerialize["parent_id"] = o.ParentId.Get()
	}
	if o.Recurrence.IsSet() {
		toSerialize["recurrence"] = o.Recurrence.Get()
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	if o.UpdaterId.IsSet() {
		toSerialize["updater_id"] = o.UpdaterId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDowntime struct {
	value *Downtime
	isSet bool
}

func (v NullableDowntime) Get() *Downtime {
	return v.value
}

func (v NullableDowntime) Set(val *Downtime) {
	v.value = val
	v.isSet = true
}

func (v NullableDowntime) IsSet() bool {
	return v.isSet
}

func (v NullableDowntime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDowntime(val *Downtime) *NullableDowntime {
	return &NullableDowntime{value: val, isSet: true}
}

func (v NullableDowntime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDowntime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

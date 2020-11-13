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

// LogsArchiveAttributes The attributes associated with the archive.
type LogsArchiveAttributes struct {
	Destination NullableLogsArchiveDestination `json:"destination"`
	// To store the tags in the archive, set the value \"true\". If it is set to \"false\", the tags will be deleted when the logs are sent to the archive.
	IncludeTags *bool `json:"include_tags,omitempty"`
	// The archive name.
	Name string `json:"name"`
	// The archive query/filter. Logs matching this query are included in the archive.
	Query string `json:"query"`
	// An array of tags to add to rehydrated logs from an archive.
	RehydrationTags *[]string         `json:"rehydration_tags,omitempty"`
	State           *LogsArchiveState `json:"state,omitempty"`
}

// NewLogsArchiveAttributes instantiates a new LogsArchiveAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsArchiveAttributes(destination NullableLogsArchiveDestination, name string, query string) *LogsArchiveAttributes {
	this := LogsArchiveAttributes{}
	this.Destination = destination
	var includeTags bool = false
	this.IncludeTags = &includeTags
	this.Name = name
	this.Query = query
	return &this
}

// NewLogsArchiveAttributesWithDefaults instantiates a new LogsArchiveAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsArchiveAttributesWithDefaults() *LogsArchiveAttributes {
	this := LogsArchiveAttributes{}
	var includeTags bool = false
	this.IncludeTags = &includeTags
	return &this
}

// GetDestination returns the Destination field value
// If the value is explicit nil, the zero value for LogsArchiveDestination will be returned
func (o *LogsArchiveAttributes) GetDestination() LogsArchiveDestination {
	if o == nil || o.Destination.Get() == nil {
		var ret LogsArchiveDestination
		return ret
	}

	return *o.Destination.Get()
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogsArchiveAttributes) GetDestinationOk() (*LogsArchiveDestination, bool) {
	if o == nil {
		return nil, false
	}
	return o.Destination.Get(), o.Destination.IsSet()
}

// SetDestination sets field value
func (o *LogsArchiveAttributes) SetDestination(v LogsArchiveDestination) {
	o.Destination.Set(&v)
}

// GetIncludeTags returns the IncludeTags field value if set, zero value otherwise.
func (o *LogsArchiveAttributes) GetIncludeTags() bool {
	if o == nil || o.IncludeTags == nil {
		var ret bool
		return ret
	}
	return *o.IncludeTags
}

// GetIncludeTagsOk returns a tuple with the IncludeTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArchiveAttributes) GetIncludeTagsOk() (*bool, bool) {
	if o == nil || o.IncludeTags == nil {
		return nil, false
	}
	return o.IncludeTags, true
}

// HasIncludeTags returns a boolean if a field has been set.
func (o *LogsArchiveAttributes) HasIncludeTags() bool {
	if o != nil && o.IncludeTags != nil {
		return true
	}

	return false
}

// SetIncludeTags gets a reference to the given bool and assigns it to the IncludeTags field.
func (o *LogsArchiveAttributes) SetIncludeTags(v bool) {
	o.IncludeTags = &v
}

// GetName returns the Name field value
func (o *LogsArchiveAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LogsArchiveAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LogsArchiveAttributes) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value
func (o *LogsArchiveAttributes) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *LogsArchiveAttributes) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *LogsArchiveAttributes) SetQuery(v string) {
	o.Query = v
}

// GetRehydrationTags returns the RehydrationTags field value if set, zero value otherwise.
func (o *LogsArchiveAttributes) GetRehydrationTags() []string {
	if o == nil || o.RehydrationTags == nil {
		var ret []string
		return ret
	}
	return *o.RehydrationTags
}

// GetRehydrationTagsOk returns a tuple with the RehydrationTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArchiveAttributes) GetRehydrationTagsOk() (*[]string, bool) {
	if o == nil || o.RehydrationTags == nil {
		return nil, false
	}
	return o.RehydrationTags, true
}

// HasRehydrationTags returns a boolean if a field has been set.
func (o *LogsArchiveAttributes) HasRehydrationTags() bool {
	if o != nil && o.RehydrationTags != nil {
		return true
	}

	return false
}

// SetRehydrationTags gets a reference to the given []string and assigns it to the RehydrationTags field.
func (o *LogsArchiveAttributes) SetRehydrationTags(v []string) {
	o.RehydrationTags = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *LogsArchiveAttributes) GetState() LogsArchiveState {
	if o == nil || o.State == nil {
		var ret LogsArchiveState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArchiveAttributes) GetStateOk() (*LogsArchiveState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *LogsArchiveAttributes) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given LogsArchiveState and assigns it to the State field.
func (o *LogsArchiveAttributes) SetState(v LogsArchiveState) {
	o.State = &v
}

func (o LogsArchiveAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["destination"] = o.Destination.Get()
	}
	if o.IncludeTags != nil {
		toSerialize["include_tags"] = o.IncludeTags
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["query"] = o.Query
	}
	if o.RehydrationTags != nil {
		toSerialize["rehydration_tags"] = o.RehydrationTags
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableLogsArchiveAttributes struct {
	value *LogsArchiveAttributes
	isSet bool
}

func (v NullableLogsArchiveAttributes) Get() *LogsArchiveAttributes {
	return v.value
}

func (v *NullableLogsArchiveAttributes) Set(val *LogsArchiveAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsArchiveAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsArchiveAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsArchiveAttributes(val *LogsArchiveAttributes) *NullableLogsArchiveAttributes {
	return &NullableLogsArchiveAttributes{value: val, isSet: true}
}

func (v NullableLogsArchiveAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsArchiveAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

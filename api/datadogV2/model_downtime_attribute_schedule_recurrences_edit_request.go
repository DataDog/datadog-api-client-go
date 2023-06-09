// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeAttributeScheduleRecurrencesEditRequest A recurring downtime schedule definition.
type DowntimeAttributeScheduleRecurrencesEditRequest struct {
	// A list of downtime recurrences.
	Recurrences []DowntimeAttributeScheduleRecurrenceCreateEditRequest `json:"recurrences,omitempty"`
	// The timezone in which to schedule the downtime.
	Timezone *string `json:"timezone,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewDowntimeAttributeScheduleRecurrencesEditRequest instantiates a new DowntimeAttributeScheduleRecurrencesEditRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDowntimeAttributeScheduleRecurrencesEditRequest() *DowntimeAttributeScheduleRecurrencesEditRequest {
	this := DowntimeAttributeScheduleRecurrencesEditRequest{}
	var timezone string = "UTC"
	this.Timezone = &timezone
	return &this
}

// NewDowntimeAttributeScheduleRecurrencesEditRequestWithDefaults instantiates a new DowntimeAttributeScheduleRecurrencesEditRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDowntimeAttributeScheduleRecurrencesEditRequestWithDefaults() *DowntimeAttributeScheduleRecurrencesEditRequest {
	this := DowntimeAttributeScheduleRecurrencesEditRequest{}
	var timezone string = "UTC"
	this.Timezone = &timezone
	return &this
}

// GetRecurrences returns the Recurrences field value if set, zero value otherwise.
func (o *DowntimeAttributeScheduleRecurrencesEditRequest) GetRecurrences() []DowntimeAttributeScheduleRecurrenceCreateEditRequest {
	if o == nil || o.Recurrences == nil {
		var ret []DowntimeAttributeScheduleRecurrenceCreateEditRequest
		return ret
	}
	return o.Recurrences
}

// GetRecurrencesOk returns a tuple with the Recurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeAttributeScheduleRecurrencesEditRequest) GetRecurrencesOk() (*[]DowntimeAttributeScheduleRecurrenceCreateEditRequest, bool) {
	if o == nil || o.Recurrences == nil {
		return nil, false
	}
	return &o.Recurrences, true
}

// HasRecurrences returns a boolean if a field has been set.
func (o *DowntimeAttributeScheduleRecurrencesEditRequest) HasRecurrences() bool {
	return o != nil && o.Recurrences != nil
}

// SetRecurrences gets a reference to the given []DowntimeAttributeScheduleRecurrenceCreateEditRequest and assigns it to the Recurrences field.
func (o *DowntimeAttributeScheduleRecurrencesEditRequest) SetRecurrences(v []DowntimeAttributeScheduleRecurrenceCreateEditRequest) {
	o.Recurrences = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *DowntimeAttributeScheduleRecurrencesEditRequest) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeAttributeScheduleRecurrencesEditRequest) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *DowntimeAttributeScheduleRecurrencesEditRequest) HasTimezone() bool {
	return o != nil && o.Timezone != nil
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *DowntimeAttributeScheduleRecurrencesEditRequest) SetTimezone(v string) {
	o.Timezone = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DowntimeAttributeScheduleRecurrencesEditRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Recurrences != nil {
		toSerialize["recurrences"] = o.Recurrences
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
func (o *DowntimeAttributeScheduleRecurrencesEditRequest) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Recurrences []DowntimeAttributeScheduleRecurrenceCreateEditRequest `json:"recurrences,omitempty"`
		Timezone    *string                                                `json:"timezone,omitempty"`
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
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"recurrences", "timezone"})
	} else {
		return err
	}
	o.Recurrences = all.Recurrences
	o.Timezone = all.Timezone
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeAttributeScheduleRecurrencesCreateRequest A recurring downtime schedule definition.
type DowntimeAttributeScheduleRecurrencesCreateRequest struct {
	// A list of downtime recurrences.
	Recurrences []DowntimeAttributeScheduleRecurrenceCreateEditRequest `json:"recurrences"`
	// The timezone in which to schedule the downtime.
	Timezone *string `json:"timezone,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewDowntimeAttributeScheduleRecurrencesCreateRequest instantiates a new DowntimeAttributeScheduleRecurrencesCreateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDowntimeAttributeScheduleRecurrencesCreateRequest(recurrences []DowntimeAttributeScheduleRecurrenceCreateEditRequest) *DowntimeAttributeScheduleRecurrencesCreateRequest {
	this := DowntimeAttributeScheduleRecurrencesCreateRequest{}
	this.Recurrences = recurrences
	var timezone string = "UTC"
	this.Timezone = &timezone
	return &this
}

// NewDowntimeAttributeScheduleRecurrencesCreateRequestWithDefaults instantiates a new DowntimeAttributeScheduleRecurrencesCreateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDowntimeAttributeScheduleRecurrencesCreateRequestWithDefaults() *DowntimeAttributeScheduleRecurrencesCreateRequest {
	this := DowntimeAttributeScheduleRecurrencesCreateRequest{}
	var timezone string = "UTC"
	this.Timezone = &timezone
	return &this
}

// GetRecurrences returns the Recurrences field value.
func (o *DowntimeAttributeScheduleRecurrencesCreateRequest) GetRecurrences() []DowntimeAttributeScheduleRecurrenceCreateEditRequest {
	if o == nil {
		var ret []DowntimeAttributeScheduleRecurrenceCreateEditRequest
		return ret
	}
	return o.Recurrences
}

// GetRecurrencesOk returns a tuple with the Recurrences field value
// and a boolean to check if the value has been set.
func (o *DowntimeAttributeScheduleRecurrencesCreateRequest) GetRecurrencesOk() (*[]DowntimeAttributeScheduleRecurrenceCreateEditRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recurrences, true
}

// SetRecurrences sets field value.
func (o *DowntimeAttributeScheduleRecurrencesCreateRequest) SetRecurrences(v []DowntimeAttributeScheduleRecurrenceCreateEditRequest) {
	o.Recurrences = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *DowntimeAttributeScheduleRecurrencesCreateRequest) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeAttributeScheduleRecurrencesCreateRequest) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *DowntimeAttributeScheduleRecurrencesCreateRequest) HasTimezone() bool {
	return o != nil && o.Timezone != nil
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *DowntimeAttributeScheduleRecurrencesCreateRequest) SetTimezone(v string) {
	o.Timezone = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DowntimeAttributeScheduleRecurrencesCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["recurrences"] = o.Recurrences
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DowntimeAttributeScheduleRecurrencesCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Recurrences *[]DowntimeAttributeScheduleRecurrenceCreateEditRequest `json:"recurrences"`
	}{}
	all := struct {
		Recurrences []DowntimeAttributeScheduleRecurrenceCreateEditRequest `json:"recurrences"`
		Timezone    *string                                                `json:"timezone,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Recurrences == nil {
		return fmt.Errorf("required field recurrences missing")
	}
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

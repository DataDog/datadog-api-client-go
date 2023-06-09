// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeAttributeScheduleRecurrencesResponse A recurring downtime schedule definition.
type DowntimeAttributeScheduleRecurrencesResponse struct {
	// The most recent actual start and end dates for a recurring downtime. For a canceled downtime,
	// this is the previously occurring downtime. For active downtimes, this is the ongoing downtime, and for scheduled
	// downtimes it is the upcoming downtime.
	CurrentDowntime *DowntimeAttributeScheduleCurrentDowntimeResponse `json:"current_downtime,omitempty"`
	// A list of downtime recurrences.
	Recurrences []DowntimeAttributeScheduleRecurrenceResponse `json:"recurrences"`
	// The timezone in which to schedule the downtime. This affects recurring start and end dates.
	// Must match `display_timezone`.
	Timezone *string `json:"timezone,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewDowntimeAttributeScheduleRecurrencesResponse instantiates a new DowntimeAttributeScheduleRecurrencesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDowntimeAttributeScheduleRecurrencesResponse(recurrences []DowntimeAttributeScheduleRecurrenceResponse) *DowntimeAttributeScheduleRecurrencesResponse {
	this := DowntimeAttributeScheduleRecurrencesResponse{}
	this.Recurrences = recurrences
	var timezone string = "UTC"
	this.Timezone = &timezone
	return &this
}

// NewDowntimeAttributeScheduleRecurrencesResponseWithDefaults instantiates a new DowntimeAttributeScheduleRecurrencesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDowntimeAttributeScheduleRecurrencesResponseWithDefaults() *DowntimeAttributeScheduleRecurrencesResponse {
	this := DowntimeAttributeScheduleRecurrencesResponse{}
	var timezone string = "UTC"
	this.Timezone = &timezone
	return &this
}

// GetCurrentDowntime returns the CurrentDowntime field value if set, zero value otherwise.
func (o *DowntimeAttributeScheduleRecurrencesResponse) GetCurrentDowntime() DowntimeAttributeScheduleCurrentDowntimeResponse {
	if o == nil || o.CurrentDowntime == nil {
		var ret DowntimeAttributeScheduleCurrentDowntimeResponse
		return ret
	}
	return *o.CurrentDowntime
}

// GetCurrentDowntimeOk returns a tuple with the CurrentDowntime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeAttributeScheduleRecurrencesResponse) GetCurrentDowntimeOk() (*DowntimeAttributeScheduleCurrentDowntimeResponse, bool) {
	if o == nil || o.CurrentDowntime == nil {
		return nil, false
	}
	return o.CurrentDowntime, true
}

// HasCurrentDowntime returns a boolean if a field has been set.
func (o *DowntimeAttributeScheduleRecurrencesResponse) HasCurrentDowntime() bool {
	return o != nil && o.CurrentDowntime != nil
}

// SetCurrentDowntime gets a reference to the given DowntimeAttributeScheduleCurrentDowntimeResponse and assigns it to the CurrentDowntime field.
func (o *DowntimeAttributeScheduleRecurrencesResponse) SetCurrentDowntime(v DowntimeAttributeScheduleCurrentDowntimeResponse) {
	o.CurrentDowntime = &v
}

// GetRecurrences returns the Recurrences field value.
func (o *DowntimeAttributeScheduleRecurrencesResponse) GetRecurrences() []DowntimeAttributeScheduleRecurrenceResponse {
	if o == nil {
		var ret []DowntimeAttributeScheduleRecurrenceResponse
		return ret
	}
	return o.Recurrences
}

// GetRecurrencesOk returns a tuple with the Recurrences field value
// and a boolean to check if the value has been set.
func (o *DowntimeAttributeScheduleRecurrencesResponse) GetRecurrencesOk() (*[]DowntimeAttributeScheduleRecurrenceResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recurrences, true
}

// SetRecurrences sets field value.
func (o *DowntimeAttributeScheduleRecurrencesResponse) SetRecurrences(v []DowntimeAttributeScheduleRecurrenceResponse) {
	o.Recurrences = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *DowntimeAttributeScheduleRecurrencesResponse) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeAttributeScheduleRecurrencesResponse) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *DowntimeAttributeScheduleRecurrencesResponse) HasTimezone() bool {
	return o != nil && o.Timezone != nil
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *DowntimeAttributeScheduleRecurrencesResponse) SetTimezone(v string) {
	o.Timezone = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DowntimeAttributeScheduleRecurrencesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.CurrentDowntime != nil {
		toSerialize["current_downtime"] = o.CurrentDowntime
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
func (o *DowntimeAttributeScheduleRecurrencesResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Recurrences *[]DowntimeAttributeScheduleRecurrenceResponse `json:"recurrences"`
	}{}
	all := struct {
		CurrentDowntime *DowntimeAttributeScheduleCurrentDowntimeResponse `json:"current_downtime,omitempty"`
		Recurrences     []DowntimeAttributeScheduleRecurrenceResponse     `json:"recurrences"`
		Timezone        *string                                           `json:"timezone,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"current_downtime", "recurrences", "timezone"})
	} else {
		return err
	}
	if all.CurrentDowntime != nil && all.CurrentDowntime.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.CurrentDowntime = all.CurrentDowntime
	o.Recurrences = all.Recurrences
	o.Timezone = all.Timezone
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

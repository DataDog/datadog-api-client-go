// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeAttributeScheduleCurrentDowntimeResponse The most recent actual start and end dates for a recurring downtime. For a canceled downtime,
// this is the previously occurring downtime. For active downtimes, this is the ongoing downtime, and for scheduled
// downtimes it is the upcoming downtime.
type DowntimeAttributeScheduleCurrentDowntimeResponse struct {
	// The end of the current downtime.
	End datadog.NullableTime `json:"end,omitempty"`
	// The start of the current downtime.
	Start *time.Time `json:"start,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewDowntimeAttributeScheduleCurrentDowntimeResponse instantiates a new DowntimeAttributeScheduleCurrentDowntimeResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDowntimeAttributeScheduleCurrentDowntimeResponse() *DowntimeAttributeScheduleCurrentDowntimeResponse {
	this := DowntimeAttributeScheduleCurrentDowntimeResponse{}
	return &this
}

// NewDowntimeAttributeScheduleCurrentDowntimeResponseWithDefaults instantiates a new DowntimeAttributeScheduleCurrentDowntimeResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDowntimeAttributeScheduleCurrentDowntimeResponseWithDefaults() *DowntimeAttributeScheduleCurrentDowntimeResponse {
	this := DowntimeAttributeScheduleCurrentDowntimeResponse{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DowntimeAttributeScheduleCurrentDowntimeResponse) GetEnd() time.Time {
	if o == nil || o.End.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.End.Get()
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DowntimeAttributeScheduleCurrentDowntimeResponse) GetEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.End.Get(), o.End.IsSet()
}

// HasEnd returns a boolean if a field has been set.
func (o *DowntimeAttributeScheduleCurrentDowntimeResponse) HasEnd() bool {
	return o != nil && o.End.IsSet()
}

// SetEnd gets a reference to the given datadog.NullableTime and assigns it to the End field.
func (o *DowntimeAttributeScheduleCurrentDowntimeResponse) SetEnd(v time.Time) {
	o.End.Set(&v)
}

// SetEndNil sets the value for End to be an explicit nil.
func (o *DowntimeAttributeScheduleCurrentDowntimeResponse) SetEndNil() {
	o.End.Set(nil)
}

// UnsetEnd ensures that no value is present for End, not even an explicit nil.
func (o *DowntimeAttributeScheduleCurrentDowntimeResponse) UnsetEnd() {
	o.End.Unset()
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *DowntimeAttributeScheduleCurrentDowntimeResponse) GetStart() time.Time {
	if o == nil || o.Start == nil {
		var ret time.Time
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeAttributeScheduleCurrentDowntimeResponse) GetStartOk() (*time.Time, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *DowntimeAttributeScheduleCurrentDowntimeResponse) HasStart() bool {
	return o != nil && o.Start != nil
}

// SetStart gets a reference to the given time.Time and assigns it to the Start field.
func (o *DowntimeAttributeScheduleCurrentDowntimeResponse) SetStart(v time.Time) {
	o.Start = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DowntimeAttributeScheduleCurrentDowntimeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.End.IsSet() {
		toSerialize["end"] = o.End.Get()
	}
	if o.Start != nil {
		if o.Start.Nanosecond() == 0 {
			toSerialize["start"] = o.Start.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["start"] = o.Start.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DowntimeAttributeScheduleCurrentDowntimeResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		End   datadog.NullableTime `json:"end,omitempty"`
		Start *time.Time           `json:"start,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"end", "start"})
	} else {
		return err
	}
	o.End = all.End
	o.Start = all.Start
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

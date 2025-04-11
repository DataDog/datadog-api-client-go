// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleUpdateRequestDataAttributesLayersItemsInterval Specifies how the rotation repeats: number of days, plus optional seconds, up to the given maximums.
type ScheduleUpdateRequestDataAttributesLayersItemsInterval struct {
	// How many days each rotation cycle should span.
	Days *int32 `json:"days,omitempty"`
	// Additional seconds to add to the rotation cycle (for example, partial days).
	Seconds *int64 `json:"seconds,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduleUpdateRequestDataAttributesLayersItemsInterval instantiates a new ScheduleUpdateRequestDataAttributesLayersItemsInterval object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleUpdateRequestDataAttributesLayersItemsInterval() *ScheduleUpdateRequestDataAttributesLayersItemsInterval {
	this := ScheduleUpdateRequestDataAttributesLayersItemsInterval{}
	return &this
}

// NewScheduleUpdateRequestDataAttributesLayersItemsIntervalWithDefaults instantiates a new ScheduleUpdateRequestDataAttributesLayersItemsInterval object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleUpdateRequestDataAttributesLayersItemsIntervalWithDefaults() *ScheduleUpdateRequestDataAttributesLayersItemsInterval {
	this := ScheduleUpdateRequestDataAttributesLayersItemsInterval{}
	return &this
}

// GetDays returns the Days field value if set, zero value otherwise.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsInterval) GetDays() int32 {
	if o == nil || o.Days == nil {
		var ret int32
		return ret
	}
	return *o.Days
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsInterval) GetDaysOk() (*int32, bool) {
	if o == nil || o.Days == nil {
		return nil, false
	}
	return o.Days, true
}

// HasDays returns a boolean if a field has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsInterval) HasDays() bool {
	return o != nil && o.Days != nil
}

// SetDays gets a reference to the given int32 and assigns it to the Days field.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsInterval) SetDays(v int32) {
	o.Days = &v
}

// GetSeconds returns the Seconds field value if set, zero value otherwise.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsInterval) GetSeconds() int64 {
	if o == nil || o.Seconds == nil {
		var ret int64
		return ret
	}
	return *o.Seconds
}

// GetSecondsOk returns a tuple with the Seconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsInterval) GetSecondsOk() (*int64, bool) {
	if o == nil || o.Seconds == nil {
		return nil, false
	}
	return o.Seconds, true
}

// HasSeconds returns a boolean if a field has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsInterval) HasSeconds() bool {
	return o != nil && o.Seconds != nil
}

// SetSeconds gets a reference to the given int64 and assigns it to the Seconds field.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsInterval) SetSeconds(v int64) {
	o.Seconds = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleUpdateRequestDataAttributesLayersItemsInterval) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Days != nil {
		toSerialize["days"] = o.Days
	}
	if o.Seconds != nil {
		toSerialize["seconds"] = o.Seconds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsInterval) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Days    *int32 `json:"days,omitempty"`
		Seconds *int64 `json:"seconds,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"days", "seconds"})
	} else {
		return err
	}
	o.Days = all.Days
	o.Seconds = all.Seconds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

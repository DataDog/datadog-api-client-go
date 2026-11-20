// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesXAxisDateFormat Test-only generated schema for the dashboard versioning proof of concept.
type TimeseriesXAxisDateFormat struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Day *TimeseriesXAxisDayDateFormat `json:"day,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Month *TimeseriesXAxisMonthDateFormat `json:"month,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTimeseriesXAxisDateFormat instantiates a new TimeseriesXAxisDateFormat object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesXAxisDateFormat() *TimeseriesXAxisDateFormat {
	this := TimeseriesXAxisDateFormat{}
	return &this
}

// NewTimeseriesXAxisDateFormatWithDefaults instantiates a new TimeseriesXAxisDateFormat object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesXAxisDateFormatWithDefaults() *TimeseriesXAxisDateFormat {
	this := TimeseriesXAxisDateFormat{}
	return &this
}

// GetDay returns the Day field value if set, zero value otherwise.
func (o *TimeseriesXAxisDateFormat) GetDay() TimeseriesXAxisDayDateFormat {
	if o == nil || o.Day == nil {
		var ret TimeseriesXAxisDayDateFormat
		return ret
	}
	return *o.Day
}

// GetDayOk returns a tuple with the Day field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesXAxisDateFormat) GetDayOk() (*TimeseriesXAxisDayDateFormat, bool) {
	if o == nil || o.Day == nil {
		return nil, false
	}
	return o.Day, true
}

// HasDay returns a boolean if a field has been set.
func (o *TimeseriesXAxisDateFormat) HasDay() bool {
	return o != nil && o.Day != nil
}

// SetDay gets a reference to the given TimeseriesXAxisDayDateFormat and assigns it to the Day field.
func (o *TimeseriesXAxisDateFormat) SetDay(v TimeseriesXAxisDayDateFormat) {
	o.Day = &v
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *TimeseriesXAxisDateFormat) GetMonth() TimeseriesXAxisMonthDateFormat {
	if o == nil || o.Month == nil {
		var ret TimeseriesXAxisMonthDateFormat
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesXAxisDateFormat) GetMonthOk() (*TimeseriesXAxisMonthDateFormat, bool) {
	if o == nil || o.Month == nil {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *TimeseriesXAxisDateFormat) HasMonth() bool {
	return o != nil && o.Month != nil
}

// SetMonth gets a reference to the given TimeseriesXAxisMonthDateFormat and assigns it to the Month field.
func (o *TimeseriesXAxisDateFormat) SetMonth(v TimeseriesXAxisMonthDateFormat) {
	o.Month = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesXAxisDateFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Day != nil {
		toSerialize["day"] = o.Day
	}
	if o.Month != nil {
		toSerialize["month"] = o.Month
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesXAxisDateFormat) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Day   *TimeseriesXAxisDayDateFormat   `json:"day,omitempty"`
		Month *TimeseriesXAxisMonthDateFormat `json:"month,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	if all.Day != nil && !all.Day.IsValid() {
		hasInvalidField = true
	} else {
		o.Day = all.Day
	}
	if all.Month != nil && !all.Month.IsValid() {
		hasInvalidField = true
	} else {
		o.Month = all.Month
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

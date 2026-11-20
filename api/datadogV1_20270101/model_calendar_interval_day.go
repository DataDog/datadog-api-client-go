// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CalendarIntervalDay Test-only generated schema for the dashboard versioning proof of concept.
type CalendarIntervalDay struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Alignment *Hour `json:"alignment,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Quantity *float64 `json:"quantity,omitempty"`
	// IANA time zone identifier, for example America/New_York.
	Timezone *string `json:"timezone,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type CalendarIntervalDayType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCalendarIntervalDay instantiates a new CalendarIntervalDay object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCalendarIntervalDay(typeVar CalendarIntervalDayType) *CalendarIntervalDay {
	this := CalendarIntervalDay{}
	this.Type = typeVar
	return &this
}

// NewCalendarIntervalDayWithDefaults instantiates a new CalendarIntervalDay object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCalendarIntervalDayWithDefaults() *CalendarIntervalDay {
	this := CalendarIntervalDay{}
	return &this
}

// GetAlignment returns the Alignment field value if set, zero value otherwise.
func (o *CalendarIntervalDay) GetAlignment() Hour {
	if o == nil || o.Alignment == nil {
		var ret Hour
		return ret
	}
	return *o.Alignment
}

// GetAlignmentOk returns a tuple with the Alignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarIntervalDay) GetAlignmentOk() (*Hour, bool) {
	if o == nil || o.Alignment == nil {
		return nil, false
	}
	return o.Alignment, true
}

// HasAlignment returns a boolean if a field has been set.
func (o *CalendarIntervalDay) HasAlignment() bool {
	return o != nil && o.Alignment != nil
}

// SetAlignment gets a reference to the given Hour and assigns it to the Alignment field.
func (o *CalendarIntervalDay) SetAlignment(v Hour) {
	o.Alignment = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *CalendarIntervalDay) GetQuantity() float64 {
	if o == nil || o.Quantity == nil {
		var ret float64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarIntervalDay) GetQuantityOk() (*float64, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *CalendarIntervalDay) HasQuantity() bool {
	return o != nil && o.Quantity != nil
}

// SetQuantity gets a reference to the given float64 and assigns it to the Quantity field.
func (o *CalendarIntervalDay) SetQuantity(v float64) {
	o.Quantity = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *CalendarIntervalDay) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarIntervalDay) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *CalendarIntervalDay) HasTimezone() bool {
	return o != nil && o.Timezone != nil
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *CalendarIntervalDay) SetTimezone(v string) {
	o.Timezone = &v
}

// GetType returns the Type field value.
func (o *CalendarIntervalDay) GetType() CalendarIntervalDayType {
	if o == nil {
		var ret CalendarIntervalDayType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CalendarIntervalDay) GetTypeOk() (*CalendarIntervalDayType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CalendarIntervalDay) SetType(v CalendarIntervalDayType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CalendarIntervalDay) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alignment != nil {
		toSerialize["alignment"] = o.Alignment
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CalendarIntervalDay) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alignment *Hour                    `json:"alignment,omitempty"`
		Quantity  *float64                 `json:"quantity,omitempty"`
		Timezone  *string                  `json:"timezone,omitempty"`
		Type      *CalendarIntervalDayType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Alignment = all.Alignment
	o.Quantity = all.Quantity
	o.Timezone = all.Timezone
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

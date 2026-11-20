// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CalendarIntervalMonth Test-only generated schema for the dashboard versioning proof of concept.
type CalendarIntervalMonth struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Alignment *DayOfTheMonth `json:"alignment,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Quantity *float64 `json:"quantity,omitempty"`
	// IANA time zone identifier, for example America/New_York.
	Timezone *string `json:"timezone,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type CalendarIntervalMonthType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCalendarIntervalMonth instantiates a new CalendarIntervalMonth object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCalendarIntervalMonth(typeVar CalendarIntervalMonthType) *CalendarIntervalMonth {
	this := CalendarIntervalMonth{}
	this.Type = typeVar
	return &this
}

// NewCalendarIntervalMonthWithDefaults instantiates a new CalendarIntervalMonth object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCalendarIntervalMonthWithDefaults() *CalendarIntervalMonth {
	this := CalendarIntervalMonth{}
	return &this
}

// GetAlignment returns the Alignment field value if set, zero value otherwise.
func (o *CalendarIntervalMonth) GetAlignment() DayOfTheMonth {
	if o == nil || o.Alignment == nil {
		var ret DayOfTheMonth
		return ret
	}
	return *o.Alignment
}

// GetAlignmentOk returns a tuple with the Alignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarIntervalMonth) GetAlignmentOk() (*DayOfTheMonth, bool) {
	if o == nil || o.Alignment == nil {
		return nil, false
	}
	return o.Alignment, true
}

// HasAlignment returns a boolean if a field has been set.
func (o *CalendarIntervalMonth) HasAlignment() bool {
	return o != nil && o.Alignment != nil
}

// SetAlignment gets a reference to the given DayOfTheMonth and assigns it to the Alignment field.
func (o *CalendarIntervalMonth) SetAlignment(v DayOfTheMonth) {
	o.Alignment = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *CalendarIntervalMonth) GetQuantity() float64 {
	if o == nil || o.Quantity == nil {
		var ret float64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarIntervalMonth) GetQuantityOk() (*float64, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *CalendarIntervalMonth) HasQuantity() bool {
	return o != nil && o.Quantity != nil
}

// SetQuantity gets a reference to the given float64 and assigns it to the Quantity field.
func (o *CalendarIntervalMonth) SetQuantity(v float64) {
	o.Quantity = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *CalendarIntervalMonth) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarIntervalMonth) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *CalendarIntervalMonth) HasTimezone() bool {
	return o != nil && o.Timezone != nil
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *CalendarIntervalMonth) SetTimezone(v string) {
	o.Timezone = &v
}

// GetType returns the Type field value.
func (o *CalendarIntervalMonth) GetType() CalendarIntervalMonthType {
	if o == nil {
		var ret CalendarIntervalMonthType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CalendarIntervalMonth) GetTypeOk() (*CalendarIntervalMonthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CalendarIntervalMonth) SetType(v CalendarIntervalMonthType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CalendarIntervalMonth) MarshalJSON() ([]byte, error) {
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
func (o *CalendarIntervalMonth) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alignment *DayOfTheMonth             `json:"alignment,omitempty"`
		Quantity  *float64                   `json:"quantity,omitempty"`
		Timezone  *string                    `json:"timezone,omitempty"`
		Type      *CalendarIntervalMonthType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	if all.Alignment != nil && !all.Alignment.IsValid() {
		hasInvalidField = true
	} else {
		o.Alignment = all.Alignment
	}
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

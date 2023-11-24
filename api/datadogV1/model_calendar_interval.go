// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CalendarInterval Calendar interval options for compute.
// Fields `interval` (numeric interval) and `rollup` (calendar interval) are mutually exclusive.
//
// For instance:
// - { type: 'day', alignment: '1pm', timezone: 'Europe/Paris' }
// - { type: 'week', alignment: 'tuesday', quantity: 2 }
// - { type: 'month', alignment: '15th' }
// - { type: 'year', alignment: 'april' }
type CalendarInterval struct {
	// Optional alignment to define period start.
	// Its possible values depend on the type field:
	//
	// - day: start hour of day as 12 or 24-hr format (for instance: 11pm, 3am, 15)
	// - week: first day of the week (for instance: tuesday, note the lowercase)
	// - month: first day of month (for instance: 1th, 2nd, 23th)
	// - year: first month of the year (for instance: april, note the lowercase)
	Alignment *string `json:"alignment,omitempty"`
	// Optional integer to specify how many units to group together.
	Quantity *int64 `json:"quantity,omitempty"`
	// Optional timezone to define the calendar interval.
	Timezone *string `json:"timezone,omitempty"`
	// Type of calendar interval (day, week, etc.).
	Type CalendarIntervalType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewCalendarInterval instantiates a new CalendarInterval object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCalendarInterval(typeVar CalendarIntervalType) *CalendarInterval {
	this := CalendarInterval{}
	var quantity int64 = 1
	this.Quantity = &quantity
	var timezone string = "UTC"
	this.Timezone = &timezone
	this.Type = typeVar
	return &this
}

// NewCalendarIntervalWithDefaults instantiates a new CalendarInterval object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCalendarIntervalWithDefaults() *CalendarInterval {
	this := CalendarInterval{}
	var quantity int64 = 1
	this.Quantity = &quantity
	var timezone string = "UTC"
	this.Timezone = &timezone
	return &this
}

// GetAlignment returns the Alignment field value if set, zero value otherwise.
func (o *CalendarInterval) GetAlignment() string {
	if o == nil || o.Alignment == nil {
		var ret string
		return ret
	}
	return *o.Alignment
}

// GetAlignmentOk returns a tuple with the Alignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarInterval) GetAlignmentOk() (*string, bool) {
	if o == nil || o.Alignment == nil {
		return nil, false
	}
	return o.Alignment, true
}

// HasAlignment returns a boolean if a field has been set.
func (o *CalendarInterval) HasAlignment() bool {
	return o != nil && o.Alignment != nil
}

// SetAlignment gets a reference to the given string and assigns it to the Alignment field.
func (o *CalendarInterval) SetAlignment(v string) {
	o.Alignment = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *CalendarInterval) GetQuantity() int64 {
	if o == nil || o.Quantity == nil {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarInterval) GetQuantityOk() (*int64, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *CalendarInterval) HasQuantity() bool {
	return o != nil && o.Quantity != nil
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *CalendarInterval) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *CalendarInterval) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarInterval) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *CalendarInterval) HasTimezone() bool {
	return o != nil && o.Timezone != nil
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *CalendarInterval) SetTimezone(v string) {
	o.Timezone = &v
}

// GetType returns the Type field value.
func (o *CalendarInterval) GetType() CalendarIntervalType {
	if o == nil {
		var ret CalendarIntervalType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CalendarInterval) GetTypeOk() (*CalendarIntervalType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CalendarInterval) SetType(v CalendarIntervalType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CalendarInterval) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CalendarInterval) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alignment *string               `json:"alignment,omitempty"`
		Quantity  *int64                `json:"quantity,omitempty"`
		Timezone  *string               `json:"timezone,omitempty"`
		Type      *CalendarIntervalType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"alignment", "quantity", "timezone", "type"})
	} else {
		return err
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

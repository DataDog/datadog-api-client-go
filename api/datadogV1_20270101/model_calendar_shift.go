// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CalendarShift Test-only generated schema for the dashboard versioning proof of concept.
type CalendarShift struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Amount float64 `json:"amount"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Unit CalendarShiftUnit `json:"unit"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCalendarShift instantiates a new CalendarShift object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCalendarShift(amount float64, unit CalendarShiftUnit) *CalendarShift {
	this := CalendarShift{}
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewCalendarShiftWithDefaults instantiates a new CalendarShift object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCalendarShiftWithDefaults() *CalendarShift {
	this := CalendarShift{}
	return &this
}

// GetAmount returns the Amount field value.
func (o *CalendarShift) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CalendarShift) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value.
func (o *CalendarShift) SetAmount(v float64) {
	o.Amount = v
}

// GetUnit returns the Unit field value.
func (o *CalendarShift) GetUnit() CalendarShiftUnit {
	if o == nil {
		var ret CalendarShiftUnit
		return ret
	}
	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *CalendarShift) GetUnitOk() (*CalendarShiftUnit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value.
func (o *CalendarShift) SetUnit(v CalendarShiftUnit) {
	o.Unit = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CalendarShift) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["amount"] = o.Amount
	toSerialize["unit"] = o.Unit
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CalendarShift) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Amount *float64           `json:"amount"`
		Unit   *CalendarShiftUnit `json:"unit"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Amount == nil {
		return fmt.Errorf("required field amount missing")
	}
	if all.Unit == nil {
		return fmt.Errorf("required field unit missing")
	}

	hasInvalidField := false
	o.Amount = *all.Amount
	if !all.Unit.IsValid() {
		hasInvalidField = true
	} else {
		o.Unit = *all.Unit
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

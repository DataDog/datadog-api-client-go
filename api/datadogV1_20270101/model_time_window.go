// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeWindow Test-only generated schema for the dashboard versioning proof of concept.
type TimeWindow struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	From float64 `json:"from"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	To float64 `json:"to"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTimeWindow instantiates a new TimeWindow object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeWindow(from float64, to float64) *TimeWindow {
	this := TimeWindow{}
	this.From = from
	this.To = to
	return &this
}

// NewTimeWindowWithDefaults instantiates a new TimeWindow object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeWindowWithDefaults() *TimeWindow {
	this := TimeWindow{}
	return &this
}

// GetFrom returns the From field value.
func (o *TimeWindow) GetFrom() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *TimeWindow) GetFromOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *TimeWindow) SetFrom(v float64) {
	o.From = v
}

// GetTo returns the To field value.
func (o *TimeWindow) GetTo() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *TimeWindow) GetToOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *TimeWindow) SetTo(v float64) {
	o.To = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeWindow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["from"] = o.From
	toSerialize["to"] = o.To
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeWindow) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		From *float64 `json:"from"`
		To   *float64 `json:"to"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.From == nil {
		return fmt.Errorf("required field from missing")
	}
	if all.To == nil {
		return fmt.Errorf("required field to missing")
	}
	o.From = *all.From
	o.To = *all.To

	return nil
}

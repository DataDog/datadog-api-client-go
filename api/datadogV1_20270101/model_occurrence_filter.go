// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OccurrenceFilter Test-only generated schema for the dashboard versioning proof of concept.
type OccurrenceFilter struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Operator OccurrenceFilterOperator `json:"operator"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Value OccurrenceFilterValue `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewOccurrenceFilter instantiates a new OccurrenceFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOccurrenceFilter(operator OccurrenceFilterOperator, value OccurrenceFilterValue) *OccurrenceFilter {
	this := OccurrenceFilter{}
	this.Operator = operator
	this.Value = value
	return &this
}

// NewOccurrenceFilterWithDefaults instantiates a new OccurrenceFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOccurrenceFilterWithDefaults() *OccurrenceFilter {
	this := OccurrenceFilter{}
	return &this
}

// GetOperator returns the Operator field value.
func (o *OccurrenceFilter) GetOperator() OccurrenceFilterOperator {
	if o == nil {
		var ret OccurrenceFilterOperator
		return ret
	}
	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *OccurrenceFilter) GetOperatorOk() (*OccurrenceFilterOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value.
func (o *OccurrenceFilter) SetOperator(v OccurrenceFilterOperator) {
	o.Operator = v
}

// GetValue returns the Value field value.
func (o *OccurrenceFilter) GetValue() OccurrenceFilterValue {
	if o == nil {
		var ret OccurrenceFilterValue
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *OccurrenceFilter) GetValueOk() (*OccurrenceFilterValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *OccurrenceFilter) SetValue(v OccurrenceFilterValue) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OccurrenceFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["operator"] = o.Operator
	toSerialize["value"] = o.Value
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OccurrenceFilter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Operator *OccurrenceFilterOperator `json:"operator"`
		Value    *OccurrenceFilterValue    `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Operator == nil {
		return fmt.Errorf("required field operator missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}

	hasInvalidField := false
	if !all.Operator.IsValid() {
		hasInvalidField = true
	} else {
		o.Operator = *all.Operator
	}
	if !all.Value.IsValid() {
		hasInvalidField = true
	} else {
		o.Value = *all.Value
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

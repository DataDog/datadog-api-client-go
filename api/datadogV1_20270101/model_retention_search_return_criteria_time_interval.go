// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionSearchReturnCriteriaTimeInterval Test-only generated schema for the dashboard versioning proof of concept.
type RetentionSearchReturnCriteriaTimeInterval struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type RetentionSearchReturnCriteriaTimeIntervalType `json:"type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Unit RetentionSearchReturnCriteriaTimeIntervalUnit `json:"unit"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Value float64 `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRetentionSearchReturnCriteriaTimeInterval instantiates a new RetentionSearchReturnCriteriaTimeInterval object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRetentionSearchReturnCriteriaTimeInterval(typeVar RetentionSearchReturnCriteriaTimeIntervalType, unit RetentionSearchReturnCriteriaTimeIntervalUnit, value float64) *RetentionSearchReturnCriteriaTimeInterval {
	this := RetentionSearchReturnCriteriaTimeInterval{}
	this.Type = typeVar
	this.Unit = unit
	this.Value = value
	return &this
}

// NewRetentionSearchReturnCriteriaTimeIntervalWithDefaults instantiates a new RetentionSearchReturnCriteriaTimeInterval object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRetentionSearchReturnCriteriaTimeIntervalWithDefaults() *RetentionSearchReturnCriteriaTimeInterval {
	this := RetentionSearchReturnCriteriaTimeInterval{}
	return &this
}

// GetType returns the Type field value.
func (o *RetentionSearchReturnCriteriaTimeInterval) GetType() RetentionSearchReturnCriteriaTimeIntervalType {
	if o == nil {
		var ret RetentionSearchReturnCriteriaTimeIntervalType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RetentionSearchReturnCriteriaTimeInterval) GetTypeOk() (*RetentionSearchReturnCriteriaTimeIntervalType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RetentionSearchReturnCriteriaTimeInterval) SetType(v RetentionSearchReturnCriteriaTimeIntervalType) {
	o.Type = v
}

// GetUnit returns the Unit field value.
func (o *RetentionSearchReturnCriteriaTimeInterval) GetUnit() RetentionSearchReturnCriteriaTimeIntervalUnit {
	if o == nil {
		var ret RetentionSearchReturnCriteriaTimeIntervalUnit
		return ret
	}
	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *RetentionSearchReturnCriteriaTimeInterval) GetUnitOk() (*RetentionSearchReturnCriteriaTimeIntervalUnit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value.
func (o *RetentionSearchReturnCriteriaTimeInterval) SetUnit(v RetentionSearchReturnCriteriaTimeIntervalUnit) {
	o.Unit = v
}

// GetValue returns the Value field value.
func (o *RetentionSearchReturnCriteriaTimeInterval) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RetentionSearchReturnCriteriaTimeInterval) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *RetentionSearchReturnCriteriaTimeInterval) SetValue(v float64) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RetentionSearchReturnCriteriaTimeInterval) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	toSerialize["unit"] = o.Unit
	toSerialize["value"] = o.Value
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RetentionSearchReturnCriteriaTimeInterval) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type  *RetentionSearchReturnCriteriaTimeIntervalType `json:"type"`
		Unit  *RetentionSearchReturnCriteriaTimeIntervalUnit `json:"unit"`
		Value *float64                                       `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Unit == nil {
		return fmt.Errorf("required field unit missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}

	hasInvalidField := false
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	if !all.Unit.IsValid() {
		hasInvalidField = true
	} else {
		o.Unit = *all.Unit
	}
	o.Value = *all.Value

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionSearchCohortCriteriaTimeInterval Test-only generated schema for the dashboard versioning proof of concept.
type RetentionSearchCohortCriteriaTimeInterval struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type RetentionSearchCohortCriteriaTimeIntervalType `json:"type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Value RetentionSearchCohortCriteriaTimeInterval20270101Value `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRetentionSearchCohortCriteriaTimeInterval instantiates a new RetentionSearchCohortCriteriaTimeInterval object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRetentionSearchCohortCriteriaTimeInterval(typeVar RetentionSearchCohortCriteriaTimeIntervalType, value RetentionSearchCohortCriteriaTimeInterval20270101Value) *RetentionSearchCohortCriteriaTimeInterval {
	this := RetentionSearchCohortCriteriaTimeInterval{}
	this.Type = typeVar
	this.Value = value
	return &this
}

// NewRetentionSearchCohortCriteriaTimeIntervalWithDefaults instantiates a new RetentionSearchCohortCriteriaTimeInterval object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRetentionSearchCohortCriteriaTimeIntervalWithDefaults() *RetentionSearchCohortCriteriaTimeInterval {
	this := RetentionSearchCohortCriteriaTimeInterval{}
	return &this
}

// GetType returns the Type field value.
func (o *RetentionSearchCohortCriteriaTimeInterval) GetType() RetentionSearchCohortCriteriaTimeIntervalType {
	if o == nil {
		var ret RetentionSearchCohortCriteriaTimeIntervalType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RetentionSearchCohortCriteriaTimeInterval) GetTypeOk() (*RetentionSearchCohortCriteriaTimeIntervalType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RetentionSearchCohortCriteriaTimeInterval) SetType(v RetentionSearchCohortCriteriaTimeIntervalType) {
	o.Type = v
}

// GetValue returns the Value field value.
func (o *RetentionSearchCohortCriteriaTimeInterval) GetValue() RetentionSearchCohortCriteriaTimeInterval20270101Value {
	if o == nil {
		var ret RetentionSearchCohortCriteriaTimeInterval20270101Value
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RetentionSearchCohortCriteriaTimeInterval) GetValueOk() (*RetentionSearchCohortCriteriaTimeInterval20270101Value, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *RetentionSearchCohortCriteriaTimeInterval) SetValue(v RetentionSearchCohortCriteriaTimeInterval20270101Value) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RetentionSearchCohortCriteriaTimeInterval) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RetentionSearchCohortCriteriaTimeInterval) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type  *RetentionSearchCohortCriteriaTimeIntervalType          `json:"type"`
		Value *RetentionSearchCohortCriteriaTimeInterval20270101Value `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
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
	o.Value = *all.Value

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

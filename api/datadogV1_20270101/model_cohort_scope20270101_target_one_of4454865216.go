// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CohortScope20270101TargetOneOf4454865216 Test-only generated schema for the dashboard versioning proof of concept.
type CohortScope20270101TargetOneOf4454865216 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type CohortScopeTargetAggregationType `json:"type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Value CohortScopeTargetAnyOfAnyOfItemValue `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCohortScope20270101TargetOneOf4454865216 instantiates a new CohortScope20270101TargetOneOf4454865216 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCohortScope20270101TargetOneOf4454865216(typeVar CohortScopeTargetAggregationType, value CohortScopeTargetAnyOfAnyOfItemValue) *CohortScope20270101TargetOneOf4454865216 {
	this := CohortScope20270101TargetOneOf4454865216{}
	this.Type = typeVar
	this.Value = value
	return &this
}

// NewCohortScope20270101TargetOneOf4454865216WithDefaults instantiates a new CohortScope20270101TargetOneOf4454865216 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCohortScope20270101TargetOneOf4454865216WithDefaults() *CohortScope20270101TargetOneOf4454865216 {
	this := CohortScope20270101TargetOneOf4454865216{}
	return &this
}

// GetType returns the Type field value.
func (o *CohortScope20270101TargetOneOf4454865216) GetType() CohortScopeTargetAggregationType {
	if o == nil {
		var ret CohortScopeTargetAggregationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CohortScope20270101TargetOneOf4454865216) GetTypeOk() (*CohortScopeTargetAggregationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CohortScope20270101TargetOneOf4454865216) SetType(v CohortScopeTargetAggregationType) {
	o.Type = v
}

// GetValue returns the Value field value.
func (o *CohortScope20270101TargetOneOf4454865216) GetValue() CohortScopeTargetAnyOfAnyOfItemValue {
	if o == nil {
		var ret CohortScopeTargetAnyOfAnyOfItemValue
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CohortScope20270101TargetOneOf4454865216) GetValueOk() (*CohortScopeTargetAnyOfAnyOfItemValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *CohortScope20270101TargetOneOf4454865216) SetValue(v CohortScopeTargetAnyOfAnyOfItemValue) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CohortScope20270101TargetOneOf4454865216) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CohortScope20270101TargetOneOf4454865216) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type  *CohortScopeTargetAggregationType     `json:"type"`
		Value *CohortScopeTargetAnyOfAnyOfItemValue `json:"value"`
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

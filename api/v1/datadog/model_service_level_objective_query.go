/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// ServiceLevelObjectiveQuery A metric SLI query. Required if type is \"metric\".
type ServiceLevelObjectiveQuery struct {
	// A Datadog metric query for total (valid) events.
	Denominator string `json:"denominator"`
	// A Datadog metric query for good events.
	Numerator string `json:"numerator"`
}

// NewServiceLevelObjectiveQuery instantiates a new ServiceLevelObjectiveQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceLevelObjectiveQuery(denominator string, numerator string) *ServiceLevelObjectiveQuery {
	this := ServiceLevelObjectiveQuery{}
	this.Denominator = denominator
	this.Numerator = numerator
	return &this
}

// NewServiceLevelObjectiveQueryWithDefaults instantiates a new ServiceLevelObjectiveQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceLevelObjectiveQueryWithDefaults() *ServiceLevelObjectiveQuery {
	this := ServiceLevelObjectiveQuery{}
	return &this
}

// GetDenominator returns the Denominator field value
func (o *ServiceLevelObjectiveQuery) GetDenominator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Denominator
}

// SetDenominator sets field value
func (o *ServiceLevelObjectiveQuery) SetDenominator(v string) {
	o.Denominator = v
}

// GetNumerator returns the Numerator field value
func (o *ServiceLevelObjectiveQuery) GetNumerator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Numerator
}

// SetNumerator sets field value
func (o *ServiceLevelObjectiveQuery) SetNumerator(v string) {
	o.Numerator = v
}

type NullableServiceLevelObjectiveQuery struct {
	Value        ServiceLevelObjectiveQuery
	ExplicitNull bool
}

func (v NullableServiceLevelObjectiveQuery) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableServiceLevelObjectiveQuery) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

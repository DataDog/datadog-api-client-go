// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CellScope20270101CohortTargetOneOf Test-only generated schema for the dashboard versioning proof of concept.
type CellScope20270101CohortTargetOneOf struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type CellScopeCohortTargetIndexType `json:"type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Value float64 `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCellScope20270101CohortTargetOneOf instantiates a new CellScope20270101CohortTargetOneOf object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCellScope20270101CohortTargetOneOf(typeVar CellScopeCohortTargetIndexType, value float64) *CellScope20270101CohortTargetOneOf {
	this := CellScope20270101CohortTargetOneOf{}
	this.Type = typeVar
	this.Value = value
	return &this
}

// NewCellScope20270101CohortTargetOneOfWithDefaults instantiates a new CellScope20270101CohortTargetOneOf object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCellScope20270101CohortTargetOneOfWithDefaults() *CellScope20270101CohortTargetOneOf {
	this := CellScope20270101CohortTargetOneOf{}
	return &this
}

// GetType returns the Type field value.
func (o *CellScope20270101CohortTargetOneOf) GetType() CellScopeCohortTargetIndexType {
	if o == nil {
		var ret CellScopeCohortTargetIndexType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CellScope20270101CohortTargetOneOf) GetTypeOk() (*CellScopeCohortTargetIndexType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CellScope20270101CohortTargetOneOf) SetType(v CellScopeCohortTargetIndexType) {
	o.Type = v
}

// GetValue returns the Value field value.
func (o *CellScope20270101CohortTargetOneOf) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CellScope20270101CohortTargetOneOf) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *CellScope20270101CohortTargetOneOf) SetValue(v float64) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CellScope20270101CohortTargetOneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CellScope20270101CohortTargetOneOf) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type  *CellScopeCohortTargetIndexType `json:"type"`
		Value *float64                        `json:"value"`
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

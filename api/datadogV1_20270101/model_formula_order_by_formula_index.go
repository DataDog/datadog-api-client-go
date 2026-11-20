// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaOrderByFormulaIndex Test-only generated schema for the dashboard versioning proof of concept.
type FormulaOrderByFormulaIndex struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Index float64 `json:"index"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Order QuerySortOrder `json:"order"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type FormulaOrderByFormulaIndexType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaOrderByFormulaIndex instantiates a new FormulaOrderByFormulaIndex object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaOrderByFormulaIndex(index float64, order QuerySortOrder, typeVar FormulaOrderByFormulaIndexType) *FormulaOrderByFormulaIndex {
	this := FormulaOrderByFormulaIndex{}
	this.Index = index
	this.Order = order
	this.Type = typeVar
	return &this
}

// NewFormulaOrderByFormulaIndexWithDefaults instantiates a new FormulaOrderByFormulaIndex object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaOrderByFormulaIndexWithDefaults() *FormulaOrderByFormulaIndex {
	this := FormulaOrderByFormulaIndex{}
	return &this
}

// GetIndex returns the Index field value.
func (o *FormulaOrderByFormulaIndex) GetIndex() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *FormulaOrderByFormulaIndex) GetIndexOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value.
func (o *FormulaOrderByFormulaIndex) SetIndex(v float64) {
	o.Index = v
}

// GetOrder returns the Order field value.
func (o *FormulaOrderByFormulaIndex) GetOrder() QuerySortOrder {
	if o == nil {
		var ret QuerySortOrder
		return ret
	}
	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *FormulaOrderByFormulaIndex) GetOrderOk() (*QuerySortOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value.
func (o *FormulaOrderByFormulaIndex) SetOrder(v QuerySortOrder) {
	o.Order = v
}

// GetType returns the Type field value.
func (o *FormulaOrderByFormulaIndex) GetType() FormulaOrderByFormulaIndexType {
	if o == nil {
		var ret FormulaOrderByFormulaIndexType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FormulaOrderByFormulaIndex) GetTypeOk() (*FormulaOrderByFormulaIndexType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *FormulaOrderByFormulaIndex) SetType(v FormulaOrderByFormulaIndexType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaOrderByFormulaIndex) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["index"] = o.Index
	toSerialize["order"] = o.Order
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaOrderByFormulaIndex) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Index *float64                        `json:"index"`
		Order *QuerySortOrder                 `json:"order"`
		Type  *FormulaOrderByFormulaIndexType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Index == nil {
		return fmt.Errorf("required field index missing")
	}
	if all.Order == nil {
		return fmt.Errorf("required field order missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Index = *all.Index
	if !all.Order.IsValid() {
		hasInvalidField = true
	} else {
		o.Order = *all.Order
	}
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

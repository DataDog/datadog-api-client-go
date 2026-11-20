// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaOrderByGroup Test-only generated schema for the dashboard versioning proof of concept.
type FormulaOrderByGroup struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Order QuerySortOrder `json:"order"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type FormulaOrderByGroupType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaOrderByGroup instantiates a new FormulaOrderByGroup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaOrderByGroup(name string, order QuerySortOrder, typeVar FormulaOrderByGroupType) *FormulaOrderByGroup {
	this := FormulaOrderByGroup{}
	this.Name = name
	this.Order = order
	this.Type = typeVar
	return &this
}

// NewFormulaOrderByGroupWithDefaults instantiates a new FormulaOrderByGroup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaOrderByGroupWithDefaults() *FormulaOrderByGroup {
	this := FormulaOrderByGroup{}
	return &this
}

// GetName returns the Name field value.
func (o *FormulaOrderByGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaOrderByGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaOrderByGroup) SetName(v string) {
	o.Name = v
}

// GetOrder returns the Order field value.
func (o *FormulaOrderByGroup) GetOrder() QuerySortOrder {
	if o == nil {
		var ret QuerySortOrder
		return ret
	}
	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *FormulaOrderByGroup) GetOrderOk() (*QuerySortOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value.
func (o *FormulaOrderByGroup) SetOrder(v QuerySortOrder) {
	o.Order = v
}

// GetType returns the Type field value.
func (o *FormulaOrderByGroup) GetType() FormulaOrderByGroupType {
	if o == nil {
		var ret FormulaOrderByGroupType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FormulaOrderByGroup) GetTypeOk() (*FormulaOrderByGroupType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *FormulaOrderByGroup) SetType(v FormulaOrderByGroupType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaOrderByGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["order"] = o.Order
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaOrderByGroup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name  *string                  `json:"name"`
		Order *QuerySortOrder          `json:"order"`
		Type  *FormulaOrderByGroupType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Order == nil {
		return fmt.Errorf("required field order missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Name = *all.Name
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

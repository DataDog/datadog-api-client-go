// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlatSortBy20270101OneOf4456386848 Test-only generated schema for the dashboard versioning proof of concept.
type FlatSortBy20270101OneOf4456386848 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Order QuerySortOrder `json:"order"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Property FlatSortByAnyOfAnyOfItemProperty4446002432 `json:"property"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFlatSortBy20270101OneOf4456386848 instantiates a new FlatSortBy20270101OneOf4456386848 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFlatSortBy20270101OneOf4456386848(order QuerySortOrder, property FlatSortByAnyOfAnyOfItemProperty4446002432) *FlatSortBy20270101OneOf4456386848 {
	this := FlatSortBy20270101OneOf4456386848{}
	this.Order = order
	this.Property = property
	return &this
}

// NewFlatSortBy20270101OneOf4456386848WithDefaults instantiates a new FlatSortBy20270101OneOf4456386848 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFlatSortBy20270101OneOf4456386848WithDefaults() *FlatSortBy20270101OneOf4456386848 {
	this := FlatSortBy20270101OneOf4456386848{}
	return &this
}

// GetOrder returns the Order field value.
func (o *FlatSortBy20270101OneOf4456386848) GetOrder() QuerySortOrder {
	if o == nil {
		var ret QuerySortOrder
		return ret
	}
	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *FlatSortBy20270101OneOf4456386848) GetOrderOk() (*QuerySortOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value.
func (o *FlatSortBy20270101OneOf4456386848) SetOrder(v QuerySortOrder) {
	o.Order = v
}

// GetProperty returns the Property field value.
func (o *FlatSortBy20270101OneOf4456386848) GetProperty() FlatSortByAnyOfAnyOfItemProperty4446002432 {
	if o == nil {
		var ret FlatSortByAnyOfAnyOfItemProperty4446002432
		return ret
	}
	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value
// and a boolean to check if the value has been set.
func (o *FlatSortBy20270101OneOf4456386848) GetPropertyOk() (*FlatSortByAnyOfAnyOfItemProperty4446002432, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Property, true
}

// SetProperty sets field value.
func (o *FlatSortBy20270101OneOf4456386848) SetProperty(v FlatSortByAnyOfAnyOfItemProperty4446002432) {
	o.Property = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FlatSortBy20270101OneOf4456386848) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["order"] = o.Order
	toSerialize["property"] = o.Property
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FlatSortBy20270101OneOf4456386848) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Order    *QuerySortOrder                             `json:"order"`
		Property *FlatSortByAnyOfAnyOfItemProperty4446002432 `json:"property"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Order == nil {
		return fmt.Errorf("required field order missing")
	}
	if all.Property == nil {
		return fmt.Errorf("required field property missing")
	}

	hasInvalidField := false
	if !all.Order.IsValid() {
		hasInvalidField = true
	} else {
		o.Order = *all.Order
	}
	if !all.Property.IsValid() {
		hasInvalidField = true
	} else {
		o.Property = *all.Property
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

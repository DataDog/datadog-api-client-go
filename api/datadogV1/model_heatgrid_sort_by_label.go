// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridSortByLabel Sort rows by their group labels.
type HeatgridSortByLabel struct {
	// Sort direction.
	Order HeatgridSortOrder `json:"order"`
	// Sort by label.
	Property HeatgridSortByLabelProperty `json:"property"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridSortByLabel instantiates a new HeatgridSortByLabel object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridSortByLabel(order HeatgridSortOrder, property HeatgridSortByLabelProperty) *HeatgridSortByLabel {
	this := HeatgridSortByLabel{}
	this.Order = order
	this.Property = property
	return &this
}

// NewHeatgridSortByLabelWithDefaults instantiates a new HeatgridSortByLabel object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridSortByLabelWithDefaults() *HeatgridSortByLabel {
	this := HeatgridSortByLabel{}
	return &this
}

// GetOrder returns the Order field value.
func (o *HeatgridSortByLabel) GetOrder() HeatgridSortOrder {
	if o == nil {
		var ret HeatgridSortOrder
		return ret
	}
	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *HeatgridSortByLabel) GetOrderOk() (*HeatgridSortOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value.
func (o *HeatgridSortByLabel) SetOrder(v HeatgridSortOrder) {
	o.Order = v
}

// GetProperty returns the Property field value.
func (o *HeatgridSortByLabel) GetProperty() HeatgridSortByLabelProperty {
	if o == nil {
		var ret HeatgridSortByLabelProperty
		return ret
	}
	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value
// and a boolean to check if the value has been set.
func (o *HeatgridSortByLabel) GetPropertyOk() (*HeatgridSortByLabelProperty, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Property, true
}

// SetProperty sets field value.
func (o *HeatgridSortByLabel) SetProperty(v HeatgridSortByLabelProperty) {
	o.Property = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridSortByLabel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["order"] = o.Order
	toSerialize["property"] = o.Property
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HeatgridSortByLabel) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Order    *HeatgridSortOrder           `json:"order"`
		Property *HeatgridSortByLabelProperty `json:"property"`
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

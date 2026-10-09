// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridSortByValue Sort rows by their aggregated values.
type HeatgridSortByValue struct {
	// Aggregation used to order rows over the displayed time range.
	Aggregation HeatgridSortAggregation `json:"aggregation"`
	// Sort direction.
	Order HeatgridSortOrder `json:"order"`
	// Sort by value.
	Property HeatgridSortByValueProperty `json:"property"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridSortByValue instantiates a new HeatgridSortByValue object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridSortByValue(aggregation HeatgridSortAggregation, order HeatgridSortOrder, property HeatgridSortByValueProperty) *HeatgridSortByValue {
	this := HeatgridSortByValue{}
	this.Aggregation = aggregation
	this.Order = order
	this.Property = property
	return &this
}

// NewHeatgridSortByValueWithDefaults instantiates a new HeatgridSortByValue object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridSortByValueWithDefaults() *HeatgridSortByValue {
	this := HeatgridSortByValue{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *HeatgridSortByValue) GetAggregation() HeatgridSortAggregation {
	if o == nil {
		var ret HeatgridSortAggregation
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *HeatgridSortByValue) GetAggregationOk() (*HeatgridSortAggregation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *HeatgridSortByValue) SetAggregation(v HeatgridSortAggregation) {
	o.Aggregation = v
}

// GetOrder returns the Order field value.
func (o *HeatgridSortByValue) GetOrder() HeatgridSortOrder {
	if o == nil {
		var ret HeatgridSortOrder
		return ret
	}
	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *HeatgridSortByValue) GetOrderOk() (*HeatgridSortOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value.
func (o *HeatgridSortByValue) SetOrder(v HeatgridSortOrder) {
	o.Order = v
}

// GetProperty returns the Property field value.
func (o *HeatgridSortByValue) GetProperty() HeatgridSortByValueProperty {
	if o == nil {
		var ret HeatgridSortByValueProperty
		return ret
	}
	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value
// and a boolean to check if the value has been set.
func (o *HeatgridSortByValue) GetPropertyOk() (*HeatgridSortByValueProperty, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Property, true
}

// SetProperty sets field value.
func (o *HeatgridSortByValue) SetProperty(v HeatgridSortByValueProperty) {
	o.Property = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridSortByValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation"] = o.Aggregation
	toSerialize["order"] = o.Order
	toSerialize["property"] = o.Property
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HeatgridSortByValue) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *HeatgridSortAggregation     `json:"aggregation"`
		Order       *HeatgridSortOrder           `json:"order"`
		Property    *HeatgridSortByValueProperty `json:"property"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Aggregation == nil {
		return fmt.Errorf("required field aggregation missing")
	}
	if all.Order == nil {
		return fmt.Errorf("required field order missing")
	}
	if all.Property == nil {
		return fmt.Errorf("required field property missing")
	}

	hasInvalidField := false
	if !all.Aggregation.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregation = *all.Aggregation
	}
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

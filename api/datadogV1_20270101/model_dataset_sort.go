// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetSort Test-only generated schema for the dashboard versioning proof of concept.
type DatasetSort struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregation DatasetSort20270101Aggregation `json:"aggregation"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Column *string `json:"column,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Order QuerySortOrder `json:"order"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDatasetSort instantiates a new DatasetSort object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatasetSort(aggregation DatasetSort20270101Aggregation, order QuerySortOrder) *DatasetSort {
	this := DatasetSort{}
	this.Aggregation = aggregation
	this.Order = order
	return &this
}

// NewDatasetSortWithDefaults instantiates a new DatasetSort object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatasetSortWithDefaults() *DatasetSort {
	this := DatasetSort{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *DatasetSort) GetAggregation() DatasetSort20270101Aggregation {
	if o == nil {
		var ret DatasetSort20270101Aggregation
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *DatasetSort) GetAggregationOk() (*DatasetSort20270101Aggregation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *DatasetSort) SetAggregation(v DatasetSort20270101Aggregation) {
	o.Aggregation = v
}

// GetColumn returns the Column field value if set, zero value otherwise.
func (o *DatasetSort) GetColumn() string {
	if o == nil || o.Column == nil {
		var ret string
		return ret
	}
	return *o.Column
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetSort) GetColumnOk() (*string, bool) {
	if o == nil || o.Column == nil {
		return nil, false
	}
	return o.Column, true
}

// HasColumn returns a boolean if a field has been set.
func (o *DatasetSort) HasColumn() bool {
	return o != nil && o.Column != nil
}

// SetColumn gets a reference to the given string and assigns it to the Column field.
func (o *DatasetSort) SetColumn(v string) {
	o.Column = &v
}

// GetOrder returns the Order field value.
func (o *DatasetSort) GetOrder() QuerySortOrder {
	if o == nil {
		var ret QuerySortOrder
		return ret
	}
	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *DatasetSort) GetOrderOk() (*QuerySortOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value.
func (o *DatasetSort) SetOrder(v QuerySortOrder) {
	o.Order = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatasetSort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation"] = o.Aggregation
	if o.Column != nil {
		toSerialize["column"] = o.Column
	}
	toSerialize["order"] = o.Order
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatasetSort) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *DatasetSort20270101Aggregation `json:"aggregation"`
		Column      *string                         `json:"column,omitempty"`
		Order       *QuerySortOrder                 `json:"order"`
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

	hasInvalidField := false
	o.Aggregation = *all.Aggregation
	o.Column = all.Column
	if !all.Order.IsValid() {
		hasInvalidField = true
	} else {
		o.Order = *all.Order
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

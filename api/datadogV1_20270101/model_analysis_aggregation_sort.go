// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnalysisAggregationSort Test-only generated schema for the dashboard versioning proof of concept.
type AnalysisAggregationSort struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Column string `json:"column"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Order SortOrder2 `json:"order"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAnalysisAggregationSort instantiates a new AnalysisAggregationSort object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnalysisAggregationSort(column string, order SortOrder2) *AnalysisAggregationSort {
	this := AnalysisAggregationSort{}
	this.Column = column
	this.Order = order
	return &this
}

// NewAnalysisAggregationSortWithDefaults instantiates a new AnalysisAggregationSort object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnalysisAggregationSortWithDefaults() *AnalysisAggregationSort {
	this := AnalysisAggregationSort{}
	return &this
}

// GetColumn returns the Column field value.
func (o *AnalysisAggregationSort) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *AnalysisAggregationSort) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *AnalysisAggregationSort) SetColumn(v string) {
	o.Column = v
}

// GetOrder returns the Order field value.
func (o *AnalysisAggregationSort) GetOrder() SortOrder2 {
	if o == nil {
		var ret SortOrder2
		return ret
	}
	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *AnalysisAggregationSort) GetOrderOk() (*SortOrder2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value.
func (o *AnalysisAggregationSort) SetOrder(v SortOrder2) {
	o.Order = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnalysisAggregationSort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["column"] = o.Column
	toSerialize["order"] = o.Order
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AnalysisAggregationSort) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Column *string     `json:"column"`
		Order  *SortOrder2 `json:"order"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Column == nil {
		return fmt.Errorf("required field column missing")
	}
	if all.Order == nil {
		return fmt.Errorf("required field order missing")
	}

	hasInvalidField := false
	o.Column = *all.Column
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

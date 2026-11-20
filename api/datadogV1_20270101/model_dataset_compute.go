// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetCompute Test-only generated schema for the dashboard versioning proof of concept.
type DatasetCompute struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregation DatasetCompute20270101Aggregation `json:"aggregation"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Column *string `json:"column,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Interval *float64 `json:"interval,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDatasetCompute instantiates a new DatasetCompute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatasetCompute(aggregation DatasetCompute20270101Aggregation, name string) *DatasetCompute {
	this := DatasetCompute{}
	this.Aggregation = aggregation
	this.Name = name
	return &this
}

// NewDatasetComputeWithDefaults instantiates a new DatasetCompute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatasetComputeWithDefaults() *DatasetCompute {
	this := DatasetCompute{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *DatasetCompute) GetAggregation() DatasetCompute20270101Aggregation {
	if o == nil {
		var ret DatasetCompute20270101Aggregation
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *DatasetCompute) GetAggregationOk() (*DatasetCompute20270101Aggregation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *DatasetCompute) SetAggregation(v DatasetCompute20270101Aggregation) {
	o.Aggregation = v
}

// GetColumn returns the Column field value if set, zero value otherwise.
func (o *DatasetCompute) GetColumn() string {
	if o == nil || o.Column == nil {
		var ret string
		return ret
	}
	return *o.Column
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetCompute) GetColumnOk() (*string, bool) {
	if o == nil || o.Column == nil {
		return nil, false
	}
	return o.Column, true
}

// HasColumn returns a boolean if a field has been set.
func (o *DatasetCompute) HasColumn() bool {
	return o != nil && o.Column != nil
}

// SetColumn gets a reference to the given string and assigns it to the Column field.
func (o *DatasetCompute) SetColumn(v string) {
	o.Column = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *DatasetCompute) GetInterval() float64 {
	if o == nil || o.Interval == nil {
		var ret float64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetCompute) GetIntervalOk() (*float64, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *DatasetCompute) HasInterval() bool {
	return o != nil && o.Interval != nil
}

// SetInterval gets a reference to the given float64 and assigns it to the Interval field.
func (o *DatasetCompute) SetInterval(v float64) {
	o.Interval = &v
}

// GetName returns the Name field value.
func (o *DatasetCompute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatasetCompute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DatasetCompute) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatasetCompute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation"] = o.Aggregation
	if o.Column != nil {
		toSerialize["column"] = o.Column
	}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	toSerialize["name"] = o.Name
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatasetCompute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *DatasetCompute20270101Aggregation `json:"aggregation"`
		Column      *string                            `json:"column,omitempty"`
		Interval    *float64                           `json:"interval,omitempty"`
		Name        *string                            `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Aggregation == nil {
		return fmt.Errorf("required field aggregation missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	o.Aggregation = *all.Aggregation
	o.Column = all.Column
	o.Interval = all.Interval
	o.Name = *all.Name

	return nil
}

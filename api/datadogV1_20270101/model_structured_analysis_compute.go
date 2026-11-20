// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisCompute Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisCompute struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregation string `json:"aggregation"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Column string `json:"column"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewStructuredAnalysisCompute instantiates a new StructuredAnalysisCompute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStructuredAnalysisCompute(aggregation string, column string) *StructuredAnalysisCompute {
	this := StructuredAnalysisCompute{}
	this.Aggregation = aggregation
	this.Column = column
	return &this
}

// NewStructuredAnalysisComputeWithDefaults instantiates a new StructuredAnalysisCompute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStructuredAnalysisComputeWithDefaults() *StructuredAnalysisCompute {
	this := StructuredAnalysisCompute{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *StructuredAnalysisCompute) GetAggregation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisCompute) GetAggregationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *StructuredAnalysisCompute) SetAggregation(v string) {
	o.Aggregation = v
}

// GetColumn returns the Column field value.
func (o *StructuredAnalysisCompute) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisCompute) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *StructuredAnalysisCompute) SetColumn(v string) {
	o.Column = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StructuredAnalysisCompute) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisCompute) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StructuredAnalysisCompute) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StructuredAnalysisCompute) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o StructuredAnalysisCompute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation"] = o.Aggregation
	toSerialize["column"] = o.Column
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StructuredAnalysisCompute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *string `json:"aggregation"`
		Column      *string `json:"column"`
		Name        *string `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Aggregation == nil {
		return fmt.Errorf("required field aggregation missing")
	}
	if all.Column == nil {
		return fmt.Errorf("required field column missing")
	}
	o.Aggregation = *all.Aggregation
	o.Column = *all.Column
	o.Name = all.Name

	return nil
}

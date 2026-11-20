// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisTimeAggregation Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisTimeAggregation struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Interval *float64 `json:"interval,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TimeColumn string `json:"time_column"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewStructuredAnalysisTimeAggregation instantiates a new StructuredAnalysisTimeAggregation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStructuredAnalysisTimeAggregation(timeColumn string) *StructuredAnalysisTimeAggregation {
	this := StructuredAnalysisTimeAggregation{}
	this.TimeColumn = timeColumn
	return &this
}

// NewStructuredAnalysisTimeAggregationWithDefaults instantiates a new StructuredAnalysisTimeAggregation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStructuredAnalysisTimeAggregationWithDefaults() *StructuredAnalysisTimeAggregation {
	this := StructuredAnalysisTimeAggregation{}
	return &this
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *StructuredAnalysisTimeAggregation) GetInterval() float64 {
	if o == nil || o.Interval == nil {
		var ret float64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisTimeAggregation) GetIntervalOk() (*float64, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *StructuredAnalysisTimeAggregation) HasInterval() bool {
	return o != nil && o.Interval != nil
}

// SetInterval gets a reference to the given float64 and assigns it to the Interval field.
func (o *StructuredAnalysisTimeAggregation) SetInterval(v float64) {
	o.Interval = &v
}

// GetTimeColumn returns the TimeColumn field value.
func (o *StructuredAnalysisTimeAggregation) GetTimeColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TimeColumn
}

// GetTimeColumnOk returns a tuple with the TimeColumn field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisTimeAggregation) GetTimeColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeColumn, true
}

// SetTimeColumn sets field value.
func (o *StructuredAnalysisTimeAggregation) SetTimeColumn(v string) {
	o.TimeColumn = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StructuredAnalysisTimeAggregation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	toSerialize["time_column"] = o.TimeColumn
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StructuredAnalysisTimeAggregation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Interval   *float64 `json:"interval,omitempty"`
		TimeColumn *string  `json:"time_column"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TimeColumn == nil {
		return fmt.Errorf("required field time_column missing")
	}
	o.Interval = all.Interval
	o.TimeColumn = *all.TimeColumn

	return nil
}

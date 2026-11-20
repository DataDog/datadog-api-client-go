// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SplitSortCompute Test-only generated schema for the dashboard versioning proof of concept.
type SplitSortCompute struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregation string `json:"aggregation"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Metric string `json:"metric"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSplitSortCompute instantiates a new SplitSortCompute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSplitSortCompute(aggregation string, metric string) *SplitSortCompute {
	this := SplitSortCompute{}
	this.Aggregation = aggregation
	this.Metric = metric
	return &this
}

// NewSplitSortComputeWithDefaults instantiates a new SplitSortCompute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSplitSortComputeWithDefaults() *SplitSortCompute {
	this := SplitSortCompute{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *SplitSortCompute) GetAggregation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *SplitSortCompute) GetAggregationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *SplitSortCompute) SetAggregation(v string) {
	o.Aggregation = v
}

// GetMetric returns the Metric field value.
func (o *SplitSortCompute) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *SplitSortCompute) GetMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value.
func (o *SplitSortCompute) SetMetric(v string) {
	o.Metric = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SplitSortCompute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation"] = o.Aggregation
	toSerialize["metric"] = o.Metric
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SplitSortCompute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *string `json:"aggregation"`
		Metric      *string `json:"metric"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Aggregation == nil {
		return fmt.Errorf("required field aggregation missing")
	}
	if all.Metric == nil {
		return fmt.Errorf("required field metric missing")
	}
	o.Aggregation = *all.Aggregation
	o.Metric = *all.Metric

	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregateCompute Test-only generated schema for the dashboard versioning proof of concept.
type AggregateCompute struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregation EventsAggregator `json:"aggregation"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Interval *float64 `json:"interval,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Metric *string `json:"metric,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Rollup *CalendarInterval `json:"rollup,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAggregateCompute instantiates a new AggregateCompute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAggregateCompute(aggregation EventsAggregator, name string) *AggregateCompute {
	this := AggregateCompute{}
	this.Aggregation = aggregation
	this.Name = name
	return &this
}

// NewAggregateComputeWithDefaults instantiates a new AggregateCompute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAggregateComputeWithDefaults() *AggregateCompute {
	this := AggregateCompute{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *AggregateCompute) GetAggregation() EventsAggregator {
	if o == nil {
		var ret EventsAggregator
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *AggregateCompute) GetAggregationOk() (*EventsAggregator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *AggregateCompute) SetAggregation(v EventsAggregator) {
	o.Aggregation = v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *AggregateCompute) GetInterval() float64 {
	if o == nil || o.Interval == nil {
		var ret float64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateCompute) GetIntervalOk() (*float64, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *AggregateCompute) HasInterval() bool {
	return o != nil && o.Interval != nil
}

// SetInterval gets a reference to the given float64 and assigns it to the Interval field.
func (o *AggregateCompute) SetInterval(v float64) {
	o.Interval = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *AggregateCompute) GetMetric() string {
	if o == nil || o.Metric == nil {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateCompute) GetMetricOk() (*string, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *AggregateCompute) HasMetric() bool {
	return o != nil && o.Metric != nil
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *AggregateCompute) SetMetric(v string) {
	o.Metric = &v
}

// GetName returns the Name field value.
func (o *AggregateCompute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AggregateCompute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AggregateCompute) SetName(v string) {
	o.Name = v
}

// GetRollup returns the Rollup field value if set, zero value otherwise.
func (o *AggregateCompute) GetRollup() CalendarInterval {
	if o == nil || o.Rollup == nil {
		var ret CalendarInterval
		return ret
	}
	return *o.Rollup
}

// GetRollupOk returns a tuple with the Rollup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateCompute) GetRollupOk() (*CalendarInterval, bool) {
	if o == nil || o.Rollup == nil {
		return nil, false
	}
	return o.Rollup, true
}

// HasRollup returns a boolean if a field has been set.
func (o *AggregateCompute) HasRollup() bool {
	return o != nil && o.Rollup != nil
}

// SetRollup gets a reference to the given CalendarInterval and assigns it to the Rollup field.
func (o *AggregateCompute) SetRollup(v CalendarInterval) {
	o.Rollup = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AggregateCompute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation"] = o.Aggregation
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
	}
	toSerialize["name"] = o.Name
	if o.Rollup != nil {
		toSerialize["rollup"] = o.Rollup
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AggregateCompute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *EventsAggregator `json:"aggregation"`
		Interval    *float64          `json:"interval,omitempty"`
		Metric      *string           `json:"metric,omitempty"`
		Name        *string           `json:"name"`
		Rollup      *CalendarInterval `json:"rollup,omitempty"`
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
	o.Interval = all.Interval
	o.Metric = all.Metric
	o.Name = *all.Name
	o.Rollup = all.Rollup

	return nil
}

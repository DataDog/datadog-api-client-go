// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaUserJourneyQueryCompute Test-only generated schema for the dashboard versioning proof of concept.
type FormulaUserJourneyQueryCompute struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregation EventsAggregator `json:"aggregation"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Interval *float64 `json:"interval,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Metric *string `json:"metric,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Rollup *CalendarInterval `json:"rollup,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Target *GraphTarget `json:"target,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaUserJourneyQueryCompute instantiates a new FormulaUserJourneyQueryCompute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaUserJourneyQueryCompute(aggregation EventsAggregator) *FormulaUserJourneyQueryCompute {
	this := FormulaUserJourneyQueryCompute{}
	this.Aggregation = aggregation
	return &this
}

// NewFormulaUserJourneyQueryComputeWithDefaults instantiates a new FormulaUserJourneyQueryCompute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaUserJourneyQueryComputeWithDefaults() *FormulaUserJourneyQueryCompute {
	this := FormulaUserJourneyQueryCompute{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *FormulaUserJourneyQueryCompute) GetAggregation() EventsAggregator {
	if o == nil {
		var ret EventsAggregator
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *FormulaUserJourneyQueryCompute) GetAggregationOk() (*EventsAggregator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *FormulaUserJourneyQueryCompute) SetAggregation(v EventsAggregator) {
	o.Aggregation = v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *FormulaUserJourneyQueryCompute) GetInterval() float64 {
	if o == nil || o.Interval == nil {
		var ret float64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaUserJourneyQueryCompute) GetIntervalOk() (*float64, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *FormulaUserJourneyQueryCompute) HasInterval() bool {
	return o != nil && o.Interval != nil
}

// SetInterval gets a reference to the given float64 and assigns it to the Interval field.
func (o *FormulaUserJourneyQueryCompute) SetInterval(v float64) {
	o.Interval = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *FormulaUserJourneyQueryCompute) GetMetric() string {
	if o == nil || o.Metric == nil {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaUserJourneyQueryCompute) GetMetricOk() (*string, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *FormulaUserJourneyQueryCompute) HasMetric() bool {
	return o != nil && o.Metric != nil
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *FormulaUserJourneyQueryCompute) SetMetric(v string) {
	o.Metric = &v
}

// GetRollup returns the Rollup field value if set, zero value otherwise.
func (o *FormulaUserJourneyQueryCompute) GetRollup() CalendarInterval {
	if o == nil || o.Rollup == nil {
		var ret CalendarInterval
		return ret
	}
	return *o.Rollup
}

// GetRollupOk returns a tuple with the Rollup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaUserJourneyQueryCompute) GetRollupOk() (*CalendarInterval, bool) {
	if o == nil || o.Rollup == nil {
		return nil, false
	}
	return o.Rollup, true
}

// HasRollup returns a boolean if a field has been set.
func (o *FormulaUserJourneyQueryCompute) HasRollup() bool {
	return o != nil && o.Rollup != nil
}

// SetRollup gets a reference to the given CalendarInterval and assigns it to the Rollup field.
func (o *FormulaUserJourneyQueryCompute) SetRollup(v CalendarInterval) {
	o.Rollup = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *FormulaUserJourneyQueryCompute) GetTarget() GraphTarget {
	if o == nil || o.Target == nil {
		var ret GraphTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaUserJourneyQueryCompute) GetTargetOk() (*GraphTarget, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *FormulaUserJourneyQueryCompute) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given GraphTarget and assigns it to the Target field.
func (o *FormulaUserJourneyQueryCompute) SetTarget(v GraphTarget) {
	o.Target = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaUserJourneyQueryCompute) MarshalJSON() ([]byte, error) {
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
	if o.Rollup != nil {
		toSerialize["rollup"] = o.Rollup
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaUserJourneyQueryCompute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *EventsAggregator `json:"aggregation"`
		Interval    *float64          `json:"interval,omitempty"`
		Metric      *string           `json:"metric,omitempty"`
		Rollup      *CalendarInterval `json:"rollup,omitempty"`
		Target      *GraphTarget      `json:"target,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Aggregation == nil {
		return fmt.Errorf("required field aggregation missing")
	}
	o.Aggregation = *all.Aggregation
	o.Interval = all.Interval
	o.Metric = all.Metric
	o.Rollup = all.Rollup
	o.Target = all.Target

	return nil
}

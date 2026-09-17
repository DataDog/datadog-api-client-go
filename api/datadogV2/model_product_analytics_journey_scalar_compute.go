// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyScalarCompute Defines the metric computed over the journey for a scalar query.
type ProductAnalyticsJourneyScalarCompute struct {
	// Aggregation function: `count`, `cardinality`, `avg`, `median`, `min`, `max`, `sum`,
	// or a percentile of the form `pc<N>` such as `pc95`. Defaults to `cardinality`.
	Aggregation string `json:"aggregation"`
	// Metric to aggregate on. Use a facet path such as `@view.time_spent`, or one of the
	// journey metrics `__dd.conversion`, `__dd.conversion_rate`, `__dd.time_to_convert`,
	// or `__dd.dropoff_rate`. Defaults to `__dd.conversion`.
	Metric *string `json:"metric,omitempty"`
	// A reference to a step, or a range of steps, in the journey.
	// Use a `node` target to name a single step, or a `path` target to name the range
	// between two steps.
	Target *ProductAnalyticsJourneyTarget `json:"target,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsJourneyScalarCompute instantiates a new ProductAnalyticsJourneyScalarCompute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsJourneyScalarCompute(aggregation string) *ProductAnalyticsJourneyScalarCompute {
	this := ProductAnalyticsJourneyScalarCompute{}
	this.Aggregation = aggregation
	return &this
}

// NewProductAnalyticsJourneyScalarComputeWithDefaults instantiates a new ProductAnalyticsJourneyScalarCompute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsJourneyScalarComputeWithDefaults() *ProductAnalyticsJourneyScalarCompute {
	this := ProductAnalyticsJourneyScalarCompute{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *ProductAnalyticsJourneyScalarCompute) GetAggregation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyScalarCompute) GetAggregationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *ProductAnalyticsJourneyScalarCompute) SetAggregation(v string) {
	o.Aggregation = v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyScalarCompute) GetMetric() string {
	if o == nil || o.Metric == nil {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyScalarCompute) GetMetricOk() (*string, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyScalarCompute) HasMetric() bool {
	return o != nil && o.Metric != nil
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *ProductAnalyticsJourneyScalarCompute) SetMetric(v string) {
	o.Metric = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyScalarCompute) GetTarget() ProductAnalyticsJourneyTarget {
	if o == nil || o.Target == nil {
		var ret ProductAnalyticsJourneyTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyScalarCompute) GetTargetOk() (*ProductAnalyticsJourneyTarget, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyScalarCompute) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given ProductAnalyticsJourneyTarget and assigns it to the Target field.
func (o *ProductAnalyticsJourneyScalarCompute) SetTarget(v ProductAnalyticsJourneyTarget) {
	o.Target = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsJourneyScalarCompute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation"] = o.Aggregation
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsJourneyScalarCompute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *string                        `json:"aggregation"`
		Metric      *string                        `json:"metric,omitempty"`
		Target      *ProductAnalyticsJourneyTarget `json:"target,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Aggregation == nil {
		return fmt.Errorf("required field aggregation missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregation", "metric", "target"})
	} else {
		return err
	}
	o.Aggregation = *all.Aggregation
	o.Metric = all.Metric
	o.Target = all.Target

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

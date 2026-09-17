// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyFunnelCompute Defines the metric computed at each funnel step.
type ProductAnalyticsJourneyFunnelCompute struct {
	// Aggregation function: `count`, `cardinality`, `avg`, `median`, `min`, `max`, `sum`,
	// or a percentile of the form `pc<N>` such as `pc95`. Defaults to `cardinality`.
	Aggregation *string `json:"aggregation,omitempty"`
	// Metric to aggregate on. Defaults to the identity join key.
	Metric *string `json:"metric,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsJourneyFunnelCompute instantiates a new ProductAnalyticsJourneyFunnelCompute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsJourneyFunnelCompute() *ProductAnalyticsJourneyFunnelCompute {
	this := ProductAnalyticsJourneyFunnelCompute{}
	return &this
}

// NewProductAnalyticsJourneyFunnelComputeWithDefaults instantiates a new ProductAnalyticsJourneyFunnelCompute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsJourneyFunnelComputeWithDefaults() *ProductAnalyticsJourneyFunnelCompute {
	this := ProductAnalyticsJourneyFunnelCompute{}
	return &this
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyFunnelCompute) GetAggregation() string {
	if o == nil || o.Aggregation == nil {
		var ret string
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyFunnelCompute) GetAggregationOk() (*string, bool) {
	if o == nil || o.Aggregation == nil {
		return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyFunnelCompute) HasAggregation() bool {
	return o != nil && o.Aggregation != nil
}

// SetAggregation gets a reference to the given string and assigns it to the Aggregation field.
func (o *ProductAnalyticsJourneyFunnelCompute) SetAggregation(v string) {
	o.Aggregation = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyFunnelCompute) GetMetric() string {
	if o == nil || o.Metric == nil {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyFunnelCompute) GetMetricOk() (*string, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyFunnelCompute) HasMetric() bool {
	return o != nil && o.Metric != nil
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *ProductAnalyticsJourneyFunnelCompute) SetMetric(v string) {
	o.Metric = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsJourneyFunnelCompute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Aggregation != nil {
		toSerialize["aggregation"] = o.Aggregation
	}
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsJourneyFunnelCompute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *string `json:"aggregation,omitempty"`
		Metric      *string `json:"metric,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregation", "metric"})
	} else {
		return err
	}
	o.Aggregation = all.Aggregation
	o.Metric = all.Metric

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

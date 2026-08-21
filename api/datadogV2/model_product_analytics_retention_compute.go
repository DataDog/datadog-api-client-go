// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionCompute The metric and aggregation applied to a retention query.
type ProductAnalyticsRetentionCompute struct {
	// The aggregation function applied to the metric, such as `count` or `avg`.
	Aggregation string `json:"aggregation"`
	// The retention metric to compute, either an absolute count or a rate.
	Metric ProductAnalyticsRetentionComputeMetric `json:"metric"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsRetentionCompute instantiates a new ProductAnalyticsRetentionCompute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsRetentionCompute(aggregation string, metric ProductAnalyticsRetentionComputeMetric) *ProductAnalyticsRetentionCompute {
	this := ProductAnalyticsRetentionCompute{}
	this.Aggregation = aggregation
	this.Metric = metric
	return &this
}

// NewProductAnalyticsRetentionComputeWithDefaults instantiates a new ProductAnalyticsRetentionCompute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsRetentionComputeWithDefaults() *ProductAnalyticsRetentionCompute {
	this := ProductAnalyticsRetentionCompute{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *ProductAnalyticsRetentionCompute) GetAggregation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionCompute) GetAggregationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *ProductAnalyticsRetentionCompute) SetAggregation(v string) {
	o.Aggregation = v
}

// GetMetric returns the Metric field value.
func (o *ProductAnalyticsRetentionCompute) GetMetric() ProductAnalyticsRetentionComputeMetric {
	if o == nil {
		var ret ProductAnalyticsRetentionComputeMetric
		return ret
	}
	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionCompute) GetMetricOk() (*ProductAnalyticsRetentionComputeMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value.
func (o *ProductAnalyticsRetentionCompute) SetMetric(v ProductAnalyticsRetentionComputeMetric) {
	o.Metric = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsRetentionCompute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation"] = o.Aggregation
	toSerialize["metric"] = o.Metric

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsRetentionCompute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *string                                 `json:"aggregation"`
		Metric      *ProductAnalyticsRetentionComputeMetric `json:"metric"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregation", "metric"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Aggregation = *all.Aggregation
	if !all.Metric.IsValid() {
		hasInvalidField = true
	} else {
		o.Metric = *all.Metric
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

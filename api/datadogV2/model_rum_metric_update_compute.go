// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMMetricUpdateCompute The compute rule properties that can be updated. Only valid for metrics whose `aggregation_type` is `distribution`.
type RUMMetricUpdateCompute struct {
	// Toggle to include or exclude percentile aggregations (p50, p75, p90, p95, p99) for distribution metrics.
	// Required when `aggregation_type` is `distribution`, and must be omitted otherwise.
	IncludePercentiles *bool `json:"include_percentiles,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMMetricUpdateCompute instantiates a new RUMMetricUpdateCompute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMMetricUpdateCompute() *RUMMetricUpdateCompute {
	this := RUMMetricUpdateCompute{}
	return &this
}

// NewRUMMetricUpdateComputeWithDefaults instantiates a new RUMMetricUpdateCompute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMMetricUpdateComputeWithDefaults() *RUMMetricUpdateCompute {
	this := RUMMetricUpdateCompute{}
	return &this
}

// GetIncludePercentiles returns the IncludePercentiles field value if set, zero value otherwise.
func (o *RUMMetricUpdateCompute) GetIncludePercentiles() bool {
	if o == nil || o.IncludePercentiles == nil {
		var ret bool
		return ret
	}
	return *o.IncludePercentiles
}

// GetIncludePercentilesOk returns a tuple with the IncludePercentiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMMetricUpdateCompute) GetIncludePercentilesOk() (*bool, bool) {
	if o == nil || o.IncludePercentiles == nil {
		return nil, false
	}
	return o.IncludePercentiles, true
}

// HasIncludePercentiles returns a boolean if a field has been set.
func (o *RUMMetricUpdateCompute) HasIncludePercentiles() bool {
	return o != nil && o.IncludePercentiles != nil
}

// SetIncludePercentiles gets a reference to the given bool and assigns it to the IncludePercentiles field.
func (o *RUMMetricUpdateCompute) SetIncludePercentiles(v bool) {
	o.IncludePercentiles = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMMetricUpdateCompute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IncludePercentiles != nil {
		toSerialize["include_percentiles"] = o.IncludePercentiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMMetricUpdateCompute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncludePercentiles *bool `json:"include_percentiles,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"include_percentiles"})
	} else {
		return err
	}
	o.IncludePercentiles = all.IncludePercentiles

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

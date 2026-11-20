// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CompositeQuerySorting Test-only generated schema for the dashboard versioning proof of concept.
type CompositeQuerySorting struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregation EventsAggregator `json:"aggregation"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Metric *string `json:"metric,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Order *QuerySortOrder `json:"order,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Source *string `json:"source,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCompositeQuerySorting instantiates a new CompositeQuerySorting object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCompositeQuerySorting(aggregation EventsAggregator) *CompositeQuerySorting {
	this := CompositeQuerySorting{}
	this.Aggregation = aggregation
	return &this
}

// NewCompositeQuerySortingWithDefaults instantiates a new CompositeQuerySorting object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCompositeQuerySortingWithDefaults() *CompositeQuerySorting {
	this := CompositeQuerySorting{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *CompositeQuerySorting) GetAggregation() EventsAggregator {
	if o == nil {
		var ret EventsAggregator
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *CompositeQuerySorting) GetAggregationOk() (*EventsAggregator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *CompositeQuerySorting) SetAggregation(v EventsAggregator) {
	o.Aggregation = v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *CompositeQuerySorting) GetMetric() string {
	if o == nil || o.Metric == nil {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeQuerySorting) GetMetricOk() (*string, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *CompositeQuerySorting) HasMetric() bool {
	return o != nil && o.Metric != nil
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *CompositeQuerySorting) SetMetric(v string) {
	o.Metric = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *CompositeQuerySorting) GetOrder() QuerySortOrder {
	if o == nil || o.Order == nil {
		var ret QuerySortOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeQuerySorting) GetOrderOk() (*QuerySortOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *CompositeQuerySorting) HasOrder() bool {
	return o != nil && o.Order != nil
}

// SetOrder gets a reference to the given QuerySortOrder and assigns it to the Order field.
func (o *CompositeQuerySorting) SetOrder(v QuerySortOrder) {
	o.Order = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CompositeQuerySorting) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeQuerySorting) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CompositeQuerySorting) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *CompositeQuerySorting) SetSource(v string) {
	o.Source = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CompositeQuerySorting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation"] = o.Aggregation
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CompositeQuerySorting) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *EventsAggregator `json:"aggregation"`
		Metric      *string           `json:"metric,omitempty"`
		Order       *QuerySortOrder   `json:"order,omitempty"`
		Source      *string           `json:"source,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Aggregation == nil {
		return fmt.Errorf("required field aggregation missing")
	}

	hasInvalidField := false
	o.Aggregation = *all.Aggregation
	o.Metric = all.Metric
	if all.Order != nil && !all.Order.IsValid() {
		hasInvalidField = true
	} else {
		o.Order = all.Order
	}
	o.Source = all.Source

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

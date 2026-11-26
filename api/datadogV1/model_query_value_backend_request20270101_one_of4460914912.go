// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryValueBackendRequest20270101OneOf4460914912 Test-only generated schema for the dashboard versioning proof of concept.
type QueryValueBackendRequest20270101OneOf4460914912 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregator *ScalarAggregator_20270101 `json:"aggregator,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Comparison *ComparisonRequest_20270101 `json:"comparison,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule_20270101 `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SyntheticsBatches LogQueryDefinition_20270101 `json:"synthetics_batches"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewQueryValueBackendRequest20270101OneOf4460914912 instantiates a new QueryValueBackendRequest20270101OneOf4460914912 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueryValueBackendRequest20270101OneOf4460914912(syntheticsBatches LogQueryDefinition_20270101) *QueryValueBackendRequest20270101OneOf4460914912 {
	this := QueryValueBackendRequest20270101OneOf4460914912{}
	this.SyntheticsBatches = syntheticsBatches
	return &this
}

// NewQueryValueBackendRequest20270101OneOf4460914912WithDefaults instantiates a new QueryValueBackendRequest20270101OneOf4460914912 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueryValueBackendRequest20270101OneOf4460914912WithDefaults() *QueryValueBackendRequest20270101OneOf4460914912 {
	this := QueryValueBackendRequest20270101OneOf4460914912{}
	return &this
}

// GetAggregator returns the Aggregator field value if set, zero value otherwise.
func (o *QueryValueBackendRequest20270101OneOf4460914912) GetAggregator() ScalarAggregator_20270101 {
	if o == nil || o.Aggregator == nil {
		var ret ScalarAggregator_20270101
		return ret
	}
	return *o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460914912) GetAggregatorOk() (*ScalarAggregator_20270101, bool) {
	if o == nil || o.Aggregator == nil {
		return nil, false
	}
	return o.Aggregator, true
}

// HasAggregator returns a boolean if a field has been set.
func (o *QueryValueBackendRequest20270101OneOf4460914912) HasAggregator() bool {
	return o != nil && o.Aggregator != nil
}

// SetAggregator gets a reference to the given ScalarAggregator_20270101 and assigns it to the Aggregator field.
func (o *QueryValueBackendRequest20270101OneOf4460914912) SetAggregator(v ScalarAggregator_20270101) {
	o.Aggregator = &v
}

// GetComparison returns the Comparison field value if set, zero value otherwise.
func (o *QueryValueBackendRequest20270101OneOf4460914912) GetComparison() ComparisonRequest_20270101 {
	if o == nil || o.Comparison == nil {
		var ret ComparisonRequest_20270101
		return ret
	}
	return *o.Comparison
}

// GetComparisonOk returns a tuple with the Comparison field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460914912) GetComparisonOk() (*ComparisonRequest_20270101, bool) {
	if o == nil || o.Comparison == nil {
		return nil, false
	}
	return o.Comparison, true
}

// HasComparison returns a boolean if a field has been set.
func (o *QueryValueBackendRequest20270101OneOf4460914912) HasComparison() bool {
	return o != nil && o.Comparison != nil
}

// SetComparison gets a reference to the given ComparisonRequest_20270101 and assigns it to the Comparison field.
func (o *QueryValueBackendRequest20270101OneOf4460914912) SetComparison(v ComparisonRequest_20270101) {
	o.Comparison = &v
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *QueryValueBackendRequest20270101OneOf4460914912) GetConditionalFormats() []ThresholdFormattingRule_20270101 {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule_20270101
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460914912) GetConditionalFormatsOk() (*[]ThresholdFormattingRule_20270101, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *QueryValueBackendRequest20270101OneOf4460914912) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule_20270101 and assigns it to the ConditionalFormats field.
func (o *QueryValueBackendRequest20270101OneOf4460914912) SetConditionalFormats(v []ThresholdFormattingRule_20270101) {
	o.ConditionalFormats = v
}

// GetSyntheticsBatches returns the SyntheticsBatches field value.
func (o *QueryValueBackendRequest20270101OneOf4460914912) GetSyntheticsBatches() LogQueryDefinition_20270101 {
	if o == nil {
		var ret LogQueryDefinition_20270101
		return ret
	}
	return o.SyntheticsBatches
}

// GetSyntheticsBatchesOk returns a tuple with the SyntheticsBatches field value
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460914912) GetSyntheticsBatchesOk() (*LogQueryDefinition_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyntheticsBatches, true
}

// SetSyntheticsBatches sets field value.
func (o *QueryValueBackendRequest20270101OneOf4460914912) SetSyntheticsBatches(v LogQueryDefinition_20270101) {
	o.SyntheticsBatches = v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueryValueBackendRequest20270101OneOf4460914912) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Aggregator != nil {
		toSerialize["aggregator"] = o.Aggregator
	}
	if o.Comparison != nil {
		toSerialize["comparison"] = o.Comparison
	}
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	toSerialize["synthetics_batches"] = o.SyntheticsBatches
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *QueryValueBackendRequest20270101OneOf4460914912) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregator         *ScalarAggregator_20270101         `json:"aggregator,omitempty"`
		Comparison         *ComparisonRequest_20270101        `json:"comparison,omitempty"`
		ConditionalFormats []ThresholdFormattingRule_20270101 `json:"conditional_formats,omitempty"`
		SyntheticsBatches  *LogQueryDefinition_20270101       `json:"synthetics_batches"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SyntheticsBatches == nil {
		return fmt.Errorf("required field synthetics_batches missing")
	}

	hasInvalidField := false
	if all.Aggregator != nil && !all.Aggregator.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregator = all.Aggregator
	}
	o.Comparison = all.Comparison
	o.ConditionalFormats = all.ConditionalFormats
	o.SyntheticsBatches = *all.SyntheticsBatches

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

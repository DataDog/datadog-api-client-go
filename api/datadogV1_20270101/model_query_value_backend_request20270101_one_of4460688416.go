// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryValueBackendRequest20270101OneOf4460688416 Test-only generated schema for the dashboard versioning proof of concept.
type QueryValueBackendRequest20270101OneOf4460688416 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregator *ScalarAggregator `json:"aggregator,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Comparison *ComparisonRequest `json:"comparison,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	LogQuery LogQueryDefinition `json:"log_query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewQueryValueBackendRequest20270101OneOf4460688416 instantiates a new QueryValueBackendRequest20270101OneOf4460688416 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueryValueBackendRequest20270101OneOf4460688416(logQuery LogQueryDefinition) *QueryValueBackendRequest20270101OneOf4460688416 {
	this := QueryValueBackendRequest20270101OneOf4460688416{}
	this.LogQuery = logQuery
	return &this
}

// NewQueryValueBackendRequest20270101OneOf4460688416WithDefaults instantiates a new QueryValueBackendRequest20270101OneOf4460688416 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueryValueBackendRequest20270101OneOf4460688416WithDefaults() *QueryValueBackendRequest20270101OneOf4460688416 {
	this := QueryValueBackendRequest20270101OneOf4460688416{}
	return &this
}

// GetAggregator returns the Aggregator field value if set, zero value otherwise.
func (o *QueryValueBackendRequest20270101OneOf4460688416) GetAggregator() ScalarAggregator {
	if o == nil || o.Aggregator == nil {
		var ret ScalarAggregator
		return ret
	}
	return *o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460688416) GetAggregatorOk() (*ScalarAggregator, bool) {
	if o == nil || o.Aggregator == nil {
		return nil, false
	}
	return o.Aggregator, true
}

// HasAggregator returns a boolean if a field has been set.
func (o *QueryValueBackendRequest20270101OneOf4460688416) HasAggregator() bool {
	return o != nil && o.Aggregator != nil
}

// SetAggregator gets a reference to the given ScalarAggregator and assigns it to the Aggregator field.
func (o *QueryValueBackendRequest20270101OneOf4460688416) SetAggregator(v ScalarAggregator) {
	o.Aggregator = &v
}

// GetComparison returns the Comparison field value if set, zero value otherwise.
func (o *QueryValueBackendRequest20270101OneOf4460688416) GetComparison() ComparisonRequest {
	if o == nil || o.Comparison == nil {
		var ret ComparisonRequest
		return ret
	}
	return *o.Comparison
}

// GetComparisonOk returns a tuple with the Comparison field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460688416) GetComparisonOk() (*ComparisonRequest, bool) {
	if o == nil || o.Comparison == nil {
		return nil, false
	}
	return o.Comparison, true
}

// HasComparison returns a boolean if a field has been set.
func (o *QueryValueBackendRequest20270101OneOf4460688416) HasComparison() bool {
	return o != nil && o.Comparison != nil
}

// SetComparison gets a reference to the given ComparisonRequest and assigns it to the Comparison field.
func (o *QueryValueBackendRequest20270101OneOf4460688416) SetComparison(v ComparisonRequest) {
	o.Comparison = &v
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *QueryValueBackendRequest20270101OneOf4460688416) GetConditionalFormats() []ThresholdFormattingRule {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460688416) GetConditionalFormatsOk() (*[]ThresholdFormattingRule, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *QueryValueBackendRequest20270101OneOf4460688416) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule and assigns it to the ConditionalFormats field.
func (o *QueryValueBackendRequest20270101OneOf4460688416) SetConditionalFormats(v []ThresholdFormattingRule) {
	o.ConditionalFormats = v
}

// GetLogQuery returns the LogQuery field value.
func (o *QueryValueBackendRequest20270101OneOf4460688416) GetLogQuery() LogQueryDefinition {
	if o == nil {
		var ret LogQueryDefinition
		return ret
	}
	return o.LogQuery
}

// GetLogQueryOk returns a tuple with the LogQuery field value
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460688416) GetLogQueryOk() (*LogQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogQuery, true
}

// SetLogQuery sets field value.
func (o *QueryValueBackendRequest20270101OneOf4460688416) SetLogQuery(v LogQueryDefinition) {
	o.LogQuery = v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueryValueBackendRequest20270101OneOf4460688416) MarshalJSON() ([]byte, error) {
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
	toSerialize["log_query"] = o.LogQuery
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *QueryValueBackendRequest20270101OneOf4460688416) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregator         *ScalarAggregator         `json:"aggregator,omitempty"`
		Comparison         *ComparisonRequest        `json:"comparison,omitempty"`
		ConditionalFormats []ThresholdFormattingRule `json:"conditional_formats,omitempty"`
		LogQuery           *LogQueryDefinition       `json:"log_query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.LogQuery == nil {
		return fmt.Errorf("required field log_query missing")
	}

	hasInvalidField := false
	if all.Aggregator != nil && !all.Aggregator.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregator = all.Aggregator
	}
	if all.Comparison != nil && all.Comparison.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Comparison = all.Comparison
	o.ConditionalFormats = all.ConditionalFormats
	if all.LogQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LogQuery = *all.LogQuery

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

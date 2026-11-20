// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryValueBackendRequest20270101OneOf4460689536 Test-only generated schema for the dashboard versioning proof of concept.
type QueryValueBackendRequest20270101OneOf4460689536 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregator *ScalarAggregator `json:"aggregator,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ApmQuery LogQueryDefinition `json:"apm_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Comparison *ComparisonRequest `json:"comparison,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule `json:"conditional_formats,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewQueryValueBackendRequest20270101OneOf4460689536 instantiates a new QueryValueBackendRequest20270101OneOf4460689536 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueryValueBackendRequest20270101OneOf4460689536(apmQuery LogQueryDefinition) *QueryValueBackendRequest20270101OneOf4460689536 {
	this := QueryValueBackendRequest20270101OneOf4460689536{}
	this.ApmQuery = apmQuery
	return &this
}

// NewQueryValueBackendRequest20270101OneOf4460689536WithDefaults instantiates a new QueryValueBackendRequest20270101OneOf4460689536 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueryValueBackendRequest20270101OneOf4460689536WithDefaults() *QueryValueBackendRequest20270101OneOf4460689536 {
	this := QueryValueBackendRequest20270101OneOf4460689536{}
	return &this
}

// GetAggregator returns the Aggregator field value if set, zero value otherwise.
func (o *QueryValueBackendRequest20270101OneOf4460689536) GetAggregator() ScalarAggregator {
	if o == nil || o.Aggregator == nil {
		var ret ScalarAggregator
		return ret
	}
	return *o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460689536) GetAggregatorOk() (*ScalarAggregator, bool) {
	if o == nil || o.Aggregator == nil {
		return nil, false
	}
	return o.Aggregator, true
}

// HasAggregator returns a boolean if a field has been set.
func (o *QueryValueBackendRequest20270101OneOf4460689536) HasAggregator() bool {
	return o != nil && o.Aggregator != nil
}

// SetAggregator gets a reference to the given ScalarAggregator and assigns it to the Aggregator field.
func (o *QueryValueBackendRequest20270101OneOf4460689536) SetAggregator(v ScalarAggregator) {
	o.Aggregator = &v
}

// GetApmQuery returns the ApmQuery field value.
func (o *QueryValueBackendRequest20270101OneOf4460689536) GetApmQuery() LogQueryDefinition {
	if o == nil {
		var ret LogQueryDefinition
		return ret
	}
	return o.ApmQuery
}

// GetApmQueryOk returns a tuple with the ApmQuery field value
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460689536) GetApmQueryOk() (*LogQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApmQuery, true
}

// SetApmQuery sets field value.
func (o *QueryValueBackendRequest20270101OneOf4460689536) SetApmQuery(v LogQueryDefinition) {
	o.ApmQuery = v
}

// GetComparison returns the Comparison field value if set, zero value otherwise.
func (o *QueryValueBackendRequest20270101OneOf4460689536) GetComparison() ComparisonRequest {
	if o == nil || o.Comparison == nil {
		var ret ComparisonRequest
		return ret
	}
	return *o.Comparison
}

// GetComparisonOk returns a tuple with the Comparison field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460689536) GetComparisonOk() (*ComparisonRequest, bool) {
	if o == nil || o.Comparison == nil {
		return nil, false
	}
	return o.Comparison, true
}

// HasComparison returns a boolean if a field has been set.
func (o *QueryValueBackendRequest20270101OneOf4460689536) HasComparison() bool {
	return o != nil && o.Comparison != nil
}

// SetComparison gets a reference to the given ComparisonRequest and assigns it to the Comparison field.
func (o *QueryValueBackendRequest20270101OneOf4460689536) SetComparison(v ComparisonRequest) {
	o.Comparison = &v
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *QueryValueBackendRequest20270101OneOf4460689536) GetConditionalFormats() []ThresholdFormattingRule {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460689536) GetConditionalFormatsOk() (*[]ThresholdFormattingRule, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *QueryValueBackendRequest20270101OneOf4460689536) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule and assigns it to the ConditionalFormats field.
func (o *QueryValueBackendRequest20270101OneOf4460689536) SetConditionalFormats(v []ThresholdFormattingRule) {
	o.ConditionalFormats = v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueryValueBackendRequest20270101OneOf4460689536) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Aggregator != nil {
		toSerialize["aggregator"] = o.Aggregator
	}
	toSerialize["apm_query"] = o.ApmQuery
	if o.Comparison != nil {
		toSerialize["comparison"] = o.Comparison
	}
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *QueryValueBackendRequest20270101OneOf4460689536) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregator         *ScalarAggregator         `json:"aggregator,omitempty"`
		ApmQuery           *LogQueryDefinition       `json:"apm_query"`
		Comparison         *ComparisonRequest        `json:"comparison,omitempty"`
		ConditionalFormats []ThresholdFormattingRule `json:"conditional_formats,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApmQuery == nil {
		return fmt.Errorf("required field apm_query missing")
	}

	hasInvalidField := false
	if all.Aggregator != nil && !all.Aggregator.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregator = all.Aggregator
	}
	if all.ApmQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ApmQuery = *all.ApmQuery
	if all.Comparison != nil && all.Comparison.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Comparison = all.Comparison
	o.ConditionalFormats = all.ConditionalFormats

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

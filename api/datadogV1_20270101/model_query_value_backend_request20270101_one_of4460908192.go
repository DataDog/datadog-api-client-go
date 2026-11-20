// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryValueBackendRequest20270101OneOf4460908192 Test-only generated schema for the dashboard versioning proof of concept.
type QueryValueBackendRequest20270101OneOf4460908192 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregator *ScalarAggregator `json:"aggregator,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Comparison *ComparisonRequest `json:"comparison,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	EventQuery LogQueryDefinition `json:"event_query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewQueryValueBackendRequest20270101OneOf4460908192 instantiates a new QueryValueBackendRequest20270101OneOf4460908192 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueryValueBackendRequest20270101OneOf4460908192(eventQuery LogQueryDefinition) *QueryValueBackendRequest20270101OneOf4460908192 {
	this := QueryValueBackendRequest20270101OneOf4460908192{}
	this.EventQuery = eventQuery
	return &this
}

// NewQueryValueBackendRequest20270101OneOf4460908192WithDefaults instantiates a new QueryValueBackendRequest20270101OneOf4460908192 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueryValueBackendRequest20270101OneOf4460908192WithDefaults() *QueryValueBackendRequest20270101OneOf4460908192 {
	this := QueryValueBackendRequest20270101OneOf4460908192{}
	return &this
}

// GetAggregator returns the Aggregator field value if set, zero value otherwise.
func (o *QueryValueBackendRequest20270101OneOf4460908192) GetAggregator() ScalarAggregator {
	if o == nil || o.Aggregator == nil {
		var ret ScalarAggregator
		return ret
	}
	return *o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460908192) GetAggregatorOk() (*ScalarAggregator, bool) {
	if o == nil || o.Aggregator == nil {
		return nil, false
	}
	return o.Aggregator, true
}

// HasAggregator returns a boolean if a field has been set.
func (o *QueryValueBackendRequest20270101OneOf4460908192) HasAggregator() bool {
	return o != nil && o.Aggregator != nil
}

// SetAggregator gets a reference to the given ScalarAggregator and assigns it to the Aggregator field.
func (o *QueryValueBackendRequest20270101OneOf4460908192) SetAggregator(v ScalarAggregator) {
	o.Aggregator = &v
}

// GetComparison returns the Comparison field value if set, zero value otherwise.
func (o *QueryValueBackendRequest20270101OneOf4460908192) GetComparison() ComparisonRequest {
	if o == nil || o.Comparison == nil {
		var ret ComparisonRequest
		return ret
	}
	return *o.Comparison
}

// GetComparisonOk returns a tuple with the Comparison field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460908192) GetComparisonOk() (*ComparisonRequest, bool) {
	if o == nil || o.Comparison == nil {
		return nil, false
	}
	return o.Comparison, true
}

// HasComparison returns a boolean if a field has been set.
func (o *QueryValueBackendRequest20270101OneOf4460908192) HasComparison() bool {
	return o != nil && o.Comparison != nil
}

// SetComparison gets a reference to the given ComparisonRequest and assigns it to the Comparison field.
func (o *QueryValueBackendRequest20270101OneOf4460908192) SetComparison(v ComparisonRequest) {
	o.Comparison = &v
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *QueryValueBackendRequest20270101OneOf4460908192) GetConditionalFormats() []ThresholdFormattingRule {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460908192) GetConditionalFormatsOk() (*[]ThresholdFormattingRule, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *QueryValueBackendRequest20270101OneOf4460908192) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule and assigns it to the ConditionalFormats field.
func (o *QueryValueBackendRequest20270101OneOf4460908192) SetConditionalFormats(v []ThresholdFormattingRule) {
	o.ConditionalFormats = v
}

// GetEventQuery returns the EventQuery field value.
func (o *QueryValueBackendRequest20270101OneOf4460908192) GetEventQuery() LogQueryDefinition {
	if o == nil {
		var ret LogQueryDefinition
		return ret
	}
	return o.EventQuery
}

// GetEventQueryOk returns a tuple with the EventQuery field value
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460908192) GetEventQueryOk() (*LogQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventQuery, true
}

// SetEventQuery sets field value.
func (o *QueryValueBackendRequest20270101OneOf4460908192) SetEventQuery(v LogQueryDefinition) {
	o.EventQuery = v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueryValueBackendRequest20270101OneOf4460908192) MarshalJSON() ([]byte, error) {
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
	toSerialize["event_query"] = o.EventQuery
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *QueryValueBackendRequest20270101OneOf4460908192) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregator         *ScalarAggregator         `json:"aggregator,omitempty"`
		Comparison         *ComparisonRequest        `json:"comparison,omitempty"`
		ConditionalFormats []ThresholdFormattingRule `json:"conditional_formats,omitempty"`
		EventQuery         *LogQueryDefinition       `json:"event_query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EventQuery == nil {
		return fmt.Errorf("required field event_query missing")
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
	if all.EventQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.EventQuery = *all.EventQuery

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

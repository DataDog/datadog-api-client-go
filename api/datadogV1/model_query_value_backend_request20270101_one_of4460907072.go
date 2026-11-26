// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryValueBackendRequest20270101OneOf4460907072 Test-only generated schema for the dashboard versioning proof of concept.
type QueryValueBackendRequest20270101OneOf4460907072 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregator *ScalarAggregator_20270101 `json:"aggregator,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Comparison *ComparisonRequest_20270101 `json:"comparison,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ComplianceFindingsQuery LogQueryDefinition_20270101 `json:"compliance_findings_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule_20270101 `json:"conditional_formats,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewQueryValueBackendRequest20270101OneOf4460907072 instantiates a new QueryValueBackendRequest20270101OneOf4460907072 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueryValueBackendRequest20270101OneOf4460907072(complianceFindingsQuery LogQueryDefinition_20270101) *QueryValueBackendRequest20270101OneOf4460907072 {
	this := QueryValueBackendRequest20270101OneOf4460907072{}
	this.ComplianceFindingsQuery = complianceFindingsQuery
	return &this
}

// NewQueryValueBackendRequest20270101OneOf4460907072WithDefaults instantiates a new QueryValueBackendRequest20270101OneOf4460907072 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueryValueBackendRequest20270101OneOf4460907072WithDefaults() *QueryValueBackendRequest20270101OneOf4460907072 {
	this := QueryValueBackendRequest20270101OneOf4460907072{}
	return &this
}

// GetAggregator returns the Aggregator field value if set, zero value otherwise.
func (o *QueryValueBackendRequest20270101OneOf4460907072) GetAggregator() ScalarAggregator_20270101 {
	if o == nil || o.Aggregator == nil {
		var ret ScalarAggregator_20270101
		return ret
	}
	return *o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460907072) GetAggregatorOk() (*ScalarAggregator_20270101, bool) {
	if o == nil || o.Aggregator == nil {
		return nil, false
	}
	return o.Aggregator, true
}

// HasAggregator returns a boolean if a field has been set.
func (o *QueryValueBackendRequest20270101OneOf4460907072) HasAggregator() bool {
	return o != nil && o.Aggregator != nil
}

// SetAggregator gets a reference to the given ScalarAggregator_20270101 and assigns it to the Aggregator field.
func (o *QueryValueBackendRequest20270101OneOf4460907072) SetAggregator(v ScalarAggregator_20270101) {
	o.Aggregator = &v
}

// GetComparison returns the Comparison field value if set, zero value otherwise.
func (o *QueryValueBackendRequest20270101OneOf4460907072) GetComparison() ComparisonRequest_20270101 {
	if o == nil || o.Comparison == nil {
		var ret ComparisonRequest_20270101
		return ret
	}
	return *o.Comparison
}

// GetComparisonOk returns a tuple with the Comparison field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460907072) GetComparisonOk() (*ComparisonRequest_20270101, bool) {
	if o == nil || o.Comparison == nil {
		return nil, false
	}
	return o.Comparison, true
}

// HasComparison returns a boolean if a field has been set.
func (o *QueryValueBackendRequest20270101OneOf4460907072) HasComparison() bool {
	return o != nil && o.Comparison != nil
}

// SetComparison gets a reference to the given ComparisonRequest_20270101 and assigns it to the Comparison field.
func (o *QueryValueBackendRequest20270101OneOf4460907072) SetComparison(v ComparisonRequest_20270101) {
	o.Comparison = &v
}

// GetComplianceFindingsQuery returns the ComplianceFindingsQuery field value.
func (o *QueryValueBackendRequest20270101OneOf4460907072) GetComplianceFindingsQuery() LogQueryDefinition_20270101 {
	if o == nil {
		var ret LogQueryDefinition_20270101
		return ret
	}
	return o.ComplianceFindingsQuery
}

// GetComplianceFindingsQueryOk returns a tuple with the ComplianceFindingsQuery field value
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460907072) GetComplianceFindingsQueryOk() (*LogQueryDefinition_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComplianceFindingsQuery, true
}

// SetComplianceFindingsQuery sets field value.
func (o *QueryValueBackendRequest20270101OneOf4460907072) SetComplianceFindingsQuery(v LogQueryDefinition_20270101) {
	o.ComplianceFindingsQuery = v
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *QueryValueBackendRequest20270101OneOf4460907072) GetConditionalFormats() []ThresholdFormattingRule_20270101 {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule_20270101
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueBackendRequest20270101OneOf4460907072) GetConditionalFormatsOk() (*[]ThresholdFormattingRule_20270101, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *QueryValueBackendRequest20270101OneOf4460907072) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule_20270101 and assigns it to the ConditionalFormats field.
func (o *QueryValueBackendRequest20270101OneOf4460907072) SetConditionalFormats(v []ThresholdFormattingRule_20270101) {
	o.ConditionalFormats = v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueryValueBackendRequest20270101OneOf4460907072) MarshalJSON() ([]byte, error) {
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
	toSerialize["compliance_findings_query"] = o.ComplianceFindingsQuery
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *QueryValueBackendRequest20270101OneOf4460907072) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregator              *ScalarAggregator_20270101         `json:"aggregator,omitempty"`
		Comparison              *ComparisonRequest_20270101        `json:"comparison,omitempty"`
		ComplianceFindingsQuery *LogQueryDefinition_20270101       `json:"compliance_findings_query"`
		ConditionalFormats      []ThresholdFormattingRule_20270101 `json:"conditional_formats,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ComplianceFindingsQuery == nil {
		return fmt.Errorf("required field compliance_findings_query missing")
	}

	hasInvalidField := false
	if all.Aggregator != nil && !all.Aggregator.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregator = all.Aggregator
	}
	o.Comparison = all.Comparison
	o.ComplianceFindingsQuery = *all.ComplianceFindingsQuery
	o.ConditionalFormats = all.ConditionalFormats

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

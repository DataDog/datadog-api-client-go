// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistBackendRequest20270101OneOf4463329120 Test-only generated schema for the dashboard versioning proof of concept.
type ToplistBackendRequest20270101OneOf4463329120 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ComplianceFindingsQuery LogQueryDefinition_20270101 `json:"compliance_findings_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule_20270101 `json:"conditional_formats,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewToplistBackendRequest20270101OneOf4463329120 instantiates a new ToplistBackendRequest20270101OneOf4463329120 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewToplistBackendRequest20270101OneOf4463329120(complianceFindingsQuery LogQueryDefinition_20270101) *ToplistBackendRequest20270101OneOf4463329120 {
	this := ToplistBackendRequest20270101OneOf4463329120{}
	this.ComplianceFindingsQuery = complianceFindingsQuery
	return &this
}

// NewToplistBackendRequest20270101OneOf4463329120WithDefaults instantiates a new ToplistBackendRequest20270101OneOf4463329120 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewToplistBackendRequest20270101OneOf4463329120WithDefaults() *ToplistBackendRequest20270101OneOf4463329120 {
	this := ToplistBackendRequest20270101OneOf4463329120{}
	return &this
}

// GetComplianceFindingsQuery returns the ComplianceFindingsQuery field value.
func (o *ToplistBackendRequest20270101OneOf4463329120) GetComplianceFindingsQuery() LogQueryDefinition_20270101 {
	if o == nil {
		var ret LogQueryDefinition_20270101
		return ret
	}
	return o.ComplianceFindingsQuery
}

// GetComplianceFindingsQueryOk returns a tuple with the ComplianceFindingsQuery field value
// and a boolean to check if the value has been set.
func (o *ToplistBackendRequest20270101OneOf4463329120) GetComplianceFindingsQueryOk() (*LogQueryDefinition_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComplianceFindingsQuery, true
}

// SetComplianceFindingsQuery sets field value.
func (o *ToplistBackendRequest20270101OneOf4463329120) SetComplianceFindingsQuery(v LogQueryDefinition_20270101) {
	o.ComplianceFindingsQuery = v
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *ToplistBackendRequest20270101OneOf4463329120) GetConditionalFormats() []ThresholdFormattingRule_20270101 {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule_20270101
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToplistBackendRequest20270101OneOf4463329120) GetConditionalFormatsOk() (*[]ThresholdFormattingRule_20270101, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *ToplistBackendRequest20270101OneOf4463329120) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule_20270101 and assigns it to the ConditionalFormats field.
func (o *ToplistBackendRequest20270101OneOf4463329120) SetConditionalFormats(v []ThresholdFormattingRule_20270101) {
	o.ConditionalFormats = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ToplistBackendRequest20270101OneOf4463329120) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["compliance_findings_query"] = o.ComplianceFindingsQuery
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ToplistBackendRequest20270101OneOf4463329120) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComplianceFindingsQuery *LogQueryDefinition_20270101       `json:"compliance_findings_query"`
		ConditionalFormats      []ThresholdFormattingRule_20270101 `json:"conditional_formats,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ComplianceFindingsQuery == nil {
		return fmt.Errorf("required field compliance_findings_query missing")
	}
	o.ComplianceFindingsQuery = *all.ComplianceFindingsQuery
	o.ConditionalFormats = all.ConditionalFormats

	return nil
}

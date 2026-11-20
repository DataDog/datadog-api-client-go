// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistBackendRequest20270101OneOf4463528192 Test-only generated schema for the dashboard versioning proof of concept.
type ToplistBackendRequest20270101OneOf4463528192 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	CiPipelineQuery LogQueryDefinition_20270101 `json:"ci_pipeline_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule_20270101 `json:"conditional_formats,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewToplistBackendRequest20270101OneOf4463528192 instantiates a new ToplistBackendRequest20270101OneOf4463528192 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewToplistBackendRequest20270101OneOf4463528192(ciPipelineQuery LogQueryDefinition_20270101) *ToplistBackendRequest20270101OneOf4463528192 {
	this := ToplistBackendRequest20270101OneOf4463528192{}
	this.CiPipelineQuery = ciPipelineQuery
	return &this
}

// NewToplistBackendRequest20270101OneOf4463528192WithDefaults instantiates a new ToplistBackendRequest20270101OneOf4463528192 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewToplistBackendRequest20270101OneOf4463528192WithDefaults() *ToplistBackendRequest20270101OneOf4463528192 {
	this := ToplistBackendRequest20270101OneOf4463528192{}
	return &this
}

// GetCiPipelineQuery returns the CiPipelineQuery field value.
func (o *ToplistBackendRequest20270101OneOf4463528192) GetCiPipelineQuery() LogQueryDefinition_20270101 {
	if o == nil {
		var ret LogQueryDefinition_20270101
		return ret
	}
	return o.CiPipelineQuery
}

// GetCiPipelineQueryOk returns a tuple with the CiPipelineQuery field value
// and a boolean to check if the value has been set.
func (o *ToplistBackendRequest20270101OneOf4463528192) GetCiPipelineQueryOk() (*LogQueryDefinition_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CiPipelineQuery, true
}

// SetCiPipelineQuery sets field value.
func (o *ToplistBackendRequest20270101OneOf4463528192) SetCiPipelineQuery(v LogQueryDefinition_20270101) {
	o.CiPipelineQuery = v
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *ToplistBackendRequest20270101OneOf4463528192) GetConditionalFormats() []ThresholdFormattingRule_20270101 {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule_20270101
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToplistBackendRequest20270101OneOf4463528192) GetConditionalFormatsOk() (*[]ThresholdFormattingRule_20270101, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *ToplistBackendRequest20270101OneOf4463528192) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule_20270101 and assigns it to the ConditionalFormats field.
func (o *ToplistBackendRequest20270101OneOf4463528192) SetConditionalFormats(v []ThresholdFormattingRule_20270101) {
	o.ConditionalFormats = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ToplistBackendRequest20270101OneOf4463528192) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["ci_pipeline_query"] = o.CiPipelineQuery
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ToplistBackendRequest20270101OneOf4463528192) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CiPipelineQuery    *LogQueryDefinition_20270101       `json:"ci_pipeline_query"`
		ConditionalFormats []ThresholdFormattingRule_20270101 `json:"conditional_formats,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CiPipelineQuery == nil {
		return fmt.Errorf("required field ci_pipeline_query missing")
	}
	o.CiPipelineQuery = *all.CiPipelineQuery
	o.ConditionalFormats = all.ConditionalFormats

	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistBackendRequest20270101OneOf4463528992 Test-only generated schema for the dashboard versioning proof of concept.
type ToplistBackendRequest20270101OneOf4463528992 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	CiTestQuery LogQueryDefinition_20270101 `json:"ci_test_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule_20270101 `json:"conditional_formats,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewToplistBackendRequest20270101OneOf4463528992 instantiates a new ToplistBackendRequest20270101OneOf4463528992 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewToplistBackendRequest20270101OneOf4463528992(ciTestQuery LogQueryDefinition_20270101) *ToplistBackendRequest20270101OneOf4463528992 {
	this := ToplistBackendRequest20270101OneOf4463528992{}
	this.CiTestQuery = ciTestQuery
	return &this
}

// NewToplistBackendRequest20270101OneOf4463528992WithDefaults instantiates a new ToplistBackendRequest20270101OneOf4463528992 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewToplistBackendRequest20270101OneOf4463528992WithDefaults() *ToplistBackendRequest20270101OneOf4463528992 {
	this := ToplistBackendRequest20270101OneOf4463528992{}
	return &this
}

// GetCiTestQuery returns the CiTestQuery field value.
func (o *ToplistBackendRequest20270101OneOf4463528992) GetCiTestQuery() LogQueryDefinition_20270101 {
	if o == nil {
		var ret LogQueryDefinition_20270101
		return ret
	}
	return o.CiTestQuery
}

// GetCiTestQueryOk returns a tuple with the CiTestQuery field value
// and a boolean to check if the value has been set.
func (o *ToplistBackendRequest20270101OneOf4463528992) GetCiTestQueryOk() (*LogQueryDefinition_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CiTestQuery, true
}

// SetCiTestQuery sets field value.
func (o *ToplistBackendRequest20270101OneOf4463528992) SetCiTestQuery(v LogQueryDefinition_20270101) {
	o.CiTestQuery = v
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *ToplistBackendRequest20270101OneOf4463528992) GetConditionalFormats() []ThresholdFormattingRule_20270101 {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule_20270101
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToplistBackendRequest20270101OneOf4463528992) GetConditionalFormatsOk() (*[]ThresholdFormattingRule_20270101, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *ToplistBackendRequest20270101OneOf4463528992) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule_20270101 and assigns it to the ConditionalFormats field.
func (o *ToplistBackendRequest20270101OneOf4463528992) SetConditionalFormats(v []ThresholdFormattingRule_20270101) {
	o.ConditionalFormats = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ToplistBackendRequest20270101OneOf4463528992) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["ci_test_query"] = o.CiTestQuery
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ToplistBackendRequest20270101OneOf4463528992) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CiTestQuery        *LogQueryDefinition_20270101       `json:"ci_test_query"`
		ConditionalFormats []ThresholdFormattingRule_20270101 `json:"conditional_formats,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CiTestQuery == nil {
		return fmt.Errorf("required field ci_test_query missing")
	}
	o.CiTestQuery = *all.CiTestQuery
	o.ConditionalFormats = all.ConditionalFormats

	return nil
}

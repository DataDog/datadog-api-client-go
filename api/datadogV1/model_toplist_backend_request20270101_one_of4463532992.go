// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistBackendRequest20270101OneOf4463532992 Test-only generated schema for the dashboard versioning proof of concept.
type ToplistBackendRequest20270101OneOf4463532992 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule_20270101 `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ProcessQuery ProcessQueryDefinition_20270101 `json:"process_query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewToplistBackendRequest20270101OneOf4463532992 instantiates a new ToplistBackendRequest20270101OneOf4463532992 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewToplistBackendRequest20270101OneOf4463532992(processQuery ProcessQueryDefinition_20270101) *ToplistBackendRequest20270101OneOf4463532992 {
	this := ToplistBackendRequest20270101OneOf4463532992{}
	this.ProcessQuery = processQuery
	return &this
}

// NewToplistBackendRequest20270101OneOf4463532992WithDefaults instantiates a new ToplistBackendRequest20270101OneOf4463532992 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewToplistBackendRequest20270101OneOf4463532992WithDefaults() *ToplistBackendRequest20270101OneOf4463532992 {
	this := ToplistBackendRequest20270101OneOf4463532992{}
	return &this
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *ToplistBackendRequest20270101OneOf4463532992) GetConditionalFormats() []ThresholdFormattingRule_20270101 {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule_20270101
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToplistBackendRequest20270101OneOf4463532992) GetConditionalFormatsOk() (*[]ThresholdFormattingRule_20270101, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *ToplistBackendRequest20270101OneOf4463532992) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule_20270101 and assigns it to the ConditionalFormats field.
func (o *ToplistBackendRequest20270101OneOf4463532992) SetConditionalFormats(v []ThresholdFormattingRule_20270101) {
	o.ConditionalFormats = v
}

// GetProcessQuery returns the ProcessQuery field value.
func (o *ToplistBackendRequest20270101OneOf4463532992) GetProcessQuery() ProcessQueryDefinition_20270101 {
	if o == nil {
		var ret ProcessQueryDefinition_20270101
		return ret
	}
	return o.ProcessQuery
}

// GetProcessQueryOk returns a tuple with the ProcessQuery field value
// and a boolean to check if the value has been set.
func (o *ToplistBackendRequest20270101OneOf4463532992) GetProcessQueryOk() (*ProcessQueryDefinition_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessQuery, true
}

// SetProcessQuery sets field value.
func (o *ToplistBackendRequest20270101OneOf4463532992) SetProcessQuery(v ProcessQueryDefinition_20270101) {
	o.ProcessQuery = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ToplistBackendRequest20270101OneOf4463532992) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	toSerialize["process_query"] = o.ProcessQuery
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ToplistBackendRequest20270101OneOf4463532992) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConditionalFormats []ThresholdFormattingRule_20270101 `json:"conditional_formats,omitempty"`
		ProcessQuery       *ProcessQueryDefinition_20270101   `json:"process_query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ProcessQuery == nil {
		return fmt.Errorf("required field process_query missing")
	}
	o.ConditionalFormats = all.ConditionalFormats
	o.ProcessQuery = *all.ProcessQuery

	return nil
}

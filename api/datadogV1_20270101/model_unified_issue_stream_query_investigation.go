// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UnifiedIssueStreamQueryInvestigation Test-only generated schema for the dashboard versioning proof of concept.
type UnifiedIssueStreamQueryInvestigation struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Completed *bool `json:"completed,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewUnifiedIssueStreamQueryInvestigation instantiates a new UnifiedIssueStreamQueryInvestigation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUnifiedIssueStreamQueryInvestigation() *UnifiedIssueStreamQueryInvestigation {
	this := UnifiedIssueStreamQueryInvestigation{}
	return &this
}

// NewUnifiedIssueStreamQueryInvestigationWithDefaults instantiates a new UnifiedIssueStreamQueryInvestigation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUnifiedIssueStreamQueryInvestigationWithDefaults() *UnifiedIssueStreamQueryInvestigation {
	this := UnifiedIssueStreamQueryInvestigation{}
	return &this
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *UnifiedIssueStreamQueryInvestigation) GetCompleted() bool {
	if o == nil || o.Completed == nil {
		var ret bool
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnifiedIssueStreamQueryInvestigation) GetCompletedOk() (*bool, bool) {
	if o == nil || o.Completed == nil {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *UnifiedIssueStreamQueryInvestigation) HasCompleted() bool {
	return o != nil && o.Completed != nil
}

// SetCompleted gets a reference to the given bool and assigns it to the Completed field.
func (o *UnifiedIssueStreamQueryInvestigation) SetCompleted(v bool) {
	o.Completed = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UnifiedIssueStreamQueryInvestigation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Completed != nil {
		toSerialize["completed"] = o.Completed
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UnifiedIssueStreamQueryInvestigation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Completed *bool `json:"completed,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.Completed = all.Completed

	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DistributionBackendRequest20270101OneOf4455771616 Test-only generated schema for the dashboard versioning proof of concept.
type DistributionBackendRequest20270101OneOf4455771616 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *RequestStyle_20270101 `json:"style,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SyntheticsTestRuns LogQueryDefinition_20270101 `json:"synthetics_test_runs"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDistributionBackendRequest20270101OneOf4455771616 instantiates a new DistributionBackendRequest20270101OneOf4455771616 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDistributionBackendRequest20270101OneOf4455771616(syntheticsTestRuns LogQueryDefinition_20270101) *DistributionBackendRequest20270101OneOf4455771616 {
	this := DistributionBackendRequest20270101OneOf4455771616{}
	this.SyntheticsTestRuns = syntheticsTestRuns
	return &this
}

// NewDistributionBackendRequest20270101OneOf4455771616WithDefaults instantiates a new DistributionBackendRequest20270101OneOf4455771616 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDistributionBackendRequest20270101OneOf4455771616WithDefaults() *DistributionBackendRequest20270101OneOf4455771616 {
	this := DistributionBackendRequest20270101OneOf4455771616{}
	return &this
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *DistributionBackendRequest20270101OneOf4455771616) GetStyle() RequestStyle_20270101 {
	if o == nil || o.Style == nil {
		var ret RequestStyle_20270101
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendRequest20270101OneOf4455771616) GetStyleOk() (*RequestStyle_20270101, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *DistributionBackendRequest20270101OneOf4455771616) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given RequestStyle_20270101 and assigns it to the Style field.
func (o *DistributionBackendRequest20270101OneOf4455771616) SetStyle(v RequestStyle_20270101) {
	o.Style = &v
}

// GetSyntheticsTestRuns returns the SyntheticsTestRuns field value.
func (o *DistributionBackendRequest20270101OneOf4455771616) GetSyntheticsTestRuns() LogQueryDefinition_20270101 {
	if o == nil {
		var ret LogQueryDefinition_20270101
		return ret
	}
	return o.SyntheticsTestRuns
}

// GetSyntheticsTestRunsOk returns a tuple with the SyntheticsTestRuns field value
// and a boolean to check if the value has been set.
func (o *DistributionBackendRequest20270101OneOf4455771616) GetSyntheticsTestRunsOk() (*LogQueryDefinition_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyntheticsTestRuns, true
}

// SetSyntheticsTestRuns sets field value.
func (o *DistributionBackendRequest20270101OneOf4455771616) SetSyntheticsTestRuns(v LogQueryDefinition_20270101) {
	o.SyntheticsTestRuns = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DistributionBackendRequest20270101OneOf4455771616) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	toSerialize["synthetics_test_runs"] = o.SyntheticsTestRuns
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DistributionBackendRequest20270101OneOf4455771616) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Style              *RequestStyle_20270101       `json:"style,omitempty"`
		SyntheticsTestRuns *LogQueryDefinition_20270101 `json:"synthetics_test_runs"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SyntheticsTestRuns == nil {
		return fmt.Errorf("required field synthetics_test_runs missing")
	}
	o.Style = all.Style
	o.SyntheticsTestRuns = *all.SyntheticsTestRuns

	return nil
}

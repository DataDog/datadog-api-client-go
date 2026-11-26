// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaRetentionQuerySampling Test-only generated schema for the dashboard versioning proof of concept.
type FormulaRetentionQuerySampling struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Enabled bool `json:"enabled"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaRetentionQuerySampling instantiates a new FormulaRetentionQuerySampling object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaRetentionQuerySampling(enabled bool) *FormulaRetentionQuerySampling {
	this := FormulaRetentionQuerySampling{}
	this.Enabled = enabled
	return &this
}

// NewFormulaRetentionQuerySamplingWithDefaults instantiates a new FormulaRetentionQuerySampling object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaRetentionQuerySamplingWithDefaults() *FormulaRetentionQuerySampling {
	this := FormulaRetentionQuerySampling{}
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *FormulaRetentionQuerySampling) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *FormulaRetentionQuerySampling) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *FormulaRetentionQuerySampling) SetEnabled(v bool) {
	o.Enabled = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaRetentionQuerySampling) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaRetentionQuerySampling) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled *bool `json:"enabled"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	o.Enabled = *all.Enabled

	return nil
}

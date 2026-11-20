// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SampledAnalysisQuerySample Test-only generated schema for the dashboard versioning proof of concept.
type SampledAnalysisQuerySample struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit *float64 `json:"limit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSampledAnalysisQuerySample instantiates a new SampledAnalysisQuerySample object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSampledAnalysisQuerySample() *SampledAnalysisQuerySample {
	this := SampledAnalysisQuerySample{}
	return &this
}

// NewSampledAnalysisQuerySampleWithDefaults instantiates a new SampledAnalysisQuerySample object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSampledAnalysisQuerySampleWithDefaults() *SampledAnalysisQuerySample {
	this := SampledAnalysisQuerySample{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *SampledAnalysisQuerySample) GetLimit() float64 {
	if o == nil || o.Limit == nil {
		var ret float64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SampledAnalysisQuerySample) GetLimitOk() (*float64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *SampledAnalysisQuerySample) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given float64 and assigns it to the Limit field.
func (o *SampledAnalysisQuerySample) SetLimit(v float64) {
	o.Limit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SampledAnalysisQuerySample) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SampledAnalysisQuerySample) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Limit *float64 `json:"limit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.Limit = all.Limit

	return nil
}

/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// MetricSearchResponse Object containing the list of metrics matching the search query.
type MetricSearchResponse struct {
	Results *MetricSearchResponseResults `json:"results,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewMetricSearchResponse instantiates a new MetricSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricSearchResponse() *MetricSearchResponse {
	this := MetricSearchResponse{}
	return &this
}

// NewMetricSearchResponseWithDefaults instantiates a new MetricSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricSearchResponseWithDefaults() *MetricSearchResponse {
	this := MetricSearchResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *MetricSearchResponse) GetResults() MetricSearchResponseResults {
	if o == nil || o.Results == nil {
		var ret MetricSearchResponseResults
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricSearchResponse) GetResultsOk() (*MetricSearchResponseResults, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *MetricSearchResponse) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given MetricSearchResponseResults and assigns it to the Results field.
func (o *MetricSearchResponse) SetResults(v MetricSearchResponseResults) {
	o.Results = &v
}

func (o MetricSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

func (o *MetricSearchResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Results *MetricSearchResponseResults `json:"results,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Results = all.Results
	return nil
}

type NullableMetricSearchResponse struct {
	value *MetricSearchResponse
	isSet bool
}

func (v NullableMetricSearchResponse) Get() *MetricSearchResponse {
	return v.value
}

func (v *NullableMetricSearchResponse) Set(val *MetricSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricSearchResponse(val *MetricSearchResponse) *NullableMetricSearchResponse {
	return &NullableMetricSearchResponse{value: val, isSet: true}
}

func (v NullableMetricSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

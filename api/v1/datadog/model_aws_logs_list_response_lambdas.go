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

// AWSLogsListResponseLambdas struct for AWSLogsListResponseLambdas
type AWSLogsListResponseLambdas struct {
	Arn *string `json:"arn,omitempty"`
}

// NewAWSLogsListResponseLambdas instantiates a new AWSLogsListResponseLambdas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSLogsListResponseLambdas() *AWSLogsListResponseLambdas {
	this := AWSLogsListResponseLambdas{}
	return &this
}

// NewAWSLogsListResponseLambdasWithDefaults instantiates a new AWSLogsListResponseLambdas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSLogsListResponseLambdasWithDefaults() *AWSLogsListResponseLambdas {
	this := AWSLogsListResponseLambdas{}
	return &this
}

// GetArn returns the Arn field value if set, zero value otherwise.
func (o *AWSLogsListResponseLambdas) GetArn() string {
	if o == nil || o.Arn == nil {
		var ret string
		return ret
	}
	return *o.Arn
}

// GetArnOk returns a tuple with the Arn field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AWSLogsListResponseLambdas) GetArnOk() (string, bool) {
	if o == nil || o.Arn == nil {
		var ret string
		return ret, false
	}
	return *o.Arn, true
}

// HasArn returns a boolean if a field has been set.
func (o *AWSLogsListResponseLambdas) HasArn() bool {
	if o != nil && o.Arn != nil {
		return true
	}

	return false
}

// SetArn gets a reference to the given string and assigns it to the Arn field.
func (o *AWSLogsListResponseLambdas) SetArn(v string) {
	o.Arn = &v
}

func (o AWSLogsListResponseLambdas) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Arn != nil {
		toSerialize["arn"] = o.Arn
	}
	return json.Marshal(toSerialize)
}

type NullableAWSLogsListResponseLambdas struct {
	value *AWSLogsListResponseLambdas
	isSet bool
}

func (v NullableAWSLogsListResponseLambdas) Get() *AWSLogsListResponseLambdas {
	return v.value
}

func (v NullableAWSLogsListResponseLambdas) Set(val *AWSLogsListResponseLambdas) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSLogsListResponseLambdas) IsSet() bool {
	return v.isSet
}

func (v NullableAWSLogsListResponseLambdas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSLogsListResponseLambdas(val *AWSLogsListResponseLambdas) *NullableAWSLogsListResponseLambdas {
	return &NullableAWSLogsListResponseLambdas{value: val, isSet: true}
}

func (v NullableAWSLogsListResponseLambdas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSLogsListResponseLambdas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// AWSAccountCreateResponse The Response returned by the AWS Create Account call.
type AWSAccountCreateResponse struct {
	// TODO.
	ExternalId *string `json:"external_id,omitempty"`
}

// NewAWSAccountCreateResponse instantiates a new AWSAccountCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSAccountCreateResponse() *AWSAccountCreateResponse {
	this := AWSAccountCreateResponse{}
	return &this
}

// NewAWSAccountCreateResponseWithDefaults instantiates a new AWSAccountCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSAccountCreateResponseWithDefaults() *AWSAccountCreateResponse {
	this := AWSAccountCreateResponse{}
	return &this
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *AWSAccountCreateResponse) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountCreateResponse) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *AWSAccountCreateResponse) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *AWSAccountCreateResponse) SetExternalId(v string) {
	o.ExternalId = &v
}

func (o AWSAccountCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalId != nil {
		toSerialize["external_id"] = o.ExternalId
	}
	return json.Marshal(toSerialize)
}

type NullableAWSAccountCreateResponse struct {
	value *AWSAccountCreateResponse
	isSet bool
}

func (v NullableAWSAccountCreateResponse) Get() *AWSAccountCreateResponse {
	return v.value
}

func (v *NullableAWSAccountCreateResponse) Set(val *AWSAccountCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSAccountCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSAccountCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSAccountCreateResponse(val *AWSAccountCreateResponse) *NullableAWSAccountCreateResponse {
	return &NullableAWSAccountCreateResponse{value: val, isSet: true}
}

func (v NullableAWSAccountCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSAccountCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// ApiKeyResponse struct for ApiKeyResponse
type ApiKeyResponse struct {
	ApiKey *ApiKey `json:"api_key,omitempty"`
}

// NewApiKeyResponse instantiates a new ApiKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyResponse() *ApiKeyResponse {
	this := ApiKeyResponse{}
	return &this
}

// NewApiKeyResponseWithDefaults instantiates a new ApiKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyResponseWithDefaults() *ApiKeyResponse {
	this := ApiKeyResponse{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *ApiKeyResponse) GetApiKey() ApiKey {
	if o == nil || o.ApiKey == nil {
		var ret ApiKey
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyResponse) GetApiKeyOk() (ApiKey, bool) {
	if o == nil || o.ApiKey == nil {
		var ret ApiKey
		return ret, false
	}
	return *o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *ApiKeyResponse) HasApiKey() bool {
	if o != nil && o.ApiKey != nil {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given ApiKey and assigns it to the ApiKey field.
func (o *ApiKeyResponse) SetApiKey(v ApiKey) {
	o.ApiKey = &v
}

type NullableApiKeyResponse struct {
	Value        ApiKeyResponse
	ExplicitNull bool
}

func (v NullableApiKeyResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableApiKeyResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

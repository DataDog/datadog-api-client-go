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

// CloudWorkloadSecurityAgentRuleResponse Response object that includes an Agent rule.
type CloudWorkloadSecurityAgentRuleResponse struct {
	Data *CloudWorkloadSecurityAgentRuleData `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewCloudWorkloadSecurityAgentRuleResponse instantiates a new CloudWorkloadSecurityAgentRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudWorkloadSecurityAgentRuleResponse() *CloudWorkloadSecurityAgentRuleResponse {
	this := CloudWorkloadSecurityAgentRuleResponse{}
	return &this
}

// NewCloudWorkloadSecurityAgentRuleResponseWithDefaults instantiates a new CloudWorkloadSecurityAgentRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudWorkloadSecurityAgentRuleResponseWithDefaults() *CloudWorkloadSecurityAgentRuleResponse {
	this := CloudWorkloadSecurityAgentRuleResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleResponse) GetData() CloudWorkloadSecurityAgentRuleData {
	if o == nil || o.Data == nil {
		var ret CloudWorkloadSecurityAgentRuleData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleResponse) GetDataOk() (*CloudWorkloadSecurityAgentRuleData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given CloudWorkloadSecurityAgentRuleData and assigns it to the Data field.
func (o *CloudWorkloadSecurityAgentRuleResponse) SetData(v CloudWorkloadSecurityAgentRuleData) {
	o.Data = &v
}

func (o CloudWorkloadSecurityAgentRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

func (o *CloudWorkloadSecurityAgentRuleResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Data *CloudWorkloadSecurityAgentRuleData `json:"data,omitempty"`
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
	o.Data = all.Data
	return nil
}

type NullableCloudWorkloadSecurityAgentRuleResponse struct {
	value *CloudWorkloadSecurityAgentRuleResponse
	isSet bool
}

func (v NullableCloudWorkloadSecurityAgentRuleResponse) Get() *CloudWorkloadSecurityAgentRuleResponse {
	return v.value
}

func (v *NullableCloudWorkloadSecurityAgentRuleResponse) Set(val *CloudWorkloadSecurityAgentRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudWorkloadSecurityAgentRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudWorkloadSecurityAgentRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudWorkloadSecurityAgentRuleResponse(val *CloudWorkloadSecurityAgentRuleResponse) *NullableCloudWorkloadSecurityAgentRuleResponse {
	return &NullableCloudWorkloadSecurityAgentRuleResponse{value: val, isSet: true}
}

func (v NullableCloudWorkloadSecurityAgentRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudWorkloadSecurityAgentRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

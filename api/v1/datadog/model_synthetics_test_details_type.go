/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// SyntheticsTestDetailsType Type of the Synthetic test, either `api` or `browser`.
type SyntheticsTestDetailsType string

// List of SyntheticsTestDetailsType
const (
	SYNTHETICSTESTDETAILSTYPE_API SyntheticsTestDetailsType = "api"
	SYNTHETICSTESTDETAILSTYPE_BROWSER SyntheticsTestDetailsType = "browser"
)

func (v *SyntheticsTestDetailsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SyntheticsTestDetailsType(value)
	for _, existing := range []SyntheticsTestDetailsType{ "api", "browser",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SyntheticsTestDetailsType", *v)
}

// Ptr returns reference to SyntheticsTestDetailsType value
func (v SyntheticsTestDetailsType) Ptr() *SyntheticsTestDetailsType {
	return &v
}

type NullableSyntheticsTestDetailsType struct {
	value *SyntheticsTestDetailsType
	isSet bool
}

func (v NullableSyntheticsTestDetailsType) Get() *SyntheticsTestDetailsType {
	return v.value
}

func (v *NullableSyntheticsTestDetailsType) Set(val *SyntheticsTestDetailsType) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsTestDetailsType) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsTestDetailsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsTestDetailsType(val *SyntheticsTestDetailsType) *NullableSyntheticsTestDetailsType {
	return &NullableSyntheticsTestDetailsType{value: val, isSet: true}
}

func (v NullableSyntheticsTestDetailsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsTestDetailsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


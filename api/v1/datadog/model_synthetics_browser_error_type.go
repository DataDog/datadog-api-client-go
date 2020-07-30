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

// SyntheticsBrowserErrorType Error type returned by a browser test.
type SyntheticsBrowserErrorType string

// List of SyntheticsBrowserErrorType
const (
	SYNTHETICSBROWSERERRORTYPE_NETWORK SyntheticsBrowserErrorType = "network"
	SYNTHETICSBROWSERERRORTYPE_JS      SyntheticsBrowserErrorType = "js"
)

func (v *SyntheticsBrowserErrorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SyntheticsBrowserErrorType(value)
	for _, existing := range []SyntheticsBrowserErrorType{"network", "js"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SyntheticsBrowserErrorType", value)
}

// Ptr returns reference to SyntheticsBrowserErrorType value
func (v SyntheticsBrowserErrorType) Ptr() *SyntheticsBrowserErrorType {
	return &v
}

type NullableSyntheticsBrowserErrorType struct {
	value *SyntheticsBrowserErrorType
	isSet bool
}

func (v NullableSyntheticsBrowserErrorType) Get() *SyntheticsBrowserErrorType {
	return v.value
}

func (v *NullableSyntheticsBrowserErrorType) Set(val *SyntheticsBrowserErrorType) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsBrowserErrorType) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsBrowserErrorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsBrowserErrorType(val *SyntheticsBrowserErrorType) *NullableSyntheticsBrowserErrorType {
	return &NullableSyntheticsBrowserErrorType{value: val, isSet: true}
}

func (v NullableSyntheticsBrowserErrorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsBrowserErrorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

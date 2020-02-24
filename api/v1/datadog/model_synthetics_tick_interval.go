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

// SyntheticsTickInterval the model 'SyntheticsTickInterval'
type SyntheticsTickInterval int64

// List of SyntheticsTickInterval
const (
	SYNTHETICSTICKINTERVAL_MINUTE          SyntheticsTickInterval = 60
	SYNTHETICSTICKINTERVAL_FIVE_MINUTES    SyntheticsTickInterval = 300
	SYNTHETICSTICKINTERVAL_FIFTEEN_MINUTES SyntheticsTickInterval = 900
	SYNTHETICSTICKINTERVAL_THIRTY_MINUTES  SyntheticsTickInterval = 1800
	SYNTHETICSTICKINTERVAL_HOUR            SyntheticsTickInterval = 3600
	SYNTHETICSTICKINTERVAL_SIX_HOURS       SyntheticsTickInterval = 21600
	SYNTHETICSTICKINTERVAL_TWELVE_HOURS    SyntheticsTickInterval = 43200
	SYNTHETICSTICKINTERVAL_DAY             SyntheticsTickInterval = 86400
	SYNTHETICSTICKINTERVAL_WEEK            SyntheticsTickInterval = 604800
)

// Ptr returns reference to SyntheticsTickInterval value
func (v SyntheticsTickInterval) Ptr() *SyntheticsTickInterval {
	return &v
}

type NullableSyntheticsTickInterval struct {
	value *SyntheticsTickInterval
	isSet bool
}

func (v NullableSyntheticsTickInterval) Get() *SyntheticsTickInterval {
	return v.value
}

func (v NullableSyntheticsTickInterval) Set(val *SyntheticsTickInterval) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsTickInterval) IsSet() bool {
	return v.isSet
}

func (v NullableSyntheticsTickInterval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsTickInterval(val *SyntheticsTickInterval) *NullableSyntheticsTickInterval {
	return &NullableSyntheticsTickInterval{value: val, isSet: true}
}

func (v NullableSyntheticsTickInterval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsTickInterval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

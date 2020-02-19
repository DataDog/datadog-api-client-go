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

// SyntheticsTestPauseStatus the model 'SyntheticsTestPauseStatus'
type SyntheticsTestPauseStatus string

// List of SyntheticsTestPauseStatus
const (
	SYNTHETICSTESTPAUSESTATUS_LIVE   SyntheticsTestPauseStatus = "live"
	SYNTHETICSTESTPAUSESTATUS_PAUSED SyntheticsTestPauseStatus = "paused"
)

// Ptr returns reference to SyntheticsTestPauseStatus value
func (v SyntheticsTestPauseStatus) Ptr() *SyntheticsTestPauseStatus {
	return &v
}

type NullableSyntheticsTestPauseStatus struct {
	value *SyntheticsTestPauseStatus
	isSet bool
}

func (v NullableSyntheticsTestPauseStatus) Get() *SyntheticsTestPauseStatus {
	return v.value
}

func (v NullableSyntheticsTestPauseStatus) Set(val *SyntheticsTestPauseStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsTestPauseStatus) IsSet() bool {
	return v.isSet
}

func (v NullableSyntheticsTestPauseStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsTestPauseStatus(val *SyntheticsTestPauseStatus) *NullableSyntheticsTestPauseStatus {
	return &NullableSyntheticsTestPauseStatus{value: val, isSet: true}
}

func (v NullableSyntheticsTestPauseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsTestPauseStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

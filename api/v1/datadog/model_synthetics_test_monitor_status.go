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

// SyntheticsTestMonitorStatus The status of your Synthetic monitor. * `O` for not triggered * `1` for triggered * `2` for no data
type SyntheticsTestMonitorStatus int64

// List of SyntheticsTestMonitorStatus
const (
	SYNTHETICSTESTMONITORSTATUS_UNTRIGGERED SyntheticsTestMonitorStatus = 0
	SYNTHETICSTESTMONITORSTATUS_TRIGGERED SyntheticsTestMonitorStatus = 1
	SYNTHETICSTESTMONITORSTATUS_NO_DATA SyntheticsTestMonitorStatus = 2
)

func (v *SyntheticsTestMonitorStatus) UnmarshalJSON(src []byte) error {
	var value int64
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SyntheticsTestMonitorStatus(value)
	for _, existing := range []SyntheticsTestMonitorStatus{ 0, 1, 2,   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SyntheticsTestMonitorStatus", *v)
}

// Ptr returns reference to SyntheticsTestMonitorStatus value
func (v SyntheticsTestMonitorStatus) Ptr() *SyntheticsTestMonitorStatus {
	return &v
}

type NullableSyntheticsTestMonitorStatus struct {
	value *SyntheticsTestMonitorStatus
	isSet bool
}

func (v NullableSyntheticsTestMonitorStatus) Get() *SyntheticsTestMonitorStatus {
	return v.value
}

func (v *NullableSyntheticsTestMonitorStatus) Set(val *SyntheticsTestMonitorStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsTestMonitorStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsTestMonitorStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsTestMonitorStatus(val *SyntheticsTestMonitorStatus) *NullableSyntheticsTestMonitorStatus {
	return &NullableSyntheticsTestMonitorStatus{value: val, isSet: true}
}

func (v NullableSyntheticsTestMonitorStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsTestMonitorStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


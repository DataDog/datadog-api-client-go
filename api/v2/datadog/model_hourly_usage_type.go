/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
	"fmt"
)

// HourlyUsageType Usage type that is being measured.
type HourlyUsageType string

// List of HourlyUsageType
const (
	HOURLYUSAGETYPE_OBSERVABILITY_PIPELINES_BYTES_PROCESSSED HourlyUsageType = "observability_pipelines_bytes_processed"
)

var allowedHourlyUsageTypeEnumValues = []HourlyUsageType{
	HOURLYUSAGETYPE_OBSERVABILITY_PIPELINES_BYTES_PROCESSSED,
}

func (w *HourlyUsageType) GetAllowedValues() []HourlyUsageType {
	return allowedHourlyUsageTypeEnumValues
}

func (v *HourlyUsageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HourlyUsageType(value)
	return nil
}

// NewHourlyUsageTypeFromValue returns a pointer to a valid HourlyUsageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHourlyUsageTypeFromValue(v string) (*HourlyUsageType, error) {
	ev := HourlyUsageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HourlyUsageType: valid values are %v", v, allowedHourlyUsageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HourlyUsageType) IsValid() bool {
	for _, existing := range allowedHourlyUsageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HourlyUsageType value
func (v HourlyUsageType) Ptr() *HourlyUsageType {
	return &v
}

type NullableHourlyUsageType struct {
	value *HourlyUsageType
	isSet bool
}

func (v NullableHourlyUsageType) Get() *HourlyUsageType {
	return v.value
}

func (v *NullableHourlyUsageType) Set(val *HourlyUsageType) {
	v.value = val
	v.isSet = true
}

func (v NullableHourlyUsageType) IsSet() bool {
	return v.isSet
}

func (v *NullableHourlyUsageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHourlyUsageType(val *HourlyUsageType) *NullableHourlyUsageType {
	return &NullableHourlyUsageType{value: val, isSet: true}
}

func (v NullableHourlyUsageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHourlyUsageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

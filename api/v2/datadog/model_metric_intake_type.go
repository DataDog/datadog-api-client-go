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

// MetricIntakeType The type of metric.
type MetricIntakeType int32

// List of MetricIntakeType
const (
	METRICINTAKETYPE_UNSPECIFIED        MetricIntakeType = 0
	METRICINTAKETYPE_COUNT              MetricIntakeType = 1
	METRICINTAKETYPE_RATE               MetricIntakeType = 2
	METRICINTAKETYPE_GAUGE              MetricIntakeType = 3
	METRICINTAKETYPE_UNSPECIFIED_LEGACY MetricIntakeType = 15
)

var allowedMetricIntakeTypeEnumValues = []MetricIntakeType{
	METRICINTAKETYPE_UNSPECIFIED,
	METRICINTAKETYPE_COUNT,
	METRICINTAKETYPE_RATE,
	METRICINTAKETYPE_GAUGE,
	METRICINTAKETYPE_UNSPECIFIED_LEGACY,
}

func (w *MetricIntakeType) GetAllowedValues() []MetricIntakeType {
	return allowedMetricIntakeTypeEnumValues
}

func (v *MetricIntakeType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricIntakeType(value)
	return nil
}

// NewMetricIntakeTypeFromValue returns a pointer to a valid MetricIntakeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMetricIntakeTypeFromValue(v int32) (*MetricIntakeType, error) {
	ev := MetricIntakeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MetricIntakeType: valid values are %v", v, allowedMetricIntakeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MetricIntakeType) IsValid() bool {
	for _, existing := range allowedMetricIntakeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricIntakeType value
func (v MetricIntakeType) Ptr() *MetricIntakeType {
	return &v
}

type NullableMetricIntakeType struct {
	value *MetricIntakeType
	isSet bool
}

func (v NullableMetricIntakeType) Get() *MetricIntakeType {
	return v.value
}

func (v *NullableMetricIntakeType) Set(val *MetricIntakeType) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricIntakeType) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricIntakeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricIntakeType(val *MetricIntakeType) *NullableMetricIntakeType {
	return &NullableMetricIntakeType{value: val, isSet: true}
}

func (v NullableMetricIntakeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricIntakeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

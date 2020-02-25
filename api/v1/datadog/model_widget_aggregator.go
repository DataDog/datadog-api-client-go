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

// WidgetAggregator Aggregator used for the request, available values are avg, last, max, min, or sum.
type WidgetAggregator string

// List of WidgetAggregator
const (
	WIDGETAGGREGATOR_AVERAGE WidgetAggregator = "avg"
	WIDGETAGGREGATOR_LAST    WidgetAggregator = "last"
	WIDGETAGGREGATOR_MAXIMUM WidgetAggregator = "max"
	WIDGETAGGREGATOR_MINIMUM WidgetAggregator = "min"
	WIDGETAGGREGATOR_SUM     WidgetAggregator = "sum"
)

// Ptr returns reference to WidgetAggregator value
func (v WidgetAggregator) Ptr() *WidgetAggregator {
	return &v
}

type NullableWidgetAggregator struct {
	value *WidgetAggregator
	isSet bool
}

func (v NullableWidgetAggregator) Get() *WidgetAggregator {
	return v.value
}

func (v NullableWidgetAggregator) Set(val *WidgetAggregator) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetAggregator) IsSet() bool {
	return v.isSet
}

func (v NullableWidgetAggregator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetAggregator(val *WidgetAggregator) *NullableWidgetAggregator {
	return &NullableWidgetAggregator{value: val, isSet: true}
}

func (v NullableWidgetAggregator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetAggregator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

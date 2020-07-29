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

// UsageMetricCategory Contains the metric category.
type UsageMetricCategory string

// List of UsageMetricCategory
const (
	USAGEMETRICCATEGORY_STANDARD UsageMetricCategory = "standard"
	USAGEMETRICCATEGORY_CUSTOM   UsageMetricCategory = "custom"
)

func (v *UsageMetricCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UsageMetricCategory(value)
	for _, existing := range []UsageMetricCategory{"standard", "custom"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UsageMetricCategory", value)
}

// Ptr returns reference to UsageMetricCategory value
func (v UsageMetricCategory) Ptr() *UsageMetricCategory {
	return &v
}

type NullableUsageMetricCategory struct {
	value *UsageMetricCategory
	isSet bool
}

func (v NullableUsageMetricCategory) Get() *UsageMetricCategory {
	return v.value
}

func (v *NullableUsageMetricCategory) Set(val *UsageMetricCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageMetricCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageMetricCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageMetricCategory(val *UsageMetricCategory) *NullableUsageMetricCategory {
	return &NullableUsageMetricCategory{value: val, isSet: true}
}

func (v NullableUsageMetricCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageMetricCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsLaunchType Launch resource type.
type ProductAnalyticsLaunchType string

// List of ProductAnalyticsLaunchType.
const (
	PRODUCTANALYTICSLAUNCHTYPE_LIFTOFF ProductAnalyticsLaunchType = "liftoff"
)

var allowedProductAnalyticsLaunchTypeEnumValues = []ProductAnalyticsLaunchType{
	PRODUCTANALYTICSLAUNCHTYPE_LIFTOFF,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsLaunchType) GetAllowedValues() []ProductAnalyticsLaunchType {
	return allowedProductAnalyticsLaunchTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsLaunchType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsLaunchType(value)
	return nil
}

// NewProductAnalyticsLaunchTypeFromValue returns a pointer to a valid ProductAnalyticsLaunchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsLaunchTypeFromValue(v string) (*ProductAnalyticsLaunchType, error) {
	ev := ProductAnalyticsLaunchType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsLaunchType: valid values are %v", v, allowedProductAnalyticsLaunchTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsLaunchType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsLaunchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsLaunchType value.
func (v ProductAnalyticsLaunchType) Ptr() *ProductAnalyticsLaunchType {
	return &v
}

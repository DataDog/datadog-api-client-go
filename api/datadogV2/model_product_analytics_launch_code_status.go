// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsLaunchCodeStatus Whether the inspected code sends the event, or attaches the property to the event.
type ProductAnalyticsLaunchCodeStatus string

// List of ProductAnalyticsLaunchCodeStatus.
const (
	PRODUCTANALYTICSLAUNCHCODESTATUS_FOUND     ProductAnalyticsLaunchCodeStatus = "found"
	PRODUCTANALYTICSLAUNCHCODESTATUS_NOT_FOUND ProductAnalyticsLaunchCodeStatus = "not_found"
)

var allowedProductAnalyticsLaunchCodeStatusEnumValues = []ProductAnalyticsLaunchCodeStatus{
	PRODUCTANALYTICSLAUNCHCODESTATUS_FOUND,
	PRODUCTANALYTICSLAUNCHCODESTATUS_NOT_FOUND,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsLaunchCodeStatus) GetAllowedValues() []ProductAnalyticsLaunchCodeStatus {
	return allowedProductAnalyticsLaunchCodeStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsLaunchCodeStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsLaunchCodeStatus(value)
	return nil
}

// NewProductAnalyticsLaunchCodeStatusFromValue returns a pointer to a valid ProductAnalyticsLaunchCodeStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsLaunchCodeStatusFromValue(v string) (*ProductAnalyticsLaunchCodeStatus, error) {
	ev := ProductAnalyticsLaunchCodeStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsLaunchCodeStatus: valid values are %v", v, allowedProductAnalyticsLaunchCodeStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsLaunchCodeStatus) IsValid() bool {
	for _, existing := range allowedProductAnalyticsLaunchCodeStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsLaunchCodeStatus value.
func (v ProductAnalyticsLaunchCodeStatus) Ptr() *ProductAnalyticsLaunchCodeStatus {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsMonitorListType Product Analytics monitor list resource type.
type ProductAnalyticsMonitorListType string

// List of ProductAnalyticsMonitorListType.
const (
	PRODUCTANALYTICSMONITORLISTTYPE_PA_MONITOR_LIST ProductAnalyticsMonitorListType = "pa_monitor_list"
)

var allowedProductAnalyticsMonitorListTypeEnumValues = []ProductAnalyticsMonitorListType{
	PRODUCTANALYTICSMONITORLISTTYPE_PA_MONITOR_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsMonitorListType) GetAllowedValues() []ProductAnalyticsMonitorListType {
	return allowedProductAnalyticsMonitorListTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsMonitorListType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsMonitorListType(value)
	return nil
}

// NewProductAnalyticsMonitorListTypeFromValue returns a pointer to a valid ProductAnalyticsMonitorListType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsMonitorListTypeFromValue(v string) (*ProductAnalyticsMonitorListType, error) {
	ev := ProductAnalyticsMonitorListType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsMonitorListType: valid values are %v", v, allowedProductAnalyticsMonitorListTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsMonitorListType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsMonitorListTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsMonitorListType value.
func (v ProductAnalyticsMonitorListType) Ptr() *ProductAnalyticsMonitorListType {
	return &v
}

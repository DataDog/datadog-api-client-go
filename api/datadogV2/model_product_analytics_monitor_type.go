// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsMonitorType Product Analytics monitor resource type.
type ProductAnalyticsMonitorType string

// List of ProductAnalyticsMonitorType.
const (
	PRODUCTANALYTICSMONITORTYPE_PA_MONITOR ProductAnalyticsMonitorType = "pa_monitor"
)

var allowedProductAnalyticsMonitorTypeEnumValues = []ProductAnalyticsMonitorType{
	PRODUCTANALYTICSMONITORTYPE_PA_MONITOR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsMonitorType) GetAllowedValues() []ProductAnalyticsMonitorType {
	return allowedProductAnalyticsMonitorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsMonitorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsMonitorType(value)
	return nil
}

// NewProductAnalyticsMonitorTypeFromValue returns a pointer to a valid ProductAnalyticsMonitorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsMonitorTypeFromValue(v string) (*ProductAnalyticsMonitorType, error) {
	ev := ProductAnalyticsMonitorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsMonitorType: valid values are %v", v, allowedProductAnalyticsMonitorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsMonitorType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsMonitorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsMonitorType value.
func (v ProductAnalyticsMonitorType) Ptr() *ProductAnalyticsMonitorType {
	return &v
}

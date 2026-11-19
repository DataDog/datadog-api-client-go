// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsMonitorRequestType Product Analytics monitor request resource type.
type ProductAnalyticsMonitorRequestType string

// List of ProductAnalyticsMonitorRequestType.
const (
	PRODUCTANALYTICSMONITORREQUESTTYPE_PA_MONITOR_REQUEST ProductAnalyticsMonitorRequestType = "pa_monitor_request"
)

var allowedProductAnalyticsMonitorRequestTypeEnumValues = []ProductAnalyticsMonitorRequestType{
	PRODUCTANALYTICSMONITORREQUESTTYPE_PA_MONITOR_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsMonitorRequestType) GetAllowedValues() []ProductAnalyticsMonitorRequestType {
	return allowedProductAnalyticsMonitorRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsMonitorRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsMonitorRequestType(value)
	return nil
}

// NewProductAnalyticsMonitorRequestTypeFromValue returns a pointer to a valid ProductAnalyticsMonitorRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsMonitorRequestTypeFromValue(v string) (*ProductAnalyticsMonitorRequestType, error) {
	ev := ProductAnalyticsMonitorRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsMonitorRequestType: valid values are %v", v, allowedProductAnalyticsMonitorRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsMonitorRequestType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsMonitorRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsMonitorRequestType value.
func (v ProductAnalyticsMonitorRequestType) Ptr() *ProductAnalyticsMonitorRequestType {
	return &v
}

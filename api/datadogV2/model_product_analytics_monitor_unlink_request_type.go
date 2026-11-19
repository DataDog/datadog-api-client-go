// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsMonitorUnlinkRequestType Product Analytics monitor unlink request resource type.
type ProductAnalyticsMonitorUnlinkRequestType string

// List of ProductAnalyticsMonitorUnlinkRequestType.
const (
	PRODUCTANALYTICSMONITORUNLINKREQUESTTYPE_PA_MONITOR_UNLINK_REQUEST ProductAnalyticsMonitorUnlinkRequestType = "pa_monitor_unlink_request"
)

var allowedProductAnalyticsMonitorUnlinkRequestTypeEnumValues = []ProductAnalyticsMonitorUnlinkRequestType{
	PRODUCTANALYTICSMONITORUNLINKREQUESTTYPE_PA_MONITOR_UNLINK_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsMonitorUnlinkRequestType) GetAllowedValues() []ProductAnalyticsMonitorUnlinkRequestType {
	return allowedProductAnalyticsMonitorUnlinkRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsMonitorUnlinkRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsMonitorUnlinkRequestType(value)
	return nil
}

// NewProductAnalyticsMonitorUnlinkRequestTypeFromValue returns a pointer to a valid ProductAnalyticsMonitorUnlinkRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsMonitorUnlinkRequestTypeFromValue(v string) (*ProductAnalyticsMonitorUnlinkRequestType, error) {
	ev := ProductAnalyticsMonitorUnlinkRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsMonitorUnlinkRequestType: valid values are %v", v, allowedProductAnalyticsMonitorUnlinkRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsMonitorUnlinkRequestType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsMonitorUnlinkRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsMonitorUnlinkRequestType value.
func (v ProductAnalyticsMonitorUnlinkRequestType) Ptr() *ProductAnalyticsMonitorUnlinkRequestType {
	return &v
}

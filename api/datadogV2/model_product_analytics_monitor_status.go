// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsMonitorStatus Current alert status of the monitor. `No Data` is returned when the monitor has not been evaluated yet.
type ProductAnalyticsMonitorStatus string

// List of ProductAnalyticsMonitorStatus.
const (
	PRODUCTANALYTICSMONITORSTATUS_OK      ProductAnalyticsMonitorStatus = "OK"
	PRODUCTANALYTICSMONITORSTATUS_ALERT   ProductAnalyticsMonitorStatus = "Alert"
	PRODUCTANALYTICSMONITORSTATUS_WARN    ProductAnalyticsMonitorStatus = "Warn"
	PRODUCTANALYTICSMONITORSTATUS_NO_DATA ProductAnalyticsMonitorStatus = "No Data"
)

var allowedProductAnalyticsMonitorStatusEnumValues = []ProductAnalyticsMonitorStatus{
	PRODUCTANALYTICSMONITORSTATUS_OK,
	PRODUCTANALYTICSMONITORSTATUS_ALERT,
	PRODUCTANALYTICSMONITORSTATUS_WARN,
	PRODUCTANALYTICSMONITORSTATUS_NO_DATA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsMonitorStatus) GetAllowedValues() []ProductAnalyticsMonitorStatus {
	return allowedProductAnalyticsMonitorStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsMonitorStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsMonitorStatus(value)
	return nil
}

// NewProductAnalyticsMonitorStatusFromValue returns a pointer to a valid ProductAnalyticsMonitorStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsMonitorStatusFromValue(v string) (*ProductAnalyticsMonitorStatus, error) {
	ev := ProductAnalyticsMonitorStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsMonitorStatus: valid values are %v", v, allowedProductAnalyticsMonitorStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsMonitorStatus) IsValid() bool {
	for _, existing := range allowedProductAnalyticsMonitorStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsMonitorStatus value.
func (v ProductAnalyticsMonitorStatus) Ptr() *ProductAnalyticsMonitorStatus {
	return &v
}

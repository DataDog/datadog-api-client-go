// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264 Test-only generated schema for the dashboard versioning proof of concept.
type BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264 string

// List of BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264.
const (
	BARCHARTBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4443923264_SCALAR BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264 = "scalar"
)

var allowedBarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264EnumValues = []BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264{
	BARCHARTBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4443923264_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264) GetAllowedValues() []BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264 {
	return allowedBarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264(value)
	return nil
}

// NewBarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264FromValue returns a pointer to a valid BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264FromValue(v string) (*BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264, error) {
	ev := BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264: valid values are %v", v, allowedBarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264) IsValid() bool {
	for _, existing := range allowedBarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264 value.
func (v BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264) Ptr() *BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443923264 {
	return &v
}

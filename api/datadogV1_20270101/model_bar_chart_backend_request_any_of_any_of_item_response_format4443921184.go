// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443921184 Test-only generated schema for the dashboard versioning proof of concept.
type BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443921184 string

// List of BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443921184.
const (
	BARCHARTBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4443921184_SCALAR BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443921184 = "scalar"
)

var allowedBarChartBackendRequestAnyOfAnyOfItemResponseFormat4443921184EnumValues = []BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443921184{
	BARCHARTBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4443921184_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443921184) GetAllowedValues() []BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443921184 {
	return allowedBarChartBackendRequestAnyOfAnyOfItemResponseFormat4443921184EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443921184) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443921184(value)
	return nil
}

// NewBarChartBackendRequestAnyOfAnyOfItemResponseFormat4443921184FromValue returns a pointer to a valid BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443921184
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBarChartBackendRequestAnyOfAnyOfItemResponseFormat4443921184FromValue(v string) (*BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443921184, error) {
	ev := BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443921184(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443921184: valid values are %v", v, allowedBarChartBackendRequestAnyOfAnyOfItemResponseFormat4443921184EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443921184) IsValid() bool {
	for _, existing := range allowedBarChartBackendRequestAnyOfAnyOfItemResponseFormat4443921184EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443921184 value.
func (v BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443921184) Ptr() *BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443921184 {
	return &v
}

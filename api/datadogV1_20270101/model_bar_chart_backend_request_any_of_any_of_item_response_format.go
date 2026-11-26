// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BarChartBackendRequestAnyOfAnyOfItemResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type BarChartBackendRequestAnyOfAnyOfItemResponseFormat string

// List of BarChartBackendRequestAnyOfAnyOfItemResponseFormat.
const (
	BARCHARTBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT_TABULAR BarChartBackendRequestAnyOfAnyOfItemResponseFormat = "tabular"
)

var allowedBarChartBackendRequestAnyOfAnyOfItemResponseFormatEnumValues = []BarChartBackendRequestAnyOfAnyOfItemResponseFormat{
	BARCHARTBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT_TABULAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BarChartBackendRequestAnyOfAnyOfItemResponseFormat) GetAllowedValues() []BarChartBackendRequestAnyOfAnyOfItemResponseFormat {
	return allowedBarChartBackendRequestAnyOfAnyOfItemResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BarChartBackendRequestAnyOfAnyOfItemResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BarChartBackendRequestAnyOfAnyOfItemResponseFormat(value)
	return nil
}

// NewBarChartBackendRequestAnyOfAnyOfItemResponseFormatFromValue returns a pointer to a valid BarChartBackendRequestAnyOfAnyOfItemResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBarChartBackendRequestAnyOfAnyOfItemResponseFormatFromValue(v string) (*BarChartBackendRequestAnyOfAnyOfItemResponseFormat, error) {
	ev := BarChartBackendRequestAnyOfAnyOfItemResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BarChartBackendRequestAnyOfAnyOfItemResponseFormat: valid values are %v", v, allowedBarChartBackendRequestAnyOfAnyOfItemResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BarChartBackendRequestAnyOfAnyOfItemResponseFormat) IsValid() bool {
	for _, existing := range allowedBarChartBackendRequestAnyOfAnyOfItemResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BarChartBackendRequestAnyOfAnyOfItemResponseFormat value.
func (v BarChartBackendRequestAnyOfAnyOfItemResponseFormat) Ptr() *BarChartBackendRequestAnyOfAnyOfItemResponseFormat {
	return &v
}

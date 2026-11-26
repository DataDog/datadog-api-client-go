// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatmapBackendRequestAnyOfAnyOfItemResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type HeatmapBackendRequestAnyOfAnyOfItemResponseFormat string

// List of HeatmapBackendRequestAnyOfAnyOfItemResponseFormat.
const (
	HEATMAPBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT_TIMESERIES HeatmapBackendRequestAnyOfAnyOfItemResponseFormat = "timeseries"
)

var allowedHeatmapBackendRequestAnyOfAnyOfItemResponseFormatEnumValues = []HeatmapBackendRequestAnyOfAnyOfItemResponseFormat{
	HEATMAPBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT_TIMESERIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatmapBackendRequestAnyOfAnyOfItemResponseFormat) GetAllowedValues() []HeatmapBackendRequestAnyOfAnyOfItemResponseFormat {
	return allowedHeatmapBackendRequestAnyOfAnyOfItemResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatmapBackendRequestAnyOfAnyOfItemResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatmapBackendRequestAnyOfAnyOfItemResponseFormat(value)
	return nil
}

// NewHeatmapBackendRequestAnyOfAnyOfItemResponseFormatFromValue returns a pointer to a valid HeatmapBackendRequestAnyOfAnyOfItemResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatmapBackendRequestAnyOfAnyOfItemResponseFormatFromValue(v string) (*HeatmapBackendRequestAnyOfAnyOfItemResponseFormat, error) {
	ev := HeatmapBackendRequestAnyOfAnyOfItemResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatmapBackendRequestAnyOfAnyOfItemResponseFormat: valid values are %v", v, allowedHeatmapBackendRequestAnyOfAnyOfItemResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatmapBackendRequestAnyOfAnyOfItemResponseFormat) IsValid() bool {
	for _, existing := range allowedHeatmapBackendRequestAnyOfAnyOfItemResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatmapBackendRequestAnyOfAnyOfItemResponseFormat value.
func (v HeatmapBackendRequestAnyOfAnyOfItemResponseFormat) Ptr() *HeatmapBackendRequestAnyOfAnyOfItemResponseFormat {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridTileDefRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridTileDefRequestResponseFormat string

// List of HeatgridTileDefRequestResponseFormat.
const (
	HEATGRIDTILEDEFREQUESTRESPONSEFORMAT_TIMESERIES HeatgridTileDefRequestResponseFormat = "timeseries"
)

var allowedHeatgridTileDefRequestResponseFormatEnumValues = []HeatgridTileDefRequestResponseFormat{
	HEATGRIDTILEDEFREQUESTRESPONSEFORMAT_TIMESERIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridTileDefRequestResponseFormat) GetAllowedValues() []HeatgridTileDefRequestResponseFormat {
	return allowedHeatgridTileDefRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridTileDefRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridTileDefRequestResponseFormat(value)
	return nil
}

// NewHeatgridTileDefRequestResponseFormatFromValue returns a pointer to a valid HeatgridTileDefRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridTileDefRequestResponseFormatFromValue(v string) (*HeatgridTileDefRequestResponseFormat, error) {
	ev := HeatgridTileDefRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridTileDefRequestResponseFormat: valid values are %v", v, allowedHeatgridTileDefRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridTileDefRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedHeatgridTileDefRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridTileDefRequestResponseFormat value.
func (v HeatgridTileDefRequestResponseFormat) Ptr() *HeatgridTileDefRequestResponseFormat {
	return &v
}

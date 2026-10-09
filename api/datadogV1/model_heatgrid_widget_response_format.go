// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridWidgetResponseFormat Response format for heatgrid queries.
type HeatgridWidgetResponseFormat string

// List of HeatgridWidgetResponseFormat.
const (
	HEATGRIDWIDGETRESPONSEFORMAT_TIMESERIES HeatgridWidgetResponseFormat = "timeseries"
)

var allowedHeatgridWidgetResponseFormatEnumValues = []HeatgridWidgetResponseFormat{
	HEATGRIDWIDGETRESPONSEFORMAT_TIMESERIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridWidgetResponseFormat) GetAllowedValues() []HeatgridWidgetResponseFormat {
	return allowedHeatgridWidgetResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridWidgetResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridWidgetResponseFormat(value)
	return nil
}

// NewHeatgridWidgetResponseFormatFromValue returns a pointer to a valid HeatgridWidgetResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridWidgetResponseFormatFromValue(v string) (*HeatgridWidgetResponseFormat, error) {
	ev := HeatgridWidgetResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridWidgetResponseFormat: valid values are %v", v, allowedHeatgridWidgetResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridWidgetResponseFormat) IsValid() bool {
	for _, existing := range allowedHeatgridWidgetResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridWidgetResponseFormat value.
func (v HeatgridWidgetResponseFormat) Ptr() *HeatgridWidgetResponseFormat {
	return &v
}

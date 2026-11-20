// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableTrendCellDisplayModeOptionsYScale Test-only generated schema for the dashboard versioning proof of concept.
type TableTrendCellDisplayModeOptionsYScale string

// List of TableTrendCellDisplayModeOptionsYScale.
const (
	TABLETRENDCELLDISPLAYMODEOPTIONSYSCALE_SHARED      TableTrendCellDisplayModeOptionsYScale = "shared"
	TABLETRENDCELLDISPLAYMODEOPTIONSYSCALE_INDEPENDENT TableTrendCellDisplayModeOptionsYScale = "independent"
)

var allowedTableTrendCellDisplayModeOptionsYScaleEnumValues = []TableTrendCellDisplayModeOptionsYScale{
	TABLETRENDCELLDISPLAYMODEOPTIONSYSCALE_SHARED,
	TABLETRENDCELLDISPLAYMODEOPTIONSYSCALE_INDEPENDENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TableTrendCellDisplayModeOptionsYScale) GetAllowedValues() []TableTrendCellDisplayModeOptionsYScale {
	return allowedTableTrendCellDisplayModeOptionsYScaleEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TableTrendCellDisplayModeOptionsYScale) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TableTrendCellDisplayModeOptionsYScale(value)
	return nil
}

// NewTableTrendCellDisplayModeOptionsYScaleFromValue returns a pointer to a valid TableTrendCellDisplayModeOptionsYScale
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTableTrendCellDisplayModeOptionsYScaleFromValue(v string) (*TableTrendCellDisplayModeOptionsYScale, error) {
	ev := TableTrendCellDisplayModeOptionsYScale(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TableTrendCellDisplayModeOptionsYScale: valid values are %v", v, allowedTableTrendCellDisplayModeOptionsYScaleEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TableTrendCellDisplayModeOptionsYScale) IsValid() bool {
	for _, existing := range allowedTableTrendCellDisplayModeOptionsYScaleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TableTrendCellDisplayModeOptionsYScale value.
func (v TableTrendCellDisplayModeOptionsYScale) Ptr() *TableTrendCellDisplayModeOptionsYScale {
	return &v
}

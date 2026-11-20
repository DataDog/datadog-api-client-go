// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CellDisplayModeOption Test-only generated schema for the dashboard versioning proof of concept.
type CellDisplayModeOption string

// List of CellDisplayModeOption.
const (
	CELLDISPLAYMODEOPTION_NUMBER CellDisplayModeOption = "number"
	CELLDISPLAYMODEOPTION_BAR    CellDisplayModeOption = "bar"
	CELLDISPLAYMODEOPTION_TREND  CellDisplayModeOption = "trend"
)

var allowedCellDisplayModeOptionEnumValues = []CellDisplayModeOption{
	CELLDISPLAYMODEOPTION_NUMBER,
	CELLDISPLAYMODEOPTION_BAR,
	CELLDISPLAYMODEOPTION_TREND,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CellDisplayModeOption) GetAllowedValues() []CellDisplayModeOption {
	return allowedCellDisplayModeOptionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CellDisplayModeOption) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CellDisplayModeOption(value)
	return nil
}

// NewCellDisplayModeOptionFromValue returns a pointer to a valid CellDisplayModeOption
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCellDisplayModeOptionFromValue(v string) (*CellDisplayModeOption, error) {
	ev := CellDisplayModeOption(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CellDisplayModeOption: valid values are %v", v, allowedCellDisplayModeOptionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CellDisplayModeOption) IsValid() bool {
	for _, existing := range allowedCellDisplayModeOptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CellDisplayModeOption value.
func (v CellDisplayModeOption) Ptr() *CellDisplayModeOption {
	return &v
}

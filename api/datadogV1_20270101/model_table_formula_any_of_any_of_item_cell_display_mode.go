// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableFormulaAnyOfAnyOfItemCellDisplayMode Test-only generated schema for the dashboard versioning proof of concept.
type TableFormulaAnyOfAnyOfItemCellDisplayMode string

// List of TableFormulaAnyOfAnyOfItemCellDisplayMode.
const (
	TABLEFORMULAANYOFANYOFITEMCELLDISPLAYMODE_NUMBER TableFormulaAnyOfAnyOfItemCellDisplayMode = "number"
	TABLEFORMULAANYOFANYOFITEMCELLDISPLAYMODE_BAR    TableFormulaAnyOfAnyOfItemCellDisplayMode = "bar"
)

var allowedTableFormulaAnyOfAnyOfItemCellDisplayModeEnumValues = []TableFormulaAnyOfAnyOfItemCellDisplayMode{
	TABLEFORMULAANYOFANYOFITEMCELLDISPLAYMODE_NUMBER,
	TABLEFORMULAANYOFANYOFITEMCELLDISPLAYMODE_BAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TableFormulaAnyOfAnyOfItemCellDisplayMode) GetAllowedValues() []TableFormulaAnyOfAnyOfItemCellDisplayMode {
	return allowedTableFormulaAnyOfAnyOfItemCellDisplayModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TableFormulaAnyOfAnyOfItemCellDisplayMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TableFormulaAnyOfAnyOfItemCellDisplayMode(value)
	return nil
}

// NewTableFormulaAnyOfAnyOfItemCellDisplayModeFromValue returns a pointer to a valid TableFormulaAnyOfAnyOfItemCellDisplayMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTableFormulaAnyOfAnyOfItemCellDisplayModeFromValue(v string) (*TableFormulaAnyOfAnyOfItemCellDisplayMode, error) {
	ev := TableFormulaAnyOfAnyOfItemCellDisplayMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TableFormulaAnyOfAnyOfItemCellDisplayMode: valid values are %v", v, allowedTableFormulaAnyOfAnyOfItemCellDisplayModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TableFormulaAnyOfAnyOfItemCellDisplayMode) IsValid() bool {
	for _, existing := range allowedTableFormulaAnyOfAnyOfItemCellDisplayModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TableFormulaAnyOfAnyOfItemCellDisplayMode value.
func (v TableFormulaAnyOfAnyOfItemCellDisplayMode) Ptr() *TableFormulaAnyOfAnyOfItemCellDisplayMode {
	return &v
}

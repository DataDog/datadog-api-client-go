// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableFormulaAnyOfAnyOfItemCellDisplayMode4452532448 Test-only generated schema for the dashboard versioning proof of concept.
type TableFormulaAnyOfAnyOfItemCellDisplayMode4452532448 string

// List of TableFormulaAnyOfAnyOfItemCellDisplayMode4452532448.
const (
	TABLEFORMULAANYOFANYOFITEMCELLDISPLAYMODE4452532448_TREND TableFormulaAnyOfAnyOfItemCellDisplayMode4452532448 = "trend"
)

var allowedTableFormulaAnyOfAnyOfItemCellDisplayMode4452532448EnumValues = []TableFormulaAnyOfAnyOfItemCellDisplayMode4452532448{
	TABLEFORMULAANYOFANYOFITEMCELLDISPLAYMODE4452532448_TREND,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TableFormulaAnyOfAnyOfItemCellDisplayMode4452532448) GetAllowedValues() []TableFormulaAnyOfAnyOfItemCellDisplayMode4452532448 {
	return allowedTableFormulaAnyOfAnyOfItemCellDisplayMode4452532448EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TableFormulaAnyOfAnyOfItemCellDisplayMode4452532448) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TableFormulaAnyOfAnyOfItemCellDisplayMode4452532448(value)
	return nil
}

// NewTableFormulaAnyOfAnyOfItemCellDisplayMode4452532448FromValue returns a pointer to a valid TableFormulaAnyOfAnyOfItemCellDisplayMode4452532448
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTableFormulaAnyOfAnyOfItemCellDisplayMode4452532448FromValue(v string) (*TableFormulaAnyOfAnyOfItemCellDisplayMode4452532448, error) {
	ev := TableFormulaAnyOfAnyOfItemCellDisplayMode4452532448(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TableFormulaAnyOfAnyOfItemCellDisplayMode4452532448: valid values are %v", v, allowedTableFormulaAnyOfAnyOfItemCellDisplayMode4452532448EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TableFormulaAnyOfAnyOfItemCellDisplayMode4452532448) IsValid() bool {
	for _, existing := range allowedTableFormulaAnyOfAnyOfItemCellDisplayMode4452532448EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TableFormulaAnyOfAnyOfItemCellDisplayMode4452532448 value.
func (v TableFormulaAnyOfAnyOfItemCellDisplayMode4452532448) Ptr() *TableFormulaAnyOfAnyOfItemCellDisplayMode4452532448 {
	return &v
}

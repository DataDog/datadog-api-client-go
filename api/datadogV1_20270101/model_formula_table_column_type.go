// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaTableColumnType Test-only generated schema for the dashboard versioning proof of concept.
type FormulaTableColumnType string

// List of FormulaTableColumnType.
const (
	FORMULATABLECOLUMNTYPE_FORMULA FormulaTableColumnType = "formula"
)

var allowedFormulaTableColumnTypeEnumValues = []FormulaTableColumnType{
	FORMULATABLECOLUMNTYPE_FORMULA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaTableColumnType) GetAllowedValues() []FormulaTableColumnType {
	return allowedFormulaTableColumnTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaTableColumnType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaTableColumnType(value)
	return nil
}

// NewFormulaTableColumnTypeFromValue returns a pointer to a valid FormulaTableColumnType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaTableColumnTypeFromValue(v string) (*FormulaTableColumnType, error) {
	ev := FormulaTableColumnType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaTableColumnType: valid values are %v", v, allowedFormulaTableColumnTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaTableColumnType) IsValid() bool {
	for _, existing := range allowedFormulaTableColumnTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaTableColumnType value.
func (v FormulaTableColumnType) Ptr() *FormulaTableColumnType {
	return &v
}

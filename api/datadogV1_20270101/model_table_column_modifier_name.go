// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableColumnModifierName Test-only generated schema for the dashboard versioning proof of concept.
type TableColumnModifierName string

// List of TableColumnModifierName.
const (
	TABLECOLUMNMODIFIERNAME_FILL     TableColumnModifierName = "fill"
	TABLECOLUMNMODIFIERNAME_AS_COUNT TableColumnModifierName = "as_count"
	TABLECOLUMNMODIFIERNAME_AS_RATE  TableColumnModifierName = "as_rate"
	TABLECOLUMNMODIFIERNAME_WEIGHTED TableColumnModifierName = "weighted"
	TABLECOLUMNMODIFIERNAME_ROLLUP   TableColumnModifierName = "rollup"
)

var allowedTableColumnModifierNameEnumValues = []TableColumnModifierName{
	TABLECOLUMNMODIFIERNAME_FILL,
	TABLECOLUMNMODIFIERNAME_AS_COUNT,
	TABLECOLUMNMODIFIERNAME_AS_RATE,
	TABLECOLUMNMODIFIERNAME_WEIGHTED,
	TABLECOLUMNMODIFIERNAME_ROLLUP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TableColumnModifierName) GetAllowedValues() []TableColumnModifierName {
	return allowedTableColumnModifierNameEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TableColumnModifierName) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TableColumnModifierName(value)
	return nil
}

// NewTableColumnModifierNameFromValue returns a pointer to a valid TableColumnModifierName
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTableColumnModifierNameFromValue(v string) (*TableColumnModifierName, error) {
	ev := TableColumnModifierName(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TableColumnModifierName: valid values are %v", v, allowedTableColumnModifierNameEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TableColumnModifierName) IsValid() bool {
	for _, existing := range allowedTableColumnModifierNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TableColumnModifierName value.
func (v TableColumnModifierName) Ptr() *TableColumnModifierName {
	return &v
}

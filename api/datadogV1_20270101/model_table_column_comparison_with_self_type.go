// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableColumnComparisonWithSelfType Test-only generated schema for the dashboard versioning proof of concept.
type TableColumnComparisonWithSelfType string

// List of TableColumnComparisonWithSelfType.
const (
	TABLECOLUMNCOMPARISONWITHSELFTYPE_PERCENT_COLUMN_TOTAL TableColumnComparisonWithSelfType = "percent_column_total"
)

var allowedTableColumnComparisonWithSelfTypeEnumValues = []TableColumnComparisonWithSelfType{
	TABLECOLUMNCOMPARISONWITHSELFTYPE_PERCENT_COLUMN_TOTAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TableColumnComparisonWithSelfType) GetAllowedValues() []TableColumnComparisonWithSelfType {
	return allowedTableColumnComparisonWithSelfTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TableColumnComparisonWithSelfType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TableColumnComparisonWithSelfType(value)
	return nil
}

// NewTableColumnComparisonWithSelfTypeFromValue returns a pointer to a valid TableColumnComparisonWithSelfType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTableColumnComparisonWithSelfTypeFromValue(v string) (*TableColumnComparisonWithSelfType, error) {
	ev := TableColumnComparisonWithSelfType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TableColumnComparisonWithSelfType: valid values are %v", v, allowedTableColumnComparisonWithSelfTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TableColumnComparisonWithSelfType) IsValid() bool {
	for _, existing := range allowedTableColumnComparisonWithSelfTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TableColumnComparisonWithSelfType value.
func (v TableColumnComparisonWithSelfType) Ptr() *TableColumnComparisonWithSelfType {
	return &v
}

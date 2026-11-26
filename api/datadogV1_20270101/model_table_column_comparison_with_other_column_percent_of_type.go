// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableColumnComparisonWithOtherColumnPercentOfType Test-only generated schema for the dashboard versioning proof of concept.
type TableColumnComparisonWithOtherColumnPercentOfType string

// List of TableColumnComparisonWithOtherColumnPercentOfType.
const (
	TABLECOLUMNCOMPARISONWITHOTHERCOLUMNPERCENTOFTYPE_PERCENT_OF TableColumnComparisonWithOtherColumnPercentOfType = "percent_of"
)

var allowedTableColumnComparisonWithOtherColumnPercentOfTypeEnumValues = []TableColumnComparisonWithOtherColumnPercentOfType{
	TABLECOLUMNCOMPARISONWITHOTHERCOLUMNPERCENTOFTYPE_PERCENT_OF,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TableColumnComparisonWithOtherColumnPercentOfType) GetAllowedValues() []TableColumnComparisonWithOtherColumnPercentOfType {
	return allowedTableColumnComparisonWithOtherColumnPercentOfTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TableColumnComparisonWithOtherColumnPercentOfType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TableColumnComparisonWithOtherColumnPercentOfType(value)
	return nil
}

// NewTableColumnComparisonWithOtherColumnPercentOfTypeFromValue returns a pointer to a valid TableColumnComparisonWithOtherColumnPercentOfType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTableColumnComparisonWithOtherColumnPercentOfTypeFromValue(v string) (*TableColumnComparisonWithOtherColumnPercentOfType, error) {
	ev := TableColumnComparisonWithOtherColumnPercentOfType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TableColumnComparisonWithOtherColumnPercentOfType: valid values are %v", v, allowedTableColumnComparisonWithOtherColumnPercentOfTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TableColumnComparisonWithOtherColumnPercentOfType) IsValid() bool {
	for _, existing := range allowedTableColumnComparisonWithOtherColumnPercentOfTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TableColumnComparisonWithOtherColumnPercentOfType value.
func (v TableColumnComparisonWithOtherColumnPercentOfType) Ptr() *TableColumnComparisonWithOtherColumnPercentOfType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableColumnComparisonWithOtherColumnDiffFromType Test-only generated schema for the dashboard versioning proof of concept.
type TableColumnComparisonWithOtherColumnDiffFromType string

// List of TableColumnComparisonWithOtherColumnDiffFromType.
const (
	TABLECOLUMNCOMPARISONWITHOTHERCOLUMNDIFFFROMTYPE_DIFF_FROM TableColumnComparisonWithOtherColumnDiffFromType = "diff_from"
)

var allowedTableColumnComparisonWithOtherColumnDiffFromTypeEnumValues = []TableColumnComparisonWithOtherColumnDiffFromType{
	TABLECOLUMNCOMPARISONWITHOTHERCOLUMNDIFFFROMTYPE_DIFF_FROM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TableColumnComparisonWithOtherColumnDiffFromType) GetAllowedValues() []TableColumnComparisonWithOtherColumnDiffFromType {
	return allowedTableColumnComparisonWithOtherColumnDiffFromTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TableColumnComparisonWithOtherColumnDiffFromType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TableColumnComparisonWithOtherColumnDiffFromType(value)
	return nil
}

// NewTableColumnComparisonWithOtherColumnDiffFromTypeFromValue returns a pointer to a valid TableColumnComparisonWithOtherColumnDiffFromType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTableColumnComparisonWithOtherColumnDiffFromTypeFromValue(v string) (*TableColumnComparisonWithOtherColumnDiffFromType, error) {
	ev := TableColumnComparisonWithOtherColumnDiffFromType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TableColumnComparisonWithOtherColumnDiffFromType: valid values are %v", v, allowedTableColumnComparisonWithOtherColumnDiffFromTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TableColumnComparisonWithOtherColumnDiffFromType) IsValid() bool {
	for _, existing := range allowedTableColumnComparisonWithOtherColumnDiffFromTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TableColumnComparisonWithOtherColumnDiffFromType value.
func (v TableColumnComparisonWithOtherColumnDiffFromType) Ptr() *TableColumnComparisonWithOtherColumnDiffFromType {
	return &v
}

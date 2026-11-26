// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableColumnComparisonWithOtherColumnPercentDiffFromType Test-only generated schema for the dashboard versioning proof of concept.
type TableColumnComparisonWithOtherColumnPercentDiffFromType string

// List of TableColumnComparisonWithOtherColumnPercentDiffFromType.
const (
	TABLECOLUMNCOMPARISONWITHOTHERCOLUMNPERCENTDIFFFROMTYPE_PERCENT_DIFF_FROM TableColumnComparisonWithOtherColumnPercentDiffFromType = "percent_diff_from"
)

var allowedTableColumnComparisonWithOtherColumnPercentDiffFromTypeEnumValues = []TableColumnComparisonWithOtherColumnPercentDiffFromType{
	TABLECOLUMNCOMPARISONWITHOTHERCOLUMNPERCENTDIFFFROMTYPE_PERCENT_DIFF_FROM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TableColumnComparisonWithOtherColumnPercentDiffFromType) GetAllowedValues() []TableColumnComparisonWithOtherColumnPercentDiffFromType {
	return allowedTableColumnComparisonWithOtherColumnPercentDiffFromTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TableColumnComparisonWithOtherColumnPercentDiffFromType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TableColumnComparisonWithOtherColumnPercentDiffFromType(value)
	return nil
}

// NewTableColumnComparisonWithOtherColumnPercentDiffFromTypeFromValue returns a pointer to a valid TableColumnComparisonWithOtherColumnPercentDiffFromType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTableColumnComparisonWithOtherColumnPercentDiffFromTypeFromValue(v string) (*TableColumnComparisonWithOtherColumnPercentDiffFromType, error) {
	ev := TableColumnComparisonWithOtherColumnPercentDiffFromType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TableColumnComparisonWithOtherColumnPercentDiffFromType: valid values are %v", v, allowedTableColumnComparisonWithOtherColumnPercentDiffFromTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TableColumnComparisonWithOtherColumnPercentDiffFromType) IsValid() bool {
	for _, existing := range allowedTableColumnComparisonWithOtherColumnPercentDiffFromTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TableColumnComparisonWithOtherColumnPercentDiffFromType value.
func (v TableColumnComparisonWithOtherColumnPercentDiffFromType) Ptr() *TableColumnComparisonWithOtherColumnPercentDiffFromType {
	return &v
}

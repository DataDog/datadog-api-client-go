// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableRowsJoinType Test-only generated schema for the dashboard versioning proof of concept.
type TableRowsJoinType string

// List of TableRowsJoinType.
const (
	TABLEROWSJOINTYPE_LEFT  TableRowsJoinType = "left"
	TABLEROWSJOINTYPE_INNER TableRowsJoinType = "inner"
)

var allowedTableRowsJoinTypeEnumValues = []TableRowsJoinType{
	TABLEROWSJOINTYPE_LEFT,
	TABLEROWSJOINTYPE_INNER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TableRowsJoinType) GetAllowedValues() []TableRowsJoinType {
	return allowedTableRowsJoinTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TableRowsJoinType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TableRowsJoinType(value)
	return nil
}

// NewTableRowsJoinTypeFromValue returns a pointer to a valid TableRowsJoinType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTableRowsJoinTypeFromValue(v string) (*TableRowsJoinType, error) {
	ev := TableRowsJoinType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TableRowsJoinType: valid values are %v", v, allowedTableRowsJoinTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TableRowsJoinType) IsValid() bool {
	for _, existing := range allowedTableRowsJoinTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TableRowsJoinType value.
func (v TableRowsJoinType) Ptr() *TableRowsJoinType {
	return &v
}

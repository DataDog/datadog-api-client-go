// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SegmentSortColumn Column to sort segments by.
type SegmentSortColumn string

// List of SegmentSortColumn.
const (
	SEGMENTSORTCOLUMN_NAME        SegmentSortColumn = "name"
	SEGMENTSORTCOLUMN_DESCRIPTION SegmentSortColumn = "description"
	SEGMENTSORTCOLUMN_CREATED_BY  SegmentSortColumn = "created_by"
	SEGMENTSORTCOLUMN_CREATED_AT  SegmentSortColumn = "created_at"
	SEGMENTSORTCOLUMN_MODIFIED_BY SegmentSortColumn = "modified_by"
	SEGMENTSORTCOLUMN_MODIFIED_AT SegmentSortColumn = "modified_at"
	SEGMENTSORTCOLUMN_DISABLED_BY SegmentSortColumn = "disabled_by"
	SEGMENTSORTCOLUMN_DISABLED_AT SegmentSortColumn = "disabled_at"
	SEGMENTSORTCOLUMN_UUID        SegmentSortColumn = "uuid"
)

var allowedSegmentSortColumnEnumValues = []SegmentSortColumn{
	SEGMENTSORTCOLUMN_NAME,
	SEGMENTSORTCOLUMN_DESCRIPTION,
	SEGMENTSORTCOLUMN_CREATED_BY,
	SEGMENTSORTCOLUMN_CREATED_AT,
	SEGMENTSORTCOLUMN_MODIFIED_BY,
	SEGMENTSORTCOLUMN_MODIFIED_AT,
	SEGMENTSORTCOLUMN_DISABLED_BY,
	SEGMENTSORTCOLUMN_DISABLED_AT,
	SEGMENTSORTCOLUMN_UUID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SegmentSortColumn) GetAllowedValues() []SegmentSortColumn {
	return allowedSegmentSortColumnEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SegmentSortColumn) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SegmentSortColumn(value)
	return nil
}

// NewSegmentSortColumnFromValue returns a pointer to a valid SegmentSortColumn
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSegmentSortColumnFromValue(v string) (*SegmentSortColumn, error) {
	ev := SegmentSortColumn(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SegmentSortColumn: valid values are %v", v, allowedSegmentSortColumnEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SegmentSortColumn) IsValid() bool {
	for _, existing := range allowedSegmentSortColumnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SegmentSortColumn value.
func (v SegmentSortColumn) Ptr() *SegmentSortColumn {
	return &v
}

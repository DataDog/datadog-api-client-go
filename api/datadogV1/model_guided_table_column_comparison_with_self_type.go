// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableColumnComparisonWithSelfType
type GuidedTableColumnComparisonWithSelfType string

// List of GuidedTableColumnComparisonWithSelfType.
const (
	GUIDEDTABLECOLUMNCOMPARISONWITHSELFTYPE_PERCENT_COLUMN_TOTAL GuidedTableColumnComparisonWithSelfType = "percent_column_total"
)

var allowedGuidedTableColumnComparisonWithSelfTypeEnumValues = []GuidedTableColumnComparisonWithSelfType{
	GUIDEDTABLECOLUMNCOMPARISONWITHSELFTYPE_PERCENT_COLUMN_TOTAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GuidedTableColumnComparisonWithSelfType) GetAllowedValues() []GuidedTableColumnComparisonWithSelfType {
	return allowedGuidedTableColumnComparisonWithSelfTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GuidedTableColumnComparisonWithSelfType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GuidedTableColumnComparisonWithSelfType(value)
	return nil
}

// NewGuidedTableColumnComparisonWithSelfTypeFromValue returns a pointer to a valid GuidedTableColumnComparisonWithSelfType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGuidedTableColumnComparisonWithSelfTypeFromValue(v string) (*GuidedTableColumnComparisonWithSelfType, error) {
	ev := GuidedTableColumnComparisonWithSelfType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GuidedTableColumnComparisonWithSelfType: valid values are %v", v, allowedGuidedTableColumnComparisonWithSelfTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GuidedTableColumnComparisonWithSelfType) IsValid() bool {
	for _, existing := range allowedGuidedTableColumnComparisonWithSelfTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GuidedTableColumnComparisonWithSelfType value.
func (v GuidedTableColumnComparisonWithSelfType) Ptr() *GuidedTableColumnComparisonWithSelfType {
	return &v
}

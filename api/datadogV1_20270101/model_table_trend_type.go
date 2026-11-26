// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableTrendType Test-only generated schema for the dashboard versioning proof of concept.
type TableTrendType string

// List of TableTrendType.
const (
	TABLETRENDTYPE_AREA          TableTrendType = "area"
	TABLETRENDTYPE_LINE          TableTrendType = "line"
	TABLETRENDTYPE_BARS          TableTrendType = "bars"
	TABLETRENDTYPE_ADJACENT_BARS TableTrendType = "adjacent_bars"
)

var allowedTableTrendTypeEnumValues = []TableTrendType{
	TABLETRENDTYPE_AREA,
	TABLETRENDTYPE_LINE,
	TABLETRENDTYPE_BARS,
	TABLETRENDTYPE_ADJACENT_BARS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TableTrendType) GetAllowedValues() []TableTrendType {
	return allowedTableTrendTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TableTrendType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TableTrendType(value)
	return nil
}

// NewTableTrendTypeFromValue returns a pointer to a valid TableTrendType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTableTrendTypeFromValue(v string) (*TableTrendType, error) {
	ev := TableTrendType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TableTrendType: valid values are %v", v, allowedTableTrendTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TableTrendType) IsValid() bool {
	for _, existing := range allowedTableTrendTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TableTrendType value.
func (v TableTrendType) Ptr() *TableTrendType {
	return &v
}

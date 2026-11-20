// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CellScopeType Test-only generated schema for the dashboard versioning proof of concept.
type CellScopeType string

// List of CellScopeType.
const (
	CELLSCOPETYPE_CELL CellScopeType = "cell"
)

var allowedCellScopeTypeEnumValues = []CellScopeType{
	CELLSCOPETYPE_CELL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CellScopeType) GetAllowedValues() []CellScopeType {
	return allowedCellScopeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CellScopeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CellScopeType(value)
	return nil
}

// NewCellScopeTypeFromValue returns a pointer to a valid CellScopeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCellScopeTypeFromValue(v string) (*CellScopeType, error) {
	ev := CellScopeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CellScopeType: valid values are %v", v, allowedCellScopeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CellScopeType) IsValid() bool {
	for _, existing := range allowedCellScopeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CellScopeType value.
func (v CellScopeType) Ptr() *CellScopeType {
	return &v
}

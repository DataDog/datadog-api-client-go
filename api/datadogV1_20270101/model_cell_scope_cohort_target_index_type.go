// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CellScopeCohortTargetIndexType Test-only generated schema for the dashboard versioning proof of concept.
type CellScopeCohortTargetIndexType string

// List of CellScopeCohortTargetIndexType.
const (
	CELLSCOPECOHORTTARGETINDEXTYPE_INDEX CellScopeCohortTargetIndexType = "index"
)

var allowedCellScopeCohortTargetIndexTypeEnumValues = []CellScopeCohortTargetIndexType{
	CELLSCOPECOHORTTARGETINDEXTYPE_INDEX,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CellScopeCohortTargetIndexType) GetAllowedValues() []CellScopeCohortTargetIndexType {
	return allowedCellScopeCohortTargetIndexTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CellScopeCohortTargetIndexType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CellScopeCohortTargetIndexType(value)
	return nil
}

// NewCellScopeCohortTargetIndexTypeFromValue returns a pointer to a valid CellScopeCohortTargetIndexType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCellScopeCohortTargetIndexTypeFromValue(v string) (*CellScopeCohortTargetIndexType, error) {
	ev := CellScopeCohortTargetIndexType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CellScopeCohortTargetIndexType: valid values are %v", v, allowedCellScopeCohortTargetIndexTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CellScopeCohortTargetIndexType) IsValid() bool {
	for _, existing := range allowedCellScopeCohortTargetIndexTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CellScopeCohortTargetIndexType value.
func (v CellScopeCohortTargetIndexType) Ptr() *CellScopeCohortTargetIndexType {
	return &v
}

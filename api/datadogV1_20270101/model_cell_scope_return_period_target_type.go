// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CellScopeReturnPeriodTargetType Test-only generated schema for the dashboard versioning proof of concept.
type CellScopeReturnPeriodTargetType string

// List of CellScopeReturnPeriodTargetType.
const (
	CELLSCOPERETURNPERIODTARGETTYPE_INDEX CellScopeReturnPeriodTargetType = "index"
)

var allowedCellScopeReturnPeriodTargetTypeEnumValues = []CellScopeReturnPeriodTargetType{
	CELLSCOPERETURNPERIODTARGETTYPE_INDEX,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CellScopeReturnPeriodTargetType) GetAllowedValues() []CellScopeReturnPeriodTargetType {
	return allowedCellScopeReturnPeriodTargetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CellScopeReturnPeriodTargetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CellScopeReturnPeriodTargetType(value)
	return nil
}

// NewCellScopeReturnPeriodTargetTypeFromValue returns a pointer to a valid CellScopeReturnPeriodTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCellScopeReturnPeriodTargetTypeFromValue(v string) (*CellScopeReturnPeriodTargetType, error) {
	ev := CellScopeReturnPeriodTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CellScopeReturnPeriodTargetType: valid values are %v", v, allowedCellScopeReturnPeriodTargetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CellScopeReturnPeriodTargetType) IsValid() bool {
	for _, existing := range allowedCellScopeReturnPeriodTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CellScopeReturnPeriodTargetType value.
func (v CellScopeReturnPeriodTargetType) Ptr() *CellScopeReturnPeriodTargetType {
	return &v
}

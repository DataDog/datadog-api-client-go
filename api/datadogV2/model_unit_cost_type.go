// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UnitCostType The JSON:API resource type for a unit cost.
type UnitCostType string

// List of UnitCostType.
const (
	UNITCOSTTYPE_UNIT_COST UnitCostType = "unit_cost"
)

var allowedUnitCostTypeEnumValues = []UnitCostType{
	UNITCOSTTYPE_UNIT_COST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UnitCostType) GetAllowedValues() []UnitCostType {
	return allowedUnitCostTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UnitCostType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UnitCostType(value)
	return nil
}

// NewUnitCostTypeFromValue returns a pointer to a valid UnitCostType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUnitCostTypeFromValue(v string) (*UnitCostType, error) {
	ev := UnitCostType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UnitCostType: valid values are %v", v, allowedUnitCostTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UnitCostType) IsValid() bool {
	for _, existing := range allowedUnitCostTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UnitCostType value.
func (v UnitCostType) Ptr() *UnitCostType {
	return &v
}

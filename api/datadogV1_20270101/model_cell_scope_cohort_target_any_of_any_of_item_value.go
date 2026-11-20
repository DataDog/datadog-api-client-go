// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CellScopeCohortTargetAnyOfAnyOfItemValue Test-only generated schema for the dashboard versioning proof of concept.
type CellScopeCohortTargetAnyOfAnyOfItemValue string

// List of CellScopeCohortTargetAnyOfAnyOfItemValue.
const (
	CELLSCOPECOHORTTARGETANYOFANYOFITEMVALUE_WEIGHTED_AVG CellScopeCohortTargetAnyOfAnyOfItemValue = "weighted_avg"
)

var allowedCellScopeCohortTargetAnyOfAnyOfItemValueEnumValues = []CellScopeCohortTargetAnyOfAnyOfItemValue{
	CELLSCOPECOHORTTARGETANYOFANYOFITEMVALUE_WEIGHTED_AVG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CellScopeCohortTargetAnyOfAnyOfItemValue) GetAllowedValues() []CellScopeCohortTargetAnyOfAnyOfItemValue {
	return allowedCellScopeCohortTargetAnyOfAnyOfItemValueEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CellScopeCohortTargetAnyOfAnyOfItemValue) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CellScopeCohortTargetAnyOfAnyOfItemValue(value)
	return nil
}

// NewCellScopeCohortTargetAnyOfAnyOfItemValueFromValue returns a pointer to a valid CellScopeCohortTargetAnyOfAnyOfItemValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCellScopeCohortTargetAnyOfAnyOfItemValueFromValue(v string) (*CellScopeCohortTargetAnyOfAnyOfItemValue, error) {
	ev := CellScopeCohortTargetAnyOfAnyOfItemValue(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CellScopeCohortTargetAnyOfAnyOfItemValue: valid values are %v", v, allowedCellScopeCohortTargetAnyOfAnyOfItemValueEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CellScopeCohortTargetAnyOfAnyOfItemValue) IsValid() bool {
	for _, existing := range allowedCellScopeCohortTargetAnyOfAnyOfItemValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CellScopeCohortTargetAnyOfAnyOfItemValue value.
func (v CellScopeCohortTargetAnyOfAnyOfItemValue) Ptr() *CellScopeCohortTargetAnyOfAnyOfItemValue {
	return &v
}

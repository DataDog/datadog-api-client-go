// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CellScopeCohortTargetAggregationType Test-only generated schema for the dashboard versioning proof of concept.
type CellScopeCohortTargetAggregationType string

// List of CellScopeCohortTargetAggregationType.
const (
	CELLSCOPECOHORTTARGETAGGREGATIONTYPE_AGGREGATION CellScopeCohortTargetAggregationType = "aggregation"
)

var allowedCellScopeCohortTargetAggregationTypeEnumValues = []CellScopeCohortTargetAggregationType{
	CELLSCOPECOHORTTARGETAGGREGATIONTYPE_AGGREGATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CellScopeCohortTargetAggregationType) GetAllowedValues() []CellScopeCohortTargetAggregationType {
	return allowedCellScopeCohortTargetAggregationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CellScopeCohortTargetAggregationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CellScopeCohortTargetAggregationType(value)
	return nil
}

// NewCellScopeCohortTargetAggregationTypeFromValue returns a pointer to a valid CellScopeCohortTargetAggregationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCellScopeCohortTargetAggregationTypeFromValue(v string) (*CellScopeCohortTargetAggregationType, error) {
	ev := CellScopeCohortTargetAggregationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CellScopeCohortTargetAggregationType: valid values are %v", v, allowedCellScopeCohortTargetAggregationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CellScopeCohortTargetAggregationType) IsValid() bool {
	for _, existing := range allowedCellScopeCohortTargetAggregationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CellScopeCohortTargetAggregationType value.
func (v CellScopeCohortTargetAggregationType) Ptr() *CellScopeCohortTargetAggregationType {
	return &v
}

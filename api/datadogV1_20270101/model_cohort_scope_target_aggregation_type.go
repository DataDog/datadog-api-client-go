// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CohortScopeTargetAggregationType Test-only generated schema for the dashboard versioning proof of concept.
type CohortScopeTargetAggregationType string

// List of CohortScopeTargetAggregationType.
const (
	COHORTSCOPETARGETAGGREGATIONTYPE_AGGREGATION CohortScopeTargetAggregationType = "aggregation"
)

var allowedCohortScopeTargetAggregationTypeEnumValues = []CohortScopeTargetAggregationType{
	COHORTSCOPETARGETAGGREGATIONTYPE_AGGREGATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CohortScopeTargetAggregationType) GetAllowedValues() []CohortScopeTargetAggregationType {
	return allowedCohortScopeTargetAggregationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CohortScopeTargetAggregationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CohortScopeTargetAggregationType(value)
	return nil
}

// NewCohortScopeTargetAggregationTypeFromValue returns a pointer to a valid CohortScopeTargetAggregationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCohortScopeTargetAggregationTypeFromValue(v string) (*CohortScopeTargetAggregationType, error) {
	ev := CohortScopeTargetAggregationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CohortScopeTargetAggregationType: valid values are %v", v, allowedCohortScopeTargetAggregationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CohortScopeTargetAggregationType) IsValid() bool {
	for _, existing := range allowedCohortScopeTargetAggregationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CohortScopeTargetAggregationType value.
func (v CohortScopeTargetAggregationType) Ptr() *CohortScopeTargetAggregationType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CohortScopeTargetAnyOfAnyOfItemValue Test-only generated schema for the dashboard versioning proof of concept.
type CohortScopeTargetAnyOfAnyOfItemValue string

// List of CohortScopeTargetAnyOfAnyOfItemValue.
const (
	COHORTSCOPETARGETANYOFANYOFITEMVALUE_WEIGHTED_AVG CohortScopeTargetAnyOfAnyOfItemValue = "weighted_avg"
)

var allowedCohortScopeTargetAnyOfAnyOfItemValueEnumValues = []CohortScopeTargetAnyOfAnyOfItemValue{
	COHORTSCOPETARGETANYOFANYOFITEMVALUE_WEIGHTED_AVG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CohortScopeTargetAnyOfAnyOfItemValue) GetAllowedValues() []CohortScopeTargetAnyOfAnyOfItemValue {
	return allowedCohortScopeTargetAnyOfAnyOfItemValueEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CohortScopeTargetAnyOfAnyOfItemValue) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CohortScopeTargetAnyOfAnyOfItemValue(value)
	return nil
}

// NewCohortScopeTargetAnyOfAnyOfItemValueFromValue returns a pointer to a valid CohortScopeTargetAnyOfAnyOfItemValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCohortScopeTargetAnyOfAnyOfItemValueFromValue(v string) (*CohortScopeTargetAnyOfAnyOfItemValue, error) {
	ev := CohortScopeTargetAnyOfAnyOfItemValue(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CohortScopeTargetAnyOfAnyOfItemValue: valid values are %v", v, allowedCohortScopeTargetAnyOfAnyOfItemValueEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CohortScopeTargetAnyOfAnyOfItemValue) IsValid() bool {
	for _, existing := range allowedCohortScopeTargetAnyOfAnyOfItemValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CohortScopeTargetAnyOfAnyOfItemValue value.
func (v CohortScopeTargetAnyOfAnyOfItemValue) Ptr() *CohortScopeTargetAnyOfAnyOfItemValue {
	return &v
}

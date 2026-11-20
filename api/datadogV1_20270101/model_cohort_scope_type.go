// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CohortScopeType Test-only generated schema for the dashboard versioning proof of concept.
type CohortScopeType string

// List of CohortScopeType.
const (
	COHORTSCOPETYPE_COHORT CohortScopeType = "cohort"
)

var allowedCohortScopeTypeEnumValues = []CohortScopeType{
	COHORTSCOPETYPE_COHORT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CohortScopeType) GetAllowedValues() []CohortScopeType {
	return allowedCohortScopeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CohortScopeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CohortScopeType(value)
	return nil
}

// NewCohortScopeTypeFromValue returns a pointer to a valid CohortScopeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCohortScopeTypeFromValue(v string) (*CohortScopeType, error) {
	ev := CohortScopeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CohortScopeType: valid values are %v", v, allowedCohortScopeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CohortScopeType) IsValid() bool {
	for _, existing := range allowedCohortScopeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CohortScopeType value.
func (v CohortScopeType) Ptr() *CohortScopeType {
	return &v
}

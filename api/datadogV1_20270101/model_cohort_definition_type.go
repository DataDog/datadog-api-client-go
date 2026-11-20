// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CohortDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type CohortDefinitionType string

// List of CohortDefinitionType.
const (
	COHORTDEFINITIONTYPE_COHORT CohortDefinitionType = "cohort"
)

var allowedCohortDefinitionTypeEnumValues = []CohortDefinitionType{
	COHORTDEFINITIONTYPE_COHORT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CohortDefinitionType) GetAllowedValues() []CohortDefinitionType {
	return allowedCohortDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CohortDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CohortDefinitionType(value)
	return nil
}

// NewCohortDefinitionTypeFromValue returns a pointer to a valid CohortDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCohortDefinitionTypeFromValue(v string) (*CohortDefinitionType, error) {
	ev := CohortDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CohortDefinitionType: valid values are %v", v, allowedCohortDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CohortDefinitionType) IsValid() bool {
	for _, existing := range allowedCohortDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CohortDefinitionType value.
func (v CohortDefinitionType) Ptr() *CohortDefinitionType {
	return &v
}

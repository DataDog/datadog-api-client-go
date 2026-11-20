// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CohortScopeTargetIndexType Test-only generated schema for the dashboard versioning proof of concept.
type CohortScopeTargetIndexType string

// List of CohortScopeTargetIndexType.
const (
	COHORTSCOPETARGETINDEXTYPE_INDEX CohortScopeTargetIndexType = "index"
)

var allowedCohortScopeTargetIndexTypeEnumValues = []CohortScopeTargetIndexType{
	COHORTSCOPETARGETINDEXTYPE_INDEX,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CohortScopeTargetIndexType) GetAllowedValues() []CohortScopeTargetIndexType {
	return allowedCohortScopeTargetIndexTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CohortScopeTargetIndexType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CohortScopeTargetIndexType(value)
	return nil
}

// NewCohortScopeTargetIndexTypeFromValue returns a pointer to a valid CohortScopeTargetIndexType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCohortScopeTargetIndexTypeFromValue(v string) (*CohortScopeTargetIndexType, error) {
	ev := CohortScopeTargetIndexType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CohortScopeTargetIndexType: valid values are %v", v, allowedCohortScopeTargetIndexTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CohortScopeTargetIndexType) IsValid() bool {
	for _, existing := range allowedCohortScopeTargetIndexTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CohortScopeTargetIndexType value.
func (v CohortScopeTargetIndexType) Ptr() *CohortScopeTargetIndexType {
	return &v
}

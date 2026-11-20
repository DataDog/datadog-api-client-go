// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionSearchCohortCriteriaTimeIntervalType Test-only generated schema for the dashboard versioning proof of concept.
type RetentionSearchCohortCriteriaTimeIntervalType string

// List of RetentionSearchCohortCriteriaTimeIntervalType.
const (
	RETENTIONSEARCHCOHORTCRITERIATIMEINTERVALTYPE_CALENDAR RetentionSearchCohortCriteriaTimeIntervalType = "calendar"
)

var allowedRetentionSearchCohortCriteriaTimeIntervalTypeEnumValues = []RetentionSearchCohortCriteriaTimeIntervalType{
	RETENTIONSEARCHCOHORTCRITERIATIMEINTERVALTYPE_CALENDAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RetentionSearchCohortCriteriaTimeIntervalType) GetAllowedValues() []RetentionSearchCohortCriteriaTimeIntervalType {
	return allowedRetentionSearchCohortCriteriaTimeIntervalTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RetentionSearchCohortCriteriaTimeIntervalType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RetentionSearchCohortCriteriaTimeIntervalType(value)
	return nil
}

// NewRetentionSearchCohortCriteriaTimeIntervalTypeFromValue returns a pointer to a valid RetentionSearchCohortCriteriaTimeIntervalType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRetentionSearchCohortCriteriaTimeIntervalTypeFromValue(v string) (*RetentionSearchCohortCriteriaTimeIntervalType, error) {
	ev := RetentionSearchCohortCriteriaTimeIntervalType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RetentionSearchCohortCriteriaTimeIntervalType: valid values are %v", v, allowedRetentionSearchCohortCriteriaTimeIntervalTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RetentionSearchCohortCriteriaTimeIntervalType) IsValid() bool {
	for _, existing := range allowedRetentionSearchCohortCriteriaTimeIntervalTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RetentionSearchCohortCriteriaTimeIntervalType value.
func (v RetentionSearchCohortCriteriaTimeIntervalType) Ptr() *RetentionSearchCohortCriteriaTimeIntervalType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionSearchReturnCriteriaTimeIntervalUnit Test-only generated schema for the dashboard versioning proof of concept.
type RetentionSearchReturnCriteriaTimeIntervalUnit string

// List of RetentionSearchReturnCriteriaTimeIntervalUnit.
const (
	RETENTIONSEARCHRETURNCRITERIATIMEINTERVALUNIT_DAY   RetentionSearchReturnCriteriaTimeIntervalUnit = "day"
	RETENTIONSEARCHRETURNCRITERIATIMEINTERVALUNIT_WEEK  RetentionSearchReturnCriteriaTimeIntervalUnit = "week"
	RETENTIONSEARCHRETURNCRITERIATIMEINTERVALUNIT_MONTH RetentionSearchReturnCriteriaTimeIntervalUnit = "month"
)

var allowedRetentionSearchReturnCriteriaTimeIntervalUnitEnumValues = []RetentionSearchReturnCriteriaTimeIntervalUnit{
	RETENTIONSEARCHRETURNCRITERIATIMEINTERVALUNIT_DAY,
	RETENTIONSEARCHRETURNCRITERIATIMEINTERVALUNIT_WEEK,
	RETENTIONSEARCHRETURNCRITERIATIMEINTERVALUNIT_MONTH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RetentionSearchReturnCriteriaTimeIntervalUnit) GetAllowedValues() []RetentionSearchReturnCriteriaTimeIntervalUnit {
	return allowedRetentionSearchReturnCriteriaTimeIntervalUnitEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RetentionSearchReturnCriteriaTimeIntervalUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RetentionSearchReturnCriteriaTimeIntervalUnit(value)
	return nil
}

// NewRetentionSearchReturnCriteriaTimeIntervalUnitFromValue returns a pointer to a valid RetentionSearchReturnCriteriaTimeIntervalUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRetentionSearchReturnCriteriaTimeIntervalUnitFromValue(v string) (*RetentionSearchReturnCriteriaTimeIntervalUnit, error) {
	ev := RetentionSearchReturnCriteriaTimeIntervalUnit(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RetentionSearchReturnCriteriaTimeIntervalUnit: valid values are %v", v, allowedRetentionSearchReturnCriteriaTimeIntervalUnitEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RetentionSearchReturnCriteriaTimeIntervalUnit) IsValid() bool {
	for _, existing := range allowedRetentionSearchReturnCriteriaTimeIntervalUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RetentionSearchReturnCriteriaTimeIntervalUnit value.
func (v RetentionSearchReturnCriteriaTimeIntervalUnit) Ptr() *RetentionSearchReturnCriteriaTimeIntervalUnit {
	return &v
}

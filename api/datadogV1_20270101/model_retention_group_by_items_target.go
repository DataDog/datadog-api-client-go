// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionGroupByItemsTarget Test-only generated schema for the dashboard versioning proof of concept.
type RetentionGroupByItemsTarget string

// List of RetentionGroupByItemsTarget.
const (
	RETENTIONGROUPBYITEMSTARGET_COHORT        RetentionGroupByItemsTarget = "cohort"
	RETENTIONGROUPBYITEMSTARGET_RETURN_PERIOD RetentionGroupByItemsTarget = "return_period"
)

var allowedRetentionGroupByItemsTargetEnumValues = []RetentionGroupByItemsTarget{
	RETENTIONGROUPBYITEMSTARGET_COHORT,
	RETENTIONGROUPBYITEMSTARGET_RETURN_PERIOD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RetentionGroupByItemsTarget) GetAllowedValues() []RetentionGroupByItemsTarget {
	return allowedRetentionGroupByItemsTargetEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RetentionGroupByItemsTarget) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RetentionGroupByItemsTarget(value)
	return nil
}

// NewRetentionGroupByItemsTargetFromValue returns a pointer to a valid RetentionGroupByItemsTarget
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRetentionGroupByItemsTargetFromValue(v string) (*RetentionGroupByItemsTarget, error) {
	ev := RetentionGroupByItemsTarget(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RetentionGroupByItemsTarget: valid values are %v", v, allowedRetentionGroupByItemsTargetEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RetentionGroupByItemsTarget) IsValid() bool {
	for _, existing := range allowedRetentionGroupByItemsTargetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RetentionGroupByItemsTarget value.
func (v RetentionGroupByItemsTarget) Ptr() *RetentionGroupByItemsTarget {
	return &v
}

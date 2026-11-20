// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionSearchReturnCriteriaTimeIntervalType Test-only generated schema for the dashboard versioning proof of concept.
type RetentionSearchReturnCriteriaTimeIntervalType string

// List of RetentionSearchReturnCriteriaTimeIntervalType.
const (
	RETENTIONSEARCHRETURNCRITERIATIMEINTERVALTYPE_FIXED RetentionSearchReturnCriteriaTimeIntervalType = "fixed"
)

var allowedRetentionSearchReturnCriteriaTimeIntervalTypeEnumValues = []RetentionSearchReturnCriteriaTimeIntervalType{
	RETENTIONSEARCHRETURNCRITERIATIMEINTERVALTYPE_FIXED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RetentionSearchReturnCriteriaTimeIntervalType) GetAllowedValues() []RetentionSearchReturnCriteriaTimeIntervalType {
	return allowedRetentionSearchReturnCriteriaTimeIntervalTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RetentionSearchReturnCriteriaTimeIntervalType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RetentionSearchReturnCriteriaTimeIntervalType(value)
	return nil
}

// NewRetentionSearchReturnCriteriaTimeIntervalTypeFromValue returns a pointer to a valid RetentionSearchReturnCriteriaTimeIntervalType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRetentionSearchReturnCriteriaTimeIntervalTypeFromValue(v string) (*RetentionSearchReturnCriteriaTimeIntervalType, error) {
	ev := RetentionSearchReturnCriteriaTimeIntervalType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RetentionSearchReturnCriteriaTimeIntervalType: valid values are %v", v, allowedRetentionSearchReturnCriteriaTimeIntervalTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RetentionSearchReturnCriteriaTimeIntervalType) IsValid() bool {
	for _, existing := range allowedRetentionSearchReturnCriteriaTimeIntervalTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RetentionSearchReturnCriteriaTimeIntervalType value.
func (v RetentionSearchReturnCriteriaTimeIntervalType) Ptr() *RetentionSearchReturnCriteriaTimeIntervalType {
	return &v
}

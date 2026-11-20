// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionSearchReturnCondition Test-only generated schema for the dashboard versioning proof of concept.
type RetentionSearchReturnCondition string

// List of RetentionSearchReturnCondition.
const (
	RETENTIONSEARCHRETURNCONDITION_CONVERSION_ON          RetentionSearchReturnCondition = "conversion_on"
	RETENTIONSEARCHRETURNCONDITION_CONVERSION_ON_OR_AFTER RetentionSearchReturnCondition = "conversion_on_or_after"
)

var allowedRetentionSearchReturnConditionEnumValues = []RetentionSearchReturnCondition{
	RETENTIONSEARCHRETURNCONDITION_CONVERSION_ON,
	RETENTIONSEARCHRETURNCONDITION_CONVERSION_ON_OR_AFTER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RetentionSearchReturnCondition) GetAllowedValues() []RetentionSearchReturnCondition {
	return allowedRetentionSearchReturnConditionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RetentionSearchReturnCondition) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RetentionSearchReturnCondition(value)
	return nil
}

// NewRetentionSearchReturnConditionFromValue returns a pointer to a valid RetentionSearchReturnCondition
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRetentionSearchReturnConditionFromValue(v string) (*RetentionSearchReturnCondition, error) {
	ev := RetentionSearchReturnCondition(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RetentionSearchReturnCondition: valid values are %v", v, allowedRetentionSearchReturnConditionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RetentionSearchReturnCondition) IsValid() bool {
	for _, existing := range allowedRetentionSearchReturnConditionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RetentionSearchReturnCondition value.
func (v RetentionSearchReturnCondition) Ptr() *RetentionSearchReturnCondition {
	return &v
}

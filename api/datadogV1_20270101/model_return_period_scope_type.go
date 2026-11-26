// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReturnPeriodScopeType Test-only generated schema for the dashboard versioning proof of concept.
type ReturnPeriodScopeType string

// List of ReturnPeriodScopeType.
const (
	RETURNPERIODSCOPETYPE_RETURN_PERIOD ReturnPeriodScopeType = "return_period"
)

var allowedReturnPeriodScopeTypeEnumValues = []ReturnPeriodScopeType{
	RETURNPERIODSCOPETYPE_RETURN_PERIOD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReturnPeriodScopeType) GetAllowedValues() []ReturnPeriodScopeType {
	return allowedReturnPeriodScopeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReturnPeriodScopeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReturnPeriodScopeType(value)
	return nil
}

// NewReturnPeriodScopeTypeFromValue returns a pointer to a valid ReturnPeriodScopeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReturnPeriodScopeTypeFromValue(v string) (*ReturnPeriodScopeType, error) {
	ev := ReturnPeriodScopeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReturnPeriodScopeType: valid values are %v", v, allowedReturnPeriodScopeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReturnPeriodScopeType) IsValid() bool {
	for _, existing := range allowedReturnPeriodScopeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReturnPeriodScopeType value.
func (v ReturnPeriodScopeType) Ptr() *ReturnPeriodScopeType {
	return &v
}

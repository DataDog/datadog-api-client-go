// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceLimitResourceType Limit resource type.
type GovernanceLimitResourceType string

// List of GovernanceLimitResourceType.
const (
	GOVERNANCELIMITRESOURCETYPE_LIMIT GovernanceLimitResourceType = "limit"
)

var allowedGovernanceLimitResourceTypeEnumValues = []GovernanceLimitResourceType{
	GOVERNANCELIMITRESOURCETYPE_LIMIT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GovernanceLimitResourceType) GetAllowedValues() []GovernanceLimitResourceType {
	return allowedGovernanceLimitResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GovernanceLimitResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GovernanceLimitResourceType(value)
	return nil
}

// NewGovernanceLimitResourceTypeFromValue returns a pointer to a valid GovernanceLimitResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGovernanceLimitResourceTypeFromValue(v string) (*GovernanceLimitResourceType, error) {
	ev := GovernanceLimitResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GovernanceLimitResourceType: valid values are %v", v, allowedGovernanceLimitResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GovernanceLimitResourceType) IsValid() bool {
	for _, existing := range allowedGovernanceLimitResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GovernanceLimitResourceType value.
func (v GovernanceLimitResourceType) Ptr() *GovernanceLimitResourceType {
	return &v
}

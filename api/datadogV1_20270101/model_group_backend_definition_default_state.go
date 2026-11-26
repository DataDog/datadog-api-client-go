// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GroupBackendDefinitionDefaultState Test-only generated schema for the dashboard versioning proof of concept.
type GroupBackendDefinitionDefaultState string

// List of GroupBackendDefinitionDefaultState.
const (
	GROUPBACKENDDEFINITIONDEFAULTSTATE_OPEN   GroupBackendDefinitionDefaultState = "open"
	GROUPBACKENDDEFINITIONDEFAULTSTATE_CLOSED GroupBackendDefinitionDefaultState = "closed"
)

var allowedGroupBackendDefinitionDefaultStateEnumValues = []GroupBackendDefinitionDefaultState{
	GROUPBACKENDDEFINITIONDEFAULTSTATE_OPEN,
	GROUPBACKENDDEFINITIONDEFAULTSTATE_CLOSED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GroupBackendDefinitionDefaultState) GetAllowedValues() []GroupBackendDefinitionDefaultState {
	return allowedGroupBackendDefinitionDefaultStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GroupBackendDefinitionDefaultState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GroupBackendDefinitionDefaultState(value)
	return nil
}

// NewGroupBackendDefinitionDefaultStateFromValue returns a pointer to a valid GroupBackendDefinitionDefaultState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGroupBackendDefinitionDefaultStateFromValue(v string) (*GroupBackendDefinitionDefaultState, error) {
	ev := GroupBackendDefinitionDefaultState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GroupBackendDefinitionDefaultState: valid values are %v", v, allowedGroupBackendDefinitionDefaultStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GroupBackendDefinitionDefaultState) IsValid() bool {
	for _, existing := range allowedGroupBackendDefinitionDefaultStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GroupBackendDefinitionDefaultState value.
func (v GroupBackendDefinitionDefaultState) Ptr() *GroupBackendDefinitionDefaultState {
	return &v
}

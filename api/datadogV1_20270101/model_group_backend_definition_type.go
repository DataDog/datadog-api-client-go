// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GroupBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type GroupBackendDefinitionType string

// List of GroupBackendDefinitionType.
const (
	GROUPBACKENDDEFINITIONTYPE_GROUP GroupBackendDefinitionType = "group"
)

var allowedGroupBackendDefinitionTypeEnumValues = []GroupBackendDefinitionType{
	GROUPBACKENDDEFINITIONTYPE_GROUP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GroupBackendDefinitionType) GetAllowedValues() []GroupBackendDefinitionType {
	return allowedGroupBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GroupBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GroupBackendDefinitionType(value)
	return nil
}

// NewGroupBackendDefinitionTypeFromValue returns a pointer to a valid GroupBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGroupBackendDefinitionTypeFromValue(v string) (*GroupBackendDefinitionType, error) {
	ev := GroupBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GroupBackendDefinitionType: valid values are %v", v, allowedGroupBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GroupBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedGroupBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GroupBackendDefinitionType value.
func (v GroupBackendDefinitionType) Ptr() *GroupBackendDefinitionType {
	return &v
}

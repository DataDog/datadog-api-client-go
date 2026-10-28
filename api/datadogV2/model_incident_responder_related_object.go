// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentResponderRelatedObject Object related to an incident responder.
type IncidentResponderRelatedObject string

// List of IncidentResponderRelatedObject.
const (
	INCIDENTRESPONDERRELATEDOBJECT_USERS            IncidentResponderRelatedObject = "users"
	INCIDENTRESPONDERRELATEDOBJECT_ROLE_ASSIGNMENTS IncidentResponderRelatedObject = "role_assignments"
)

var allowedIncidentResponderRelatedObjectEnumValues = []IncidentResponderRelatedObject{
	INCIDENTRESPONDERRELATEDOBJECT_USERS,
	INCIDENTRESPONDERRELATEDOBJECT_ROLE_ASSIGNMENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentResponderRelatedObject) GetAllowedValues() []IncidentResponderRelatedObject {
	return allowedIncidentResponderRelatedObjectEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentResponderRelatedObject) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentResponderRelatedObject(value)
	return nil
}

// NewIncidentResponderRelatedObjectFromValue returns a pointer to a valid IncidentResponderRelatedObject
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentResponderRelatedObjectFromValue(v string) (*IncidentResponderRelatedObject, error) {
	ev := IncidentResponderRelatedObject(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentResponderRelatedObject: valid values are %v", v, allowedIncidentResponderRelatedObjectEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentResponderRelatedObject) IsValid() bool {
	for _, existing := range allowedIncidentResponderRelatedObjectEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentResponderRelatedObject value.
func (v IncidentResponderRelatedObject) Ptr() *IncidentResponderRelatedObject {
	return &v
}

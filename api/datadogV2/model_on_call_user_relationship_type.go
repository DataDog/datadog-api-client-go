// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallUserRelationshipType The definition of `OnCallUserRelationshipType` object.
type OnCallUserRelationshipType string

// List of OnCallUserRelationshipType.
const (
	ONCALLUSERRELATIONSHIPTYPE_USERS OnCallUserRelationshipType = "users"
)

var allowedOnCallUserRelationshipTypeEnumValues = []OnCallUserRelationshipType{
	ONCALLUSERRELATIONSHIPTYPE_USERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnCallUserRelationshipType) GetAllowedValues() []OnCallUserRelationshipType {
	return allowedOnCallUserRelationshipTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnCallUserRelationshipType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnCallUserRelationshipType(value)
	return nil
}

// NewOnCallUserRelationshipTypeFromValue returns a pointer to a valid OnCallUserRelationshipType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnCallUserRelationshipTypeFromValue(v string) (*OnCallUserRelationshipType, error) {
	ev := OnCallUserRelationshipType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnCallUserRelationshipType: valid values are %v", v, allowedOnCallUserRelationshipTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnCallUserRelationshipType) IsValid() bool {
	for _, existing := range allowedOnCallUserRelationshipTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnCallUserRelationshipType value.
func (v OnCallUserRelationshipType) Ptr() *OnCallUserRelationshipType {
	return &v
}

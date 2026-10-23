// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageTargetRelationshipDataType The resource type of the page target.
type OnCallPageTargetRelationshipDataType string

// List of OnCallPageTargetRelationshipDataType.
const (
	ONCALLPAGETARGETRELATIONSHIPDATATYPE_TEAMS             OnCallPageTargetRelationshipDataType = "teams"
	ONCALLPAGETARGETRELATIONSHIPDATATYPE_USERS             OnCallPageTargetRelationshipDataType = "users"
	ONCALLPAGETARGETRELATIONSHIPDATATYPE_CROSS_ORG_HANDLES OnCallPageTargetRelationshipDataType = "cross_org_handles"
)

var allowedOnCallPageTargetRelationshipDataTypeEnumValues = []OnCallPageTargetRelationshipDataType{
	ONCALLPAGETARGETRELATIONSHIPDATATYPE_TEAMS,
	ONCALLPAGETARGETRELATIONSHIPDATATYPE_USERS,
	ONCALLPAGETARGETRELATIONSHIPDATATYPE_CROSS_ORG_HANDLES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnCallPageTargetRelationshipDataType) GetAllowedValues() []OnCallPageTargetRelationshipDataType {
	return allowedOnCallPageTargetRelationshipDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnCallPageTargetRelationshipDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnCallPageTargetRelationshipDataType(value)
	return nil
}

// NewOnCallPageTargetRelationshipDataTypeFromValue returns a pointer to a valid OnCallPageTargetRelationshipDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnCallPageTargetRelationshipDataTypeFromValue(v string) (*OnCallPageTargetRelationshipDataType, error) {
	ev := OnCallPageTargetRelationshipDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnCallPageTargetRelationshipDataType: valid values are %v", v, allowedOnCallPageTargetRelationshipDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnCallPageTargetRelationshipDataType) IsValid() bool {
	for _, existing := range allowedOnCallPageTargetRelationshipDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnCallPageTargetRelationshipDataType value.
func (v OnCallPageTargetRelationshipDataType) Ptr() *OnCallPageTargetRelationshipDataType {
	return &v
}

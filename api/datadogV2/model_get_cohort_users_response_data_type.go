// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetCohortUsersResponseDataType
type GetCohortUsersResponseDataType string

// List of GetCohortUsersResponseDataType.
const (
	GETCOHORTUSERSRESPONSEDATATYPE_COHORT_USERS_RESPONSE GetCohortUsersResponseDataType = "cohort_users_response"
)

var allowedGetCohortUsersResponseDataTypeEnumValues = []GetCohortUsersResponseDataType{
	GETCOHORTUSERSRESPONSEDATATYPE_COHORT_USERS_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GetCohortUsersResponseDataType) GetAllowedValues() []GetCohortUsersResponseDataType {
	return allowedGetCohortUsersResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GetCohortUsersResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GetCohortUsersResponseDataType(value)
	return nil
}

// NewGetCohortUsersResponseDataTypeFromValue returns a pointer to a valid GetCohortUsersResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGetCohortUsersResponseDataTypeFromValue(v string) (*GetCohortUsersResponseDataType, error) {
	ev := GetCohortUsersResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GetCohortUsersResponseDataType: valid values are %v", v, allowedGetCohortUsersResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GetCohortUsersResponseDataType) IsValid() bool {
	for _, existing := range allowedGetCohortUsersResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GetCohortUsersResponseDataType value.
func (v GetCohortUsersResponseDataType) Ptr() *GetCohortUsersResponseDataType {
	return &v
}

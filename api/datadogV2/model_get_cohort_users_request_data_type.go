// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetCohortUsersRequestDataType
type GetCohortUsersRequestDataType string

// List of GetCohortUsersRequestDataType.
const (
	GETCOHORTUSERSREQUESTDATATYPE_COHORT_USERS_REQUEST GetCohortUsersRequestDataType = "cohort_users_request"
)

var allowedGetCohortUsersRequestDataTypeEnumValues = []GetCohortUsersRequestDataType{
	GETCOHORTUSERSREQUESTDATATYPE_COHORT_USERS_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GetCohortUsersRequestDataType) GetAllowedValues() []GetCohortUsersRequestDataType {
	return allowedGetCohortUsersRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GetCohortUsersRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GetCohortUsersRequestDataType(value)
	return nil
}

// NewGetCohortUsersRequestDataTypeFromValue returns a pointer to a valid GetCohortUsersRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGetCohortUsersRequestDataTypeFromValue(v string) (*GetCohortUsersRequestDataType, error) {
	ev := GetCohortUsersRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GetCohortUsersRequestDataType: valid values are %v", v, allowedGetCohortUsersRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GetCohortUsersRequestDataType) IsValid() bool {
	for _, existing := range allowedGetCohortUsersRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GetCohortUsersRequestDataType value.
func (v GetCohortUsersRequestDataType) Ptr() *GetCohortUsersRequestDataType {
	return &v
}

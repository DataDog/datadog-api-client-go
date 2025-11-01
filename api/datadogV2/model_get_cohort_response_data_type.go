// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetCohortResponseDataType
type GetCohortResponseDataType string

// List of GetCohortResponseDataType.
const (
	GETCOHORTRESPONSEDATATYPE_COHORT_RESPONSE GetCohortResponseDataType = "cohort_response"
)

var allowedGetCohortResponseDataTypeEnumValues = []GetCohortResponseDataType{
	GETCOHORTRESPONSEDATATYPE_COHORT_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GetCohortResponseDataType) GetAllowedValues() []GetCohortResponseDataType {
	return allowedGetCohortResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GetCohortResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GetCohortResponseDataType(value)
	return nil
}

// NewGetCohortResponseDataTypeFromValue returns a pointer to a valid GetCohortResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGetCohortResponseDataTypeFromValue(v string) (*GetCohortResponseDataType, error) {
	ev := GetCohortResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GetCohortResponseDataType: valid values are %v", v, allowedGetCohortResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GetCohortResponseDataType) IsValid() bool {
	for _, existing := range allowedGetCohortResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GetCohortResponseDataType value.
func (v GetCohortResponseDataType) Ptr() *GetCohortResponseDataType {
	return &v
}

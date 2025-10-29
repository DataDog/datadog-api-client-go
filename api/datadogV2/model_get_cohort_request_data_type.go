// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetCohortRequestDataType
type GetCohortRequestDataType string

// List of GetCohortRequestDataType.
const (
	GETCOHORTREQUESTDATATYPE_COHORT_REQUEST GetCohortRequestDataType = "cohort_request"
)

var allowedGetCohortRequestDataTypeEnumValues = []GetCohortRequestDataType{
	GETCOHORTREQUESTDATATYPE_COHORT_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GetCohortRequestDataType) GetAllowedValues() []GetCohortRequestDataType {
	return allowedGetCohortRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GetCohortRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GetCohortRequestDataType(value)
	return nil
}

// NewGetCohortRequestDataTypeFromValue returns a pointer to a valid GetCohortRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGetCohortRequestDataTypeFromValue(v string) (*GetCohortRequestDataType, error) {
	ev := GetCohortRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GetCohortRequestDataType: valid values are %v", v, allowedGetCohortRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GetCohortRequestDataType) IsValid() bool {
	for _, existing := range allowedGetCohortRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GetCohortRequestDataType value.
func (v GetCohortRequestDataType) Ptr() *GetCohortRequestDataType {
	return &v
}

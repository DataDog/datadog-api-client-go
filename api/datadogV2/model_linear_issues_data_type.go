// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LinearIssuesDataType Linear issues resource type.
type LinearIssuesDataType string

// List of LinearIssuesDataType.
const (
	LINEARISSUESDATATYPE_LINEAR_ISSUES LinearIssuesDataType = "linear_issues"
)

var allowedLinearIssuesDataTypeEnumValues = []LinearIssuesDataType{
	LINEARISSUESDATATYPE_LINEAR_ISSUES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LinearIssuesDataType) GetAllowedValues() []LinearIssuesDataType {
	return allowedLinearIssuesDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LinearIssuesDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LinearIssuesDataType(value)
	return nil
}

// NewLinearIssuesDataTypeFromValue returns a pointer to a valid LinearIssuesDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLinearIssuesDataTypeFromValue(v string) (*LinearIssuesDataType, error) {
	ev := LinearIssuesDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LinearIssuesDataType: valid values are %v", v, allowedLinearIssuesDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LinearIssuesDataType) IsValid() bool {
	for _, existing := range allowedLinearIssuesDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LinearIssuesDataType value.
func (v LinearIssuesDataType) Ptr() *LinearIssuesDataType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueDefinitionDataType Issue definitions resource type.
type IssueDefinitionDataType string

// List of IssueDefinitionDataType.
const (
	ISSUEDEFINITIONDATATYPE_ISSUE_DEFINITIONS IssueDefinitionDataType = "issue_definitions"
)

var allowedIssueDefinitionDataTypeEnumValues = []IssueDefinitionDataType{
	ISSUEDEFINITIONDATATYPE_ISSUE_DEFINITIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssueDefinitionDataType) GetAllowedValues() []IssueDefinitionDataType {
	return allowedIssueDefinitionDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssueDefinitionDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssueDefinitionDataType(value)
	return nil
}

// NewIssueDefinitionDataTypeFromValue returns a pointer to a valid IssueDefinitionDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssueDefinitionDataTypeFromValue(v string) (*IssueDefinitionDataType, error) {
	ev := IssueDefinitionDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssueDefinitionDataType: valid values are %v", v, allowedIssueDefinitionDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssueDefinitionDataType) IsValid() bool {
	for _, existing := range allowedIssueDefinitionDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueDefinitionDataType value.
func (v IssueDefinitionDataType) Ptr() *IssueDefinitionDataType {
	return &v
}

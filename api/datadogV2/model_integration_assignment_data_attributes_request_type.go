// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationAssignmentDataAttributesRequestType Type of the security issues. Can be "findings" or "vulnerabilities".
type IntegrationAssignmentDataAttributesRequestType string

// List of IntegrationAssignmentDataAttributesRequestType.
const (
	INTEGRATIONASSIGNMENTDATAATTRIBUTESREQUESTTYPE_FINDINGS        IntegrationAssignmentDataAttributesRequestType = "findings"
	INTEGRATIONASSIGNMENTDATAATTRIBUTESREQUESTTYPE_VULNERABILITIES IntegrationAssignmentDataAttributesRequestType = "vulnerabilities"
)

var allowedIntegrationAssignmentDataAttributesRequestTypeEnumValues = []IntegrationAssignmentDataAttributesRequestType{
	INTEGRATIONASSIGNMENTDATAATTRIBUTESREQUESTTYPE_FINDINGS,
	INTEGRATIONASSIGNMENTDATAATTRIBUTESREQUESTTYPE_VULNERABILITIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IntegrationAssignmentDataAttributesRequestType) GetAllowedValues() []IntegrationAssignmentDataAttributesRequestType {
	return allowedIntegrationAssignmentDataAttributesRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IntegrationAssignmentDataAttributesRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IntegrationAssignmentDataAttributesRequestType(value)
	return nil
}

// NewIntegrationAssignmentDataAttributesRequestTypeFromValue returns a pointer to a valid IntegrationAssignmentDataAttributesRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIntegrationAssignmentDataAttributesRequestTypeFromValue(v string) (*IntegrationAssignmentDataAttributesRequestType, error) {
	ev := IntegrationAssignmentDataAttributesRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IntegrationAssignmentDataAttributesRequestType: valid values are %v", v, allowedIntegrationAssignmentDataAttributesRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IntegrationAssignmentDataAttributesRequestType) IsValid() bool {
	for _, existing := range allowedIntegrationAssignmentDataAttributesRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IntegrationAssignmentDataAttributesRequestType value.
func (v IntegrationAssignmentDataAttributesRequestType) Ptr() *IntegrationAssignmentDataAttributesRequestType {
	return &v
}

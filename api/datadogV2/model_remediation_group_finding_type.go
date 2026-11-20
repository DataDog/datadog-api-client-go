// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationGroupFindingType Value of the findings' `@finding_type` attribute that a grouping definition covers.
type RemediationGroupFindingType string

// List of RemediationGroupFindingType.
const (
	REMEDIATIONGROUPFINDINGTYPE_LIBRARY_VULNERABILITY            RemediationGroupFindingType = "library_vulnerability"
	REMEDIATIONGROUPFINDINGTYPE_HOST_AND_CONTAINER_VULNERABILITY RemediationGroupFindingType = "host_and_container_vulnerability"
)

var allowedRemediationGroupFindingTypeEnumValues = []RemediationGroupFindingType{
	REMEDIATIONGROUPFINDINGTYPE_LIBRARY_VULNERABILITY,
	REMEDIATIONGROUPFINDINGTYPE_HOST_AND_CONTAINER_VULNERABILITY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RemediationGroupFindingType) GetAllowedValues() []RemediationGroupFindingType {
	return allowedRemediationGroupFindingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RemediationGroupFindingType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RemediationGroupFindingType(value)
	return nil
}

// NewRemediationGroupFindingTypeFromValue returns a pointer to a valid RemediationGroupFindingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRemediationGroupFindingTypeFromValue(v string) (*RemediationGroupFindingType, error) {
	ev := RemediationGroupFindingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RemediationGroupFindingType: valid values are %v", v, allowedRemediationGroupFindingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RemediationGroupFindingType) IsValid() bool {
	for _, existing := range allowedRemediationGroupFindingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RemediationGroupFindingType value.
func (v RemediationGroupFindingType) Ptr() *RemediationGroupFindingType {
	return &v
}

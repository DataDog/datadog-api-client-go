// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationGroupType Remediation groups resource type.
type RemediationGroupType string

// List of RemediationGroupType.
const (
	REMEDIATIONGROUPTYPE_REMEDIATION_GROUPS RemediationGroupType = "remediation_groups"
)

var allowedRemediationGroupTypeEnumValues = []RemediationGroupType{
	REMEDIATIONGROUPTYPE_REMEDIATION_GROUPS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RemediationGroupType) GetAllowedValues() []RemediationGroupType {
	return allowedRemediationGroupTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RemediationGroupType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RemediationGroupType(value)
	return nil
}

// NewRemediationGroupTypeFromValue returns a pointer to a valid RemediationGroupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRemediationGroupTypeFromValue(v string) (*RemediationGroupType, error) {
	ev := RemediationGroupType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RemediationGroupType: valid values are %v", v, allowedRemediationGroupTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RemediationGroupType) IsValid() bool {
	for _, existing := range allowedRemediationGroupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RemediationGroupType value.
func (v RemediationGroupType) Ptr() *RemediationGroupType {
	return &v
}

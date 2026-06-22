// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationLaunchType ECS launch type.
type RemediationLaunchType string

// List of RemediationLaunchType.
const (
	REMEDIATIONLAUNCHTYPE_EC2     RemediationLaunchType = "EC2"
	REMEDIATIONLAUNCHTYPE_FARGATE RemediationLaunchType = "FARGATE"
)

var allowedRemediationLaunchTypeEnumValues = []RemediationLaunchType{
	REMEDIATIONLAUNCHTYPE_EC2,
	REMEDIATIONLAUNCHTYPE_FARGATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RemediationLaunchType) GetAllowedValues() []RemediationLaunchType {
	return allowedRemediationLaunchTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RemediationLaunchType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RemediationLaunchType(value)
	return nil
}

// NewRemediationLaunchTypeFromValue returns a pointer to a valid RemediationLaunchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRemediationLaunchTypeFromValue(v string) (*RemediationLaunchType, error) {
	ev := RemediationLaunchType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RemediationLaunchType: valid values are %v", v, allowedRemediationLaunchTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RemediationLaunchType) IsValid() bool {
	for _, existing := range allowedRemediationLaunchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RemediationLaunchType value.
func (v RemediationLaunchType) Ptr() *RemediationLaunchType {
	return &v
}

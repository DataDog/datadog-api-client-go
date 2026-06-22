// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationWorkloadType The kind of ECS workload that owns the problematic tasks.
type RemediationWorkloadType string

// List of RemediationWorkloadType.
const (
	REMEDIATIONWORKLOADTYPE_WORKLOAD_TYPE_UNSPECIFIED RemediationWorkloadType = "WORKLOAD_TYPE_UNSPECIFIED"
	REMEDIATIONWORKLOADTYPE_SERVICE                   RemediationWorkloadType = "SERVICE"
	REMEDIATIONWORKLOADTYPE_STANDALONE_TASK           RemediationWorkloadType = "STANDALONE_TASK"
	REMEDIATIONWORKLOADTYPE_DAEMON                    RemediationWorkloadType = "DAEMON"
)

var allowedRemediationWorkloadTypeEnumValues = []RemediationWorkloadType{
	REMEDIATIONWORKLOADTYPE_WORKLOAD_TYPE_UNSPECIFIED,
	REMEDIATIONWORKLOADTYPE_SERVICE,
	REMEDIATIONWORKLOADTYPE_STANDALONE_TASK,
	REMEDIATIONWORKLOADTYPE_DAEMON,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RemediationWorkloadType) GetAllowedValues() []RemediationWorkloadType {
	return allowedRemediationWorkloadTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RemediationWorkloadType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RemediationWorkloadType(value)
	return nil
}

// NewRemediationWorkloadTypeFromValue returns a pointer to a valid RemediationWorkloadType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRemediationWorkloadTypeFromValue(v string) (*RemediationWorkloadType, error) {
	ev := RemediationWorkloadType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RemediationWorkloadType: valid values are %v", v, allowedRemediationWorkloadTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RemediationWorkloadType) IsValid() bool {
	for _, existing := range allowedRemediationWorkloadTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RemediationWorkloadType value.
func (v RemediationWorkloadType) Ptr() *RemediationWorkloadType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudWorkloadSecurityAgentPolicyBatchUpdateDataType Type of the batch update operation
type CloudWorkloadSecurityAgentPolicyBatchUpdateDataType string

// List of CloudWorkloadSecurityAgentPolicyBatchUpdateDataType.
const (
	CLOUDWORKLOADSECURITYAGENTPOLICYBATCHUPDATEDATATYPE_POLICIES CloudWorkloadSecurityAgentPolicyBatchUpdateDataType = "policies"
)

var allowedCloudWorkloadSecurityAgentPolicyBatchUpdateDataTypeEnumValues = []CloudWorkloadSecurityAgentPolicyBatchUpdateDataType{
	CLOUDWORKLOADSECURITYAGENTPOLICYBATCHUPDATEDATATYPE_POLICIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CloudWorkloadSecurityAgentPolicyBatchUpdateDataType) GetAllowedValues() []CloudWorkloadSecurityAgentPolicyBatchUpdateDataType {
	return allowedCloudWorkloadSecurityAgentPolicyBatchUpdateDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CloudWorkloadSecurityAgentPolicyBatchUpdateDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CloudWorkloadSecurityAgentPolicyBatchUpdateDataType(value)
	return nil
}

// NewCloudWorkloadSecurityAgentPolicyBatchUpdateDataTypeFromValue returns a pointer to a valid CloudWorkloadSecurityAgentPolicyBatchUpdateDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCloudWorkloadSecurityAgentPolicyBatchUpdateDataTypeFromValue(v string) (*CloudWorkloadSecurityAgentPolicyBatchUpdateDataType, error) {
	ev := CloudWorkloadSecurityAgentPolicyBatchUpdateDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CloudWorkloadSecurityAgentPolicyBatchUpdateDataType: valid values are %v", v, allowedCloudWorkloadSecurityAgentPolicyBatchUpdateDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CloudWorkloadSecurityAgentPolicyBatchUpdateDataType) IsValid() bool {
	for _, existing := range allowedCloudWorkloadSecurityAgentPolicyBatchUpdateDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudWorkloadSecurityAgentPolicyBatchUpdateDataType value.
func (v CloudWorkloadSecurityAgentPolicyBatchUpdateDataType) Ptr() *CloudWorkloadSecurityAgentPolicyBatchUpdateDataType {
	return &v
}

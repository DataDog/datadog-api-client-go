// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PolicyResultType The type of the resource. The value should always be `policy_result`.
type PolicyResultType string

// List of PolicyResultType.
const (
	POLICYRESULTTYPE_POLICY_RESULT PolicyResultType = "policy_result"
)

var allowedPolicyResultTypeEnumValues = []PolicyResultType{
	POLICYRESULTTYPE_POLICY_RESULT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PolicyResultType) GetAllowedValues() []PolicyResultType {
	return allowedPolicyResultTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PolicyResultType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PolicyResultType(value)
	return nil
}

// NewPolicyResultTypeFromValue returns a pointer to a valid PolicyResultType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPolicyResultTypeFromValue(v string) (*PolicyResultType, error) {
	ev := PolicyResultType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PolicyResultType: valid values are %v", v, allowedPolicyResultTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PolicyResultType) IsValid() bool {
	for _, existing := range allowedPolicyResultTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PolicyResultType value.
func (v PolicyResultType) Ptr() *PolicyResultType {
	return &v
}

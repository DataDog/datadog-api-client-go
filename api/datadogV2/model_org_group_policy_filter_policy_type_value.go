// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPolicyFilterPolicyTypeValue The type of the policy to filter by. `org_config` indicates a policy backed by an organization configuration setting. `role` indicates a policy backed by a Datadog custom role.
type OrgGroupPolicyFilterPolicyTypeValue string

// List of OrgGroupPolicyFilterPolicyTypeValue.
const (
	ORGGROUPPOLICYFILTERPOLICYTYPEVALUE_ORG_CONFIG OrgGroupPolicyFilterPolicyTypeValue = "org_config"
	ORGGROUPPOLICYFILTERPOLICYTYPEVALUE_ROLE       OrgGroupPolicyFilterPolicyTypeValue = "role"
)

var allowedOrgGroupPolicyFilterPolicyTypeValueEnumValues = []OrgGroupPolicyFilterPolicyTypeValue{
	ORGGROUPPOLICYFILTERPOLICYTYPEVALUE_ORG_CONFIG,
	ORGGROUPPOLICYFILTERPOLICYTYPEVALUE_ROLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrgGroupPolicyFilterPolicyTypeValue) GetAllowedValues() []OrgGroupPolicyFilterPolicyTypeValue {
	return allowedOrgGroupPolicyFilterPolicyTypeValueEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgGroupPolicyFilterPolicyTypeValue) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgGroupPolicyFilterPolicyTypeValue(value)
	return nil
}

// NewOrgGroupPolicyFilterPolicyTypeValueFromValue returns a pointer to a valid OrgGroupPolicyFilterPolicyTypeValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgGroupPolicyFilterPolicyTypeValueFromValue(v string) (*OrgGroupPolicyFilterPolicyTypeValue, error) {
	ev := OrgGroupPolicyFilterPolicyTypeValue(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgGroupPolicyFilterPolicyTypeValue: valid values are %v", v, allowedOrgGroupPolicyFilterPolicyTypeValueEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgGroupPolicyFilterPolicyTypeValue) IsValid() bool {
	for _, existing := range allowedOrgGroupPolicyFilterPolicyTypeValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgGroupPolicyFilterPolicyTypeValue value.
func (v OrgGroupPolicyFilterPolicyTypeValue) Ptr() *OrgGroupPolicyFilterPolicyTypeValue {
	return &v
}

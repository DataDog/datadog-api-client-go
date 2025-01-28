// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityRuleTypesItems Security rule type
type SecurityRuleTypesItems string

// List of SecurityRuleTypesItems.
const (
	SECURITYRULETYPESITEMS_APPLICATION_CODE_VULNERABILITY    SecurityRuleTypesItems = "application_code_vulnerability"
	SECURITYRULETYPESITEMS_APPLICATION_LIBRARY_VULNERABILITY SecurityRuleTypesItems = "application_library_vulnerability"
	SECURITYRULETYPESITEMS_ATTACK_PATH                       SecurityRuleTypesItems = "attack_path"
	SECURITYRULETYPESITEMS_CONTAINER_IMAGE_VULNERABILITY     SecurityRuleTypesItems = "container_image_vulnerability"
	SECURITYRULETYPESITEMS_HOST_VULNERABILITY                SecurityRuleTypesItems = "host_vulnerability"
	SECURITYRULETYPESITEMS_IDENTITY_RISK                     SecurityRuleTypesItems = "identity_risk"
	SECURITYRULETYPESITEMS_MISCONFIGURATION                  SecurityRuleTypesItems = "misconfiguration"
	SECURITYRULETYPESITEMS_API_SECURITY                      SecurityRuleTypesItems = "api_security"
)

var allowedSecurityRuleTypesItemsEnumValues = []SecurityRuleTypesItems{
	SECURITYRULETYPESITEMS_APPLICATION_CODE_VULNERABILITY,
	SECURITYRULETYPESITEMS_APPLICATION_LIBRARY_VULNERABILITY,
	SECURITYRULETYPESITEMS_ATTACK_PATH,
	SECURITYRULETYPESITEMS_CONTAINER_IMAGE_VULNERABILITY,
	SECURITYRULETYPESITEMS_HOST_VULNERABILITY,
	SECURITYRULETYPESITEMS_IDENTITY_RISK,
	SECURITYRULETYPESITEMS_MISCONFIGURATION,
	SECURITYRULETYPESITEMS_API_SECURITY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityRuleTypesItems) GetAllowedValues() []SecurityRuleTypesItems {
	return allowedSecurityRuleTypesItemsEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityRuleTypesItems) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityRuleTypesItems(value)
	return nil
}

// NewSecurityRuleTypesItemsFromValue returns a pointer to a valid SecurityRuleTypesItems
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityRuleTypesItemsFromValue(v string) (*SecurityRuleTypesItems, error) {
	ev := SecurityRuleTypesItems(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityRuleTypesItems: valid values are %v", v, allowedSecurityRuleTypesItemsEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityRuleTypesItems) IsValid() bool {
	for _, existing := range allowedSecurityRuleTypesItemsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityRuleTypesItems value.
func (v SecurityRuleTypesItems) Ptr() *SecurityRuleTypesItems {
	return &v
}

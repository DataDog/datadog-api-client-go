// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RuleTypesItems Security rule type
type RuleTypesItems string

// List of RuleTypesItems.
const (
	RULETYPESITEMS_APPLICATION_CODE_VULNERABILITY    RuleTypesItems = "application_code_vulnerability"
	RULETYPESITEMS_APPLICATION_LIBRARY_VULNERABILITY RuleTypesItems = "application_library_vulnerability"
	RULETYPESITEMS_ATTACK_PATH                       RuleTypesItems = "attack_path"
	RULETYPESITEMS_CONTAINER_IMAGE_VULNERABILITY     RuleTypesItems = "container_image_vulnerability"
	RULETYPESITEMS_HOST_VULNERABILITY                RuleTypesItems = "host_vulnerability"
	RULETYPESITEMS_IDENTITY_RISK                     RuleTypesItems = "identity_risk"
	RULETYPESITEMS_MISCONFIGURATION                  RuleTypesItems = "misconfiguration"
	RULETYPESITEMS_API_SECURITY                      RuleTypesItems = "api_security"
)

var allowedRuleTypesItemsEnumValues = []RuleTypesItems{
	RULETYPESITEMS_APPLICATION_CODE_VULNERABILITY,
	RULETYPESITEMS_APPLICATION_LIBRARY_VULNERABILITY,
	RULETYPESITEMS_ATTACK_PATH,
	RULETYPESITEMS_CONTAINER_IMAGE_VULNERABILITY,
	RULETYPESITEMS_HOST_VULNERABILITY,
	RULETYPESITEMS_IDENTITY_RISK,
	RULETYPESITEMS_MISCONFIGURATION,
	RULETYPESITEMS_API_SECURITY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RuleTypesItems) GetAllowedValues() []RuleTypesItems {
	return allowedRuleTypesItemsEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RuleTypesItems) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RuleTypesItems(value)
	return nil
}

// NewRuleTypesItemsFromValue returns a pointer to a valid RuleTypesItems
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRuleTypesItemsFromValue(v string) (*RuleTypesItems, error) {
	ev := RuleTypesItems(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RuleTypesItems: valid values are %v", v, allowedRuleTypesItemsEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RuleTypesItems) IsValid() bool {
	for _, existing := range allowedRuleTypesItemsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RuleTypesItems value.
func (v RuleTypesItems) Ptr() *RuleTypesItems {
	return &v
}

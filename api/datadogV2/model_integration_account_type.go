// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationAccountType The type of the integration account resource. Always `integration-account`.
type IntegrationAccountType string

// List of IntegrationAccountType.
const (
	INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT IntegrationAccountType = "integration-account"
)

var allowedIntegrationAccountTypeEnumValues = []IntegrationAccountType{
	INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IntegrationAccountType) GetAllowedValues() []IntegrationAccountType {
	return allowedIntegrationAccountTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IntegrationAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IntegrationAccountType(value)
	return nil
}

// NewIntegrationAccountTypeFromValue returns a pointer to a valid IntegrationAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIntegrationAccountTypeFromValue(v string) (*IntegrationAccountType, error) {
	ev := IntegrationAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IntegrationAccountType: valid values are %v", v, allowedIntegrationAccountTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IntegrationAccountType) IsValid() bool {
	for _, existing := range allowedIntegrationAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IntegrationAccountType value.
func (v IntegrationAccountType) Ptr() *IntegrationAccountType {
	return &v
}

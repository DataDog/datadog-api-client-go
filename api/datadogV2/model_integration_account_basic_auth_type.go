// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationAccountBasicAuthType The authentication method type.
type IntegrationAccountBasicAuthType string

// List of IntegrationAccountBasicAuthType.
const (
	INTEGRATIONACCOUNTBASICAUTHTYPE_BASIC IntegrationAccountBasicAuthType = "basic"
)

var allowedIntegrationAccountBasicAuthTypeEnumValues = []IntegrationAccountBasicAuthType{
	INTEGRATIONACCOUNTBASICAUTHTYPE_BASIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IntegrationAccountBasicAuthType) GetAllowedValues() []IntegrationAccountBasicAuthType {
	return allowedIntegrationAccountBasicAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IntegrationAccountBasicAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IntegrationAccountBasicAuthType(value)
	return nil
}

// NewIntegrationAccountBasicAuthTypeFromValue returns a pointer to a valid IntegrationAccountBasicAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIntegrationAccountBasicAuthTypeFromValue(v string) (*IntegrationAccountBasicAuthType, error) {
	ev := IntegrationAccountBasicAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IntegrationAccountBasicAuthType: valid values are %v", v, allowedIntegrationAccountBasicAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IntegrationAccountBasicAuthType) IsValid() bool {
	for _, existing := range allowedIntegrationAccountBasicAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IntegrationAccountBasicAuthType value.
func (v IntegrationAccountBasicAuthType) Ptr() *IntegrationAccountBasicAuthType {
	return &v
}

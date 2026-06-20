// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AmsIntegrationAccountType The JSON:API type for web integration accounts.
type AmsIntegrationAccountType string

// List of AmsIntegrationAccountType.
const (
	AMSINTEGRATIONACCOUNTTYPE_ACCOUNT AmsIntegrationAccountType = "Account"
)

var allowedAmsIntegrationAccountTypeEnumValues = []AmsIntegrationAccountType{
	AMSINTEGRATIONACCOUNTTYPE_ACCOUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AmsIntegrationAccountType) GetAllowedValues() []AmsIntegrationAccountType {
	return allowedAmsIntegrationAccountTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AmsIntegrationAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AmsIntegrationAccountType(value)
	return nil
}

// NewAmsIntegrationAccountTypeFromValue returns a pointer to a valid AmsIntegrationAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAmsIntegrationAccountTypeFromValue(v string) (*AmsIntegrationAccountType, error) {
	ev := AmsIntegrationAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AmsIntegrationAccountType: valid values are %v", v, allowedAmsIntegrationAccountTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AmsIntegrationAccountType) IsValid() bool {
	for _, existing := range allowedAmsIntegrationAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AmsIntegrationAccountType value.
func (v AmsIntegrationAccountType) Ptr() *AmsIntegrationAccountType {
	return &v
}

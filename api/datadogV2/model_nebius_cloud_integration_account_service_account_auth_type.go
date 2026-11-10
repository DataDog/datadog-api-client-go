// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NebiusCloudIntegrationAccountServiceAccountAuthType The authentication method type.
type NebiusCloudIntegrationAccountServiceAccountAuthType string

// List of NebiusCloudIntegrationAccountServiceAccountAuthType.
const (
	NEBIUSCLOUDINTEGRATIONACCOUNTSERVICEACCOUNTAUTHTYPE_NEBIUS_SERVICE_ACCOUNT NebiusCloudIntegrationAccountServiceAccountAuthType = "nebius_service_account"
)

var allowedNebiusCloudIntegrationAccountServiceAccountAuthTypeEnumValues = []NebiusCloudIntegrationAccountServiceAccountAuthType{
	NEBIUSCLOUDINTEGRATIONACCOUNTSERVICEACCOUNTAUTHTYPE_NEBIUS_SERVICE_ACCOUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NebiusCloudIntegrationAccountServiceAccountAuthType) GetAllowedValues() []NebiusCloudIntegrationAccountServiceAccountAuthType {
	return allowedNebiusCloudIntegrationAccountServiceAccountAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NebiusCloudIntegrationAccountServiceAccountAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NebiusCloudIntegrationAccountServiceAccountAuthType(value)
	return nil
}

// NewNebiusCloudIntegrationAccountServiceAccountAuthTypeFromValue returns a pointer to a valid NebiusCloudIntegrationAccountServiceAccountAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNebiusCloudIntegrationAccountServiceAccountAuthTypeFromValue(v string) (*NebiusCloudIntegrationAccountServiceAccountAuthType, error) {
	ev := NebiusCloudIntegrationAccountServiceAccountAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NebiusCloudIntegrationAccountServiceAccountAuthType: valid values are %v", v, allowedNebiusCloudIntegrationAccountServiceAccountAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NebiusCloudIntegrationAccountServiceAccountAuthType) IsValid() bool {
	for _, existing := range allowedNebiusCloudIntegrationAccountServiceAccountAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NebiusCloudIntegrationAccountServiceAccountAuthType value.
func (v NebiusCloudIntegrationAccountServiceAccountAuthType) Ptr() *NebiusCloudIntegrationAccountServiceAccountAuthType {
	return &v
}

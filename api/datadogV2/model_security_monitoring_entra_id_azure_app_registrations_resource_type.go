// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringEntraIdAzureAppRegistrationsResourceType The type of the resource. The value should always be `entra_id_azure_app_registrations`.
type SecurityMonitoringEntraIdAzureAppRegistrationsResourceType string

// List of SecurityMonitoringEntraIdAzureAppRegistrationsResourceType.
const (
	SECURITYMONITORINGENTRAIDAZUREAPPREGISTRATIONSRESOURCETYPE_ENTRA_ID_AZURE_APP_REGISTRATIONS SecurityMonitoringEntraIdAzureAppRegistrationsResourceType = "entra_id_azure_app_registrations"
)

var allowedSecurityMonitoringEntraIdAzureAppRegistrationsResourceTypeEnumValues = []SecurityMonitoringEntraIdAzureAppRegistrationsResourceType{
	SECURITYMONITORINGENTRAIDAZUREAPPREGISTRATIONSRESOURCETYPE_ENTRA_ID_AZURE_APP_REGISTRATIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringEntraIdAzureAppRegistrationsResourceType) GetAllowedValues() []SecurityMonitoringEntraIdAzureAppRegistrationsResourceType {
	return allowedSecurityMonitoringEntraIdAzureAppRegistrationsResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringEntraIdAzureAppRegistrationsResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringEntraIdAzureAppRegistrationsResourceType(value)
	return nil
}

// NewSecurityMonitoringEntraIdAzureAppRegistrationsResourceTypeFromValue returns a pointer to a valid SecurityMonitoringEntraIdAzureAppRegistrationsResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringEntraIdAzureAppRegistrationsResourceTypeFromValue(v string) (*SecurityMonitoringEntraIdAzureAppRegistrationsResourceType, error) {
	ev := SecurityMonitoringEntraIdAzureAppRegistrationsResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringEntraIdAzureAppRegistrationsResourceType: valid values are %v", v, allowedSecurityMonitoringEntraIdAzureAppRegistrationsResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringEntraIdAzureAppRegistrationsResourceType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringEntraIdAzureAppRegistrationsResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringEntraIdAzureAppRegistrationsResourceType value.
func (v SecurityMonitoringEntraIdAzureAppRegistrationsResourceType) Ptr() *SecurityMonitoringEntraIdAzureAppRegistrationsResourceType {
	return &v
}

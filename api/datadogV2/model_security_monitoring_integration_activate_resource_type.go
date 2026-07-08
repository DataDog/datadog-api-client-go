// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringIntegrationActivateResourceType The type of the resource. The value should always be `activate_entra_id_request`.
type SecurityMonitoringIntegrationActivateResourceType string

// List of SecurityMonitoringIntegrationActivateResourceType.
const (
	SECURITYMONITORINGINTEGRATIONACTIVATERESOURCETYPE_ACTIVATE_ENTRA_ID_REQUEST SecurityMonitoringIntegrationActivateResourceType = "activate_entra_id_request"
)

var allowedSecurityMonitoringIntegrationActivateResourceTypeEnumValues = []SecurityMonitoringIntegrationActivateResourceType{
	SECURITYMONITORINGINTEGRATIONACTIVATERESOURCETYPE_ACTIVATE_ENTRA_ID_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringIntegrationActivateResourceType) GetAllowedValues() []SecurityMonitoringIntegrationActivateResourceType {
	return allowedSecurityMonitoringIntegrationActivateResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringIntegrationActivateResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringIntegrationActivateResourceType(value)
	return nil
}

// NewSecurityMonitoringIntegrationActivateResourceTypeFromValue returns a pointer to a valid SecurityMonitoringIntegrationActivateResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringIntegrationActivateResourceTypeFromValue(v string) (*SecurityMonitoringIntegrationActivateResourceType, error) {
	ev := SecurityMonitoringIntegrationActivateResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringIntegrationActivateResourceType: valid values are %v", v, allowedSecurityMonitoringIntegrationActivateResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringIntegrationActivateResourceType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringIntegrationActivateResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringIntegrationActivateResourceType value.
func (v SecurityMonitoringIntegrationActivateResourceType) Ptr() *SecurityMonitoringIntegrationActivateResourceType {
	return &v
}

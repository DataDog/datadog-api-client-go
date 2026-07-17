// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringIntegrationTypeEntraId The source type for an Entra ID entity context sync.
type SecurityMonitoringIntegrationTypeEntraId string

// List of SecurityMonitoringIntegrationTypeEntraId.
const (
	SECURITYMONITORINGINTEGRATIONTYPEENTRAID_ENTRA_ID SecurityMonitoringIntegrationTypeEntraId = "ENTRA_ID"
)

var allowedSecurityMonitoringIntegrationTypeEntraIdEnumValues = []SecurityMonitoringIntegrationTypeEntraId{
	SECURITYMONITORINGINTEGRATIONTYPEENTRAID_ENTRA_ID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringIntegrationTypeEntraId) GetAllowedValues() []SecurityMonitoringIntegrationTypeEntraId {
	return allowedSecurityMonitoringIntegrationTypeEntraIdEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringIntegrationTypeEntraId) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringIntegrationTypeEntraId(value)
	return nil
}

// NewSecurityMonitoringIntegrationTypeEntraIdFromValue returns a pointer to a valid SecurityMonitoringIntegrationTypeEntraId
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringIntegrationTypeEntraIdFromValue(v string) (*SecurityMonitoringIntegrationTypeEntraId, error) {
	ev := SecurityMonitoringIntegrationTypeEntraId(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringIntegrationTypeEntraId: valid values are %v", v, allowedSecurityMonitoringIntegrationTypeEntraIdEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringIntegrationTypeEntraId) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringIntegrationTypeEntraIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringIntegrationTypeEntraId value.
func (v SecurityMonitoringIntegrationTypeEntraId) Ptr() *SecurityMonitoringIntegrationTypeEntraId {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringIntegrationTypeGoogleWorkspace The source type for a Google Workspace entity context sync.
type SecurityMonitoringIntegrationTypeGoogleWorkspace string

// List of SecurityMonitoringIntegrationTypeGoogleWorkspace.
const (
	SECURITYMONITORINGINTEGRATIONTYPEGOOGLEWORKSPACE_GOOGLE_WORKSPACE SecurityMonitoringIntegrationTypeGoogleWorkspace = "GOOGLE_WORKSPACE"
)

var allowedSecurityMonitoringIntegrationTypeGoogleWorkspaceEnumValues = []SecurityMonitoringIntegrationTypeGoogleWorkspace{
	SECURITYMONITORINGINTEGRATIONTYPEGOOGLEWORKSPACE_GOOGLE_WORKSPACE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringIntegrationTypeGoogleWorkspace) GetAllowedValues() []SecurityMonitoringIntegrationTypeGoogleWorkspace {
	return allowedSecurityMonitoringIntegrationTypeGoogleWorkspaceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringIntegrationTypeGoogleWorkspace) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringIntegrationTypeGoogleWorkspace(value)
	return nil
}

// NewSecurityMonitoringIntegrationTypeGoogleWorkspaceFromValue returns a pointer to a valid SecurityMonitoringIntegrationTypeGoogleWorkspace
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringIntegrationTypeGoogleWorkspaceFromValue(v string) (*SecurityMonitoringIntegrationTypeGoogleWorkspace, error) {
	ev := SecurityMonitoringIntegrationTypeGoogleWorkspace(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringIntegrationTypeGoogleWorkspace: valid values are %v", v, allowedSecurityMonitoringIntegrationTypeGoogleWorkspaceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringIntegrationTypeGoogleWorkspace) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringIntegrationTypeGoogleWorkspaceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringIntegrationTypeGoogleWorkspace value.
func (v SecurityMonitoringIntegrationTypeGoogleWorkspace) Ptr() *SecurityMonitoringIntegrationTypeGoogleWorkspace {
	return &v
}

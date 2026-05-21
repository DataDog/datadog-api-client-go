// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentJiraIssueIntegrationType Incident integration resource type.
type IncidentJiraIssueIntegrationType string

// List of IncidentJiraIssueIntegrationType.
const (
	INCIDENTJIRAISSUEINTEGRATIONTYPE_INCIDENT_INTEGRATIONS IncidentJiraIssueIntegrationType = "incident_integrations"
)

var allowedIncidentJiraIssueIntegrationTypeEnumValues = []IncidentJiraIssueIntegrationType{
	INCIDENTJIRAISSUEINTEGRATIONTYPE_INCIDENT_INTEGRATIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentJiraIssueIntegrationType) GetAllowedValues() []IncidentJiraIssueIntegrationType {
	return allowedIncidentJiraIssueIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentJiraIssueIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentJiraIssueIntegrationType(value)
	return nil
}

// NewIncidentJiraIssueIntegrationTypeFromValue returns a pointer to a valid IncidentJiraIssueIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentJiraIssueIntegrationTypeFromValue(v string) (*IncidentJiraIssueIntegrationType, error) {
	ev := IncidentJiraIssueIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentJiraIssueIntegrationType: valid values are %v", v, allowedIncidentJiraIssueIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentJiraIssueIntegrationType) IsValid() bool {
	for _, existing := range allowedIncidentJiraIssueIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentJiraIssueIntegrationType value.
func (v IncidentJiraIssueIntegrationType) Ptr() *IncidentJiraIssueIntegrationType {
	return &v
}

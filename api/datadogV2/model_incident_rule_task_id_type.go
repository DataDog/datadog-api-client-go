// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentRuleTaskIDType The task ID for an incident rule.
type IncidentRuleTaskIDType string

// List of IncidentRuleTaskIDType.
const (
	INCIDENTRULETASKIDTYPE_JIRA_CREATE_ISSUE_JOB                  IncidentRuleTaskIDType = "jira-create-issue-job"
	INCIDENTRULETASKIDTYPE_NOTIFY_INCIDENT_HANDLES_JOB            IncidentRuleTaskIDType = "notify-incident-handles-job"
	INCIDENTRULETASKIDTYPE_SERVICENOW_CREATE_INCIDENT_JOB         IncidentRuleTaskIDType = "servicenow-create-incident-job"
	INCIDENTRULETASKIDTYPE_SLACK_CREATE_CHANNEL_JOB               IncidentRuleTaskIDType = "slack-create-channel-job"
	INCIDENTRULETASKIDTYPE_ZOOM_CREATE_MEETING_JOB                IncidentRuleTaskIDType = "zoom-create-meeting-job"
	INCIDENTRULETASKIDTYPE_GOOGLE_MEET_CREATE_MEETING_JOB         IncidentRuleTaskIDType = "google-meet-create-meeting-job"
	INCIDENTRULETASKIDTYPE_WORKFLOW_AUTOMATION_JOB                IncidentRuleTaskIDType = "workflow-automation-job"
	INCIDENTRULETASKIDTYPE_MS_TEAMS_CREATE_MEETING_JOB            IncidentRuleTaskIDType = "ms-teams-create-meeting-job"
	INCIDENTRULETASKIDTYPE_GOOGLE_CHAT_CREATE_SPACE_JOB           IncidentRuleTaskIDType = "google-chat-create-space-job"
	INCIDENTRULETASKIDTYPE_ZOOM_SUPPRESS_SUMMARIZATION_JOB        IncidentRuleTaskIDType = "zoom-suppress-summarization-job"
	INCIDENTRULETASKIDTYPE_MS_TEAMS_SUPPRESS_SUMMARIZATION_JOB    IncidentRuleTaskIDType = "ms-teams-suppress-summarization-job"
	INCIDENTRULETASKIDTYPE_GOOGLE_MEET_SUPPRESS_SUMMARIZATION_JOB IncidentRuleTaskIDType = "google-meet-suppress-summarization-job"
)

var allowedIncidentRuleTaskIDTypeEnumValues = []IncidentRuleTaskIDType{
	INCIDENTRULETASKIDTYPE_JIRA_CREATE_ISSUE_JOB,
	INCIDENTRULETASKIDTYPE_NOTIFY_INCIDENT_HANDLES_JOB,
	INCIDENTRULETASKIDTYPE_SERVICENOW_CREATE_INCIDENT_JOB,
	INCIDENTRULETASKIDTYPE_SLACK_CREATE_CHANNEL_JOB,
	INCIDENTRULETASKIDTYPE_ZOOM_CREATE_MEETING_JOB,
	INCIDENTRULETASKIDTYPE_GOOGLE_MEET_CREATE_MEETING_JOB,
	INCIDENTRULETASKIDTYPE_WORKFLOW_AUTOMATION_JOB,
	INCIDENTRULETASKIDTYPE_MS_TEAMS_CREATE_MEETING_JOB,
	INCIDENTRULETASKIDTYPE_GOOGLE_CHAT_CREATE_SPACE_JOB,
	INCIDENTRULETASKIDTYPE_ZOOM_SUPPRESS_SUMMARIZATION_JOB,
	INCIDENTRULETASKIDTYPE_MS_TEAMS_SUPPRESS_SUMMARIZATION_JOB,
	INCIDENTRULETASKIDTYPE_GOOGLE_MEET_SUPPRESS_SUMMARIZATION_JOB,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentRuleTaskIDType) GetAllowedValues() []IncidentRuleTaskIDType {
	return allowedIncidentRuleTaskIDTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentRuleTaskIDType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentRuleTaskIDType(value)
	return nil
}

// NewIncidentRuleTaskIDTypeFromValue returns a pointer to a valid IncidentRuleTaskIDType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentRuleTaskIDTypeFromValue(v string) (*IncidentRuleTaskIDType, error) {
	ev := IncidentRuleTaskIDType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentRuleTaskIDType: valid values are %v", v, allowedIncidentRuleTaskIDTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentRuleTaskIDType) IsValid() bool {
	for _, existing := range allowedIncidentRuleTaskIDTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentRuleTaskIDType value.
func (v IncidentRuleTaskIDType) Ptr() *IncidentRuleTaskIDType {
	return &v
}

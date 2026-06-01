// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentAIPostmortemResponseType AI postmortem response resource type.
type IncidentAIPostmortemResponseType string

// List of IncidentAIPostmortemResponseType.
const (
	INCIDENTAIPOSTMORTEMRESPONSETYPE_GET_INCIDENT_AI_POSTMORTEM_RESPONSE IncidentAIPostmortemResponseType = "get_incident_ai_postmortem_response"
)

var allowedIncidentAIPostmortemResponseTypeEnumValues = []IncidentAIPostmortemResponseType{
	INCIDENTAIPOSTMORTEMRESPONSETYPE_GET_INCIDENT_AI_POSTMORTEM_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentAIPostmortemResponseType) GetAllowedValues() []IncidentAIPostmortemResponseType {
	return allowedIncidentAIPostmortemResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentAIPostmortemResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentAIPostmortemResponseType(value)
	return nil
}

// NewIncidentAIPostmortemResponseTypeFromValue returns a pointer to a valid IncidentAIPostmortemResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentAIPostmortemResponseTypeFromValue(v string) (*IncidentAIPostmortemResponseType, error) {
	ev := IncidentAIPostmortemResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentAIPostmortemResponseType: valid values are %v", v, allowedIncidentAIPostmortemResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentAIPostmortemResponseType) IsValid() bool {
	for _, existing := range allowedIncidentAIPostmortemResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentAIPostmortemResponseType value.
func (v IncidentAIPostmortemResponseType) Ptr() *IncidentAIPostmortemResponseType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentAutomationDataType Incident automation data resource type.
type IncidentAutomationDataType string

// List of IncidentAutomationDataType.
const (
	INCIDENTAUTOMATIONDATATYPE_INCIDENTS_AUTOMATION_DATA IncidentAutomationDataType = "incidents_automation_data"
)

var allowedIncidentAutomationDataTypeEnumValues = []IncidentAutomationDataType{
	INCIDENTAUTOMATIONDATATYPE_INCIDENTS_AUTOMATION_DATA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentAutomationDataType) GetAllowedValues() []IncidentAutomationDataType {
	return allowedIncidentAutomationDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentAutomationDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentAutomationDataType(value)
	return nil
}

// NewIncidentAutomationDataTypeFromValue returns a pointer to a valid IncidentAutomationDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentAutomationDataTypeFromValue(v string) (*IncidentAutomationDataType, error) {
	ev := IncidentAutomationDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentAutomationDataType: valid values are %v", v, allowedIncidentAutomationDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentAutomationDataType) IsValid() bool {
	for _, existing := range allowedIncidentAutomationDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentAutomationDataType value.
func (v IncidentAutomationDataType) Ptr() *IncidentAutomationDataType {
	return &v
}

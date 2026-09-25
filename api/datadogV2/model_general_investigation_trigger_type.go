// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeneralInvestigationTriggerType The type of general investigation trigger.
type GeneralInvestigationTriggerType string

// List of GeneralInvestigationTriggerType.
const (
	GENERALINVESTIGATIONTRIGGERTYPE_GENERAL_INVESTIGATION GeneralInvestigationTriggerType = "general_investigation"
)

var allowedGeneralInvestigationTriggerTypeEnumValues = []GeneralInvestigationTriggerType{
	GENERALINVESTIGATIONTRIGGERTYPE_GENERAL_INVESTIGATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GeneralInvestigationTriggerType) GetAllowedValues() []GeneralInvestigationTriggerType {
	return allowedGeneralInvestigationTriggerTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GeneralInvestigationTriggerType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GeneralInvestigationTriggerType(value)
	return nil
}

// NewGeneralInvestigationTriggerTypeFromValue returns a pointer to a valid GeneralInvestigationTriggerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGeneralInvestigationTriggerTypeFromValue(v string) (*GeneralInvestigationTriggerType, error) {
	ev := GeneralInvestigationTriggerType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GeneralInvestigationTriggerType: valid values are %v", v, allowedGeneralInvestigationTriggerTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GeneralInvestigationTriggerType) IsValid() bool {
	for _, existing := range allowedGeneralInvestigationTriggerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GeneralInvestigationTriggerType value.
func (v GeneralInvestigationTriggerType) Ptr() *GeneralInvestigationTriggerType {
	return &v
}

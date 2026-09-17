// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DemRumStepType The type of a RUM journey step.
type DemRumStepType string

// List of DemRumStepType.
const (
	DEMRUMSTEPTYPE_START DemRumStepType = "start"
	DEMRUMSTEPTYPE_STOP  DemRumStepType = "stop"
	DEMRUMSTEPTYPE_STEP  DemRumStepType = "step"
)

var allowedDemRumStepTypeEnumValues = []DemRumStepType{
	DEMRUMSTEPTYPE_START,
	DEMRUMSTEPTYPE_STOP,
	DEMRUMSTEPTYPE_STEP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DemRumStepType) GetAllowedValues() []DemRumStepType {
	return allowedDemRumStepTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DemRumStepType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DemRumStepType(value)
	return nil
}

// NewDemRumStepTypeFromValue returns a pointer to a valid DemRumStepType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDemRumStepTypeFromValue(v string) (*DemRumStepType, error) {
	ev := DemRumStepType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DemRumStepType: valid values are %v", v, allowedDemRumStepTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DemRumStepType) IsValid() bool {
	for _, existing := range allowedDemRumStepTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DemRumStepType value.
func (v DemRumStepType) Ptr() *DemRumStepType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMOperationJourneyStepType The type of a step within a RUM operation's journey.
type RUMOperationJourneyStepType string

// List of RUMOperationJourneyStepType.
const (
	RUMOPERATIONJOURNEYSTEPTYPE_START     RUMOperationJourneyStepType = "start"
	RUMOPERATIONJOURNEYSTEPTYPE_UPDATE    RUMOperationJourneyStepType = "update"
	RUMOPERATIONJOURNEYSTEPTYPE_STOP      RUMOperationJourneyStepType = "stop"
	RUMOPERATIONJOURNEYSTEPTYPE_ERROR     RUMOperationJourneyStepType = "error"
	RUMOPERATIONJOURNEYSTEPTYPE_ABANDONED RUMOperationJourneyStepType = "abandoned"
)

var allowedRUMOperationJourneyStepTypeEnumValues = []RUMOperationJourneyStepType{
	RUMOPERATIONJOURNEYSTEPTYPE_START,
	RUMOPERATIONJOURNEYSTEPTYPE_UPDATE,
	RUMOPERATIONJOURNEYSTEPTYPE_STOP,
	RUMOPERATIONJOURNEYSTEPTYPE_ERROR,
	RUMOPERATIONJOURNEYSTEPTYPE_ABANDONED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMOperationJourneyStepType) GetAllowedValues() []RUMOperationJourneyStepType {
	return allowedRUMOperationJourneyStepTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMOperationJourneyStepType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMOperationJourneyStepType(value)
	return nil
}

// NewRUMOperationJourneyStepTypeFromValue returns a pointer to a valid RUMOperationJourneyStepType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMOperationJourneyStepTypeFromValue(v string) (*RUMOperationJourneyStepType, error) {
	ev := RUMOperationJourneyStepType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMOperationJourneyStepType: valid values are %v", v, allowedRUMOperationJourneyStepTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMOperationJourneyStepType) IsValid() bool {
	for _, existing := range allowedRUMOperationJourneyStepTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMOperationJourneyStepType value.
func (v RUMOperationJourneyStepType) Ptr() *RUMOperationJourneyStepType {
	return &v
}

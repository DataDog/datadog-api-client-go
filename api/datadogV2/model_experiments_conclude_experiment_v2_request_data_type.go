// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsConcludeExperimentV2RequestDataType Conclude experiment request resource type.
type ExperimentsConcludeExperimentV2RequestDataType string

// List of ExperimentsConcludeExperimentV2RequestDataType.
const (
	EXPERIMENTSCONCLUDEEXPERIMENTV2REQUESTDATATYPE_CONCLUDE_EXPERIMENT_REQUEST ExperimentsConcludeExperimentV2RequestDataType = "conclude-experiment-request"
)

var allowedExperimentsConcludeExperimentV2RequestDataTypeEnumValues = []ExperimentsConcludeExperimentV2RequestDataType{
	EXPERIMENTSCONCLUDEEXPERIMENTV2REQUESTDATATYPE_CONCLUDE_EXPERIMENT_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsConcludeExperimentV2RequestDataType) GetAllowedValues() []ExperimentsConcludeExperimentV2RequestDataType {
	return allowedExperimentsConcludeExperimentV2RequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsConcludeExperimentV2RequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsConcludeExperimentV2RequestDataType(value)
	return nil
}

// NewExperimentsConcludeExperimentV2RequestDataTypeFromValue returns a pointer to a valid ExperimentsConcludeExperimentV2RequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsConcludeExperimentV2RequestDataTypeFromValue(v string) (*ExperimentsConcludeExperimentV2RequestDataType, error) {
	ev := ExperimentsConcludeExperimentV2RequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsConcludeExperimentV2RequestDataType: valid values are %v", v, allowedExperimentsConcludeExperimentV2RequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsConcludeExperimentV2RequestDataType) IsValid() bool {
	for _, existing := range allowedExperimentsConcludeExperimentV2RequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsConcludeExperimentV2RequestDataType value.
func (v ExperimentsConcludeExperimentV2RequestDataType) Ptr() *ExperimentsConcludeExperimentV2RequestDataType {
	return &v
}

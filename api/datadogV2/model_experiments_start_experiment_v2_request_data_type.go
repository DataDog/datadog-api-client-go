// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsStartExperimentV2RequestDataType Start experiment request resource type.
type ExperimentsStartExperimentV2RequestDataType string

// List of ExperimentsStartExperimentV2RequestDataType.
const (
	EXPERIMENTSSTARTEXPERIMENTV2REQUESTDATATYPE_START_EXPERIMENT_REQUEST ExperimentsStartExperimentV2RequestDataType = "start-experiment-request"
)

var allowedExperimentsStartExperimentV2RequestDataTypeEnumValues = []ExperimentsStartExperimentV2RequestDataType{
	EXPERIMENTSSTARTEXPERIMENTV2REQUESTDATATYPE_START_EXPERIMENT_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsStartExperimentV2RequestDataType) GetAllowedValues() []ExperimentsStartExperimentV2RequestDataType {
	return allowedExperimentsStartExperimentV2RequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsStartExperimentV2RequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsStartExperimentV2RequestDataType(value)
	return nil
}

// NewExperimentsStartExperimentV2RequestDataTypeFromValue returns a pointer to a valid ExperimentsStartExperimentV2RequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsStartExperimentV2RequestDataTypeFromValue(v string) (*ExperimentsStartExperimentV2RequestDataType, error) {
	ev := ExperimentsStartExperimentV2RequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsStartExperimentV2RequestDataType: valid values are %v", v, allowedExperimentsStartExperimentV2RequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsStartExperimentV2RequestDataType) IsValid() bool {
	for _, existing := range allowedExperimentsStartExperimentV2RequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsStartExperimentV2RequestDataType value.
func (v ExperimentsStartExperimentV2RequestDataType) Ptr() *ExperimentsStartExperimentV2RequestDataType {
	return &v
}

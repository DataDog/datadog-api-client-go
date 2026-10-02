// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCancelExperimentV2RequestDataType Cancel experiment request resource type.
type ExperimentsCancelExperimentV2RequestDataType string

// List of ExperimentsCancelExperimentV2RequestDataType.
const (
	EXPERIMENTSCANCELEXPERIMENTV2REQUESTDATATYPE_CANCEL_EXPERIMENT_REQUEST ExperimentsCancelExperimentV2RequestDataType = "cancel-experiment-request"
)

var allowedExperimentsCancelExperimentV2RequestDataTypeEnumValues = []ExperimentsCancelExperimentV2RequestDataType{
	EXPERIMENTSCANCELEXPERIMENTV2REQUESTDATATYPE_CANCEL_EXPERIMENT_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsCancelExperimentV2RequestDataType) GetAllowedValues() []ExperimentsCancelExperimentV2RequestDataType {
	return allowedExperimentsCancelExperimentV2RequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsCancelExperimentV2RequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsCancelExperimentV2RequestDataType(value)
	return nil
}

// NewExperimentsCancelExperimentV2RequestDataTypeFromValue returns a pointer to a valid ExperimentsCancelExperimentV2RequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsCancelExperimentV2RequestDataTypeFromValue(v string) (*ExperimentsCancelExperimentV2RequestDataType, error) {
	ev := ExperimentsCancelExperimentV2RequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsCancelExperimentV2RequestDataType: valid values are %v", v, allowedExperimentsCancelExperimentV2RequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsCancelExperimentV2RequestDataType) IsValid() bool {
	for _, existing := range allowedExperimentsCancelExperimentV2RequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsCancelExperimentV2RequestDataType value.
func (v ExperimentsCancelExperimentV2RequestDataType) Ptr() *ExperimentsCancelExperimentV2RequestDataType {
	return &v
}

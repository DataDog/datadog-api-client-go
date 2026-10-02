// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPatchExperimentMetricGroupV2RequestDataType Experiment metric groups resource type.
type ExperimentsPatchExperimentMetricGroupV2RequestDataType string

// List of ExperimentsPatchExperimentMetricGroupV2RequestDataType.
const (
	EXPERIMENTSPATCHEXPERIMENTMETRICGROUPV2REQUESTDATATYPE_EXPERIMENT_METRIC_GROUPS ExperimentsPatchExperimentMetricGroupV2RequestDataType = "experiment-metric-groups"
)

var allowedExperimentsPatchExperimentMetricGroupV2RequestDataTypeEnumValues = []ExperimentsPatchExperimentMetricGroupV2RequestDataType{
	EXPERIMENTSPATCHEXPERIMENTMETRICGROUPV2REQUESTDATATYPE_EXPERIMENT_METRIC_GROUPS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsPatchExperimentMetricGroupV2RequestDataType) GetAllowedValues() []ExperimentsPatchExperimentMetricGroupV2RequestDataType {
	return allowedExperimentsPatchExperimentMetricGroupV2RequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsPatchExperimentMetricGroupV2RequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsPatchExperimentMetricGroupV2RequestDataType(value)
	return nil
}

// NewExperimentsPatchExperimentMetricGroupV2RequestDataTypeFromValue returns a pointer to a valid ExperimentsPatchExperimentMetricGroupV2RequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsPatchExperimentMetricGroupV2RequestDataTypeFromValue(v string) (*ExperimentsPatchExperimentMetricGroupV2RequestDataType, error) {
	ev := ExperimentsPatchExperimentMetricGroupV2RequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsPatchExperimentMetricGroupV2RequestDataType: valid values are %v", v, allowedExperimentsPatchExperimentMetricGroupV2RequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsPatchExperimentMetricGroupV2RequestDataType) IsValid() bool {
	for _, existing := range allowedExperimentsPatchExperimentMetricGroupV2RequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsPatchExperimentMetricGroupV2RequestDataType value.
func (v ExperimentsPatchExperimentMetricGroupV2RequestDataType) Ptr() *ExperimentsPatchExperimentMetricGroupV2RequestDataType {
	return &v
}

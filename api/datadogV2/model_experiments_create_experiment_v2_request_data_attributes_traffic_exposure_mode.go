// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode Whether exposure uses a fixed fraction or a sequence of steps.
type ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode string

// List of ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode.
const (
	EXPERIMENTSCREATEEXPERIMENTV2REQUESTDATAATTRIBUTESTRAFFICEXPOSUREMODE_STATIC ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode = "STATIC"
	EXPERIMENTSCREATEEXPERIMENTV2REQUESTDATAATTRIBUTESTRAFFICEXPOSUREMODE_STEPS  ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode = "STEPS"
)

var allowedExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureModeEnumValues = []ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode{
	EXPERIMENTSCREATEEXPERIMENTV2REQUESTDATAATTRIBUTESTRAFFICEXPOSUREMODE_STATIC,
	EXPERIMENTSCREATEEXPERIMENTV2REQUESTDATAATTRIBUTESTRAFFICEXPOSUREMODE_STEPS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode) GetAllowedValues() []ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode {
	return allowedExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode(value)
	return nil
}

// NewExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureModeFromValue returns a pointer to a valid ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureModeFromValue(v string) (*ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode, error) {
	ev := ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode: valid values are %v", v, allowedExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode) IsValid() bool {
	for _, existing := range allowedExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode value.
func (v ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode) Ptr() *ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode {
	return &v
}

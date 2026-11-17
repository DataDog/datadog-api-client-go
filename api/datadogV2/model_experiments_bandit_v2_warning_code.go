// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditV2WarningCode Code that identifies the warning. `datadog_flag_configuration_omitted` means the caller lacks feature flag or environment read permissions, so the response leaves out `datadog_flag_configuration`.
type ExperimentsBanditV2WarningCode string

// List of ExperimentsBanditV2WarningCode.
const (
	EXPERIMENTSBANDITV2WARNINGCODE_DATADOG_FLAG_CONFIGURATION_OMITTED ExperimentsBanditV2WarningCode = "datadog_flag_configuration_omitted"
)

var allowedExperimentsBanditV2WarningCodeEnumValues = []ExperimentsBanditV2WarningCode{
	EXPERIMENTSBANDITV2WARNINGCODE_DATADOG_FLAG_CONFIGURATION_OMITTED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsBanditV2WarningCode) GetAllowedValues() []ExperimentsBanditV2WarningCode {
	return allowedExperimentsBanditV2WarningCodeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsBanditV2WarningCode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsBanditV2WarningCode(value)
	return nil
}

// NewExperimentsBanditV2WarningCodeFromValue returns a pointer to a valid ExperimentsBanditV2WarningCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsBanditV2WarningCodeFromValue(v string) (*ExperimentsBanditV2WarningCode, error) {
	ev := ExperimentsBanditV2WarningCode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsBanditV2WarningCode: valid values are %v", v, allowedExperimentsBanditV2WarningCodeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsBanditV2WarningCode) IsValid() bool {
	for _, existing := range allowedExperimentsBanditV2WarningCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsBanditV2WarningCode value.
func (v ExperimentsBanditV2WarningCode) Ptr() *ExperimentsBanditV2WarningCode {
	return &v
}

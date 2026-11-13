// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditPolicyReason Why the policy was applied.
type ExperimentsBanditPolicyReason string

// List of ExperimentsBanditPolicyReason.
const (
	EXPERIMENTSBANDITPOLICYREASON_ACTIVATION    ExperimentsBanditPolicyReason = "ACTIVATION"
	EXPERIMENTSBANDITPOLICYREASON_REFRESH       ExperimentsBanditPolicyReason = "REFRESH"
	EXPERIMENTSBANDITPOLICYREASON_ARM_CHANGE    ExperimentsBanditPolicyReason = "ARM_CHANGE"
	EXPERIMENTSBANDITPOLICYREASON_SCHEDULED_END ExperimentsBanditPolicyReason = "SCHEDULED_END"
	EXPERIMENTSBANDITPOLICYREASON_STOP          ExperimentsBanditPolicyReason = "STOP"
)

var allowedExperimentsBanditPolicyReasonEnumValues = []ExperimentsBanditPolicyReason{
	EXPERIMENTSBANDITPOLICYREASON_ACTIVATION,
	EXPERIMENTSBANDITPOLICYREASON_REFRESH,
	EXPERIMENTSBANDITPOLICYREASON_ARM_CHANGE,
	EXPERIMENTSBANDITPOLICYREASON_SCHEDULED_END,
	EXPERIMENTSBANDITPOLICYREASON_STOP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsBanditPolicyReason) GetAllowedValues() []ExperimentsBanditPolicyReason {
	return allowedExperimentsBanditPolicyReasonEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsBanditPolicyReason) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsBanditPolicyReason(value)
	return nil
}

// NewExperimentsBanditPolicyReasonFromValue returns a pointer to a valid ExperimentsBanditPolicyReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsBanditPolicyReasonFromValue(v string) (*ExperimentsBanditPolicyReason, error) {
	ev := ExperimentsBanditPolicyReason(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsBanditPolicyReason: valid values are %v", v, allowedExperimentsBanditPolicyReasonEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsBanditPolicyReason) IsValid() bool {
	for _, existing := range allowedExperimentsBanditPolicyReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsBanditPolicyReason value.
func (v ExperimentsBanditPolicyReason) Ptr() *ExperimentsBanditPolicyReason {
	return &v
}

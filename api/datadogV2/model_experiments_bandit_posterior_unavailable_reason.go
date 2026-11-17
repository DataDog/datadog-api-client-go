// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditPosteriorUnavailableReason Why the arm's estimate and interval are unavailable. Absent when they are available.
type ExperimentsBanditPosteriorUnavailableReason string

// List of ExperimentsBanditPosteriorUnavailableReason.
const (
	EXPERIMENTSBANDITPOSTERIORUNAVAILABLEREASON_NO_DATA           ExperimentsBanditPosteriorUnavailableReason = "no_data"
	EXPERIMENTSBANDITPOSTERIORUNAVAILABLEREASON_INSUFFICIENT_DATA ExperimentsBanditPosteriorUnavailableReason = "insufficient_data"
	EXPERIMENTSBANDITPOSTERIORUNAVAILABLEREASON_ZERO_VARIANCE     ExperimentsBanditPosteriorUnavailableReason = "zero_variance"
)

var allowedExperimentsBanditPosteriorUnavailableReasonEnumValues = []ExperimentsBanditPosteriorUnavailableReason{
	EXPERIMENTSBANDITPOSTERIORUNAVAILABLEREASON_NO_DATA,
	EXPERIMENTSBANDITPOSTERIORUNAVAILABLEREASON_INSUFFICIENT_DATA,
	EXPERIMENTSBANDITPOSTERIORUNAVAILABLEREASON_ZERO_VARIANCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsBanditPosteriorUnavailableReason) GetAllowedValues() []ExperimentsBanditPosteriorUnavailableReason {
	return allowedExperimentsBanditPosteriorUnavailableReasonEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsBanditPosteriorUnavailableReason) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsBanditPosteriorUnavailableReason(value)
	return nil
}

// NewExperimentsBanditPosteriorUnavailableReasonFromValue returns a pointer to a valid ExperimentsBanditPosteriorUnavailableReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsBanditPosteriorUnavailableReasonFromValue(v string) (*ExperimentsBanditPosteriorUnavailableReason, error) {
	ev := ExperimentsBanditPosteriorUnavailableReason(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsBanditPosteriorUnavailableReason: valid values are %v", v, allowedExperimentsBanditPosteriorUnavailableReasonEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsBanditPosteriorUnavailableReason) IsValid() bool {
	for _, existing := range allowedExperimentsBanditPosteriorUnavailableReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsBanditPosteriorUnavailableReason value.
func (v ExperimentsBanditPosteriorUnavailableReason) Ptr() *ExperimentsBanditPosteriorUnavailableReason {
	return &v
}

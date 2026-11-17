// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditOptimalVariantUnavailableReason Why no optimal variant is reported. Absent when one is, or when the cause isn't known.
type ExperimentsBanditOptimalVariantUnavailableReason string

// List of ExperimentsBanditOptimalVariantUnavailableReason.
const (
	EXPERIMENTSBANDITOPTIMALVARIANTUNAVAILABLEREASON_NO_RESULTS_YET    ExperimentsBanditOptimalVariantUnavailableReason = "NO_RESULTS_YET"
	EXPERIMENTSBANDITOPTIMALVARIANTUNAVAILABLEREASON_WARMUP            ExperimentsBanditOptimalVariantUnavailableReason = "WARMUP"
	EXPERIMENTSBANDITOPTIMALVARIANTUNAVAILABLEREASON_NOT_ENOUGH_DATA   ExperimentsBanditOptimalVariantUnavailableReason = "NOT_ENOUGH_DATA"
	EXPERIMENTSBANDITOPTIMALVARIANTUNAVAILABLEREASON_NO_DIFFERENCE     ExperimentsBanditOptimalVariantUnavailableReason = "NO_DIFFERENCE"
	EXPERIMENTSBANDITOPTIMALVARIANTUNAVAILABLEREASON_STATS_UNAVAILABLE ExperimentsBanditOptimalVariantUnavailableReason = "STATS_UNAVAILABLE"
)

var allowedExperimentsBanditOptimalVariantUnavailableReasonEnumValues = []ExperimentsBanditOptimalVariantUnavailableReason{
	EXPERIMENTSBANDITOPTIMALVARIANTUNAVAILABLEREASON_NO_RESULTS_YET,
	EXPERIMENTSBANDITOPTIMALVARIANTUNAVAILABLEREASON_WARMUP,
	EXPERIMENTSBANDITOPTIMALVARIANTUNAVAILABLEREASON_NOT_ENOUGH_DATA,
	EXPERIMENTSBANDITOPTIMALVARIANTUNAVAILABLEREASON_NO_DIFFERENCE,
	EXPERIMENTSBANDITOPTIMALVARIANTUNAVAILABLEREASON_STATS_UNAVAILABLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsBanditOptimalVariantUnavailableReason) GetAllowedValues() []ExperimentsBanditOptimalVariantUnavailableReason {
	return allowedExperimentsBanditOptimalVariantUnavailableReasonEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsBanditOptimalVariantUnavailableReason) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsBanditOptimalVariantUnavailableReason(value)
	return nil
}

// NewExperimentsBanditOptimalVariantUnavailableReasonFromValue returns a pointer to a valid ExperimentsBanditOptimalVariantUnavailableReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsBanditOptimalVariantUnavailableReasonFromValue(v string) (*ExperimentsBanditOptimalVariantUnavailableReason, error) {
	ev := ExperimentsBanditOptimalVariantUnavailableReason(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsBanditOptimalVariantUnavailableReason: valid values are %v", v, allowedExperimentsBanditOptimalVariantUnavailableReasonEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsBanditOptimalVariantUnavailableReason) IsValid() bool {
	for _, existing := range allowedExperimentsBanditOptimalVariantUnavailableReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsBanditOptimalVariantUnavailableReason value.
func (v ExperimentsBanditOptimalVariantUnavailableReason) Ptr() *ExperimentsBanditOptimalVariantUnavailableReason {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineTagCardinalityLimitProcessorTrackingModeMode The cardinality tracking algorithm to use.
type ObservabilityPipelineTagCardinalityLimitProcessorTrackingModeMode string

// List of ObservabilityPipelineTagCardinalityLimitProcessorTrackingModeMode.
const (
	OBSERVABILITYPIPELINETAGCARDINALITYLIMITPROCESSORTRACKINGMODEMODE_EXACT_FINGERPRINT ObservabilityPipelineTagCardinalityLimitProcessorTrackingModeMode = "exact_fingerprint"
	OBSERVABILITYPIPELINETAGCARDINALITYLIMITPROCESSORTRACKINGMODEMODE_PROBABILISTIC     ObservabilityPipelineTagCardinalityLimitProcessorTrackingModeMode = "probabilistic"
)

var allowedObservabilityPipelineTagCardinalityLimitProcessorTrackingModeModeEnumValues = []ObservabilityPipelineTagCardinalityLimitProcessorTrackingModeMode{
	OBSERVABILITYPIPELINETAGCARDINALITYLIMITPROCESSORTRACKINGMODEMODE_EXACT_FINGERPRINT,
	OBSERVABILITYPIPELINETAGCARDINALITYLIMITPROCESSORTRACKINGMODEMODE_PROBABILISTIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineTagCardinalityLimitProcessorTrackingModeMode) GetAllowedValues() []ObservabilityPipelineTagCardinalityLimitProcessorTrackingModeMode {
	return allowedObservabilityPipelineTagCardinalityLimitProcessorTrackingModeModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineTagCardinalityLimitProcessorTrackingModeMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineTagCardinalityLimitProcessorTrackingModeMode(value)
	return nil
}

// NewObservabilityPipelineTagCardinalityLimitProcessorTrackingModeModeFromValue returns a pointer to a valid ObservabilityPipelineTagCardinalityLimitProcessorTrackingModeMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineTagCardinalityLimitProcessorTrackingModeModeFromValue(v string) (*ObservabilityPipelineTagCardinalityLimitProcessorTrackingModeMode, error) {
	ev := ObservabilityPipelineTagCardinalityLimitProcessorTrackingModeMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineTagCardinalityLimitProcessorTrackingModeMode: valid values are %v", v, allowedObservabilityPipelineTagCardinalityLimitProcessorTrackingModeModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineTagCardinalityLimitProcessorTrackingModeMode) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineTagCardinalityLimitProcessorTrackingModeModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineTagCardinalityLimitProcessorTrackingModeMode value.
func (v ObservabilityPipelineTagCardinalityLimitProcessorTrackingModeMode) Ptr() *ObservabilityPipelineTagCardinalityLimitProcessorTrackingModeMode {
	return &v
}

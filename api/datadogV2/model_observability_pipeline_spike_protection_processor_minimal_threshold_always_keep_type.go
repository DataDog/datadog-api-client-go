// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepType The minimal threshold type. The value should always be `always_keep`.
type ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepType string

// List of ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepType.
const (
	OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORMINIMALTHRESHOLDALWAYSKEEPTYPE_ALWAYS_KEEP ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepType = "always_keep"
)

var allowedObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepTypeEnumValues = []ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepType{
	OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORMINIMALTHRESHOLDALWAYSKEEPTYPE_ALWAYS_KEEP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepType) GetAllowedValues() []ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepType {
	return allowedObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepType(value)
	return nil
}

// NewObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepTypeFromValue returns a pointer to a valid ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepTypeFromValue(v string) (*ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepType, error) {
	ev := ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepType: valid values are %v", v, allowedObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepType value.
func (v ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepType) Ptr() *ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeepType {
	return &v
}

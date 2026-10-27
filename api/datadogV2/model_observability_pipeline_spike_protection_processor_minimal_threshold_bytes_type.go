// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType The minimal threshold type. The value should always be `threshold`.
type ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType string

// List of ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType.
const (
	OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORMINIMALTHRESHOLDBYTESTYPE_THRESHOLD ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType = "threshold"
)

var allowedObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesTypeEnumValues = []ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType{
	OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORMINIMALTHRESHOLDBYTESTYPE_THRESHOLD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType) GetAllowedValues() []ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType {
	return allowedObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType(value)
	return nil
}

// NewObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesTypeFromValue returns a pointer to a valid ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesTypeFromValue(v string) (*ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType, error) {
	ev := ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType: valid values are %v", v, allowedObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType value.
func (v ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType) Ptr() *ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytesType {
	return &v
}

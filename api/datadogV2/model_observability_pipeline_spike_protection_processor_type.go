// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSpikeProtectionProcessorType The processor type. The value should always be `spike_protection`.
type ObservabilityPipelineSpikeProtectionProcessorType string

// List of ObservabilityPipelineSpikeProtectionProcessorType.
const (
	OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORTYPE_SPIKE_PROTECTION ObservabilityPipelineSpikeProtectionProcessorType = "spike_protection"
)

var allowedObservabilityPipelineSpikeProtectionProcessorTypeEnumValues = []ObservabilityPipelineSpikeProtectionProcessorType{
	OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORTYPE_SPIKE_PROTECTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSpikeProtectionProcessorType) GetAllowedValues() []ObservabilityPipelineSpikeProtectionProcessorType {
	return allowedObservabilityPipelineSpikeProtectionProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSpikeProtectionProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSpikeProtectionProcessorType(value)
	return nil
}

// NewObservabilityPipelineSpikeProtectionProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineSpikeProtectionProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSpikeProtectionProcessorTypeFromValue(v string) (*ObservabilityPipelineSpikeProtectionProcessorType, error) {
	ev := ObservabilityPipelineSpikeProtectionProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSpikeProtectionProcessorType: valid values are %v", v, allowedObservabilityPipelineSpikeProtectionProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSpikeProtectionProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSpikeProtectionProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSpikeProtectionProcessorType value.
func (v ObservabilityPipelineSpikeProtectionProcessorType) Ptr() *ObservabilityPipelineSpikeProtectionProcessorType {
	return &v
}

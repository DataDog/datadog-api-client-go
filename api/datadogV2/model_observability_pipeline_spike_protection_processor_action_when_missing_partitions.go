// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions How to handle logs that are missing one of the `partition_fields`. Options:
// - `ignore`: Forward the event without applying spike protection.
// - `default_bucket`: Assign the event to a default partition.
type ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions string

// List of ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions.
const (
	OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORACTIONWHENMISSINGPARTITIONS_IGNORE         ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions = "ignore"
	OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORACTIONWHENMISSINGPARTITIONS_DEFAULT_BUCKET ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions = "default_bucket"
)

var allowedObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitionsEnumValues = []ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions{
	OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORACTIONWHENMISSINGPARTITIONS_IGNORE,
	OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORACTIONWHENMISSINGPARTITIONS_DEFAULT_BUCKET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions) GetAllowedValues() []ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions {
	return allowedObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitionsEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions(value)
	return nil
}

// NewObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitionsFromValue returns a pointer to a valid ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitionsFromValue(v string) (*ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions, error) {
	ev := ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions: valid values are %v", v, allowedObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitionsEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions value.
func (v ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions) Ptr() *ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions {
	return &v
}

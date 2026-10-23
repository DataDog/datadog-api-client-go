// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSpikeProtectionProcessorSpikeAction The action to take with matching logs while spike protection is active. Options:
// - `drop`: Drop the event.
// - `no_action`: Let the event pass through.
// - `overflow_routing`: Route to an overflow destination.
//
// When `overflow_routing` is used, there must be a destination whose `inputs` reference this processor with the `<processor-id>.overflow_events` suffix to route the overflowing events.
type ObservabilityPipelineSpikeProtectionProcessorSpikeAction string

// List of ObservabilityPipelineSpikeProtectionProcessorSpikeAction.
const (
	OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORSPIKEACTION_DROP             ObservabilityPipelineSpikeProtectionProcessorSpikeAction = "drop"
	OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORSPIKEACTION_NO_ACTION        ObservabilityPipelineSpikeProtectionProcessorSpikeAction = "no_action"
	OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORSPIKEACTION_OVERFLOW_ROUTING ObservabilityPipelineSpikeProtectionProcessorSpikeAction = "overflow_routing"
)

var allowedObservabilityPipelineSpikeProtectionProcessorSpikeActionEnumValues = []ObservabilityPipelineSpikeProtectionProcessorSpikeAction{
	OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORSPIKEACTION_DROP,
	OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORSPIKEACTION_NO_ACTION,
	OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORSPIKEACTION_OVERFLOW_ROUTING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSpikeProtectionProcessorSpikeAction) GetAllowedValues() []ObservabilityPipelineSpikeProtectionProcessorSpikeAction {
	return allowedObservabilityPipelineSpikeProtectionProcessorSpikeActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSpikeProtectionProcessorSpikeAction) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSpikeProtectionProcessorSpikeAction(value)
	return nil
}

// NewObservabilityPipelineSpikeProtectionProcessorSpikeActionFromValue returns a pointer to a valid ObservabilityPipelineSpikeProtectionProcessorSpikeAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSpikeProtectionProcessorSpikeActionFromValue(v string) (*ObservabilityPipelineSpikeProtectionProcessorSpikeAction, error) {
	ev := ObservabilityPipelineSpikeProtectionProcessorSpikeAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSpikeProtectionProcessorSpikeAction: valid values are %v", v, allowedObservabilityPipelineSpikeProtectionProcessorSpikeActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSpikeProtectionProcessorSpikeAction) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSpikeProtectionProcessorSpikeActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSpikeProtectionProcessorSpikeAction value.
func (v ObservabilityPipelineSpikeProtectionProcessorSpikeAction) Ptr() *ObservabilityPipelineSpikeProtectionProcessorSpikeAction {
	return &v
}

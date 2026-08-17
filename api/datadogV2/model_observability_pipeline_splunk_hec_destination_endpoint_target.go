// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSplunkHecDestinationEndpointTarget The Splunk HEC endpoint to send events to. Use `event` to send structured events to the `/event` endpoint, or `raw` to send the raw message to the `/raw` endpoint.
type ObservabilityPipelineSplunkHecDestinationEndpointTarget string

// List of ObservabilityPipelineSplunkHecDestinationEndpointTarget.
const (
	OBSERVABILITYPIPELINESPLUNKHECDESTINATIONENDPOINTTARGET_EVENT ObservabilityPipelineSplunkHecDestinationEndpointTarget = "event"
	OBSERVABILITYPIPELINESPLUNKHECDESTINATIONENDPOINTTARGET_RAW   ObservabilityPipelineSplunkHecDestinationEndpointTarget = "raw"
)

var allowedObservabilityPipelineSplunkHecDestinationEndpointTargetEnumValues = []ObservabilityPipelineSplunkHecDestinationEndpointTarget{
	OBSERVABILITYPIPELINESPLUNKHECDESTINATIONENDPOINTTARGET_EVENT,
	OBSERVABILITYPIPELINESPLUNKHECDESTINATIONENDPOINTTARGET_RAW,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSplunkHecDestinationEndpointTarget) GetAllowedValues() []ObservabilityPipelineSplunkHecDestinationEndpointTarget {
	return allowedObservabilityPipelineSplunkHecDestinationEndpointTargetEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSplunkHecDestinationEndpointTarget) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSplunkHecDestinationEndpointTarget(value)
	return nil
}

// NewObservabilityPipelineSplunkHecDestinationEndpointTargetFromValue returns a pointer to a valid ObservabilityPipelineSplunkHecDestinationEndpointTarget
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSplunkHecDestinationEndpointTargetFromValue(v string) (*ObservabilityPipelineSplunkHecDestinationEndpointTarget, error) {
	ev := ObservabilityPipelineSplunkHecDestinationEndpointTarget(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSplunkHecDestinationEndpointTarget: valid values are %v", v, allowedObservabilityPipelineSplunkHecDestinationEndpointTargetEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSplunkHecDestinationEndpointTarget) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSplunkHecDestinationEndpointTargetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSplunkHecDestinationEndpointTarget value.
func (v ObservabilityPipelineSplunkHecDestinationEndpointTarget) Ptr() *ObservabilityPipelineSplunkHecDestinationEndpointTarget {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAggregateProcessorAggregationTimingType Determines whether metrics are assigned to aggregation windows based on when they are processed or their timestamps.
type ObservabilityPipelineAggregateProcessorAggregationTimingType string

// List of ObservabilityPipelineAggregateProcessorAggregationTimingType.
const (
	OBSERVABILITYPIPELINEAGGREGATEPROCESSORAGGREGATIONTIMINGTYPE_SYSTEM_TIME ObservabilityPipelineAggregateProcessorAggregationTimingType = "system_time"
	OBSERVABILITYPIPELINEAGGREGATEPROCESSORAGGREGATIONTIMINGTYPE_EVENT_TIME  ObservabilityPipelineAggregateProcessorAggregationTimingType = "event_time"
)

var allowedObservabilityPipelineAggregateProcessorAggregationTimingTypeEnumValues = []ObservabilityPipelineAggregateProcessorAggregationTimingType{
	OBSERVABILITYPIPELINEAGGREGATEPROCESSORAGGREGATIONTIMINGTYPE_SYSTEM_TIME,
	OBSERVABILITYPIPELINEAGGREGATEPROCESSORAGGREGATIONTIMINGTYPE_EVENT_TIME,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAggregateProcessorAggregationTimingType) GetAllowedValues() []ObservabilityPipelineAggregateProcessorAggregationTimingType {
	return allowedObservabilityPipelineAggregateProcessorAggregationTimingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAggregateProcessorAggregationTimingType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAggregateProcessorAggregationTimingType(value)
	return nil
}

// NewObservabilityPipelineAggregateProcessorAggregationTimingTypeFromValue returns a pointer to a valid ObservabilityPipelineAggregateProcessorAggregationTimingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAggregateProcessorAggregationTimingTypeFromValue(v string) (*ObservabilityPipelineAggregateProcessorAggregationTimingType, error) {
	ev := ObservabilityPipelineAggregateProcessorAggregationTimingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAggregateProcessorAggregationTimingType: valid values are %v", v, allowedObservabilityPipelineAggregateProcessorAggregationTimingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAggregateProcessorAggregationTimingType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAggregateProcessorAggregationTimingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAggregateProcessorAggregationTimingType value.
func (v ObservabilityPipelineAggregateProcessorAggregationTimingType) Ptr() *ObservabilityPipelineAggregateProcessorAggregationTimingType {
	return &v
}

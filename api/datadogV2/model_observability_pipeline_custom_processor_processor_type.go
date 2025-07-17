// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineCustomProcessorProcessorType The processor type. The value should always be `custom_processor`.
type ObservabilityPipelineCustomProcessorProcessorType string

// List of ObservabilityPipelineCustomProcessorProcessorType.
const (
	OBSERVABILITYPIPELINECUSTOMPROCESSORPROCESSORTYPE_CUSTOM_PROCESSOR ObservabilityPipelineCustomProcessorProcessorType = "custom_processor"
)

var allowedObservabilityPipelineCustomProcessorProcessorTypeEnumValues = []ObservabilityPipelineCustomProcessorProcessorType{
	OBSERVABILITYPIPELINECUSTOMPROCESSORPROCESSORTYPE_CUSTOM_PROCESSOR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineCustomProcessorProcessorType) GetAllowedValues() []ObservabilityPipelineCustomProcessorProcessorType {
	return allowedObservabilityPipelineCustomProcessorProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineCustomProcessorProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineCustomProcessorProcessorType(value)
	return nil
}

// NewObservabilityPipelineCustomProcessorProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineCustomProcessorProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineCustomProcessorProcessorTypeFromValue(v string) (*ObservabilityPipelineCustomProcessorProcessorType, error) {
	ev := ObservabilityPipelineCustomProcessorProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineCustomProcessorProcessorType: valid values are %v", v, allowedObservabilityPipelineCustomProcessorProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineCustomProcessorProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineCustomProcessorProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineCustomProcessorProcessorType value.
func (v ObservabilityPipelineCustomProcessorProcessorType) Ptr() *ObservabilityPipelineCustomProcessorProcessorType {
	return &v
}

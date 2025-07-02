// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineRemapVrlProcessorType The processor type. The value should always be `remap_vrl`.
type ObservabilityPipelineRemapVrlProcessorType string

// List of ObservabilityPipelineRemapVrlProcessorType.
const (
	OBSERVABILITYPIPELINEREMAPVRLPROCESSORTYPE_REMAP_VRL ObservabilityPipelineRemapVrlProcessorType = "remap_vrl"
)

var allowedObservabilityPipelineRemapVrlProcessorTypeEnumValues = []ObservabilityPipelineRemapVrlProcessorType{
	OBSERVABILITYPIPELINEREMAPVRLPROCESSORTYPE_REMAP_VRL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineRemapVrlProcessorType) GetAllowedValues() []ObservabilityPipelineRemapVrlProcessorType {
	return allowedObservabilityPipelineRemapVrlProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineRemapVrlProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineRemapVrlProcessorType(value)
	return nil
}

// NewObservabilityPipelineRemapVrlProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineRemapVrlProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineRemapVrlProcessorTypeFromValue(v string) (*ObservabilityPipelineRemapVrlProcessorType, error) {
	ev := ObservabilityPipelineRemapVrlProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineRemapVrlProcessorType: valid values are %v", v, allowedObservabilityPipelineRemapVrlProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineRemapVrlProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineRemapVrlProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineRemapVrlProcessorType value.
func (v ObservabilityPipelineRemapVrlProcessorType) Ptr() *ObservabilityPipelineRemapVrlProcessorType {
	return &v
}

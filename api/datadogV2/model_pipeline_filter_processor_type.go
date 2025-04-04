// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PipelineFilterProcessorType The processor type. The value should always be `filter`.
type PipelineFilterProcessorType string

// List of PipelineFilterProcessorType.
const (
	PIPELINEFILTERPROCESSORTYPE_FILTER PipelineFilterProcessorType = "filter"
)

var allowedPipelineFilterProcessorTypeEnumValues = []PipelineFilterProcessorType{
	PIPELINEFILTERPROCESSORTYPE_FILTER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PipelineFilterProcessorType) GetAllowedValues() []PipelineFilterProcessorType {
	return allowedPipelineFilterProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PipelineFilterProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PipelineFilterProcessorType(value)
	return nil
}

// NewPipelineFilterProcessorTypeFromValue returns a pointer to a valid PipelineFilterProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPipelineFilterProcessorTypeFromValue(v string) (*PipelineFilterProcessorType, error) {
	ev := PipelineFilterProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PipelineFilterProcessorType: valid values are %v", v, allowedPipelineFilterProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PipelineFilterProcessorType) IsValid() bool {
	for _, existing := range allowedPipelineFilterProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PipelineFilterProcessorType value.
func (v PipelineFilterProcessorType) Ptr() *PipelineFilterProcessorType {
	return &v
}

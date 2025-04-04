// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PipelineParseJSONProcessorType The processor type. The value should always be `parse_json`.
type PipelineParseJSONProcessorType string

// List of PipelineParseJSONProcessorType.
const (
	PIPELINEPARSEJSONPROCESSORTYPE_PARSE_JSON PipelineParseJSONProcessorType = "parse_json"
)

var allowedPipelineParseJSONProcessorTypeEnumValues = []PipelineParseJSONProcessorType{
	PIPELINEPARSEJSONPROCESSORTYPE_PARSE_JSON,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PipelineParseJSONProcessorType) GetAllowedValues() []PipelineParseJSONProcessorType {
	return allowedPipelineParseJSONProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PipelineParseJSONProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PipelineParseJSONProcessorType(value)
	return nil
}

// NewPipelineParseJSONProcessorTypeFromValue returns a pointer to a valid PipelineParseJSONProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPipelineParseJSONProcessorTypeFromValue(v string) (*PipelineParseJSONProcessorType, error) {
	ev := PipelineParseJSONProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PipelineParseJSONProcessorType: valid values are %v", v, allowedPipelineParseJSONProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PipelineParseJSONProcessorType) IsValid() bool {
	for _, existing := range allowedPipelineParseJSONProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PipelineParseJSONProcessorType value.
func (v PipelineParseJSONProcessorType) Ptr() *PipelineParseJSONProcessorType {
	return &v
}

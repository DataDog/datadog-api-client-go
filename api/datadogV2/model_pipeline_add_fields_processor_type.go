// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PipelineAddFieldsProcessorType The processor type. The value should always be `add_fields`.
type PipelineAddFieldsProcessorType string

// List of PipelineAddFieldsProcessorType.
const (
	PIPELINEADDFIELDSPROCESSORTYPE_ADD_FIELDS PipelineAddFieldsProcessorType = "add_fields"
)

var allowedPipelineAddFieldsProcessorTypeEnumValues = []PipelineAddFieldsProcessorType{
	PIPELINEADDFIELDSPROCESSORTYPE_ADD_FIELDS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PipelineAddFieldsProcessorType) GetAllowedValues() []PipelineAddFieldsProcessorType {
	return allowedPipelineAddFieldsProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PipelineAddFieldsProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PipelineAddFieldsProcessorType(value)
	return nil
}

// NewPipelineAddFieldsProcessorTypeFromValue returns a pointer to a valid PipelineAddFieldsProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPipelineAddFieldsProcessorTypeFromValue(v string) (*PipelineAddFieldsProcessorType, error) {
	ev := PipelineAddFieldsProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PipelineAddFieldsProcessorType: valid values are %v", v, allowedPipelineAddFieldsProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PipelineAddFieldsProcessorType) IsValid() bool {
	for _, existing := range allowedPipelineAddFieldsProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PipelineAddFieldsProcessorType value.
func (v PipelineAddFieldsProcessorType) Ptr() *PipelineAddFieldsProcessorType {
	return &v
}

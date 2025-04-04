// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PipelineRenameFieldsProcessorType The processor type. The value should always be `rename_fields`.
type PipelineRenameFieldsProcessorType string

// List of PipelineRenameFieldsProcessorType.
const (
	PIPELINERENAMEFIELDSPROCESSORTYPE_RENAME_FIELDS PipelineRenameFieldsProcessorType = "rename_fields"
)

var allowedPipelineRenameFieldsProcessorTypeEnumValues = []PipelineRenameFieldsProcessorType{
	PIPELINERENAMEFIELDSPROCESSORTYPE_RENAME_FIELDS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PipelineRenameFieldsProcessorType) GetAllowedValues() []PipelineRenameFieldsProcessorType {
	return allowedPipelineRenameFieldsProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PipelineRenameFieldsProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PipelineRenameFieldsProcessorType(value)
	return nil
}

// NewPipelineRenameFieldsProcessorTypeFromValue returns a pointer to a valid PipelineRenameFieldsProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPipelineRenameFieldsProcessorTypeFromValue(v string) (*PipelineRenameFieldsProcessorType, error) {
	ev := PipelineRenameFieldsProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PipelineRenameFieldsProcessorType: valid values are %v", v, allowedPipelineRenameFieldsProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PipelineRenameFieldsProcessorType) IsValid() bool {
	for _, existing := range allowedPipelineRenameFieldsProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PipelineRenameFieldsProcessorType value.
func (v PipelineRenameFieldsProcessorType) Ptr() *PipelineRenameFieldsProcessorType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PipelineQuotaProcessorType The processor type. The value should always be `quota`.
type PipelineQuotaProcessorType string

// List of PipelineQuotaProcessorType.
const (
	PIPELINEQUOTAPROCESSORTYPE_QUOTA PipelineQuotaProcessorType = "quota"
)

var allowedPipelineQuotaProcessorTypeEnumValues = []PipelineQuotaProcessorType{
	PIPELINEQUOTAPROCESSORTYPE_QUOTA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PipelineQuotaProcessorType) GetAllowedValues() []PipelineQuotaProcessorType {
	return allowedPipelineQuotaProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PipelineQuotaProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PipelineQuotaProcessorType(value)
	return nil
}

// NewPipelineQuotaProcessorTypeFromValue returns a pointer to a valid PipelineQuotaProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPipelineQuotaProcessorTypeFromValue(v string) (*PipelineQuotaProcessorType, error) {
	ev := PipelineQuotaProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PipelineQuotaProcessorType: valid values are %v", v, allowedPipelineQuotaProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PipelineQuotaProcessorType) IsValid() bool {
	for _, existing := range allowedPipelineQuotaProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PipelineQuotaProcessorType value.
func (v PipelineQuotaProcessorType) Ptr() *PipelineQuotaProcessorType {
	return &v
}

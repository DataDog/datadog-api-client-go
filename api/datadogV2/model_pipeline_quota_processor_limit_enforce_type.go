// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PipelineQuotaProcessorLimitEnforceType Unit for quota enforcement in bytes for data size or events for count.
type PipelineQuotaProcessorLimitEnforceType string

// List of PipelineQuotaProcessorLimitEnforceType.
const (
	PIPELINEQUOTAPROCESSORLIMITENFORCETYPE_BYTES  PipelineQuotaProcessorLimitEnforceType = "bytes"
	PIPELINEQUOTAPROCESSORLIMITENFORCETYPE_EVENTS PipelineQuotaProcessorLimitEnforceType = "events"
)

var allowedPipelineQuotaProcessorLimitEnforceTypeEnumValues = []PipelineQuotaProcessorLimitEnforceType{
	PIPELINEQUOTAPROCESSORLIMITENFORCETYPE_BYTES,
	PIPELINEQUOTAPROCESSORLIMITENFORCETYPE_EVENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PipelineQuotaProcessorLimitEnforceType) GetAllowedValues() []PipelineQuotaProcessorLimitEnforceType {
	return allowedPipelineQuotaProcessorLimitEnforceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PipelineQuotaProcessorLimitEnforceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PipelineQuotaProcessorLimitEnforceType(value)
	return nil
}

// NewPipelineQuotaProcessorLimitEnforceTypeFromValue returns a pointer to a valid PipelineQuotaProcessorLimitEnforceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPipelineQuotaProcessorLimitEnforceTypeFromValue(v string) (*PipelineQuotaProcessorLimitEnforceType, error) {
	ev := PipelineQuotaProcessorLimitEnforceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PipelineQuotaProcessorLimitEnforceType: valid values are %v", v, allowedPipelineQuotaProcessorLimitEnforceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PipelineQuotaProcessorLimitEnforceType) IsValid() bool {
	for _, existing := range allowedPipelineQuotaProcessorLimitEnforceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PipelineQuotaProcessorLimitEnforceType value.
func (v PipelineQuotaProcessorLimitEnforceType) Ptr() *PipelineQuotaProcessorLimitEnforceType {
	return &v
}

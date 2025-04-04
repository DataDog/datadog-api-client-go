// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PipelineDatadogAgentSourceType The source type. The value should always be `datadog_agent`.
type PipelineDatadogAgentSourceType string

// List of PipelineDatadogAgentSourceType.
const (
	PIPELINEDATADOGAGENTSOURCETYPE_DATADOG_AGENT PipelineDatadogAgentSourceType = "datadog_agent"
)

var allowedPipelineDatadogAgentSourceTypeEnumValues = []PipelineDatadogAgentSourceType{
	PIPELINEDATADOGAGENTSOURCETYPE_DATADOG_AGENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PipelineDatadogAgentSourceType) GetAllowedValues() []PipelineDatadogAgentSourceType {
	return allowedPipelineDatadogAgentSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PipelineDatadogAgentSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PipelineDatadogAgentSourceType(value)
	return nil
}

// NewPipelineDatadogAgentSourceTypeFromValue returns a pointer to a valid PipelineDatadogAgentSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPipelineDatadogAgentSourceTypeFromValue(v string) (*PipelineDatadogAgentSourceType, error) {
	ev := PipelineDatadogAgentSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PipelineDatadogAgentSourceType: valid values are %v", v, allowedPipelineDatadogAgentSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PipelineDatadogAgentSourceType) IsValid() bool {
	for _, existing := range allowedPipelineDatadogAgentSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PipelineDatadogAgentSourceType value.
func (v PipelineDatadogAgentSourceType) Ptr() *PipelineDatadogAgentSourceType {
	return &v
}

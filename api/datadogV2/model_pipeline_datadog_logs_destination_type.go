// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PipelineDatadogLogsDestinationType The destination type. The value should always be `datadog_logs`.
type PipelineDatadogLogsDestinationType string

// List of PipelineDatadogLogsDestinationType.
const (
	PIPELINEDATADOGLOGSDESTINATIONTYPE_DATADOG_LOGS PipelineDatadogLogsDestinationType = "datadog_logs"
)

var allowedPipelineDatadogLogsDestinationTypeEnumValues = []PipelineDatadogLogsDestinationType{
	PIPELINEDATADOGLOGSDESTINATIONTYPE_DATADOG_LOGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PipelineDatadogLogsDestinationType) GetAllowedValues() []PipelineDatadogLogsDestinationType {
	return allowedPipelineDatadogLogsDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PipelineDatadogLogsDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PipelineDatadogLogsDestinationType(value)
	return nil
}

// NewPipelineDatadogLogsDestinationTypeFromValue returns a pointer to a valid PipelineDatadogLogsDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPipelineDatadogLogsDestinationTypeFromValue(v string) (*PipelineDatadogLogsDestinationType, error) {
	ev := PipelineDatadogLogsDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PipelineDatadogLogsDestinationType: valid values are %v", v, allowedPipelineDatadogLogsDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PipelineDatadogLogsDestinationType) IsValid() bool {
	for _, existing := range allowedPipelineDatadogLogsDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PipelineDatadogLogsDestinationType value.
func (v PipelineDatadogLogsDestinationType) Ptr() *PipelineDatadogLogsDestinationType {
	return &v
}

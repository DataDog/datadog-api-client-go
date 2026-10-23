// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineDatadogTracesDestinationType The destination type. The value should always be `datadog_traces`.
type ObservabilityPipelineDatadogTracesDestinationType string

// List of ObservabilityPipelineDatadogTracesDestinationType.
const (
	OBSERVABILITYPIPELINEDATADOGTRACESDESTINATIONTYPE_DATADOG_TRACES ObservabilityPipelineDatadogTracesDestinationType = "datadog_traces"
)

var allowedObservabilityPipelineDatadogTracesDestinationTypeEnumValues = []ObservabilityPipelineDatadogTracesDestinationType{
	OBSERVABILITYPIPELINEDATADOGTRACESDESTINATIONTYPE_DATADOG_TRACES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineDatadogTracesDestinationType) GetAllowedValues() []ObservabilityPipelineDatadogTracesDestinationType {
	return allowedObservabilityPipelineDatadogTracesDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineDatadogTracesDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineDatadogTracesDestinationType(value)
	return nil
}

// NewObservabilityPipelineDatadogTracesDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineDatadogTracesDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineDatadogTracesDestinationTypeFromValue(v string) (*ObservabilityPipelineDatadogTracesDestinationType, error) {
	ev := ObservabilityPipelineDatadogTracesDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineDatadogTracesDestinationType: valid values are %v", v, allowedObservabilityPipelineDatadogTracesDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineDatadogTracesDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineDatadogTracesDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineDatadogTracesDestinationType value.
func (v ObservabilityPipelineDatadogTracesDestinationType) Ptr() *ObservabilityPipelineDatadogTracesDestinationType {
	return &v
}

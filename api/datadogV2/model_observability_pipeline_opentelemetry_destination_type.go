// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineOpentelemetryDestinationType The destination type. Always `opentelemetry`.
type ObservabilityPipelineOpentelemetryDestinationType string

// List of ObservabilityPipelineOpentelemetryDestinationType.
const (
	OBSERVABILITYPIPELINEOPENTELEMETRYDESTINATIONTYPE_OPENTELEMETRY ObservabilityPipelineOpentelemetryDestinationType = "opentelemetry"
)

var allowedObservabilityPipelineOpentelemetryDestinationTypeEnumValues = []ObservabilityPipelineOpentelemetryDestinationType{
	OBSERVABILITYPIPELINEOPENTELEMETRYDESTINATIONTYPE_OPENTELEMETRY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineOpentelemetryDestinationType) GetAllowedValues() []ObservabilityPipelineOpentelemetryDestinationType {
	return allowedObservabilityPipelineOpentelemetryDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineOpentelemetryDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineOpentelemetryDestinationType(value)
	return nil
}

// NewObservabilityPipelineOpentelemetryDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineOpentelemetryDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineOpentelemetryDestinationTypeFromValue(v string) (*ObservabilityPipelineOpentelemetryDestinationType, error) {
	ev := ObservabilityPipelineOpentelemetryDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineOpentelemetryDestinationType: valid values are %v", v, allowedObservabilityPipelineOpentelemetryDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineOpentelemetryDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineOpentelemetryDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineOpentelemetryDestinationType value.
func (v ObservabilityPipelineOpentelemetryDestinationType) Ptr() *ObservabilityPipelineOpentelemetryDestinationType {
	return &v
}

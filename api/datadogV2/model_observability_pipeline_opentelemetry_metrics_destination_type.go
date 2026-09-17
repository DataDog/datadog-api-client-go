// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineOpentelemetryMetricsDestinationType The destination type. Always `opentelemetry`.
type ObservabilityPipelineOpentelemetryMetricsDestinationType string

// List of ObservabilityPipelineOpentelemetryMetricsDestinationType.
const (
	OBSERVABILITYPIPELINEOPENTELEMETRYMETRICSDESTINATIONTYPE_OPENTELEMETRY ObservabilityPipelineOpentelemetryMetricsDestinationType = "opentelemetry"
)

var allowedObservabilityPipelineOpentelemetryMetricsDestinationTypeEnumValues = []ObservabilityPipelineOpentelemetryMetricsDestinationType{
	OBSERVABILITYPIPELINEOPENTELEMETRYMETRICSDESTINATIONTYPE_OPENTELEMETRY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineOpentelemetryMetricsDestinationType) GetAllowedValues() []ObservabilityPipelineOpentelemetryMetricsDestinationType {
	return allowedObservabilityPipelineOpentelemetryMetricsDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineOpentelemetryMetricsDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineOpentelemetryMetricsDestinationType(value)
	return nil
}

// NewObservabilityPipelineOpentelemetryMetricsDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineOpentelemetryMetricsDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineOpentelemetryMetricsDestinationTypeFromValue(v string) (*ObservabilityPipelineOpentelemetryMetricsDestinationType, error) {
	ev := ObservabilityPipelineOpentelemetryMetricsDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineOpentelemetryMetricsDestinationType: valid values are %v", v, allowedObservabilityPipelineOpentelemetryMetricsDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineOpentelemetryMetricsDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineOpentelemetryMetricsDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineOpentelemetryMetricsDestinationType value.
func (v ObservabilityPipelineOpentelemetryMetricsDestinationType) Ptr() *ObservabilityPipelineOpentelemetryMetricsDestinationType {
	return &v
}

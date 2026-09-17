// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelinePrometheusRemoteWriteDestinationType The destination type. The value should always be `prometheus_remote_write`.
type ObservabilityPipelinePrometheusRemoteWriteDestinationType string

// List of ObservabilityPipelinePrometheusRemoteWriteDestinationType.
const (
	OBSERVABILITYPIPELINEPROMETHEUSREMOTEWRITEDESTINATIONTYPE_PROMETHEUS_REMOTE_WRITE ObservabilityPipelinePrometheusRemoteWriteDestinationType = "prometheus_remote_write"
)

var allowedObservabilityPipelinePrometheusRemoteWriteDestinationTypeEnumValues = []ObservabilityPipelinePrometheusRemoteWriteDestinationType{
	OBSERVABILITYPIPELINEPROMETHEUSREMOTEWRITEDESTINATIONTYPE_PROMETHEUS_REMOTE_WRITE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelinePrometheusRemoteWriteDestinationType) GetAllowedValues() []ObservabilityPipelinePrometheusRemoteWriteDestinationType {
	return allowedObservabilityPipelinePrometheusRemoteWriteDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelinePrometheusRemoteWriteDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelinePrometheusRemoteWriteDestinationType(value)
	return nil
}

// NewObservabilityPipelinePrometheusRemoteWriteDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelinePrometheusRemoteWriteDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelinePrometheusRemoteWriteDestinationTypeFromValue(v string) (*ObservabilityPipelinePrometheusRemoteWriteDestinationType, error) {
	ev := ObservabilityPipelinePrometheusRemoteWriteDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelinePrometheusRemoteWriteDestinationType: valid values are %v", v, allowedObservabilityPipelinePrometheusRemoteWriteDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelinePrometheusRemoteWriteDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelinePrometheusRemoteWriteDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelinePrometheusRemoteWriteDestinationType value.
func (v ObservabilityPipelinePrometheusRemoteWriteDestinationType) Ptr() *ObservabilityPipelinePrometheusRemoteWriteDestinationType {
	return &v
}

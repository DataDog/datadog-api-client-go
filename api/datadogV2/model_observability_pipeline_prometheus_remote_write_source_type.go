// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelinePrometheusRemoteWriteSourceType The source type. The value should always be `prometheus_remote_write`.
type ObservabilityPipelinePrometheusRemoteWriteSourceType string

// List of ObservabilityPipelinePrometheusRemoteWriteSourceType.
const (
	OBSERVABILITYPIPELINEPROMETHEUSREMOTEWRITESOURCETYPE_PROMETHEUS_REMOTE_WRITE ObservabilityPipelinePrometheusRemoteWriteSourceType = "prometheus_remote_write"
)

var allowedObservabilityPipelinePrometheusRemoteWriteSourceTypeEnumValues = []ObservabilityPipelinePrometheusRemoteWriteSourceType{
	OBSERVABILITYPIPELINEPROMETHEUSREMOTEWRITESOURCETYPE_PROMETHEUS_REMOTE_WRITE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelinePrometheusRemoteWriteSourceType) GetAllowedValues() []ObservabilityPipelinePrometheusRemoteWriteSourceType {
	return allowedObservabilityPipelinePrometheusRemoteWriteSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelinePrometheusRemoteWriteSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelinePrometheusRemoteWriteSourceType(value)
	return nil
}

// NewObservabilityPipelinePrometheusRemoteWriteSourceTypeFromValue returns a pointer to a valid ObservabilityPipelinePrometheusRemoteWriteSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelinePrometheusRemoteWriteSourceTypeFromValue(v string) (*ObservabilityPipelinePrometheusRemoteWriteSourceType, error) {
	ev := ObservabilityPipelinePrometheusRemoteWriteSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelinePrometheusRemoteWriteSourceType: valid values are %v", v, allowedObservabilityPipelinePrometheusRemoteWriteSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelinePrometheusRemoteWriteSourceType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelinePrometheusRemoteWriteSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelinePrometheusRemoteWriteSourceType value.
func (v ObservabilityPipelinePrometheusRemoteWriteSourceType) Ptr() *ObservabilityPipelinePrometheusRemoteWriteSourceType {
	return &v
}

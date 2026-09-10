// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy HTTP authentication method.
type ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy string

// List of ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy.
const (
	OBSERVABILITYPIPELINEPROMETHEUSREMOTEWRITESOURCEAUTHSTRATEGY_NONE  ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy = "none"
	OBSERVABILITYPIPELINEPROMETHEUSREMOTEWRITESOURCEAUTHSTRATEGY_PLAIN ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy = "plain"
)

var allowedObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategyEnumValues = []ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy{
	OBSERVABILITYPIPELINEPROMETHEUSREMOTEWRITESOURCEAUTHSTRATEGY_NONE,
	OBSERVABILITYPIPELINEPROMETHEUSREMOTEWRITESOURCEAUTHSTRATEGY_PLAIN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy) GetAllowedValues() []ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy {
	return allowedObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy(value)
	return nil
}

// NewObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategyFromValue returns a pointer to a valid ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategyFromValue(v string) (*ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy, error) {
	ev := ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy: valid values are %v", v, allowedObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy) IsValid() bool {
	for _, existing := range allowedObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy value.
func (v ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy) Ptr() *ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy {
	return &v
}

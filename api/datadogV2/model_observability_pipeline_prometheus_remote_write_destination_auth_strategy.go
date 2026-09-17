// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy The authentication strategy to use for outgoing Prometheus Remote Write requests.
type ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy string

// List of ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy.
const (
	OBSERVABILITYPIPELINEPROMETHEUSREMOTEWRITEDESTINATIONAUTHSTRATEGY_NONE   ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy = "none"
	OBSERVABILITYPIPELINEPROMETHEUSREMOTEWRITEDESTINATIONAUTHSTRATEGY_BASIC  ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy = "basic"
	OBSERVABILITYPIPELINEPROMETHEUSREMOTEWRITEDESTINATIONAUTHSTRATEGY_BEARER ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy = "bearer"
)

var allowedObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategyEnumValues = []ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy{
	OBSERVABILITYPIPELINEPROMETHEUSREMOTEWRITEDESTINATIONAUTHSTRATEGY_NONE,
	OBSERVABILITYPIPELINEPROMETHEUSREMOTEWRITEDESTINATIONAUTHSTRATEGY_BASIC,
	OBSERVABILITYPIPELINEPROMETHEUSREMOTEWRITEDESTINATIONAUTHSTRATEGY_BEARER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy) GetAllowedValues() []ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy {
	return allowedObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy(value)
	return nil
}

// NewObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategyFromValue returns a pointer to a valid ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategyFromValue(v string) (*ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy, error) {
	ev := ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy: valid values are %v", v, allowedObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy) IsValid() bool {
	for _, existing := range allowedObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy value.
func (v ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy) Ptr() *ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy {
	return &v
}

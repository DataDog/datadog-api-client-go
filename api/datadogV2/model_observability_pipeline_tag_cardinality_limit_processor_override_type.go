// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineTagCardinalityLimitProcessorOverrideType How the override is applied. `limit_override` enforces a custom limit; `excluded` omits the metric or tag from cardinality tracking.
type ObservabilityPipelineTagCardinalityLimitProcessorOverrideType string

// List of ObservabilityPipelineTagCardinalityLimitProcessorOverrideType.
const (
	OBSERVABILITYPIPELINETAGCARDINALITYLIMITPROCESSOROVERRIDETYPE_LIMIT_OVERRIDE ObservabilityPipelineTagCardinalityLimitProcessorOverrideType = "limit_override"
	OBSERVABILITYPIPELINETAGCARDINALITYLIMITPROCESSOROVERRIDETYPE_EXCLUDED       ObservabilityPipelineTagCardinalityLimitProcessorOverrideType = "excluded"
)

var allowedObservabilityPipelineTagCardinalityLimitProcessorOverrideTypeEnumValues = []ObservabilityPipelineTagCardinalityLimitProcessorOverrideType{
	OBSERVABILITYPIPELINETAGCARDINALITYLIMITPROCESSOROVERRIDETYPE_LIMIT_OVERRIDE,
	OBSERVABILITYPIPELINETAGCARDINALITYLIMITPROCESSOROVERRIDETYPE_EXCLUDED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineTagCardinalityLimitProcessorOverrideType) GetAllowedValues() []ObservabilityPipelineTagCardinalityLimitProcessorOverrideType {
	return allowedObservabilityPipelineTagCardinalityLimitProcessorOverrideTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineTagCardinalityLimitProcessorOverrideType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineTagCardinalityLimitProcessorOverrideType(value)
	return nil
}

// NewObservabilityPipelineTagCardinalityLimitProcessorOverrideTypeFromValue returns a pointer to a valid ObservabilityPipelineTagCardinalityLimitProcessorOverrideType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineTagCardinalityLimitProcessorOverrideTypeFromValue(v string) (*ObservabilityPipelineTagCardinalityLimitProcessorOverrideType, error) {
	ev := ObservabilityPipelineTagCardinalityLimitProcessorOverrideType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineTagCardinalityLimitProcessorOverrideType: valid values are %v", v, allowedObservabilityPipelineTagCardinalityLimitProcessorOverrideTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineTagCardinalityLimitProcessorOverrideType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineTagCardinalityLimitProcessorOverrideTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineTagCardinalityLimitProcessorOverrideType value.
func (v ObservabilityPipelineTagCardinalityLimitProcessorOverrideType) Ptr() *ObservabilityPipelineTagCardinalityLimitProcessorOverrideType {
	return &v
}

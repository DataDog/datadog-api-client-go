// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineMetricEnrichmentTableTagLookupType The lookup source type. The value should always be `tag`.
type ObservabilityPipelineMetricEnrichmentTableTagLookupType string

// List of ObservabilityPipelineMetricEnrichmentTableTagLookupType.
const (
	OBSERVABILITYPIPELINEMETRICENRICHMENTTABLETAGLOOKUPTYPE_TAG ObservabilityPipelineMetricEnrichmentTableTagLookupType = "tag"
)

var allowedObservabilityPipelineMetricEnrichmentTableTagLookupTypeEnumValues = []ObservabilityPipelineMetricEnrichmentTableTagLookupType{
	OBSERVABILITYPIPELINEMETRICENRICHMENTTABLETAGLOOKUPTYPE_TAG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineMetricEnrichmentTableTagLookupType) GetAllowedValues() []ObservabilityPipelineMetricEnrichmentTableTagLookupType {
	return allowedObservabilityPipelineMetricEnrichmentTableTagLookupTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineMetricEnrichmentTableTagLookupType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineMetricEnrichmentTableTagLookupType(value)
	return nil
}

// NewObservabilityPipelineMetricEnrichmentTableTagLookupTypeFromValue returns a pointer to a valid ObservabilityPipelineMetricEnrichmentTableTagLookupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineMetricEnrichmentTableTagLookupTypeFromValue(v string) (*ObservabilityPipelineMetricEnrichmentTableTagLookupType, error) {
	ev := ObservabilityPipelineMetricEnrichmentTableTagLookupType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineMetricEnrichmentTableTagLookupType: valid values are %v", v, allowedObservabilityPipelineMetricEnrichmentTableTagLookupTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineMetricEnrichmentTableTagLookupType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineMetricEnrichmentTableTagLookupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineMetricEnrichmentTableTagLookupType value.
func (v ObservabilityPipelineMetricEnrichmentTableTagLookupType) Ptr() *ObservabilityPipelineMetricEnrichmentTableTagLookupType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineMetricEnrichmentTableMetricNameLookupType The lookup source type. The value should always be `metric_name`.
type ObservabilityPipelineMetricEnrichmentTableMetricNameLookupType string

// List of ObservabilityPipelineMetricEnrichmentTableMetricNameLookupType.
const (
	OBSERVABILITYPIPELINEMETRICENRICHMENTTABLEMETRICNAMELOOKUPTYPE_METRIC_NAME ObservabilityPipelineMetricEnrichmentTableMetricNameLookupType = "metric_name"
)

var allowedObservabilityPipelineMetricEnrichmentTableMetricNameLookupTypeEnumValues = []ObservabilityPipelineMetricEnrichmentTableMetricNameLookupType{
	OBSERVABILITYPIPELINEMETRICENRICHMENTTABLEMETRICNAMELOOKUPTYPE_METRIC_NAME,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineMetricEnrichmentTableMetricNameLookupType) GetAllowedValues() []ObservabilityPipelineMetricEnrichmentTableMetricNameLookupType {
	return allowedObservabilityPipelineMetricEnrichmentTableMetricNameLookupTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineMetricEnrichmentTableMetricNameLookupType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineMetricEnrichmentTableMetricNameLookupType(value)
	return nil
}

// NewObservabilityPipelineMetricEnrichmentTableMetricNameLookupTypeFromValue returns a pointer to a valid ObservabilityPipelineMetricEnrichmentTableMetricNameLookupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineMetricEnrichmentTableMetricNameLookupTypeFromValue(v string) (*ObservabilityPipelineMetricEnrichmentTableMetricNameLookupType, error) {
	ev := ObservabilityPipelineMetricEnrichmentTableMetricNameLookupType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineMetricEnrichmentTableMetricNameLookupType: valid values are %v", v, allowedObservabilityPipelineMetricEnrichmentTableMetricNameLookupTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineMetricEnrichmentTableMetricNameLookupType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineMetricEnrichmentTableMetricNameLookupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineMetricEnrichmentTableMetricNameLookupType value.
func (v ObservabilityPipelineMetricEnrichmentTableMetricNameLookupType) Ptr() *ObservabilityPipelineMetricEnrichmentTableMetricNameLookupType {
	return &v
}

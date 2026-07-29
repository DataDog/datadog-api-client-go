// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HistoricalMetricsConfigurationType The historical metrics configuration resource type.
type HistoricalMetricsConfigurationType string

// List of HistoricalMetricsConfigurationType.
const (
	HISTORICALMETRICSCONFIGURATIONTYPE_HISTORICAL_METRICS_CONFIGURATIONS HistoricalMetricsConfigurationType = "historical_metrics_configurations"
)

var allowedHistoricalMetricsConfigurationTypeEnumValues = []HistoricalMetricsConfigurationType{
	HISTORICALMETRICSCONFIGURATIONTYPE_HISTORICAL_METRICS_CONFIGURATIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HistoricalMetricsConfigurationType) GetAllowedValues() []HistoricalMetricsConfigurationType {
	return allowedHistoricalMetricsConfigurationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HistoricalMetricsConfigurationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HistoricalMetricsConfigurationType(value)
	return nil
}

// NewHistoricalMetricsConfigurationTypeFromValue returns a pointer to a valid HistoricalMetricsConfigurationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHistoricalMetricsConfigurationTypeFromValue(v string) (*HistoricalMetricsConfigurationType, error) {
	ev := HistoricalMetricsConfigurationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HistoricalMetricsConfigurationType: valid values are %v", v, allowedHistoricalMetricsConfigurationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HistoricalMetricsConfigurationType) IsValid() bool {
	for _, existing := range allowedHistoricalMetricsConfigurationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HistoricalMetricsConfigurationType value.
func (v HistoricalMetricsConfigurationType) Ptr() *HistoricalMetricsConfigurationType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricMetadataType The metric metadata resource type.
type MetricMetadataType string

// List of MetricMetadataType.
const (
	METRICMETADATATYPE_METRIC_METADATA MetricMetadataType = "metric_metadata"
)

var allowedMetricMetadataTypeEnumValues = []MetricMetadataType{
	METRICMETADATATYPE_METRIC_METADATA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MetricMetadataType) GetAllowedValues() []MetricMetadataType {
	return allowedMetricMetadataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MetricMetadataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricMetadataType(value)
	return nil
}

// NewMetricMetadataTypeFromValue returns a pointer to a valid MetricMetadataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMetricMetadataTypeFromValue(v string) (*MetricMetadataType, error) {
	ev := MetricMetadataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MetricMetadataType: valid values are %v", v, allowedMetricMetadataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MetricMetadataType) IsValid() bool {
	for _, existing := range allowedMetricMetadataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricMetadataType value.
func (v MetricMetadataType) Ptr() *MetricMetadataType {
	return &v
}

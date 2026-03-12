// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricVolumesInclude Comma-separated list of additional data to include in the response. Allowed values are `metric_volumes`.
type MetricVolumesInclude string

// List of MetricVolumesInclude.
const (
	METRICVOLUMESINCLUDE_METRIC_VOLUMES              MetricVolumesInclude = "metric_volumes"
	METRICVOLUMESINCLUDE_GENERATED_METRIC_ATTRIBUTES MetricVolumesInclude = "generated_metric_attributes"
)

var allowedMetricVolumesIncludeEnumValues = []MetricVolumesInclude{
	METRICVOLUMESINCLUDE_METRIC_VOLUMES,
	METRICVOLUMESINCLUDE_GENERATED_METRIC_ATTRIBUTES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MetricVolumesInclude) GetAllowedValues() []MetricVolumesInclude {
	return allowedMetricVolumesIncludeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MetricVolumesInclude) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricVolumesInclude(value)
	return nil
}

// NewMetricVolumesIncludeFromValue returns a pointer to a valid MetricVolumesInclude
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMetricVolumesIncludeFromValue(v string) (*MetricVolumesInclude, error) {
	ev := MetricVolumesInclude(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MetricVolumesInclude: valid values are %v", v, allowedMetricVolumesIncludeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MetricVolumesInclude) IsValid() bool {
	for _, existing := range allowedMetricVolumesIncludeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricVolumesInclude value.
func (v MetricVolumesInclude) Ptr() *MetricVolumesInclude {
	return &v
}

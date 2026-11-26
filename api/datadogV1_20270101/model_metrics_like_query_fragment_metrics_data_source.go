// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricsLikeQueryFragmentMetricsDataSource Test-only generated schema for the dashboard versioning proof of concept.
type MetricsLikeQueryFragmentMetricsDataSource string

// List of MetricsLikeQueryFragmentMetricsDataSource.
const (
	METRICSLIKEQUERYFRAGMENTMETRICSDATASOURCE_METRICS MetricsLikeQueryFragmentMetricsDataSource = "metrics"
)

var allowedMetricsLikeQueryFragmentMetricsDataSourceEnumValues = []MetricsLikeQueryFragmentMetricsDataSource{
	METRICSLIKEQUERYFRAGMENTMETRICSDATASOURCE_METRICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MetricsLikeQueryFragmentMetricsDataSource) GetAllowedValues() []MetricsLikeQueryFragmentMetricsDataSource {
	return allowedMetricsLikeQueryFragmentMetricsDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MetricsLikeQueryFragmentMetricsDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricsLikeQueryFragmentMetricsDataSource(value)
	return nil
}

// NewMetricsLikeQueryFragmentMetricsDataSourceFromValue returns a pointer to a valid MetricsLikeQueryFragmentMetricsDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMetricsLikeQueryFragmentMetricsDataSourceFromValue(v string) (*MetricsLikeQueryFragmentMetricsDataSource, error) {
	ev := MetricsLikeQueryFragmentMetricsDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MetricsLikeQueryFragmentMetricsDataSource: valid values are %v", v, allowedMetricsLikeQueryFragmentMetricsDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MetricsLikeQueryFragmentMetricsDataSource) IsValid() bool {
	for _, existing := range allowedMetricsLikeQueryFragmentMetricsDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricsLikeQueryFragmentMetricsDataSource value.
func (v MetricsLikeQueryFragmentMetricsDataSource) Ptr() *MetricsLikeQueryFragmentMetricsDataSource {
	return &v
}

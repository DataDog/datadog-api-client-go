// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableMetricsQueryDataSource Metrics data source.
type GuidedTableMetricsQueryDataSource string

// List of GuidedTableMetricsQueryDataSource.
const (
	GUIDEDTABLEMETRICSQUERYDATASOURCE_METRICS    GuidedTableMetricsQueryDataSource = "metrics"
	GUIDEDTABLEMETRICSQUERYDATASOURCE_CLOUD_COST GuidedTableMetricsQueryDataSource = "cloud_cost"
)

var allowedGuidedTableMetricsQueryDataSourceEnumValues = []GuidedTableMetricsQueryDataSource{
	GUIDEDTABLEMETRICSQUERYDATASOURCE_METRICS,
	GUIDEDTABLEMETRICSQUERYDATASOURCE_CLOUD_COST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GuidedTableMetricsQueryDataSource) GetAllowedValues() []GuidedTableMetricsQueryDataSource {
	return allowedGuidedTableMetricsQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GuidedTableMetricsQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GuidedTableMetricsQueryDataSource(value)
	return nil
}

// NewGuidedTableMetricsQueryDataSourceFromValue returns a pointer to a valid GuidedTableMetricsQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGuidedTableMetricsQueryDataSourceFromValue(v string) (*GuidedTableMetricsQueryDataSource, error) {
	ev := GuidedTableMetricsQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GuidedTableMetricsQueryDataSource: valid values are %v", v, allowedGuidedTableMetricsQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GuidedTableMetricsQueryDataSource) IsValid() bool {
	for _, existing := range allowedGuidedTableMetricsQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GuidedTableMetricsQueryDataSource value.
func (v GuidedTableMetricsQueryDataSource) Ptr() *GuidedTableMetricsQueryDataSource {
	return &v
}

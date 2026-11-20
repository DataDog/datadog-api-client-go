// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricsLikeQueryFragmentCloudCostDataSource Test-only generated schema for the dashboard versioning proof of concept.
type MetricsLikeQueryFragmentCloudCostDataSource string

// List of MetricsLikeQueryFragmentCloudCostDataSource.
const (
	METRICSLIKEQUERYFRAGMENTCLOUDCOSTDATASOURCE_CLOUD_COST MetricsLikeQueryFragmentCloudCostDataSource = "cloud_cost"
)

var allowedMetricsLikeQueryFragmentCloudCostDataSourceEnumValues = []MetricsLikeQueryFragmentCloudCostDataSource{
	METRICSLIKEQUERYFRAGMENTCLOUDCOSTDATASOURCE_CLOUD_COST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MetricsLikeQueryFragmentCloudCostDataSource) GetAllowedValues() []MetricsLikeQueryFragmentCloudCostDataSource {
	return allowedMetricsLikeQueryFragmentCloudCostDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MetricsLikeQueryFragmentCloudCostDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricsLikeQueryFragmentCloudCostDataSource(value)
	return nil
}

// NewMetricsLikeQueryFragmentCloudCostDataSourceFromValue returns a pointer to a valid MetricsLikeQueryFragmentCloudCostDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMetricsLikeQueryFragmentCloudCostDataSourceFromValue(v string) (*MetricsLikeQueryFragmentCloudCostDataSource, error) {
	ev := MetricsLikeQueryFragmentCloudCostDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MetricsLikeQueryFragmentCloudCostDataSource: valid values are %v", v, allowedMetricsLikeQueryFragmentCloudCostDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MetricsLikeQueryFragmentCloudCostDataSource) IsValid() bool {
	for _, existing := range allowedMetricsLikeQueryFragmentCloudCostDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricsLikeQueryFragmentCloudCostDataSource value.
func (v MetricsLikeQueryFragmentCloudCostDataSource) Ptr() *MetricsLikeQueryFragmentCloudCostDataSource {
	return &v
}

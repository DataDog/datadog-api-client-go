// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApmMetricsQueryFragmentDataSource Test-only generated schema for the dashboard versioning proof of concept.
type ApmMetricsQueryFragmentDataSource string

// List of ApmMetricsQueryFragmentDataSource.
const (
	APMMETRICSQUERYFRAGMENTDATASOURCE_APM_METRICS ApmMetricsQueryFragmentDataSource = "apm_metrics"
)

var allowedApmMetricsQueryFragmentDataSourceEnumValues = []ApmMetricsQueryFragmentDataSource{
	APMMETRICSQUERYFRAGMENTDATASOURCE_APM_METRICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApmMetricsQueryFragmentDataSource) GetAllowedValues() []ApmMetricsQueryFragmentDataSource {
	return allowedApmMetricsQueryFragmentDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApmMetricsQueryFragmentDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApmMetricsQueryFragmentDataSource(value)
	return nil
}

// NewApmMetricsQueryFragmentDataSourceFromValue returns a pointer to a valid ApmMetricsQueryFragmentDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApmMetricsQueryFragmentDataSourceFromValue(v string) (*ApmMetricsQueryFragmentDataSource, error) {
	ev := ApmMetricsQueryFragmentDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApmMetricsQueryFragmentDataSource: valid values are %v", v, allowedApmMetricsQueryFragmentDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApmMetricsQueryFragmentDataSource) IsValid() bool {
	for _, existing := range allowedApmMetricsQueryFragmentDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApmMetricsQueryFragmentDataSource value.
func (v ApmMetricsQueryFragmentDataSource) Ptr() *ApmMetricsQueryFragmentDataSource {
	return &v
}

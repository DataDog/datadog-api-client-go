// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApmRecommendationsStreamQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type ApmRecommendationsStreamQueryDataSource string

// List of ApmRecommendationsStreamQueryDataSource.
const (
	APMRECOMMENDATIONSSTREAMQUERYDATASOURCE_APM_RECOMMENDATIONS_STREAM ApmRecommendationsStreamQueryDataSource = "apm_recommendations_stream"
)

var allowedApmRecommendationsStreamQueryDataSourceEnumValues = []ApmRecommendationsStreamQueryDataSource{
	APMRECOMMENDATIONSSTREAMQUERYDATASOURCE_APM_RECOMMENDATIONS_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApmRecommendationsStreamQueryDataSource) GetAllowedValues() []ApmRecommendationsStreamQueryDataSource {
	return allowedApmRecommendationsStreamQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApmRecommendationsStreamQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApmRecommendationsStreamQueryDataSource(value)
	return nil
}

// NewApmRecommendationsStreamQueryDataSourceFromValue returns a pointer to a valid ApmRecommendationsStreamQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApmRecommendationsStreamQueryDataSourceFromValue(v string) (*ApmRecommendationsStreamQueryDataSource, error) {
	ev := ApmRecommendationsStreamQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApmRecommendationsStreamQueryDataSource: valid values are %v", v, allowedApmRecommendationsStreamQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApmRecommendationsStreamQueryDataSource) IsValid() bool {
	for _, existing := range allowedApmRecommendationsStreamQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApmRecommendationsStreamQueryDataSource value.
func (v ApmRecommendationsStreamQueryDataSource) Ptr() *ApmRecommendationsStreamQueryDataSource {
	return &v
}

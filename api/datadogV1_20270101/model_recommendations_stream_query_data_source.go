// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RecommendationsStreamQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type RecommendationsStreamQueryDataSource string

// List of RecommendationsStreamQueryDataSource.
const (
	RECOMMENDATIONSSTREAMQUERYDATASOURCE_RECOMMENDATIONS_STREAM RecommendationsStreamQueryDataSource = "recommendations_stream"
)

var allowedRecommendationsStreamQueryDataSourceEnumValues = []RecommendationsStreamQueryDataSource{
	RECOMMENDATIONSSTREAMQUERYDATASOURCE_RECOMMENDATIONS_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RecommendationsStreamQueryDataSource) GetAllowedValues() []RecommendationsStreamQueryDataSource {
	return allowedRecommendationsStreamQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RecommendationsStreamQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RecommendationsStreamQueryDataSource(value)
	return nil
}

// NewRecommendationsStreamQueryDataSourceFromValue returns a pointer to a valid RecommendationsStreamQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRecommendationsStreamQueryDataSourceFromValue(v string) (*RecommendationsStreamQueryDataSource, error) {
	ev := RecommendationsStreamQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RecommendationsStreamQueryDataSource: valid values are %v", v, allowedRecommendationsStreamQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RecommendationsStreamQueryDataSource) IsValid() bool {
	for _, existing := range allowedRecommendationsStreamQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RecommendationsStreamQueryDataSource value.
func (v RecommendationsStreamQueryDataSource) Ptr() *RecommendationsStreamQueryDataSource {
	return &v
}

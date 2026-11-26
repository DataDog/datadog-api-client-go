// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabaseRecommendationsStreamQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type DatabaseRecommendationsStreamQueryDataSource string

// List of DatabaseRecommendationsStreamQueryDataSource.
const (
	DATABASERECOMMENDATIONSSTREAMQUERYDATASOURCE_DATABASE_RECOMMENDATIONS_STREAM DatabaseRecommendationsStreamQueryDataSource = "database_recommendations_stream"
)

var allowedDatabaseRecommendationsStreamQueryDataSourceEnumValues = []DatabaseRecommendationsStreamQueryDataSource{
	DATABASERECOMMENDATIONSSTREAMQUERYDATASOURCE_DATABASE_RECOMMENDATIONS_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DatabaseRecommendationsStreamQueryDataSource) GetAllowedValues() []DatabaseRecommendationsStreamQueryDataSource {
	return allowedDatabaseRecommendationsStreamQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DatabaseRecommendationsStreamQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DatabaseRecommendationsStreamQueryDataSource(value)
	return nil
}

// NewDatabaseRecommendationsStreamQueryDataSourceFromValue returns a pointer to a valid DatabaseRecommendationsStreamQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDatabaseRecommendationsStreamQueryDataSourceFromValue(v string) (*DatabaseRecommendationsStreamQueryDataSource, error) {
	ev := DatabaseRecommendationsStreamQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DatabaseRecommendationsStreamQueryDataSource: valid values are %v", v, allowedDatabaseRecommendationsStreamQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DatabaseRecommendationsStreamQueryDataSource) IsValid() bool {
	for _, existing := range allowedDatabaseRecommendationsStreamQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatabaseRecommendationsStreamQueryDataSource value.
func (v DatabaseRecommendationsStreamQueryDataSource) Ptr() *DatabaseRecommendationsStreamQueryDataSource {
	return &v
}

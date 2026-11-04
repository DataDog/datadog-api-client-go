// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsQueryDefinitionDataSource The data source of the query.
type CommitmentsQueryDefinitionDataSource string

// List of CommitmentsQueryDefinitionDataSource.
const (
	COMMITMENTSQUERYDEFINITIONDATASOURCE_CLOUD_COST CommitmentsQueryDefinitionDataSource = "cloud_cost"
	COMMITMENTSQUERYDEFINITIONDATASOURCE_METRICS    CommitmentsQueryDefinitionDataSource = "metrics"
)

var allowedCommitmentsQueryDefinitionDataSourceEnumValues = []CommitmentsQueryDefinitionDataSource{
	COMMITMENTSQUERYDEFINITIONDATASOURCE_CLOUD_COST,
	COMMITMENTSQUERYDEFINITIONDATASOURCE_METRICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CommitmentsQueryDefinitionDataSource) GetAllowedValues() []CommitmentsQueryDefinitionDataSource {
	return allowedCommitmentsQueryDefinitionDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CommitmentsQueryDefinitionDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CommitmentsQueryDefinitionDataSource(value)
	return nil
}

// NewCommitmentsQueryDefinitionDataSourceFromValue returns a pointer to a valid CommitmentsQueryDefinitionDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCommitmentsQueryDefinitionDataSourceFromValue(v string) (*CommitmentsQueryDefinitionDataSource, error) {
	ev := CommitmentsQueryDefinitionDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CommitmentsQueryDefinitionDataSource: valid values are %v", v, allowedCommitmentsQueryDefinitionDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CommitmentsQueryDefinitionDataSource) IsValid() bool {
	for _, existing := range allowedCommitmentsQueryDefinitionDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CommitmentsQueryDefinitionDataSource value.
func (v CommitmentsQueryDefinitionDataSource) Ptr() *CommitmentsQueryDefinitionDataSource {
	return &v
}

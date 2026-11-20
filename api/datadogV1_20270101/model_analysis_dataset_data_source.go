// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnalysisDatasetDataSource Test-only generated schema for the dashboard versioning proof of concept.
type AnalysisDatasetDataSource string

// List of AnalysisDatasetDataSource.
const (
	ANALYSISDATASETDATASOURCE_ANALYSIS_DATASET AnalysisDatasetDataSource = "analysis_dataset"
)

var allowedAnalysisDatasetDataSourceEnumValues = []AnalysisDatasetDataSource{
	ANALYSISDATASETDATASOURCE_ANALYSIS_DATASET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AnalysisDatasetDataSource) GetAllowedValues() []AnalysisDatasetDataSource {
	return allowedAnalysisDatasetDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AnalysisDatasetDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AnalysisDatasetDataSource(value)
	return nil
}

// NewAnalysisDatasetDataSourceFromValue returns a pointer to a valid AnalysisDatasetDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAnalysisDatasetDataSourceFromValue(v string) (*AnalysisDatasetDataSource, error) {
	ev := AnalysisDatasetDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AnalysisDatasetDataSource: valid values are %v", v, allowedAnalysisDatasetDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AnalysisDatasetDataSource) IsValid() bool {
	for _, existing := range allowedAnalysisDatasetDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AnalysisDatasetDataSource value.
func (v AnalysisDatasetDataSource) Ptr() *AnalysisDatasetDataSource {
	return &v
}

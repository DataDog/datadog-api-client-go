// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaApmDependencyStatsQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type FormulaApmDependencyStatsQueryDataSource string

// List of FormulaApmDependencyStatsQueryDataSource.
const (
	FORMULAAPMDEPENDENCYSTATSQUERYDATASOURCE_APM_DEPENDENCY_STATS FormulaApmDependencyStatsQueryDataSource = "apm_dependency_stats"
)

var allowedFormulaApmDependencyStatsQueryDataSourceEnumValues = []FormulaApmDependencyStatsQueryDataSource{
	FORMULAAPMDEPENDENCYSTATSQUERYDATASOURCE_APM_DEPENDENCY_STATS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaApmDependencyStatsQueryDataSource) GetAllowedValues() []FormulaApmDependencyStatsQueryDataSource {
	return allowedFormulaApmDependencyStatsQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaApmDependencyStatsQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaApmDependencyStatsQueryDataSource(value)
	return nil
}

// NewFormulaApmDependencyStatsQueryDataSourceFromValue returns a pointer to a valid FormulaApmDependencyStatsQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaApmDependencyStatsQueryDataSourceFromValue(v string) (*FormulaApmDependencyStatsQueryDataSource, error) {
	ev := FormulaApmDependencyStatsQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaApmDependencyStatsQueryDataSource: valid values are %v", v, allowedFormulaApmDependencyStatsQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaApmDependencyStatsQueryDataSource) IsValid() bool {
	for _, existing := range allowedFormulaApmDependencyStatsQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaApmDependencyStatsQueryDataSource value.
func (v FormulaApmDependencyStatsQueryDataSource) Ptr() *FormulaApmDependencyStatsQueryDataSource {
	return &v
}

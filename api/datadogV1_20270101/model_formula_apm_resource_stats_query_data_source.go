// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaApmResourceStatsQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type FormulaApmResourceStatsQueryDataSource string

// List of FormulaApmResourceStatsQueryDataSource.
const (
	FORMULAAPMRESOURCESTATSQUERYDATASOURCE_APM_RESOURCE_STATS FormulaApmResourceStatsQueryDataSource = "apm_resource_stats"
)

var allowedFormulaApmResourceStatsQueryDataSourceEnumValues = []FormulaApmResourceStatsQueryDataSource{
	FORMULAAPMRESOURCESTATSQUERYDATASOURCE_APM_RESOURCE_STATS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaApmResourceStatsQueryDataSource) GetAllowedValues() []FormulaApmResourceStatsQueryDataSource {
	return allowedFormulaApmResourceStatsQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaApmResourceStatsQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaApmResourceStatsQueryDataSource(value)
	return nil
}

// NewFormulaApmResourceStatsQueryDataSourceFromValue returns a pointer to a valid FormulaApmResourceStatsQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaApmResourceStatsQueryDataSourceFromValue(v string) (*FormulaApmResourceStatsQueryDataSource, error) {
	ev := FormulaApmResourceStatsQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaApmResourceStatsQueryDataSource: valid values are %v", v, allowedFormulaApmResourceStatsQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaApmResourceStatsQueryDataSource) IsValid() bool {
	for _, existing := range allowedFormulaApmResourceStatsQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaApmResourceStatsQueryDataSource value.
func (v FormulaApmResourceStatsQueryDataSource) Ptr() *FormulaApmResourceStatsQueryDataSource {
	return &v
}

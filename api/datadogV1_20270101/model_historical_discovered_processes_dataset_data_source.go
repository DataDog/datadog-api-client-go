// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HistoricalDiscoveredProcessesDatasetDataSource Test-only generated schema for the dashboard versioning proof of concept.
type HistoricalDiscoveredProcessesDatasetDataSource string

// List of HistoricalDiscoveredProcessesDatasetDataSource.
const (
	HISTORICALDISCOVEREDPROCESSESDATASETDATASOURCE_HISTORICAL_DISCOVERED_PROCESSES HistoricalDiscoveredProcessesDatasetDataSource = "historical_discovered_processes"
)

var allowedHistoricalDiscoveredProcessesDatasetDataSourceEnumValues = []HistoricalDiscoveredProcessesDatasetDataSource{
	HISTORICALDISCOVEREDPROCESSESDATASETDATASOURCE_HISTORICAL_DISCOVERED_PROCESSES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HistoricalDiscoveredProcessesDatasetDataSource) GetAllowedValues() []HistoricalDiscoveredProcessesDatasetDataSource {
	return allowedHistoricalDiscoveredProcessesDatasetDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HistoricalDiscoveredProcessesDatasetDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HistoricalDiscoveredProcessesDatasetDataSource(value)
	return nil
}

// NewHistoricalDiscoveredProcessesDatasetDataSourceFromValue returns a pointer to a valid HistoricalDiscoveredProcessesDatasetDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHistoricalDiscoveredProcessesDatasetDataSourceFromValue(v string) (*HistoricalDiscoveredProcessesDatasetDataSource, error) {
	ev := HistoricalDiscoveredProcessesDatasetDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HistoricalDiscoveredProcessesDatasetDataSource: valid values are %v", v, allowedHistoricalDiscoveredProcessesDatasetDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HistoricalDiscoveredProcessesDatasetDataSource) IsValid() bool {
	for _, existing := range allowedHistoricalDiscoveredProcessesDatasetDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HistoricalDiscoveredProcessesDatasetDataSource value.
func (v HistoricalDiscoveredProcessesDatasetDataSource) Ptr() *HistoricalDiscoveredProcessesDatasetDataSource {
	return &v
}

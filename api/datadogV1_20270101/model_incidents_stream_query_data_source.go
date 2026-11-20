// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentsStreamQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type IncidentsStreamQueryDataSource string

// List of IncidentsStreamQueryDataSource.
const (
	INCIDENTSSTREAMQUERYDATASOURCE_INCIDENTS_STREAM IncidentsStreamQueryDataSource = "incidents_stream"
)

var allowedIncidentsStreamQueryDataSourceEnumValues = []IncidentsStreamQueryDataSource{
	INCIDENTSSTREAMQUERYDATASOURCE_INCIDENTS_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentsStreamQueryDataSource) GetAllowedValues() []IncidentsStreamQueryDataSource {
	return allowedIncidentsStreamQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentsStreamQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentsStreamQueryDataSource(value)
	return nil
}

// NewIncidentsStreamQueryDataSourceFromValue returns a pointer to a valid IncidentsStreamQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentsStreamQueryDataSourceFromValue(v string) (*IncidentsStreamQueryDataSource, error) {
	ev := IncidentsStreamQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentsStreamQueryDataSource: valid values are %v", v, allowedIncidentsStreamQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentsStreamQueryDataSource) IsValid() bool {
	for _, existing := range allowedIncidentsStreamQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentsStreamQueryDataSource value.
func (v IncidentsStreamQueryDataSource) Ptr() *IncidentsStreamQueryDataSource {
	return &v
}

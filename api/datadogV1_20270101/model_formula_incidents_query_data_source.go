// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaIncidentsQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type FormulaIncidentsQueryDataSource string

// List of FormulaIncidentsQueryDataSource.
const (
	FORMULAINCIDENTSQUERYDATASOURCE_INCIDENT_ANALYTICS FormulaIncidentsQueryDataSource = "incident_analytics"
)

var allowedFormulaIncidentsQueryDataSourceEnumValues = []FormulaIncidentsQueryDataSource{
	FORMULAINCIDENTSQUERYDATASOURCE_INCIDENT_ANALYTICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaIncidentsQueryDataSource) GetAllowedValues() []FormulaIncidentsQueryDataSource {
	return allowedFormulaIncidentsQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaIncidentsQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaIncidentsQueryDataSource(value)
	return nil
}

// NewFormulaIncidentsQueryDataSourceFromValue returns a pointer to a valid FormulaIncidentsQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaIncidentsQueryDataSourceFromValue(v string) (*FormulaIncidentsQueryDataSource, error) {
	ev := FormulaIncidentsQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaIncidentsQueryDataSource: valid values are %v", v, allowedFormulaIncidentsQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaIncidentsQueryDataSource) IsValid() bool {
	for _, existing := range allowedFormulaIncidentsQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaIncidentsQueryDataSource value.
func (v FormulaIncidentsQueryDataSource) Ptr() *FormulaIncidentsQueryDataSource {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaChangeTrackingQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type FormulaChangeTrackingQueryDataSource string

// List of FormulaChangeTrackingQueryDataSource.
const (
	FORMULACHANGETRACKINGQUERYDATASOURCE_CHANGE_TRACKING FormulaChangeTrackingQueryDataSource = "change_tracking"
)

var allowedFormulaChangeTrackingQueryDataSourceEnumValues = []FormulaChangeTrackingQueryDataSource{
	FORMULACHANGETRACKINGQUERYDATASOURCE_CHANGE_TRACKING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaChangeTrackingQueryDataSource) GetAllowedValues() []FormulaChangeTrackingQueryDataSource {
	return allowedFormulaChangeTrackingQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaChangeTrackingQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaChangeTrackingQueryDataSource(value)
	return nil
}

// NewFormulaChangeTrackingQueryDataSourceFromValue returns a pointer to a valid FormulaChangeTrackingQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaChangeTrackingQueryDataSourceFromValue(v string) (*FormulaChangeTrackingQueryDataSource, error) {
	ev := FormulaChangeTrackingQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaChangeTrackingQueryDataSource: valid values are %v", v, allowedFormulaChangeTrackingQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaChangeTrackingQueryDataSource) IsValid() bool {
	for _, existing := range allowedFormulaChangeTrackingQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaChangeTrackingQueryDataSource value.
func (v FormulaChangeTrackingQueryDataSource) Ptr() *FormulaChangeTrackingQueryDataSource {
	return &v
}

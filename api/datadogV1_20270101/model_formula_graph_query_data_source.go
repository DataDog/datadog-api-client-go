// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaGraphQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type FormulaGraphQueryDataSource string

// List of FormulaGraphQueryDataSource.
const (
	FORMULAGRAPHQUERYDATASOURCE_TRACES FormulaGraphQueryDataSource = "traces"
)

var allowedFormulaGraphQueryDataSourceEnumValues = []FormulaGraphQueryDataSource{
	FORMULAGRAPHQUERYDATASOURCE_TRACES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaGraphQueryDataSource) GetAllowedValues() []FormulaGraphQueryDataSource {
	return allowedFormulaGraphQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaGraphQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaGraphQueryDataSource(value)
	return nil
}

// NewFormulaGraphQueryDataSourceFromValue returns a pointer to a valid FormulaGraphQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaGraphQueryDataSourceFromValue(v string) (*FormulaGraphQueryDataSource, error) {
	ev := FormulaGraphQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaGraphQueryDataSource: valid values are %v", v, allowedFormulaGraphQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaGraphQueryDataSource) IsValid() bool {
	for _, existing := range allowedFormulaGraphQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaGraphQueryDataSource value.
func (v FormulaGraphQueryDataSource) Ptr() *FormulaGraphQueryDataSource {
	return &v
}

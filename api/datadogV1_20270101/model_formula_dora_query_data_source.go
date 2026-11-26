// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaDoraQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type FormulaDoraQueryDataSource string

// List of FormulaDoraQueryDataSource.
const (
	FORMULADORAQUERYDATASOURCE_DORA FormulaDoraQueryDataSource = "dora"
)

var allowedFormulaDoraQueryDataSourceEnumValues = []FormulaDoraQueryDataSource{
	FORMULADORAQUERYDATASOURCE_DORA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaDoraQueryDataSource) GetAllowedValues() []FormulaDoraQueryDataSource {
	return allowedFormulaDoraQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaDoraQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaDoraQueryDataSource(value)
	return nil
}

// NewFormulaDoraQueryDataSourceFromValue returns a pointer to a valid FormulaDoraQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaDoraQueryDataSourceFromValue(v string) (*FormulaDoraQueryDataSource, error) {
	ev := FormulaDoraQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaDoraQueryDataSource: valid values are %v", v, allowedFormulaDoraQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaDoraQueryDataSource) IsValid() bool {
	for _, existing := range allowedFormulaDoraQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaDoraQueryDataSource value.
func (v FormulaDoraQueryDataSource) Ptr() *FormulaDoraQueryDataSource {
	return &v
}

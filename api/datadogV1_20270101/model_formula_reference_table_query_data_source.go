// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaReferenceTableQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type FormulaReferenceTableQueryDataSource string

// List of FormulaReferenceTableQueryDataSource.
const (
	FORMULAREFERENCETABLEQUERYDATASOURCE_REFERENCE_TABLE FormulaReferenceTableQueryDataSource = "reference_table"
)

var allowedFormulaReferenceTableQueryDataSourceEnumValues = []FormulaReferenceTableQueryDataSource{
	FORMULAREFERENCETABLEQUERYDATASOURCE_REFERENCE_TABLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaReferenceTableQueryDataSource) GetAllowedValues() []FormulaReferenceTableQueryDataSource {
	return allowedFormulaReferenceTableQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaReferenceTableQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaReferenceTableQueryDataSource(value)
	return nil
}

// NewFormulaReferenceTableQueryDataSourceFromValue returns a pointer to a valid FormulaReferenceTableQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaReferenceTableQueryDataSourceFromValue(v string) (*FormulaReferenceTableQueryDataSource, error) {
	ev := FormulaReferenceTableQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaReferenceTableQueryDataSource: valid values are %v", v, allowedFormulaReferenceTableQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaReferenceTableQueryDataSource) IsValid() bool {
	for _, existing := range allowedFormulaReferenceTableQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaReferenceTableQueryDataSource value.
func (v FormulaReferenceTableQueryDataSource) Ptr() *FormulaReferenceTableQueryDataSource {
	return &v
}

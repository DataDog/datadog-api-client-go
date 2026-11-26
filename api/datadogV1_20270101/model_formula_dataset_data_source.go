// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaDatasetDataSource Test-only generated schema for the dashboard versioning proof of concept.
type FormulaDatasetDataSource string

// List of FormulaDatasetDataSource.
const (
	FORMULADATASETDATASOURCE_FORMULA FormulaDatasetDataSource = "formula"
)

var allowedFormulaDatasetDataSourceEnumValues = []FormulaDatasetDataSource{
	FORMULADATASETDATASOURCE_FORMULA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaDatasetDataSource) GetAllowedValues() []FormulaDatasetDataSource {
	return allowedFormulaDatasetDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaDatasetDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaDatasetDataSource(value)
	return nil
}

// NewFormulaDatasetDataSourceFromValue returns a pointer to a valid FormulaDatasetDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaDatasetDataSourceFromValue(v string) (*FormulaDatasetDataSource, error) {
	ev := FormulaDatasetDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaDatasetDataSource: valid values are %v", v, allowedFormulaDatasetDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaDatasetDataSource) IsValid() bool {
	for _, existing := range allowedFormulaDatasetDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaDatasetDataSource value.
func (v FormulaDatasetDataSource) Ptr() *FormulaDatasetDataSource {
	return &v
}

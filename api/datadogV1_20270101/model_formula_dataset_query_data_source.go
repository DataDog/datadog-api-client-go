// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaDatasetQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type FormulaDatasetQueryDataSource string

// List of FormulaDatasetQueryDataSource.
const (
	FORMULADATASETQUERYDATASOURCE_DATASET FormulaDatasetQueryDataSource = "dataset"
)

var allowedFormulaDatasetQueryDataSourceEnumValues = []FormulaDatasetQueryDataSource{
	FORMULADATASETQUERYDATASOURCE_DATASET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaDatasetQueryDataSource) GetAllowedValues() []FormulaDatasetQueryDataSource {
	return allowedFormulaDatasetQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaDatasetQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaDatasetQueryDataSource(value)
	return nil
}

// NewFormulaDatasetQueryDataSourceFromValue returns a pointer to a valid FormulaDatasetQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaDatasetQueryDataSourceFromValue(v string) (*FormulaDatasetQueryDataSource, error) {
	ev := FormulaDatasetQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaDatasetQueryDataSource: valid values are %v", v, allowedFormulaDatasetQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaDatasetQueryDataSource) IsValid() bool {
	for _, existing := range allowedFormulaDatasetQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaDatasetQueryDataSource value.
func (v FormulaDatasetQueryDataSource) Ptr() *FormulaDatasetQueryDataSource {
	return &v
}

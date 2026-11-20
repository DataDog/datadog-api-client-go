// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReferentialDatasetReferenceTableDataSource Test-only generated schema for the dashboard versioning proof of concept.
type ReferentialDatasetReferenceTableDataSource string

// List of ReferentialDatasetReferenceTableDataSource.
const (
	REFERENTIALDATASETREFERENCETABLEDATASOURCE_REFERENCE_TABLE ReferentialDatasetReferenceTableDataSource = "reference_table"
)

var allowedReferentialDatasetReferenceTableDataSourceEnumValues = []ReferentialDatasetReferenceTableDataSource{
	REFERENTIALDATASETREFERENCETABLEDATASOURCE_REFERENCE_TABLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReferentialDatasetReferenceTableDataSource) GetAllowedValues() []ReferentialDatasetReferenceTableDataSource {
	return allowedReferentialDatasetReferenceTableDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReferentialDatasetReferenceTableDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReferentialDatasetReferenceTableDataSource(value)
	return nil
}

// NewReferentialDatasetReferenceTableDataSourceFromValue returns a pointer to a valid ReferentialDatasetReferenceTableDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReferentialDatasetReferenceTableDataSourceFromValue(v string) (*ReferentialDatasetReferenceTableDataSource, error) {
	ev := ReferentialDatasetReferenceTableDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReferentialDatasetReferenceTableDataSource: valid values are %v", v, allowedReferentialDatasetReferenceTableDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReferentialDatasetReferenceTableDataSource) IsValid() bool {
	for _, existing := range allowedReferentialDatasetReferenceTableDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReferentialDatasetReferenceTableDataSource value.
func (v ReferentialDatasetReferenceTableDataSource) Ptr() *ReferentialDatasetReferenceTableDataSource {
	return &v
}

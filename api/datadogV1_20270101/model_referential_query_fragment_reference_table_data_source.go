// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReferentialQueryFragmentReferenceTableDataSource Test-only generated schema for the dashboard versioning proof of concept.
type ReferentialQueryFragmentReferenceTableDataSource string

// List of ReferentialQueryFragmentReferenceTableDataSource.
const (
	REFERENTIALQUERYFRAGMENTREFERENCETABLEDATASOURCE_REFERENCE_TABLE ReferentialQueryFragmentReferenceTableDataSource = "reference_table"
)

var allowedReferentialQueryFragmentReferenceTableDataSourceEnumValues = []ReferentialQueryFragmentReferenceTableDataSource{
	REFERENTIALQUERYFRAGMENTREFERENCETABLEDATASOURCE_REFERENCE_TABLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReferentialQueryFragmentReferenceTableDataSource) GetAllowedValues() []ReferentialQueryFragmentReferenceTableDataSource {
	return allowedReferentialQueryFragmentReferenceTableDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReferentialQueryFragmentReferenceTableDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReferentialQueryFragmentReferenceTableDataSource(value)
	return nil
}

// NewReferentialQueryFragmentReferenceTableDataSourceFromValue returns a pointer to a valid ReferentialQueryFragmentReferenceTableDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReferentialQueryFragmentReferenceTableDataSourceFromValue(v string) (*ReferentialQueryFragmentReferenceTableDataSource, error) {
	ev := ReferentialQueryFragmentReferenceTableDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReferentialQueryFragmentReferenceTableDataSource: valid values are %v", v, allowedReferentialQueryFragmentReferenceTableDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReferentialQueryFragmentReferenceTableDataSource) IsValid() bool {
	for _, existing := range allowedReferentialQueryFragmentReferenceTableDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReferentialQueryFragmentReferenceTableDataSource value.
func (v ReferentialQueryFragmentReferenceTableDataSource) Ptr() *ReferentialQueryFragmentReferenceTableDataSource {
	return &v
}

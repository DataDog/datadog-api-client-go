// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReferentialDatasetManagedResourceDataSource Test-only generated schema for the dashboard versioning proof of concept.
type ReferentialDatasetManagedResourceDataSource string

// List of ReferentialDatasetManagedResourceDataSource.
const (
	REFERENTIALDATASETMANAGEDRESOURCEDATASOURCE_MANAGED_RESOURCE ReferentialDatasetManagedResourceDataSource = "managed_resource"
)

var allowedReferentialDatasetManagedResourceDataSourceEnumValues = []ReferentialDatasetManagedResourceDataSource{
	REFERENTIALDATASETMANAGEDRESOURCEDATASOURCE_MANAGED_RESOURCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReferentialDatasetManagedResourceDataSource) GetAllowedValues() []ReferentialDatasetManagedResourceDataSource {
	return allowedReferentialDatasetManagedResourceDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReferentialDatasetManagedResourceDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReferentialDatasetManagedResourceDataSource(value)
	return nil
}

// NewReferentialDatasetManagedResourceDataSourceFromValue returns a pointer to a valid ReferentialDatasetManagedResourceDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReferentialDatasetManagedResourceDataSourceFromValue(v string) (*ReferentialDatasetManagedResourceDataSource, error) {
	ev := ReferentialDatasetManagedResourceDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReferentialDatasetManagedResourceDataSource: valid values are %v", v, allowedReferentialDatasetManagedResourceDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReferentialDatasetManagedResourceDataSource) IsValid() bool {
	for _, existing := range allowedReferentialDatasetManagedResourceDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReferentialDatasetManagedResourceDataSource value.
func (v ReferentialDatasetManagedResourceDataSource) Ptr() *ReferentialDatasetManagedResourceDataSource {
	return &v
}

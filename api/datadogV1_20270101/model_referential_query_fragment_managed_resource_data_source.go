// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReferentialQueryFragmentManagedResourceDataSource Test-only generated schema for the dashboard versioning proof of concept.
type ReferentialQueryFragmentManagedResourceDataSource string

// List of ReferentialQueryFragmentManagedResourceDataSource.
const (
	REFERENTIALQUERYFRAGMENTMANAGEDRESOURCEDATASOURCE_MANAGED_RESOURCE ReferentialQueryFragmentManagedResourceDataSource = "managed_resource"
)

var allowedReferentialQueryFragmentManagedResourceDataSourceEnumValues = []ReferentialQueryFragmentManagedResourceDataSource{
	REFERENTIALQUERYFRAGMENTMANAGEDRESOURCEDATASOURCE_MANAGED_RESOURCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReferentialQueryFragmentManagedResourceDataSource) GetAllowedValues() []ReferentialQueryFragmentManagedResourceDataSource {
	return allowedReferentialQueryFragmentManagedResourceDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReferentialQueryFragmentManagedResourceDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReferentialQueryFragmentManagedResourceDataSource(value)
	return nil
}

// NewReferentialQueryFragmentManagedResourceDataSourceFromValue returns a pointer to a valid ReferentialQueryFragmentManagedResourceDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReferentialQueryFragmentManagedResourceDataSourceFromValue(v string) (*ReferentialQueryFragmentManagedResourceDataSource, error) {
	ev := ReferentialQueryFragmentManagedResourceDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReferentialQueryFragmentManagedResourceDataSource: valid values are %v", v, allowedReferentialQueryFragmentManagedResourceDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReferentialQueryFragmentManagedResourceDataSource) IsValid() bool {
	for _, existing := range allowedReferentialQueryFragmentManagedResourceDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReferentialQueryFragmentManagedResourceDataSource value.
func (v ReferentialQueryFragmentManagedResourceDataSource) Ptr() *ReferentialQueryFragmentManagedResourceDataSource {
	return &v
}

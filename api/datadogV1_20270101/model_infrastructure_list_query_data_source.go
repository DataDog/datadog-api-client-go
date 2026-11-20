// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// InfrastructureListQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type InfrastructureListQueryDataSource string

// List of InfrastructureListQueryDataSource.
const (
	INFRASTRUCTURELISTQUERYDATASOURCE_INFRASTRUCTURE InfrastructureListQueryDataSource = "infrastructure"
)

var allowedInfrastructureListQueryDataSourceEnumValues = []InfrastructureListQueryDataSource{
	INFRASTRUCTURELISTQUERYDATASOURCE_INFRASTRUCTURE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *InfrastructureListQueryDataSource) GetAllowedValues() []InfrastructureListQueryDataSource {
	return allowedInfrastructureListQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InfrastructureListQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InfrastructureListQueryDataSource(value)
	return nil
}

// NewInfrastructureListQueryDataSourceFromValue returns a pointer to a valid InfrastructureListQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInfrastructureListQueryDataSourceFromValue(v string) (*InfrastructureListQueryDataSource, error) {
	ev := InfrastructureListQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InfrastructureListQueryDataSource: valid values are %v", v, allowedInfrastructureListQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InfrastructureListQueryDataSource) IsValid() bool {
	for _, existing := range allowedInfrastructureListQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InfrastructureListQueryDataSource value.
func (v InfrastructureListQueryDataSource) Ptr() *InfrastructureListQueryDataSource {
	return &v
}

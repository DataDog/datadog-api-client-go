// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaManagedResourceQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type FormulaManagedResourceQueryDataSource string

// List of FormulaManagedResourceQueryDataSource.
const (
	FORMULAMANAGEDRESOURCEQUERYDATASOURCE_MANAGED_RESOURCE FormulaManagedResourceQueryDataSource = "managed_resource"
)

var allowedFormulaManagedResourceQueryDataSourceEnumValues = []FormulaManagedResourceQueryDataSource{
	FORMULAMANAGEDRESOURCEQUERYDATASOURCE_MANAGED_RESOURCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaManagedResourceQueryDataSource) GetAllowedValues() []FormulaManagedResourceQueryDataSource {
	return allowedFormulaManagedResourceQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaManagedResourceQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaManagedResourceQueryDataSource(value)
	return nil
}

// NewFormulaManagedResourceQueryDataSourceFromValue returns a pointer to a valid FormulaManagedResourceQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaManagedResourceQueryDataSourceFromValue(v string) (*FormulaManagedResourceQueryDataSource, error) {
	ev := FormulaManagedResourceQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaManagedResourceQueryDataSource: valid values are %v", v, allowedFormulaManagedResourceQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaManagedResourceQueryDataSource) IsValid() bool {
	for _, existing := range allowedFormulaManagedResourceQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaManagedResourceQueryDataSource value.
func (v FormulaManagedResourceQueryDataSource) Ptr() *FormulaManagedResourceQueryDataSource {
	return &v
}

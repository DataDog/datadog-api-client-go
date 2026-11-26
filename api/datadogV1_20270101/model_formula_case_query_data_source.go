// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaCaseQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type FormulaCaseQueryDataSource string

// List of FormulaCaseQueryDataSource.
const (
	FORMULACASEQUERYDATASOURCE_CASES FormulaCaseQueryDataSource = "cases"
)

var allowedFormulaCaseQueryDataSourceEnumValues = []FormulaCaseQueryDataSource{
	FORMULACASEQUERYDATASOURCE_CASES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaCaseQueryDataSource) GetAllowedValues() []FormulaCaseQueryDataSource {
	return allowedFormulaCaseQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaCaseQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaCaseQueryDataSource(value)
	return nil
}

// NewFormulaCaseQueryDataSourceFromValue returns a pointer to a valid FormulaCaseQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaCaseQueryDataSourceFromValue(v string) (*FormulaCaseQueryDataSource, error) {
	ev := FormulaCaseQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaCaseQueryDataSource: valid values are %v", v, allowedFormulaCaseQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaCaseQueryDataSource) IsValid() bool {
	for _, existing := range allowedFormulaCaseQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaCaseQueryDataSource value.
func (v FormulaCaseQueryDataSource) Ptr() *FormulaCaseQueryDataSource {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaSloQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type FormulaSloQueryDataSource string

// List of FormulaSloQueryDataSource.
const (
	FORMULASLOQUERYDATASOURCE_SLO FormulaSloQueryDataSource = "slo"
)

var allowedFormulaSloQueryDataSourceEnumValues = []FormulaSloQueryDataSource{
	FORMULASLOQUERYDATASOURCE_SLO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaSloQueryDataSource) GetAllowedValues() []FormulaSloQueryDataSource {
	return allowedFormulaSloQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaSloQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaSloQueryDataSource(value)
	return nil
}

// NewFormulaSloQueryDataSourceFromValue returns a pointer to a valid FormulaSloQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaSloQueryDataSourceFromValue(v string) (*FormulaSloQueryDataSource, error) {
	ev := FormulaSloQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaSloQueryDataSource: valid values are %v", v, allowedFormulaSloQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaSloQueryDataSource) IsValid() bool {
	for _, existing := range allowedFormulaSloQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaSloQueryDataSource value.
func (v FormulaSloQueryDataSource) Ptr() *FormulaSloQueryDataSource {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CompositeReferentialDataQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type CompositeReferentialDataQueryDataSource string

// List of CompositeReferentialDataQueryDataSource.
const (
	COMPOSITEREFERENTIALDATAQUERYDATASOURCE_REFERENCE_TABLE CompositeReferentialDataQueryDataSource = "reference_table"
)

var allowedCompositeReferentialDataQueryDataSourceEnumValues = []CompositeReferentialDataQueryDataSource{
	COMPOSITEREFERENTIALDATAQUERYDATASOURCE_REFERENCE_TABLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CompositeReferentialDataQueryDataSource) GetAllowedValues() []CompositeReferentialDataQueryDataSource {
	return allowedCompositeReferentialDataQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CompositeReferentialDataQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CompositeReferentialDataQueryDataSource(value)
	return nil
}

// NewCompositeReferentialDataQueryDataSourceFromValue returns a pointer to a valid CompositeReferentialDataQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCompositeReferentialDataQueryDataSourceFromValue(v string) (*CompositeReferentialDataQueryDataSource, error) {
	ev := CompositeReferentialDataQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CompositeReferentialDataQueryDataSource: valid values are %v", v, allowedCompositeReferentialDataQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CompositeReferentialDataQueryDataSource) IsValid() bool {
	for _, existing := range allowedCompositeReferentialDataQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CompositeReferentialDataQueryDataSource value.
func (v CompositeReferentialDataQueryDataSource) Ptr() *CompositeReferentialDataQueryDataSource {
	return &v
}

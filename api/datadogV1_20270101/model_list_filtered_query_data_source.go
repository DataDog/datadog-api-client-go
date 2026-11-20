// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListFilteredQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type ListFilteredQueryDataSource string

// List of ListFilteredQueryDataSource.
const (
	LISTFILTEREDQUERYDATASOURCE_LIST_FILTERED_QUERY ListFilteredQueryDataSource = "list_filtered_query"
)

var allowedListFilteredQueryDataSourceEnumValues = []ListFilteredQueryDataSource{
	LISTFILTEREDQUERYDATASOURCE_LIST_FILTERED_QUERY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ListFilteredQueryDataSource) GetAllowedValues() []ListFilteredQueryDataSource {
	return allowedListFilteredQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ListFilteredQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ListFilteredQueryDataSource(value)
	return nil
}

// NewListFilteredQueryDataSourceFromValue returns a pointer to a valid ListFilteredQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewListFilteredQueryDataSourceFromValue(v string) (*ListFilteredQueryDataSource, error) {
	ev := ListFilteredQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ListFilteredQueryDataSource: valid values are %v", v, allowedListFilteredQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ListFilteredQueryDataSource) IsValid() bool {
	for _, existing := range allowedListFilteredQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListFilteredQueryDataSource value.
func (v ListFilteredQueryDataSource) Ptr() *ListFilteredQueryDataSource {
	return &v
}

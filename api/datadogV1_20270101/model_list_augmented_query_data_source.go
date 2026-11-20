// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListAugmentedQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type ListAugmentedQueryDataSource string

// List of ListAugmentedQueryDataSource.
const (
	LISTAUGMENTEDQUERYDATASOURCE_LIST_AUGMENTED_QUERY ListAugmentedQueryDataSource = "list_augmented_query"
)

var allowedListAugmentedQueryDataSourceEnumValues = []ListAugmentedQueryDataSource{
	LISTAUGMENTEDQUERYDATASOURCE_LIST_AUGMENTED_QUERY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ListAugmentedQueryDataSource) GetAllowedValues() []ListAugmentedQueryDataSource {
	return allowedListAugmentedQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ListAugmentedQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ListAugmentedQueryDataSource(value)
	return nil
}

// NewListAugmentedQueryDataSourceFromValue returns a pointer to a valid ListAugmentedQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewListAugmentedQueryDataSourceFromValue(v string) (*ListAugmentedQueryDataSource, error) {
	ev := ListAugmentedQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ListAugmentedQueryDataSource: valid values are %v", v, allowedListAugmentedQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ListAugmentedQueryDataSource) IsValid() bool {
	for _, existing := range allowedListAugmentedQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListAugmentedQueryDataSource value.
func (v ListAugmentedQueryDataSource) Ptr() *ListAugmentedQueryDataSource {
	return &v
}

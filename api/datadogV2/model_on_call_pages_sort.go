// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPagesSort Sorting options for pages. Prefix a field with `-` to sort in descending order.
type OnCallPagesSort string

// List of OnCallPagesSort.
const (
	ONCALLPAGESSORT_CREATED_AT_ASCENDING   OnCallPagesSort = "created_at"
	ONCALLPAGESSORT_CREATED_AT_DESCENDING  OnCallPagesSort = "-created_at"
	ONCALLPAGESSORT_MODIFIED_AT_ASCENDING  OnCallPagesSort = "modified_at"
	ONCALLPAGESSORT_MODIFIED_AT_DESCENDING OnCallPagesSort = "-modified_at"
	ONCALLPAGESSORT_STATUS_ASCENDING       OnCallPagesSort = "status"
	ONCALLPAGESSORT_STATUS_DESCENDING      OnCallPagesSort = "-status"
	ONCALLPAGESSORT_PRIORITY_ASCENDING     OnCallPagesSort = "priority"
	ONCALLPAGESSORT_PRIORITY_DESCENDING    OnCallPagesSort = "-priority"
)

var allowedOnCallPagesSortEnumValues = []OnCallPagesSort{
	ONCALLPAGESSORT_CREATED_AT_ASCENDING,
	ONCALLPAGESSORT_CREATED_AT_DESCENDING,
	ONCALLPAGESSORT_MODIFIED_AT_ASCENDING,
	ONCALLPAGESSORT_MODIFIED_AT_DESCENDING,
	ONCALLPAGESSORT_STATUS_ASCENDING,
	ONCALLPAGESSORT_STATUS_DESCENDING,
	ONCALLPAGESSORT_PRIORITY_ASCENDING,
	ONCALLPAGESSORT_PRIORITY_DESCENDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnCallPagesSort) GetAllowedValues() []OnCallPagesSort {
	return allowedOnCallPagesSortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnCallPagesSort) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnCallPagesSort(value)
	return nil
}

// NewOnCallPagesSortFromValue returns a pointer to a valid OnCallPagesSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnCallPagesSortFromValue(v string) (*OnCallPagesSort, error) {
	ev := OnCallPagesSort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnCallPagesSort: valid values are %v", v, allowedOnCallPagesSortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnCallPagesSort) IsValid() bool {
	for _, existing := range allowedOnCallPagesSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnCallPagesSort value.
func (v OnCallPagesSort) Ptr() *OnCallPagesSort {
	return &v
}

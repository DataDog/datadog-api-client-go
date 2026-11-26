// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SortOrder Test-only generated schema for the dashboard versioning proof of concept.
type SortOrder string

// List of SortOrder.
const (
	SORTORDER_DESC SortOrder = "desc"
	SORTORDER_ASC  SortOrder = "asc"
)

var allowedSortOrderEnumValues = []SortOrder{
	SORTORDER_DESC,
	SORTORDER_ASC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SortOrder) GetAllowedValues() []SortOrder {
	return allowedSortOrderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SortOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SortOrder(value)
	return nil
}

// NewSortOrderFromValue returns a pointer to a valid SortOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSortOrderFromValue(v string) (*SortOrder, error) {
	ev := SortOrder(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SortOrder: valid values are %v", v, allowedSortOrderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SortOrder) IsValid() bool {
	for _, existing := range allowedSortOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SortOrder value.
func (v SortOrder) Ptr() *SortOrder {
	return &v
}

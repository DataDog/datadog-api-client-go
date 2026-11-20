// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SplitConfigSortOrder Test-only generated schema for the dashboard versioning proof of concept.
type SplitConfigSortOrder string

// List of SplitConfigSortOrder.
const (
	SPLITCONFIGSORTORDER_DESC SplitConfigSortOrder = "desc"
	SPLITCONFIGSORTORDER_ASC  SplitConfigSortOrder = "asc"
)

var allowedSplitConfigSortOrderEnumValues = []SplitConfigSortOrder{
	SPLITCONFIGSORTORDER_DESC,
	SPLITCONFIGSORTORDER_ASC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SplitConfigSortOrder) GetAllowedValues() []SplitConfigSortOrder {
	return allowedSplitConfigSortOrderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SplitConfigSortOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SplitConfigSortOrder(value)
	return nil
}

// NewSplitConfigSortOrderFromValue returns a pointer to a valid SplitConfigSortOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSplitConfigSortOrderFromValue(v string) (*SplitConfigSortOrder, error) {
	ev := SplitConfigSortOrder(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SplitConfigSortOrder: valid values are %v", v, allowedSplitConfigSortOrderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SplitConfigSortOrder) IsValid() bool {
	for _, existing := range allowedSplitConfigSortOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SplitConfigSortOrder value.
func (v SplitConfigSortOrder) Ptr() *SplitConfigSortOrder {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SortOrder2 Test-only generated schema for the dashboard versioning proof of concept.
type SortOrder2 string

// List of SortOrder2.
const (
	SORTORDER2_DESC SortOrder2 = "desc"
	SORTORDER2_ASC  SortOrder2 = "asc"
)

var allowedSortOrder2EnumValues = []SortOrder2{
	SORTORDER2_DESC,
	SORTORDER2_ASC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SortOrder2) GetAllowedValues() []SortOrder2 {
	return allowedSortOrder2EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SortOrder2) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SortOrder2(value)
	return nil
}

// NewSortOrder2FromValue returns a pointer to a valid SortOrder2
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSortOrder2FromValue(v string) (*SortOrder2, error) {
	ev := SortOrder2(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SortOrder2: valid values are %v", v, allowedSortOrder2EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SortOrder2) IsValid() bool {
	for _, existing := range allowedSortOrder2EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SortOrder2 value.
func (v SortOrder2) Ptr() *SortOrder2 {
	return &v
}

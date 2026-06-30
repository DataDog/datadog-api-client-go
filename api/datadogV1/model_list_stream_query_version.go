// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListStreamQueryVersion Version of the query for the logs transaction stream widget. When omitted, v1 query behavior is
// preserved. Set to `sequential_query` to use v2 behavior. **This feature is in Preview.**
type ListStreamQueryVersion string

// List of ListStreamQueryVersion.
const (
	LISTSTREAMQUERYVERSION_SEQUENTIAL_QUERY ListStreamQueryVersion = "sequential_query"
)

var allowedListStreamQueryVersionEnumValues = []ListStreamQueryVersion{
	LISTSTREAMQUERYVERSION_SEQUENTIAL_QUERY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ListStreamQueryVersion) GetAllowedValues() []ListStreamQueryVersion {
	return allowedListStreamQueryVersionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ListStreamQueryVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ListStreamQueryVersion(value)
	return nil
}

// NewListStreamQueryVersionFromValue returns a pointer to a valid ListStreamQueryVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewListStreamQueryVersionFromValue(v string) (*ListStreamQueryVersion, error) {
	ev := ListStreamQueryVersion(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ListStreamQueryVersion: valid values are %v", v, allowedListStreamQueryVersionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ListStreamQueryVersion) IsValid() bool {
	for _, existing := range allowedListStreamQueryVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListStreamQueryVersion value.
func (v ListStreamQueryVersion) Ptr() *ListStreamQueryVersion {
	return &v
}

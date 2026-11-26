// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HasSearchBar Test-only generated schema for the dashboard versioning proof of concept.
type HasSearchBar string

// List of HasSearchBar.
const (
	HASSEARCHBAR_AUTO   HasSearchBar = "auto"
	HASSEARCHBAR_ALWAYS HasSearchBar = "always"
	HASSEARCHBAR_NEVER  HasSearchBar = "never"
)

var allowedHasSearchBarEnumValues = []HasSearchBar{
	HASSEARCHBAR_AUTO,
	HASSEARCHBAR_ALWAYS,
	HASSEARCHBAR_NEVER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HasSearchBar) GetAllowedValues() []HasSearchBar {
	return allowedHasSearchBarEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HasSearchBar) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HasSearchBar(value)
	return nil
}

// NewHasSearchBarFromValue returns a pointer to a valid HasSearchBar
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHasSearchBarFromValue(v string) (*HasSearchBar, error) {
	ev := HasSearchBar(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HasSearchBar: valid values are %v", v, allowedHasSearchBarEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HasSearchBar) IsValid() bool {
	for _, existing := range allowedHasSearchBarEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HasSearchBar value.
func (v HasSearchBar) Ptr() *HasSearchBar {
	return &v
}

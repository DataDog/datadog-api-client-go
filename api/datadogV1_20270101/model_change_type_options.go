// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeTypeOptions Test-only generated schema for the dashboard versioning proof of concept.
type ChangeTypeOptions string

// List of ChangeTypeOptions.
const (
	CHANGETYPEOPTIONS_RELATIVE ChangeTypeOptions = "relative"
	CHANGETYPEOPTIONS_ABSOLUTE ChangeTypeOptions = "absolute"
)

var allowedChangeTypeOptionsEnumValues = []ChangeTypeOptions{
	CHANGETYPEOPTIONS_RELATIVE,
	CHANGETYPEOPTIONS_ABSOLUTE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ChangeTypeOptions) GetAllowedValues() []ChangeTypeOptions {
	return allowedChangeTypeOptionsEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ChangeTypeOptions) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ChangeTypeOptions(value)
	return nil
}

// NewChangeTypeOptionsFromValue returns a pointer to a valid ChangeTypeOptions
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewChangeTypeOptionsFromValue(v string) (*ChangeTypeOptions, error) {
	ev := ChangeTypeOptions(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ChangeTypeOptions: valid values are %v", v, allowedChangeTypeOptionsEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ChangeTypeOptions) IsValid() bool {
	for _, existing := range allowedChangeTypeOptionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChangeTypeOptions value.
func (v ChangeTypeOptions) Ptr() *ChangeTypeOptions {
	return &v
}

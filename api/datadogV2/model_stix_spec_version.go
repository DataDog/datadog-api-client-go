// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// STIXSpecVersion The supported STIX specification version.
type STIXSpecVersion string

// List of STIXSpecVersion.
const (
	STIXSPECVERSION_VERSION_2_1 STIXSpecVersion = "2.1"
)

var allowedSTIXSpecVersionEnumValues = []STIXSpecVersion{
	STIXSPECVERSION_VERSION_2_1,
}

// GetAllowedValues reeturns the list of possible values.
func (v *STIXSpecVersion) GetAllowedValues() []STIXSpecVersion {
	return allowedSTIXSpecVersionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *STIXSpecVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = STIXSpecVersion(value)
	return nil
}

// NewSTIXSpecVersionFromValue returns a pointer to a valid STIXSpecVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSTIXSpecVersionFromValue(v string) (*STIXSpecVersion, error) {
	ev := STIXSpecVersion(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for STIXSpecVersion: valid values are %v", v, allowedSTIXSpecVersionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v STIXSpecVersion) IsValid() bool {
	for _, existing := range allowedSTIXSpecVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to STIXSpecVersion value.
func (v STIXSpecVersion) Ptr() *STIXSpecVersion {
	return &v
}

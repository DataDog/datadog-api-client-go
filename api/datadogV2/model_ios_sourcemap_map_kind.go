// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IOSSourcemapMapKind The type of source map. Always `ios` for iOS dSYM source maps.
type IOSSourcemapMapKind string

// List of IOSSourcemapMapKind.
const (
	IOSSOURCEMAPMAPKIND_IOS IOSSourcemapMapKind = "ios"
)

var allowedIOSSourcemapMapKindEnumValues = []IOSSourcemapMapKind{
	IOSSOURCEMAPMAPKIND_IOS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IOSSourcemapMapKind) GetAllowedValues() []IOSSourcemapMapKind {
	return allowedIOSSourcemapMapKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IOSSourcemapMapKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IOSSourcemapMapKind(value)
	return nil
}

// NewIOSSourcemapMapKindFromValue returns a pointer to a valid IOSSourcemapMapKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIOSSourcemapMapKindFromValue(v string) (*IOSSourcemapMapKind, error) {
	ev := IOSSourcemapMapKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IOSSourcemapMapKind: valid values are %v", v, allowedIOSSourcemapMapKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IOSSourcemapMapKind) IsValid() bool {
	for _, existing := range allowedIOSSourcemapMapKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IOSSourcemapMapKind value.
func (v IOSSourcemapMapKind) Ptr() *IOSSourcemapMapKind {
	return &v
}

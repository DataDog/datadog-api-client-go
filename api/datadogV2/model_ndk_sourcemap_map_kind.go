// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NDKSourcemapMapKind The type of source map. Always `ndk` for Android NDK source maps.
type NDKSourcemapMapKind string

// List of NDKSourcemapMapKind.
const (
	NDKSOURCEMAPMAPKIND_NDK NDKSourcemapMapKind = "ndk"
)

var allowedNDKSourcemapMapKindEnumValues = []NDKSourcemapMapKind{
	NDKSOURCEMAPMAPKIND_NDK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NDKSourcemapMapKind) GetAllowedValues() []NDKSourcemapMapKind {
	return allowedNDKSourcemapMapKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NDKSourcemapMapKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NDKSourcemapMapKind(value)
	return nil
}

// NewNDKSourcemapMapKindFromValue returns a pointer to a valid NDKSourcemapMapKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNDKSourcemapMapKindFromValue(v string) (*NDKSourcemapMapKind, error) {
	ev := NDKSourcemapMapKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NDKSourcemapMapKind: valid values are %v", v, allowedNDKSourcemapMapKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NDKSourcemapMapKind) IsValid() bool {
	for _, existing := range allowedNDKSourcemapMapKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NDKSourcemapMapKind value.
func (v NDKSourcemapMapKind) Ptr() *NDKSourcemapMapKind {
	return &v
}

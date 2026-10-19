// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ELFSourcemapMapKind The type of source map. Always `elf` for ELF source maps.
type ELFSourcemapMapKind string

// List of ELFSourcemapMapKind.
const (
	ELFSOURCEMAPMAPKIND_ELF ELFSourcemapMapKind = "elf"
)

var allowedELFSourcemapMapKindEnumValues = []ELFSourcemapMapKind{
	ELFSOURCEMAPMAPKIND_ELF,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ELFSourcemapMapKind) GetAllowedValues() []ELFSourcemapMapKind {
	return allowedELFSourcemapMapKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ELFSourcemapMapKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ELFSourcemapMapKind(value)
	return nil
}

// NewELFSourcemapMapKindFromValue returns a pointer to a valid ELFSourcemapMapKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewELFSourcemapMapKindFromValue(v string) (*ELFSourcemapMapKind, error) {
	ev := ELFSourcemapMapKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ELFSourcemapMapKind: valid values are %v", v, allowedELFSourcemapMapKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ELFSourcemapMapKind) IsValid() bool {
	for _, existing := range allowedELFSourcemapMapKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ELFSourcemapMapKind value.
func (v ELFSourcemapMapKind) Ptr() *ELFSourcemapMapKind {
	return &v
}

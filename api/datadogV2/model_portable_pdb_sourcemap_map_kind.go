// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PortablePdbSourcemapMapKind The type of source map. Always `portable_pdb` for Portable PDB files.
type PortablePdbSourcemapMapKind string

// List of PortablePdbSourcemapMapKind.
const (
	PORTABLEPDBSOURCEMAPMAPKIND_PORTABLE_PDB PortablePdbSourcemapMapKind = "portable_pdb"
)

var allowedPortablePdbSourcemapMapKindEnumValues = []PortablePdbSourcemapMapKind{
	PORTABLEPDBSOURCEMAPMAPKIND_PORTABLE_PDB,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PortablePdbSourcemapMapKind) GetAllowedValues() []PortablePdbSourcemapMapKind {
	return allowedPortablePdbSourcemapMapKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PortablePdbSourcemapMapKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PortablePdbSourcemapMapKind(value)
	return nil
}

// NewPortablePdbSourcemapMapKindFromValue returns a pointer to a valid PortablePdbSourcemapMapKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPortablePdbSourcemapMapKindFromValue(v string) (*PortablePdbSourcemapMapKind, error) {
	ev := PortablePdbSourcemapMapKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PortablePdbSourcemapMapKind: valid values are %v", v, allowedPortablePdbSourcemapMapKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PortablePdbSourcemapMapKind) IsValid() bool {
	for _, existing := range allowedPortablePdbSourcemapMapKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortablePdbSourcemapMapKind value.
func (v PortablePdbSourcemapMapKind) Ptr() *PortablePdbSourcemapMapKind {
	return &v
}

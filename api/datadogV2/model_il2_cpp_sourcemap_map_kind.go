// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IL2CPPSourcemapMapKind The type of source map. Always `il2cpp` for IL2CPP source maps.
type IL2CPPSourcemapMapKind string

// List of IL2CPPSourcemapMapKind.
const (
	IL2CPPSOURCEMAPMAPKIND_IL2CPP IL2CPPSourcemapMapKind = "il2cpp"
)

var allowedIL2CPPSourcemapMapKindEnumValues = []IL2CPPSourcemapMapKind{
	IL2CPPSOURCEMAPMAPKIND_IL2CPP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IL2CPPSourcemapMapKind) GetAllowedValues() []IL2CPPSourcemapMapKind {
	return allowedIL2CPPSourcemapMapKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IL2CPPSourcemapMapKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IL2CPPSourcemapMapKind(value)
	return nil
}

// NewIL2CPPSourcemapMapKindFromValue returns a pointer to a valid IL2CPPSourcemapMapKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIL2CPPSourcemapMapKindFromValue(v string) (*IL2CPPSourcemapMapKind, error) {
	ev := IL2CPPSourcemapMapKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IL2CPPSourcemapMapKind: valid values are %v", v, allowedIL2CPPSourcemapMapKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IL2CPPSourcemapMapKind) IsValid() bool {
	for _, existing := range allowedIL2CPPSourcemapMapKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IL2CPPSourcemapMapKind value.
func (v IL2CPPSourcemapMapKind) Ptr() *IL2CPPSourcemapMapKind {
	return &v
}

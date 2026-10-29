// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SourcemapSymbolType The WASM symbol representation.
type SourcemapSymbolType string

// List of SourcemapSymbolType.
const (
	SOURCEMAPSYMBOLTYPE_DWARF     SourcemapSymbolType = "dwarf"
	SOURCEMAPSYMBOLTYPE_SOURCEMAP SourcemapSymbolType = "sourcemap"
)

var allowedSourcemapSymbolTypeEnumValues = []SourcemapSymbolType{
	SOURCEMAPSYMBOLTYPE_DWARF,
	SOURCEMAPSYMBOLTYPE_SOURCEMAP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SourcemapSymbolType) GetAllowedValues() []SourcemapSymbolType {
	return allowedSourcemapSymbolTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SourcemapSymbolType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SourcemapSymbolType(value)
	return nil
}

// NewSourcemapSymbolTypeFromValue returns a pointer to a valid SourcemapSymbolType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSourcemapSymbolTypeFromValue(v string) (*SourcemapSymbolType, error) {
	ev := SourcemapSymbolType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SourcemapSymbolType: valid values are %v", v, allowedSourcemapSymbolTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SourcemapSymbolType) IsValid() bool {
	for _, existing := range allowedSourcemapSymbolTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SourcemapSymbolType value.
func (v SourcemapSymbolType) Ptr() *SourcemapSymbolType {
	return &v
}

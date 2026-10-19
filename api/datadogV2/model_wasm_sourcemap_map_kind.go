// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WASMSourcemapMapKind The type of source map. Always `wasm` for WebAssembly symbol files.
type WASMSourcemapMapKind string

// List of WASMSourcemapMapKind.
const (
	WASMSOURCEMAPMAPKIND_WASM WASMSourcemapMapKind = "wasm"
)

var allowedWASMSourcemapMapKindEnumValues = []WASMSourcemapMapKind{
	WASMSOURCEMAPMAPKIND_WASM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WASMSourcemapMapKind) GetAllowedValues() []WASMSourcemapMapKind {
	return allowedWASMSourcemapMapKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WASMSourcemapMapKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WASMSourcemapMapKind(value)
	return nil
}

// NewWASMSourcemapMapKindFromValue returns a pointer to a valid WASMSourcemapMapKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWASMSourcemapMapKindFromValue(v string) (*WASMSourcemapMapKind, error) {
	ev := WASMSourcemapMapKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WASMSourcemapMapKind: valid values are %v", v, allowedWASMSourcemapMapKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WASMSourcemapMapKind) IsValid() bool {
	for _, existing := range allowedWASMSourcemapMapKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WASMSourcemapMapKind value.
func (v WASMSourcemapMapKind) Ptr() *WASMSourcemapMapKind {
	return &v
}

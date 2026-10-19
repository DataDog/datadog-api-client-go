// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlutterSourcemapMapKind The type of source map. Always `flutter` for Flutter source maps.
type FlutterSourcemapMapKind string

// List of FlutterSourcemapMapKind.
const (
	FLUTTERSOURCEMAPMAPKIND_FLUTTER FlutterSourcemapMapKind = "flutter"
)

var allowedFlutterSourcemapMapKindEnumValues = []FlutterSourcemapMapKind{
	FLUTTERSOURCEMAPMAPKIND_FLUTTER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FlutterSourcemapMapKind) GetAllowedValues() []FlutterSourcemapMapKind {
	return allowedFlutterSourcemapMapKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FlutterSourcemapMapKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FlutterSourcemapMapKind(value)
	return nil
}

// NewFlutterSourcemapMapKindFromValue returns a pointer to a valid FlutterSourcemapMapKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFlutterSourcemapMapKindFromValue(v string) (*FlutterSourcemapMapKind, error) {
	ev := FlutterSourcemapMapKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FlutterSourcemapMapKind: valid values are %v", v, allowedFlutterSourcemapMapKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FlutterSourcemapMapKind) IsValid() bool {
	for _, existing := range allowedFlutterSourcemapMapKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FlutterSourcemapMapKind value.
func (v FlutterSourcemapMapKind) Ptr() *FlutterSourcemapMapKind {
	return &v
}

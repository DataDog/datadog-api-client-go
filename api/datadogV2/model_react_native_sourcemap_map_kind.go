// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReactNativeSourcemapMapKind The type of source map. Always `react` for React Native source maps.
type ReactNativeSourcemapMapKind string

// List of ReactNativeSourcemapMapKind.
const (
	REACTNATIVESOURCEMAPMAPKIND_REACT ReactNativeSourcemapMapKind = "react"
)

var allowedReactNativeSourcemapMapKindEnumValues = []ReactNativeSourcemapMapKind{
	REACTNATIVESOURCEMAPMAPKIND_REACT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReactNativeSourcemapMapKind) GetAllowedValues() []ReactNativeSourcemapMapKind {
	return allowedReactNativeSourcemapMapKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReactNativeSourcemapMapKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReactNativeSourcemapMapKind(value)
	return nil
}

// NewReactNativeSourcemapMapKindFromValue returns a pointer to a valid ReactNativeSourcemapMapKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReactNativeSourcemapMapKindFromValue(v string) (*ReactNativeSourcemapMapKind, error) {
	ev := ReactNativeSourcemapMapKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReactNativeSourcemapMapKind: valid values are %v", v, allowedReactNativeSourcemapMapKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReactNativeSourcemapMapKind) IsValid() bool {
	for _, existing := range allowedReactNativeSourcemapMapKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReactNativeSourcemapMapKind value.
func (v ReactNativeSourcemapMapKind) Ptr() *ReactNativeSourcemapMapKind {
	return &v
}

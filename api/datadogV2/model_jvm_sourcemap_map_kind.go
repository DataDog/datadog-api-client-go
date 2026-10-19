// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JVMSourcemapMapKind The type of source map. Always `jvm` for JVM source maps.
type JVMSourcemapMapKind string

// List of JVMSourcemapMapKind.
const (
	JVMSOURCEMAPMAPKIND_JVM JVMSourcemapMapKind = "jvm"
)

var allowedJVMSourcemapMapKindEnumValues = []JVMSourcemapMapKind{
	JVMSOURCEMAPMAPKIND_JVM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *JVMSourcemapMapKind) GetAllowedValues() []JVMSourcemapMapKind {
	return allowedJVMSourcemapMapKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *JVMSourcemapMapKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = JVMSourcemapMapKind(value)
	return nil
}

// NewJVMSourcemapMapKindFromValue returns a pointer to a valid JVMSourcemapMapKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewJVMSourcemapMapKindFromValue(v string) (*JVMSourcemapMapKind, error) {
	ev := JVMSourcemapMapKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for JVMSourcemapMapKind: valid values are %v", v, allowedJVMSourcemapMapKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v JVMSourcemapMapKind) IsValid() bool {
	for _, existing := range allowedJVMSourcemapMapKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JVMSourcemapMapKind value.
func (v JVMSourcemapMapKind) Ptr() *JVMSourcemapMapKind {
	return &v
}

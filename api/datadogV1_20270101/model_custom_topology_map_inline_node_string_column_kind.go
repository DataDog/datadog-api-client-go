// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomTopologyMapInlineNodeStringColumnKind Test-only generated schema for the dashboard versioning proof of concept.
type CustomTopologyMapInlineNodeStringColumnKind string

// List of CustomTopologyMapInlineNodeStringColumnKind.
const (
	CUSTOMTOPOLOGYMAPINLINENODESTRINGCOLUMNKIND_STRING CustomTopologyMapInlineNodeStringColumnKind = "string"
)

var allowedCustomTopologyMapInlineNodeStringColumnKindEnumValues = []CustomTopologyMapInlineNodeStringColumnKind{
	CUSTOMTOPOLOGYMAPINLINENODESTRINGCOLUMNKIND_STRING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomTopologyMapInlineNodeStringColumnKind) GetAllowedValues() []CustomTopologyMapInlineNodeStringColumnKind {
	return allowedCustomTopologyMapInlineNodeStringColumnKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomTopologyMapInlineNodeStringColumnKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomTopologyMapInlineNodeStringColumnKind(value)
	return nil
}

// NewCustomTopologyMapInlineNodeStringColumnKindFromValue returns a pointer to a valid CustomTopologyMapInlineNodeStringColumnKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomTopologyMapInlineNodeStringColumnKindFromValue(v string) (*CustomTopologyMapInlineNodeStringColumnKind, error) {
	ev := CustomTopologyMapInlineNodeStringColumnKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomTopologyMapInlineNodeStringColumnKind: valid values are %v", v, allowedCustomTopologyMapInlineNodeStringColumnKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomTopologyMapInlineNodeStringColumnKind) IsValid() bool {
	for _, existing := range allowedCustomTopologyMapInlineNodeStringColumnKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomTopologyMapInlineNodeStringColumnKind value.
func (v CustomTopologyMapInlineNodeStringColumnKind) Ptr() *CustomTopologyMapInlineNodeStringColumnKind {
	return &v
}

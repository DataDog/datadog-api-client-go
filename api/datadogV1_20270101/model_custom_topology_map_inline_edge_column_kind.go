// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomTopologyMapInlineEdgeColumnKind Test-only generated schema for the dashboard versioning proof of concept.
type CustomTopologyMapInlineEdgeColumnKind string

// List of CustomTopologyMapInlineEdgeColumnKind.
const (
	CUSTOMTOPOLOGYMAPINLINEEDGECOLUMNKIND_STRING CustomTopologyMapInlineEdgeColumnKind = "string"
)

var allowedCustomTopologyMapInlineEdgeColumnKindEnumValues = []CustomTopologyMapInlineEdgeColumnKind{
	CUSTOMTOPOLOGYMAPINLINEEDGECOLUMNKIND_STRING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomTopologyMapInlineEdgeColumnKind) GetAllowedValues() []CustomTopologyMapInlineEdgeColumnKind {
	return allowedCustomTopologyMapInlineEdgeColumnKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomTopologyMapInlineEdgeColumnKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomTopologyMapInlineEdgeColumnKind(value)
	return nil
}

// NewCustomTopologyMapInlineEdgeColumnKindFromValue returns a pointer to a valid CustomTopologyMapInlineEdgeColumnKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomTopologyMapInlineEdgeColumnKindFromValue(v string) (*CustomTopologyMapInlineEdgeColumnKind, error) {
	ev := CustomTopologyMapInlineEdgeColumnKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomTopologyMapInlineEdgeColumnKind: valid values are %v", v, allowedCustomTopologyMapInlineEdgeColumnKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomTopologyMapInlineEdgeColumnKind) IsValid() bool {
	for _, existing := range allowedCustomTopologyMapInlineEdgeColumnKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomTopologyMapInlineEdgeColumnKind value.
func (v CustomTopologyMapInlineEdgeColumnKind) Ptr() *CustomTopologyMapInlineEdgeColumnKind {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomTopologyMapInlineNodeIdColumnKind Test-only generated schema for the dashboard versioning proof of concept.
type CustomTopologyMapInlineNodeIdColumnKind string

// List of CustomTopologyMapInlineNodeIdColumnKind.
const (
	CUSTOMTOPOLOGYMAPINLINENODEIDCOLUMNKIND_STRING CustomTopologyMapInlineNodeIdColumnKind = "string"
)

var allowedCustomTopologyMapInlineNodeIdColumnKindEnumValues = []CustomTopologyMapInlineNodeIdColumnKind{
	CUSTOMTOPOLOGYMAPINLINENODEIDCOLUMNKIND_STRING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomTopologyMapInlineNodeIdColumnKind) GetAllowedValues() []CustomTopologyMapInlineNodeIdColumnKind {
	return allowedCustomTopologyMapInlineNodeIdColumnKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomTopologyMapInlineNodeIdColumnKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomTopologyMapInlineNodeIdColumnKind(value)
	return nil
}

// NewCustomTopologyMapInlineNodeIdColumnKindFromValue returns a pointer to a valid CustomTopologyMapInlineNodeIdColumnKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomTopologyMapInlineNodeIdColumnKindFromValue(v string) (*CustomTopologyMapInlineNodeIdColumnKind, error) {
	ev := CustomTopologyMapInlineNodeIdColumnKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomTopologyMapInlineNodeIdColumnKind: valid values are %v", v, allowedCustomTopologyMapInlineNodeIdColumnKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomTopologyMapInlineNodeIdColumnKind) IsValid() bool {
	for _, existing := range allowedCustomTopologyMapInlineNodeIdColumnKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomTopologyMapInlineNodeIdColumnKind value.
func (v CustomTopologyMapInlineNodeIdColumnKind) Ptr() *CustomTopologyMapInlineNodeIdColumnKind {
	return &v
}

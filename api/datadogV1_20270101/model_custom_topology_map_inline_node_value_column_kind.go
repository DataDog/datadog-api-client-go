// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomTopologyMapInlineNodeValueColumnKind Test-only generated schema for the dashboard versioning proof of concept.
type CustomTopologyMapInlineNodeValueColumnKind string

// List of CustomTopologyMapInlineNodeValueColumnKind.
const (
	CUSTOMTOPOLOGYMAPINLINENODEVALUECOLUMNKIND_SCALAR CustomTopologyMapInlineNodeValueColumnKind = "scalar"
)

var allowedCustomTopologyMapInlineNodeValueColumnKindEnumValues = []CustomTopologyMapInlineNodeValueColumnKind{
	CUSTOMTOPOLOGYMAPINLINENODEVALUECOLUMNKIND_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomTopologyMapInlineNodeValueColumnKind) GetAllowedValues() []CustomTopologyMapInlineNodeValueColumnKind {
	return allowedCustomTopologyMapInlineNodeValueColumnKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomTopologyMapInlineNodeValueColumnKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomTopologyMapInlineNodeValueColumnKind(value)
	return nil
}

// NewCustomTopologyMapInlineNodeValueColumnKindFromValue returns a pointer to a valid CustomTopologyMapInlineNodeValueColumnKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomTopologyMapInlineNodeValueColumnKindFromValue(v string) (*CustomTopologyMapInlineNodeValueColumnKind, error) {
	ev := CustomTopologyMapInlineNodeValueColumnKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomTopologyMapInlineNodeValueColumnKind: valid values are %v", v, allowedCustomTopologyMapInlineNodeValueColumnKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomTopologyMapInlineNodeValueColumnKind) IsValid() bool {
	for _, existing := range allowedCustomTopologyMapInlineNodeValueColumnKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomTopologyMapInlineNodeValueColumnKind value.
func (v CustomTopologyMapInlineNodeValueColumnKind) Ptr() *CustomTopologyMapInlineNodeValueColumnKind {
	return &v
}

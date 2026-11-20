// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomTopologyMapInlineEdgeColumnId Test-only generated schema for the dashboard versioning proof of concept.
type CustomTopologyMapInlineEdgeColumnId string

// List of CustomTopologyMapInlineEdgeColumnId.
const (
	CUSTOMTOPOLOGYMAPINLINEEDGECOLUMNID_TARGET CustomTopologyMapInlineEdgeColumnId = "target"
	CUSTOMTOPOLOGYMAPINLINEEDGECOLUMNID_SOURCE CustomTopologyMapInlineEdgeColumnId = "source"
)

var allowedCustomTopologyMapInlineEdgeColumnIdEnumValues = []CustomTopologyMapInlineEdgeColumnId{
	CUSTOMTOPOLOGYMAPINLINEEDGECOLUMNID_TARGET,
	CUSTOMTOPOLOGYMAPINLINEEDGECOLUMNID_SOURCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomTopologyMapInlineEdgeColumnId) GetAllowedValues() []CustomTopologyMapInlineEdgeColumnId {
	return allowedCustomTopologyMapInlineEdgeColumnIdEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomTopologyMapInlineEdgeColumnId) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomTopologyMapInlineEdgeColumnId(value)
	return nil
}

// NewCustomTopologyMapInlineEdgeColumnIdFromValue returns a pointer to a valid CustomTopologyMapInlineEdgeColumnId
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomTopologyMapInlineEdgeColumnIdFromValue(v string) (*CustomTopologyMapInlineEdgeColumnId, error) {
	ev := CustomTopologyMapInlineEdgeColumnId(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomTopologyMapInlineEdgeColumnId: valid values are %v", v, allowedCustomTopologyMapInlineEdgeColumnIdEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomTopologyMapInlineEdgeColumnId) IsValid() bool {
	for _, existing := range allowedCustomTopologyMapInlineEdgeColumnIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomTopologyMapInlineEdgeColumnId value.
func (v CustomTopologyMapInlineEdgeColumnId) Ptr() *CustomTopologyMapInlineEdgeColumnId {
	return &v
}

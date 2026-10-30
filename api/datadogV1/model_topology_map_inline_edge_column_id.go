// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyMapInlineEdgeColumnId Edge endpoint the column holds.
type TopologyMapInlineEdgeColumnId string

// List of TopologyMapInlineEdgeColumnId.
const (
	TOPOLOGYMAPINLINEEDGECOLUMNID_SOURCE TopologyMapInlineEdgeColumnId = "source"
	TOPOLOGYMAPINLINEEDGECOLUMNID_TARGET TopologyMapInlineEdgeColumnId = "target"
)

var allowedTopologyMapInlineEdgeColumnIdEnumValues = []TopologyMapInlineEdgeColumnId{
	TOPOLOGYMAPINLINEEDGECOLUMNID_SOURCE,
	TOPOLOGYMAPINLINEEDGECOLUMNID_TARGET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TopologyMapInlineEdgeColumnId) GetAllowedValues() []TopologyMapInlineEdgeColumnId {
	return allowedTopologyMapInlineEdgeColumnIdEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TopologyMapInlineEdgeColumnId) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TopologyMapInlineEdgeColumnId(value)
	return nil
}

// NewTopologyMapInlineEdgeColumnIdFromValue returns a pointer to a valid TopologyMapInlineEdgeColumnId
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTopologyMapInlineEdgeColumnIdFromValue(v string) (*TopologyMapInlineEdgeColumnId, error) {
	ev := TopologyMapInlineEdgeColumnId(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TopologyMapInlineEdgeColumnId: valid values are %v", v, allowedTopologyMapInlineEdgeColumnIdEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TopologyMapInlineEdgeColumnId) IsValid() bool {
	for _, existing := range allowedTopologyMapInlineEdgeColumnIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TopologyMapInlineEdgeColumnId value.
func (v TopologyMapInlineEdgeColumnId) Ptr() *TopologyMapInlineEdgeColumnId {
	return &v
}

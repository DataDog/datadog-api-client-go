// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyMapInlineNodeIdColumnId ID of the node ID column.
type TopologyMapInlineNodeIdColumnId string

// List of TopologyMapInlineNodeIdColumnId.
const (
	TOPOLOGYMAPINLINENODEIDCOLUMNID_ID TopologyMapInlineNodeIdColumnId = "id"
)

var allowedTopologyMapInlineNodeIdColumnIdEnumValues = []TopologyMapInlineNodeIdColumnId{
	TOPOLOGYMAPINLINENODEIDCOLUMNID_ID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TopologyMapInlineNodeIdColumnId) GetAllowedValues() []TopologyMapInlineNodeIdColumnId {
	return allowedTopologyMapInlineNodeIdColumnIdEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TopologyMapInlineNodeIdColumnId) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TopologyMapInlineNodeIdColumnId(value)
	return nil
}

// NewTopologyMapInlineNodeIdColumnIdFromValue returns a pointer to a valid TopologyMapInlineNodeIdColumnId
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTopologyMapInlineNodeIdColumnIdFromValue(v string) (*TopologyMapInlineNodeIdColumnId, error) {
	ev := TopologyMapInlineNodeIdColumnId(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TopologyMapInlineNodeIdColumnId: valid values are %v", v, allowedTopologyMapInlineNodeIdColumnIdEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TopologyMapInlineNodeIdColumnId) IsValid() bool {
	for _, existing := range allowedTopologyMapInlineNodeIdColumnIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TopologyMapInlineNodeIdColumnId value.
func (v TopologyMapInlineNodeIdColumnId) Ptr() *TopologyMapInlineNodeIdColumnId {
	return &v
}

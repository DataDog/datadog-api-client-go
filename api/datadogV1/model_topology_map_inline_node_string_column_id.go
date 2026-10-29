// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyMapInlineNodeStringColumnId Node field the column fills. `title` is bold text at the top of the node; `subtitle` is smaller text below it.
type TopologyMapInlineNodeStringColumnId string

// List of TopologyMapInlineNodeStringColumnId.
const (
	TOPOLOGYMAPINLINENODESTRINGCOLUMNID_TITLE    TopologyMapInlineNodeStringColumnId = "title"
	TOPOLOGYMAPINLINENODESTRINGCOLUMNID_SUBTITLE TopologyMapInlineNodeStringColumnId = "subtitle"
)

var allowedTopologyMapInlineNodeStringColumnIdEnumValues = []TopologyMapInlineNodeStringColumnId{
	TOPOLOGYMAPINLINENODESTRINGCOLUMNID_TITLE,
	TOPOLOGYMAPINLINENODESTRINGCOLUMNID_SUBTITLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TopologyMapInlineNodeStringColumnId) GetAllowedValues() []TopologyMapInlineNodeStringColumnId {
	return allowedTopologyMapInlineNodeStringColumnIdEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TopologyMapInlineNodeStringColumnId) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TopologyMapInlineNodeStringColumnId(value)
	return nil
}

// NewTopologyMapInlineNodeStringColumnIdFromValue returns a pointer to a valid TopologyMapInlineNodeStringColumnId
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTopologyMapInlineNodeStringColumnIdFromValue(v string) (*TopologyMapInlineNodeStringColumnId, error) {
	ev := TopologyMapInlineNodeStringColumnId(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TopologyMapInlineNodeStringColumnId: valid values are %v", v, allowedTopologyMapInlineNodeStringColumnIdEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TopologyMapInlineNodeStringColumnId) IsValid() bool {
	for _, existing := range allowedTopologyMapInlineNodeStringColumnIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TopologyMapInlineNodeStringColumnId value.
func (v TopologyMapInlineNodeStringColumnId) Ptr() *TopologyMapInlineNodeStringColumnId {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyRequestCustomProjectionType Projection type. Each result row is a node.
type TopologyRequestCustomProjectionType string

// List of TopologyRequestCustomProjectionType.
const (
	TOPOLOGYREQUESTCUSTOMPROJECTIONTYPE_CUSTOM_TOPOLOGY_MAP_NODE TopologyRequestCustomProjectionType = "custom_topology_map_node"
)

var allowedTopologyRequestCustomProjectionTypeEnumValues = []TopologyRequestCustomProjectionType{
	TOPOLOGYREQUESTCUSTOMPROJECTIONTYPE_CUSTOM_TOPOLOGY_MAP_NODE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TopologyRequestCustomProjectionType) GetAllowedValues() []TopologyRequestCustomProjectionType {
	return allowedTopologyRequestCustomProjectionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TopologyRequestCustomProjectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TopologyRequestCustomProjectionType(value)
	return nil
}

// NewTopologyRequestCustomProjectionTypeFromValue returns a pointer to a valid TopologyRequestCustomProjectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTopologyRequestCustomProjectionTypeFromValue(v string) (*TopologyRequestCustomProjectionType, error) {
	ev := TopologyRequestCustomProjectionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TopologyRequestCustomProjectionType: valid values are %v", v, allowedTopologyRequestCustomProjectionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TopologyRequestCustomProjectionType) IsValid() bool {
	for _, existing := range allowedTopologyRequestCustomProjectionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TopologyRequestCustomProjectionType value.
func (v TopologyRequestCustomProjectionType) Ptr() *TopologyRequestCustomProjectionType {
	return &v
}

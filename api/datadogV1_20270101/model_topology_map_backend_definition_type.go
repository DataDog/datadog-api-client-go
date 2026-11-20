// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyMapBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type TopologyMapBackendDefinitionType string

// List of TopologyMapBackendDefinitionType.
const (
	TOPOLOGYMAPBACKENDDEFINITIONTYPE_TOPOLOGY_MAP TopologyMapBackendDefinitionType = "topology_map"
)

var allowedTopologyMapBackendDefinitionTypeEnumValues = []TopologyMapBackendDefinitionType{
	TOPOLOGYMAPBACKENDDEFINITIONTYPE_TOPOLOGY_MAP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TopologyMapBackendDefinitionType) GetAllowedValues() []TopologyMapBackendDefinitionType {
	return allowedTopologyMapBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TopologyMapBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TopologyMapBackendDefinitionType(value)
	return nil
}

// NewTopologyMapBackendDefinitionTypeFromValue returns a pointer to a valid TopologyMapBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTopologyMapBackendDefinitionTypeFromValue(v string) (*TopologyMapBackendDefinitionType, error) {
	ev := TopologyMapBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TopologyMapBackendDefinitionType: valid values are %v", v, allowedTopologyMapBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TopologyMapBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedTopologyMapBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TopologyMapBackendDefinitionType value.
func (v TopologyMapBackendDefinitionType) Ptr() *TopologyMapBackendDefinitionType {
	return &v
}

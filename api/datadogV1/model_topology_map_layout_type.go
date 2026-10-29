// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyMapLayoutType Type of layout algorithm.
type TopologyMapLayoutType string

// List of TopologyMapLayoutType.
const (
	TOPOLOGYMAPLAYOUTTYPE_HIERARCHICAL TopologyMapLayoutType = "hierarchical"
)

var allowedTopologyMapLayoutTypeEnumValues = []TopologyMapLayoutType{
	TOPOLOGYMAPLAYOUTTYPE_HIERARCHICAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TopologyMapLayoutType) GetAllowedValues() []TopologyMapLayoutType {
	return allowedTopologyMapLayoutTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TopologyMapLayoutType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TopologyMapLayoutType(value)
	return nil
}

// NewTopologyMapLayoutTypeFromValue returns a pointer to a valid TopologyMapLayoutType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTopologyMapLayoutTypeFromValue(v string) (*TopologyMapLayoutType, error) {
	ev := TopologyMapLayoutType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TopologyMapLayoutType: valid values are %v", v, allowedTopologyMapLayoutTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TopologyMapLayoutType) IsValid() bool {
	for _, existing := range allowedTopologyMapLayoutTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TopologyMapLayoutType value.
func (v TopologyMapLayoutType) Ptr() *TopologyMapLayoutType {
	return &v
}

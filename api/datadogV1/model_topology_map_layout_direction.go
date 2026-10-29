// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyMapLayoutDirection Direction of ranks. For example, `top-to-bottom` places first-rank nodes at the top and last-rank nodes at the bottom.
type TopologyMapLayoutDirection string

// List of TopologyMapLayoutDirection.
const (
	TOPOLOGYMAPLAYOUTDIRECTION_TOP_TO_BOTTOM TopologyMapLayoutDirection = "top-to-bottom"
	TOPOLOGYMAPLAYOUTDIRECTION_BOTTOM_TO_TOP TopologyMapLayoutDirection = "bottom-to-top"
	TOPOLOGYMAPLAYOUTDIRECTION_LEFT_TO_RIGHT TopologyMapLayoutDirection = "left-to-right"
	TOPOLOGYMAPLAYOUTDIRECTION_RIGHT_TO_LEFT TopologyMapLayoutDirection = "right-to-left"
)

var allowedTopologyMapLayoutDirectionEnumValues = []TopologyMapLayoutDirection{
	TOPOLOGYMAPLAYOUTDIRECTION_TOP_TO_BOTTOM,
	TOPOLOGYMAPLAYOUTDIRECTION_BOTTOM_TO_TOP,
	TOPOLOGYMAPLAYOUTDIRECTION_LEFT_TO_RIGHT,
	TOPOLOGYMAPLAYOUTDIRECTION_RIGHT_TO_LEFT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TopologyMapLayoutDirection) GetAllowedValues() []TopologyMapLayoutDirection {
	return allowedTopologyMapLayoutDirectionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TopologyMapLayoutDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TopologyMapLayoutDirection(value)
	return nil
}

// NewTopologyMapLayoutDirectionFromValue returns a pointer to a valid TopologyMapLayoutDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTopologyMapLayoutDirectionFromValue(v string) (*TopologyMapLayoutDirection, error) {
	ev := TopologyMapLayoutDirection(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TopologyMapLayoutDirection: valid values are %v", v, allowedTopologyMapLayoutDirectionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TopologyMapLayoutDirection) IsValid() bool {
	for _, existing := range allowedTopologyMapLayoutDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TopologyMapLayoutDirection value.
func (v TopologyMapLayoutDirection) Ptr() *TopologyMapLayoutDirection {
	return &v
}

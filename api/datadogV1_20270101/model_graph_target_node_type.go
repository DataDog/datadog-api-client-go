// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GraphTargetNodeType Test-only generated schema for the dashboard versioning proof of concept.
type GraphTargetNodeType string

// List of GraphTargetNodeType.
const (
	GRAPHTARGETNODETYPE_NODE GraphTargetNodeType = "node"
)

var allowedGraphTargetNodeTypeEnumValues = []GraphTargetNodeType{
	GRAPHTARGETNODETYPE_NODE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GraphTargetNodeType) GetAllowedValues() []GraphTargetNodeType {
	return allowedGraphTargetNodeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GraphTargetNodeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GraphTargetNodeType(value)
	return nil
}

// NewGraphTargetNodeTypeFromValue returns a pointer to a valid GraphTargetNodeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGraphTargetNodeTypeFromValue(v string) (*GraphTargetNodeType, error) {
	ev := GraphTargetNodeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GraphTargetNodeType: valid values are %v", v, allowedGraphTargetNodeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GraphTargetNodeType) IsValid() bool {
	for _, existing := range allowedGraphTargetNodeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GraphTargetNodeType value.
func (v GraphTargetNodeType) Ptr() *GraphTargetNodeType {
	return &v
}

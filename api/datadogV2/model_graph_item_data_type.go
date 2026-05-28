// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GraphItemDataType Graph items resource type.
type GraphItemDataType string

// List of GraphItemDataType.
const (
	GRAPHITEMDATATYPE_GRAPH_ITEMS GraphItemDataType = "graph_items"
)

var allowedGraphItemDataTypeEnumValues = []GraphItemDataType{
	GRAPHITEMDATATYPE_GRAPH_ITEMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GraphItemDataType) GetAllowedValues() []GraphItemDataType {
	return allowedGraphItemDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GraphItemDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GraphItemDataType(value)
	return nil
}

// NewGraphItemDataTypeFromValue returns a pointer to a valid GraphItemDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGraphItemDataTypeFromValue(v string) (*GraphItemDataType, error) {
	ev := GraphItemDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GraphItemDataType: valid values are %v", v, allowedGraphItemDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GraphItemDataType) IsValid() bool {
	for _, existing := range allowedGraphItemDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GraphItemDataType value.
func (v GraphItemDataType) Ptr() *GraphItemDataType {
	return &v
}

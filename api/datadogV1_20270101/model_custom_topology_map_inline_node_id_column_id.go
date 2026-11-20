// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomTopologyMapInlineNodeIdColumnId Test-only generated schema for the dashboard versioning proof of concept.
type CustomTopologyMapInlineNodeIdColumnId string

// List of CustomTopologyMapInlineNodeIdColumnId.
const (
	CUSTOMTOPOLOGYMAPINLINENODEIDCOLUMNID_ID CustomTopologyMapInlineNodeIdColumnId = "id"
)

var allowedCustomTopologyMapInlineNodeIdColumnIdEnumValues = []CustomTopologyMapInlineNodeIdColumnId{
	CUSTOMTOPOLOGYMAPINLINENODEIDCOLUMNID_ID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomTopologyMapInlineNodeIdColumnId) GetAllowedValues() []CustomTopologyMapInlineNodeIdColumnId {
	return allowedCustomTopologyMapInlineNodeIdColumnIdEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomTopologyMapInlineNodeIdColumnId) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomTopologyMapInlineNodeIdColumnId(value)
	return nil
}

// NewCustomTopologyMapInlineNodeIdColumnIdFromValue returns a pointer to a valid CustomTopologyMapInlineNodeIdColumnId
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomTopologyMapInlineNodeIdColumnIdFromValue(v string) (*CustomTopologyMapInlineNodeIdColumnId, error) {
	ev := CustomTopologyMapInlineNodeIdColumnId(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomTopologyMapInlineNodeIdColumnId: valid values are %v", v, allowedCustomTopologyMapInlineNodeIdColumnIdEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomTopologyMapInlineNodeIdColumnId) IsValid() bool {
	for _, existing := range allowedCustomTopologyMapInlineNodeIdColumnIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomTopologyMapInlineNodeIdColumnId value.
func (v CustomTopologyMapInlineNodeIdColumnId) Ptr() *CustomTopologyMapInlineNodeIdColumnId {
	return &v
}

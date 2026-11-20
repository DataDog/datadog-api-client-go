// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomTopologyMapInlineNodeStringColumnId Test-only generated schema for the dashboard versioning proof of concept.
type CustomTopologyMapInlineNodeStringColumnId string

// List of CustomTopologyMapInlineNodeStringColumnId.
const (
	CUSTOMTOPOLOGYMAPINLINENODESTRINGCOLUMNID_TITLE    CustomTopologyMapInlineNodeStringColumnId = "title"
	CUSTOMTOPOLOGYMAPINLINENODESTRINGCOLUMNID_SUBTITLE CustomTopologyMapInlineNodeStringColumnId = "subtitle"
)

var allowedCustomTopologyMapInlineNodeStringColumnIdEnumValues = []CustomTopologyMapInlineNodeStringColumnId{
	CUSTOMTOPOLOGYMAPINLINENODESTRINGCOLUMNID_TITLE,
	CUSTOMTOPOLOGYMAPINLINENODESTRINGCOLUMNID_SUBTITLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomTopologyMapInlineNodeStringColumnId) GetAllowedValues() []CustomTopologyMapInlineNodeStringColumnId {
	return allowedCustomTopologyMapInlineNodeStringColumnIdEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomTopologyMapInlineNodeStringColumnId) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomTopologyMapInlineNodeStringColumnId(value)
	return nil
}

// NewCustomTopologyMapInlineNodeStringColumnIdFromValue returns a pointer to a valid CustomTopologyMapInlineNodeStringColumnId
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomTopologyMapInlineNodeStringColumnIdFromValue(v string) (*CustomTopologyMapInlineNodeStringColumnId, error) {
	ev := CustomTopologyMapInlineNodeStringColumnId(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomTopologyMapInlineNodeStringColumnId: valid values are %v", v, allowedCustomTopologyMapInlineNodeStringColumnIdEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomTopologyMapInlineNodeStringColumnId) IsValid() bool {
	for _, existing := range allowedCustomTopologyMapInlineNodeStringColumnIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomTopologyMapInlineNodeStringColumnId value.
func (v CustomTopologyMapInlineNodeStringColumnId) Ptr() *CustomTopologyMapInlineNodeStringColumnId {
	return &v
}

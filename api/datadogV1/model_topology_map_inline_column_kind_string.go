// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyMapInlineColumnKindString Column of strings.
type TopologyMapInlineColumnKindString string

// List of TopologyMapInlineColumnKindString.
const (
	TOPOLOGYMAPINLINECOLUMNKINDSTRING_STRING TopologyMapInlineColumnKindString = "string"
)

var allowedTopologyMapInlineColumnKindStringEnumValues = []TopologyMapInlineColumnKindString{
	TOPOLOGYMAPINLINECOLUMNKINDSTRING_STRING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TopologyMapInlineColumnKindString) GetAllowedValues() []TopologyMapInlineColumnKindString {
	return allowedTopologyMapInlineColumnKindStringEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TopologyMapInlineColumnKindString) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TopologyMapInlineColumnKindString(value)
	return nil
}

// NewTopologyMapInlineColumnKindStringFromValue returns a pointer to a valid TopologyMapInlineColumnKindString
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTopologyMapInlineColumnKindStringFromValue(v string) (*TopologyMapInlineColumnKindString, error) {
	ev := TopologyMapInlineColumnKindString(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TopologyMapInlineColumnKindString: valid values are %v", v, allowedTopologyMapInlineColumnKindStringEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TopologyMapInlineColumnKindString) IsValid() bool {
	for _, existing := range allowedTopologyMapInlineColumnKindStringEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TopologyMapInlineColumnKindString value.
func (v TopologyMapInlineColumnKindString) Ptr() *TopologyMapInlineColumnKindString {
	return &v
}

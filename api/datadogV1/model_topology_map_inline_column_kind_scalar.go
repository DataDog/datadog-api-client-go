// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyMapInlineColumnKindScalar Column of numbers.
type TopologyMapInlineColumnKindScalar string

// List of TopologyMapInlineColumnKindScalar.
const (
	TOPOLOGYMAPINLINECOLUMNKINDSCALAR_SCALAR TopologyMapInlineColumnKindScalar = "scalar"
)

var allowedTopologyMapInlineColumnKindScalarEnumValues = []TopologyMapInlineColumnKindScalar{
	TOPOLOGYMAPINLINECOLUMNKINDSCALAR_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TopologyMapInlineColumnKindScalar) GetAllowedValues() []TopologyMapInlineColumnKindScalar {
	return allowedTopologyMapInlineColumnKindScalarEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TopologyMapInlineColumnKindScalar) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TopologyMapInlineColumnKindScalar(value)
	return nil
}

// NewTopologyMapInlineColumnKindScalarFromValue returns a pointer to a valid TopologyMapInlineColumnKindScalar
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTopologyMapInlineColumnKindScalarFromValue(v string) (*TopologyMapInlineColumnKindScalar, error) {
	ev := TopologyMapInlineColumnKindScalar(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TopologyMapInlineColumnKindScalar: valid values are %v", v, allowedTopologyMapInlineColumnKindScalarEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TopologyMapInlineColumnKindScalar) IsValid() bool {
	for _, existing := range allowedTopologyMapInlineColumnKindScalarEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TopologyMapInlineColumnKindScalar value.
func (v TopologyMapInlineColumnKindScalar) Ptr() *TopologyMapInlineColumnKindScalar {
	return &v
}

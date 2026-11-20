// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostmapDimension Test-only generated schema for the dashboard versioning proof of concept.
type HostmapDimension string

// List of HostmapDimension.
const (
	HOSTMAPDIMENSION_GROUP HostmapDimension = "group"
	HOSTMAPDIMENSION_NODE  HostmapDimension = "node"
	HOSTMAPDIMENSION_FILL  HostmapDimension = "fill"
	HOSTMAPDIMENSION_SIZE  HostmapDimension = "size"
)

var allowedHostmapDimensionEnumValues = []HostmapDimension{
	HOSTMAPDIMENSION_GROUP,
	HOSTMAPDIMENSION_NODE,
	HOSTMAPDIMENSION_FILL,
	HOSTMAPDIMENSION_SIZE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HostmapDimension) GetAllowedValues() []HostmapDimension {
	return allowedHostmapDimensionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HostmapDimension) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HostmapDimension(value)
	return nil
}

// NewHostmapDimensionFromValue returns a pointer to a valid HostmapDimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHostmapDimensionFromValue(v string) (*HostmapDimension, error) {
	ev := HostmapDimension(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HostmapDimension: valid values are %v", v, allowedHostmapDimensionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HostmapDimension) IsValid() bool {
	for _, existing := range allowedHostmapDimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostmapDimension value.
func (v HostmapDimension) Ptr() *HostmapDimension {
	return &v
}

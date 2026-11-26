// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostmapProjectionType Test-only generated schema for the dashboard versioning proof of concept.
type HostmapProjectionType string

// List of HostmapProjectionType.
const (
	HOSTMAPPROJECTIONTYPE_HOSTMAP HostmapProjectionType = "hostmap"
)

var allowedHostmapProjectionTypeEnumValues = []HostmapProjectionType{
	HOSTMAPPROJECTIONTYPE_HOSTMAP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HostmapProjectionType) GetAllowedValues() []HostmapProjectionType {
	return allowedHostmapProjectionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HostmapProjectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HostmapProjectionType(value)
	return nil
}

// NewHostmapProjectionTypeFromValue returns a pointer to a valid HostmapProjectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHostmapProjectionTypeFromValue(v string) (*HostmapProjectionType, error) {
	ev := HostmapProjectionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HostmapProjectionType: valid values are %v", v, allowedHostmapProjectionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HostmapProjectionType) IsValid() bool {
	for _, existing := range allowedHostmapProjectionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostmapProjectionType value.
func (v HostmapProjectionType) Ptr() *HostmapProjectionType {
	return &v
}

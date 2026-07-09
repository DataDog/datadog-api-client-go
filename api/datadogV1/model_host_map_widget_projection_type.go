// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostMapWidgetProjectionType Type of the host map projection.
type HostMapWidgetProjectionType string

// List of HostMapWidgetProjectionType.
const (
	HOSTMAPWIDGETPROJECTIONTYPE_HOSTMAP HostMapWidgetProjectionType = "hostmap"
)

var allowedHostMapWidgetProjectionTypeEnumValues = []HostMapWidgetProjectionType{
	HOSTMAPWIDGETPROJECTIONTYPE_HOSTMAP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HostMapWidgetProjectionType) GetAllowedValues() []HostMapWidgetProjectionType {
	return allowedHostMapWidgetProjectionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HostMapWidgetProjectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HostMapWidgetProjectionType(value)
	return nil
}

// NewHostMapWidgetProjectionTypeFromValue returns a pointer to a valid HostMapWidgetProjectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHostMapWidgetProjectionTypeFromValue(v string) (*HostMapWidgetProjectionType, error) {
	ev := HostMapWidgetProjectionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HostMapWidgetProjectionType: valid values are %v", v, allowedHostMapWidgetProjectionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HostMapWidgetProjectionType) IsValid() bool {
	for _, existing := range allowedHostMapWidgetProjectionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostMapWidgetProjectionType value.
func (v HostMapWidgetProjectionType) Ptr() *HostMapWidgetProjectionType {
	return &v
}

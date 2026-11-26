// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostmapDataProjectionRequestRequestType Test-only generated schema for the dashboard versioning proof of concept.
type HostmapDataProjectionRequestRequestType string

// List of HostmapDataProjectionRequestRequestType.
const (
	HOSTMAPDATAPROJECTIONREQUESTREQUESTTYPE_DATA_PROJECTION HostmapDataProjectionRequestRequestType = "data_projection"
)

var allowedHostmapDataProjectionRequestRequestTypeEnumValues = []HostmapDataProjectionRequestRequestType{
	HOSTMAPDATAPROJECTIONREQUESTREQUESTTYPE_DATA_PROJECTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HostmapDataProjectionRequestRequestType) GetAllowedValues() []HostmapDataProjectionRequestRequestType {
	return allowedHostmapDataProjectionRequestRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HostmapDataProjectionRequestRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HostmapDataProjectionRequestRequestType(value)
	return nil
}

// NewHostmapDataProjectionRequestRequestTypeFromValue returns a pointer to a valid HostmapDataProjectionRequestRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHostmapDataProjectionRequestRequestTypeFromValue(v string) (*HostmapDataProjectionRequestRequestType, error) {
	ev := HostmapDataProjectionRequestRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HostmapDataProjectionRequestRequestType: valid values are %v", v, allowedHostmapDataProjectionRequestRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HostmapDataProjectionRequestRequestType) IsValid() bool {
	for _, existing := range allowedHostmapDataProjectionRequestRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostmapDataProjectionRequestRequestType value.
func (v HostmapDataProjectionRequestRequestType) Ptr() *HostmapDataProjectionRequestRequestType {
	return &v
}

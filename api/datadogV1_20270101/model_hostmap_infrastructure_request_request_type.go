// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostmapInfrastructureRequestRequestType Test-only generated schema for the dashboard versioning proof of concept.
type HostmapInfrastructureRequestRequestType string

// List of HostmapInfrastructureRequestRequestType.
const (
	HOSTMAPINFRASTRUCTUREREQUESTREQUESTTYPE_INFRASTRUCTURE_HOSTMAP HostmapInfrastructureRequestRequestType = "infrastructure_hostmap"
)

var allowedHostmapInfrastructureRequestRequestTypeEnumValues = []HostmapInfrastructureRequestRequestType{
	HOSTMAPINFRASTRUCTUREREQUESTREQUESTTYPE_INFRASTRUCTURE_HOSTMAP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HostmapInfrastructureRequestRequestType) GetAllowedValues() []HostmapInfrastructureRequestRequestType {
	return allowedHostmapInfrastructureRequestRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HostmapInfrastructureRequestRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HostmapInfrastructureRequestRequestType(value)
	return nil
}

// NewHostmapInfrastructureRequestRequestTypeFromValue returns a pointer to a valid HostmapInfrastructureRequestRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHostmapInfrastructureRequestRequestTypeFromValue(v string) (*HostmapInfrastructureRequestRequestType, error) {
	ev := HostmapInfrastructureRequestRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HostmapInfrastructureRequestRequestType: valid values are %v", v, allowedHostmapInfrastructureRequestRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HostmapInfrastructureRequestRequestType) IsValid() bool {
	for _, existing := range allowedHostmapInfrastructureRequestRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostmapInfrastructureRequestRequestType value.
func (v HostmapInfrastructureRequestRequestType) Ptr() *HostmapInfrastructureRequestRequestType {
	return &v
}

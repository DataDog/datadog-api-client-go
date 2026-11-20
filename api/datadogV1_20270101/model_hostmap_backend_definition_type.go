// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostmapBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type HostmapBackendDefinitionType string

// List of HostmapBackendDefinitionType.
const (
	HOSTMAPBACKENDDEFINITIONTYPE_HOSTMAP HostmapBackendDefinitionType = "hostmap"
)

var allowedHostmapBackendDefinitionTypeEnumValues = []HostmapBackendDefinitionType{
	HOSTMAPBACKENDDEFINITIONTYPE_HOSTMAP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HostmapBackendDefinitionType) GetAllowedValues() []HostmapBackendDefinitionType {
	return allowedHostmapBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HostmapBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HostmapBackendDefinitionType(value)
	return nil
}

// NewHostmapBackendDefinitionTypeFromValue returns a pointer to a valid HostmapBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHostmapBackendDefinitionTypeFromValue(v string) (*HostmapBackendDefinitionType, error) {
	ev := HostmapBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HostmapBackendDefinitionType: valid values are %v", v, allowedHostmapBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HostmapBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedHostmapBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostmapBackendDefinitionType value.
func (v HostmapBackendDefinitionType) Ptr() *HostmapBackendDefinitionType {
	return &v
}

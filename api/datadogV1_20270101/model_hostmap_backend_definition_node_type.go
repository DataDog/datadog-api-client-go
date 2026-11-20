// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostmapBackendDefinitionNodeType Test-only generated schema for the dashboard versioning proof of concept.
type HostmapBackendDefinitionNodeType string

// List of HostmapBackendDefinitionNodeType.
const (
	HOSTMAPBACKENDDEFINITIONNODETYPE_CONTAINER HostmapBackendDefinitionNodeType = "container"
	HOSTMAPBACKENDDEFINITIONNODETYPE_HOST      HostmapBackendDefinitionNodeType = "host"
)

var allowedHostmapBackendDefinitionNodeTypeEnumValues = []HostmapBackendDefinitionNodeType{
	HOSTMAPBACKENDDEFINITIONNODETYPE_CONTAINER,
	HOSTMAPBACKENDDEFINITIONNODETYPE_HOST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HostmapBackendDefinitionNodeType) GetAllowedValues() []HostmapBackendDefinitionNodeType {
	return allowedHostmapBackendDefinitionNodeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HostmapBackendDefinitionNodeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HostmapBackendDefinitionNodeType(value)
	return nil
}

// NewHostmapBackendDefinitionNodeTypeFromValue returns a pointer to a valid HostmapBackendDefinitionNodeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHostmapBackendDefinitionNodeTypeFromValue(v string) (*HostmapBackendDefinitionNodeType, error) {
	ev := HostmapBackendDefinitionNodeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HostmapBackendDefinitionNodeType: valid values are %v", v, allowedHostmapBackendDefinitionNodeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HostmapBackendDefinitionNodeType) IsValid() bool {
	for _, existing := range allowedHostmapBackendDefinitionNodeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostmapBackendDefinitionNodeType value.
func (v HostmapBackendDefinitionNodeType) Ptr() *HostmapBackendDefinitionNodeType {
	return &v
}

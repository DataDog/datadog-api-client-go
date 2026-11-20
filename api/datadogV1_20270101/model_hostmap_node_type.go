// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostmapNodeType Test-only generated schema for the dashboard versioning proof of concept.
type HostmapNodeType string

// List of HostmapNodeType.
const (
	HOSTMAPNODETYPE_CLUSTER   HostmapNodeType = "cluster"
	HOSTMAPNODETYPE_CONTAINER HostmapNodeType = "container"
	HOSTMAPNODETYPE_HOST      HostmapNodeType = "host"
	HOSTMAPNODETYPE_POD       HostmapNodeType = "pod"
)

var allowedHostmapNodeTypeEnumValues = []HostmapNodeType{
	HOSTMAPNODETYPE_CLUSTER,
	HOSTMAPNODETYPE_CONTAINER,
	HOSTMAPNODETYPE_HOST,
	HOSTMAPNODETYPE_POD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HostmapNodeType) GetAllowedValues() []HostmapNodeType {
	return allowedHostmapNodeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HostmapNodeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HostmapNodeType(value)
	return nil
}

// NewHostmapNodeTypeFromValue returns a pointer to a valid HostmapNodeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHostmapNodeTypeFromValue(v string) (*HostmapNodeType, error) {
	ev := HostmapNodeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HostmapNodeType: valid values are %v", v, allowedHostmapNodeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HostmapNodeType) IsValid() bool {
	for _, existing := range allowedHostmapNodeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostmapNodeType value.
func (v HostmapNodeType) Ptr() *HostmapNodeType {
	return &v
}

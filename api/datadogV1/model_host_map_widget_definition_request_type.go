// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostMapWidgetDefinitionRequestType Identifies which host map request format the sibling fields on `HostMapWidgetDefinitionRequests` describe: an infrastructure-backed request or a DDSQL published-dataset request.
type HostMapWidgetDefinitionRequestType string

// List of HostMapWidgetDefinitionRequestType.
const (
	HOSTMAPWIDGETDEFINITIONREQUESTTYPE_INFRASTRUCTURE_HOSTMAP HostMapWidgetDefinitionRequestType = "infrastructure_hostmap"
	HOSTMAPWIDGETDEFINITIONREQUESTTYPE_DATA_PROJECTION        HostMapWidgetDefinitionRequestType = "data_projection"
)

var allowedHostMapWidgetDefinitionRequestTypeEnumValues = []HostMapWidgetDefinitionRequestType{
	HOSTMAPWIDGETDEFINITIONREQUESTTYPE_INFRASTRUCTURE_HOSTMAP,
	HOSTMAPWIDGETDEFINITIONREQUESTTYPE_DATA_PROJECTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HostMapWidgetDefinitionRequestType) GetAllowedValues() []HostMapWidgetDefinitionRequestType {
	return allowedHostMapWidgetDefinitionRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HostMapWidgetDefinitionRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HostMapWidgetDefinitionRequestType(value)
	return nil
}

// NewHostMapWidgetDefinitionRequestTypeFromValue returns a pointer to a valid HostMapWidgetDefinitionRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHostMapWidgetDefinitionRequestTypeFromValue(v string) (*HostMapWidgetDefinitionRequestType, error) {
	ev := HostMapWidgetDefinitionRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HostMapWidgetDefinitionRequestType: valid values are %v", v, allowedHostMapWidgetDefinitionRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HostMapWidgetDefinitionRequestType) IsValid() bool {
	for _, existing := range allowedHostMapWidgetDefinitionRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostMapWidgetDefinitionRequestType value.
func (v HostMapWidgetDefinitionRequestType) Ptr() *HostMapWidgetDefinitionRequestType {
	return &v
}

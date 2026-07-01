// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTypeSlugSource When set to `servicenow`, incidents will display the ServiceNow record ID instead of the public ID. If no ServiceNow integration exists, the public ID will be displayed.
type IncidentTypeSlugSource string

// List of IncidentTypeSlugSource.
const (
	INCIDENTTYPESLUGSOURCE_DEFAULT    IncidentTypeSlugSource = "default"
	INCIDENTTYPESLUGSOURCE_SERVICENOW IncidentTypeSlugSource = "servicenow"
)

var allowedIncidentTypeSlugSourceEnumValues = []IncidentTypeSlugSource{
	INCIDENTTYPESLUGSOURCE_DEFAULT,
	INCIDENTTYPESLUGSOURCE_SERVICENOW,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentTypeSlugSource) GetAllowedValues() []IncidentTypeSlugSource {
	return allowedIncidentTypeSlugSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentTypeSlugSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentTypeSlugSource(value)
	return nil
}

// NewIncidentTypeSlugSourceFromValue returns a pointer to a valid IncidentTypeSlugSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentTypeSlugSourceFromValue(v string) (*IncidentTypeSlugSource, error) {
	ev := IncidentTypeSlugSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentTypeSlugSource: valid values are %v", v, allowedIncidentTypeSlugSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentTypeSlugSource) IsValid() bool {
	for _, existing := range allowedIncidentTypeSlugSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentTypeSlugSource value.
func (v IncidentTypeSlugSource) Ptr() *IncidentTypeSlugSource {
	return &v
}

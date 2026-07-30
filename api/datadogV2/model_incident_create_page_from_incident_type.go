// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentCreatePageFromIncidentType Resource type for a page creation request.
type IncidentCreatePageFromIncidentType string

// List of IncidentCreatePageFromIncidentType.
const (
	INCIDENTCREATEPAGEFROMINCIDENTTYPE_PAGE IncidentCreatePageFromIncidentType = "page"
)

var allowedIncidentCreatePageFromIncidentTypeEnumValues = []IncidentCreatePageFromIncidentType{
	INCIDENTCREATEPAGEFROMINCIDENTTYPE_PAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentCreatePageFromIncidentType) GetAllowedValues() []IncidentCreatePageFromIncidentType {
	return allowedIncidentCreatePageFromIncidentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentCreatePageFromIncidentType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentCreatePageFromIncidentType(value)
	return nil
}

// NewIncidentCreatePageFromIncidentTypeFromValue returns a pointer to a valid IncidentCreatePageFromIncidentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentCreatePageFromIncidentTypeFromValue(v string) (*IncidentCreatePageFromIncidentType, error) {
	ev := IncidentCreatePageFromIncidentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentCreatePageFromIncidentType: valid values are %v", v, allowedIncidentCreatePageFromIncidentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentCreatePageFromIncidentType) IsValid() bool {
	for _, existing := range allowedIncidentCreatePageFromIncidentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentCreatePageFromIncidentType value.
func (v IncidentCreatePageFromIncidentType) Ptr() *IncidentCreatePageFromIncidentType {
	return &v
}

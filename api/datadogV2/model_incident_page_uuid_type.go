// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentPageUUIDType Resource type for a page UUID response.
type IncidentPageUUIDType string

// List of IncidentPageUUIDType.
const (
	INCIDENTPAGEUUIDTYPE_PAGE_UUID IncidentPageUUIDType = "page_uuid"
)

var allowedIncidentPageUUIDTypeEnumValues = []IncidentPageUUIDType{
	INCIDENTPAGEUUIDTYPE_PAGE_UUID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentPageUUIDType) GetAllowedValues() []IncidentPageUUIDType {
	return allowedIncidentPageUUIDTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentPageUUIDType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentPageUUIDType(value)
	return nil
}

// NewIncidentPageUUIDTypeFromValue returns a pointer to a valid IncidentPageUUIDType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentPageUUIDTypeFromValue(v string) (*IncidentPageUUIDType, error) {
	ev := IncidentPageUUIDType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentPageUUIDType: valid values are %v", v, allowedIncidentPageUUIDTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentPageUUIDType) IsValid() bool {
	for _, existing := range allowedIncidentPageUUIDTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentPageUUIDType value.
func (v IncidentPageUUIDType) Ptr() *IncidentPageUUIDType {
	return &v
}

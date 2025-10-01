// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentPageIdType Incident page ID type.
type IncidentPageIdType string

// List of IncidentPageIdType.
const (
	INCIDENTPAGEIDTYPE_PAGE_UUID IncidentPageIdType = "page_uuid"
)

var allowedIncidentPageIdTypeEnumValues = []IncidentPageIdType{
	INCIDENTPAGEIDTYPE_PAGE_UUID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentPageIdType) GetAllowedValues() []IncidentPageIdType {
	return allowedIncidentPageIdTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentPageIdType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentPageIdType(value)
	return nil
}

// NewIncidentPageIdTypeFromValue returns a pointer to a valid IncidentPageIdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentPageIdTypeFromValue(v string) (*IncidentPageIdType, error) {
	ev := IncidentPageIdType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentPageIdType: valid values are %v", v, allowedIncidentPageIdTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentPageIdType) IsValid() bool {
	for _, existing := range allowedIncidentPageIdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentPageIdType value.
func (v IncidentPageIdType) Ptr() *IncidentPageIdType {
	return &v
}

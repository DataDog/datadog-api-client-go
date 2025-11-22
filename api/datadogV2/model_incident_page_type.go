// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentPageType Incident page type.
type IncidentPageType string

// List of IncidentPageType.
const (
	INCIDENTPAGETYPE_PAGE IncidentPageType = "page"
)

var allowedIncidentPageTypeEnumValues = []IncidentPageType{
	INCIDENTPAGETYPE_PAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentPageType) GetAllowedValues() []IncidentPageType {
	return allowedIncidentPageTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentPageType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentPageType(value)
	return nil
}

// NewIncidentPageTypeFromValue returns a pointer to a valid IncidentPageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentPageTypeFromValue(v string) (*IncidentPageType, error) {
	ev := IncidentPageType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentPageType: valid values are %v", v, allowedIncidentPageTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentPageType) IsValid() bool {
	for _, existing := range allowedIncidentPageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentPageType value.
func (v IncidentPageType) Ptr() *IncidentPageType {
	return &v
}

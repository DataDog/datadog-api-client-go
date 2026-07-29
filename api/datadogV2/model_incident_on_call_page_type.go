// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentOnCallPageType On-call page resource type.
type IncidentOnCallPageType string

// List of IncidentOnCallPageType.
const (
	INCIDENTONCALLPAGETYPE_PAGE IncidentOnCallPageType = "page"
)

var allowedIncidentOnCallPageTypeEnumValues = []IncidentOnCallPageType{
	INCIDENTONCALLPAGETYPE_PAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentOnCallPageType) GetAllowedValues() []IncidentOnCallPageType {
	return allowedIncidentOnCallPageTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentOnCallPageType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentOnCallPageType(value)
	return nil
}

// NewIncidentOnCallPageTypeFromValue returns a pointer to a valid IncidentOnCallPageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentOnCallPageTypeFromValue(v string) (*IncidentOnCallPageType, error) {
	ev := IncidentOnCallPageType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentOnCallPageType: valid values are %v", v, allowedIncidentOnCallPageTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentOnCallPageType) IsValid() bool {
	for _, existing := range allowedIncidentOnCallPageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentOnCallPageType value.
func (v IncidentOnCallPageType) Ptr() *IncidentOnCallPageType {
	return &v
}

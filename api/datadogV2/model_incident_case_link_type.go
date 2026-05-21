// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentCaseLinkType Case link resource type.
type IncidentCaseLinkType string

// List of IncidentCaseLinkType.
const (
	INCIDENTCASELINKTYPE_CASE_LINK IncidentCaseLinkType = "case_link"
)

var allowedIncidentCaseLinkTypeEnumValues = []IncidentCaseLinkType{
	INCIDENTCASELINKTYPE_CASE_LINK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentCaseLinkType) GetAllowedValues() []IncidentCaseLinkType {
	return allowedIncidentCaseLinkTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentCaseLinkType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentCaseLinkType(value)
	return nil
}

// NewIncidentCaseLinkTypeFromValue returns a pointer to a valid IncidentCaseLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentCaseLinkTypeFromValue(v string) (*IncidentCaseLinkType, error) {
	ev := IncidentCaseLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentCaseLinkType: valid values are %v", v, allowedIncidentCaseLinkTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentCaseLinkType) IsValid() bool {
	for _, existing := range allowedIncidentCaseLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentCaseLinkType value.
func (v IncidentCaseLinkType) Ptr() *IncidentCaseLinkType {
	return &v
}

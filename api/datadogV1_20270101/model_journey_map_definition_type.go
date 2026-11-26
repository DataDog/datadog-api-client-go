// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JourneyMapDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type JourneyMapDefinitionType string

// List of JourneyMapDefinitionType.
const (
	JOURNEYMAPDEFINITIONTYPE_JOURNEY_MAP JourneyMapDefinitionType = "journey_map"
)

var allowedJourneyMapDefinitionTypeEnumValues = []JourneyMapDefinitionType{
	JOURNEYMAPDEFINITIONTYPE_JOURNEY_MAP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *JourneyMapDefinitionType) GetAllowedValues() []JourneyMapDefinitionType {
	return allowedJourneyMapDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *JourneyMapDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = JourneyMapDefinitionType(value)
	return nil
}

// NewJourneyMapDefinitionTypeFromValue returns a pointer to a valid JourneyMapDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewJourneyMapDefinitionTypeFromValue(v string) (*JourneyMapDefinitionType, error) {
	ev := JourneyMapDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for JourneyMapDefinitionType: valid values are %v", v, allowedJourneyMapDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v JourneyMapDefinitionType) IsValid() bool {
	for _, existing := range allowedJourneyMapDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JourneyMapDefinitionType value.
func (v JourneyMapDefinitionType) Ptr() *JourneyMapDefinitionType {
	return &v
}

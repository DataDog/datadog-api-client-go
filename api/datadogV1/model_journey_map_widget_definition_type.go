// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JourneyMapWidgetDefinitionType Type of the Journey Paths widget.
type JourneyMapWidgetDefinitionType string

// List of JourneyMapWidgetDefinitionType.
const (
	JOURNEYMAPWIDGETDEFINITIONTYPE_JOURNEY_MAP JourneyMapWidgetDefinitionType = "journey_map"
)

var allowedJourneyMapWidgetDefinitionTypeEnumValues = []JourneyMapWidgetDefinitionType{
	JOURNEYMAPWIDGETDEFINITIONTYPE_JOURNEY_MAP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *JourneyMapWidgetDefinitionType) GetAllowedValues() []JourneyMapWidgetDefinitionType {
	return allowedJourneyMapWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *JourneyMapWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = JourneyMapWidgetDefinitionType(value)
	return nil
}

// NewJourneyMapWidgetDefinitionTypeFromValue returns a pointer to a valid JourneyMapWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewJourneyMapWidgetDefinitionTypeFromValue(v string) (*JourneyMapWidgetDefinitionType, error) {
	ev := JourneyMapWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for JourneyMapWidgetDefinitionType: valid values are %v", v, allowedJourneyMapWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v JourneyMapWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedJourneyMapWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JourneyMapWidgetDefinitionType value.
func (v JourneyMapWidgetDefinitionType) Ptr() *JourneyMapWidgetDefinitionType {
	return &v
}

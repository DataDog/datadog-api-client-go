// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SloWidgetViewMode Test-only generated schema for the dashboard versioning proof of concept.
type SloWidgetViewMode string

// List of SloWidgetViewMode.
const (
	SLOWIDGETVIEWMODE_OVERALL   SloWidgetViewMode = "overall"
	SLOWIDGETVIEWMODE_BOTH      SloWidgetViewMode = "both"
	SLOWIDGETVIEWMODE_COMPONENT SloWidgetViewMode = "component"
)

var allowedSloWidgetViewModeEnumValues = []SloWidgetViewMode{
	SLOWIDGETVIEWMODE_OVERALL,
	SLOWIDGETVIEWMODE_BOTH,
	SLOWIDGETVIEWMODE_COMPONENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SloWidgetViewMode) GetAllowedValues() []SloWidgetViewMode {
	return allowedSloWidgetViewModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SloWidgetViewMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SloWidgetViewMode(value)
	return nil
}

// NewSloWidgetViewModeFromValue returns a pointer to a valid SloWidgetViewMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSloWidgetViewModeFromValue(v string) (*SloWidgetViewMode, error) {
	ev := SloWidgetViewMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SloWidgetViewMode: valid values are %v", v, allowedSloWidgetViewModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SloWidgetViewMode) IsValid() bool {
	for _, existing := range allowedSloWidgetViewModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SloWidgetViewMode value.
func (v SloWidgetViewMode) Ptr() *SloWidgetViewMode {
	return &v
}

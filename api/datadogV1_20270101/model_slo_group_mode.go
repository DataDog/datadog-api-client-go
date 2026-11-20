// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SloGroupMode Test-only generated schema for the dashboard versioning proof of concept.
type SloGroupMode string

// List of SloGroupMode.
const (
	SLOGROUPMODE_OVERALL    SloGroupMode = "overall"
	SLOGROUPMODE_COMPONENTS SloGroupMode = "components"
)

var allowedSloGroupModeEnumValues = []SloGroupMode{
	SLOGROUPMODE_OVERALL,
	SLOGROUPMODE_COMPONENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SloGroupMode) GetAllowedValues() []SloGroupMode {
	return allowedSloGroupModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SloGroupMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SloGroupMode(value)
	return nil
}

// NewSloGroupModeFromValue returns a pointer to a valid SloGroupMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSloGroupModeFromValue(v string) (*SloGroupMode, error) {
	ev := SloGroupMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SloGroupMode: valid values are %v", v, allowedSloGroupModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SloGroupMode) IsValid() bool {
	for _, existing := range allowedSloGroupModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SloGroupMode value.
func (v SloGroupMode) Ptr() *SloGroupMode {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableTrendColumnFormatMode
type GuidedTableTrendColumnFormatMode string

// List of GuidedTableTrendColumnFormatMode.
const (
	GUIDEDTABLETRENDCOLUMNFORMATMODE_TREND GuidedTableTrendColumnFormatMode = "trend"
)

var allowedGuidedTableTrendColumnFormatModeEnumValues = []GuidedTableTrendColumnFormatMode{
	GUIDEDTABLETRENDCOLUMNFORMATMODE_TREND,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GuidedTableTrendColumnFormatMode) GetAllowedValues() []GuidedTableTrendColumnFormatMode {
	return allowedGuidedTableTrendColumnFormatModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GuidedTableTrendColumnFormatMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GuidedTableTrendColumnFormatMode(value)
	return nil
}

// NewGuidedTableTrendColumnFormatModeFromValue returns a pointer to a valid GuidedTableTrendColumnFormatMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGuidedTableTrendColumnFormatModeFromValue(v string) (*GuidedTableTrendColumnFormatMode, error) {
	ev := GuidedTableTrendColumnFormatMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GuidedTableTrendColumnFormatMode: valid values are %v", v, allowedGuidedTableTrendColumnFormatModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GuidedTableTrendColumnFormatMode) IsValid() bool {
	for _, existing := range allowedGuidedTableTrendColumnFormatModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GuidedTableTrendColumnFormatMode value.
func (v GuidedTableTrendColumnFormatMode) Ptr() *GuidedTableTrendColumnFormatMode {
	return &v
}

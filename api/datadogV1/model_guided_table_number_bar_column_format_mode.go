// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableNumberBarColumnFormatMode
type GuidedTableNumberBarColumnFormatMode string

// List of GuidedTableNumberBarColumnFormatMode.
const (
	GUIDEDTABLENUMBERBARCOLUMNFORMATMODE_NUMBER GuidedTableNumberBarColumnFormatMode = "number"
	GUIDEDTABLENUMBERBARCOLUMNFORMATMODE_BAR    GuidedTableNumberBarColumnFormatMode = "bar"
)

var allowedGuidedTableNumberBarColumnFormatModeEnumValues = []GuidedTableNumberBarColumnFormatMode{
	GUIDEDTABLENUMBERBARCOLUMNFORMATMODE_NUMBER,
	GUIDEDTABLENUMBERBARCOLUMNFORMATMODE_BAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GuidedTableNumberBarColumnFormatMode) GetAllowedValues() []GuidedTableNumberBarColumnFormatMode {
	return allowedGuidedTableNumberBarColumnFormatModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GuidedTableNumberBarColumnFormatMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GuidedTableNumberBarColumnFormatMode(value)
	return nil
}

// NewGuidedTableNumberBarColumnFormatModeFromValue returns a pointer to a valid GuidedTableNumberBarColumnFormatMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGuidedTableNumberBarColumnFormatModeFromValue(v string) (*GuidedTableNumberBarColumnFormatMode, error) {
	ev := GuidedTableNumberBarColumnFormatMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GuidedTableNumberBarColumnFormatMode: valid values are %v", v, allowedGuidedTableNumberBarColumnFormatModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GuidedTableNumberBarColumnFormatMode) IsValid() bool {
	for _, existing := range allowedGuidedTableNumberBarColumnFormatModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GuidedTableNumberBarColumnFormatMode value.
func (v GuidedTableNumberBarColumnFormatMode) Ptr() *GuidedTableNumberBarColumnFormatMode {
	return &v
}

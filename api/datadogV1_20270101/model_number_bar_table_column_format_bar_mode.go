// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NumberBarTableColumnFormatBarMode Test-only generated schema for the dashboard versioning proof of concept.
type NumberBarTableColumnFormatBarMode string

// List of NumberBarTableColumnFormatBarMode.
const (
	NUMBERBARTABLECOLUMNFORMATBARMODE_BAR NumberBarTableColumnFormatBarMode = "bar"
)

var allowedNumberBarTableColumnFormatBarModeEnumValues = []NumberBarTableColumnFormatBarMode{
	NUMBERBARTABLECOLUMNFORMATBARMODE_BAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NumberBarTableColumnFormatBarMode) GetAllowedValues() []NumberBarTableColumnFormatBarMode {
	return allowedNumberBarTableColumnFormatBarModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NumberBarTableColumnFormatBarMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NumberBarTableColumnFormatBarMode(value)
	return nil
}

// NewNumberBarTableColumnFormatBarModeFromValue returns a pointer to a valid NumberBarTableColumnFormatBarMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNumberBarTableColumnFormatBarModeFromValue(v string) (*NumberBarTableColumnFormatBarMode, error) {
	ev := NumberBarTableColumnFormatBarMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NumberBarTableColumnFormatBarMode: valid values are %v", v, allowedNumberBarTableColumnFormatBarModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NumberBarTableColumnFormatBarMode) IsValid() bool {
	for _, existing := range allowedNumberBarTableColumnFormatBarModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NumberBarTableColumnFormatBarMode value.
func (v NumberBarTableColumnFormatBarMode) Ptr() *NumberBarTableColumnFormatBarMode {
	return &v
}

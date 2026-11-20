// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NumberBarTableColumnFormatMode Test-only generated schema for the dashboard versioning proof of concept.
type NumberBarTableColumnFormatMode string

// List of NumberBarTableColumnFormatMode.
const (
	NUMBERBARTABLECOLUMNFORMATMODE_NUMBER NumberBarTableColumnFormatMode = "number"
	NUMBERBARTABLECOLUMNFORMATMODE_BAR    NumberBarTableColumnFormatMode = "bar"
)

var allowedNumberBarTableColumnFormatModeEnumValues = []NumberBarTableColumnFormatMode{
	NUMBERBARTABLECOLUMNFORMATMODE_NUMBER,
	NUMBERBARTABLECOLUMNFORMATMODE_BAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NumberBarTableColumnFormatMode) GetAllowedValues() []NumberBarTableColumnFormatMode {
	return allowedNumberBarTableColumnFormatModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NumberBarTableColumnFormatMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NumberBarTableColumnFormatMode(value)
	return nil
}

// NewNumberBarTableColumnFormatModeFromValue returns a pointer to a valid NumberBarTableColumnFormatMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNumberBarTableColumnFormatModeFromValue(v string) (*NumberBarTableColumnFormatMode, error) {
	ev := NumberBarTableColumnFormatMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NumberBarTableColumnFormatMode: valid values are %v", v, allowedNumberBarTableColumnFormatModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NumberBarTableColumnFormatMode) IsValid() bool {
	for _, existing := range allowedNumberBarTableColumnFormatModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NumberBarTableColumnFormatMode value.
func (v NumberBarTableColumnFormatMode) Ptr() *NumberBarTableColumnFormatMode {
	return &v
}

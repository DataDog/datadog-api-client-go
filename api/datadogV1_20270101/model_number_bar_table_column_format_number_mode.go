// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NumberBarTableColumnFormatNumberMode Test-only generated schema for the dashboard versioning proof of concept.
type NumberBarTableColumnFormatNumberMode string

// List of NumberBarTableColumnFormatNumberMode.
const (
	NUMBERBARTABLECOLUMNFORMATNUMBERMODE_NUMBER NumberBarTableColumnFormatNumberMode = "number"
)

var allowedNumberBarTableColumnFormatNumberModeEnumValues = []NumberBarTableColumnFormatNumberMode{
	NUMBERBARTABLECOLUMNFORMATNUMBERMODE_NUMBER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NumberBarTableColumnFormatNumberMode) GetAllowedValues() []NumberBarTableColumnFormatNumberMode {
	return allowedNumberBarTableColumnFormatNumberModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NumberBarTableColumnFormatNumberMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NumberBarTableColumnFormatNumberMode(value)
	return nil
}

// NewNumberBarTableColumnFormatNumberModeFromValue returns a pointer to a valid NumberBarTableColumnFormatNumberMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNumberBarTableColumnFormatNumberModeFromValue(v string) (*NumberBarTableColumnFormatNumberMode, error) {
	ev := NumberBarTableColumnFormatNumberMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NumberBarTableColumnFormatNumberMode: valid values are %v", v, allowedNumberBarTableColumnFormatNumberModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NumberBarTableColumnFormatNumberMode) IsValid() bool {
	for _, existing := range allowedNumberBarTableColumnFormatNumberModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NumberBarTableColumnFormatNumberMode value.
func (v NumberBarTableColumnFormatNumberMode) Ptr() *NumberBarTableColumnFormatNumberMode {
	return &v
}

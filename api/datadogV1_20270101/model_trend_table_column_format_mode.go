// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TrendTableColumnFormatMode Test-only generated schema for the dashboard versioning proof of concept.
type TrendTableColumnFormatMode string

// List of TrendTableColumnFormatMode.
const (
	TRENDTABLECOLUMNFORMATMODE_TREND TrendTableColumnFormatMode = "trend"
)

var allowedTrendTableColumnFormatModeEnumValues = []TrendTableColumnFormatMode{
	TRENDTABLECOLUMNFORMATMODE_TREND,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TrendTableColumnFormatMode) GetAllowedValues() []TrendTableColumnFormatMode {
	return allowedTrendTableColumnFormatModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TrendTableColumnFormatMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TrendTableColumnFormatMode(value)
	return nil
}

// NewTrendTableColumnFormatModeFromValue returns a pointer to a valid TrendTableColumnFormatMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTrendTableColumnFormatModeFromValue(v string) (*TrendTableColumnFormatMode, error) {
	ev := TrendTableColumnFormatMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TrendTableColumnFormatMode: valid values are %v", v, allowedTrendTableColumnFormatModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TrendTableColumnFormatMode) IsValid() bool {
	for _, existing := range allowedTrendTableColumnFormatModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TrendTableColumnFormatMode value.
func (v TrendTableColumnFormatMode) Ptr() *TrendTableColumnFormatMode {
	return &v
}

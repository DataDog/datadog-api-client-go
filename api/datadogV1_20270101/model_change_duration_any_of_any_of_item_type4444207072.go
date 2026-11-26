// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeDurationAnyOfAnyOfItemType4444207072 Test-only generated schema for the dashboard versioning proof of concept.
type ChangeDurationAnyOfAnyOfItemType4444207072 string

// List of ChangeDurationAnyOfAnyOfItemType4444207072.
const (
	CHANGEDURATIONANYOFANYOFITEMTYPE4444207072_PREVIOUS_TIMEFRAME ChangeDurationAnyOfAnyOfItemType4444207072 = "previous_timeframe"
	CHANGEDURATIONANYOFANYOFITEMTYPE4444207072_PREVIOUS_DAY       ChangeDurationAnyOfAnyOfItemType4444207072 = "previous_day"
	CHANGEDURATIONANYOFANYOFITEMTYPE4444207072_PREVIOUS_WEEK      ChangeDurationAnyOfAnyOfItemType4444207072 = "previous_week"
	CHANGEDURATIONANYOFANYOFITEMTYPE4444207072_PREVIOUS_MONTH     ChangeDurationAnyOfAnyOfItemType4444207072 = "previous_month"
)

var allowedChangeDurationAnyOfAnyOfItemType4444207072EnumValues = []ChangeDurationAnyOfAnyOfItemType4444207072{
	CHANGEDURATIONANYOFANYOFITEMTYPE4444207072_PREVIOUS_TIMEFRAME,
	CHANGEDURATIONANYOFANYOFITEMTYPE4444207072_PREVIOUS_DAY,
	CHANGEDURATIONANYOFANYOFITEMTYPE4444207072_PREVIOUS_WEEK,
	CHANGEDURATIONANYOFANYOFITEMTYPE4444207072_PREVIOUS_MONTH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ChangeDurationAnyOfAnyOfItemType4444207072) GetAllowedValues() []ChangeDurationAnyOfAnyOfItemType4444207072 {
	return allowedChangeDurationAnyOfAnyOfItemType4444207072EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ChangeDurationAnyOfAnyOfItemType4444207072) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ChangeDurationAnyOfAnyOfItemType4444207072(value)
	return nil
}

// NewChangeDurationAnyOfAnyOfItemType4444207072FromValue returns a pointer to a valid ChangeDurationAnyOfAnyOfItemType4444207072
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewChangeDurationAnyOfAnyOfItemType4444207072FromValue(v string) (*ChangeDurationAnyOfAnyOfItemType4444207072, error) {
	ev := ChangeDurationAnyOfAnyOfItemType4444207072(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ChangeDurationAnyOfAnyOfItemType4444207072: valid values are %v", v, allowedChangeDurationAnyOfAnyOfItemType4444207072EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ChangeDurationAnyOfAnyOfItemType4444207072) IsValid() bool {
	for _, existing := range allowedChangeDurationAnyOfAnyOfItemType4444207072EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChangeDurationAnyOfAnyOfItemType4444207072 value.
func (v ChangeDurationAnyOfAnyOfItemType4444207072) Ptr() *ChangeDurationAnyOfAnyOfItemType4444207072 {
	return &v
}

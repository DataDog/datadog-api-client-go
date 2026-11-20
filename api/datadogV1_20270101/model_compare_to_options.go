// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CompareToOptions Test-only generated schema for the dashboard versioning proof of concept.
type CompareToOptions string

// List of CompareToOptions.
const (
	COMPARETOOPTIONS_HOUR_BEFORE  CompareToOptions = "hour_before"
	COMPARETOOPTIONS_DAY_BEFORE   CompareToOptions = "day_before"
	COMPARETOOPTIONS_WEEK_BEFORE  CompareToOptions = "week_before"
	COMPARETOOPTIONS_MONTH_BEFORE CompareToOptions = "month_before"
)

var allowedCompareToOptionsEnumValues = []CompareToOptions{
	COMPARETOOPTIONS_HOUR_BEFORE,
	COMPARETOOPTIONS_DAY_BEFORE,
	COMPARETOOPTIONS_WEEK_BEFORE,
	COMPARETOOPTIONS_MONTH_BEFORE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CompareToOptions) GetAllowedValues() []CompareToOptions {
	return allowedCompareToOptionsEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CompareToOptions) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CompareToOptions(value)
	return nil
}

// NewCompareToOptionsFromValue returns a pointer to a valid CompareToOptions
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCompareToOptionsFromValue(v string) (*CompareToOptions, error) {
	ev := CompareToOptions(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CompareToOptions: valid values are %v", v, allowedCompareToOptionsEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CompareToOptions) IsValid() bool {
	for _, existing := range allowedCompareToOptionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CompareToOptions value.
func (v CompareToOptions) Ptr() *CompareToOptions {
	return &v
}

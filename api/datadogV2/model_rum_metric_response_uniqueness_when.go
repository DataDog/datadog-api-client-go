// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumMetricResponseUniquenessWhen When to count updatable events. "match" when the event is first seen, or "end" when the event is complete.
type RumMetricResponseUniquenessWhen string

// List of RumMetricResponseUniquenessWhen.
const (
	RUMMETRICRESPONSEUNIQUENESSWHEN_MATCH RumMetricResponseUniquenessWhen = "match"
	RUMMETRICRESPONSEUNIQUENESSWHEN_END   RumMetricResponseUniquenessWhen = "end"
)

var allowedRumMetricResponseUniquenessWhenEnumValues = []RumMetricResponseUniquenessWhen{
	RUMMETRICRESPONSEUNIQUENESSWHEN_MATCH,
	RUMMETRICRESPONSEUNIQUENESSWHEN_END,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumMetricResponseUniquenessWhen) GetAllowedValues() []RumMetricResponseUniquenessWhen {
	return allowedRumMetricResponseUniquenessWhenEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumMetricResponseUniquenessWhen) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumMetricResponseUniquenessWhen(value)
	return nil
}

// NewRumMetricResponseUniquenessWhenFromValue returns a pointer to a valid RumMetricResponseUniquenessWhen
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumMetricResponseUniquenessWhenFromValue(v string) (*RumMetricResponseUniquenessWhen, error) {
	ev := RumMetricResponseUniquenessWhen(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumMetricResponseUniquenessWhen: valid values are %v", v, allowedRumMetricResponseUniquenessWhenEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumMetricResponseUniquenessWhen) IsValid() bool {
	for _, existing := range allowedRumMetricResponseUniquenessWhenEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumMetricResponseUniquenessWhen value.
func (v RumMetricResponseUniquenessWhen) Ptr() *RumMetricResponseUniquenessWhen {
	return &v
}

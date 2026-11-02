// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JourneyMapCountingMethod Whether to count unique journeys or all journeys.
type JourneyMapCountingMethod string

// List of JourneyMapCountingMethod.
const (
	JOURNEYMAPCOUNTINGMETHOD_UNIQUES JourneyMapCountingMethod = "uniques"
	JOURNEYMAPCOUNTINGMETHOD_TOTALS  JourneyMapCountingMethod = "totals"
)

var allowedJourneyMapCountingMethodEnumValues = []JourneyMapCountingMethod{
	JOURNEYMAPCOUNTINGMETHOD_UNIQUES,
	JOURNEYMAPCOUNTINGMETHOD_TOTALS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *JourneyMapCountingMethod) GetAllowedValues() []JourneyMapCountingMethod {
	return allowedJourneyMapCountingMethodEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *JourneyMapCountingMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = JourneyMapCountingMethod(value)
	return nil
}

// NewJourneyMapCountingMethodFromValue returns a pointer to a valid JourneyMapCountingMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewJourneyMapCountingMethodFromValue(v string) (*JourneyMapCountingMethod, error) {
	ev := JourneyMapCountingMethod(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for JourneyMapCountingMethod: valid values are %v", v, allowedJourneyMapCountingMethodEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v JourneyMapCountingMethod) IsValid() bool {
	for _, existing := range allowedJourneyMapCountingMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JourneyMapCountingMethod value.
func (v JourneyMapCountingMethod) Ptr() *JourneyMapCountingMethod {
	return &v
}

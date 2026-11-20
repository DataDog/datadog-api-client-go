// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ValueLabelsDensity Test-only generated schema for the dashboard versioning proof of concept.
type ValueLabelsDensity string

// List of ValueLabelsDensity.
const (
	VALUELABELSDENSITY_MAXIMIZE ValueLabelsDensity = "maximize"
)

var allowedValueLabelsDensityEnumValues = []ValueLabelsDensity{
	VALUELABELSDENSITY_MAXIMIZE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ValueLabelsDensity) GetAllowedValues() []ValueLabelsDensity {
	return allowedValueLabelsDensityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ValueLabelsDensity) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ValueLabelsDensity(value)
	return nil
}

// NewValueLabelsDensityFromValue returns a pointer to a valid ValueLabelsDensity
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewValueLabelsDensityFromValue(v string) (*ValueLabelsDensity, error) {
	ev := ValueLabelsDensity(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ValueLabelsDensity: valid values are %v", v, allowedValueLabelsDensityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ValueLabelsDensity) IsValid() bool {
	for _, existing := range allowedValueLabelsDensityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ValueLabelsDensity value.
func (v ValueLabelsDensity) Ptr() *ValueLabelsDensity {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeDirectionality Test-only generated schema for the dashboard versioning proof of concept.
type ChangeDirectionality string

// List of ChangeDirectionality.
const (
	CHANGEDIRECTIONALITY_INCREASE_BETTER ChangeDirectionality = "increase_better"
	CHANGEDIRECTIONALITY_DECREASE_BETTER ChangeDirectionality = "decrease_better"
	CHANGEDIRECTIONALITY_NEUTRAL         ChangeDirectionality = "neutral"
)

var allowedChangeDirectionalityEnumValues = []ChangeDirectionality{
	CHANGEDIRECTIONALITY_INCREASE_BETTER,
	CHANGEDIRECTIONALITY_DECREASE_BETTER,
	CHANGEDIRECTIONALITY_NEUTRAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ChangeDirectionality) GetAllowedValues() []ChangeDirectionality {
	return allowedChangeDirectionalityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ChangeDirectionality) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ChangeDirectionality(value)
	return nil
}

// NewChangeDirectionalityFromValue returns a pointer to a valid ChangeDirectionality
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewChangeDirectionalityFromValue(v string) (*ChangeDirectionality, error) {
	ev := ChangeDirectionality(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ChangeDirectionality: valid values are %v", v, allowedChangeDirectionalityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ChangeDirectionality) IsValid() bool {
	for _, existing := range allowedChangeDirectionalityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChangeDirectionality value.
func (v ChangeDirectionality) Ptr() *ChangeDirectionality {
	return &v
}

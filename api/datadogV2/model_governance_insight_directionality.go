// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceInsightDirectionality Whether an increase in the insight's value is good, bad, or neutral.
type GovernanceInsightDirectionality string

// List of GovernanceInsightDirectionality.
const (
	GOVERNANCEINSIGHTDIRECTIONALITY_NEUTRAL         GovernanceInsightDirectionality = "neutral"
	GOVERNANCEINSIGHTDIRECTIONALITY_INCREASE_BETTER GovernanceInsightDirectionality = "increase_better"
	GOVERNANCEINSIGHTDIRECTIONALITY_DECREASE_BETTER GovernanceInsightDirectionality = "decrease_better"
)

var allowedGovernanceInsightDirectionalityEnumValues = []GovernanceInsightDirectionality{
	GOVERNANCEINSIGHTDIRECTIONALITY_NEUTRAL,
	GOVERNANCEINSIGHTDIRECTIONALITY_INCREASE_BETTER,
	GOVERNANCEINSIGHTDIRECTIONALITY_DECREASE_BETTER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GovernanceInsightDirectionality) GetAllowedValues() []GovernanceInsightDirectionality {
	return allowedGovernanceInsightDirectionalityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GovernanceInsightDirectionality) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GovernanceInsightDirectionality(value)
	return nil
}

// NewGovernanceInsightDirectionalityFromValue returns a pointer to a valid GovernanceInsightDirectionality
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGovernanceInsightDirectionalityFromValue(v string) (*GovernanceInsightDirectionality, error) {
	ev := GovernanceInsightDirectionality(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GovernanceInsightDirectionality: valid values are %v", v, allowedGovernanceInsightDirectionalityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GovernanceInsightDirectionality) IsValid() bool {
	for _, existing := range allowedGovernanceInsightDirectionalityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GovernanceInsightDirectionality value.
func (v GovernanceInsightDirectionality) Ptr() *GovernanceInsightDirectionality {
	return &v
}

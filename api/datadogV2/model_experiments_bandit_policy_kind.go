// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditPolicyKind WARMUP serves the arms equally. ADAPTIVE serves the initial split or a split from a refresh. FALLBACK serves one variant after the bandit ends.
type ExperimentsBanditPolicyKind string

// List of ExperimentsBanditPolicyKind.
const (
	EXPERIMENTSBANDITPOLICYKIND_WARMUP   ExperimentsBanditPolicyKind = "WARMUP"
	EXPERIMENTSBANDITPOLICYKIND_ADAPTIVE ExperimentsBanditPolicyKind = "ADAPTIVE"
	EXPERIMENTSBANDITPOLICYKIND_FALLBACK ExperimentsBanditPolicyKind = "FALLBACK"
)

var allowedExperimentsBanditPolicyKindEnumValues = []ExperimentsBanditPolicyKind{
	EXPERIMENTSBANDITPOLICYKIND_WARMUP,
	EXPERIMENTSBANDITPOLICYKIND_ADAPTIVE,
	EXPERIMENTSBANDITPOLICYKIND_FALLBACK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsBanditPolicyKind) GetAllowedValues() []ExperimentsBanditPolicyKind {
	return allowedExperimentsBanditPolicyKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsBanditPolicyKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsBanditPolicyKind(value)
	return nil
}

// NewExperimentsBanditPolicyKindFromValue returns a pointer to a valid ExperimentsBanditPolicyKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsBanditPolicyKindFromValue(v string) (*ExperimentsBanditPolicyKind, error) {
	ev := ExperimentsBanditPolicyKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsBanditPolicyKind: valid values are %v", v, allowedExperimentsBanditPolicyKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsBanditPolicyKind) IsValid() bool {
	for _, existing := range allowedExperimentsBanditPolicyKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsBanditPolicyKind value.
func (v ExperimentsBanditPolicyKind) Ptr() *ExperimentsBanditPolicyKind {
	return &v
}

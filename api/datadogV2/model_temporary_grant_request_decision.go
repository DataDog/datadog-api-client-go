// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TemporaryGrantRequestDecision Approve the exact submitted total and expiry, or deny the request.
type TemporaryGrantRequestDecision string

// List of TemporaryGrantRequestDecision.
const (
	TEMPORARYGRANTREQUESTDECISION_APPROVED TemporaryGrantRequestDecision = "approved"
	TEMPORARYGRANTREQUESTDECISION_DENIED   TemporaryGrantRequestDecision = "denied"
)

var allowedTemporaryGrantRequestDecisionEnumValues = []TemporaryGrantRequestDecision{
	TEMPORARYGRANTREQUESTDECISION_APPROVED,
	TEMPORARYGRANTREQUESTDECISION_DENIED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TemporaryGrantRequestDecision) GetAllowedValues() []TemporaryGrantRequestDecision {
	return allowedTemporaryGrantRequestDecisionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TemporaryGrantRequestDecision) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TemporaryGrantRequestDecision(value)
	return nil
}

// NewTemporaryGrantRequestDecisionFromValue returns a pointer to a valid TemporaryGrantRequestDecision
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTemporaryGrantRequestDecisionFromValue(v string) (*TemporaryGrantRequestDecision, error) {
	ev := TemporaryGrantRequestDecision(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TemporaryGrantRequestDecision: valid values are %v", v, allowedTemporaryGrantRequestDecisionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TemporaryGrantRequestDecision) IsValid() bool {
	for _, existing := range allowedTemporaryGrantRequestDecisionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TemporaryGrantRequestDecision value.
func (v TemporaryGrantRequestDecision) Ptr() *TemporaryGrantRequestDecision {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TemporaryGrantRequestGrantState Current state of the grant associated with an approved request. Replaced means a successor superseded this grant; the successor need not still be active. Absent for requests that were not approved.
type TemporaryGrantRequestGrantState string

// List of TemporaryGrantRequestGrantState.
const (
	TEMPORARYGRANTREQUESTGRANTSTATE_ACTIVE   TemporaryGrantRequestGrantState = "active"
	TEMPORARYGRANTREQUESTGRANTSTATE_EXPIRED  TemporaryGrantRequestGrantState = "expired"
	TEMPORARYGRANTREQUESTGRANTSTATE_REVOKED  TemporaryGrantRequestGrantState = "revoked"
	TEMPORARYGRANTREQUESTGRANTSTATE_REPLACED TemporaryGrantRequestGrantState = "replaced"
)

var allowedTemporaryGrantRequestGrantStateEnumValues = []TemporaryGrantRequestGrantState{
	TEMPORARYGRANTREQUESTGRANTSTATE_ACTIVE,
	TEMPORARYGRANTREQUESTGRANTSTATE_EXPIRED,
	TEMPORARYGRANTREQUESTGRANTSTATE_REVOKED,
	TEMPORARYGRANTREQUESTGRANTSTATE_REPLACED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TemporaryGrantRequestGrantState) GetAllowedValues() []TemporaryGrantRequestGrantState {
	return allowedTemporaryGrantRequestGrantStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TemporaryGrantRequestGrantState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TemporaryGrantRequestGrantState(value)
	return nil
}

// NewTemporaryGrantRequestGrantStateFromValue returns a pointer to a valid TemporaryGrantRequestGrantState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTemporaryGrantRequestGrantStateFromValue(v string) (*TemporaryGrantRequestGrantState, error) {
	ev := TemporaryGrantRequestGrantState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TemporaryGrantRequestGrantState: valid values are %v", v, allowedTemporaryGrantRequestGrantStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TemporaryGrantRequestGrantState) IsValid() bool {
	for _, existing := range allowedTemporaryGrantRequestGrantStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TemporaryGrantRequestGrantState value.
func (v TemporaryGrantRequestGrantState) Ptr() *TemporaryGrantRequestGrantState {
	return &v
}

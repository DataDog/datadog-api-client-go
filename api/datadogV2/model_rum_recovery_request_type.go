// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRecoveryRequestType JSON:API resource type.
type RumRecoveryRequestType string

// List of RumRecoveryRequestType.
const (
	RUMRECOVERYREQUESTTYPE_RUM_RECOVERY_REQUEST RumRecoveryRequestType = "rum_recovery_request"
)

var allowedRumRecoveryRequestTypeEnumValues = []RumRecoveryRequestType{
	RUMRECOVERYREQUESTTYPE_RUM_RECOVERY_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumRecoveryRequestType) GetAllowedValues() []RumRecoveryRequestType {
	return allowedRumRecoveryRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRecoveryRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRecoveryRequestType(value)
	return nil
}

// NewRumRecoveryRequestTypeFromValue returns a pointer to a valid RumRecoveryRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRecoveryRequestTypeFromValue(v string) (*RumRecoveryRequestType, error) {
	ev := RumRecoveryRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRecoveryRequestType: valid values are %v", v, allowedRumRecoveryRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRecoveryRequestType) IsValid() bool {
	for _, existing := range allowedRumRecoveryRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRecoveryRequestType value.
func (v RumRecoveryRequestType) Ptr() *RumRecoveryRequestType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRecoveryRequestType JSON:API resource type.
type RUMRecoveryRequestType string

// List of RUMRecoveryRequestType.
const (
	RUMRECOVERYREQUESTTYPE_RUM_RECOVERY_REQUEST RUMRecoveryRequestType = "rum_recovery_request"
)

var allowedRUMRecoveryRequestTypeEnumValues = []RUMRecoveryRequestType{
	RUMRECOVERYREQUESTTYPE_RUM_RECOVERY_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMRecoveryRequestType) GetAllowedValues() []RUMRecoveryRequestType {
	return allowedRUMRecoveryRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMRecoveryRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMRecoveryRequestType(value)
	return nil
}

// NewRUMRecoveryRequestTypeFromValue returns a pointer to a valid RUMRecoveryRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMRecoveryRequestTypeFromValue(v string) (*RUMRecoveryRequestType, error) {
	ev := RUMRecoveryRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMRecoveryRequestType: valid values are %v", v, allowedRUMRecoveryRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMRecoveryRequestType) IsValid() bool {
	for _, existing := range allowedRUMRecoveryRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMRecoveryRequestType value.
func (v RUMRecoveryRequestType) Ptr() *RUMRecoveryRequestType {
	return &v
}

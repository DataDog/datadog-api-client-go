// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRecoveryCountResponseDataType JSON:API resource type.
type RumRecoveryCountResponseDataType string

// List of RumRecoveryCountResponseDataType.
const (
	RUMRECOVERYCOUNTRESPONSEDATATYPE_RUM_RECOVERY_COUNT RumRecoveryCountResponseDataType = "rum_recovery_count"
)

var allowedRumRecoveryCountResponseDataTypeEnumValues = []RumRecoveryCountResponseDataType{
	RUMRECOVERYCOUNTRESPONSEDATATYPE_RUM_RECOVERY_COUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumRecoveryCountResponseDataType) GetAllowedValues() []RumRecoveryCountResponseDataType {
	return allowedRumRecoveryCountResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRecoveryCountResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRecoveryCountResponseDataType(value)
	return nil
}

// NewRumRecoveryCountResponseDataTypeFromValue returns a pointer to a valid RumRecoveryCountResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRecoveryCountResponseDataTypeFromValue(v string) (*RumRecoveryCountResponseDataType, error) {
	ev := RumRecoveryCountResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRecoveryCountResponseDataType: valid values are %v", v, allowedRumRecoveryCountResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRecoveryCountResponseDataType) IsValid() bool {
	for _, existing := range allowedRumRecoveryCountResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRecoveryCountResponseDataType value.
func (v RumRecoveryCountResponseDataType) Ptr() *RumRecoveryCountResponseDataType {
	return &v
}

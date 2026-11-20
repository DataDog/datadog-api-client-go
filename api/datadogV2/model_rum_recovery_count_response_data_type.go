// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRecoveryCountResponseDataType JSON:API resource type.
type RUMRecoveryCountResponseDataType string

// List of RUMRecoveryCountResponseDataType.
const (
	RUMRECOVERYCOUNTRESPONSEDATATYPE_RUM_RECOVERY_COUNT RUMRecoveryCountResponseDataType = "rum_recovery_count"
)

var allowedRUMRecoveryCountResponseDataTypeEnumValues = []RUMRecoveryCountResponseDataType{
	RUMRECOVERYCOUNTRESPONSEDATATYPE_RUM_RECOVERY_COUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMRecoveryCountResponseDataType) GetAllowedValues() []RUMRecoveryCountResponseDataType {
	return allowedRUMRecoveryCountResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMRecoveryCountResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMRecoveryCountResponseDataType(value)
	return nil
}

// NewRUMRecoveryCountResponseDataTypeFromValue returns a pointer to a valid RUMRecoveryCountResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMRecoveryCountResponseDataTypeFromValue(v string) (*RUMRecoveryCountResponseDataType, error) {
	ev := RUMRecoveryCountResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMRecoveryCountResponseDataType: valid values are %v", v, allowedRUMRecoveryCountResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMRecoveryCountResponseDataType) IsValid() bool {
	for _, existing := range allowedRUMRecoveryCountResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMRecoveryCountResponseDataType value.
func (v RUMRecoveryCountResponseDataType) Ptr() *RUMRecoveryCountResponseDataType {
	return &v
}

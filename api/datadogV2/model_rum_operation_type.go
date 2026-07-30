// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMOperationType The JSON:API type for RUM operation resources.
type RUMOperationType string

// List of RUMOperationType.
const (
	RUMOPERATIONTYPE_OPERATIONS RUMOperationType = "operations"
)

var allowedRUMOperationTypeEnumValues = []RUMOperationType{
	RUMOPERATIONTYPE_OPERATIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMOperationType) GetAllowedValues() []RUMOperationType {
	return allowedRUMOperationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMOperationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMOperationType(value)
	return nil
}

// NewRUMOperationTypeFromValue returns a pointer to a valid RUMOperationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMOperationTypeFromValue(v string) (*RUMOperationType, error) {
	ev := RUMOperationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMOperationType: valid values are %v", v, allowedRUMOperationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMOperationType) IsValid() bool {
	for _, existing := range allowedRUMOperationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMOperationType value.
func (v RUMOperationType) Ptr() *RUMOperationType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntakeAPIKeyType The resource type for an intake API key.
type IntakeAPIKeyType string

// List of IntakeAPIKeyType.
const (
	INTAKEAPIKEYTYPE_INTAKE_API_KEY IntakeAPIKeyType = "intake_api_key"
)

var allowedIntakeAPIKeyTypeEnumValues = []IntakeAPIKeyType{
	INTAKEAPIKEYTYPE_INTAKE_API_KEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IntakeAPIKeyType) GetAllowedValues() []IntakeAPIKeyType {
	return allowedIntakeAPIKeyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IntakeAPIKeyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IntakeAPIKeyType(value)
	return nil
}

// NewIntakeAPIKeyTypeFromValue returns a pointer to a valid IntakeAPIKeyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIntakeAPIKeyTypeFromValue(v string) (*IntakeAPIKeyType, error) {
	ev := IntakeAPIKeyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IntakeAPIKeyType: valid values are %v", v, allowedIntakeAPIKeyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IntakeAPIKeyType) IsValid() bool {
	for _, existing := range allowedIntakeAPIKeyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IntakeAPIKeyType value.
func (v IntakeAPIKeyType) Ptr() *IntakeAPIKeyType {
	return &v
}

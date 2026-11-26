// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type ChangeBackendDefinitionType string

// List of ChangeBackendDefinitionType.
const (
	CHANGEBACKENDDEFINITIONTYPE_CHANGE ChangeBackendDefinitionType = "change"
)

var allowedChangeBackendDefinitionTypeEnumValues = []ChangeBackendDefinitionType{
	CHANGEBACKENDDEFINITIONTYPE_CHANGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ChangeBackendDefinitionType) GetAllowedValues() []ChangeBackendDefinitionType {
	return allowedChangeBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ChangeBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ChangeBackendDefinitionType(value)
	return nil
}

// NewChangeBackendDefinitionTypeFromValue returns a pointer to a valid ChangeBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewChangeBackendDefinitionTypeFromValue(v string) (*ChangeBackendDefinitionType, error) {
	ev := ChangeBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ChangeBackendDefinitionType: valid values are %v", v, allowedChangeBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ChangeBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedChangeBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChangeBackendDefinitionType value.
func (v ChangeBackendDefinitionType) Ptr() *ChangeBackendDefinitionType {
	return &v
}

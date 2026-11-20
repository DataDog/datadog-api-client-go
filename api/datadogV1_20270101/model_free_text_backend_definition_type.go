// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FreeTextBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type FreeTextBackendDefinitionType string

// List of FreeTextBackendDefinitionType.
const (
	FREETEXTBACKENDDEFINITIONTYPE_FREE_TEXT FreeTextBackendDefinitionType = "free_text"
)

var allowedFreeTextBackendDefinitionTypeEnumValues = []FreeTextBackendDefinitionType{
	FREETEXTBACKENDDEFINITIONTYPE_FREE_TEXT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FreeTextBackendDefinitionType) GetAllowedValues() []FreeTextBackendDefinitionType {
	return allowedFreeTextBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FreeTextBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FreeTextBackendDefinitionType(value)
	return nil
}

// NewFreeTextBackendDefinitionTypeFromValue returns a pointer to a valid FreeTextBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFreeTextBackendDefinitionTypeFromValue(v string) (*FreeTextBackendDefinitionType, error) {
	ev := FreeTextBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FreeTextBackendDefinitionType: valid values are %v", v, allowedFreeTextBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FreeTextBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedFreeTextBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FreeTextBackendDefinitionType value.
func (v FreeTextBackendDefinitionType) Ptr() *FreeTextBackendDefinitionType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityContextEntityType The type of entity to retrieve. Only `siem_entity_identity` is currently supported.
type EntityContextEntityType string

// List of EntityContextEntityType.
const (
	ENTITYCONTEXTENTITYTYPE_SIEM_ENTITY_IDENTITY EntityContextEntityType = "siem_entity_identity"
)

var allowedEntityContextEntityTypeEnumValues = []EntityContextEntityType{
	ENTITYCONTEXTENTITYTYPE_SIEM_ENTITY_IDENTITY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EntityContextEntityType) GetAllowedValues() []EntityContextEntityType {
	return allowedEntityContextEntityTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EntityContextEntityType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EntityContextEntityType(value)
	return nil
}

// NewEntityContextEntityTypeFromValue returns a pointer to a valid EntityContextEntityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEntityContextEntityTypeFromValue(v string) (*EntityContextEntityType, error) {
	ev := EntityContextEntityType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EntityContextEntityType: valid values are %v", v, allowedEntityContextEntityTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EntityContextEntityType) IsValid() bool {
	for _, existing := range allowedEntityContextEntityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityContextEntityType value.
func (v EntityContextEntityType) Ptr() *EntityContextEntityType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntitySchemaVersion Entity schema version for conversion.
type EntitySchemaVersion string

// List of EntitySchemaVersion.
const (
	ENTITYSCHEMAVERSION_V2   EntitySchemaVersion = "v2"
	ENTITYSCHEMAVERSION_V2_1 EntitySchemaVersion = "v2.1"
	ENTITYSCHEMAVERSION_V2_2 EntitySchemaVersion = "v2.2"
	ENTITYSCHEMAVERSION_V3   EntitySchemaVersion = "v3"
)

var allowedEntitySchemaVersionEnumValues = []EntitySchemaVersion{
	ENTITYSCHEMAVERSION_V2,
	ENTITYSCHEMAVERSION_V2_1,
	ENTITYSCHEMAVERSION_V2_2,
	ENTITYSCHEMAVERSION_V3,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EntitySchemaVersion) GetAllowedValues() []EntitySchemaVersion {
	return allowedEntitySchemaVersionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EntitySchemaVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EntitySchemaVersion(value)
	return nil
}

// NewEntitySchemaVersionFromValue returns a pointer to a valid EntitySchemaVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEntitySchemaVersionFromValue(v string) (*EntitySchemaVersion, error) {
	ev := EntitySchemaVersion(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EntitySchemaVersion: valid values are %v", v, allowedEntitySchemaVersionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EntitySchemaVersion) IsValid() bool {
	for _, existing := range allowedEntitySchemaVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntitySchemaVersion value.
func (v EntitySchemaVersion) Ptr() *EntitySchemaVersion {
	return &v
}

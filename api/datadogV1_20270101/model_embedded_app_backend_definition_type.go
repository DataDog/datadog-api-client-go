// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EmbeddedAppBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type EmbeddedAppBackendDefinitionType string

// List of EmbeddedAppBackendDefinitionType.
const (
	EMBEDDEDAPPBACKENDDEFINITIONTYPE_EMBEDDED_APP EmbeddedAppBackendDefinitionType = "embedded_app"
)

var allowedEmbeddedAppBackendDefinitionTypeEnumValues = []EmbeddedAppBackendDefinitionType{
	EMBEDDEDAPPBACKENDDEFINITIONTYPE_EMBEDDED_APP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EmbeddedAppBackendDefinitionType) GetAllowedValues() []EmbeddedAppBackendDefinitionType {
	return allowedEmbeddedAppBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EmbeddedAppBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EmbeddedAppBackendDefinitionType(value)
	return nil
}

// NewEmbeddedAppBackendDefinitionTypeFromValue returns a pointer to a valid EmbeddedAppBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEmbeddedAppBackendDefinitionTypeFromValue(v string) (*EmbeddedAppBackendDefinitionType, error) {
	ev := EmbeddedAppBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EmbeddedAppBackendDefinitionType: valid values are %v", v, allowedEmbeddedAppBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EmbeddedAppBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedEmbeddedAppBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EmbeddedAppBackendDefinitionType value.
func (v EmbeddedAppBackendDefinitionType) Ptr() *EmbeddedAppBackendDefinitionType {
	return &v
}

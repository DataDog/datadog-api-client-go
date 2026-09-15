// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EmbeddedAppWidgetDefinitionType Type of the embedded app widget.
type EmbeddedAppWidgetDefinitionType string

// List of EmbeddedAppWidgetDefinitionType.
const (
	EMBEDDEDAPPWIDGETDEFINITIONTYPE_EMBEDDED_APP EmbeddedAppWidgetDefinitionType = "embedded_app"
)

var allowedEmbeddedAppWidgetDefinitionTypeEnumValues = []EmbeddedAppWidgetDefinitionType{
	EMBEDDEDAPPWIDGETDEFINITIONTYPE_EMBEDDED_APP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EmbeddedAppWidgetDefinitionType) GetAllowedValues() []EmbeddedAppWidgetDefinitionType {
	return allowedEmbeddedAppWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EmbeddedAppWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EmbeddedAppWidgetDefinitionType(value)
	return nil
}

// NewEmbeddedAppWidgetDefinitionTypeFromValue returns a pointer to a valid EmbeddedAppWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEmbeddedAppWidgetDefinitionTypeFromValue(v string) (*EmbeddedAppWidgetDefinitionType, error) {
	ev := EmbeddedAppWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EmbeddedAppWidgetDefinitionType: valid values are %v", v, allowedEmbeddedAppWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EmbeddedAppWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedEmbeddedAppWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EmbeddedAppWidgetDefinitionType value.
func (v EmbeddedAppWidgetDefinitionType) Ptr() *EmbeddedAppWidgetDefinitionType {
	return &v
}

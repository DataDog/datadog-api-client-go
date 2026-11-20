// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FreeTextBackendDefinitionTextAlign Test-only generated schema for the dashboard versioning proof of concept.
type FreeTextBackendDefinitionTextAlign string

// List of FreeTextBackendDefinitionTextAlign.
const (
	FREETEXTBACKENDDEFINITIONTEXTALIGN_LEFT   FreeTextBackendDefinitionTextAlign = "left"
	FREETEXTBACKENDDEFINITIONTEXTALIGN_CENTER FreeTextBackendDefinitionTextAlign = "center"
	FREETEXTBACKENDDEFINITIONTEXTALIGN_RIGHT  FreeTextBackendDefinitionTextAlign = "right"
)

var allowedFreeTextBackendDefinitionTextAlignEnumValues = []FreeTextBackendDefinitionTextAlign{
	FREETEXTBACKENDDEFINITIONTEXTALIGN_LEFT,
	FREETEXTBACKENDDEFINITIONTEXTALIGN_CENTER,
	FREETEXTBACKENDDEFINITIONTEXTALIGN_RIGHT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FreeTextBackendDefinitionTextAlign) GetAllowedValues() []FreeTextBackendDefinitionTextAlign {
	return allowedFreeTextBackendDefinitionTextAlignEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FreeTextBackendDefinitionTextAlign) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FreeTextBackendDefinitionTextAlign(value)
	return nil
}

// NewFreeTextBackendDefinitionTextAlignFromValue returns a pointer to a valid FreeTextBackendDefinitionTextAlign
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFreeTextBackendDefinitionTextAlignFromValue(v string) (*FreeTextBackendDefinitionTextAlign, error) {
	ev := FreeTextBackendDefinitionTextAlign(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FreeTextBackendDefinitionTextAlign: valid values are %v", v, allowedFreeTextBackendDefinitionTextAlignEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FreeTextBackendDefinitionTextAlign) IsValid() bool {
	for _, existing := range allowedFreeTextBackendDefinitionTextAlignEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FreeTextBackendDefinitionTextAlign value.
func (v FreeTextBackendDefinitionTextAlign) Ptr() *FreeTextBackendDefinitionTextAlign {
	return &v
}

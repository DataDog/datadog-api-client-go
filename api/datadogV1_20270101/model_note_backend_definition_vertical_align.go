// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NoteBackendDefinitionVerticalAlign Test-only generated schema for the dashboard versioning proof of concept.
type NoteBackendDefinitionVerticalAlign string

// List of NoteBackendDefinitionVerticalAlign.
const (
	NOTEBACKENDDEFINITIONVERTICALALIGN_CENTER NoteBackendDefinitionVerticalAlign = "center"
	NOTEBACKENDDEFINITIONVERTICALALIGN_TOP    NoteBackendDefinitionVerticalAlign = "top"
	NOTEBACKENDDEFINITIONVERTICALALIGN_BOTTOM NoteBackendDefinitionVerticalAlign = "bottom"
)

var allowedNoteBackendDefinitionVerticalAlignEnumValues = []NoteBackendDefinitionVerticalAlign{
	NOTEBACKENDDEFINITIONVERTICALALIGN_CENTER,
	NOTEBACKENDDEFINITIONVERTICALALIGN_TOP,
	NOTEBACKENDDEFINITIONVERTICALALIGN_BOTTOM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NoteBackendDefinitionVerticalAlign) GetAllowedValues() []NoteBackendDefinitionVerticalAlign {
	return allowedNoteBackendDefinitionVerticalAlignEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NoteBackendDefinitionVerticalAlign) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NoteBackendDefinitionVerticalAlign(value)
	return nil
}

// NewNoteBackendDefinitionVerticalAlignFromValue returns a pointer to a valid NoteBackendDefinitionVerticalAlign
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNoteBackendDefinitionVerticalAlignFromValue(v string) (*NoteBackendDefinitionVerticalAlign, error) {
	ev := NoteBackendDefinitionVerticalAlign(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NoteBackendDefinitionVerticalAlign: valid values are %v", v, allowedNoteBackendDefinitionVerticalAlignEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NoteBackendDefinitionVerticalAlign) IsValid() bool {
	for _, existing := range allowedNoteBackendDefinitionVerticalAlignEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NoteBackendDefinitionVerticalAlign value.
func (v NoteBackendDefinitionVerticalAlign) Ptr() *NoteBackendDefinitionVerticalAlign {
	return &v
}

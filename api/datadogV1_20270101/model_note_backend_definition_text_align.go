// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NoteBackendDefinitionTextAlign Test-only generated schema for the dashboard versioning proof of concept.
type NoteBackendDefinitionTextAlign string

// List of NoteBackendDefinitionTextAlign.
const (
	NOTEBACKENDDEFINITIONTEXTALIGN_LEFT   NoteBackendDefinitionTextAlign = "left"
	NOTEBACKENDDEFINITIONTEXTALIGN_CENTER NoteBackendDefinitionTextAlign = "center"
	NOTEBACKENDDEFINITIONTEXTALIGN_RIGHT  NoteBackendDefinitionTextAlign = "right"
)

var allowedNoteBackendDefinitionTextAlignEnumValues = []NoteBackendDefinitionTextAlign{
	NOTEBACKENDDEFINITIONTEXTALIGN_LEFT,
	NOTEBACKENDDEFINITIONTEXTALIGN_CENTER,
	NOTEBACKENDDEFINITIONTEXTALIGN_RIGHT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NoteBackendDefinitionTextAlign) GetAllowedValues() []NoteBackendDefinitionTextAlign {
	return allowedNoteBackendDefinitionTextAlignEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NoteBackendDefinitionTextAlign) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NoteBackendDefinitionTextAlign(value)
	return nil
}

// NewNoteBackendDefinitionTextAlignFromValue returns a pointer to a valid NoteBackendDefinitionTextAlign
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNoteBackendDefinitionTextAlignFromValue(v string) (*NoteBackendDefinitionTextAlign, error) {
	ev := NoteBackendDefinitionTextAlign(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NoteBackendDefinitionTextAlign: valid values are %v", v, allowedNoteBackendDefinitionTextAlignEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NoteBackendDefinitionTextAlign) IsValid() bool {
	for _, existing := range allowedNoteBackendDefinitionTextAlignEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NoteBackendDefinitionTextAlign value.
func (v NoteBackendDefinitionTextAlign) Ptr() *NoteBackendDefinitionTextAlign {
	return &v
}

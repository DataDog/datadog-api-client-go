// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NoteBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type NoteBackendDefinitionType string

// List of NoteBackendDefinitionType.
const (
	NOTEBACKENDDEFINITIONTYPE_NOTE NoteBackendDefinitionType = "note"
)

var allowedNoteBackendDefinitionTypeEnumValues = []NoteBackendDefinitionType{
	NOTEBACKENDDEFINITIONTYPE_NOTE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NoteBackendDefinitionType) GetAllowedValues() []NoteBackendDefinitionType {
	return allowedNoteBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NoteBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NoteBackendDefinitionType(value)
	return nil
}

// NewNoteBackendDefinitionTypeFromValue returns a pointer to a valid NoteBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNoteBackendDefinitionTypeFromValue(v string) (*NoteBackendDefinitionType, error) {
	ev := NoteBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NoteBackendDefinitionType: valid values are %v", v, allowedNoteBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NoteBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedNoteBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NoteBackendDefinitionType value.
func (v NoteBackendDefinitionType) Ptr() *NoteBackendDefinitionType {
	return &v
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NoteBackendDefinitionTickEdge Test-only generated schema for the dashboard versioning proof of concept.
type NoteBackendDefinitionTickEdge string

// List of NoteBackendDefinitionTickEdge.
const (
	NOTEBACKENDDEFINITIONTICKEDGE_LEFT   NoteBackendDefinitionTickEdge = "left"
	NOTEBACKENDDEFINITIONTICKEDGE_RIGHT  NoteBackendDefinitionTickEdge = "right"
	NOTEBACKENDDEFINITIONTICKEDGE_TOP    NoteBackendDefinitionTickEdge = "top"
	NOTEBACKENDDEFINITIONTICKEDGE_BOTTOM NoteBackendDefinitionTickEdge = "bottom"
)

var allowedNoteBackendDefinitionTickEdgeEnumValues = []NoteBackendDefinitionTickEdge{
	NOTEBACKENDDEFINITIONTICKEDGE_LEFT,
	NOTEBACKENDDEFINITIONTICKEDGE_RIGHT,
	NOTEBACKENDDEFINITIONTICKEDGE_TOP,
	NOTEBACKENDDEFINITIONTICKEDGE_BOTTOM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NoteBackendDefinitionTickEdge) GetAllowedValues() []NoteBackendDefinitionTickEdge {
	return allowedNoteBackendDefinitionTickEdgeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NoteBackendDefinitionTickEdge) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NoteBackendDefinitionTickEdge(value)
	return nil
}

// NewNoteBackendDefinitionTickEdgeFromValue returns a pointer to a valid NoteBackendDefinitionTickEdge
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNoteBackendDefinitionTickEdgeFromValue(v string) (*NoteBackendDefinitionTickEdge, error) {
	ev := NoteBackendDefinitionTickEdge(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NoteBackendDefinitionTickEdge: valid values are %v", v, allowedNoteBackendDefinitionTickEdgeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NoteBackendDefinitionTickEdge) IsValid() bool {
	for _, existing := range allowedNoteBackendDefinitionTickEdgeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NoteBackendDefinitionTickEdge value.
func (v NoteBackendDefinitionTickEdge) Ptr() *NoteBackendDefinitionTickEdge {
	return &v
}

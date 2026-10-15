// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SavedFilterCreationType How the saved filter was authored. RULES indicates rule-based authoring and LIST indicates list import. Both types return targeting rules and conditions.
type SavedFilterCreationType string

// List of SavedFilterCreationType.
const (
	SAVEDFILTERCREATIONTYPE_RULES SavedFilterCreationType = "RULES"
	SAVEDFILTERCREATIONTYPE_LIST  SavedFilterCreationType = "LIST"
)

var allowedSavedFilterCreationTypeEnumValues = []SavedFilterCreationType{
	SAVEDFILTERCREATIONTYPE_RULES,
	SAVEDFILTERCREATIONTYPE_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SavedFilterCreationType) GetAllowedValues() []SavedFilterCreationType {
	return allowedSavedFilterCreationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SavedFilterCreationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SavedFilterCreationType(value)
	return nil
}

// NewSavedFilterCreationTypeFromValue returns a pointer to a valid SavedFilterCreationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSavedFilterCreationTypeFromValue(v string) (*SavedFilterCreationType, error) {
	ev := SavedFilterCreationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SavedFilterCreationType: valid values are %v", v, allowedSavedFilterCreationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SavedFilterCreationType) IsValid() bool {
	for _, existing := range allowedSavedFilterCreationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SavedFilterCreationType value.
func (v SavedFilterCreationType) Ptr() *SavedFilterCreationType {
	return &v
}

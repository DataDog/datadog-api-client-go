// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SavedFilterDataType The saved filter resource type.
type SavedFilterDataType string

// List of SavedFilterDataType.
const (
	SAVEDFILTERDATATYPE_SAVED_FILTERS SavedFilterDataType = "saved-filters"
)

var allowedSavedFilterDataTypeEnumValues = []SavedFilterDataType{
	SAVEDFILTERDATATYPE_SAVED_FILTERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SavedFilterDataType) GetAllowedValues() []SavedFilterDataType {
	return allowedSavedFilterDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SavedFilterDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SavedFilterDataType(value)
	return nil
}

// NewSavedFilterDataTypeFromValue returns a pointer to a valid SavedFilterDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSavedFilterDataTypeFromValue(v string) (*SavedFilterDataType, error) {
	ev := SavedFilterDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SavedFilterDataType: valid values are %v", v, allowedSavedFilterDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SavedFilterDataType) IsValid() bool {
	for _, existing := range allowedSavedFilterDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SavedFilterDataType value.
func (v SavedFilterDataType) Ptr() *SavedFilterDataType {
	return &v
}

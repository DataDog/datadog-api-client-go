// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SavedFilterPermissionDeniedType The resource type for this saved filter update error.
type SavedFilterPermissionDeniedType string

// List of SavedFilterPermissionDeniedType.
const (
	SAVEDFILTERPERMISSIONDENIEDTYPE_UPDATE_GRACE_DENIED SavedFilterPermissionDeniedType = "update-grace-denied"
)

var allowedSavedFilterPermissionDeniedTypeEnumValues = []SavedFilterPermissionDeniedType{
	SAVEDFILTERPERMISSIONDENIEDTYPE_UPDATE_GRACE_DENIED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SavedFilterPermissionDeniedType) GetAllowedValues() []SavedFilterPermissionDeniedType {
	return allowedSavedFilterPermissionDeniedTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SavedFilterPermissionDeniedType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SavedFilterPermissionDeniedType(value)
	return nil
}

// NewSavedFilterPermissionDeniedTypeFromValue returns a pointer to a valid SavedFilterPermissionDeniedType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSavedFilterPermissionDeniedTypeFromValue(v string) (*SavedFilterPermissionDeniedType, error) {
	ev := SavedFilterPermissionDeniedType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SavedFilterPermissionDeniedType: valid values are %v", v, allowedSavedFilterPermissionDeniedTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SavedFilterPermissionDeniedType) IsValid() bool {
	for _, existing := range allowedSavedFilterPermissionDeniedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SavedFilterPermissionDeniedType value.
func (v SavedFilterPermissionDeniedType) Ptr() *SavedFilterPermissionDeniedType {
	return &v
}

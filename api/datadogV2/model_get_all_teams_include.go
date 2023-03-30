// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// GetAllTeamsInclude Included related resources optionally requested.
type GetAllTeamsInclude string

// List of GetAllTeamsInclude.
const (
	GETALLTEAMSINCLUDE_TEAM_LINKS            GetAllTeamsInclude = "team_links"
	GETALLTEAMSINCLUDE_USER_TEAM_PERMISSIONS GetAllTeamsInclude = "user_team_permissions"
)

var allowedGetAllTeamsIncludeEnumValues = []GetAllTeamsInclude{
	GETALLTEAMSINCLUDE_TEAM_LINKS,
	GETALLTEAMSINCLUDE_USER_TEAM_PERMISSIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GetAllTeamsInclude) GetAllowedValues() []GetAllTeamsInclude {
	return allowedGetAllTeamsIncludeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GetAllTeamsInclude) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GetAllTeamsInclude(value)
	return nil
}

// NewGetAllTeamsIncludeFromValue returns a pointer to a valid GetAllTeamsInclude
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGetAllTeamsIncludeFromValue(v string) (*GetAllTeamsInclude, error) {
	ev := GetAllTeamsInclude(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GetAllTeamsInclude: valid values are %v", v, allowedGetAllTeamsIncludeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GetAllTeamsInclude) IsValid() bool {
	for _, existing := range allowedGetAllTeamsIncludeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GetAllTeamsInclude value.
func (v GetAllTeamsInclude) Ptr() *GetAllTeamsInclude {
	return &v
}

// NullableGetAllTeamsInclude handles when a null is used for GetAllTeamsInclude.
type NullableGetAllTeamsInclude struct {
	value *GetAllTeamsInclude
	isSet bool
}

// Get returns the associated value.
func (v NullableGetAllTeamsInclude) Get() *GetAllTeamsInclude {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableGetAllTeamsInclude) Set(val *GetAllTeamsInclude) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableGetAllTeamsInclude) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableGetAllTeamsInclude) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableGetAllTeamsInclude initializes the struct as if Set has been called.
func NewNullableGetAllTeamsInclude(val *GetAllTeamsInclude) *NullableGetAllTeamsInclude {
	return &NullableGetAllTeamsInclude{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableGetAllTeamsInclude) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableGetAllTeamsInclude) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

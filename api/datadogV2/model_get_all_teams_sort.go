// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// GetAllTeamsSort Specifies the order of the returned teams
type GetAllTeamsSort string

// List of GetAllTeamsSort.
const (
	GETALLTEAMSSORT_NAME        GetAllTeamsSort = "name"
	GETALLTEAMSSORT__NAME       GetAllTeamsSort = "-name"
	GETALLTEAMSSORT_USER_COUNT  GetAllTeamsSort = "user_count"
	GETALLTEAMSSORT__USER_COUNT GetAllTeamsSort = "-user_count"
)

var allowedGetAllTeamsSortEnumValues = []GetAllTeamsSort{
	GETALLTEAMSSORT_NAME,
	GETALLTEAMSSORT__NAME,
	GETALLTEAMSSORT_USER_COUNT,
	GETALLTEAMSSORT__USER_COUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GetAllTeamsSort) GetAllowedValues() []GetAllTeamsSort {
	return allowedGetAllTeamsSortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GetAllTeamsSort) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GetAllTeamsSort(value)
	return nil
}

// NewGetAllTeamsSortFromValue returns a pointer to a valid GetAllTeamsSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGetAllTeamsSortFromValue(v string) (*GetAllTeamsSort, error) {
	ev := GetAllTeamsSort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GetAllTeamsSort: valid values are %v", v, allowedGetAllTeamsSortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GetAllTeamsSort) IsValid() bool {
	for _, existing := range allowedGetAllTeamsSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GetAllTeamsSort value.
func (v GetAllTeamsSort) Ptr() *GetAllTeamsSort {
	return &v
}

// NullableGetAllTeamsSort handles when a null is used for GetAllTeamsSort.
type NullableGetAllTeamsSort struct {
	value *GetAllTeamsSort
	isSet bool
}

// Get returns the associated value.
func (v NullableGetAllTeamsSort) Get() *GetAllTeamsSort {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableGetAllTeamsSort) Set(val *GetAllTeamsSort) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableGetAllTeamsSort) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableGetAllTeamsSort) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableGetAllTeamsSort initializes the struct as if Set has been called.
func NewNullableGetAllTeamsSort(val *GetAllTeamsSort) *NullableGetAllTeamsSort {
	return &NullableGetAllTeamsSort{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableGetAllTeamsSort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableGetAllTeamsSort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

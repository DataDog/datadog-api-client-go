/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// UserInvitationsType User invitations type.
type UserInvitationsType string

// List of UserInvitationsType
const (
	USERINVITATIONSTYPE_USER_INVITATIONS UserInvitationsType = "user_invitations"
)

func (v *UserInvitationsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserInvitationsType(value)
	for _, existing := range []UserInvitationsType{"user_invitations"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserInvitationsType", value)
}

// Ptr returns reference to UserInvitationsType value
func (v UserInvitationsType) Ptr() *UserInvitationsType {
	return &v
}

type NullableUserInvitationsType struct {
	value *UserInvitationsType
	isSet bool
}

func (v NullableUserInvitationsType) Get() *UserInvitationsType {
	return v.value
}

func (v *NullableUserInvitationsType) Set(val *UserInvitationsType) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInvitationsType) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInvitationsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInvitationsType(val *UserInvitationsType) *NullableUserInvitationsType {
	return &NullableUserInvitationsType{value: val, isSet: true}
}

func (v NullableUserInvitationsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInvitationsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

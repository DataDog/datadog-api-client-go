/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// ApplicationKeyResponseIncludedItem - An object related to an application key.
type ApplicationKeyResponseIncludedItem struct {
	User *User
	Role *Role

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// UserAsApplicationKeyResponseIncludedItem is a convenience function that returns User wrapped in ApplicationKeyResponseIncludedItem
func UserAsApplicationKeyResponseIncludedItem(v *User) ApplicationKeyResponseIncludedItem {
	return ApplicationKeyResponseIncludedItem{User: v}
}

// RoleAsApplicationKeyResponseIncludedItem is a convenience function that returns Role wrapped in ApplicationKeyResponseIncludedItem
func RoleAsApplicationKeyResponseIncludedItem(v *Role) ApplicationKeyResponseIncludedItem {
	return ApplicationKeyResponseIncludedItem{Role: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ApplicationKeyResponseIncludedItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into User
	err = json.Unmarshal(data, &dst.User)
	if err == nil {
		if dst.User != nil && dst.User.UnparsedObject == nil {
			jsonUser, _ := json.Marshal(dst.User)
			if string(jsonUser) == "{}" { // empty struct
				dst.User = nil
			} else {
				match++
			}
		} else {
			dst.User = nil
		}
	} else {
		dst.User = nil
	}

	// try to unmarshal data into Role
	err = json.Unmarshal(data, &dst.Role)
	if err == nil {
		if dst.Role != nil && dst.Role.UnparsedObject == nil {
			jsonRole, _ := json.Marshal(dst.Role)
			if string(jsonRole) == "{}" { // empty struct
				dst.Role = nil
			} else {
				match++
			}
		} else {
			dst.Role = nil
		}
	} else {
		dst.Role = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.User = nil
		dst.Role = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ApplicationKeyResponseIncludedItem) MarshalJSON() ([]byte, error) {
	if src.User != nil {
		return json.Marshal(&src.User)
	}

	if src.Role != nil {
		return json.Marshal(&src.Role)
	}

	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ApplicationKeyResponseIncludedItem) GetActualInstance() interface{} {
	if obj.User != nil {
		return obj.User
	}

	if obj.Role != nil {
		return obj.Role
	}

	// all schemas are nil
	return nil
}

type NullableApplicationKeyResponseIncludedItem struct {
	value *ApplicationKeyResponseIncludedItem
	isSet bool
}

func (v NullableApplicationKeyResponseIncludedItem) Get() *ApplicationKeyResponseIncludedItem {
	return v.value
}

func (v *NullableApplicationKeyResponseIncludedItem) Set(val *ApplicationKeyResponseIncludedItem) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationKeyResponseIncludedItem) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationKeyResponseIncludedItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationKeyResponseIncludedItem(val *ApplicationKeyResponseIncludedItem) *NullableApplicationKeyResponseIncludedItem {
	return &NullableApplicationKeyResponseIncludedItem{value: val, isSet: true}
}

func (v NullableApplicationKeyResponseIncludedItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationKeyResponseIncludedItem) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}

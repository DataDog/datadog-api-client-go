// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
)

// DowntimeResponseIncludedItem - An object related to a downtime.
type DowntimeResponseIncludedItem struct {
	User                        *User
	DowntimeMonitorIncludedItem *DowntimeMonitorIncludedItem

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// UserAsDowntimeResponseIncludedItem is a convenience function that returns User wrapped in DowntimeResponseIncludedItem.
func UserAsDowntimeResponseIncludedItem(v *User) DowntimeResponseIncludedItem {
	return DowntimeResponseIncludedItem{User: v}
}

// DowntimeMonitorIncludedItemAsDowntimeResponseIncludedItem is a convenience function that returns DowntimeMonitorIncludedItem wrapped in DowntimeResponseIncludedItem.
func DowntimeMonitorIncludedItemAsDowntimeResponseIncludedItem(v *DowntimeMonitorIncludedItem) DowntimeResponseIncludedItem {
	return DowntimeResponseIncludedItem{DowntimeMonitorIncludedItem: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DowntimeResponseIncludedItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into User
	err = json.Unmarshal(data, &obj.User)
	if err == nil {
		if obj.User != nil && obj.User.UnparsedObject == nil {
			jsonUser, _ := json.Marshal(obj.User)
			if string(jsonUser) == "{}" { // empty struct
				obj.User = nil
			} else {
				match++
			}
		} else {
			obj.User = nil
		}
	} else {
		obj.User = nil
	}

	// try to unmarshal data into DowntimeMonitorIncludedItem
	err = json.Unmarshal(data, &obj.DowntimeMonitorIncludedItem)
	if err == nil {
		if obj.DowntimeMonitorIncludedItem != nil && obj.DowntimeMonitorIncludedItem.UnparsedObject == nil {
			jsonDowntimeMonitorIncludedItem, _ := json.Marshal(obj.DowntimeMonitorIncludedItem)
			if string(jsonDowntimeMonitorIncludedItem) == "{}" { // empty struct
				obj.DowntimeMonitorIncludedItem = nil
			} else {
				match++
			}
		} else {
			obj.DowntimeMonitorIncludedItem = nil
		}
	} else {
		obj.DowntimeMonitorIncludedItem = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.User = nil
		obj.DowntimeMonitorIncludedItem = nil
		return json.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DowntimeResponseIncludedItem) MarshalJSON() ([]byte, error) {
	if obj.User != nil {
		return json.Marshal(&obj.User)
	}

	if obj.DowntimeMonitorIncludedItem != nil {
		return json.Marshal(&obj.DowntimeMonitorIncludedItem)
	}

	if obj.UnparsedObject != nil {
		return json.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DowntimeResponseIncludedItem) GetActualInstance() interface{} {
	if obj.User != nil {
		return obj.User
	}

	if obj.DowntimeMonitorIncludedItem != nil {
		return obj.DowntimeMonitorIncludedItem
	}

	// all schemas are nil
	return nil
}

// NullableDowntimeResponseIncludedItem handles when a null is used for DowntimeResponseIncludedItem.
type NullableDowntimeResponseIncludedItem struct {
	value *DowntimeResponseIncludedItem
	isSet bool
}

// Get returns the associated value.
func (v NullableDowntimeResponseIncludedItem) Get() *DowntimeResponseIncludedItem {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDowntimeResponseIncludedItem) Set(val *DowntimeResponseIncludedItem) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDowntimeResponseIncludedItem) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableDowntimeResponseIncludedItem) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDowntimeResponseIncludedItem initializes the struct as if Set has been called.
func NewNullableDowntimeResponseIncludedItem(val *DowntimeResponseIncludedItem) *NullableDowntimeResponseIncludedItem {
	return &NullableDowntimeResponseIncludedItem{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDowntimeResponseIncludedItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDowntimeResponseIncludedItem) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}

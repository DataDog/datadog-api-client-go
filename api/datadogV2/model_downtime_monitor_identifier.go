// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"
)

// DowntimeMonitorIdentifier - Monitor identifier for the downtime.
type DowntimeMonitorIdentifier struct {
	DowntimeMonitorIdentifierId   *DowntimeMonitorIdentifierId
	DowntimeMonitorIdentifierTags *DowntimeMonitorIdentifierTags

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DowntimeMonitorIdentifierIdAsDowntimeMonitorIdentifier is a convenience function that returns DowntimeMonitorIdentifierId wrapped in DowntimeMonitorIdentifier.
func DowntimeMonitorIdentifierIdAsDowntimeMonitorIdentifier(v *DowntimeMonitorIdentifierId) DowntimeMonitorIdentifier {
	return DowntimeMonitorIdentifier{DowntimeMonitorIdentifierId: v}
}

// DowntimeMonitorIdentifierTagsAsDowntimeMonitorIdentifier is a convenience function that returns DowntimeMonitorIdentifierTags wrapped in DowntimeMonitorIdentifier.
func DowntimeMonitorIdentifierTagsAsDowntimeMonitorIdentifier(v *DowntimeMonitorIdentifierTags) DowntimeMonitorIdentifier {
	return DowntimeMonitorIdentifier{DowntimeMonitorIdentifierTags: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DowntimeMonitorIdentifier) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DowntimeMonitorIdentifierId
	err = json.Unmarshal(data, &obj.DowntimeMonitorIdentifierId)
	if err == nil {
		if obj.DowntimeMonitorIdentifierId != nil && obj.DowntimeMonitorIdentifierId.UnparsedObject == nil {
			jsonDowntimeMonitorIdentifierId, _ := json.Marshal(obj.DowntimeMonitorIdentifierId)
			if string(jsonDowntimeMonitorIdentifierId) == "{}" { // empty struct
				obj.DowntimeMonitorIdentifierId = nil
			} else {
				match++
			}
		} else {
			obj.DowntimeMonitorIdentifierId = nil
		}
	} else {
		obj.DowntimeMonitorIdentifierId = nil
	}

	// try to unmarshal data into DowntimeMonitorIdentifierTags
	err = json.Unmarshal(data, &obj.DowntimeMonitorIdentifierTags)
	if err == nil {
		if obj.DowntimeMonitorIdentifierTags != nil && obj.DowntimeMonitorIdentifierTags.UnparsedObject == nil {
			jsonDowntimeMonitorIdentifierTags, _ := json.Marshal(obj.DowntimeMonitorIdentifierTags)
			if string(jsonDowntimeMonitorIdentifierTags) == "{}" { // empty struct
				obj.DowntimeMonitorIdentifierTags = nil
			} else {
				match++
			}
		} else {
			obj.DowntimeMonitorIdentifierTags = nil
		}
	} else {
		obj.DowntimeMonitorIdentifierTags = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DowntimeMonitorIdentifierId = nil
		obj.DowntimeMonitorIdentifierTags = nil
		return json.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DowntimeMonitorIdentifier) MarshalJSON() ([]byte, error) {
	if obj.DowntimeMonitorIdentifierId != nil {
		return json.Marshal(&obj.DowntimeMonitorIdentifierId)
	}

	if obj.DowntimeMonitorIdentifierTags != nil {
		return json.Marshal(&obj.DowntimeMonitorIdentifierTags)
	}

	if obj.UnparsedObject != nil {
		return json.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DowntimeMonitorIdentifier) GetActualInstance() interface{} {
	if obj.DowntimeMonitorIdentifierId != nil {
		return obj.DowntimeMonitorIdentifierId
	}

	if obj.DowntimeMonitorIdentifierTags != nil {
		return obj.DowntimeMonitorIdentifierTags
	}

	// all schemas are nil
	return nil
}

// NullableDowntimeMonitorIdentifier handles when a null is used for DowntimeMonitorIdentifier.
type NullableDowntimeMonitorIdentifier struct {
	value *DowntimeMonitorIdentifier
	isSet bool
}

// Get returns the associated value.
func (v NullableDowntimeMonitorIdentifier) Get() *DowntimeMonitorIdentifier {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDowntimeMonitorIdentifier) Set(val *DowntimeMonitorIdentifier) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDowntimeMonitorIdentifier) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableDowntimeMonitorIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDowntimeMonitorIdentifier initializes the struct as if Set has been called.
func NewNullableDowntimeMonitorIdentifier(val *DowntimeMonitorIdentifier) *NullableDowntimeMonitorIdentifier {
	return &NullableDowntimeMonitorIdentifier{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDowntimeMonitorIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDowntimeMonitorIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}

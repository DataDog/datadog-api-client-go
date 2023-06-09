// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
)

// DowntimeAttributeMonitorIdentifier - Monitor identifier for the downtime.
type DowntimeAttributeMonitorIdentifier struct {
	DowntimeAttributeMonitorIdentifierId   *DowntimeAttributeMonitorIdentifierId
	DowntimeAttributeMonitorIdentifierTags *DowntimeAttributeMonitorIdentifierTags

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DowntimeAttributeMonitorIdentifierIdAsDowntimeAttributeMonitorIdentifier is a convenience function that returns DowntimeAttributeMonitorIdentifierId wrapped in DowntimeAttributeMonitorIdentifier.
func DowntimeAttributeMonitorIdentifierIdAsDowntimeAttributeMonitorIdentifier(v *DowntimeAttributeMonitorIdentifierId) DowntimeAttributeMonitorIdentifier {
	return DowntimeAttributeMonitorIdentifier{DowntimeAttributeMonitorIdentifierId: v}
}

// DowntimeAttributeMonitorIdentifierTagsAsDowntimeAttributeMonitorIdentifier is a convenience function that returns DowntimeAttributeMonitorIdentifierTags wrapped in DowntimeAttributeMonitorIdentifier.
func DowntimeAttributeMonitorIdentifierTagsAsDowntimeAttributeMonitorIdentifier(v *DowntimeAttributeMonitorIdentifierTags) DowntimeAttributeMonitorIdentifier {
	return DowntimeAttributeMonitorIdentifier{DowntimeAttributeMonitorIdentifierTags: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DowntimeAttributeMonitorIdentifier) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DowntimeAttributeMonitorIdentifierId
	err = json.Unmarshal(data, &obj.DowntimeAttributeMonitorIdentifierId)
	if err == nil {
		if obj.DowntimeAttributeMonitorIdentifierId != nil && obj.DowntimeAttributeMonitorIdentifierId.UnparsedObject == nil {
			jsonDowntimeAttributeMonitorIdentifierId, _ := json.Marshal(obj.DowntimeAttributeMonitorIdentifierId)
			if string(jsonDowntimeAttributeMonitorIdentifierId) == "{}" { // empty struct
				obj.DowntimeAttributeMonitorIdentifierId = nil
			} else {
				match++
			}
		} else {
			obj.DowntimeAttributeMonitorIdentifierId = nil
		}
	} else {
		obj.DowntimeAttributeMonitorIdentifierId = nil
	}

	// try to unmarshal data into DowntimeAttributeMonitorIdentifierTags
	err = json.Unmarshal(data, &obj.DowntimeAttributeMonitorIdentifierTags)
	if err == nil {
		if obj.DowntimeAttributeMonitorIdentifierTags != nil && obj.DowntimeAttributeMonitorIdentifierTags.UnparsedObject == nil {
			jsonDowntimeAttributeMonitorIdentifierTags, _ := json.Marshal(obj.DowntimeAttributeMonitorIdentifierTags)
			if string(jsonDowntimeAttributeMonitorIdentifierTags) == "{}" { // empty struct
				obj.DowntimeAttributeMonitorIdentifierTags = nil
			} else {
				match++
			}
		} else {
			obj.DowntimeAttributeMonitorIdentifierTags = nil
		}
	} else {
		obj.DowntimeAttributeMonitorIdentifierTags = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DowntimeAttributeMonitorIdentifierId = nil
		obj.DowntimeAttributeMonitorIdentifierTags = nil
		return json.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DowntimeAttributeMonitorIdentifier) MarshalJSON() ([]byte, error) {
	if obj.DowntimeAttributeMonitorIdentifierId != nil {
		return json.Marshal(&obj.DowntimeAttributeMonitorIdentifierId)
	}

	if obj.DowntimeAttributeMonitorIdentifierTags != nil {
		return json.Marshal(&obj.DowntimeAttributeMonitorIdentifierTags)
	}

	if obj.UnparsedObject != nil {
		return json.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DowntimeAttributeMonitorIdentifier) GetActualInstance() interface{} {
	if obj.DowntimeAttributeMonitorIdentifierId != nil {
		return obj.DowntimeAttributeMonitorIdentifierId
	}

	if obj.DowntimeAttributeMonitorIdentifierTags != nil {
		return obj.DowntimeAttributeMonitorIdentifierTags
	}

	// all schemas are nil
	return nil
}

// NullableDowntimeAttributeMonitorIdentifier handles when a null is used for DowntimeAttributeMonitorIdentifier.
type NullableDowntimeAttributeMonitorIdentifier struct {
	value *DowntimeAttributeMonitorIdentifier
	isSet bool
}

// Get returns the associated value.
func (v NullableDowntimeAttributeMonitorIdentifier) Get() *DowntimeAttributeMonitorIdentifier {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDowntimeAttributeMonitorIdentifier) Set(val *DowntimeAttributeMonitorIdentifier) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDowntimeAttributeMonitorIdentifier) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableDowntimeAttributeMonitorIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDowntimeAttributeMonitorIdentifier initializes the struct as if Set has been called.
func NewNullableDowntimeAttributeMonitorIdentifier(val *DowntimeAttributeMonitorIdentifier) *NullableDowntimeAttributeMonitorIdentifier {
	return &NullableDowntimeAttributeMonitorIdentifier{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDowntimeAttributeMonitorIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDowntimeAttributeMonitorIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}

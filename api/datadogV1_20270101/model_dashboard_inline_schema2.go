// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema2 - Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema2 struct {
	String  *string
	Float64 *float64
	Bool    *bool

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StringAsDashboardInlineSchema2 is a convenience function that returns string wrapped in DashboardInlineSchema2.
func StringAsDashboardInlineSchema2(v *string) DashboardInlineSchema2 {
	return DashboardInlineSchema2{String: v}
}

// Float64AsDashboardInlineSchema2 is a convenience function that returns float64 wrapped in DashboardInlineSchema2.
func Float64AsDashboardInlineSchema2(v *float64) DashboardInlineSchema2 {
	return DashboardInlineSchema2{Float64: v}
}

// BoolAsDashboardInlineSchema2 is a convenience function that returns bool wrapped in DashboardInlineSchema2.
func BoolAsDashboardInlineSchema2(v *bool) DashboardInlineSchema2 {
	return DashboardInlineSchema2{Bool: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DashboardInlineSchema2) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into String
	err = datadog.Unmarshal(data, &obj.String)
	if err == nil {
		if obj.String != nil {
			jsonString, _ := datadog.Marshal(obj.String)
			if string(jsonString) == "{}" { // empty struct
				obj.String = nil
			} else {
				match++
			}
		} else {
			obj.String = nil
		}
	} else {
		obj.String = nil
	}

	// try to unmarshal data into Float64
	err = datadog.Unmarshal(data, &obj.Float64)
	if err == nil {
		if obj.Float64 != nil {
			jsonFloat64, _ := datadog.Marshal(obj.Float64)
			if string(jsonFloat64) == "{}" { // empty struct
				obj.Float64 = nil
			} else {
				match++
			}
		} else {
			obj.Float64 = nil
		}
	} else {
		obj.Float64 = nil
	}

	// try to unmarshal data into Bool
	err = datadog.Unmarshal(data, &obj.Bool)
	if err == nil {
		if obj.Bool != nil {
			jsonBool, _ := datadog.Marshal(obj.Bool)
			if string(jsonBool) == "{}" { // empty struct
				obj.Bool = nil
			} else {
				match++
			}
		} else {
			obj.Bool = nil
		}
	} else {
		obj.Bool = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.String = nil
		obj.Float64 = nil
		obj.Bool = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DashboardInlineSchema2) MarshalJSON() ([]byte, error) {
	if obj.String != nil {
		return datadog.Marshal(&obj.String)
	}

	if obj.Float64 != nil {
		return datadog.Marshal(&obj.Float64)
	}

	if obj.Bool != nil {
		return datadog.Marshal(&obj.Bool)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DashboardInlineSchema2) GetActualInstance() interface{} {
	if obj.String != nil {
		return obj.String
	}

	if obj.Float64 != nil {
		return obj.Float64
	}

	if obj.Bool != nil {
		return obj.Bool
	}

	// all schemas are nil
	return nil
}

// NullableDashboardInlineSchema2 handles when a null is used for DashboardInlineSchema2.
type NullableDashboardInlineSchema2 struct {
	value *DashboardInlineSchema2
	isSet bool
}

// Get returns the associated value.
func (v NullableDashboardInlineSchema2) Get() *DashboardInlineSchema2 {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDashboardInlineSchema2) Set(val *DashboardInlineSchema2) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDashboardInlineSchema2) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableDashboardInlineSchema2) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDashboardInlineSchema2 initializes the struct as if Set has been called.
func NewNullableDashboardInlineSchema2(val *DashboardInlineSchema2) *NullableDashboardInlineSchema2 {
	return &NullableDashboardInlineSchema2{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDashboardInlineSchema2) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDashboardInlineSchema2) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}

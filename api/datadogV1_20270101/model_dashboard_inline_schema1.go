// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema1 - Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema1 struct {
	String  *string
	Float64 *float64
	Bool    *bool

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StringAsDashboardInlineSchema1 is a convenience function that returns string wrapped in DashboardInlineSchema1.
func StringAsDashboardInlineSchema1(v *string) DashboardInlineSchema1 {
	return DashboardInlineSchema1{String: v}
}

// Float64AsDashboardInlineSchema1 is a convenience function that returns float64 wrapped in DashboardInlineSchema1.
func Float64AsDashboardInlineSchema1(v *float64) DashboardInlineSchema1 {
	return DashboardInlineSchema1{Float64: v}
}

// BoolAsDashboardInlineSchema1 is a convenience function that returns bool wrapped in DashboardInlineSchema1.
func BoolAsDashboardInlineSchema1(v *bool) DashboardInlineSchema1 {
	return DashboardInlineSchema1{Bool: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DashboardInlineSchema1) UnmarshalJSON(data []byte) error {
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
func (obj DashboardInlineSchema1) MarshalJSON() ([]byte, error) {
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
func (obj *DashboardInlineSchema1) GetActualInstance() interface{} {
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

// NullableDashboardInlineSchema1 handles when a null is used for DashboardInlineSchema1.
type NullableDashboardInlineSchema1 struct {
	value *DashboardInlineSchema1
	isSet bool
}

// Get returns the associated value.
func (v NullableDashboardInlineSchema1) Get() *DashboardInlineSchema1 {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDashboardInlineSchema1) Set(val *DashboardInlineSchema1) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDashboardInlineSchema1) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableDashboardInlineSchema1) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDashboardInlineSchema1 initializes the struct as if Set has been called.
func NewNullableDashboardInlineSchema1(val *DashboardInlineSchema1) *NullableDashboardInlineSchema1 {
	return &NullableDashboardInlineSchema1{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDashboardInlineSchema1) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDashboardInlineSchema1) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}

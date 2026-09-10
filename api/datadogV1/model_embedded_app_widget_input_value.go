// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EmbeddedAppWidgetInputValue - Value of the app input. This can be a string, number, boolean, object, or a non-empty homogeneous array of those types.
type EmbeddedAppWidgetInputValue struct {
	String                                  *string
	Float64                                 *float64
	Bool                                    *bool
	EmbeddedAppWidgetInputValueObject       map[string]interface{}
	EmbeddedAppWidgetInputValueStringArray  *EmbeddedAppWidgetInputValueStringArray
	EmbeddedAppWidgetInputValueNumberArray  *EmbeddedAppWidgetInputValueNumberArray
	EmbeddedAppWidgetInputValueBooleanArray *EmbeddedAppWidgetInputValueBooleanArray
	EmbeddedAppWidgetInputValueObjectArray  *EmbeddedAppWidgetInputValueObjectArray

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StringAsEmbeddedAppWidgetInputValue is a convenience function that returns string wrapped in EmbeddedAppWidgetInputValue.
func StringAsEmbeddedAppWidgetInputValue(v *string) EmbeddedAppWidgetInputValue {
	return EmbeddedAppWidgetInputValue{String: v}
}

// Float64AsEmbeddedAppWidgetInputValue is a convenience function that returns float64 wrapped in EmbeddedAppWidgetInputValue.
func Float64AsEmbeddedAppWidgetInputValue(v *float64) EmbeddedAppWidgetInputValue {
	return EmbeddedAppWidgetInputValue{Float64: v}
}

// BoolAsEmbeddedAppWidgetInputValue is a convenience function that returns bool wrapped in EmbeddedAppWidgetInputValue.
func BoolAsEmbeddedAppWidgetInputValue(v *bool) EmbeddedAppWidgetInputValue {
	return EmbeddedAppWidgetInputValue{Bool: v}
}

// EmbeddedAppWidgetInputValueObjectAsEmbeddedAppWidgetInputValue is a convenience function that returns map[string]interface{} wrapped in EmbeddedAppWidgetInputValue.
func EmbeddedAppWidgetInputValueObjectAsEmbeddedAppWidgetInputValue(v map[string]interface{}) EmbeddedAppWidgetInputValue {
	return EmbeddedAppWidgetInputValue{EmbeddedAppWidgetInputValueObject: v}
}

// EmbeddedAppWidgetInputValueStringArrayAsEmbeddedAppWidgetInputValue is a convenience function that returns EmbeddedAppWidgetInputValueStringArray wrapped in EmbeddedAppWidgetInputValue.
func EmbeddedAppWidgetInputValueStringArrayAsEmbeddedAppWidgetInputValue(v *EmbeddedAppWidgetInputValueStringArray) EmbeddedAppWidgetInputValue {
	return EmbeddedAppWidgetInputValue{EmbeddedAppWidgetInputValueStringArray: v}
}

// EmbeddedAppWidgetInputValueNumberArrayAsEmbeddedAppWidgetInputValue is a convenience function that returns EmbeddedAppWidgetInputValueNumberArray wrapped in EmbeddedAppWidgetInputValue.
func EmbeddedAppWidgetInputValueNumberArrayAsEmbeddedAppWidgetInputValue(v *EmbeddedAppWidgetInputValueNumberArray) EmbeddedAppWidgetInputValue {
	return EmbeddedAppWidgetInputValue{EmbeddedAppWidgetInputValueNumberArray: v}
}

// EmbeddedAppWidgetInputValueBooleanArrayAsEmbeddedAppWidgetInputValue is a convenience function that returns EmbeddedAppWidgetInputValueBooleanArray wrapped in EmbeddedAppWidgetInputValue.
func EmbeddedAppWidgetInputValueBooleanArrayAsEmbeddedAppWidgetInputValue(v *EmbeddedAppWidgetInputValueBooleanArray) EmbeddedAppWidgetInputValue {
	return EmbeddedAppWidgetInputValue{EmbeddedAppWidgetInputValueBooleanArray: v}
}

// EmbeddedAppWidgetInputValueObjectArrayAsEmbeddedAppWidgetInputValue is a convenience function that returns EmbeddedAppWidgetInputValueObjectArray wrapped in EmbeddedAppWidgetInputValue.
func EmbeddedAppWidgetInputValueObjectArrayAsEmbeddedAppWidgetInputValue(v *EmbeddedAppWidgetInputValueObjectArray) EmbeddedAppWidgetInputValue {
	return EmbeddedAppWidgetInputValue{EmbeddedAppWidgetInputValueObjectArray: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *EmbeddedAppWidgetInputValue) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into EmbeddedAppWidgetInputValueObject
	err = datadog.Unmarshal(data, &obj.EmbeddedAppWidgetInputValueObject)
	if err == nil {
		if obj.EmbeddedAppWidgetInputValueObject != nil {
			jsonEmbeddedAppWidgetInputValueObject, _ := datadog.Marshal(obj.EmbeddedAppWidgetInputValueObject)
			if string(jsonEmbeddedAppWidgetInputValueObject) == "{}" && string(data) != "{}" { // empty struct
				obj.EmbeddedAppWidgetInputValueObject = nil
			} else {
				match++
			}
		} else {
			obj.EmbeddedAppWidgetInputValueObject = nil
		}
	} else {
		obj.EmbeddedAppWidgetInputValueObject = nil
	}

	// try to unmarshal data into EmbeddedAppWidgetInputValueStringArray
	err = datadog.Unmarshal(data, &obj.EmbeddedAppWidgetInputValueStringArray)
	if err == nil {
		if obj.EmbeddedAppWidgetInputValueStringArray != nil {
			jsonEmbeddedAppWidgetInputValueStringArray, _ := datadog.Marshal(obj.EmbeddedAppWidgetInputValueStringArray)
			if string(jsonEmbeddedAppWidgetInputValueStringArray) == "{}" && string(data) != "{}" { // empty struct
				obj.EmbeddedAppWidgetInputValueStringArray = nil
			} else {
				match++
			}
		} else {
			obj.EmbeddedAppWidgetInputValueStringArray = nil
		}
	} else {
		obj.EmbeddedAppWidgetInputValueStringArray = nil
	}

	// try to unmarshal data into EmbeddedAppWidgetInputValueNumberArray
	err = datadog.Unmarshal(data, &obj.EmbeddedAppWidgetInputValueNumberArray)
	if err == nil {
		if obj.EmbeddedAppWidgetInputValueNumberArray != nil {
			jsonEmbeddedAppWidgetInputValueNumberArray, _ := datadog.Marshal(obj.EmbeddedAppWidgetInputValueNumberArray)
			if string(jsonEmbeddedAppWidgetInputValueNumberArray) == "{}" && string(data) != "{}" { // empty struct
				obj.EmbeddedAppWidgetInputValueNumberArray = nil
			} else {
				match++
			}
		} else {
			obj.EmbeddedAppWidgetInputValueNumberArray = nil
		}
	} else {
		obj.EmbeddedAppWidgetInputValueNumberArray = nil
	}

	// try to unmarshal data into EmbeddedAppWidgetInputValueBooleanArray
	err = datadog.Unmarshal(data, &obj.EmbeddedAppWidgetInputValueBooleanArray)
	if err == nil {
		if obj.EmbeddedAppWidgetInputValueBooleanArray != nil {
			jsonEmbeddedAppWidgetInputValueBooleanArray, _ := datadog.Marshal(obj.EmbeddedAppWidgetInputValueBooleanArray)
			if string(jsonEmbeddedAppWidgetInputValueBooleanArray) == "{}" && string(data) != "{}" { // empty struct
				obj.EmbeddedAppWidgetInputValueBooleanArray = nil
			} else {
				match++
			}
		} else {
			obj.EmbeddedAppWidgetInputValueBooleanArray = nil
		}
	} else {
		obj.EmbeddedAppWidgetInputValueBooleanArray = nil
	}

	// try to unmarshal data into EmbeddedAppWidgetInputValueObjectArray
	err = datadog.Unmarshal(data, &obj.EmbeddedAppWidgetInputValueObjectArray)
	if err == nil {
		if obj.EmbeddedAppWidgetInputValueObjectArray != nil {
			jsonEmbeddedAppWidgetInputValueObjectArray, _ := datadog.Marshal(obj.EmbeddedAppWidgetInputValueObjectArray)
			if string(jsonEmbeddedAppWidgetInputValueObjectArray) == "{}" && string(data) != "{}" { // empty struct
				obj.EmbeddedAppWidgetInputValueObjectArray = nil
			} else {
				match++
			}
		} else {
			obj.EmbeddedAppWidgetInputValueObjectArray = nil
		}
	} else {
		obj.EmbeddedAppWidgetInputValueObjectArray = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.String = nil
		obj.Float64 = nil
		obj.Bool = nil
		obj.EmbeddedAppWidgetInputValueObject = nil
		obj.EmbeddedAppWidgetInputValueStringArray = nil
		obj.EmbeddedAppWidgetInputValueNumberArray = nil
		obj.EmbeddedAppWidgetInputValueBooleanArray = nil
		obj.EmbeddedAppWidgetInputValueObjectArray = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj EmbeddedAppWidgetInputValue) MarshalJSON() ([]byte, error) {
	if obj.String != nil {
		return datadog.Marshal(&obj.String)
	}

	if obj.Float64 != nil {
		return datadog.Marshal(&obj.Float64)
	}

	if obj.Bool != nil {
		return datadog.Marshal(&obj.Bool)
	}

	if obj.EmbeddedAppWidgetInputValueObject != nil {
		return datadog.Marshal(&obj.EmbeddedAppWidgetInputValueObject)
	}

	if obj.EmbeddedAppWidgetInputValueStringArray != nil {
		return datadog.Marshal(&obj.EmbeddedAppWidgetInputValueStringArray)
	}

	if obj.EmbeddedAppWidgetInputValueNumberArray != nil {
		return datadog.Marshal(&obj.EmbeddedAppWidgetInputValueNumberArray)
	}

	if obj.EmbeddedAppWidgetInputValueBooleanArray != nil {
		return datadog.Marshal(&obj.EmbeddedAppWidgetInputValueBooleanArray)
	}

	if obj.EmbeddedAppWidgetInputValueObjectArray != nil {
		return datadog.Marshal(&obj.EmbeddedAppWidgetInputValueObjectArray)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *EmbeddedAppWidgetInputValue) GetActualInstance() interface{} {
	if obj.String != nil {
		return obj.String
	}

	if obj.Float64 != nil {
		return obj.Float64
	}

	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.EmbeddedAppWidgetInputValueObject != nil {
		return obj.EmbeddedAppWidgetInputValueObject
	}

	if obj.EmbeddedAppWidgetInputValueStringArray != nil {
		return obj.EmbeddedAppWidgetInputValueStringArray
	}

	if obj.EmbeddedAppWidgetInputValueNumberArray != nil {
		return obj.EmbeddedAppWidgetInputValueNumberArray
	}

	if obj.EmbeddedAppWidgetInputValueBooleanArray != nil {
		return obj.EmbeddedAppWidgetInputValueBooleanArray
	}

	if obj.EmbeddedAppWidgetInputValueObjectArray != nil {
		return obj.EmbeddedAppWidgetInputValueObjectArray
	}

	// all schemas are nil
	return nil
}

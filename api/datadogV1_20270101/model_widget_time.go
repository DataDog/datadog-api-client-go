// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

// WidgetTime - Test-only generated schema for the dashboard versioning proof of concept.
type WidgetTime struct {
	WidgetTime20270101OneOf           *WidgetTime20270101OneOf
	WidgetTime20270101OneOf4464103552 *WidgetTime20270101OneOf4464103552
	WidgetTime20270101OneOf4464104512 *WidgetTime20270101OneOf4464104512
	WidgetTime20270101OneOf4464105312 *WidgetTime20270101OneOf4464105312
	WidgetTime20270101OneOf4464105952 *datadogV1.WidgetTime20270101OneOf4464105952
	EmptyObject_20270101              *datadogV1.EmptyObject_20270101

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// WidgetTime20270101OneOfAsWidgetTime is a convenience function that returns WidgetTime20270101OneOf wrapped in WidgetTime.
func WidgetTime20270101OneOfAsWidgetTime(v *WidgetTime20270101OneOf) WidgetTime {
	return WidgetTime{WidgetTime20270101OneOf: v}
}

// WidgetTime20270101OneOf4464103552AsWidgetTime is a convenience function that returns WidgetTime20270101OneOf4464103552 wrapped in WidgetTime.
func WidgetTime20270101OneOf4464103552AsWidgetTime(v *WidgetTime20270101OneOf4464103552) WidgetTime {
	return WidgetTime{WidgetTime20270101OneOf4464103552: v}
}

// WidgetTime20270101OneOf4464104512AsWidgetTime is a convenience function that returns WidgetTime20270101OneOf4464104512 wrapped in WidgetTime.
func WidgetTime20270101OneOf4464104512AsWidgetTime(v *WidgetTime20270101OneOf4464104512) WidgetTime {
	return WidgetTime{WidgetTime20270101OneOf4464104512: v}
}

// WidgetTime20270101OneOf4464105312AsWidgetTime is a convenience function that returns WidgetTime20270101OneOf4464105312 wrapped in WidgetTime.
func WidgetTime20270101OneOf4464105312AsWidgetTime(v *WidgetTime20270101OneOf4464105312) WidgetTime {
	return WidgetTime{WidgetTime20270101OneOf4464105312: v}
}

// WidgetTime20270101OneOf4464105952AsWidgetTime is a convenience function that returns datadogV1.WidgetTime20270101OneOf4464105952 wrapped in WidgetTime.
func WidgetTime20270101OneOf4464105952AsWidgetTime(v *datadogV1.WidgetTime20270101OneOf4464105952) WidgetTime {
	return WidgetTime{WidgetTime20270101OneOf4464105952: v}
}

// EmptyObject_20270101AsWidgetTime is a convenience function that returns datadogV1.EmptyObject_20270101 wrapped in WidgetTime.
func EmptyObject_20270101AsWidgetTime(v *datadogV1.EmptyObject_20270101) WidgetTime {
	return WidgetTime{EmptyObject_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *WidgetTime) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into WidgetTime20270101OneOf
	err = datadog.Unmarshal(data, &obj.WidgetTime20270101OneOf)
	if err == nil {
		if obj.WidgetTime20270101OneOf != nil && obj.WidgetTime20270101OneOf.UnparsedObject == nil {
			jsonWidgetTime20270101OneOf, _ := datadog.Marshal(obj.WidgetTime20270101OneOf)
			if string(jsonWidgetTime20270101OneOf) == "{}" { // empty struct
				obj.WidgetTime20270101OneOf = nil
			} else {
				match++
			}
		} else {
			obj.WidgetTime20270101OneOf = nil
		}
	} else {
		obj.WidgetTime20270101OneOf = nil
	}

	// try to unmarshal data into WidgetTime20270101OneOf4464103552
	err = datadog.Unmarshal(data, &obj.WidgetTime20270101OneOf4464103552)
	if err == nil {
		if obj.WidgetTime20270101OneOf4464103552 != nil && obj.WidgetTime20270101OneOf4464103552.UnparsedObject == nil {
			jsonWidgetTime20270101OneOf4464103552, _ := datadog.Marshal(obj.WidgetTime20270101OneOf4464103552)
			if string(jsonWidgetTime20270101OneOf4464103552) == "{}" { // empty struct
				obj.WidgetTime20270101OneOf4464103552 = nil
			} else {
				match++
			}
		} else {
			obj.WidgetTime20270101OneOf4464103552 = nil
		}
	} else {
		obj.WidgetTime20270101OneOf4464103552 = nil
	}

	// try to unmarshal data into WidgetTime20270101OneOf4464104512
	err = datadog.Unmarshal(data, &obj.WidgetTime20270101OneOf4464104512)
	if err == nil {
		if obj.WidgetTime20270101OneOf4464104512 != nil && obj.WidgetTime20270101OneOf4464104512.UnparsedObject == nil {
			jsonWidgetTime20270101OneOf4464104512, _ := datadog.Marshal(obj.WidgetTime20270101OneOf4464104512)
			if string(jsonWidgetTime20270101OneOf4464104512) == "{}" { // empty struct
				obj.WidgetTime20270101OneOf4464104512 = nil
			} else {
				match++
			}
		} else {
			obj.WidgetTime20270101OneOf4464104512 = nil
		}
	} else {
		obj.WidgetTime20270101OneOf4464104512 = nil
	}

	// try to unmarshal data into WidgetTime20270101OneOf4464105312
	err = datadog.Unmarshal(data, &obj.WidgetTime20270101OneOf4464105312)
	if err == nil {
		if obj.WidgetTime20270101OneOf4464105312 != nil && obj.WidgetTime20270101OneOf4464105312.UnparsedObject == nil {
			jsonWidgetTime20270101OneOf4464105312, _ := datadog.Marshal(obj.WidgetTime20270101OneOf4464105312)
			if string(jsonWidgetTime20270101OneOf4464105312) == "{}" { // empty struct
				obj.WidgetTime20270101OneOf4464105312 = nil
			} else {
				match++
			}
		} else {
			obj.WidgetTime20270101OneOf4464105312 = nil
		}
	} else {
		obj.WidgetTime20270101OneOf4464105312 = nil
	}

	// try to unmarshal data into WidgetTime20270101OneOf4464105952
	err = datadog.Unmarshal(data, &obj.WidgetTime20270101OneOf4464105952)
	if err == nil {
		if obj.WidgetTime20270101OneOf4464105952 != nil && obj.WidgetTime20270101OneOf4464105952.UnparsedObject == nil {
			jsonWidgetTime20270101OneOf4464105952, _ := datadog.Marshal(obj.WidgetTime20270101OneOf4464105952)
			if string(jsonWidgetTime20270101OneOf4464105952) == "{}" && string(data) != "{}" { // empty struct
				obj.WidgetTime20270101OneOf4464105952 = nil
			} else {
				match++
			}
		} else {
			obj.WidgetTime20270101OneOf4464105952 = nil
		}
	} else {
		obj.WidgetTime20270101OneOf4464105952 = nil
	}

	// try to unmarshal data into EmptyObject_20270101
	err = datadog.Unmarshal(data, &obj.EmptyObject_20270101)
	if err == nil {
		if obj.EmptyObject_20270101 != nil && obj.EmptyObject_20270101.UnparsedObject == nil {
			jsonEmptyObject_20270101, _ := datadog.Marshal(obj.EmptyObject_20270101)
			if string(jsonEmptyObject_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.EmptyObject_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.EmptyObject_20270101 = nil
		}
	} else {
		obj.EmptyObject_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.WidgetTime20270101OneOf = nil
		obj.WidgetTime20270101OneOf4464103552 = nil
		obj.WidgetTime20270101OneOf4464104512 = nil
		obj.WidgetTime20270101OneOf4464105312 = nil
		obj.WidgetTime20270101OneOf4464105952 = nil
		obj.EmptyObject_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj WidgetTime) MarshalJSON() ([]byte, error) {
	if obj.WidgetTime20270101OneOf != nil {
		return datadog.Marshal(&obj.WidgetTime20270101OneOf)
	}

	if obj.WidgetTime20270101OneOf4464103552 != nil {
		return datadog.Marshal(&obj.WidgetTime20270101OneOf4464103552)
	}

	if obj.WidgetTime20270101OneOf4464104512 != nil {
		return datadog.Marshal(&obj.WidgetTime20270101OneOf4464104512)
	}

	if obj.WidgetTime20270101OneOf4464105312 != nil {
		return datadog.Marshal(&obj.WidgetTime20270101OneOf4464105312)
	}

	if obj.WidgetTime20270101OneOf4464105952 != nil {
		return datadog.Marshal(&obj.WidgetTime20270101OneOf4464105952)
	}

	if obj.EmptyObject_20270101 != nil {
		return datadog.Marshal(&obj.EmptyObject_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *WidgetTime) GetActualInstance() interface{} {
	if obj.WidgetTime20270101OneOf != nil {
		return obj.WidgetTime20270101OneOf
	}

	if obj.WidgetTime20270101OneOf4464103552 != nil {
		return obj.WidgetTime20270101OneOf4464103552
	}

	if obj.WidgetTime20270101OneOf4464104512 != nil {
		return obj.WidgetTime20270101OneOf4464104512
	}

	if obj.WidgetTime20270101OneOf4464105312 != nil {
		return obj.WidgetTime20270101OneOf4464105312
	}

	if obj.WidgetTime20270101OneOf4464105952 != nil {
		return obj.WidgetTime20270101OneOf4464105952
	}

	if obj.EmptyObject_20270101 != nil {
		return obj.EmptyObject_20270101
	}

	// all schemas are nil
	return nil
}

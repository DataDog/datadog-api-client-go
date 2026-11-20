// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NumberFormatUnit - Test-only generated schema for the dashboard versioning proof of concept.
type NumberFormatUnit struct {
	NumberFormatUnitCanonical_20270101 *NumberFormatUnitCanonical
	NumberFormatUnitCustom_20270101    *NumberFormatUnitCustom

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// NumberFormatUnitCanonical_20270101AsNumberFormatUnit is a convenience function that returns NumberFormatUnitCanonical wrapped in NumberFormatUnit.
func NumberFormatUnitCanonical_20270101AsNumberFormatUnit(v *NumberFormatUnitCanonical) NumberFormatUnit {
	return NumberFormatUnit{NumberFormatUnitCanonical_20270101: v}
}

// NumberFormatUnitCustom_20270101AsNumberFormatUnit is a convenience function that returns NumberFormatUnitCustom wrapped in NumberFormatUnit.
func NumberFormatUnitCustom_20270101AsNumberFormatUnit(v *NumberFormatUnitCustom) NumberFormatUnit {
	return NumberFormatUnit{NumberFormatUnitCustom_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *NumberFormatUnit) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NumberFormatUnitCanonical_20270101
	err = datadog.Unmarshal(data, &obj.NumberFormatUnitCanonical_20270101)
	if err == nil {
		if obj.NumberFormatUnitCanonical_20270101 != nil && obj.NumberFormatUnitCanonical_20270101.UnparsedObject == nil {
			jsonNumberFormatUnitCanonical_20270101, _ := datadog.Marshal(obj.NumberFormatUnitCanonical_20270101)
			if string(jsonNumberFormatUnitCanonical_20270101) == "{}" { // empty struct
				obj.NumberFormatUnitCanonical_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.NumberFormatUnitCanonical_20270101 = nil
		}
	} else {
		obj.NumberFormatUnitCanonical_20270101 = nil
	}

	// try to unmarshal data into NumberFormatUnitCustom_20270101
	err = datadog.Unmarshal(data, &obj.NumberFormatUnitCustom_20270101)
	if err == nil {
		if obj.NumberFormatUnitCustom_20270101 != nil && obj.NumberFormatUnitCustom_20270101.UnparsedObject == nil {
			jsonNumberFormatUnitCustom_20270101, _ := datadog.Marshal(obj.NumberFormatUnitCustom_20270101)
			if string(jsonNumberFormatUnitCustom_20270101) == "{}" { // empty struct
				obj.NumberFormatUnitCustom_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.NumberFormatUnitCustom_20270101 = nil
		}
	} else {
		obj.NumberFormatUnitCustom_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.NumberFormatUnitCanonical_20270101 = nil
		obj.NumberFormatUnitCustom_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj NumberFormatUnit) MarshalJSON() ([]byte, error) {
	if obj.NumberFormatUnitCanonical_20270101 != nil {
		return datadog.Marshal(&obj.NumberFormatUnitCanonical_20270101)
	}

	if obj.NumberFormatUnitCustom_20270101 != nil {
		return datadog.Marshal(&obj.NumberFormatUnitCustom_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *NumberFormatUnit) GetActualInstance() interface{} {
	if obj.NumberFormatUnitCanonical_20270101 != nil {
		return obj.NumberFormatUnitCanonical_20270101
	}

	if obj.NumberFormatUnitCustom_20270101 != nil {
		return obj.NumberFormatUnitCustom_20270101
	}

	// all schemas are nil
	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaUserJourneyQuery20270101Search - Test-only generated schema for the dashboard versioning proof of concept.
type FormulaUserJourneyQuery20270101Search struct {
	LegacyUserJourneySearch_20270101 *LegacyUserJourneySearch
	UserJourneySearch_20270101       *UserJourneySearch

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// LegacyUserJourneySearch_20270101AsFormulaUserJourneyQuery20270101Search is a convenience function that returns LegacyUserJourneySearch wrapped in FormulaUserJourneyQuery20270101Search.
func LegacyUserJourneySearch_20270101AsFormulaUserJourneyQuery20270101Search(v *LegacyUserJourneySearch) FormulaUserJourneyQuery20270101Search {
	return FormulaUserJourneyQuery20270101Search{LegacyUserJourneySearch_20270101: v}
}

// UserJourneySearch_20270101AsFormulaUserJourneyQuery20270101Search is a convenience function that returns UserJourneySearch wrapped in FormulaUserJourneyQuery20270101Search.
func UserJourneySearch_20270101AsFormulaUserJourneyQuery20270101Search(v *UserJourneySearch) FormulaUserJourneyQuery20270101Search {
	return FormulaUserJourneyQuery20270101Search{UserJourneySearch_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *FormulaUserJourneyQuery20270101Search) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LegacyUserJourneySearch_20270101
	err = datadog.Unmarshal(data, &obj.LegacyUserJourneySearch_20270101)
	if err == nil {
		if obj.LegacyUserJourneySearch_20270101 != nil && obj.LegacyUserJourneySearch_20270101.UnparsedObject == nil {
			jsonLegacyUserJourneySearch_20270101, _ := datadog.Marshal(obj.LegacyUserJourneySearch_20270101)
			if string(jsonLegacyUserJourneySearch_20270101) == "{}" { // empty struct
				obj.LegacyUserJourneySearch_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.LegacyUserJourneySearch_20270101 = nil
		}
	} else {
		obj.LegacyUserJourneySearch_20270101 = nil
	}

	// try to unmarshal data into UserJourneySearch_20270101
	err = datadog.Unmarshal(data, &obj.UserJourneySearch_20270101)
	if err == nil {
		if obj.UserJourneySearch_20270101 != nil && obj.UserJourneySearch_20270101.UnparsedObject == nil {
			jsonUserJourneySearch_20270101, _ := datadog.Marshal(obj.UserJourneySearch_20270101)
			if string(jsonUserJourneySearch_20270101) == "{}" { // empty struct
				obj.UserJourneySearch_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.UserJourneySearch_20270101 = nil
		}
	} else {
		obj.UserJourneySearch_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.LegacyUserJourneySearch_20270101 = nil
		obj.UserJourneySearch_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj FormulaUserJourneyQuery20270101Search) MarshalJSON() ([]byte, error) {
	if obj.LegacyUserJourneySearch_20270101 != nil {
		return datadog.Marshal(&obj.LegacyUserJourneySearch_20270101)
	}

	if obj.UserJourneySearch_20270101 != nil {
		return datadog.Marshal(&obj.UserJourneySearch_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *FormulaUserJourneyQuery20270101Search) GetActualInstance() interface{} {
	if obj.LegacyUserJourneySearch_20270101 != nil {
		return obj.LegacyUserJourneySearch_20270101
	}

	if obj.UserJourneySearch_20270101 != nil {
		return obj.UserJourneySearch_20270101
	}

	// all schemas are nil
	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeverityOverrideAttributes - Severity override to apply to the findings.
// Set `action` to `set` to apply a manual severity override with the given `value`.
// Set `action` to `clear` to remove a manual severity override.
type SeverityOverrideAttributes struct {
	SeverityOverrideSet   *SeverityOverrideSet
	SeverityOverrideClear *SeverityOverrideClear

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SeverityOverrideSetAsSeverityOverrideAttributes is a convenience function that returns SeverityOverrideSet wrapped in SeverityOverrideAttributes.
func SeverityOverrideSetAsSeverityOverrideAttributes(v *SeverityOverrideSet) SeverityOverrideAttributes {
	return SeverityOverrideAttributes{SeverityOverrideSet: v}
}

// SeverityOverrideClearAsSeverityOverrideAttributes is a convenience function that returns SeverityOverrideClear wrapped in SeverityOverrideAttributes.
func SeverityOverrideClearAsSeverityOverrideAttributes(v *SeverityOverrideClear) SeverityOverrideAttributes {
	return SeverityOverrideAttributes{SeverityOverrideClear: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SeverityOverrideAttributes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SeverityOverrideSet
	err = datadog.Unmarshal(data, &obj.SeverityOverrideSet)
	if err == nil {
		if obj.SeverityOverrideSet != nil && obj.SeverityOverrideSet.UnparsedObject == nil {
			jsonSeverityOverrideSet, _ := datadog.Marshal(obj.SeverityOverrideSet)
			if string(jsonSeverityOverrideSet) == "{}" { // empty struct
				obj.SeverityOverrideSet = nil
			} else {
				match++
			}
		} else {
			obj.SeverityOverrideSet = nil
		}
	} else {
		obj.SeverityOverrideSet = nil
	}

	// try to unmarshal data into SeverityOverrideClear
	err = datadog.Unmarshal(data, &obj.SeverityOverrideClear)
	if err == nil {
		if obj.SeverityOverrideClear != nil && obj.SeverityOverrideClear.UnparsedObject == nil {
			jsonSeverityOverrideClear, _ := datadog.Marshal(obj.SeverityOverrideClear)
			if string(jsonSeverityOverrideClear) == "{}" { // empty struct
				obj.SeverityOverrideClear = nil
			} else {
				match++
			}
		} else {
			obj.SeverityOverrideClear = nil
		}
	} else {
		obj.SeverityOverrideClear = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SeverityOverrideSet = nil
		obj.SeverityOverrideClear = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SeverityOverrideAttributes) MarshalJSON() ([]byte, error) {
	if obj.SeverityOverrideSet != nil {
		return datadog.Marshal(&obj.SeverityOverrideSet)
	}

	if obj.SeverityOverrideClear != nil {
		return datadog.Marshal(&obj.SeverityOverrideClear)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SeverityOverrideAttributes) GetActualInstance() interface{} {
	if obj.SeverityOverrideSet != nil {
		return obj.SeverityOverrideSet
	}

	if obj.SeverityOverrideClear != nil {
		return obj.SeverityOverrideClear
	}

	// all schemas are nil
	return nil
}

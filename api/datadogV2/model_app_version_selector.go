// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AppVersionSelector - The version selector parameter used in endpoints such as Get App. Version numbers start at 1 and increment with each update. 0 is a special value that always selects the latest version.
type AppVersionSelector struct {
	AppVersionSelectorConstants *AppVersionSelectorConstants
	Int64                       *int64

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AppVersionSelectorConstantsAsAppVersionSelector is a convenience function that returns AppVersionSelectorConstants wrapped in AppVersionSelector.
func AppVersionSelectorConstantsAsAppVersionSelector(v *AppVersionSelectorConstants) AppVersionSelector {
	return AppVersionSelector{AppVersionSelectorConstants: v}
}

// Int64AsAppVersionSelector is a convenience function that returns int64 wrapped in AppVersionSelector.
func Int64AsAppVersionSelector(v *int64) AppVersionSelector {
	return AppVersionSelector{Int64: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AppVersionSelector) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AppVersionSelectorConstants
	err = datadog.Unmarshal(data, &obj.AppVersionSelectorConstants)
	if err == nil {
		if obj.AppVersionSelectorConstants != nil {
			jsonAppVersionSelectorConstants, _ := datadog.Marshal(obj.AppVersionSelectorConstants)
			if string(jsonAppVersionSelectorConstants) == "{}" && string(data) != "{}" { // empty struct
				obj.AppVersionSelectorConstants = nil
			} else {
				match++
			}
		} else {
			obj.AppVersionSelectorConstants = nil
		}
	} else {
		obj.AppVersionSelectorConstants = nil
	}

	// try to unmarshal data into Int64
	err = datadog.Unmarshal(data, &obj.Int64)
	if err == nil {
		if obj.Int64 != nil {
			jsonInt64, _ := datadog.Marshal(obj.Int64)
			if string(jsonInt64) == "{}" { // empty struct
				obj.Int64 = nil
			} else {
				match++
			}
		} else {
			obj.Int64 = nil
		}
	} else {
		obj.Int64 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AppVersionSelectorConstants = nil
		obj.Int64 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AppVersionSelector) MarshalJSON() ([]byte, error) {
	if obj.AppVersionSelectorConstants != nil {
		return datadog.Marshal(&obj.AppVersionSelectorConstants)
	}

	if obj.Int64 != nil {
		return datadog.Marshal(&obj.Int64)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AppVersionSelector) GetActualInstance() interface{} {
	if obj.AppVersionSelectorConstants != nil {
		return obj.AppVersionSelectorConstants
	}

	if obj.Int64 != nil {
		return obj.Int64
	}

	// all schemas are nil
	return nil
}

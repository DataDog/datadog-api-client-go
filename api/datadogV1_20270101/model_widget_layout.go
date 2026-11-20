// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetLayout - Test-only generated schema for the dashboard versioning proof of concept.
type WidgetLayout struct {
	FreeLayout_20270101             *FreeLayout
	MSLWidgetLayout_20270101        *MSLWidgetLayout
	MSLBackendWidgetLayout_20270101 *MSLBackendWidgetLayout

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// FreeLayout_20270101AsWidgetLayout is a convenience function that returns FreeLayout wrapped in WidgetLayout.
func FreeLayout_20270101AsWidgetLayout(v *FreeLayout) WidgetLayout {
	return WidgetLayout{FreeLayout_20270101: v}
}

// MSLWidgetLayout_20270101AsWidgetLayout is a convenience function that returns MSLWidgetLayout wrapped in WidgetLayout.
func MSLWidgetLayout_20270101AsWidgetLayout(v *MSLWidgetLayout) WidgetLayout {
	return WidgetLayout{MSLWidgetLayout_20270101: v}
}

// MSLBackendWidgetLayout_20270101AsWidgetLayout is a convenience function that returns MSLBackendWidgetLayout wrapped in WidgetLayout.
func MSLBackendWidgetLayout_20270101AsWidgetLayout(v *MSLBackendWidgetLayout) WidgetLayout {
	return WidgetLayout{MSLBackendWidgetLayout_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *WidgetLayout) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FreeLayout_20270101
	err = datadog.Unmarshal(data, &obj.FreeLayout_20270101)
	if err == nil {
		if obj.FreeLayout_20270101 != nil && obj.FreeLayout_20270101.UnparsedObject == nil {
			jsonFreeLayout_20270101, _ := datadog.Marshal(obj.FreeLayout_20270101)
			if string(jsonFreeLayout_20270101) == "{}" { // empty struct
				obj.FreeLayout_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FreeLayout_20270101 = nil
		}
	} else {
		obj.FreeLayout_20270101 = nil
	}

	// try to unmarshal data into MSLWidgetLayout_20270101
	err = datadog.Unmarshal(data, &obj.MSLWidgetLayout_20270101)
	if err == nil {
		if obj.MSLWidgetLayout_20270101 != nil && obj.MSLWidgetLayout_20270101.UnparsedObject == nil {
			jsonMSLWidgetLayout_20270101, _ := datadog.Marshal(obj.MSLWidgetLayout_20270101)
			if string(jsonMSLWidgetLayout_20270101) == "{}" { // empty struct
				obj.MSLWidgetLayout_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.MSLWidgetLayout_20270101 = nil
		}
	} else {
		obj.MSLWidgetLayout_20270101 = nil
	}

	// try to unmarshal data into MSLBackendWidgetLayout_20270101
	err = datadog.Unmarshal(data, &obj.MSLBackendWidgetLayout_20270101)
	if err == nil {
		if obj.MSLBackendWidgetLayout_20270101 != nil && obj.MSLBackendWidgetLayout_20270101.UnparsedObject == nil {
			jsonMSLBackendWidgetLayout_20270101, _ := datadog.Marshal(obj.MSLBackendWidgetLayout_20270101)
			if string(jsonMSLBackendWidgetLayout_20270101) == "{}" { // empty struct
				obj.MSLBackendWidgetLayout_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.MSLBackendWidgetLayout_20270101 = nil
		}
	} else {
		obj.MSLBackendWidgetLayout_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.FreeLayout_20270101 = nil
		obj.MSLWidgetLayout_20270101 = nil
		obj.MSLBackendWidgetLayout_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj WidgetLayout) MarshalJSON() ([]byte, error) {
	if obj.FreeLayout_20270101 != nil {
		return datadog.Marshal(&obj.FreeLayout_20270101)
	}

	if obj.MSLWidgetLayout_20270101 != nil {
		return datadog.Marshal(&obj.MSLWidgetLayout_20270101)
	}

	if obj.MSLBackendWidgetLayout_20270101 != nil {
		return datadog.Marshal(&obj.MSLBackendWidgetLayout_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *WidgetLayout) GetActualInstance() interface{} {
	if obj.FreeLayout_20270101 != nil {
		return obj.FreeLayout_20270101
	}

	if obj.MSLWidgetLayout_20270101 != nil {
		return obj.MSLWidgetLayout_20270101
	}

	if obj.MSLBackendWidgetLayout_20270101 != nil {
		return obj.MSLBackendWidgetLayout_20270101
	}

	// all schemas are nil
	return nil
}

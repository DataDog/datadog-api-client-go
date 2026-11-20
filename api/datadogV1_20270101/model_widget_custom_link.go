// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetCustomLink - Test-only generated schema for the dashboard versioning proof of concept.
type WidgetCustomLink struct {
	ContextMenuLink_20270101 *ContextMenuLink
	UserDefinedLink_20270101 *UserDefinedLink

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ContextMenuLink_20270101AsWidgetCustomLink is a convenience function that returns ContextMenuLink wrapped in WidgetCustomLink.
func ContextMenuLink_20270101AsWidgetCustomLink(v *ContextMenuLink) WidgetCustomLink {
	return WidgetCustomLink{ContextMenuLink_20270101: v}
}

// UserDefinedLink_20270101AsWidgetCustomLink is a convenience function that returns UserDefinedLink wrapped in WidgetCustomLink.
func UserDefinedLink_20270101AsWidgetCustomLink(v *UserDefinedLink) WidgetCustomLink {
	return WidgetCustomLink{UserDefinedLink_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *WidgetCustomLink) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ContextMenuLink_20270101
	err = datadog.Unmarshal(data, &obj.ContextMenuLink_20270101)
	if err == nil {
		if obj.ContextMenuLink_20270101 != nil && obj.ContextMenuLink_20270101.UnparsedObject == nil {
			jsonContextMenuLink_20270101, _ := datadog.Marshal(obj.ContextMenuLink_20270101)
			if string(jsonContextMenuLink_20270101) == "{}" { // empty struct
				obj.ContextMenuLink_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ContextMenuLink_20270101 = nil
		}
	} else {
		obj.ContextMenuLink_20270101 = nil
	}

	// try to unmarshal data into UserDefinedLink_20270101
	err = datadog.Unmarshal(data, &obj.UserDefinedLink_20270101)
	if err == nil {
		if obj.UserDefinedLink_20270101 != nil && obj.UserDefinedLink_20270101.UnparsedObject == nil {
			jsonUserDefinedLink_20270101, _ := datadog.Marshal(obj.UserDefinedLink_20270101)
			if string(jsonUserDefinedLink_20270101) == "{}" { // empty struct
				obj.UserDefinedLink_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.UserDefinedLink_20270101 = nil
		}
	} else {
		obj.UserDefinedLink_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ContextMenuLink_20270101 = nil
		obj.UserDefinedLink_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj WidgetCustomLink) MarshalJSON() ([]byte, error) {
	if obj.ContextMenuLink_20270101 != nil {
		return datadog.Marshal(&obj.ContextMenuLink_20270101)
	}

	if obj.UserDefinedLink_20270101 != nil {
		return datadog.Marshal(&obj.UserDefinedLink_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *WidgetCustomLink) GetActualInstance() interface{} {
	if obj.ContextMenuLink_20270101 != nil {
		return obj.ContextMenuLink_20270101
	}

	if obj.UserDefinedLink_20270101 != nil {
		return obj.UserDefinedLink_20270101
	}

	// all schemas are nil
	return nil
}

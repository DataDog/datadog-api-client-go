// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableWidgetDefinitionRequestsItem -
type TableWidgetDefinitionRequestsItem struct {
	TableWidgetRequest *TableWidgetRequest
	GuidedTableRequest *GuidedTableRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TableWidgetRequestAsTableWidgetDefinitionRequestsItem is a convenience function that returns TableWidgetRequest wrapped in TableWidgetDefinitionRequestsItem.
func TableWidgetRequestAsTableWidgetDefinitionRequestsItem(v *TableWidgetRequest) TableWidgetDefinitionRequestsItem {
	return TableWidgetDefinitionRequestsItem{TableWidgetRequest: v}
}

// GuidedTableRequestAsTableWidgetDefinitionRequestsItem is a convenience function that returns GuidedTableRequest wrapped in TableWidgetDefinitionRequestsItem.
func GuidedTableRequestAsTableWidgetDefinitionRequestsItem(v *GuidedTableRequest) TableWidgetDefinitionRequestsItem {
	return TableWidgetDefinitionRequestsItem{GuidedTableRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TableWidgetDefinitionRequestsItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TableWidgetRequest
	err = datadog.Unmarshal(data, &obj.TableWidgetRequest)
	if err == nil {
		if obj.TableWidgetRequest != nil && obj.TableWidgetRequest.UnparsedObject == nil {
			jsonTableWidgetRequest, _ := datadog.Marshal(obj.TableWidgetRequest)
			if string(jsonTableWidgetRequest) == "{}" && string(data) != "{}" { // empty struct
				obj.TableWidgetRequest = nil
			} else {
				match++
			}
		} else {
			obj.TableWidgetRequest = nil
		}
	} else {
		obj.TableWidgetRequest = nil
	}

	// try to unmarshal data into GuidedTableRequest
	err = datadog.Unmarshal(data, &obj.GuidedTableRequest)
	if err == nil {
		if obj.GuidedTableRequest != nil && obj.GuidedTableRequest.UnparsedObject == nil {
			jsonGuidedTableRequest, _ := datadog.Marshal(obj.GuidedTableRequest)
			if string(jsonGuidedTableRequest) == "{}" { // empty struct
				obj.GuidedTableRequest = nil
			} else {
				match++
			}
		} else {
			obj.GuidedTableRequest = nil
		}
	} else {
		obj.GuidedTableRequest = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TableWidgetRequest = nil
		obj.GuidedTableRequest = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TableWidgetDefinitionRequestsItem) MarshalJSON() ([]byte, error) {
	if obj.TableWidgetRequest != nil {
		return datadog.Marshal(&obj.TableWidgetRequest)
	}

	if obj.GuidedTableRequest != nil {
		return datadog.Marshal(&obj.GuidedTableRequest)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TableWidgetDefinitionRequestsItem) GetActualInstance() interface{} {
	if obj.TableWidgetRequest != nil {
		return obj.TableWidgetRequest
	}

	if obj.GuidedTableRequest != nil {
		return obj.GuidedTableRequest
	}

	// all schemas are nil
	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchSavedFilterConflictResponse - A stale update, duplicate name, or other conflict, or a resource listing consumer flags requiring approval.
type PatchSavedFilterConflictResponse struct {
	SavedFilterJSONAPIErrorResponse     *SavedFilterJSONAPIErrorResponse
	SavedFilterApprovalConflictResponse *SavedFilterApprovalConflictResponse

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SavedFilterJSONAPIErrorResponseAsPatchSavedFilterConflictResponse is a convenience function that returns SavedFilterJSONAPIErrorResponse wrapped in PatchSavedFilterConflictResponse.
func SavedFilterJSONAPIErrorResponseAsPatchSavedFilterConflictResponse(v *SavedFilterJSONAPIErrorResponse) PatchSavedFilterConflictResponse {
	return PatchSavedFilterConflictResponse{SavedFilterJSONAPIErrorResponse: v}
}

// SavedFilterApprovalConflictResponseAsPatchSavedFilterConflictResponse is a convenience function that returns SavedFilterApprovalConflictResponse wrapped in PatchSavedFilterConflictResponse.
func SavedFilterApprovalConflictResponseAsPatchSavedFilterConflictResponse(v *SavedFilterApprovalConflictResponse) PatchSavedFilterConflictResponse {
	return PatchSavedFilterConflictResponse{SavedFilterApprovalConflictResponse: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *PatchSavedFilterConflictResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SavedFilterJSONAPIErrorResponse
	err = datadog.Unmarshal(data, &obj.SavedFilterJSONAPIErrorResponse)
	if err == nil {
		if obj.SavedFilterJSONAPIErrorResponse != nil && obj.SavedFilterJSONAPIErrorResponse.UnparsedObject == nil {
			jsonSavedFilterJSONAPIErrorResponse, _ := datadog.Marshal(obj.SavedFilterJSONAPIErrorResponse)
			if string(jsonSavedFilterJSONAPIErrorResponse) == "{}" { // empty struct
				obj.SavedFilterJSONAPIErrorResponse = nil
			} else {
				match++
			}
		} else {
			obj.SavedFilterJSONAPIErrorResponse = nil
		}
	} else {
		obj.SavedFilterJSONAPIErrorResponse = nil
	}

	// try to unmarshal data into SavedFilterApprovalConflictResponse
	err = datadog.Unmarshal(data, &obj.SavedFilterApprovalConflictResponse)
	if err == nil {
		if obj.SavedFilterApprovalConflictResponse != nil && obj.SavedFilterApprovalConflictResponse.UnparsedObject == nil {
			jsonSavedFilterApprovalConflictResponse, _ := datadog.Marshal(obj.SavedFilterApprovalConflictResponse)
			if string(jsonSavedFilterApprovalConflictResponse) == "{}" { // empty struct
				obj.SavedFilterApprovalConflictResponse = nil
			} else {
				match++
			}
		} else {
			obj.SavedFilterApprovalConflictResponse = nil
		}
	} else {
		obj.SavedFilterApprovalConflictResponse = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SavedFilterJSONAPIErrorResponse = nil
		obj.SavedFilterApprovalConflictResponse = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj PatchSavedFilterConflictResponse) MarshalJSON() ([]byte, error) {
	if obj.SavedFilterJSONAPIErrorResponse != nil {
		return datadog.Marshal(&obj.SavedFilterJSONAPIErrorResponse)
	}

	if obj.SavedFilterApprovalConflictResponse != nil {
		return datadog.Marshal(&obj.SavedFilterApprovalConflictResponse)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *PatchSavedFilterConflictResponse) GetActualInstance() interface{} {
	if obj.SavedFilterJSONAPIErrorResponse != nil {
		return obj.SavedFilterJSONAPIErrorResponse
	}

	if obj.SavedFilterApprovalConflictResponse != nil {
		return obj.SavedFilterApprovalConflictResponse
	}

	// all schemas are nil
	return nil
}

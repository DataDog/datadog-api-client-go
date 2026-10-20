// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchSavedFilterForbiddenResponse - An authentication or permission error, or a resource listing consumer flags that the caller cannot edit.
type PatchSavedFilterForbiddenResponse struct {
	SavedFilterAPIErrorResponse         *SavedFilterAPIErrorResponse
	SavedFilterJSONAPIErrorResponse     *SavedFilterJSONAPIErrorResponse
	SavedFilterPermissionDeniedResponse *SavedFilterPermissionDeniedResponse

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SavedFilterAPIErrorResponseAsPatchSavedFilterForbiddenResponse is a convenience function that returns SavedFilterAPIErrorResponse wrapped in PatchSavedFilterForbiddenResponse.
func SavedFilterAPIErrorResponseAsPatchSavedFilterForbiddenResponse(v *SavedFilterAPIErrorResponse) PatchSavedFilterForbiddenResponse {
	return PatchSavedFilterForbiddenResponse{SavedFilterAPIErrorResponse: v}
}

// SavedFilterJSONAPIErrorResponseAsPatchSavedFilterForbiddenResponse is a convenience function that returns SavedFilterJSONAPIErrorResponse wrapped in PatchSavedFilterForbiddenResponse.
func SavedFilterJSONAPIErrorResponseAsPatchSavedFilterForbiddenResponse(v *SavedFilterJSONAPIErrorResponse) PatchSavedFilterForbiddenResponse {
	return PatchSavedFilterForbiddenResponse{SavedFilterJSONAPIErrorResponse: v}
}

// SavedFilterPermissionDeniedResponseAsPatchSavedFilterForbiddenResponse is a convenience function that returns SavedFilterPermissionDeniedResponse wrapped in PatchSavedFilterForbiddenResponse.
func SavedFilterPermissionDeniedResponseAsPatchSavedFilterForbiddenResponse(v *SavedFilterPermissionDeniedResponse) PatchSavedFilterForbiddenResponse {
	return PatchSavedFilterForbiddenResponse{SavedFilterPermissionDeniedResponse: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *PatchSavedFilterForbiddenResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SavedFilterAPIErrorResponse
	err = datadog.Unmarshal(data, &obj.SavedFilterAPIErrorResponse)
	if err == nil {
		if obj.SavedFilterAPIErrorResponse != nil && obj.SavedFilterAPIErrorResponse.UnparsedObject == nil {
			jsonSavedFilterAPIErrorResponse, _ := datadog.Marshal(obj.SavedFilterAPIErrorResponse)
			if string(jsonSavedFilterAPIErrorResponse) == "{}" { // empty struct
				obj.SavedFilterAPIErrorResponse = nil
			} else {
				match++
			}
		} else {
			obj.SavedFilterAPIErrorResponse = nil
		}
	} else {
		obj.SavedFilterAPIErrorResponse = nil
	}

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

	// try to unmarshal data into SavedFilterPermissionDeniedResponse
	err = datadog.Unmarshal(data, &obj.SavedFilterPermissionDeniedResponse)
	if err == nil {
		if obj.SavedFilterPermissionDeniedResponse != nil && obj.SavedFilterPermissionDeniedResponse.UnparsedObject == nil {
			jsonSavedFilterPermissionDeniedResponse, _ := datadog.Marshal(obj.SavedFilterPermissionDeniedResponse)
			if string(jsonSavedFilterPermissionDeniedResponse) == "{}" { // empty struct
				obj.SavedFilterPermissionDeniedResponse = nil
			} else {
				match++
			}
		} else {
			obj.SavedFilterPermissionDeniedResponse = nil
		}
	} else {
		obj.SavedFilterPermissionDeniedResponse = nil
	}

	if match > 1 {
		// more than one variant matched, so tell them apart by the items of the arrays they hold
		if obj.SavedFilterJSONAPIErrorResponse != nil {
			allErrorsUnparsed := len(obj.SavedFilterJSONAPIErrorResponse.Errors) > 0
			for _, item := range obj.SavedFilterJSONAPIErrorResponse.Errors {
				if item.UnparsedObject == nil {
					allErrorsUnparsed = false
					break
				}
			}
			if allErrorsUnparsed {
				obj.SavedFilterJSONAPIErrorResponse = nil
				match--
			}
		}
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SavedFilterAPIErrorResponse = nil
		obj.SavedFilterJSONAPIErrorResponse = nil
		obj.SavedFilterPermissionDeniedResponse = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj PatchSavedFilterForbiddenResponse) MarshalJSON() ([]byte, error) {
	if obj.SavedFilterAPIErrorResponse != nil {
		return datadog.Marshal(&obj.SavedFilterAPIErrorResponse)
	}

	if obj.SavedFilterJSONAPIErrorResponse != nil {
		return datadog.Marshal(&obj.SavedFilterJSONAPIErrorResponse)
	}

	if obj.SavedFilterPermissionDeniedResponse != nil {
		return datadog.Marshal(&obj.SavedFilterPermissionDeniedResponse)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *PatchSavedFilterForbiddenResponse) GetActualInstance() interface{} {
	if obj.SavedFilterAPIErrorResponse != nil {
		return obj.SavedFilterAPIErrorResponse
	}

	if obj.SavedFilterJSONAPIErrorResponse != nil {
		return obj.SavedFilterJSONAPIErrorResponse
	}

	if obj.SavedFilterPermissionDeniedResponse != nil {
		return obj.SavedFilterPermissionDeniedResponse
	}

	// all schemas are nil
	return nil
}

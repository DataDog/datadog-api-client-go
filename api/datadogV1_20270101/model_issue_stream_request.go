// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueStreamRequest - Test-only generated schema for the dashboard versioning proof of concept.
type IssueStreamRequest struct {
	LegacyIssueStreamRequest_20270101  *LegacyIssueStreamRequest
	UnifiedIssueStreamRequest_20270101 *UnifiedIssueStreamRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// LegacyIssueStreamRequest_20270101AsIssueStreamRequest is a convenience function that returns LegacyIssueStreamRequest wrapped in IssueStreamRequest.
func LegacyIssueStreamRequest_20270101AsIssueStreamRequest(v *LegacyIssueStreamRequest) IssueStreamRequest {
	return IssueStreamRequest{LegacyIssueStreamRequest_20270101: v}
}

// UnifiedIssueStreamRequest_20270101AsIssueStreamRequest is a convenience function that returns UnifiedIssueStreamRequest wrapped in IssueStreamRequest.
func UnifiedIssueStreamRequest_20270101AsIssueStreamRequest(v *UnifiedIssueStreamRequest) IssueStreamRequest {
	return IssueStreamRequest{UnifiedIssueStreamRequest_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *IssueStreamRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LegacyIssueStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.LegacyIssueStreamRequest_20270101)
	if err == nil {
		if obj.LegacyIssueStreamRequest_20270101 != nil && obj.LegacyIssueStreamRequest_20270101.UnparsedObject == nil {
			jsonLegacyIssueStreamRequest_20270101, _ := datadog.Marshal(obj.LegacyIssueStreamRequest_20270101)
			if string(jsonLegacyIssueStreamRequest_20270101) == "{}" { // empty struct
				obj.LegacyIssueStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.LegacyIssueStreamRequest_20270101 = nil
		}
	} else {
		obj.LegacyIssueStreamRequest_20270101 = nil
	}

	// try to unmarshal data into UnifiedIssueStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.UnifiedIssueStreamRequest_20270101)
	if err == nil {
		if obj.UnifiedIssueStreamRequest_20270101 != nil && obj.UnifiedIssueStreamRequest_20270101.UnparsedObject == nil {
			jsonUnifiedIssueStreamRequest_20270101, _ := datadog.Marshal(obj.UnifiedIssueStreamRequest_20270101)
			if string(jsonUnifiedIssueStreamRequest_20270101) == "{}" { // empty struct
				obj.UnifiedIssueStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.UnifiedIssueStreamRequest_20270101 = nil
		}
	} else {
		obj.UnifiedIssueStreamRequest_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.LegacyIssueStreamRequest_20270101 = nil
		obj.UnifiedIssueStreamRequest_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj IssueStreamRequest) MarshalJSON() ([]byte, error) {
	if obj.LegacyIssueStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.LegacyIssueStreamRequest_20270101)
	}

	if obj.UnifiedIssueStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.UnifiedIssueStreamRequest_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *IssueStreamRequest) GetActualInstance() interface{} {
	if obj.LegacyIssueStreamRequest_20270101 != nil {
		return obj.LegacyIssueStreamRequest_20270101
	}

	if obj.UnifiedIssueStreamRequest_20270101 != nil {
		return obj.UnifiedIssueStreamRequest_20270101
	}

	// all schemas are nil
	return nil
}

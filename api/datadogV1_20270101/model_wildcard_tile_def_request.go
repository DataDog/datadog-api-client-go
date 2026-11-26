// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WildcardTileDefRequest - Test-only generated schema for the dashboard versioning proof of concept.
type WildcardTileDefRequest struct {
	TabularDataRequest_20270101                *TabularDataRequest
	DeprecatedWildcardTileDefRequests_20270101 *DeprecatedWildcardTileDefRequests

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TabularDataRequest_20270101AsWildcardTileDefRequest is a convenience function that returns TabularDataRequest wrapped in WildcardTileDefRequest.
func TabularDataRequest_20270101AsWildcardTileDefRequest(v *TabularDataRequest) WildcardTileDefRequest {
	return WildcardTileDefRequest{TabularDataRequest_20270101: v}
}

// DeprecatedWildcardTileDefRequests_20270101AsWildcardTileDefRequest is a convenience function that returns DeprecatedWildcardTileDefRequests wrapped in WildcardTileDefRequest.
func DeprecatedWildcardTileDefRequests_20270101AsWildcardTileDefRequest(v *DeprecatedWildcardTileDefRequests) WildcardTileDefRequest {
	return WildcardTileDefRequest{DeprecatedWildcardTileDefRequests_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *WildcardTileDefRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TabularDataRequest_20270101
	err = datadog.Unmarshal(data, &obj.TabularDataRequest_20270101)
	if err == nil {
		if obj.TabularDataRequest_20270101 != nil && obj.TabularDataRequest_20270101.UnparsedObject == nil {
			jsonTabularDataRequest_20270101, _ := datadog.Marshal(obj.TabularDataRequest_20270101)
			if string(jsonTabularDataRequest_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.TabularDataRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TabularDataRequest_20270101 = nil
		}
	} else {
		obj.TabularDataRequest_20270101 = nil
	}

	// try to unmarshal data into DeprecatedWildcardTileDefRequests_20270101
	err = datadog.Unmarshal(data, &obj.DeprecatedWildcardTileDefRequests_20270101)
	if err == nil {
		if obj.DeprecatedWildcardTileDefRequests_20270101 != nil && obj.DeprecatedWildcardTileDefRequests_20270101.UnparsedObject == nil {
			jsonDeprecatedWildcardTileDefRequests_20270101, _ := datadog.Marshal(obj.DeprecatedWildcardTileDefRequests_20270101)
			if string(jsonDeprecatedWildcardTileDefRequests_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.DeprecatedWildcardTileDefRequests_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.DeprecatedWildcardTileDefRequests_20270101 = nil
		}
	} else {
		obj.DeprecatedWildcardTileDefRequests_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TabularDataRequest_20270101 = nil
		obj.DeprecatedWildcardTileDefRequests_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj WildcardTileDefRequest) MarshalJSON() ([]byte, error) {
	if obj.TabularDataRequest_20270101 != nil {
		return datadog.Marshal(&obj.TabularDataRequest_20270101)
	}

	if obj.DeprecatedWildcardTileDefRequests_20270101 != nil {
		return datadog.Marshal(&obj.DeprecatedWildcardTileDefRequests_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *WildcardTileDefRequest) GetActualInstance() interface{} {
	if obj.TabularDataRequest_20270101 != nil {
		return obj.TabularDataRequest_20270101
	}

	if obj.DeprecatedWildcardTileDefRequests_20270101 != nil {
		return obj.DeprecatedWildcardTileDefRequests_20270101
	}

	// all schemas are nil
	return nil
}

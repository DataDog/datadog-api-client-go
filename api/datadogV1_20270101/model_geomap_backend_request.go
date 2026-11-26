// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapBackendRequest - Test-only generated schema for the dashboard versioning proof of concept.
type GeomapBackendRequest struct {
	GeomapBackendRegionsRequest_20270101 *GeomapBackendRegionsRequest
	GeomapBackendPointsRequest_20270101  *GeomapBackendPointsRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// GeomapBackendRegionsRequest_20270101AsGeomapBackendRequest is a convenience function that returns GeomapBackendRegionsRequest wrapped in GeomapBackendRequest.
func GeomapBackendRegionsRequest_20270101AsGeomapBackendRequest(v *GeomapBackendRegionsRequest) GeomapBackendRequest {
	return GeomapBackendRequest{GeomapBackendRegionsRequest_20270101: v}
}

// GeomapBackendPointsRequest_20270101AsGeomapBackendRequest is a convenience function that returns GeomapBackendPointsRequest wrapped in GeomapBackendRequest.
func GeomapBackendPointsRequest_20270101AsGeomapBackendRequest(v *GeomapBackendPointsRequest) GeomapBackendRequest {
	return GeomapBackendRequest{GeomapBackendPointsRequest_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GeomapBackendRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GeomapBackendRegionsRequest_20270101
	err = datadog.Unmarshal(data, &obj.GeomapBackendRegionsRequest_20270101)
	if err == nil {
		if obj.GeomapBackendRegionsRequest_20270101 != nil && obj.GeomapBackendRegionsRequest_20270101.UnparsedObject == nil {
			jsonGeomapBackendRegionsRequest_20270101, _ := datadog.Marshal(obj.GeomapBackendRegionsRequest_20270101)
			if string(jsonGeomapBackendRegionsRequest_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.GeomapBackendRegionsRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.GeomapBackendRegionsRequest_20270101 = nil
		}
	} else {
		obj.GeomapBackendRegionsRequest_20270101 = nil
	}

	// try to unmarshal data into GeomapBackendPointsRequest_20270101
	err = datadog.Unmarshal(data, &obj.GeomapBackendPointsRequest_20270101)
	if err == nil {
		if obj.GeomapBackendPointsRequest_20270101 != nil && obj.GeomapBackendPointsRequest_20270101.UnparsedObject == nil {
			jsonGeomapBackendPointsRequest_20270101, _ := datadog.Marshal(obj.GeomapBackendPointsRequest_20270101)
			if string(jsonGeomapBackendPointsRequest_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.GeomapBackendPointsRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.GeomapBackendPointsRequest_20270101 = nil
		}
	} else {
		obj.GeomapBackendPointsRequest_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.GeomapBackendRegionsRequest_20270101 = nil
		obj.GeomapBackendPointsRequest_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GeomapBackendRequest) MarshalJSON() ([]byte, error) {
	if obj.GeomapBackendRegionsRequest_20270101 != nil {
		return datadog.Marshal(&obj.GeomapBackendRegionsRequest_20270101)
	}

	if obj.GeomapBackendPointsRequest_20270101 != nil {
		return datadog.Marshal(&obj.GeomapBackendPointsRequest_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GeomapBackendRequest) GetActualInstance() interface{} {
	if obj.GeomapBackendRegionsRequest_20270101 != nil {
		return obj.GeomapBackendRegionsRequest_20270101
	}

	if obj.GeomapBackendPointsRequest_20270101 != nil {
		return obj.GeomapBackendPointsRequest_20270101
	}

	// all schemas are nil
	return nil
}

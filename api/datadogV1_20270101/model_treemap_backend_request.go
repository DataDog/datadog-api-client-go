// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreemapBackendRequest - Test-only generated schema for the dashboard versioning proof of concept.
type TreemapBackendRequest struct {
	TreemapBackendRequest20270101OneOf           *TreemapBackendRequest20270101OneOf
	TreemapBackendRequest20270101OneOf4463750144 *TreemapBackendRequest20270101OneOf4463750144
	TreemapBackendRequest20270101OneOf4463751424 *TreemapBackendRequest20270101OneOf4463751424
	TreemapBackendRequest20270101OneOf4463752384 *TreemapBackendRequest20270101OneOf4463752384
	TreemapBackendRequest20270101OneOf4463753824 *TreemapBackendRequest20270101OneOf4463753824

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TreemapBackendRequest20270101OneOfAsTreemapBackendRequest is a convenience function that returns TreemapBackendRequest20270101OneOf wrapped in TreemapBackendRequest.
func TreemapBackendRequest20270101OneOfAsTreemapBackendRequest(v *TreemapBackendRequest20270101OneOf) TreemapBackendRequest {
	return TreemapBackendRequest{TreemapBackendRequest20270101OneOf: v}
}

// TreemapBackendRequest20270101OneOf4463750144AsTreemapBackendRequest is a convenience function that returns TreemapBackendRequest20270101OneOf4463750144 wrapped in TreemapBackendRequest.
func TreemapBackendRequest20270101OneOf4463750144AsTreemapBackendRequest(v *TreemapBackendRequest20270101OneOf4463750144) TreemapBackendRequest {
	return TreemapBackendRequest{TreemapBackendRequest20270101OneOf4463750144: v}
}

// TreemapBackendRequest20270101OneOf4463751424AsTreemapBackendRequest is a convenience function that returns TreemapBackendRequest20270101OneOf4463751424 wrapped in TreemapBackendRequest.
func TreemapBackendRequest20270101OneOf4463751424AsTreemapBackendRequest(v *TreemapBackendRequest20270101OneOf4463751424) TreemapBackendRequest {
	return TreemapBackendRequest{TreemapBackendRequest20270101OneOf4463751424: v}
}

// TreemapBackendRequest20270101OneOf4463752384AsTreemapBackendRequest is a convenience function that returns TreemapBackendRequest20270101OneOf4463752384 wrapped in TreemapBackendRequest.
func TreemapBackendRequest20270101OneOf4463752384AsTreemapBackendRequest(v *TreemapBackendRequest20270101OneOf4463752384) TreemapBackendRequest {
	return TreemapBackendRequest{TreemapBackendRequest20270101OneOf4463752384: v}
}

// TreemapBackendRequest20270101OneOf4463753824AsTreemapBackendRequest is a convenience function that returns TreemapBackendRequest20270101OneOf4463753824 wrapped in TreemapBackendRequest.
func TreemapBackendRequest20270101OneOf4463753824AsTreemapBackendRequest(v *TreemapBackendRequest20270101OneOf4463753824) TreemapBackendRequest {
	return TreemapBackendRequest{TreemapBackendRequest20270101OneOf4463753824: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TreemapBackendRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TreemapBackendRequest20270101OneOf
	err = datadog.Unmarshal(data, &obj.TreemapBackendRequest20270101OneOf)
	if err == nil {
		if obj.TreemapBackendRequest20270101OneOf != nil && obj.TreemapBackendRequest20270101OneOf.UnparsedObject == nil {
			jsonTreemapBackendRequest20270101OneOf, _ := datadog.Marshal(obj.TreemapBackendRequest20270101OneOf)
			if string(jsonTreemapBackendRequest20270101OneOf) == "{}" { // empty struct
				obj.TreemapBackendRequest20270101OneOf = nil
			} else {
				match++
			}
		} else {
			obj.TreemapBackendRequest20270101OneOf = nil
		}
	} else {
		obj.TreemapBackendRequest20270101OneOf = nil
	}

	// try to unmarshal data into TreemapBackendRequest20270101OneOf4463750144
	err = datadog.Unmarshal(data, &obj.TreemapBackendRequest20270101OneOf4463750144)
	if err == nil {
		if obj.TreemapBackendRequest20270101OneOf4463750144 != nil && obj.TreemapBackendRequest20270101OneOf4463750144.UnparsedObject == nil {
			jsonTreemapBackendRequest20270101OneOf4463750144, _ := datadog.Marshal(obj.TreemapBackendRequest20270101OneOf4463750144)
			if string(jsonTreemapBackendRequest20270101OneOf4463750144) == "{}" { // empty struct
				obj.TreemapBackendRequest20270101OneOf4463750144 = nil
			} else {
				match++
			}
		} else {
			obj.TreemapBackendRequest20270101OneOf4463750144 = nil
		}
	} else {
		obj.TreemapBackendRequest20270101OneOf4463750144 = nil
	}

	// try to unmarshal data into TreemapBackendRequest20270101OneOf4463751424
	err = datadog.Unmarshal(data, &obj.TreemapBackendRequest20270101OneOf4463751424)
	if err == nil {
		if obj.TreemapBackendRequest20270101OneOf4463751424 != nil && obj.TreemapBackendRequest20270101OneOf4463751424.UnparsedObject == nil {
			jsonTreemapBackendRequest20270101OneOf4463751424, _ := datadog.Marshal(obj.TreemapBackendRequest20270101OneOf4463751424)
			if string(jsonTreemapBackendRequest20270101OneOf4463751424) == "{}" { // empty struct
				obj.TreemapBackendRequest20270101OneOf4463751424 = nil
			} else {
				match++
			}
		} else {
			obj.TreemapBackendRequest20270101OneOf4463751424 = nil
		}
	} else {
		obj.TreemapBackendRequest20270101OneOf4463751424 = nil
	}

	// try to unmarshal data into TreemapBackendRequest20270101OneOf4463752384
	err = datadog.Unmarshal(data, &obj.TreemapBackendRequest20270101OneOf4463752384)
	if err == nil {
		if obj.TreemapBackendRequest20270101OneOf4463752384 != nil && obj.TreemapBackendRequest20270101OneOf4463752384.UnparsedObject == nil {
			jsonTreemapBackendRequest20270101OneOf4463752384, _ := datadog.Marshal(obj.TreemapBackendRequest20270101OneOf4463752384)
			if string(jsonTreemapBackendRequest20270101OneOf4463752384) == "{}" { // empty struct
				obj.TreemapBackendRequest20270101OneOf4463752384 = nil
			} else {
				match++
			}
		} else {
			obj.TreemapBackendRequest20270101OneOf4463752384 = nil
		}
	} else {
		obj.TreemapBackendRequest20270101OneOf4463752384 = nil
	}

	// try to unmarshal data into TreemapBackendRequest20270101OneOf4463753824
	err = datadog.Unmarshal(data, &obj.TreemapBackendRequest20270101OneOf4463753824)
	if err == nil {
		if obj.TreemapBackendRequest20270101OneOf4463753824 != nil && obj.TreemapBackendRequest20270101OneOf4463753824.UnparsedObject == nil {
			jsonTreemapBackendRequest20270101OneOf4463753824, _ := datadog.Marshal(obj.TreemapBackendRequest20270101OneOf4463753824)
			if string(jsonTreemapBackendRequest20270101OneOf4463753824) == "{}" { // empty struct
				obj.TreemapBackendRequest20270101OneOf4463753824 = nil
			} else {
				match++
			}
		} else {
			obj.TreemapBackendRequest20270101OneOf4463753824 = nil
		}
	} else {
		obj.TreemapBackendRequest20270101OneOf4463753824 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TreemapBackendRequest20270101OneOf = nil
		obj.TreemapBackendRequest20270101OneOf4463750144 = nil
		obj.TreemapBackendRequest20270101OneOf4463751424 = nil
		obj.TreemapBackendRequest20270101OneOf4463752384 = nil
		obj.TreemapBackendRequest20270101OneOf4463753824 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TreemapBackendRequest) MarshalJSON() ([]byte, error) {
	if obj.TreemapBackendRequest20270101OneOf != nil {
		return datadog.Marshal(&obj.TreemapBackendRequest20270101OneOf)
	}

	if obj.TreemapBackendRequest20270101OneOf4463750144 != nil {
		return datadog.Marshal(&obj.TreemapBackendRequest20270101OneOf4463750144)
	}

	if obj.TreemapBackendRequest20270101OneOf4463751424 != nil {
		return datadog.Marshal(&obj.TreemapBackendRequest20270101OneOf4463751424)
	}

	if obj.TreemapBackendRequest20270101OneOf4463752384 != nil {
		return datadog.Marshal(&obj.TreemapBackendRequest20270101OneOf4463752384)
	}

	if obj.TreemapBackendRequest20270101OneOf4463753824 != nil {
		return datadog.Marshal(&obj.TreemapBackendRequest20270101OneOf4463753824)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TreemapBackendRequest) GetActualInstance() interface{} {
	if obj.TreemapBackendRequest20270101OneOf != nil {
		return obj.TreemapBackendRequest20270101OneOf
	}

	if obj.TreemapBackendRequest20270101OneOf4463750144 != nil {
		return obj.TreemapBackendRequest20270101OneOf4463750144
	}

	if obj.TreemapBackendRequest20270101OneOf4463751424 != nil {
		return obj.TreemapBackendRequest20270101OneOf4463751424
	}

	if obj.TreemapBackendRequest20270101OneOf4463752384 != nil {
		return obj.TreemapBackendRequest20270101OneOf4463752384
	}

	if obj.TreemapBackendRequest20270101OneOf4463753824 != nil {
		return obj.TreemapBackendRequest20270101OneOf4463753824
	}

	// all schemas are nil
	return nil
}

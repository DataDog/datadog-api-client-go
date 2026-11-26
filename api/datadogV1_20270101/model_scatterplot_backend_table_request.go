// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotBackendTableRequest - Test-only generated schema for the dashboard versioning proof of concept.
type ScatterplotBackendTableRequest struct {
	ScatterplotBackendTableRequest20270101OneOf           *ScatterplotBackendTableRequest20270101OneOf
	ScatterplotBackendTableRequest20270101OneOf4461521920 *ScatterplotBackendTableRequest20270101OneOf4461521920
	ScatterplotBackendTableRequest20270101OneOf4461523200 *ScatterplotBackendTableRequest20270101OneOf4461523200
	ScatterplotBackendTableRequest20270101OneOf4461524480 *ScatterplotBackendTableRequest20270101OneOf4461524480
	ScatterplotBackendTableRequest20270101OneOf4461525600 *ScatterplotBackendTableRequest20270101OneOf4461525600
	ScatterplotBackendTableRequest20270101OneOf4461756480 *ScatterplotBackendTableRequest20270101OneOf4461756480

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ScatterplotBackendTableRequest20270101OneOfAsScatterplotBackendTableRequest is a convenience function that returns ScatterplotBackendTableRequest20270101OneOf wrapped in ScatterplotBackendTableRequest.
func ScatterplotBackendTableRequest20270101OneOfAsScatterplotBackendTableRequest(v *ScatterplotBackendTableRequest20270101OneOf) ScatterplotBackendTableRequest {
	return ScatterplotBackendTableRequest{ScatterplotBackendTableRequest20270101OneOf: v}
}

// ScatterplotBackendTableRequest20270101OneOf4461521920AsScatterplotBackendTableRequest is a convenience function that returns ScatterplotBackendTableRequest20270101OneOf4461521920 wrapped in ScatterplotBackendTableRequest.
func ScatterplotBackendTableRequest20270101OneOf4461521920AsScatterplotBackendTableRequest(v *ScatterplotBackendTableRequest20270101OneOf4461521920) ScatterplotBackendTableRequest {
	return ScatterplotBackendTableRequest{ScatterplotBackendTableRequest20270101OneOf4461521920: v}
}

// ScatterplotBackendTableRequest20270101OneOf4461523200AsScatterplotBackendTableRequest is a convenience function that returns ScatterplotBackendTableRequest20270101OneOf4461523200 wrapped in ScatterplotBackendTableRequest.
func ScatterplotBackendTableRequest20270101OneOf4461523200AsScatterplotBackendTableRequest(v *ScatterplotBackendTableRequest20270101OneOf4461523200) ScatterplotBackendTableRequest {
	return ScatterplotBackendTableRequest{ScatterplotBackendTableRequest20270101OneOf4461523200: v}
}

// ScatterplotBackendTableRequest20270101OneOf4461524480AsScatterplotBackendTableRequest is a convenience function that returns ScatterplotBackendTableRequest20270101OneOf4461524480 wrapped in ScatterplotBackendTableRequest.
func ScatterplotBackendTableRequest20270101OneOf4461524480AsScatterplotBackendTableRequest(v *ScatterplotBackendTableRequest20270101OneOf4461524480) ScatterplotBackendTableRequest {
	return ScatterplotBackendTableRequest{ScatterplotBackendTableRequest20270101OneOf4461524480: v}
}

// ScatterplotBackendTableRequest20270101OneOf4461525600AsScatterplotBackendTableRequest is a convenience function that returns ScatterplotBackendTableRequest20270101OneOf4461525600 wrapped in ScatterplotBackendTableRequest.
func ScatterplotBackendTableRequest20270101OneOf4461525600AsScatterplotBackendTableRequest(v *ScatterplotBackendTableRequest20270101OneOf4461525600) ScatterplotBackendTableRequest {
	return ScatterplotBackendTableRequest{ScatterplotBackendTableRequest20270101OneOf4461525600: v}
}

// ScatterplotBackendTableRequest20270101OneOf4461756480AsScatterplotBackendTableRequest is a convenience function that returns ScatterplotBackendTableRequest20270101OneOf4461756480 wrapped in ScatterplotBackendTableRequest.
func ScatterplotBackendTableRequest20270101OneOf4461756480AsScatterplotBackendTableRequest(v *ScatterplotBackendTableRequest20270101OneOf4461756480) ScatterplotBackendTableRequest {
	return ScatterplotBackendTableRequest{ScatterplotBackendTableRequest20270101OneOf4461756480: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ScatterplotBackendTableRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ScatterplotBackendTableRequest20270101OneOf
	err = datadog.Unmarshal(data, &obj.ScatterplotBackendTableRequest20270101OneOf)
	if err == nil {
		if obj.ScatterplotBackendTableRequest20270101OneOf != nil && obj.ScatterplotBackendTableRequest20270101OneOf.UnparsedObject == nil {
			jsonScatterplotBackendTableRequest20270101OneOf, _ := datadog.Marshal(obj.ScatterplotBackendTableRequest20270101OneOf)
			if string(jsonScatterplotBackendTableRequest20270101OneOf) == "{}" { // empty struct
				obj.ScatterplotBackendTableRequest20270101OneOf = nil
			} else {
				match++
			}
		} else {
			obj.ScatterplotBackendTableRequest20270101OneOf = nil
		}
	} else {
		obj.ScatterplotBackendTableRequest20270101OneOf = nil
	}

	// try to unmarshal data into ScatterplotBackendTableRequest20270101OneOf4461521920
	err = datadog.Unmarshal(data, &obj.ScatterplotBackendTableRequest20270101OneOf4461521920)
	if err == nil {
		if obj.ScatterplotBackendTableRequest20270101OneOf4461521920 != nil && obj.ScatterplotBackendTableRequest20270101OneOf4461521920.UnparsedObject == nil {
			jsonScatterplotBackendTableRequest20270101OneOf4461521920, _ := datadog.Marshal(obj.ScatterplotBackendTableRequest20270101OneOf4461521920)
			if string(jsonScatterplotBackendTableRequest20270101OneOf4461521920) == "{}" { // empty struct
				obj.ScatterplotBackendTableRequest20270101OneOf4461521920 = nil
			} else {
				match++
			}
		} else {
			obj.ScatterplotBackendTableRequest20270101OneOf4461521920 = nil
		}
	} else {
		obj.ScatterplotBackendTableRequest20270101OneOf4461521920 = nil
	}

	// try to unmarshal data into ScatterplotBackendTableRequest20270101OneOf4461523200
	err = datadog.Unmarshal(data, &obj.ScatterplotBackendTableRequest20270101OneOf4461523200)
	if err == nil {
		if obj.ScatterplotBackendTableRequest20270101OneOf4461523200 != nil && obj.ScatterplotBackendTableRequest20270101OneOf4461523200.UnparsedObject == nil {
			jsonScatterplotBackendTableRequest20270101OneOf4461523200, _ := datadog.Marshal(obj.ScatterplotBackendTableRequest20270101OneOf4461523200)
			if string(jsonScatterplotBackendTableRequest20270101OneOf4461523200) == "{}" { // empty struct
				obj.ScatterplotBackendTableRequest20270101OneOf4461523200 = nil
			} else {
				match++
			}
		} else {
			obj.ScatterplotBackendTableRequest20270101OneOf4461523200 = nil
		}
	} else {
		obj.ScatterplotBackendTableRequest20270101OneOf4461523200 = nil
	}

	// try to unmarshal data into ScatterplotBackendTableRequest20270101OneOf4461524480
	err = datadog.Unmarshal(data, &obj.ScatterplotBackendTableRequest20270101OneOf4461524480)
	if err == nil {
		if obj.ScatterplotBackendTableRequest20270101OneOf4461524480 != nil && obj.ScatterplotBackendTableRequest20270101OneOf4461524480.UnparsedObject == nil {
			jsonScatterplotBackendTableRequest20270101OneOf4461524480, _ := datadog.Marshal(obj.ScatterplotBackendTableRequest20270101OneOf4461524480)
			if string(jsonScatterplotBackendTableRequest20270101OneOf4461524480) == "{}" { // empty struct
				obj.ScatterplotBackendTableRequest20270101OneOf4461524480 = nil
			} else {
				match++
			}
		} else {
			obj.ScatterplotBackendTableRequest20270101OneOf4461524480 = nil
		}
	} else {
		obj.ScatterplotBackendTableRequest20270101OneOf4461524480 = nil
	}

	// try to unmarshal data into ScatterplotBackendTableRequest20270101OneOf4461525600
	err = datadog.Unmarshal(data, &obj.ScatterplotBackendTableRequest20270101OneOf4461525600)
	if err == nil {
		if obj.ScatterplotBackendTableRequest20270101OneOf4461525600 != nil && obj.ScatterplotBackendTableRequest20270101OneOf4461525600.UnparsedObject == nil {
			jsonScatterplotBackendTableRequest20270101OneOf4461525600, _ := datadog.Marshal(obj.ScatterplotBackendTableRequest20270101OneOf4461525600)
			if string(jsonScatterplotBackendTableRequest20270101OneOf4461525600) == "{}" { // empty struct
				obj.ScatterplotBackendTableRequest20270101OneOf4461525600 = nil
			} else {
				match++
			}
		} else {
			obj.ScatterplotBackendTableRequest20270101OneOf4461525600 = nil
		}
	} else {
		obj.ScatterplotBackendTableRequest20270101OneOf4461525600 = nil
	}

	// try to unmarshal data into ScatterplotBackendTableRequest20270101OneOf4461756480
	err = datadog.Unmarshal(data, &obj.ScatterplotBackendTableRequest20270101OneOf4461756480)
	if err == nil {
		if obj.ScatterplotBackendTableRequest20270101OneOf4461756480 != nil && obj.ScatterplotBackendTableRequest20270101OneOf4461756480.UnparsedObject == nil {
			jsonScatterplotBackendTableRequest20270101OneOf4461756480, _ := datadog.Marshal(obj.ScatterplotBackendTableRequest20270101OneOf4461756480)
			if string(jsonScatterplotBackendTableRequest20270101OneOf4461756480) == "{}" { // empty struct
				obj.ScatterplotBackendTableRequest20270101OneOf4461756480 = nil
			} else {
				match++
			}
		} else {
			obj.ScatterplotBackendTableRequest20270101OneOf4461756480 = nil
		}
	} else {
		obj.ScatterplotBackendTableRequest20270101OneOf4461756480 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ScatterplotBackendTableRequest20270101OneOf = nil
		obj.ScatterplotBackendTableRequest20270101OneOf4461521920 = nil
		obj.ScatterplotBackendTableRequest20270101OneOf4461523200 = nil
		obj.ScatterplotBackendTableRequest20270101OneOf4461524480 = nil
		obj.ScatterplotBackendTableRequest20270101OneOf4461525600 = nil
		obj.ScatterplotBackendTableRequest20270101OneOf4461756480 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ScatterplotBackendTableRequest) MarshalJSON() ([]byte, error) {
	if obj.ScatterplotBackendTableRequest20270101OneOf != nil {
		return datadog.Marshal(&obj.ScatterplotBackendTableRequest20270101OneOf)
	}

	if obj.ScatterplotBackendTableRequest20270101OneOf4461521920 != nil {
		return datadog.Marshal(&obj.ScatterplotBackendTableRequest20270101OneOf4461521920)
	}

	if obj.ScatterplotBackendTableRequest20270101OneOf4461523200 != nil {
		return datadog.Marshal(&obj.ScatterplotBackendTableRequest20270101OneOf4461523200)
	}

	if obj.ScatterplotBackendTableRequest20270101OneOf4461524480 != nil {
		return datadog.Marshal(&obj.ScatterplotBackendTableRequest20270101OneOf4461524480)
	}

	if obj.ScatterplotBackendTableRequest20270101OneOf4461525600 != nil {
		return datadog.Marshal(&obj.ScatterplotBackendTableRequest20270101OneOf4461525600)
	}

	if obj.ScatterplotBackendTableRequest20270101OneOf4461756480 != nil {
		return datadog.Marshal(&obj.ScatterplotBackendTableRequest20270101OneOf4461756480)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ScatterplotBackendTableRequest) GetActualInstance() interface{} {
	if obj.ScatterplotBackendTableRequest20270101OneOf != nil {
		return obj.ScatterplotBackendTableRequest20270101OneOf
	}

	if obj.ScatterplotBackendTableRequest20270101OneOf4461521920 != nil {
		return obj.ScatterplotBackendTableRequest20270101OneOf4461521920
	}

	if obj.ScatterplotBackendTableRequest20270101OneOf4461523200 != nil {
		return obj.ScatterplotBackendTableRequest20270101OneOf4461523200
	}

	if obj.ScatterplotBackendTableRequest20270101OneOf4461524480 != nil {
		return obj.ScatterplotBackendTableRequest20270101OneOf4461524480
	}

	if obj.ScatterplotBackendTableRequest20270101OneOf4461525600 != nil {
		return obj.ScatterplotBackendTableRequest20270101OneOf4461525600
	}

	if obj.ScatterplotBackendTableRequest20270101OneOf4461756480 != nil {
		return obj.ScatterplotBackendTableRequest20270101OneOf4461756480
	}

	// all schemas are nil
	return nil
}

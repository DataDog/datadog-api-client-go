// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BarChartBackendRequest - Test-only generated schema for the dashboard versioning proof of concept.
type BarChartBackendRequest struct {
	BarChartBackendRequest20270101OneOf           *BarChartBackendRequest20270101OneOf
	BarChartBackendRequest20270101OneOf4454434112 *BarChartBackendRequest20270101OneOf4454434112
	BarChartBackendRequest20270101OneOf4454435552 *BarChartBackendRequest20270101OneOf4454435552
	BarChartBackendRequest20270101OneOf4454436672 *BarChartBackendRequest20270101OneOf4454436672
	BarChartBackendRequest20270101OneOf4454438272 *BarChartBackendRequest20270101OneOf4454438272

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// BarChartBackendRequest20270101OneOfAsBarChartBackendRequest is a convenience function that returns BarChartBackendRequest20270101OneOf wrapped in BarChartBackendRequest.
func BarChartBackendRequest20270101OneOfAsBarChartBackendRequest(v *BarChartBackendRequest20270101OneOf) BarChartBackendRequest {
	return BarChartBackendRequest{BarChartBackendRequest20270101OneOf: v}
}

// BarChartBackendRequest20270101OneOf4454434112AsBarChartBackendRequest is a convenience function that returns BarChartBackendRequest20270101OneOf4454434112 wrapped in BarChartBackendRequest.
func BarChartBackendRequest20270101OneOf4454434112AsBarChartBackendRequest(v *BarChartBackendRequest20270101OneOf4454434112) BarChartBackendRequest {
	return BarChartBackendRequest{BarChartBackendRequest20270101OneOf4454434112: v}
}

// BarChartBackendRequest20270101OneOf4454435552AsBarChartBackendRequest is a convenience function that returns BarChartBackendRequest20270101OneOf4454435552 wrapped in BarChartBackendRequest.
func BarChartBackendRequest20270101OneOf4454435552AsBarChartBackendRequest(v *BarChartBackendRequest20270101OneOf4454435552) BarChartBackendRequest {
	return BarChartBackendRequest{BarChartBackendRequest20270101OneOf4454435552: v}
}

// BarChartBackendRequest20270101OneOf4454436672AsBarChartBackendRequest is a convenience function that returns BarChartBackendRequest20270101OneOf4454436672 wrapped in BarChartBackendRequest.
func BarChartBackendRequest20270101OneOf4454436672AsBarChartBackendRequest(v *BarChartBackendRequest20270101OneOf4454436672) BarChartBackendRequest {
	return BarChartBackendRequest{BarChartBackendRequest20270101OneOf4454436672: v}
}

// BarChartBackendRequest20270101OneOf4454438272AsBarChartBackendRequest is a convenience function that returns BarChartBackendRequest20270101OneOf4454438272 wrapped in BarChartBackendRequest.
func BarChartBackendRequest20270101OneOf4454438272AsBarChartBackendRequest(v *BarChartBackendRequest20270101OneOf4454438272) BarChartBackendRequest {
	return BarChartBackendRequest{BarChartBackendRequest20270101OneOf4454438272: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *BarChartBackendRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BarChartBackendRequest20270101OneOf
	err = datadog.Unmarshal(data, &obj.BarChartBackendRequest20270101OneOf)
	if err == nil {
		if obj.BarChartBackendRequest20270101OneOf != nil && obj.BarChartBackendRequest20270101OneOf.UnparsedObject == nil {
			jsonBarChartBackendRequest20270101OneOf, _ := datadog.Marshal(obj.BarChartBackendRequest20270101OneOf)
			if string(jsonBarChartBackendRequest20270101OneOf) == "{}" { // empty struct
				obj.BarChartBackendRequest20270101OneOf = nil
			} else {
				match++
			}
		} else {
			obj.BarChartBackendRequest20270101OneOf = nil
		}
	} else {
		obj.BarChartBackendRequest20270101OneOf = nil
	}

	// try to unmarshal data into BarChartBackendRequest20270101OneOf4454434112
	err = datadog.Unmarshal(data, &obj.BarChartBackendRequest20270101OneOf4454434112)
	if err == nil {
		if obj.BarChartBackendRequest20270101OneOf4454434112 != nil && obj.BarChartBackendRequest20270101OneOf4454434112.UnparsedObject == nil {
			jsonBarChartBackendRequest20270101OneOf4454434112, _ := datadog.Marshal(obj.BarChartBackendRequest20270101OneOf4454434112)
			if string(jsonBarChartBackendRequest20270101OneOf4454434112) == "{}" { // empty struct
				obj.BarChartBackendRequest20270101OneOf4454434112 = nil
			} else {
				match++
			}
		} else {
			obj.BarChartBackendRequest20270101OneOf4454434112 = nil
		}
	} else {
		obj.BarChartBackendRequest20270101OneOf4454434112 = nil
	}

	// try to unmarshal data into BarChartBackendRequest20270101OneOf4454435552
	err = datadog.Unmarshal(data, &obj.BarChartBackendRequest20270101OneOf4454435552)
	if err == nil {
		if obj.BarChartBackendRequest20270101OneOf4454435552 != nil && obj.BarChartBackendRequest20270101OneOf4454435552.UnparsedObject == nil {
			jsonBarChartBackendRequest20270101OneOf4454435552, _ := datadog.Marshal(obj.BarChartBackendRequest20270101OneOf4454435552)
			if string(jsonBarChartBackendRequest20270101OneOf4454435552) == "{}" { // empty struct
				obj.BarChartBackendRequest20270101OneOf4454435552 = nil
			} else {
				match++
			}
		} else {
			obj.BarChartBackendRequest20270101OneOf4454435552 = nil
		}
	} else {
		obj.BarChartBackendRequest20270101OneOf4454435552 = nil
	}

	// try to unmarshal data into BarChartBackendRequest20270101OneOf4454436672
	err = datadog.Unmarshal(data, &obj.BarChartBackendRequest20270101OneOf4454436672)
	if err == nil {
		if obj.BarChartBackendRequest20270101OneOf4454436672 != nil && obj.BarChartBackendRequest20270101OneOf4454436672.UnparsedObject == nil {
			jsonBarChartBackendRequest20270101OneOf4454436672, _ := datadog.Marshal(obj.BarChartBackendRequest20270101OneOf4454436672)
			if string(jsonBarChartBackendRequest20270101OneOf4454436672) == "{}" { // empty struct
				obj.BarChartBackendRequest20270101OneOf4454436672 = nil
			} else {
				match++
			}
		} else {
			obj.BarChartBackendRequest20270101OneOf4454436672 = nil
		}
	} else {
		obj.BarChartBackendRequest20270101OneOf4454436672 = nil
	}

	// try to unmarshal data into BarChartBackendRequest20270101OneOf4454438272
	err = datadog.Unmarshal(data, &obj.BarChartBackendRequest20270101OneOf4454438272)
	if err == nil {
		if obj.BarChartBackendRequest20270101OneOf4454438272 != nil && obj.BarChartBackendRequest20270101OneOf4454438272.UnparsedObject == nil {
			jsonBarChartBackendRequest20270101OneOf4454438272, _ := datadog.Marshal(obj.BarChartBackendRequest20270101OneOf4454438272)
			if string(jsonBarChartBackendRequest20270101OneOf4454438272) == "{}" { // empty struct
				obj.BarChartBackendRequest20270101OneOf4454438272 = nil
			} else {
				match++
			}
		} else {
			obj.BarChartBackendRequest20270101OneOf4454438272 = nil
		}
	} else {
		obj.BarChartBackendRequest20270101OneOf4454438272 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.BarChartBackendRequest20270101OneOf = nil
		obj.BarChartBackendRequest20270101OneOf4454434112 = nil
		obj.BarChartBackendRequest20270101OneOf4454435552 = nil
		obj.BarChartBackendRequest20270101OneOf4454436672 = nil
		obj.BarChartBackendRequest20270101OneOf4454438272 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj BarChartBackendRequest) MarshalJSON() ([]byte, error) {
	if obj.BarChartBackendRequest20270101OneOf != nil {
		return datadog.Marshal(&obj.BarChartBackendRequest20270101OneOf)
	}

	if obj.BarChartBackendRequest20270101OneOf4454434112 != nil {
		return datadog.Marshal(&obj.BarChartBackendRequest20270101OneOf4454434112)
	}

	if obj.BarChartBackendRequest20270101OneOf4454435552 != nil {
		return datadog.Marshal(&obj.BarChartBackendRequest20270101OneOf4454435552)
	}

	if obj.BarChartBackendRequest20270101OneOf4454436672 != nil {
		return datadog.Marshal(&obj.BarChartBackendRequest20270101OneOf4454436672)
	}

	if obj.BarChartBackendRequest20270101OneOf4454438272 != nil {
		return datadog.Marshal(&obj.BarChartBackendRequest20270101OneOf4454438272)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *BarChartBackendRequest) GetActualInstance() interface{} {
	if obj.BarChartBackendRequest20270101OneOf != nil {
		return obj.BarChartBackendRequest20270101OneOf
	}

	if obj.BarChartBackendRequest20270101OneOf4454434112 != nil {
		return obj.BarChartBackendRequest20270101OneOf4454434112
	}

	if obj.BarChartBackendRequest20270101OneOf4454435552 != nil {
		return obj.BarChartBackendRequest20270101OneOf4454435552
	}

	if obj.BarChartBackendRequest20270101OneOf4454436672 != nil {
		return obj.BarChartBackendRequest20270101OneOf4454436672
	}

	if obj.BarChartBackendRequest20270101OneOf4454438272 != nil {
		return obj.BarChartBackendRequest20270101OneOf4454438272
	}

	// all schemas are nil
	return nil
}

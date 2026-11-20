// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatmapBackendRequest - Test-only generated schema for the dashboard versioning proof of concept.
type HeatmapBackendRequest struct {
	HeatmapBackendRequest20270101OneOf           *HeatmapBackendRequest20270101OneOf
	HeatmapBackendRequest20270101OneOf4458138496 *HeatmapBackendRequest20270101OneOf4458138496
	HeatmapBackendRequest20270101OneOf4458139936 *HeatmapBackendRequest20270101OneOf4458139936
	HeatmapBackendRequest20270101OneOf4458140576 *HeatmapBackendRequest20270101OneOf4458140576

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// HeatmapBackendRequest20270101OneOfAsHeatmapBackendRequest is a convenience function that returns HeatmapBackendRequest20270101OneOf wrapped in HeatmapBackendRequest.
func HeatmapBackendRequest20270101OneOfAsHeatmapBackendRequest(v *HeatmapBackendRequest20270101OneOf) HeatmapBackendRequest {
	return HeatmapBackendRequest{HeatmapBackendRequest20270101OneOf: v}
}

// HeatmapBackendRequest20270101OneOf4458138496AsHeatmapBackendRequest is a convenience function that returns HeatmapBackendRequest20270101OneOf4458138496 wrapped in HeatmapBackendRequest.
func HeatmapBackendRequest20270101OneOf4458138496AsHeatmapBackendRequest(v *HeatmapBackendRequest20270101OneOf4458138496) HeatmapBackendRequest {
	return HeatmapBackendRequest{HeatmapBackendRequest20270101OneOf4458138496: v}
}

// HeatmapBackendRequest20270101OneOf4458139936AsHeatmapBackendRequest is a convenience function that returns HeatmapBackendRequest20270101OneOf4458139936 wrapped in HeatmapBackendRequest.
func HeatmapBackendRequest20270101OneOf4458139936AsHeatmapBackendRequest(v *HeatmapBackendRequest20270101OneOf4458139936) HeatmapBackendRequest {
	return HeatmapBackendRequest{HeatmapBackendRequest20270101OneOf4458139936: v}
}

// HeatmapBackendRequest20270101OneOf4458140576AsHeatmapBackendRequest is a convenience function that returns HeatmapBackendRequest20270101OneOf4458140576 wrapped in HeatmapBackendRequest.
func HeatmapBackendRequest20270101OneOf4458140576AsHeatmapBackendRequest(v *HeatmapBackendRequest20270101OneOf4458140576) HeatmapBackendRequest {
	return HeatmapBackendRequest{HeatmapBackendRequest20270101OneOf4458140576: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *HeatmapBackendRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into HeatmapBackendRequest20270101OneOf
	err = datadog.Unmarshal(data, &obj.HeatmapBackendRequest20270101OneOf)
	if err == nil {
		if obj.HeatmapBackendRequest20270101OneOf != nil && obj.HeatmapBackendRequest20270101OneOf.UnparsedObject == nil {
			jsonHeatmapBackendRequest20270101OneOf, _ := datadog.Marshal(obj.HeatmapBackendRequest20270101OneOf)
			if string(jsonHeatmapBackendRequest20270101OneOf) == "{}" { // empty struct
				obj.HeatmapBackendRequest20270101OneOf = nil
			} else {
				match++
			}
		} else {
			obj.HeatmapBackendRequest20270101OneOf = nil
		}
	} else {
		obj.HeatmapBackendRequest20270101OneOf = nil
	}

	// try to unmarshal data into HeatmapBackendRequest20270101OneOf4458138496
	err = datadog.Unmarshal(data, &obj.HeatmapBackendRequest20270101OneOf4458138496)
	if err == nil {
		if obj.HeatmapBackendRequest20270101OneOf4458138496 != nil && obj.HeatmapBackendRequest20270101OneOf4458138496.UnparsedObject == nil {
			jsonHeatmapBackendRequest20270101OneOf4458138496, _ := datadog.Marshal(obj.HeatmapBackendRequest20270101OneOf4458138496)
			if string(jsonHeatmapBackendRequest20270101OneOf4458138496) == "{}" { // empty struct
				obj.HeatmapBackendRequest20270101OneOf4458138496 = nil
			} else {
				match++
			}
		} else {
			obj.HeatmapBackendRequest20270101OneOf4458138496 = nil
		}
	} else {
		obj.HeatmapBackendRequest20270101OneOf4458138496 = nil
	}

	// try to unmarshal data into HeatmapBackendRequest20270101OneOf4458139936
	err = datadog.Unmarshal(data, &obj.HeatmapBackendRequest20270101OneOf4458139936)
	if err == nil {
		if obj.HeatmapBackendRequest20270101OneOf4458139936 != nil && obj.HeatmapBackendRequest20270101OneOf4458139936.UnparsedObject == nil {
			jsonHeatmapBackendRequest20270101OneOf4458139936, _ := datadog.Marshal(obj.HeatmapBackendRequest20270101OneOf4458139936)
			if string(jsonHeatmapBackendRequest20270101OneOf4458139936) == "{}" { // empty struct
				obj.HeatmapBackendRequest20270101OneOf4458139936 = nil
			} else {
				match++
			}
		} else {
			obj.HeatmapBackendRequest20270101OneOf4458139936 = nil
		}
	} else {
		obj.HeatmapBackendRequest20270101OneOf4458139936 = nil
	}

	// try to unmarshal data into HeatmapBackendRequest20270101OneOf4458140576
	err = datadog.Unmarshal(data, &obj.HeatmapBackendRequest20270101OneOf4458140576)
	if err == nil {
		if obj.HeatmapBackendRequest20270101OneOf4458140576 != nil && obj.HeatmapBackendRequest20270101OneOf4458140576.UnparsedObject == nil {
			jsonHeatmapBackendRequest20270101OneOf4458140576, _ := datadog.Marshal(obj.HeatmapBackendRequest20270101OneOf4458140576)
			if string(jsonHeatmapBackendRequest20270101OneOf4458140576) == "{}" { // empty struct
				obj.HeatmapBackendRequest20270101OneOf4458140576 = nil
			} else {
				match++
			}
		} else {
			obj.HeatmapBackendRequest20270101OneOf4458140576 = nil
		}
	} else {
		obj.HeatmapBackendRequest20270101OneOf4458140576 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.HeatmapBackendRequest20270101OneOf = nil
		obj.HeatmapBackendRequest20270101OneOf4458138496 = nil
		obj.HeatmapBackendRequest20270101OneOf4458139936 = nil
		obj.HeatmapBackendRequest20270101OneOf4458140576 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj HeatmapBackendRequest) MarshalJSON() ([]byte, error) {
	if obj.HeatmapBackendRequest20270101OneOf != nil {
		return datadog.Marshal(&obj.HeatmapBackendRequest20270101OneOf)
	}

	if obj.HeatmapBackendRequest20270101OneOf4458138496 != nil {
		return datadog.Marshal(&obj.HeatmapBackendRequest20270101OneOf4458138496)
	}

	if obj.HeatmapBackendRequest20270101OneOf4458139936 != nil {
		return datadog.Marshal(&obj.HeatmapBackendRequest20270101OneOf4458139936)
	}

	if obj.HeatmapBackendRequest20270101OneOf4458140576 != nil {
		return datadog.Marshal(&obj.HeatmapBackendRequest20270101OneOf4458140576)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *HeatmapBackendRequest) GetActualInstance() interface{} {
	if obj.HeatmapBackendRequest20270101OneOf != nil {
		return obj.HeatmapBackendRequest20270101OneOf
	}

	if obj.HeatmapBackendRequest20270101OneOf4458138496 != nil {
		return obj.HeatmapBackendRequest20270101OneOf4458138496
	}

	if obj.HeatmapBackendRequest20270101OneOf4458139936 != nil {
		return obj.HeatmapBackendRequest20270101OneOf4458139936
	}

	if obj.HeatmapBackendRequest20270101OneOf4458140576 != nil {
		return obj.HeatmapBackendRequest20270101OneOf4458140576
	}

	// all schemas are nil
	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AllListRequestVariants - Test-only generated schema for the dashboard versioning proof of concept.
type AllListRequestVariants struct {
	ListStreamRequest_20270101         *ListStreamRequest
	CompositeStreamRequest_20270101    *CompositeStreamRequest
	DatasetListRequest_20270101        *DatasetListRequest
	InfrastructureListRequest_20270101 *InfrastructureListRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ListStreamRequest_20270101AsAllListRequestVariants is a convenience function that returns ListStreamRequest wrapped in AllListRequestVariants.
func ListStreamRequest_20270101AsAllListRequestVariants(v *ListStreamRequest) AllListRequestVariants {
	return AllListRequestVariants{ListStreamRequest_20270101: v}
}

// CompositeStreamRequest_20270101AsAllListRequestVariants is a convenience function that returns CompositeStreamRequest wrapped in AllListRequestVariants.
func CompositeStreamRequest_20270101AsAllListRequestVariants(v *CompositeStreamRequest) AllListRequestVariants {
	return AllListRequestVariants{CompositeStreamRequest_20270101: v}
}

// DatasetListRequest_20270101AsAllListRequestVariants is a convenience function that returns DatasetListRequest wrapped in AllListRequestVariants.
func DatasetListRequest_20270101AsAllListRequestVariants(v *DatasetListRequest) AllListRequestVariants {
	return AllListRequestVariants{DatasetListRequest_20270101: v}
}

// InfrastructureListRequest_20270101AsAllListRequestVariants is a convenience function that returns InfrastructureListRequest wrapped in AllListRequestVariants.
func InfrastructureListRequest_20270101AsAllListRequestVariants(v *InfrastructureListRequest) AllListRequestVariants {
	return AllListRequestVariants{InfrastructureListRequest_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AllListRequestVariants) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.ListStreamRequest_20270101)
	if err == nil {
		if obj.ListStreamRequest_20270101 != nil && obj.ListStreamRequest_20270101.UnparsedObject == nil {
			jsonListStreamRequest_20270101, _ := datadog.Marshal(obj.ListStreamRequest_20270101)
			if string(jsonListStreamRequest_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.ListStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ListStreamRequest_20270101 = nil
		}
	} else {
		obj.ListStreamRequest_20270101 = nil
	}

	// try to unmarshal data into CompositeStreamRequest_20270101
	err = datadog.Unmarshal(data, &obj.CompositeStreamRequest_20270101)
	if err == nil {
		if obj.CompositeStreamRequest_20270101 != nil && obj.CompositeStreamRequest_20270101.UnparsedObject == nil {
			jsonCompositeStreamRequest_20270101, _ := datadog.Marshal(obj.CompositeStreamRequest_20270101)
			if string(jsonCompositeStreamRequest_20270101) == "{}" { // empty struct
				obj.CompositeStreamRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.CompositeStreamRequest_20270101 = nil
		}
	} else {
		obj.CompositeStreamRequest_20270101 = nil
	}

	// try to unmarshal data into DatasetListRequest_20270101
	err = datadog.Unmarshal(data, &obj.DatasetListRequest_20270101)
	if err == nil {
		if obj.DatasetListRequest_20270101 != nil && obj.DatasetListRequest_20270101.UnparsedObject == nil {
			jsonDatasetListRequest_20270101, _ := datadog.Marshal(obj.DatasetListRequest_20270101)
			if string(jsonDatasetListRequest_20270101) == "{}" { // empty struct
				obj.DatasetListRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.DatasetListRequest_20270101 = nil
		}
	} else {
		obj.DatasetListRequest_20270101 = nil
	}

	// try to unmarshal data into InfrastructureListRequest_20270101
	err = datadog.Unmarshal(data, &obj.InfrastructureListRequest_20270101)
	if err == nil {
		if obj.InfrastructureListRequest_20270101 != nil && obj.InfrastructureListRequest_20270101.UnparsedObject == nil {
			jsonInfrastructureListRequest_20270101, _ := datadog.Marshal(obj.InfrastructureListRequest_20270101)
			if string(jsonInfrastructureListRequest_20270101) == "{}" { // empty struct
				obj.InfrastructureListRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.InfrastructureListRequest_20270101 = nil
		}
	} else {
		obj.InfrastructureListRequest_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ListStreamRequest_20270101 = nil
		obj.CompositeStreamRequest_20270101 = nil
		obj.DatasetListRequest_20270101 = nil
		obj.InfrastructureListRequest_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AllListRequestVariants) MarshalJSON() ([]byte, error) {
	if obj.ListStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.ListStreamRequest_20270101)
	}

	if obj.CompositeStreamRequest_20270101 != nil {
		return datadog.Marshal(&obj.CompositeStreamRequest_20270101)
	}

	if obj.DatasetListRequest_20270101 != nil {
		return datadog.Marshal(&obj.DatasetListRequest_20270101)
	}

	if obj.InfrastructureListRequest_20270101 != nil {
		return datadog.Marshal(&obj.InfrastructureListRequest_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AllListRequestVariants) GetActualInstance() interface{} {
	if obj.ListStreamRequest_20270101 != nil {
		return obj.ListStreamRequest_20270101
	}

	if obj.CompositeStreamRequest_20270101 != nil {
		return obj.CompositeStreamRequest_20270101
	}

	if obj.DatasetListRequest_20270101 != nil {
		return obj.DatasetListRequest_20270101
	}

	if obj.InfrastructureListRequest_20270101 != nil {
		return obj.InfrastructureListRequest_20270101
	}

	// all schemas are nil
	return nil
}

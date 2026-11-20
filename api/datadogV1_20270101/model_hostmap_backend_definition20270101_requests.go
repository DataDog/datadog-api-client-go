// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostmapBackendDefinition20270101Requests - Test-only generated schema for the dashboard versioning proof of concept.
type HostmapBackendDefinition20270101Requests struct {
	MetricHostmapRequests_20270101        *MetricHostmapRequests
	HostmapInfrastructureRequest_20270101 *HostmapInfrastructureRequest
	HostmapDataProjectionRequest_20270101 *HostmapDataProjectionRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// MetricHostmapRequests_20270101AsHostmapBackendDefinition20270101Requests is a convenience function that returns MetricHostmapRequests wrapped in HostmapBackendDefinition20270101Requests.
func MetricHostmapRequests_20270101AsHostmapBackendDefinition20270101Requests(v *MetricHostmapRequests) HostmapBackendDefinition20270101Requests {
	return HostmapBackendDefinition20270101Requests{MetricHostmapRequests_20270101: v}
}

// HostmapInfrastructureRequest_20270101AsHostmapBackendDefinition20270101Requests is a convenience function that returns HostmapInfrastructureRequest wrapped in HostmapBackendDefinition20270101Requests.
func HostmapInfrastructureRequest_20270101AsHostmapBackendDefinition20270101Requests(v *HostmapInfrastructureRequest) HostmapBackendDefinition20270101Requests {
	return HostmapBackendDefinition20270101Requests{HostmapInfrastructureRequest_20270101: v}
}

// HostmapDataProjectionRequest_20270101AsHostmapBackendDefinition20270101Requests is a convenience function that returns HostmapDataProjectionRequest wrapped in HostmapBackendDefinition20270101Requests.
func HostmapDataProjectionRequest_20270101AsHostmapBackendDefinition20270101Requests(v *HostmapDataProjectionRequest) HostmapBackendDefinition20270101Requests {
	return HostmapBackendDefinition20270101Requests{HostmapDataProjectionRequest_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *HostmapBackendDefinition20270101Requests) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MetricHostmapRequests_20270101
	err = datadog.Unmarshal(data, &obj.MetricHostmapRequests_20270101)
	if err == nil {
		if obj.MetricHostmapRequests_20270101 != nil && obj.MetricHostmapRequests_20270101.UnparsedObject == nil {
			jsonMetricHostmapRequests_20270101, _ := datadog.Marshal(obj.MetricHostmapRequests_20270101)
			if string(jsonMetricHostmapRequests_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.MetricHostmapRequests_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.MetricHostmapRequests_20270101 = nil
		}
	} else {
		obj.MetricHostmapRequests_20270101 = nil
	}

	// try to unmarshal data into HostmapInfrastructureRequest_20270101
	err = datadog.Unmarshal(data, &obj.HostmapInfrastructureRequest_20270101)
	if err == nil {
		if obj.HostmapInfrastructureRequest_20270101 != nil && obj.HostmapInfrastructureRequest_20270101.UnparsedObject == nil {
			jsonHostmapInfrastructureRequest_20270101, _ := datadog.Marshal(obj.HostmapInfrastructureRequest_20270101)
			if string(jsonHostmapInfrastructureRequest_20270101) == "{}" { // empty struct
				obj.HostmapInfrastructureRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.HostmapInfrastructureRequest_20270101 = nil
		}
	} else {
		obj.HostmapInfrastructureRequest_20270101 = nil
	}

	// try to unmarshal data into HostmapDataProjectionRequest_20270101
	err = datadog.Unmarshal(data, &obj.HostmapDataProjectionRequest_20270101)
	if err == nil {
		if obj.HostmapDataProjectionRequest_20270101 != nil && obj.HostmapDataProjectionRequest_20270101.UnparsedObject == nil {
			jsonHostmapDataProjectionRequest_20270101, _ := datadog.Marshal(obj.HostmapDataProjectionRequest_20270101)
			if string(jsonHostmapDataProjectionRequest_20270101) == "{}" { // empty struct
				obj.HostmapDataProjectionRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.HostmapDataProjectionRequest_20270101 = nil
		}
	} else {
		obj.HostmapDataProjectionRequest_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.MetricHostmapRequests_20270101 = nil
		obj.HostmapInfrastructureRequest_20270101 = nil
		obj.HostmapDataProjectionRequest_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj HostmapBackendDefinition20270101Requests) MarshalJSON() ([]byte, error) {
	if obj.MetricHostmapRequests_20270101 != nil {
		return datadog.Marshal(&obj.MetricHostmapRequests_20270101)
	}

	if obj.HostmapInfrastructureRequest_20270101 != nil {
		return datadog.Marshal(&obj.HostmapInfrastructureRequest_20270101)
	}

	if obj.HostmapDataProjectionRequest_20270101 != nil {
		return datadog.Marshal(&obj.HostmapDataProjectionRequest_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *HostmapBackendDefinition20270101Requests) GetActualInstance() interface{} {
	if obj.MetricHostmapRequests_20270101 != nil {
		return obj.MetricHostmapRequests_20270101
	}

	if obj.HostmapInfrastructureRequest_20270101 != nil {
		return obj.HostmapInfrastructureRequest_20270101
	}

	if obj.HostmapDataProjectionRequest_20270101 != nil {
		return obj.HostmapDataProjectionRequest_20270101
	}

	// all schemas are nil
	return nil
}

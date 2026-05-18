// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyQuery - A topology data source query.
type TopologyQuery struct {
	TopologyQueryDataStreamsOrServiceMap *TopologyQueryDataStreamsOrServiceMap

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TopologyQueryDataStreamsOrServiceMapAsTopologyQuery is a convenience function that returns TopologyQueryDataStreamsOrServiceMap wrapped in TopologyQuery.
func TopologyQueryDataStreamsOrServiceMapAsTopologyQuery(v *TopologyQueryDataStreamsOrServiceMap) TopologyQuery {
	return TopologyQuery{TopologyQueryDataStreamsOrServiceMap: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TopologyQuery) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TopologyQueryDataStreamsOrServiceMap
	err = datadog.Unmarshal(data, &obj.TopologyQueryDataStreamsOrServiceMap)
	if err == nil {
		if obj.TopologyQueryDataStreamsOrServiceMap != nil && obj.TopologyQueryDataStreamsOrServiceMap.UnparsedObject == nil {
			jsonTopologyQueryDataStreamsOrServiceMap, _ := datadog.Marshal(obj.TopologyQueryDataStreamsOrServiceMap)
			if string(jsonTopologyQueryDataStreamsOrServiceMap) == "{}" { // empty struct
				obj.TopologyQueryDataStreamsOrServiceMap = nil
			} else {
				match++
			}
		} else {
			obj.TopologyQueryDataStreamsOrServiceMap = nil
		}
	} else {
		obj.TopologyQueryDataStreamsOrServiceMap = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TopologyQueryDataStreamsOrServiceMap = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TopologyQuery) MarshalJSON() ([]byte, error) {
	if obj.TopologyQueryDataStreamsOrServiceMap != nil {
		return datadog.Marshal(&obj.TopologyQueryDataStreamsOrServiceMap)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TopologyQuery) GetActualInstance() interface{} {
	if obj.TopologyQueryDataStreamsOrServiceMap != nil {
		return obj.TopologyQueryDataStreamsOrServiceMap
	}

	// all schemas are nil
	return nil
}

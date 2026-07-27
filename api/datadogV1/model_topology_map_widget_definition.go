// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyMapWidgetDefinition - This widget displays a topology of nodes and edges for different data sources. It replaces the service map widget.
type TopologyMapWidgetDefinition struct {
	TopologyMapWidgetDefinitionDataStreams *TopologyMapWidgetDefinitionDataStreams
	TopologyMapWidgetDefinitionServiceMap  *TopologyMapWidgetDefinitionServiceMap

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TopologyMapWidgetDefinitionDataStreamsAsTopologyMapWidgetDefinition is a convenience function that returns TopologyMapWidgetDefinitionDataStreams wrapped in TopologyMapWidgetDefinition.
func TopologyMapWidgetDefinitionDataStreamsAsTopologyMapWidgetDefinition(v *TopologyMapWidgetDefinitionDataStreams) TopologyMapWidgetDefinition {
	return TopologyMapWidgetDefinition{TopologyMapWidgetDefinitionDataStreams: v}
}

// TopologyMapWidgetDefinitionServiceMapAsTopologyMapWidgetDefinition is a convenience function that returns TopologyMapWidgetDefinitionServiceMap wrapped in TopologyMapWidgetDefinition.
func TopologyMapWidgetDefinitionServiceMapAsTopologyMapWidgetDefinition(v *TopologyMapWidgetDefinitionServiceMap) TopologyMapWidgetDefinition {
	return TopologyMapWidgetDefinition{TopologyMapWidgetDefinitionServiceMap: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TopologyMapWidgetDefinition) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TopologyMapWidgetDefinitionDataStreams
	err = datadog.Unmarshal(data, &obj.TopologyMapWidgetDefinitionDataStreams)
	if err == nil {
		if obj.TopologyMapWidgetDefinitionDataStreams != nil && obj.TopologyMapWidgetDefinitionDataStreams.UnparsedObject == nil {
			jsonTopologyMapWidgetDefinitionDataStreams, _ := datadog.Marshal(obj.TopologyMapWidgetDefinitionDataStreams)
			if string(jsonTopologyMapWidgetDefinitionDataStreams) == "{}" { // empty struct
				obj.TopologyMapWidgetDefinitionDataStreams = nil
			} else {
				match++
			}
		} else {
			obj.TopologyMapWidgetDefinitionDataStreams = nil
		}
	} else {
		obj.TopologyMapWidgetDefinitionDataStreams = nil
	}

	// try to unmarshal data into TopologyMapWidgetDefinitionServiceMap
	err = datadog.Unmarshal(data, &obj.TopologyMapWidgetDefinitionServiceMap)
	if err == nil {
		if obj.TopologyMapWidgetDefinitionServiceMap != nil && obj.TopologyMapWidgetDefinitionServiceMap.UnparsedObject == nil {
			jsonTopologyMapWidgetDefinitionServiceMap, _ := datadog.Marshal(obj.TopologyMapWidgetDefinitionServiceMap)
			if string(jsonTopologyMapWidgetDefinitionServiceMap) == "{}" { // empty struct
				obj.TopologyMapWidgetDefinitionServiceMap = nil
			} else {
				match++
			}
		} else {
			obj.TopologyMapWidgetDefinitionServiceMap = nil
		}
	} else {
		obj.TopologyMapWidgetDefinitionServiceMap = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TopologyMapWidgetDefinitionDataStreams = nil
		obj.TopologyMapWidgetDefinitionServiceMap = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TopologyMapWidgetDefinition) MarshalJSON() ([]byte, error) {
	if obj.TopologyMapWidgetDefinitionDataStreams != nil {
		return datadog.Marshal(&obj.TopologyMapWidgetDefinitionDataStreams)
	}

	if obj.TopologyMapWidgetDefinitionServiceMap != nil {
		return datadog.Marshal(&obj.TopologyMapWidgetDefinitionServiceMap)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TopologyMapWidgetDefinition) GetActualInstance() interface{} {
	if obj.TopologyMapWidgetDefinitionDataStreams != nil {
		return obj.TopologyMapWidgetDefinitionDataStreams
	}

	if obj.TopologyMapWidgetDefinitionServiceMap != nil {
		return obj.TopologyMapWidgetDefinitionServiceMap
	}

	// all schemas are nil
	return nil
}

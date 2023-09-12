// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"
)

// CustomDestinationForwarderDestination - The destination to forward events to.
type CustomDestinationForwarderDestination struct {
	HttpDestination          *HttpDestination
	ElasticsearchDestination *ElasticsearchDestination
	SplunkHecDestination     *SplunkHecDestination

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// HttpDestinationAsCustomDestinationForwarderDestination is a convenience function that returns HttpDestination wrapped in CustomDestinationForwarderDestination.
func HttpDestinationAsCustomDestinationForwarderDestination(v *HttpDestination) CustomDestinationForwarderDestination {
	return CustomDestinationForwarderDestination{HttpDestination: v}
}

// ElasticsearchDestinationAsCustomDestinationForwarderDestination is a convenience function that returns ElasticsearchDestination wrapped in CustomDestinationForwarderDestination.
func ElasticsearchDestinationAsCustomDestinationForwarderDestination(v *ElasticsearchDestination) CustomDestinationForwarderDestination {
	return CustomDestinationForwarderDestination{ElasticsearchDestination: v}
}

// SplunkHecDestinationAsCustomDestinationForwarderDestination is a convenience function that returns SplunkHecDestination wrapped in CustomDestinationForwarderDestination.
func SplunkHecDestinationAsCustomDestinationForwarderDestination(v *SplunkHecDestination) CustomDestinationForwarderDestination {
	return CustomDestinationForwarderDestination{SplunkHecDestination: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CustomDestinationForwarderDestination) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into HttpDestination
	err = json.Unmarshal(data, &obj.HttpDestination)
	if err == nil {
		if obj.HttpDestination != nil && obj.HttpDestination.UnparsedObject == nil {
			jsonHttpDestination, _ := json.Marshal(obj.HttpDestination)
			if string(jsonHttpDestination) == "{}" { // empty struct
				obj.HttpDestination = nil
			} else {
				match++
			}
		} else {
			obj.HttpDestination = nil
		}
	} else {
		obj.HttpDestination = nil
	}

	// try to unmarshal data into ElasticsearchDestination
	err = json.Unmarshal(data, &obj.ElasticsearchDestination)
	if err == nil {
		if obj.ElasticsearchDestination != nil && obj.ElasticsearchDestination.UnparsedObject == nil {
			jsonElasticsearchDestination, _ := json.Marshal(obj.ElasticsearchDestination)
			if string(jsonElasticsearchDestination) == "{}" { // empty struct
				obj.ElasticsearchDestination = nil
			} else {
				match++
			}
		} else {
			obj.ElasticsearchDestination = nil
		}
	} else {
		obj.ElasticsearchDestination = nil
	}

	// try to unmarshal data into SplunkHecDestination
	err = json.Unmarshal(data, &obj.SplunkHecDestination)
	if err == nil {
		if obj.SplunkHecDestination != nil && obj.SplunkHecDestination.UnparsedObject == nil {
			jsonSplunkHecDestination, _ := json.Marshal(obj.SplunkHecDestination)
			if string(jsonSplunkHecDestination) == "{}" { // empty struct
				obj.SplunkHecDestination = nil
			} else {
				match++
			}
		} else {
			obj.SplunkHecDestination = nil
		}
	} else {
		obj.SplunkHecDestination = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.HttpDestination = nil
		obj.ElasticsearchDestination = nil
		obj.SplunkHecDestination = nil
		return json.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CustomDestinationForwarderDestination) MarshalJSON() ([]byte, error) {
	if obj.HttpDestination != nil {
		return json.Marshal(&obj.HttpDestination)
	}

	if obj.ElasticsearchDestination != nil {
		return json.Marshal(&obj.ElasticsearchDestination)
	}

	if obj.SplunkHecDestination != nil {
		return json.Marshal(&obj.SplunkHecDestination)
	}

	if obj.UnparsedObject != nil {
		return json.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CustomDestinationForwarderDestination) GetActualInstance() interface{} {
	if obj.HttpDestination != nil {
		return obj.HttpDestination
	}

	if obj.ElasticsearchDestination != nil {
		return obj.ElasticsearchDestination
	}

	if obj.SplunkHecDestination != nil {
		return obj.SplunkHecDestination
	}

	// all schemas are nil
	return nil
}

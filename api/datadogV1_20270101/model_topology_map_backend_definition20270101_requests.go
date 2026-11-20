// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.


package datadogV1_20270101

import (
	_io "io"
	"time"
	"github.com/google/uuid"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"

)


// TopologyMapBackendDefinition20270101Requests - Test-only generated schema for the dashboard versioning proof of concept.
type TopologyMapBackendDefinition20270101Requests struct {
	[]interface{} *[]interface{}
	[]interface{} *[]interface{}
	[]interface{} *[]interface{}
	[]interface{} *[]interface{}
	[]CustomTopologyMapTileDefRequest *[]CustomTopologyMapTileDefRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// []interface{}AsTopologyMapBackendDefinition20270101Requests is a convenience function that returns []interface{} wrapped in TopologyMapBackendDefinition20270101Requests.
func []interface{}AsTopologyMapBackendDefinition20270101Requests(v *[]interface{}) TopologyMapBackendDefinition20270101Requests {
	return TopologyMapBackendDefinition20270101Requests{[]interface{}: v}
}

// []interface{}AsTopologyMapBackendDefinition20270101Requests is a convenience function that returns []interface{} wrapped in TopologyMapBackendDefinition20270101Requests.
func []interface{}AsTopologyMapBackendDefinition20270101Requests(v *[]interface{}) TopologyMapBackendDefinition20270101Requests {
	return TopologyMapBackendDefinition20270101Requests{[]interface{}: v}
}

// []interface{}AsTopologyMapBackendDefinition20270101Requests is a convenience function that returns []interface{} wrapped in TopologyMapBackendDefinition20270101Requests.
func []interface{}AsTopologyMapBackendDefinition20270101Requests(v *[]interface{}) TopologyMapBackendDefinition20270101Requests {
	return TopologyMapBackendDefinition20270101Requests{[]interface{}: v}
}

// []interface{}AsTopologyMapBackendDefinition20270101Requests is a convenience function that returns []interface{} wrapped in TopologyMapBackendDefinition20270101Requests.
func []interface{}AsTopologyMapBackendDefinition20270101Requests(v *[]interface{}) TopologyMapBackendDefinition20270101Requests {
	return TopologyMapBackendDefinition20270101Requests{[]interface{}: v}
}

// []CustomTopologyMapTileDefRequestAsTopologyMapBackendDefinition20270101Requests is a convenience function that returns []CustomTopologyMapTileDefRequest wrapped in TopologyMapBackendDefinition20270101Requests.
func []CustomTopologyMapTileDefRequestAsTopologyMapBackendDefinition20270101Requests(v *[]CustomTopologyMapTileDefRequest) TopologyMapBackendDefinition20270101Requests {
	return TopologyMapBackendDefinition20270101Requests{[]CustomTopologyMapTileDefRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TopologyMapBackendDefinition20270101Requests) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into []interface{}
	err = datadog.Unmarshal(data, &obj.[]interface{})
	if err == nil {
		if obj.[]interface{} != nil {
			json[]interface{}, _ := datadog.Marshal(obj.[]interface{})
			if string(json[]interface{}) == "{}" && string(data) != "{}"  { // empty struct
				obj.[]interface{} = nil
			} else {
				match++
			}
		} else {
			obj.[]interface{} = nil
		}
	} else {
		obj.[]interface{} = nil
	}

	// try to unmarshal data into []interface{}
	err = datadog.Unmarshal(data, &obj.[]interface{})
	if err == nil {
		if obj.[]interface{} != nil {
			json[]interface{}, _ := datadog.Marshal(obj.[]interface{})
			if string(json[]interface{}) == "{}" && string(data) != "{}"  { // empty struct
				obj.[]interface{} = nil
			} else {
				match++
			}
		} else {
			obj.[]interface{} = nil
		}
	} else {
		obj.[]interface{} = nil
	}

	// try to unmarshal data into []interface{}
	err = datadog.Unmarshal(data, &obj.[]interface{})
	if err == nil {
		if obj.[]interface{} != nil {
			json[]interface{}, _ := datadog.Marshal(obj.[]interface{})
			if string(json[]interface{}) == "{}" && string(data) != "{}"  { // empty struct
				obj.[]interface{} = nil
			} else {
				match++
			}
		} else {
			obj.[]interface{} = nil
		}
	} else {
		obj.[]interface{} = nil
	}

	// try to unmarshal data into []interface{}
	err = datadog.Unmarshal(data, &obj.[]interface{})
	if err == nil {
		if obj.[]interface{} != nil {
			json[]interface{}, _ := datadog.Marshal(obj.[]interface{})
			if string(json[]interface{}) == "{}" && string(data) != "{}"  { // empty struct
				obj.[]interface{} = nil
			} else {
				match++
			}
		} else {
			obj.[]interface{} = nil
		}
	} else {
		obj.[]interface{} = nil
	}

	// try to unmarshal data into []CustomTopologyMapTileDefRequest
	err = datadog.Unmarshal(data, &obj.[]CustomTopologyMapTileDefRequest)
	if err == nil {
		if obj.[]CustomTopologyMapTileDefRequest != nil {
			json[]CustomTopologyMapTileDefRequest, _ := datadog.Marshal(obj.[]CustomTopologyMapTileDefRequest)
			if string(json[]CustomTopologyMapTileDefRequest) == "{}" && string(data) != "{}"  { // empty struct
				obj.[]CustomTopologyMapTileDefRequest = nil
			} else {
				match++
			}
		} else {
			obj.[]CustomTopologyMapTileDefRequest = nil
		}
	} else {
		obj.[]CustomTopologyMapTileDefRequest = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.[]interface{} = nil
		obj.[]interface{} = nil
		obj.[]interface{} = nil
		obj.[]interface{} = nil
		obj.[]CustomTopologyMapTileDefRequest = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TopologyMapBackendDefinition20270101Requests) MarshalJSON() ([]byte, error) {
	if obj.[]interface{} != nil {
		return datadog.Marshal(&obj.[]interface{})
	}


	if obj.[]interface{} != nil {
		return datadog.Marshal(&obj.[]interface{})
	}


	if obj.[]interface{} != nil {
		return datadog.Marshal(&obj.[]interface{})
	}


	if obj.[]interface{} != nil {
		return datadog.Marshal(&obj.[]interface{})
	}


	if obj.[]CustomTopologyMapTileDefRequest != nil {
		return datadog.Marshal(&obj.[]CustomTopologyMapTileDefRequest)
	}


	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TopologyMapBackendDefinition20270101Requests) GetActualInstance() (interface{}) {
	if obj.[]interface{} != nil {
		return obj.[]interface{}
	}


	if obj.[]interface{} != nil {
		return obj.[]interface{}
	}


	if obj.[]interface{} != nil {
		return obj.[]interface{}
	}


	if obj.[]interface{} != nil {
		return obj.[]interface{}
	}


	if obj.[]CustomTopologyMapTileDefRequest != nil {
		return obj.[]CustomTopologyMapTileDefRequest
	}


	// all schemas are nil
	return nil
}

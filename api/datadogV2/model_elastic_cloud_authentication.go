// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudAuthentication - Authentication methods supported by the Elastic Cloud interface. Exactly one is set, selected by its `type`.
type ElasticCloudAuthentication struct {
	ElasticCloudBasicAuth *ElasticCloudBasicAuth

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ElasticCloudBasicAuthAsElasticCloudAuthentication is a convenience function that returns ElasticCloudBasicAuth wrapped in ElasticCloudAuthentication.
func ElasticCloudBasicAuthAsElasticCloudAuthentication(v *ElasticCloudBasicAuth) ElasticCloudAuthentication {
	return ElasticCloudAuthentication{ElasticCloudBasicAuth: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ElasticCloudAuthentication) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ElasticCloudBasicAuth
	err = datadog.Unmarshal(data, &obj.ElasticCloudBasicAuth)
	if err == nil {
		if obj.ElasticCloudBasicAuth != nil && obj.ElasticCloudBasicAuth.UnparsedObject == nil {
			jsonElasticCloudBasicAuth, _ := datadog.Marshal(obj.ElasticCloudBasicAuth)
			if string(jsonElasticCloudBasicAuth) == "{}" { // empty struct
				obj.ElasticCloudBasicAuth = nil
			} else {
				match++
			}
		} else {
			obj.ElasticCloudBasicAuth = nil
		}
	} else {
		obj.ElasticCloudBasicAuth = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ElasticCloudBasicAuth = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ElasticCloudAuthentication) MarshalJSON() ([]byte, error) {
	if obj.ElasticCloudBasicAuth != nil {
		return datadog.Marshal(&obj.ElasticCloudBasicAuth)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ElasticCloudAuthentication) GetActualInstance() interface{} {
	if obj.ElasticCloudBasicAuth != nil {
		return obj.ElasticCloudBasicAuth
	}

	// all schemas are nil
	return nil
}

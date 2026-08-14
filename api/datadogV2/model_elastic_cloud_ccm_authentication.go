// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudCcmAuthentication - Authentication methods supported by the Elastic Cloud CCM interface. Exactly one is set, selected by its `type`.
type ElasticCloudCcmAuthentication struct {
	ElasticCloudCcmTokenAuth *ElasticCloudCcmTokenAuth

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ElasticCloudCcmTokenAuthAsElasticCloudCcmAuthentication is a convenience function that returns ElasticCloudCcmTokenAuth wrapped in ElasticCloudCcmAuthentication.
func ElasticCloudCcmTokenAuthAsElasticCloudCcmAuthentication(v *ElasticCloudCcmTokenAuth) ElasticCloudCcmAuthentication {
	return ElasticCloudCcmAuthentication{ElasticCloudCcmTokenAuth: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ElasticCloudCcmAuthentication) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ElasticCloudCcmTokenAuth
	err = datadog.Unmarshal(data, &obj.ElasticCloudCcmTokenAuth)
	if err == nil {
		if obj.ElasticCloudCcmTokenAuth != nil && obj.ElasticCloudCcmTokenAuth.UnparsedObject == nil {
			jsonElasticCloudCcmTokenAuth, _ := datadog.Marshal(obj.ElasticCloudCcmTokenAuth)
			if string(jsonElasticCloudCcmTokenAuth) == "{}" { // empty struct
				obj.ElasticCloudCcmTokenAuth = nil
			} else {
				match++
			}
		} else {
			obj.ElasticCloudCcmTokenAuth = nil
		}
	} else {
		obj.ElasticCloudCcmTokenAuth = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ElasticCloudCcmTokenAuth = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ElasticCloudCcmAuthentication) MarshalJSON() ([]byte, error) {
	if obj.ElasticCloudCcmTokenAuth != nil {
		return datadog.Marshal(&obj.ElasticCloudCcmTokenAuth)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ElasticCloudCcmAuthentication) GetActualInstance() interface{} {
	if obj.ElasticCloudCcmTokenAuth != nil {
		return obj.ElasticCloudCcmTokenAuth
	}

	// all schemas are nil
	return nil
}

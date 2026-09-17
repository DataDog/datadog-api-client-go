// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudIntegrationAccountAuthenticationRequest - Authentication for creating the Elastic Cloud integration account. Exactly one method is set.
type ElasticCloudIntegrationAccountAuthenticationRequest struct {
	ElasticCloudIntegrationAccountBasicAuthRequest *ElasticCloudIntegrationAccountBasicAuthRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ElasticCloudIntegrationAccountBasicAuthRequestAsElasticCloudIntegrationAccountAuthenticationRequest is a convenience function that returns ElasticCloudIntegrationAccountBasicAuthRequest wrapped in ElasticCloudIntegrationAccountAuthenticationRequest.
func ElasticCloudIntegrationAccountBasicAuthRequestAsElasticCloudIntegrationAccountAuthenticationRequest(v *ElasticCloudIntegrationAccountBasicAuthRequest) ElasticCloudIntegrationAccountAuthenticationRequest {
	return ElasticCloudIntegrationAccountAuthenticationRequest{ElasticCloudIntegrationAccountBasicAuthRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ElasticCloudIntegrationAccountAuthenticationRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ElasticCloudIntegrationAccountBasicAuthRequest
	err = datadog.Unmarshal(data, &obj.ElasticCloudIntegrationAccountBasicAuthRequest)
	if err == nil {
		if obj.ElasticCloudIntegrationAccountBasicAuthRequest != nil && obj.ElasticCloudIntegrationAccountBasicAuthRequest.UnparsedObject == nil {
			jsonElasticCloudIntegrationAccountBasicAuthRequest, _ := datadog.Marshal(obj.ElasticCloudIntegrationAccountBasicAuthRequest)
			if string(jsonElasticCloudIntegrationAccountBasicAuthRequest) == "{}" { // empty struct
				obj.ElasticCloudIntegrationAccountBasicAuthRequest = nil
			} else {
				match++
			}
		} else {
			obj.ElasticCloudIntegrationAccountBasicAuthRequest = nil
		}
	} else {
		obj.ElasticCloudIntegrationAccountBasicAuthRequest = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ElasticCloudIntegrationAccountBasicAuthRequest = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ElasticCloudIntegrationAccountAuthenticationRequest) MarshalJSON() ([]byte, error) {
	if obj.ElasticCloudIntegrationAccountBasicAuthRequest != nil {
		return datadog.Marshal(&obj.ElasticCloudIntegrationAccountBasicAuthRequest)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ElasticCloudIntegrationAccountAuthenticationRequest) GetActualInstance() interface{} {
	if obj.ElasticCloudIntegrationAccountBasicAuthRequest != nil {
		return obj.ElasticCloudIntegrationAccountBasicAuthRequest
	}

	// all schemas are nil
	return nil
}

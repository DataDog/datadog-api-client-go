// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudIntegrationAccountAuthenticationRequest - Authentication for creating the Elastic Cloud integration account. Exactly one method is set.
type ElasticCloudIntegrationAccountAuthenticationRequest struct {
	IntegrationAccountBasicAuthRequest *IntegrationAccountBasicAuthRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// IntegrationAccountBasicAuthRequestAsElasticCloudIntegrationAccountAuthenticationRequest is a convenience function that returns IntegrationAccountBasicAuthRequest wrapped in ElasticCloudIntegrationAccountAuthenticationRequest.
func IntegrationAccountBasicAuthRequestAsElasticCloudIntegrationAccountAuthenticationRequest(v *IntegrationAccountBasicAuthRequest) ElasticCloudIntegrationAccountAuthenticationRequest {
	return ElasticCloudIntegrationAccountAuthenticationRequest{IntegrationAccountBasicAuthRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ElasticCloudIntegrationAccountAuthenticationRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IntegrationAccountBasicAuthRequest
	err = datadog.Unmarshal(data, &obj.IntegrationAccountBasicAuthRequest)
	if err == nil {
		if obj.IntegrationAccountBasicAuthRequest != nil && obj.IntegrationAccountBasicAuthRequest.UnparsedObject == nil {
			jsonIntegrationAccountBasicAuthRequest, _ := datadog.Marshal(obj.IntegrationAccountBasicAuthRequest)
			if string(jsonIntegrationAccountBasicAuthRequest) == "{}" { // empty struct
				obj.IntegrationAccountBasicAuthRequest = nil
			} else {
				match++
			}
		} else {
			obj.IntegrationAccountBasicAuthRequest = nil
		}
	} else {
		obj.IntegrationAccountBasicAuthRequest = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.IntegrationAccountBasicAuthRequest = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ElasticCloudIntegrationAccountAuthenticationRequest) MarshalJSON() ([]byte, error) {
	if obj.IntegrationAccountBasicAuthRequest != nil {
		return datadog.Marshal(&obj.IntegrationAccountBasicAuthRequest)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ElasticCloudIntegrationAccountAuthenticationRequest) GetActualInstance() interface{} {
	if obj.IntegrationAccountBasicAuthRequest != nil {
		return obj.IntegrationAccountBasicAuthRequest
	}

	// all schemas are nil
	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudIntegrationAccountAuthenticationResponse - Authentication configured on the Elastic Cloud integration account.
type ElasticCloudIntegrationAccountAuthenticationResponse struct {
	IntegrationAccountBasicAuthResponse *IntegrationAccountBasicAuthResponse

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// IntegrationAccountBasicAuthResponseAsElasticCloudIntegrationAccountAuthenticationResponse is a convenience function that returns IntegrationAccountBasicAuthResponse wrapped in ElasticCloudIntegrationAccountAuthenticationResponse.
func IntegrationAccountBasicAuthResponseAsElasticCloudIntegrationAccountAuthenticationResponse(v *IntegrationAccountBasicAuthResponse) ElasticCloudIntegrationAccountAuthenticationResponse {
	return ElasticCloudIntegrationAccountAuthenticationResponse{IntegrationAccountBasicAuthResponse: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ElasticCloudIntegrationAccountAuthenticationResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IntegrationAccountBasicAuthResponse
	err = datadog.Unmarshal(data, &obj.IntegrationAccountBasicAuthResponse)
	if err == nil {
		if obj.IntegrationAccountBasicAuthResponse != nil && obj.IntegrationAccountBasicAuthResponse.UnparsedObject == nil {
			jsonIntegrationAccountBasicAuthResponse, _ := datadog.Marshal(obj.IntegrationAccountBasicAuthResponse)
			if string(jsonIntegrationAccountBasicAuthResponse) == "{}" { // empty struct
				obj.IntegrationAccountBasicAuthResponse = nil
			} else {
				match++
			}
		} else {
			obj.IntegrationAccountBasicAuthResponse = nil
		}
	} else {
		obj.IntegrationAccountBasicAuthResponse = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.IntegrationAccountBasicAuthResponse = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ElasticCloudIntegrationAccountAuthenticationResponse) MarshalJSON() ([]byte, error) {
	if obj.IntegrationAccountBasicAuthResponse != nil {
		return datadog.Marshal(&obj.IntegrationAccountBasicAuthResponse)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ElasticCloudIntegrationAccountAuthenticationResponse) GetActualInstance() interface{} {
	if obj.IntegrationAccountBasicAuthResponse != nil {
		return obj.IntegrationAccountBasicAuthResponse
	}

	// all schemas are nil
	return nil
}

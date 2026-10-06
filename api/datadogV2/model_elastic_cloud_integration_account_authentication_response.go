// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudIntegrationAccountAuthenticationResponse - Authentication configured on the Elastic Cloud integration account.
type ElasticCloudIntegrationAccountAuthenticationResponse struct {
	ElasticCloudIntegrationAccountBasicAuthResponse *ElasticCloudIntegrationAccountBasicAuthResponse

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ElasticCloudIntegrationAccountBasicAuthResponseAsElasticCloudIntegrationAccountAuthenticationResponse is a convenience function that returns ElasticCloudIntegrationAccountBasicAuthResponse wrapped in ElasticCloudIntegrationAccountAuthenticationResponse.
func ElasticCloudIntegrationAccountBasicAuthResponseAsElasticCloudIntegrationAccountAuthenticationResponse(v *ElasticCloudIntegrationAccountBasicAuthResponse) ElasticCloudIntegrationAccountAuthenticationResponse {
	return ElasticCloudIntegrationAccountAuthenticationResponse{ElasticCloudIntegrationAccountBasicAuthResponse: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ElasticCloudIntegrationAccountAuthenticationResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ElasticCloudIntegrationAccountBasicAuthResponse
	err = datadog.Unmarshal(data, &obj.ElasticCloudIntegrationAccountBasicAuthResponse)
	if err == nil {
		if obj.ElasticCloudIntegrationAccountBasicAuthResponse != nil && obj.ElasticCloudIntegrationAccountBasicAuthResponse.UnparsedObject == nil {
			jsonElasticCloudIntegrationAccountBasicAuthResponse, _ := datadog.Marshal(obj.ElasticCloudIntegrationAccountBasicAuthResponse)
			if string(jsonElasticCloudIntegrationAccountBasicAuthResponse) == "{}" { // empty struct
				obj.ElasticCloudIntegrationAccountBasicAuthResponse = nil
			} else {
				match++
			}
		} else {
			obj.ElasticCloudIntegrationAccountBasicAuthResponse = nil
		}
	} else {
		obj.ElasticCloudIntegrationAccountBasicAuthResponse = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ElasticCloudIntegrationAccountBasicAuthResponse = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ElasticCloudIntegrationAccountAuthenticationResponse) MarshalJSON() ([]byte, error) {
	if obj.ElasticCloudIntegrationAccountBasicAuthResponse != nil {
		return datadog.Marshal(&obj.ElasticCloudIntegrationAccountBasicAuthResponse)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ElasticCloudIntegrationAccountAuthenticationResponse) GetActualInstance() interface{} {
	if obj.ElasticCloudIntegrationAccountBasicAuthResponse != nil {
		return obj.ElasticCloudIntegrationAccountBasicAuthResponse
	}

	// all schemas are nil
	return nil
}

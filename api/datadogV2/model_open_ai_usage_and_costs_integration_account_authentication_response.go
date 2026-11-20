// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAIUsageAndCostsIntegrationAccountAuthenticationResponse - Authentication configured on the OpenAI Usage and Costs integration account.
type OpenAIUsageAndCostsIntegrationAccountAuthenticationResponse struct {
	OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse *OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponseAsOpenAIUsageAndCostsIntegrationAccountAuthenticationResponse is a convenience function that returns OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse wrapped in OpenAIUsageAndCostsIntegrationAccountAuthenticationResponse.
func OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponseAsOpenAIUsageAndCostsIntegrationAccountAuthenticationResponse(v *OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse) OpenAIUsageAndCostsIntegrationAccountAuthenticationResponse {
	return OpenAIUsageAndCostsIntegrationAccountAuthenticationResponse{OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *OpenAIUsageAndCostsIntegrationAccountAuthenticationResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse
	err = datadog.Unmarshal(data, &obj.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse)
	if err == nil {
		if obj.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse != nil && obj.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse.UnparsedObject == nil {
			jsonOpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse, _ := datadog.Marshal(obj.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse)
			if string(jsonOpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse) == "{}" { // empty struct
				obj.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse = nil
			} else {
				match++
			}
		} else {
			obj.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse = nil
		}
	} else {
		obj.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj OpenAIUsageAndCostsIntegrationAccountAuthenticationResponse) MarshalJSON() ([]byte, error) {
	if obj.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse != nil {
		return datadog.Marshal(&obj.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *OpenAIUsageAndCostsIntegrationAccountAuthenticationResponse) GetActualInstance() interface{} {
	if obj.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse != nil {
		return obj.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse
	}

	// all schemas are nil
	return nil
}

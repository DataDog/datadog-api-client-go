// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAIUsageAndCostsIntegrationAccountAuthenticationRequest - Authentication for creating the OpenAI Usage and Costs integration account. Exactly one method is set.
type OpenAIUsageAndCostsIntegrationAccountAuthenticationRequest struct {
	OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthRequest *OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthRequestAsOpenAIUsageAndCostsIntegrationAccountAuthenticationRequest is a convenience function that returns OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthRequest wrapped in OpenAIUsageAndCostsIntegrationAccountAuthenticationRequest.
func OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthRequestAsOpenAIUsageAndCostsIntegrationAccountAuthenticationRequest(v *OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthRequest) OpenAIUsageAndCostsIntegrationAccountAuthenticationRequest {
	return OpenAIUsageAndCostsIntegrationAccountAuthenticationRequest{OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *OpenAIUsageAndCostsIntegrationAccountAuthenticationRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthRequest
	err = datadog.Unmarshal(data, &obj.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthRequest)
	if err == nil {
		if obj.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthRequest != nil && obj.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthRequest.UnparsedObject == nil {
			jsonOpenAIUsageAndCostsIntegrationAccountBearerTokenAuthRequest, _ := datadog.Marshal(obj.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthRequest)
			if string(jsonOpenAIUsageAndCostsIntegrationAccountBearerTokenAuthRequest) == "{}" { // empty struct
				obj.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthRequest = nil
			} else {
				match++
			}
		} else {
			obj.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthRequest = nil
		}
	} else {
		obj.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthRequest = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthRequest = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj OpenAIUsageAndCostsIntegrationAccountAuthenticationRequest) MarshalJSON() ([]byte, error) {
	if obj.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthRequest != nil {
		return datadog.Marshal(&obj.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthRequest)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *OpenAIUsageAndCostsIntegrationAccountAuthenticationRequest) GetActualInstance() interface{} {
	if obj.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthRequest != nil {
		return obj.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthRequest
	}

	// all schemas are nil
	return nil
}

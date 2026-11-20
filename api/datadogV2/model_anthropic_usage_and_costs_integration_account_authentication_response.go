// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnthropicUsageAndCostsIntegrationAccountAuthenticationResponse - Authentication configured on the Anthropic Usage and Costs integration account.
type AnthropicUsageAndCostsIntegrationAccountAuthenticationResponse struct {
	AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponseAsAnthropicUsageAndCostsIntegrationAccountAuthenticationResponse is a convenience function that returns AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse wrapped in AnthropicUsageAndCostsIntegrationAccountAuthenticationResponse.
func AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponseAsAnthropicUsageAndCostsIntegrationAccountAuthenticationResponse(v *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse) AnthropicUsageAndCostsIntegrationAccountAuthenticationResponse {
	return AnthropicUsageAndCostsIntegrationAccountAuthenticationResponse{AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AnthropicUsageAndCostsIntegrationAccountAuthenticationResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse
	err = datadog.Unmarshal(data, &obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse)
	if err == nil {
		if obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse != nil && obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse.UnparsedObject == nil {
			jsonAnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse, _ := datadog.Marshal(obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse)
			if string(jsonAnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse) == "{}" { // empty struct
				obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse = nil
			} else {
				match++
			}
		} else {
			obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse = nil
		}
	} else {
		obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AnthropicUsageAndCostsIntegrationAccountAuthenticationResponse) MarshalJSON() ([]byte, error) {
	if obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse != nil {
		return datadog.Marshal(&obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AnthropicUsageAndCostsIntegrationAccountAuthenticationResponse) GetActualInstance() interface{} {
	if obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse != nil {
		return obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse
	}

	// all schemas are nil
	return nil
}

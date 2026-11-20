// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnthropicUsageAndCostsIntegrationAccountAuthenticationRequest - Authentication for creating the Anthropic Usage and Costs integration account. Exactly one method is set.
type AnthropicUsageAndCostsIntegrationAccountAuthenticationRequest struct {
	AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequestAsAnthropicUsageAndCostsIntegrationAccountAuthenticationRequest is a convenience function that returns AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest wrapped in AnthropicUsageAndCostsIntegrationAccountAuthenticationRequest.
func AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequestAsAnthropicUsageAndCostsIntegrationAccountAuthenticationRequest(v *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest) AnthropicUsageAndCostsIntegrationAccountAuthenticationRequest {
	return AnthropicUsageAndCostsIntegrationAccountAuthenticationRequest{AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AnthropicUsageAndCostsIntegrationAccountAuthenticationRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest
	err = datadog.Unmarshal(data, &obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest)
	if err == nil {
		if obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest != nil && obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest.UnparsedObject == nil {
			jsonAnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest, _ := datadog.Marshal(obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest)
			if string(jsonAnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest) == "{}" { // empty struct
				obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest = nil
			} else {
				match++
			}
		} else {
			obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest = nil
		}
	} else {
		obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AnthropicUsageAndCostsIntegrationAccountAuthenticationRequest) MarshalJSON() ([]byte, error) {
	if obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest != nil {
		return datadog.Marshal(&obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AnthropicUsageAndCostsIntegrationAccountAuthenticationRequest) GetActualInstance() interface{} {
	if obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest != nil {
		return obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest
	}

	// all schemas are nil
	return nil
}

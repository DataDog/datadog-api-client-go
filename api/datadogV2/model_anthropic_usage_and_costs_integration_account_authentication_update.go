// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnthropicUsageAndCostsIntegrationAccountAuthenticationUpdate - Authentication for updating the Anthropic Usage and Costs integration account. Exactly one method is set.
type AnthropicUsageAndCostsIntegrationAccountAuthenticationUpdate struct {
	AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthUpdate *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthUpdateAsAnthropicUsageAndCostsIntegrationAccountAuthenticationUpdate is a convenience function that returns AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthUpdate wrapped in AnthropicUsageAndCostsIntegrationAccountAuthenticationUpdate.
func AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthUpdateAsAnthropicUsageAndCostsIntegrationAccountAuthenticationUpdate(v *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthUpdate) AnthropicUsageAndCostsIntegrationAccountAuthenticationUpdate {
	return AnthropicUsageAndCostsIntegrationAccountAuthenticationUpdate{AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AnthropicUsageAndCostsIntegrationAccountAuthenticationUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthUpdate
	err = datadog.Unmarshal(data, &obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthUpdate)
	if err == nil {
		if obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthUpdate != nil && obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthUpdate.UnparsedObject == nil {
			jsonAnthropicUsageAndCostsIntegrationAccountBearerTokenAuthUpdate, _ := datadog.Marshal(obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthUpdate)
			if string(jsonAnthropicUsageAndCostsIntegrationAccountBearerTokenAuthUpdate) == "{}" { // empty struct
				obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthUpdate = nil
			} else {
				match++
			}
		} else {
			obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthUpdate = nil
		}
	} else {
		obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AnthropicUsageAndCostsIntegrationAccountAuthenticationUpdate) MarshalJSON() ([]byte, error) {
	if obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthUpdate != nil {
		return datadog.Marshal(&obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AnthropicUsageAndCostsIntegrationAccountAuthenticationUpdate) GetActualInstance() interface{} {
	if obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthUpdate != nil {
		return obj.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthUpdate
	}

	// all schemas are nil
	return nil
}

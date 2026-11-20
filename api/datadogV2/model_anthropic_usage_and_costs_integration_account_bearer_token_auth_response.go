// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse The bearer token authentication method configured on the account.
type AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse struct {
	// The authentication method type.
	AuthType AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType `json:"auth_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse instantiates a new AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse(authType AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType) *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse {
	this := AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse{}
	this.AuthType = authType
	return &this
}

// NewAnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponseWithDefaults instantiates a new AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponseWithDefaults() *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse {
	this := AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse{}
	var authType AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType = ANTHROPICUSAGEANDCOSTSINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse) GetAuthType() AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType {
	if o == nil {
		var ret AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse) GetAuthTypeOk() (*AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse) SetAuthType(v AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType) {
	o.AuthType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth_type"] = o.AuthType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType `json:"auth_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthType == nil {
		return fmt.Errorf("required field auth_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth_type"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.AuthType.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthType = *all.AuthType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse The bearer token authentication method configured on the account.
type OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse struct {
	// The authentication method type.
	AuthType OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType `json:"auth_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse instantiates a new OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse(authType OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType) *OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse {
	this := OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse{}
	this.AuthType = authType
	return &this
}

// NewOpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponseWithDefaults instantiates a new OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponseWithDefaults() *OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse {
	this := OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse{}
	var authType OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType = OPENAIUSAGEANDCOSTSINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse) GetAuthType() OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType {
	if o == nil {
		var ret OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse) GetAuthTypeOk() (*OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse) SetAuthType(v OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType) {
	o.AuthType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse) MarshalJSON() ([]byte, error) {
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
func (o *OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType *OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthType `json:"auth_type"`
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

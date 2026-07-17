// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringIntegrationConfigOktaSecrets Credentials for an Okta entity context sync.
type SecurityMonitoringIntegrationConfigOktaSecrets struct {
	// The Okta API token used to authenticate against the Okta API.
	ApiToken string `json:"api_token"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringIntegrationConfigOktaSecrets instantiates a new SecurityMonitoringIntegrationConfigOktaSecrets object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringIntegrationConfigOktaSecrets(apiToken string) *SecurityMonitoringIntegrationConfigOktaSecrets {
	this := SecurityMonitoringIntegrationConfigOktaSecrets{}
	this.ApiToken = apiToken
	return &this
}

// NewSecurityMonitoringIntegrationConfigOktaSecretsWithDefaults instantiates a new SecurityMonitoringIntegrationConfigOktaSecrets object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringIntegrationConfigOktaSecretsWithDefaults() *SecurityMonitoringIntegrationConfigOktaSecrets {
	this := SecurityMonitoringIntegrationConfigOktaSecrets{}
	return &this
}

// GetApiToken returns the ApiToken field value.
func (o *SecurityMonitoringIntegrationConfigOktaSecrets) GetApiToken() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApiToken
}

// GetApiTokenOk returns a tuple with the ApiToken field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigOktaSecrets) GetApiTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiToken, true
}

// SetApiToken sets field value.
func (o *SecurityMonitoringIntegrationConfigOktaSecrets) SetApiToken(v string) {
	o.ApiToken = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringIntegrationConfigOktaSecrets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["api_token"] = o.ApiToken

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringIntegrationConfigOktaSecrets) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiToken *string `json:"api_token"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApiToken == nil {
		return fmt.Errorf("required field api_token missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"api_token"})
	} else {
		return err
	}
	o.ApiToken = *all.ApiToken

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

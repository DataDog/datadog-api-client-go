// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringIntegrationConfigSentinelOneSecrets Credentials for a SentinelOne entity context sync.
type SecurityMonitoringIntegrationConfigSentinelOneSecrets struct {
	// The SentinelOne API token.
	ApiToken string `json:"api_token"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringIntegrationConfigSentinelOneSecrets instantiates a new SecurityMonitoringIntegrationConfigSentinelOneSecrets object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringIntegrationConfigSentinelOneSecrets(apiToken string) *SecurityMonitoringIntegrationConfigSentinelOneSecrets {
	this := SecurityMonitoringIntegrationConfigSentinelOneSecrets{}
	this.ApiToken = apiToken
	return &this
}

// NewSecurityMonitoringIntegrationConfigSentinelOneSecretsWithDefaults instantiates a new SecurityMonitoringIntegrationConfigSentinelOneSecrets object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringIntegrationConfigSentinelOneSecretsWithDefaults() *SecurityMonitoringIntegrationConfigSentinelOneSecrets {
	this := SecurityMonitoringIntegrationConfigSentinelOneSecrets{}
	return &this
}

// GetApiToken returns the ApiToken field value.
func (o *SecurityMonitoringIntegrationConfigSentinelOneSecrets) GetApiToken() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApiToken
}

// GetApiTokenOk returns a tuple with the ApiToken field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigSentinelOneSecrets) GetApiTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiToken, true
}

// SetApiToken sets field value.
func (o *SecurityMonitoringIntegrationConfigSentinelOneSecrets) SetApiToken(v string) {
	o.ApiToken = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringIntegrationConfigSentinelOneSecrets) MarshalJSON() ([]byte, error) {
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
func (o *SecurityMonitoringIntegrationConfigSentinelOneSecrets) UnmarshalJSON(bytes []byte) (err error) {
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
